CREATE TABLE IF NOT EXISTS CashFlow (
    symbol VARCHAR(10) NOT NULL, -- Ticker symbol for the company
    fiscalDateEnding DATE NOT NULL, -- Date of the annual report
    reportedCurrency TEXT NOT NULL, -- Currency in which the amounts are reported
    operatingCashflow BIGINT,
    paymentsForOperatingActivities BIGINT,
    proceedsFromOperatingActivities BIGINT,
    changeInOperatingLiabilities BIGINT,
    changeInOperatingAssets BIGINT,
    depreciationDepletionAndAmortization BIGINT,
    capitalExpenditures BIGINT,
    changeInReceivables BIGINT,
    changeInInventory BIGINT,
    profitLoss BIGINT,
    cashflowFromInvestment BIGINT,
    cashflowFromFinancing BIGINT,
    proceedsFromRepaymentsOfShortTermDebt BIGINT,
    paymentsForRepurchaseOfCommonStock BIGINT,
    paymentsForRepurchaseOfEquity BIGINT,
    paymentsForRepurchaseOfPreferredStock BIGINT,
    dividendPayout BIGINT,
    dividendPayoutCommonStock BIGINT,
    dividendPayoutPreferredStock BIGINT,
    proceedsFromIssuanceOfCommonStock BIGINT,
    proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet BIGINT,
    proceedsFromIssuanceOfPreferredStock BIGINT,
    proceedsFromRepurchaseOfEquity BIGINT,
    proceedsFromSaleOfTreasuryStock BIGINT,
    changeInCashAndCashEquivalents BIGINT,
    changeInExchangeRate BIGINT,
    netIncome BIGINT,
    PRIMARY KEY(symbol, fiscalDateEnding, reportedCurrency)
);

CREATE INDEX IF NOT EXISTS cashflowidx ON CashFlow(symbol, fiscalDateEnding);
