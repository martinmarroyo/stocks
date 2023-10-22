CREATE TABLE IF NOT EXISTS DailyStocks (
    Symbol TEXT,
    fiscalDate DATE,
    Open DOUBLE PRECISION,
    High DOUBLE PRECISION,
    Low DOUBLE PRECISION,
    Close DOUBLE PRECISION,
    Volume INT,
    PRIMARY KEY (Symbol, fiscalDate)
);

CREATE INDEX IF NOT EXISTS dailystocksidx ON DailyStocks(Symbol, fiscalDate);