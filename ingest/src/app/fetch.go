package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func buildUrl(apiKey, function, symbol string) string {
	baseUrl := "https://www.alphavantage.co/query?"
	funcParam := "function=" + function + "&"
	symbolParam := "symbol=" + symbol + "&"
	apiKeyParam := "apikey=" + apiKey
	finalUrl := baseUrl + funcParam + symbolParam + apiKeyParam
	return finalUrl
}

func getApiResponseBody(apiKey, function, symbol string) ([]byte, error) {
	// Build url
	url := buildUrl(apiKey, function, symbol)
	// Make request
	request, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(request.Body)
	return body, err
}

func getDailyStockData(apiKey, function, symbol string) DailyStockResponse {
	body, err := getApiResponseBody(apiKey, function, symbol)
	if err != nil {
		log.Fatal(err)
	}
	var response DailyStockResponse
	json.Unmarshal(body, &response)
	return response
}

func getCompanyOverview(apiKey, function, symbol string) CompanyOverview {
	body, err := getApiResponseBody(apiKey, function, symbol)
	if err != nil {
		log.Fatal(err)
	}
	var response CompanyOverview
	json.Unmarshal(body, &response)
	return response
}

func getBalanceSheet(apiKey, function, symbol string) BalanceSheet {
	body, err := getApiResponseBody(apiKey, function, symbol)
	if err != nil {
		log.Fatal(err)
	}
	var response BalanceSheet
	json.Unmarshal(body, &response)
	return response
}

func getIncomeStatement(apiKey, function, symbol string) IncomeStatement {
	body, err := getApiResponseBody(apiKey, function, symbol)
	if err != nil {
		log.Fatal(err)
	}
	var response IncomeStatement
	json.Unmarshal(body, &response)
	return response
}

func getCashFlowReport(apiKey, function, symbol string) CashflowReport {
	body, err := getApiResponseBody(apiKey, function, symbol)
	if err != nil {
		log.Fatal(err)
	}
	var response CashflowReport
	json.Unmarshal(body, &response)
	return response
}
