package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pgx "github.com/jackc/pgx/v4"
)

func main() {
	// Get environment variables
	dbUrl := os.Getenv("DB_URL")
	apiKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
	ctx := context.Background()
	conn, err := connectToDB(ctx, dbUrl)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	defer conn.Close(ctx)
	// Open file with symbols
	symbols, err := os.Open("symbols.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer symbols.Close()
	scanner := bufio.NewScanner(symbols)
	rateLimit := 1
	for scanner.Scan() {
		if rateLimit > 6 {
			rateLimit = 1
			time.Sleep(1 * time.Minute)
		}
		symbol := scanner.Text()
		fmt.Printf("Getting data for %s\n", symbol)
		getStockData(&ctx, conn, apiKey, symbol)
		fmt.Println("Total rows inserted for Daily Stocks: ", ctx.Value("DailyStocksRowsInserted"))
		fmt.Println("Total rows inserted for CompanyOverview: ", ctx.Value("CompanyOverviewRowsInserted"))
		fmt.Println("Total rows inserted for BalanceSheet: ", ctx.Value("BalanceSheetRowsInserted"))
		fmt.Println("Total rows inserted for IncomeStatement: ", ctx.Value("IncomeStatementRowsInserted"))
		fmt.Println("Total rows inserted for CashFlowReport: ", ctx.Value("CashFlowRowsInserted"))
		rateLimit++
	}

}

// getStockData fetches stock data for a given symbol, collecting the Company Overview, Balance Sheets,
// Income Statements, Cash Flow Statements, Annual Earnings Reports, and daily Stock Prices.
func getStockData(ctx *context.Context, connection *pgx.Conn, apiKey, symbol string) {
	dailyStocks := getDailyStockData(apiKey, "TIME_SERIES_DAILY", symbol)
	companyOverview := getCompanyOverview(apiKey, "OVERVIEW", symbol)
	balanceSheets := getBalanceSheet(apiKey, "BALANCE_SHEET", symbol)
	incomeStatements := getIncomeStatement(apiKey, "INCOME_STATEMENT", symbol)
	cashFlowReport := getCashFlowReport(apiKey, "CASH_FLOW", symbol)
	upsertDailyStock(ctx, connection, dailyStocks)
	upsertCompanyOverview(ctx, connection, companyOverview)
	upsertBalanceSheet(ctx, connection, balanceSheets)
	upsertIncomeStatement(ctx, connection, incomeStatements)
	upsertCashFlowReport(ctx, connection, cashFlowReport)
}
