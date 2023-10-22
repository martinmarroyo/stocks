package main

// DailyStock
type DailyStockResponse struct {
	Metadata Metadata                `json:"Meta Data"`
	Readings map[string]DailyReading `json:"Time Series (Daily)"`
}

type Metadata struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

type DailyReading struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

// CompanyOverview
type CompanyOverview struct {
	Symbol                     string `json:"Symbol"`
	AssetType                  string `json:"AssetType"`
	Name                       string `json:"Name"`
	Description                string `json:"Description"`
	CIK                        string `json:"CIK"`
	Exchange                   string `json:"Exchange"`
	Currency                   string `json:"Currency"`
	Country                    string `json:"Country"`
	Sector                     string `json:"Sector"`
	Industry                   string `json:"Industry"`
	Address                    string `json:"Address"`
	FiscalYearEnd              string `json:"FiscalYearEnd"`
	LatestQuarter              string `json:"LatestQuarter"`
	MarketCapitalization       string `json:"MarketCapitalization"`
	EBITDA                     string `json:"EBITDA"`
	PERatio                    string `json:"PERatio"`
	PEGRatio                   string `json:"PEGRatio"`
	BookValue                  string `json:"BookValue"`
	DividendPerShare           string `json:"DividendPerShare"`
	DividendYield              string `json:"DividendYield"`
	EPS                        string `json:"EPS"`
	RevenuePerShareTTM         string `json:"RevenuePerShareTTM"`
	ProfitMargin               string `json:"ProfitMargin"`
	OperatingMarginTTM         string `json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM          string `json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM          string `json:"ReturnOnEquityTTM"`
	RevenueTTM                 string `json:"RevenueTTM"`
	GrossProfitTTM             string `json:"GrossProfitTTM"`
	DilutedEPSTTM              string `json:"DilutedEPSTTM"`
	QuarterlyEarningsGrowthYOY string `json:"QuarterlyEarningsGrowthYOY"`
	QuarterlyRevenueGrowthYOY  string `json:"QuarterlyRevenueGrowthYOY"`
	AnalystTargetPrice         string `json:"AnalystTargetPrice"`
	TrailingPE                 string `json:"TrailingPE"`
	ForwardPE                  string `json:"ForwardPE"`
	PriceToSalesRatioTTM       string `json:"PriceToSalesRatioTTM"`
	PriceToBookRatio           string `json:"PriceToBookRatio"`
	EVToRevenue                string `json:"EVToRevenue"`
	EVToEBITDA                 string `json:"EVToEBITDA"`
	Beta                       string `json:"Beta"`
	Fifty2WeekHigh             string `json:"52WeekHigh"`
	Fifty2WeekLow              string `json:"52WeekLow"`
	FiftyDayMovingAverage      string `json:"50DayMovingAverage"`
	TwoHundredDayMovingAverage string `json:"200DayMovingAverage"`
	SharesOutstanding          string `json:"SharesOutstanding"`
	DividendDate               string `json:"DividendDate"`
	ExDividendDate             string `json:"ExDividendDate"`
}

// BalanceSheet
type BalanceSheet struct {
	Symbol              string               `json:"symbol"`
	AnnualBalanceSheets []AnnualBalanceSheet `json:"annualReports"`
}

type AnnualBalanceSheet struct {
	FiscalDateEnding                       string `json:"fiscalDateEnding"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	TotalAssets                            int64  `json:"totalAssets,string"`
	TotalCurrentAssets                     int64  `json:"totalCurrentAssets,string"`
	CashAndCashEquivalentsAtCarryingValue  int64  `json:"cashAndCashEquivalentsAtCarryingValue,string"`
	CashAndShortTermInvestments            int64  `json:"cashAndShortTermInvestments,string"`
	Inventory                              int64  `json:"inventory,string"`
	CurrentNetReceivables                  int64  `json:"currentNetReceivables,string"`
	TotalNonCurrentAssets                  int64  `json:"totalNonCurrentAssets,string"`
	PropertyPlantEquipment                 int64  `json:"propertyPlantEquipment,string"`
	AccumulatedDepreciationAmortizationPPE int64  `json:"accumulatedDepreciationAmortizationPPE,string"`
	IntangibleAssets                       int64  `json:"intangibleAssets,string"`
	IntangibleAssetsExcludingGoodwill      int64  `json:"intangibleAssetsExcludingGoodwill,string"`
	Goodwill                               int64  `json:"goodwill,string"`
	Investments                            int64  `json:"investments"`
	LongTermInvestments                    int64  `json:"longTermInvestments,string"`
	ShortTermInvestments                   int64  `json:"shortTermInvestments,string"`
	OtherCurrentAssets                     int64  `json:"otherCurrentAssets,string"`
	OtherNonCurrentAssets                  int64  `json:"otherNonCurrentAssets"`
	TotalLiabilities                       int64  `json:"totalLiabilities,string"`
	TotalCurrentLiabilities                int64  `json:"totalCurrentLiabilities,string"`
	CurrentAccountsPayable                 int64  `json:"currentAccountsPayable,string"`
	DeferredRevenue                        int64  `json:"deferredRevenue,string"`
	CurrentDebt                            int64  `json:"currentDebt,string"`
	ShortTermDebt                          int64  `json:"shortTermDebt,string"`
	TotalNonCurrentLiabilities             int64  `json:"totalNonCurrentLiabilities,string"`
	CapitalLeaseObligations                int64  `json:"capitalLeaseObligations,string"`
	LongTermDebt                           int64  `json:"longTermDebt,string"`
	CurrentLongTermDebt                    int64  `json:"currentLongTermDebt,string"`
	LongTermDebtNoncurrent                 int64  `json:"longTermDebtNoncurrent,string"`
	ShortLongTermDebtTotal                 int64  `json:"shortLongTermDebtTotal,string"`
	OtherCurrentLiabilities                int64  `json:"otherCurrentLiabilities,string"`
	OtherNonCurrentLiabilities             int64  `json:"otherNonCurrentLiabilities,string"`
	TotalShareholderEquity                 int64  `json:"totalShareholderEquity,string"`
	TreasuryStock                          int64  `json:"treasuryStock,string"`
	RetainedEarnings                       int64  `json:"retainedEarnings,string"`
	CommonStock                            int64  `json:"commonStock,string"`
	CommonStockSharesOutstanding           int64  `json:"commonStockSharesOutstanding,string"`
}

// IncomeStatement
type IncomeStatement struct {
	Symbol              string               `json:"symbol"`
	AnnualIncomeReports []AnnualIncomeReport `json:"annualReports"`
}

type AnnualIncomeReport struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding"`
	ReportedCurrency                  string `json:"reportedCurrency"`
	GrossProfit                       int64  `json:"grossProfit,string"`
	TotalRevenue                      int64  `json:"totalRevenue,string"`
	CostOfRevenue                     int64  `json:"costOfRevenue,string"`
	CostofGoodsAndServicesSold        int64  `json:"costofGoodsAndServicesSold,string"`
	OperatingIncome                   int64  `json:"operatingIncome,string"`
	SellingGeneralAndAdministrative   int64  `json:"sellingGeneralAndAdministrative,string"`
	ResearchAndDevelopment            int64  `json:"researchAndDevelopment,string"`
	OperatingExpenses                 int64  `json:"operatingExpenses,string"`
	InvestmentIncomeNet               int64  `json:"investmentIncomeNet"`
	NetInterestIncome                 int64  `json:"netInterestIncome,string"`
	InterestIncome                    int64  `json:"interestIncome,string"`
	InterestExpense                   int64  `json:"interestExpense,string"`
	NonInterestIncome                 int64  `json:"nonInterestIncome,string"`
	OtherNonOperatingIncome           int64  `json:"otherNonOperatingIncome,string"`
	Depreciation                      int64  `json:"depreciation,string"`
	DepreciationAndAmortization       int64  `json:"depreciationAndAmortization,string"`
	IncomeBeforeTax                   int64  `json:"incomeBeforeTax,string"`
	IncomeTaxExpense                  int64  `json:"incomeTaxExpense,string"`
	InterestAndDebtExpense            int64  `json:"interestAndDebtExpense,string"`
	NetIncomeFromContinuingOperations int64  `json:"netIncomeFromContinuingOperations,string"`
	ComprehensiveIncomeNetOfTax       int64  `json:"comprehensiveIncomeNetOfTax,string"`
	Ebit                              int64  `json:"ebit,string"`
	Ebitda                            int64  `json:"ebitda,string"`
	NetIncome                         int64  `json:"netIncome,string"`
}

// CashFlows
type CashflowReport struct {
	Symbol        string         `json:"symbol"`
	AnnualReports []CashflowData `json:"annualReports"`
}

type CashflowData struct {
	FiscalDateEnding                                          string `json:"fiscalDateEnding"`
	ReportedCurrency                                          string `json:"reportedCurrency"`
	OperatingCashflow                                         int64  `json:"operatingCashflow,string"`
	PaymentsForOperatingActivities                            int64  `json:"paymentsForOperatingActivities,string"`
	ProceedsFromOperatingActivities                           int64  `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              int64  `json:"changeInOperatingLiabilities,string"`
	ChangeInOperatingAssets                                   int64  `json:"changeInOperatingAssets,string"`
	DepreciationDepletionAndAmortization                      int64  `json:"depreciationDepletionAndAmortization,string"`
	CapitalExpenditures                                       int64  `json:"capitalExpenditures,string"`
	ChangeInReceivables                                       int64  `json:"changeInReceivables,string"`
	ChangeInInventory                                         int64  `json:"changeInInventory,string"`
	ProfitLoss                                                int64  `json:"profitLoss,string"`
	CashflowFromInvestment                                    int64  `json:"cashflowFromInvestment,string"`
	CashflowFromFinancing                                     int64  `json:"cashflowFromFinancing,string"`
	ProceedsFromRepaymentsOfShortTermDebt                     int64  `json:"proceedsFromRepaymentsOfShortTermDebt,string"`
	PaymentsForRepurchaseOfCommonStock                        int64  `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             int64  `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     int64  `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            int64  `json:"dividendPayout,string"`
	DividendPayoutCommonStock                                 int64  `json:"dividendPayoutCommonStock,string"`
	DividendPayoutPreferredStock                              int64  `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         int64  `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet int64  `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet,string"`
	ProceedsFromIssuanceOfPreferredStock                      int64  `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            int64  `json:"proceedsFromRepurchaseOfEquity,string"`
	ProceedsFromSaleOfTreasuryStock                           int64  `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents                            int64  `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                                      int64  `json:"changeInExchangeRate"`
	NetIncome                                                 int64  `json:"netIncome,string"`
}
