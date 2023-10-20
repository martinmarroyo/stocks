package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// get API key
	apiKey := os.Getenv("ALPHA_VANTAGE_API_KEY") // 7GQ6FTK7V0XHYM0H
	symbol := "IBM"
	stock := getDailyStockData(apiKey, "TIME_SERIES_DAILY", symbol)
	// companyOverview, _ := getApiResponseBody(apiKey, "OVERVIEW", symbol)
	// fmt.Println(companyOverview.Symbol, companyOverview.Name, companyOverview.Currency, companyOverview.Exchange)
	// time.Sleep(1 * time.Second)
	fmt.Println("Date | \tOpen | \tClose\n")
	for date, reading := range stock.Readings {
		fmt.Printf("%s\t%s\t%s\n", date, reading.Open, reading.Close)
		time.Sleep(1 * time.Second)
	}
	// test, err := json.Marshal(companyOverview)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(test))
	// fmt.Println(string(companyOverview))
}
