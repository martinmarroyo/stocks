package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	// Get environment variables
	dbUrl := os.Getenv("DB_URL")
	apiKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
	compact := flag.Bool("compact", true, "Flag that indicates whether to get first 100 data points or full history")
	flag.Parse()
	if *compact == true {
		fmt.Println("Payload size: Last 100 data points")
	} else {
		fmt.Println("Payload size: Last 20+ years historical")
	}
	channel := make(chan bool)
	var wg sync.WaitGroup
	ctx := context.Background()
	conn, err := connectToDB(ctx, dbUrl)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	defer conn.Close()
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
		wg.Add(1)
		symbol := scanner.Text()
		fmt.Printf("Getting data for %s\n", symbol)
		go getStockData(&ctx, channel, &wg, conn, *compact, apiKey, symbol)
		fmt.Println("Total rows inserted for Daily Stocks: ", ctx.Value("DailyStocksRowsInserted"))
		fmt.Println("Total rows inserted for CompanyOverview: ", ctx.Value("CompanyOverviewRowsInserted"))
		fmt.Println("Total rows inserted for BalanceSheet: ", ctx.Value("BalanceSheetRowsInserted"))
		fmt.Println("Total rows inserted for IncomeStatement: ", ctx.Value("IncomeStatementRowsInserted"))
		fmt.Println("Total rows inserted for CashFlowReport: ", ctx.Value("CashFlowRowsInserted"))
		rateLimit++
	}
	go func() {
		wg.Wait()
		close(channel)
	}()
}

// getStockData fetches stock data for a given symbol, collecting the Company Overview, Balance Sheets,
// Income Statements, Cash Flow Statements, Annual Earnings Reports, and daily Stock Prices.
func getStockData(ctx *context.Context, channel chan bool, wg *sync.WaitGroup, connection *pgxpool.Pool, compact bool, apiKey, symbol string) {
	defer wg.Done()
	dailyStocks := getDailyStockData(compact, apiKey, "TIME_SERIES_DAILY", symbol)
	companyOverview := getCompanyOverview(apiKey, "OVERVIEW", symbol)
	balanceSheets := getBalanceSheet(apiKey, "BALANCE_SHEET", symbol)
	incomeStatements := getIncomeStatement(apiKey, "INCOME_STATEMENT", symbol)
	cashFlowReport := getCashFlowReport(apiKey, "CASH_FLOW", symbol)
	upsertDailyStock(ctx, connection, dailyStocks)
	upsertCompanyOverview(ctx, connection, companyOverview)
	upsertBalanceSheet(ctx, connection, balanceSheets)
	upsertIncomeStatement(ctx, connection, incomeStatements)
	upsertCashFlowReport(ctx, connection, cashFlowReport)
	channel <- true
}
