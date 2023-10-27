package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func connectToDB(ctx context.Context, url string) (*pgxpool.Pool, error) {
	return pgxpool.Connect(ctx, url)
}

// nullableString takes in a string that represents an empty string (in our case, either "" or "None")
// and returns a string that can be interpreted as null by the database. This is used as a wrapper around
// string objects that could potentially be nullable.
func nullableString(val string) sql.NullString {
	if val == "" || val == "None" || val == "-" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: val,
		Valid:  true,
	}
}

func upsertDailyStock(ctx *context.Context, connection *pgxpool.Pool, dailyStock DailyStockResponse) {
	query := `INSERT INTO DailyStocks 
			(Symbol, fiscalDate, Open, High, Low, Close, Volume) 
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			ON CONFLICT (Symbol, fiscalDate) DO NOTHING;
			`
	symbol := dailyStock.Metadata.Symbol
	rowsInserted := 0
	for date, reading := range dailyStock.Readings {
		command, err := connection.Exec(
			*ctx,
			query,
			symbol, date, reading.Open, reading.High,
			reading.Low, reading.Close, reading.Volume)
		rowsInserted += int(command.RowsAffected())
		if err != nil {
			fmt.Println("Error encountered in upsertDailyStock:")
			fmt.Println(err)
			continue
		}
	}
	*ctx = context.WithValue(*ctx, "DailyStocksRowsInserted", rowsInserted)
}

func upsertCompanyOverview(ctx *context.Context, connection *pgxpool.Pool, data CompanyOverview) {
	insertStatement := `INSERT INTO CompanyOverview (
		symbol, assetType, name, description, cik, exchange, currency, country, 
		sector, industry, address, fiscalYearEnd, latestQuarter, marketCapitalization, 
		ebitda, peRatio, pegRatio, bookValue, dividendPerShare, dividendYield, 
		eps, revenuePerShareTTM, profitMargin, operatingMarginTTM, returnOnAssetsTTM, 
		returnOnEquityTTM, revenueTTM, grossProfitTTM, dilutedEPSTTM, quarterlyEarningsGrowthYOY, 
		quarterlyRevenueGrowthYOY, analystTargetPrice, trailingPE, forwardPE, priceToSalesRatioTTM, 
		priceToBookRatio, evToRevenue, evToEBITDA, beta, _52WeekHigh, _52WeekLow, 
		_50DayMovingAverage, _200DayMovingAverage, sharesOutstanding, dividendDate, exDividendDate
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, 
		$11, $12, $13, $14, $15, $16, $17, $18, $19, $20, 
		$21, $22, $23, $24, $25, $26, $27, $28, $29, $30, 
		$31, $32, $33, $34, $35, $36, $37, $38, $39, $40, 
		$41, $42, $43, $44, $45, $46
	) ON CONFLICT (symbol) DO NOTHING;`
	command, err := connection.Exec(
		*ctx,
		insertStatement,
		data.Symbol, data.AssetType, data.Name, data.Description,
		data.CIK, data.Exchange, data.Currency, data.Country, data.Sector,
		data.Industry, data.Address, data.FiscalYearEnd, nullableString(data.LatestQuarter),
		nullableString(data.MarketCapitalization), nullableString(data.EBITDA), nullableString(data.PERatio), nullableString(data.PEGRatio),
		nullableString(data.BookValue), nullableString(data.DividendPerShare), nullableString(data.DividendYield), nullableString(data.EPS),
		nullableString(data.RevenuePerShareTTM), nullableString(data.ProfitMargin), nullableString(data.OperatingMarginTTM), nullableString(data.ReturnOnAssetsTTM),
		nullableString(data.ReturnOnEquityTTM), nullableString(data.RevenueTTM), nullableString(data.GrossProfitTTM), nullableString(data.DilutedEPSTTM),
		nullableString(data.QuarterlyEarningsGrowthYOY), nullableString(data.QuarterlyRevenueGrowthYOY), nullableString(data.AnalystTargetPrice), nullableString(data.TrailingPE),
		nullableString(data.ForwardPE), nullableString(data.PriceToSalesRatioTTM), nullableString(data.PriceToBookRatio), nullableString(data.EVToRevenue),
		nullableString(data.EVToEBITDA), nullableString(data.Beta), nullableString(data.Fifty2WeekHigh), nullableString(data.Fifty2WeekLow), nullableString(data.FiftyDayMovingAverage),
		nullableString(data.TwoHundredDayMovingAverage), nullableString(data.SharesOutstanding), nullableString(data.DividendDate),
		nullableString(data.ExDividendDate))

	if err != nil {
		fmt.Println("Error encountered in upsertCompanyOverview:")
		fmt.Println(err)
	}

	*ctx = context.WithValue(*ctx, "CompanyOverviewRowsInserted", int(command.RowsAffected()))
}

func upsertBalanceSheet(ctx *context.Context, connection *pgxpool.Pool, data BalanceSheet) {
	insertStatement := `INSERT INTO BalanceSheet (
		symbol, fiscalDateEnding, reportedCurrency, totalAssets, totalCurrentAssets, 
		cashAndCashEquivalentsAtCarryingValue, cashAndShortTermInvestments, inventory, 
		currentNetReceivables, totalNonCurrentAssets, propertyPlantEquipment, 
		accumulatedDepreciationAmortizationPPE, intangibleAssets, intangibleAssetsExcludingGoodwill, 
		goodwill, investments, longTermInvestments, shortTermInvestments, otherCurrentAssets, 
		otherNonCurrentAssets, totalLiabilities, totalCurrentLiabilities, currentAccountsPayable, 
		deferredRevenue, currentDebt, shortTermDebt, totalNonCurrentLiabilities, capitalLeaseObligations, 
		longTermDebt, currentLongTermDebt, longTermDebtNoncurrent, shortLongTermDebtTotal, 
		otherCurrentLiabilities, otherNonCurrentLiabilities, totalShareholderEquity, treasuryStock, 
		retainedEarnings, commonStock, commonStockSharesOutstanding
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, 
		$11, $12, $13, $14, $15, $16, $17, $18, $19, $20, 
		$21, $22, $23, $24, $25, $26, $27, $28, $29, $30, 
		$31, $32, $33, $34, $35, $36, $37, $38, $39
	) ON CONFLICT (symbol, fiscalDateEnding, reportedCurrency) DO NOTHING;`
	rowsInserted := 0
	for _, balanceSheet := range data.AnnualBalanceSheets {
		command, err := connection.Exec(
			*ctx,
			insertStatement,
			data.Symbol, nullableString(balanceSheet.FiscalDateEnding), balanceSheet.ReportedCurrency, balanceSheet.TotalAssets,
			balanceSheet.TotalCurrentAssets, balanceSheet.CashAndCashEquivalentsAtCarryingValue, balanceSheet.CashAndShortTermInvestments, balanceSheet.Inventory,
			balanceSheet.CurrentNetReceivables, balanceSheet.TotalNonCurrentAssets, balanceSheet.PropertyPlantEquipment, balanceSheet.AccumulatedDepreciationAmortizationPPE,
			balanceSheet.IntangibleAssets, balanceSheet.IntangibleAssetsExcludingGoodwill, balanceSheet.Goodwill, balanceSheet.Investments,
			balanceSheet.LongTermInvestments, balanceSheet.ShortTermInvestments, balanceSheet.OtherCurrentAssets, balanceSheet.OtherNonCurrentAssets, balanceSheet.TotalLiabilities,
			balanceSheet.TotalCurrentLiabilities, balanceSheet.CurrentAccountsPayable, balanceSheet.DeferredRevenue, balanceSheet.CurrentDebt,
			balanceSheet.ShortTermDebt, balanceSheet.TotalNonCurrentLiabilities, balanceSheet.CapitalLeaseObligations, balanceSheet.LongTermDebt,
			balanceSheet.CurrentLongTermDebt, balanceSheet.LongTermDebtNoncurrent, balanceSheet.ShortLongTermDebtTotal, balanceSheet.OtherCurrentLiabilities,
			balanceSheet.OtherNonCurrentLiabilities, balanceSheet.TotalShareholderEquity, balanceSheet.TreasuryStock, balanceSheet.RetainedEarnings,
			balanceSheet.CommonStock, balanceSheet.CommonStockSharesOutstanding)
		if err != nil {
			fmt.Println("Error encountered in upsertBalanceSheet")
			fmt.Println(err)
			continue
		}
		rowsInserted += int(command.RowsAffected())
	}
	*ctx = context.WithValue(*ctx, "BalanceSheetRowsInserted", rowsInserted)
}

func upsertIncomeStatement(ctx *context.Context, connection *pgxpool.Pool, data IncomeStatement) {
	insertStatement := `INSERT INTO IncomeStatement (
		symbol, fiscalDateEnding, reportedCurrency, grossProfit, totalRevenue, 
		costOfRevenue, costofGoodsAndServicesSold, operatingIncome, 
		sellingGeneralAndAdministrative, researchAndDevelopment, operatingExpenses, 
		investmentIncomeNet, netInterestIncome, interestIncome, interestExpense, 
		nonInterestIncome, otherNonOperatingIncome, depreciation, 
		depreciationAndAmortization, incomeBeforeTax, incomeTaxExpense, 
		interestAndDebtExpense, netIncomeFromContinuingOperations, 
		comprehensiveIncomeNetOfTax, ebit, ebitda, netIncome
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, 
		$11, $12, $13, $14, $15, $16, $17, $18, $19, $20, 
		$21, $22, $23, $24, $25, $26, $27
	) ON CONFLICT(symbol, fiscalDateEnding, reportedCurrency) DO NOTHING;`
	rowsInserted := 0
	for _, incomeStatement := range data.AnnualIncomeReports {
		command, err := connection.Exec(
			*ctx,
			insertStatement,
			data.Symbol,
			nullableString(incomeStatement.FiscalDateEnding), incomeStatement.ReportedCurrency, incomeStatement.GrossProfit,
			incomeStatement.TotalRevenue, incomeStatement.CostOfRevenue, incomeStatement.CostofGoodsAndServicesSold,
			incomeStatement.OperatingIncome, incomeStatement.SellingGeneralAndAdministrative, incomeStatement.ResearchAndDevelopment,
			incomeStatement.OperatingExpenses, incomeStatement.InvestmentIncomeNet, incomeStatement.NetInterestIncome, incomeStatement.InterestIncome,
			incomeStatement.InterestExpense, incomeStatement.NonInterestIncome, incomeStatement.OtherNonOperatingIncome, incomeStatement.Depreciation,
			incomeStatement.DepreciationAndAmortization, incomeStatement.IncomeBeforeTax, incomeStatement.IncomeTaxExpense, incomeStatement.InterestAndDebtExpense,
			incomeStatement.NetIncomeFromContinuingOperations, incomeStatement.ComprehensiveIncomeNetOfTax, incomeStatement.Ebit, incomeStatement.Ebitda, incomeStatement.NetIncome)
		if err != nil {
			fmt.Println("Error encountered in upsertIncomeStatement:")
			fmt.Println(err)
			continue
		}
		rowsInserted += int(command.RowsAffected())
	}
	*ctx = context.WithValue(*ctx, "IncomeStatementRowsInserted", rowsInserted)
}

func upsertCashFlowReport(ctx *context.Context, connection *pgxpool.Pool, data CashflowReport) {
	insertStatement := `INSERT INTO CashFlow (
		symbol, fiscalDateEnding, reportedCurrency, operatingCashflow, 
		paymentsForOperatingActivities, proceedsFromOperatingActivities, 
		changeInOperatingLiabilities, changeInOperatingAssets, 
		depreciationDepletionAndAmortization, capitalExpenditures, 
		changeInReceivables, changeInInventory, profitLoss, 
		cashflowFromInvestment, cashflowFromFinancing, 
		proceedsFromRepaymentsOfShortTermDebt, paymentsForRepurchaseOfCommonStock, 
		paymentsForRepurchaseOfEquity, paymentsForRepurchaseOfPreferredStock, 
		dividendPayout, dividendPayoutCommonStock, dividendPayoutPreferredStock, 
		proceedsFromIssuanceOfCommonStock, 
		proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet, 
		proceedsFromIssuanceOfPreferredStock, proceedsFromRepurchaseOfEquity, 
		proceedsFromSaleOfTreasuryStock, changeInCashAndCashEquivalents, 
		changeInExchangeRate, netIncome
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, 
		$11, $12, $13, $14, $15, $16, $17, $18, $19, $20, 
		$21, $22, $23, $24, $25, $26, $27, $28, $29, $30
	) ON CONFLICT (symbol, fiscalDateEnding, reportedCurrency) DO NOTHING;`
	rowsInserted := 0
	for _, cashFlowReport := range data.AnnualReports {
		command, err := connection.Exec(
			*ctx,
			insertStatement,
			data.Symbol, nullableString(cashFlowReport.FiscalDateEnding), cashFlowReport.ReportedCurrency, cashFlowReport.OperatingCashflow,
			cashFlowReport.PaymentsForOperatingActivities, cashFlowReport.ProceedsFromOperatingActivities, cashFlowReport.ChangeInOperatingLiabilities,
			cashFlowReport.ChangeInOperatingAssets, cashFlowReport.DepreciationDepletionAndAmortization, cashFlowReport.CapitalExpenditures,
			cashFlowReport.ChangeInReceivables, cashFlowReport.ChangeInInventory, cashFlowReport.ProfitLoss, cashFlowReport.CashflowFromInvestment,
			cashFlowReport.CashflowFromFinancing, cashFlowReport.ProceedsFromRepaymentsOfShortTermDebt, cashFlowReport.PaymentsForRepurchaseOfCommonStock,
			cashFlowReport.PaymentsForRepurchaseOfEquity, cashFlowReport.PaymentsForRepurchaseOfPreferredStock, cashFlowReport.DividendPayout,
			cashFlowReport.DividendPayoutCommonStock, cashFlowReport.DividendPayoutPreferredStock, cashFlowReport.ProceedsFromIssuanceOfCommonStock,
			cashFlowReport.ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet, cashFlowReport.ProceedsFromIssuanceOfPreferredStock,
			cashFlowReport.ProceedsFromRepurchaseOfEquity, cashFlowReport.ProceedsFromSaleOfTreasuryStock, cashFlowReport.ChangeInCashAndCashEquivalents,
			cashFlowReport.ChangeInExchangeRate, cashFlowReport.NetIncome)
		if err != nil {
			fmt.Println("Error encountered in upsertCashFlowReport:")
			fmt.Println(err)
			continue
		}
		rowsInserted += int(command.RowsAffected())
	}
	*ctx = context.WithValue(*ctx, "CashFlowRowsInserted", rowsInserted)
}
