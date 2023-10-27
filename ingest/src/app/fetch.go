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

func getApiResponseBody(url, apiKey, function, symbol string) ([]byte, error) {
	// Make request
	request, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(request.Body)
	return body, err
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getDailyStockData(compact bool, apiKey, function, symbol string) DailyStockResponse {
	var outputSize string
	if compact {
		outputSize = "compact"
	} else {
		outputSize = "full"
	}
	url := buildUrl(apiKey, function, symbol) + "&outputsize=" + outputSize
	body, err := getApiResponseBody(url, apiKey, function, symbol)
	handleError(err)
	var response DailyStockResponse
	json.Unmarshal(body, &response)
	return response
}

func getCompanyOverview(apiKey, function, symbol string) CompanyOverview {
	url := buildUrl(apiKey, function, symbol)
	body, err := getApiResponseBody(url, apiKey, function, symbol)
	handleError(err)
	var response CompanyOverview
	json.Unmarshal(body, &response)
	return response
}

func getBalanceSheet(apiKey, function, symbol string) BalanceSheet {
	url := buildUrl(apiKey, function, symbol)
	body, err := getApiResponseBody(url, apiKey, function, symbol)
	handleError(err)
	var response BalanceSheet
	json.Unmarshal(body, &response)
	return response
}

func getIncomeStatement(apiKey, function, symbol string) IncomeStatement {
	url := buildUrl(apiKey, function, symbol)
	body, err := getApiResponseBody(url, apiKey, function, symbol)
	handleError(err)
	var response IncomeStatement
	json.Unmarshal(body, &response)
	return response
}

func getCashFlowReport(apiKey, function, symbol string) CashflowReport {
	url := buildUrl(apiKey, function, symbol)
	body, err := getApiResponseBody(url, apiKey, function, symbol)
	handleError(err)
	var response CashflowReport
	json.Unmarshal(body, &response)
	return response
}
