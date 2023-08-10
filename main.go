package main

import (
	"StockDataSDK/APIs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

// accepts ewz overview, overview ewz etx.
// if not given with script invocation, will ask
// accepts multiple args and turns them into one string
// makes all caps, puts relevant thingslike bond, overview etc. first
func GetTickerFromUser() string {
	var userInput string

	if len(os.Args) < 2 {
		fmt.Println("Enter Stock Ticker. N.b. prefereds use a hyphen (PBR-A):")
		// reader := bufio.NewReader(os.Stdin)
		// ticker, err := reader.ReadString('\n')
		// check(err)
		fmt.Scanln(&userInput) // replaces need for reader!
		userInput = strings.TrimSuffix(userInput, "\n")
	} else {
		userInput = strings.Join(os.Args[1:], " ")
	}
	userInput = strings.ToUpper(userInput)

	// put overview, bond etc. as the first word before returning, so it can accept ewz overview, also.
	if strings.Contains(userInput, ".") {
		userInput = strings.ReplaceAll(userInput, ".", "-")
	}
	// looking for these, e.g. bond..
	words := []string{"OVERVIEW", "NONFARMPAYROLL", "NONFARM", "PAYROLL", "EMPLOYMENT", "TGLAT", "GAINERS", "LOSERS", "TOPGAINERSLOSERS",
		"OVERVIEW", "RETAIL", "INFLATION", "CPI", "FEDFUNDSRATE", "FUNDS", "EFFECTIVEFEDERALFUNDSRATE", "EFFR",
		"GDPPC", "GDPPERCAP", "GDP", "EARNINGS", "CASHFLOW", "BALANCE_SHEET", "BALANCE", "BALANCESHEET", "INCOME", "INCOMESTATEMENT", "RETAILSALES",
		"BOND", "YIELD", "TREASURY", "TREASURY_YIELD",
	}
	for _, word := range words {
		if strings.Contains(userInput, word) {
			userInput = strings.ReplaceAll(userInput, word, "")
			userInput = word + " " + userInput
		}
	}

	// Move the date to beginning
	words = strings.Fields(userInput)
	for _, word := range words {
		if regexp.MustCompile(`\b\d{4}-\d{2}\b`).MatchString(word) { // date in 2003-01 format
			userInput = strings.ReplaceAll(userInput, word, "")
			userInput = word + " " + userInput
		}
	}

	return userInput

}

// build baseUR, fetching the APIkey from env
// os.Getenv or github.com/joho/godotenv ?
// viper has a lot of dependencies...
//
//	func buildBaseURLViper() string {
//		viper.SetConfigFile("local.env")
//		err := viper.ReadInConfig()
//		check(err)
//		apiKey, ok := viper.Get("APIKEY").(string)
//		if !ok {
//			log.Fatalf("Add API Key to .env")
//		}
// apiKey = "?apikey=" + apiKey
// return "https://www.alphavantage.co//query" + apiKey + "&function="

//	}

// lazy implementation from godotenv to reduce dependencies
func buildBaseURL() string {
	f, err := (os.Open(".env"))
	check(err)
	defer f.Close()

	var envMap map[string]string
	err = json.NewDecoder(f).Decode(&envMap)
	check(err)

	currentEnv := map[string]bool{}
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentEnv[key] = true
	}
	for key, value := range envMap {
		if !currentEnv[key] {
			_ = os.Setenv(key, value)
		}
	}

	apiKey := os.Getenv("APIKEY")
	// apiKey, ok := os.LookupEnv("APIKEY")
	// if !ok {
	// 	log.Fatalf("Add API Key to .env")
	// }
	apiKey = "?apikey=" + apiKey
	return "https://www.alphavantage.co//query" + apiKey + "&function="

}

// example querry:
// https://www.alphavantage.co//query?function=TIME_SERIES_DAILY&symbol=EWZ&apikey= examplekey'
// ticker's a string with spaces, check the first word in the switch statement (e.g. overview ewz -> OVERVIEW EWZ)
// this frist word picks the func/querry type
func QueryBuilder(ticker string) (url string) {

	// the actual ticker comes after
	tickerFirst := strings.Fields(ticker)[0]

	var dateRegexIsTrue string
	if regexp.MustCompile(`\b\d{4}-\d{2}\b`).MatchString(tickerFirst) {
		dateRegexIsTrue = tickerFirst
	}

	switch tickerFirst {

	// News sentiment - complicated beast - figure out later

	// TOP_GAINTERS_LOSERS and most active...
	case "TGLAT", "TGLATS", "GAINERS", "LOSERS", "TOPGAINERSLOSERS":
		url = baseUrl + "APIs.TOP_GAINTERS_LOSERS"
		structType = "TGLATs"
		return
	// overview OVERVIEW
	case "OVERVIEW":
		url = baseUrl + "OVERVIEW" + "&symbol=" + ticker
		structType = "APIs.StockOverview"
		return
	// income INCOME_STATEMENT  // "EARNINGS", "CASHFLOW", "BALANCE", "BALANCESHEET", "INCOME", "INCOMESTATEMENT"
	// balance 	BALANCE_SHEET
	case "BALANCE_SHEET", "BALANCESHEET", "BALANCE":
		url = baseUrl + "BALANCE_SHEET" + "&symbol=" + ticker
		structType = "APIs.BalanceSheets"
		return
	// cashflow CASH_FLOW
	case "CASH_FLOW", "CASHFLOW":
		url = baseUrl + "CASH_FLOW" + "&symbol=" + ticker
		structType = "APIs.CashFlowStatements"
		return
	// earnings	EARNINGS
	case "EARNINGS":
		url = baseUrl + "EARNINGS" + "&symbol=" + ticker
		structType = "APIs.EarningsData"
		return

	// commodities and macro indicators use the same structs, but are funcs instead of ticker
	// WTI, BRENT, NATURAL_GAS, COPPER, ALUMINUM, WHEAT, CORN, COTTON, SUGAR, COFFEE, ALL_COMMODITIES
	// WTI
	case "WTI":
		url = baseUrl + "WTI" + "&interval=daily"
		structType = "APIs.CommodityPrices"
		return
	// BRENT
	case "BRENT":
		url = baseUrl + "BRENT" + "&interval=daily"
		structType = "APIs.CommodityPrices"
		return
	// nat gas
	case "NATURAL_GAS", "GAS": // check if GAS is a stock...
		url = baseUrl + "NATURAL_GAS" + "&interval=daily"
		structType = "APIs.CommodityPrices"
		return
	// COPPER
	case "COPPER":
		url = baseUrl + "COPPER" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// ALUMINUM
	case "ALUMINUM":
		url = baseUrl + "ALUMINUM" + "&interval=quarterly"
		structType = "CommodityPrices"
		return
	// WHEAT
	case "WHEAT":
		url = baseUrl + "WHEAT" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// CORN
	case "CORN":
		url = baseUrl + "CORN" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// COTTON
	case "COTTON":
		url = baseUrl + "COTTON" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// SUGAR
	case "SUGAR":
		url = baseUrl + "SUGAR" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// COFFEE
	case "COFFEE":
		url = baseUrl + "COFFEE" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// ALL_COMMODITIES
	case "ALL_COMMODITIES":
		url = baseUrl + "ALL_COMMODITIES" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// REAL_GDP
	case "GDP":
		url = baseUrl + "REAL_GDP" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	// real gdp per cap
	case "GDPPC", "GDPPERCAP":
		url = baseUrl + "REAL_GDP_PER_CAPITA" + "&interval=quarterly"
		structType = "APIs.CommodityPrices"
		return
	case "FEDFUNDSRATE", "FEDFUNDS", "FUNDS", "EFFECTIVEFEDERALFUNDSRATE", "EFFR":
		url = baseUrl + "FEDERAL_FUNDS_RATE" + "&interval=daily"
		structType = "APIs.CommodityPrices"
		return
	case "CPI":
		url = baseUrl + "CPI" + "&interval=monthly"
		structType = "APIs.CommodityPrices"
		return
	// inflation
	case "INFLATION":
		url = baseUrl + "INFLATION"
		structType = "APIs.CommodityPrices"
		return
	// retail sales - RETAIL_SALES
	case "RETAILSALES", "RETAIL":
		url = baseUrl + "RETAIL_SALES"
		structType = "APIs.CommodityPrices"
		return
		// durable goods orders - DURABLES
	case "DURABLES":
		url = baseUrl + "DURABLES"
		structType = "APIs.CommodityPrices"
		return
	// unemployment - UNEMPLOYMENT
	case "UNEMPLOYMENT":
		url = baseUrl + "UNEMPLOYMENT"
		structType = "APIs.CommodityPrices"
		return
	// nonfarm payroll
	case "NONFARMPAYROLL", "NONFARM", "PAYROLL", "EMPLOYMENT":
		url = baseUrl + "NONFARM_PAYROLL"
		structType = "APIs.CommodityPrices"
		return

		// treasury yield - in percent, monthly or daily?
		//          // maturities: 3month, 2year, 5year, 7year, 10year
		// TREASURY_YIELD  &maturity 3month, 2year,5,year,7year, 10 year, 30year
	case "BOND", "YIELD", "TREASURY", "TREASURY_YIELD":
		// assuming they type e.g. "bond 3"
		maturity := strings.Fields(ticker)[1]
		switch maturity {
		case "3", "3m", "3month":
			maturity = "3month"
		case "2", "2y", "2yr", "2year":
			maturity = "2year"
		case "5", "5y", "5yr", "5year":
			maturity = "5year"
		case "7", "7y", "7yr", "7year":
			maturity = "7year"
		case "10", "10y", "10yr", "10year":
			maturity = "10year"
		case "30", "30y", "30yr", "30year":
			maturity = "30year"
		}

		url = baseUrl + "TREASURY_YIELD" + "&interval=daily" + "&maturity=" + maturity
		structType = "APIs.CommodityPrices"
		return

	// time series intraday   	 ?interval=1min  extended true/false?
	// e.g. month=2009-01, since 2000-01
	case dateRegexIsTrue: // looks for 2001-01 format

		// Parse date to check if it's before "2000-01".
		date, err := time.Parse("2006-01", tickerFirst)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}
		refDate, _ := time.Parse("2006-01", "2000-01")
		if date.Before(refDate) {
			fmt.Println("Error: Date is before 2000-01")
		}
		// strongly expect this to fail, it doesn't add the month...
		url = baseUrl + "TIME_SERIES_INTRADAY" + "&symbol" + ticker + "&outputsize=full"
		structType = "APIs.IntradayOHLCVs"
		return

	// daily time series, DailyOHLCVs
	// &outputsize=full gets 20 years of data, remove it when testing defaulting to compact with 100 data points...
	default:
		url = baseUrl + "TIME_SERIES_DAILY" + "&symbol=" + ticker // + "&outputsize=full"
		structType = "APIs.DailyOHLCVs"
		return

	}
}

// unmarshal into correct structs, because of the different nesting forms, can't avoid it
// to marshal into better format, use huge switch statement, which checks the global var structType
// is this (9/func) less loc than the requisite interface satisfying?

// what if I don't remarshal?

func ReformatJson(resp io.Reader) string {
	// declare encoder
	decoder := json.NewDecoder(resp)

	// the switch checks global var structType, then uses it as the marshaling struct type
	switch structType {
	case "APIs.TGLATs":
		var seriesDataMap APIs.TGLATs
		err := decoder.Decode(&seriesDataMap)
		check(err)
		output, err := json.Marshal(seriesDataMap) // perhaps change, but 3 maps
		check(err)
		return string(output)
	case "APIs.StockOverview":
		var seriesDataMap APIs.StockOverview
		err := decoder.Decode(&seriesDataMap)
		check(err)
		output, err := json.Marshal(seriesDataMap)
		check(err)
		return string(output)
	case "APIs.BalanceSheets":
		var seriesDataMap APIs.BalanceSheets
		err := decoder.Decode(&seriesDataMap.QuarterlyReports)
		check(err)
		output, err := json.Marshal(seriesDataMap.QuarterlyReports)
		check(err)
		return string(output)
	case "APIs.CashFlowStatements":
		var seriesDataMap APIs.CashFlowStatements
		err := decoder.Decode(&seriesDataMap)
		check(err)
		output, err := json.Marshal(seriesDataMap.QuarterlyReports)
		check(err)
		return string(output)
	case "APIs.EarningsData":
		var seriesDataMap APIs.EarningsData
		err := decoder.Decode(&seriesDataMap)
		check(err)
		output, err := json.Marshal(seriesDataMap.QuarterlyEarnings)
		check(err)
		return string(output)
	// Commodities and Economic Indicators - use same structure
	// WTI, BRENT, nat gas, COPPER, ALUMINUM, WHEAT, CORN, COTTON, SUGAR, COFFEE
	case "APIs.CommodityPrices":
		var seriesDataMap APIs.CommodityPrices
		err := decoder.Decode(&seriesDataMap)
		check(err)
		output, err := json.Marshal(seriesDataMap)
		check(err)
		return string(output)

	case "APIs.IntradayOHLCVs":
		var seriesDataMap APIs.IntradayOHLCVs
		err := decoder.Decode(&seriesDataMap)
		check(err)
		output, err := json.Marshal(seriesDataMap.TimeSeries5min)
		check(err)
		return string(output)
	case "APIs.DailyOHLCVs":
		var seriesDataMap APIs.DailyOHLCVs
		err := decoder.Decode(&seriesDataMap)
		check(err)
		// what if I don't remarshal?
		// fmt.Println(seriesDataMap)
		// return fmt.Sprint(seriesDataMap.TimeSeries)
		output, err := json.Marshal(seriesDataMap.TimeSeries)
		check(err)
		return string(output)
	default: // why do i need this? wont trigger, hm
		panic("confident I don't need this")

	}
}

// original
// func ReformatJson(body io.Reader, structType string) string {
// trying without structs, since I don't know how to declare the structType in accordance to the querry
// 	// use structType to pick structs
// 	// might need to make funcs to marshall and unmarshal for other programs later

// 	// structType will be the var instead of DailyOHLCVs, picking it...
// 	// declare seriesDataMap in switch cases!
// 	// but there will be scope issues, so predeclare them globally
// 	// or more elegantly, use interfaces...
// 	err := json.NewDecoder(body).Decode(&seriesDataMap)
// 	check(err)

// 	// isolate just the time data
// 	jsondata, err := json.Marshal(seriesdatamap.TimeSeries) // TimeSeries must be modular too...
// 	check(err)
// 	return string(jsondata) // return a pointer? these are big items...
// }

func WriteToFile(filename, data string) {
	// if the file doesn't exist, make it, using ticker name (incl day/full querry)
	// n.b. will duplicate data if file exists - remove append or
	// have the sql builder have aditional logic?
	// Move the date to end
	words := strings.Fields(filename)
	for _, word := range words {
		if regexp.MustCompile(`\b\d{4}-\d{2}\b`).MatchString(word) { // date in 2003-01 format
			filename = strings.ReplaceAll(filename, word, "")
			filename = strings.ReplaceAll(filename, "  ", " ")
			filename = filename[1:] + " " + word // [1:] to remove a space, or: word = word + " "
		}
	}

	filename = strings.ReplaceAll(filename, " ", "_") // remove spaces from file name
	f, err := os.OpenFile("data/"+filename+".txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	check(err)
	defer f.Close()
	fmt.Fprint(f, data)
}

var baseUrl string    // global
var structType string // global, is specified in query builder, then used it reformat Json

func init() {
	baseUrl = buildBaseURL()
}

func main() {
	ticker := GetTickerFromUser()
	url := QueryBuilder(ticker)

	resp, err := http.Get(url) // later become new func?
	check(err)
	defer resp.Body.Close()
	fmt.Println(url)

	fmt.Println(resp.Body)
	finalData := ReformatJson(resp.Body)
	fmt.Println(finalData) // for dev purposes, remove when everything works
	WriteToFile(ticker, finalData)

}
