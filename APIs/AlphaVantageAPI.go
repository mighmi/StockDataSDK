package APIs

import (
	"encoding/json"
	"strconv"
)

// AlphaVantage API https://www.alphavantage.co/documentation/

///////////////////////
// Stock Ticker Data:

// TIME_SERIES_DAILY
type DailyOHLCVs struct {
	MetaData   DailyOHLCVMetaData    `json:"Meta Data"`
	TimeSeries map[string]DailyOHLCV `json:"Time Series (Daily)"`
}

// reformat stuff for this?
type DailyOHLCVMetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

type DailyOHLCV struct {
	Open   float64 `json:"1. open,string"`
	High   float64 `json:"2. high,string"`
	Low    float64 `json:"3. low,string"`
	Close  float64 `json:"4. close,string"`
	Volume int64   `json:"5. volume,string"`
}

// TIME_SERIES_INTRADAY
type IntradayOHLCVs struct {
	MetaData       IntradayOHLCVMetaData    `json:"Meta Data"`
	TimeSeries1min map[string]IntradayOHLCV `json:"Time Series (1min)"`
}

type IntradayOHLCVMetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	Interval      string `json:"4. Interval"`
	OutputSize    string `json:"5. Output Size"`
	TimeZone      string `json:"6. Time Zone"`
}

type IntradayOHLCV struct {
	Open   float64 `json:"1. open,string"`
	High   float64 `json:"2. high,string"`
	Low    float64 `json:"3. low,string"`
	Close  float64 `json:"4. close,string"`
	Volume int64   `json:"5. volume,string"`
}

//////////////////
// Stock Fundementals:
// The 4 financial statements + "overview"

// overview
type StockOverview struct {
	Symbol                     string  `json:"Symbol"`
	AssetType                  string  `json:"AssetType"`
	Name                       string  `json:"Name"`
	Description                string  `json:"Description"`
	CIK                        string  `json:"CIK"`
	Exchange                   string  `json:"Exchange"`
	Currency                   string  `json:"Currency"`
	Country                    string  `json:"Country"`
	Sector                     string  `json:"Sector"`
	Industry                   string  `json:"Industry"`
	Address                    string  `json:"Address"`
	FiscalYearEnd              string  `json:"FiscalYearEnd"`
	LatestQuarter              string  `json:"LatestQuarter"`
	MarketCapitalization       float64 `json:"MarketCapitalization,string"`
	EBITDA                     float64 `json:"EBITDA,string"`
	PERatio                    float64 `json:"PERatio,string"`
	PEGRatio                   float64 `json:"PEGRatio,string"`
	BookValue                  float64 `json:"BookValue,string"`
	DividendPerShare           float64 `json:"DividendPerShare,string"`
	DividendYield              float64 `json:"DividendYield,string"`
	EPS                        float64 `json:"EPS,string"`
	RevenuePerShareTTM         float64 `json:"RevenuePerShareTTM,string"`
	ProfitMargin               float64 `json:"ProfitMargin,string"`
	OperatingMarginTTM         float64 `json:"OperatingMarginTTM,string"`
	ReturnOnAssetsTTM          float64 `json:"ReturnOnAssetsTTM,string"`
	ReturnOnEquityTTM          float64 `json:"ReturnOnEquityTTM,string"`
	RevenueTTM                 float64 `json:"RevenueTTM,string"`
	GrossProfitTTM             float64 `json:"GrossProfitTTM,string"`
	DilutedEPSTTM              float64 `json:"DilutedEPSTTM,string"`
	QuarterlyEarningsGrowthYOY float64 `json:"QuarterlyEarningsGrowthYOY,string"`
	QuarterlyRevenueGrowthYOY  float64 `json:"QuarterlyRevenueGrowthYOY,string"`
	AnalystTargetPrice         float64 `json:"AnalystTargetPrice,string"`
	TrailingPE                 float64 `json:"TrailingPE,string"`
	ForwardPE                  float64 `json:"ForwardPE,string"`
	PriceToSalesRatioTTM       float64 `json:"PriceToSalesRatioTTM,string"`
	PriceToBookRatio           float64 `json:"PriceToBookRatio,string"`
	EVToRevenue                float64 `json:"EVToRevenue,string"`
	EVToEBITDA                 float64 `json:"EVToEBITDA,string"`
	Beta                       float64 `json:"Beta,string"`
	WeekHigh                   float64 `json:"52WeekHigh,string"`
	WeekLow                    float64 `json:"52WeekLow,string"`
	DayMovingAverage50         float64 `json:"50DayMovingAverage,string"`
	DayMovingAverage200        float64 `json:"200DayMovingAverage,string"`
	SharesOutstanding          float64 `json:"SharesOutstanding,string"`
	DividendDate               string  `json:"DividendDate"`
	ExDividendDate             string  `json:"ExDividendDate"`
}

// income statement
// json: invalid use of ,string struct tag, trying to unmarshal "None" into float64
// unsure why, so unmarshaling into map any

type IncomeStatements struct {
	Symbol           string            `json:"symbol"`
	AnnualReports    []IncomeStatement `json:"annualReports"`
	QuarterlyReports []IncomeStatement `json:"quarterlyReports"`
}

type IncomeStatement struct {
	FiscalDateEnding                  string  `json:"fiscalDateEnding"`
	ReportedCurrency                  string  `json:"reportedCurrency"`
	GrossProfit                       float64 `json:"grossProfit,string"`
	TotalRevenue                      float64 `json:"totalRevenue,string"`
	CostOfRevenue                     float64 `json:"costOfRevenue,string"`
	CostofGoodsAndServicesSold        float64 `json:"costofGoodsAndServicesSold,string"`
	OperatingIncome                   float64 `json:"operatingIncome,string"`
	SellingGeneralAndAdministrative   float64 `json:"sellingGeneralAndAdministrative,string"`
	ResearchAndDevelopment            float64 `json:"researchAndDevelopment,string"`
	OperatingExpenses                 float64 `json:"operatingExpenses,string"`
	InvestmentIncomeNet               float64 `json:"investmentIncomeNet,string"`
	NetInterestIncome                 float64 `json:"netInterestIncome,string"`
	InterestIncome                    float64 `json:"interestIncome,string"`
	InterestExpense                   float64 `json:"interestExpense,string"`
	NonInterestIncome                 float64 `json:"nonInterestIncome,string"`
	OtherNonOperatingIncome           float64 `json:"otherNonOperatingIncome,string"`
	Depreciation                      float64 `json:"depreciation,string"`
	DepreciationAndAmortization       float64 `json:"depreciationAndAmortization,string"`
	IncomeBeforeTax                   float64 `json:"incomeBeforeTax,string"`
	IncomeTaxExpense                  float64 `json:"incomeTaxExpense,string"`
	InterestAndDebtExpense            float64 `json:"interestAndDebtExpense,string"`
	NetIncomeFromContinuingOperations float64 `json:"netIncomeFromContinuingOperations,string"`
	ComprehensiveIncomeNetOfTax       float64 `json:"comprehensiveIncomeNetOfTax,string"`
	EBIT                              float64 `json:"ebit,string"`
	EBITDA                            float64 `json:"ebitda,string"`
	NetIncome                         float64 `json:"netIncome,string"`
}

// https://github.com/veqqq/StockDataSDK/issues/6
// implement UnmarshalJSON to avoid this error:
// json: invalid use of ,string struct tag, trying to unmarshal "None" into float64
// should add every "None" possible field to aux and give it a similar if statement below
//
// turn nullable fields into *float64, for clf that's at least InvestmentIncomeNet and ResearchAndDeveloplemtnt
//
// func (f *IncomeStatement) UnmarshalJSON(data []byte) error {
// 	type Alias IncomeStatement // Alias avoids recursive call to UnmarshalJSON
// 	aux := &struct {
// 		*Alias
// 		ResearchAndDevelopment json.RawMessage `json:"researchAndDevelopment"`
// 		InvestmentIncomeNet    json.RawMessage `json:"investmentIncomeNet"`
// 	}{Alias: (*Alias)(f)}

// 	if err := json.Unmarshal(data, &aux); err != nil {
// 		return err
// 	}

// 	if aux.ResearchAndDevelopment != nil {
// 		var value float64
// 		if err := json.Unmarshal(aux.ResearchAndDevelopment, &value); err != nil {
// 			if string(aux.ResearchAndDevelopment) == `"None"` {
// 				f.ResearchAndDevelopment = nil
// 			} else {
// 				return err
// 			}
// 		} else {
// 			f.ResearchAndDevelopment = &value
// 		}
// 	}

// 	if aux.InvestmentIncomeNet != nil {
// 		var value float64
// 		if err := json.Unmarshal(aux.InvestmentIncomeNet, &value); err != nil {
// 			if string(aux.InvestmentIncomeNet) == `"None"` {
// 				f.InvestmentIncomeNet = nil
// 			} else {
// 				return err
// 			}
// 		} else {
// 			f.InvestmentIncomeNet = &value
// 		}
// 	}

// 	return nil
// }

// balance sheet

type BalanceSheets struct {
	Symbol           string         `json:"symbol"`
	AnnualReports    []BalanceSheet `json:"annualReports"`
	QuarterlyReports []BalanceSheet `json:"quarterlyReports"`
}

type BalanceSheet struct {
	FiscalDateEnding                       string  `json:"fiscalDateEnding"`
	ReportedCurrency                       string  `json:"reportedCurrency"`
	TotalAssets                            float64 `json:"totalAssets,string"`
	TotalCurrentAssets                     float64 `json:"totalCurrentAssets,string"`
	CashAndCashEquivalentsAtCarryingValue  float64 `json:"cashAndCashEquivalentsAtCarryingValue,string"`
	CashAndShortTermInvestments            float64 `json:"cashAndShortTermInvestments,string"`
	Inventory                              float64 `json:"inventory,string"`
	CurrentNetReceivables                  float64 `json:"currentNetReceivables,string"`
	TotalNonCurrentAssets                  float64 `json:"totalNonCurrentAssets,string"`
	PropertyPlantEquipment                 float64 `json:"propertyPlantEquipment,string"`
	AccumulatedDepreciationAmortizationPPE float64 `json:"accumulatedDepreciationAmortizationPPE,string"`
	IntangibleAssets                       float64 `json:"intangibleAssets,string"`
	IntangibleAssetsExcludingGoodwill      float64 `json:"intangibleAssetsExcludingGoodwill,string"`
	Goodwill                               float64 `json:"goodwill,string"`
	Investments                            float64 `json:"investments,string"`
	LongTermInvestments                    float64 `json:"longTermInvestments,string"`
	ShortTermInvestments                   float64 `json:"shortTermInvestments,string"`
	OtherCurrentAssets                     float64 `json:"otherCurrentAssets,string"`
	OtherNonCurrentAssets                  float64 `json:"otherNonCurrentAssets,string"`
	TotalLiabilities                       float64 `json:"totalLiabilities,string"`
	TotalCurrentLiabilities                float64 `json:"totalCurrentLiabilities,string"`
	CurrentAccountsPayable                 float64 `json:"currentAccountsPayable,string"`
	DeferredRevenue                        float64 `json:"deferredRevenue,string"`
	CurrentDebt                            float64 `json:"currentDebt,string"`
	ShortTermDebt                          float64 `json:"shortTermDebt,string"`
	TotalNonCurrentLiabilities             float64 `json:"totalNonCurrentLiabilities,string"`
	CapitalLeaseObligations                float64 `json:"capitalLeaseObligations,string"`
	LongTermDebt                           float64 `json:"longTermDebt,string"`
	CurrentLongTermDebt                    float64 `json:"currentLongTermDebt,string"`
	LongTermDebtNoncurrent                 float64 `json:"longTermDebtNoncurrent,string"`
	ShortLongTermDebtTotal                 float64 `json:"shortLongTermDebtTotal,string"`
	OtherCurrentLiabilities                float64 `json:"otherCurrentLiabilities,string"`
	OtherNonCurrentLiabilities             float64 `json:"otherNonCurrentLiabilities,string"`
	TotalShareholderEquity                 float64 `json:"totalShareholderEquity,string"`
	TreasuryStock                          float64 `json:"treasuryStock,string"`
	RetainedEarnings                       float64 `json:"retainedEarnings,string"`
	CommonStock                            float64 `json:"commonStock,string"`
	CommonStockSharesOutstanding           float64 `json:"commonStockSharesOutstanding,string"`
}

// Cash flow

type CashFlowStatements struct {
	Symbol           string              `json:"symbol"`
	AnnualReports    []CashFlowStatement `json:"annualReports"`
	QuarterlyReports []CashFlowStatement `json:"quarterlyReports"`
}

type CashFlowStatement struct {
	FiscalDateEnding                                          string  `json:"fiscalDateEnding"`
	ReportedCurrency                                          string  `json:"reportedCurrency"`
	OperatingCashflow                                         float64 `json:"operatingCashflow,string"`
	PaymentsForOperatingActivities                            float64 `json:"paymentsForOperatingActivities,string"`
	ProceedsFromOperatingActivities                           float64 `json:"proceedsFromOperatingActivities,string"`
	ChangeInOperatingLiabilities                              float64 `json:"changeInOperatingLiabilities,string"`
	ChangeInOperatingAssets                                   float64 `json:"changeInOperatingAssets,string"`
	DepreciationDepletionAndAmortization                      float64 `json:"depreciationDepletionAndAmortization,string"`
	CapitalExpenditures                                       float64 `json:"capitalExpenditures,string"`
	ChangeInReceivables                                       float64 `json:"changeInReceivables,string"`
	ChangeInInventory                                         float64 `json:"changeInInventory,string"`
	ProfitLoss                                                float64 `json:"profitLoss,string"`
	CashflowFromInvestment                                    float64 `json:"cashflowFromInvestment,string"`
	CashflowFromFinancing                                     float64 `json:"cashflowFromFinancing,string"`
	ProceedsFromRepaymentsOfShortTermDebt                     float64 `json:"proceedsFromRepaymentsOfShortTermDebt,string"`
	PaymentsForRepurchaseOfCommonStock                        float64 `json:"paymentsForRepurchaseOfCommonStock,string"`
	PaymentsForRepurchaseOfEquity                             float64 `json:"paymentsForRepurchaseOfEquity,string"`
	PaymentsForRepurchaseOfPreferredStock                     float64 `json:"paymentsForRepurchaseOfPreferredStock,string"`
	DividendPayout                                            float64 `json:"dividendPayout,string"`
	DividendPayoutCommonStock                                 float64 `json:"dividendPayoutCommonStock,string"`
	DividendPayoutPreferredStock                              float64 `json:"dividendPayoutPreferredStock,string"`
	ProceedsFromIssuanceOfCommonStock                         float64 `json:"proceedsFromIssuanceOfCommonStock,string"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet float64 `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet,string"`
	ProceedsFromIssuanceOfPreferredStock                      float64 `json:"proceedsFromIssuanceOfPreferredStock,string"`
	ProceedsFromRepurchaseOfEquity                            float64 `json:"proceedsFromRepurchaseOfEquity,string"`
	ProceedsFromSaleOfTreasuryStock                           float64 `json:"proceedsFromSaleOfTreasuryStock,string"`
	ChangeInCashAndCashEquivalents                            float64 `json:"changeInCashAndCashEquivalents,string"`
	ChangeInExchangeRate                                      float64 `json:"changeInExchangeRate,string"`
	NetIncome                                                 float64 `json:"netIncome,string"`
}

// Earnings
type EarningsData struct {
	Symbol            string              `json:"symbol"`
	AnnualEarnings    []AnnualEarnings    `json:"annualEarnings"`
	QuarterlyEarnings []QuarterlyEarnings `json:"quarterlyEarnings"`
}

type AnnualEarnings struct { // lol pointless
	FiscalDateEnding string  `json:"fiscalDateEnding"`
	ReportedEPS      float64 `json:"reportedEPS,string"`
}

type QuarterlyEarnings struct {
	FiscalDateEnding   string  `json:"fiscalDateEnding"`
	ReportedDate       string  `json:"reportedDate"`
	ReportedEPS        float64 `json:"reportedEPS,string"`
	EstimatedEPS       float64 `json:"estimatedEPS,string"`
	Surprise           float64 `json:"surprise,string"`
	SurprisePercentage float64 `json:"surprisePercentage,string"`
}

///////

///////
// Commodities and Economic Indicators - use same structure

// WTI, BRENT, nat gas, COPPER, ALUMINUM, WHEAT, CORN, COTTON, SUGAR, COFFEE

// REAL_GDP in billions of dollars - same structure as commodities...
// real gdp per cap - in "chained 2012 dollars"
// fed funds rate - in percent - daily? monthly?
// cpi - "index 1982-1984=100" - monthly
// inflation - in percent, only annual
// retail sales - in millions, only monthly
// durable goods orders - in millions, only monthly
// unemployment - in percent, only monthly
// nonfarm payroll - in thousands of people, only monthly
// treasury yield - in percent, monthly or daily?
//          // maturities: 3month, 2year, 5year, 7year, 10year

type CommodityPrices struct { // rename
	Name     string           `json:"name"`     // e.g. global price of copper or henry hub...
	Interval string           `json:"interval"` // should be daily always
	Unit     string           `json:"unit"`
	Data     []CommodityPrice `json:"data"`
}

type CommodityPrice struct { // rename
	Date  string  `json:"date,omitempty"`
	Value float64 `json:"value,string,omitempty"`
}

// CommodityPrice satisfies Unmarshaler
func (c *CommodityPrice) UnmarshalJSON(data []byte) error {
	var aux struct {
		Date  string `json:"date"`
		Value string `json:"value"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	// If val or date are useless, ignore struct
	// some dates have "." val, which normal unmarshaler doesn't handle, hence writing this
	if aux.Value == "null" || aux.Value == "" || aux.Value == "." || aux.Value == "0" || aux.Value == "0.0" || aux.Date == "" || aux.Date == "null" || aux.Date == "." {
		return nil
	}
	value, err := strconv.ParseFloat(aux.Value, 64)
	if err != nil { // check declared in main.go, inaccessible here
		return err
	}
	c.Date = aux.Date
	c.Value = value
	return nil
}

/////////
//// Other

// exchange rates, daily

type ForexPrices struct {
	MetaData     ForexMetaData         `json:"Meta Data"`
	TimeSeriesFX map[string]ForexPrice `json:"Time Series FX (Daily)"` // monthly is the same, just uses (Monthly here)
}

type ForexMetaData struct {
	Information   string `json:"1. Information"` // "1. Information": "Forex Daily Prices (open, high, low, close)",
	FromSymbol    string `json:"2. From Symbol"` // "2. From Symbol": "EUR",
	ToSymbol      string `json:"3. To Symbol"`   // "3. To Symbol": "USD",
	OutputSize    string `json:"4. Output Size"`
	LastRefreshed string `json:"5. Last Refreshed"`
	TimeZone      string `json:"6. Time Zone"`
}

type ForexPrice struct {
	Open  float64 `json:"1. open,string"` // "1. open": "1.10020",
	High  float64 `json:"2. high,string"`
	Low   float64 `json:"3. low,string"`
	Close float64 `json:"4. close,string"`
}

// 20 TOP_GAINERS_LOSERS  and most actively traded

type TGLATs struct {
	Metadata     string  `json:"metadata"`
	LastUpdated  string  `json:"last_updated"`
	TopGainers   []TGLAT `json:"top_gainers"`
	TopLosers    []TGLAT `json:"top_losers"`
	MostActively []TGLAT `json:"most_actively_traded"`
}

type TGLAT struct {
	Ticker           string  `json:"ticker"`
	Price            float64 `json:"price,string"`
	ChangeAmount     float64 `json:"change_amount,string"`
	ChangePercentage string  `json:"change_percentage"` // has %, later implement json unmarshaler for this
	Volume           int64   `json:"volume,string"`
}

// News sentiment - complicated beast

type SentimentData struct {
	Items                    string     `json:"items"`
	SentimentScoreDefinition string     `json:"sentiment_score_definition"`
	RelevanceScoreDefinition string     `json:"relevance_score_definition"`
	Feed                     []FeedData `json:"feed"`
}

type FeedData struct {
	Title                 string                `json:"title"`
	URL                   string                `json:"url"`
	TimePublished         string                `json:"time_published"` // "20230805T122000",
	Authors               []string              `json:"authors"`
	Summary               string                `json:"summary"`
	BannerImage           string                `json:"banner_image"`           // lol, for a spam farm?
	Source                string                `json:"source"`                 // "CNBC",
	CategoryWithinSource  string                `json:"category_within_source"` // "Top News",
	SourceDomain          string                `json:"source_domain"`          // "www.cnbc.com",
	Topics                []TopicData           `json:"topics"`
	OverallSentimentScore float64               `json:"overall_sentiment_score"` //	0.072673,
	OverallSentimentLabel string                `json:"overall_sentiment_label"` //	"Neutral",
	TickerSentiment       []TickerSentimentData `json:"ticker_sentiment"`
}

type TopicData struct {
	Topic          string  `json:"topic"`                  // "Technology",
	RelevanceScore float64 `json:"relevance_score,string"` // "1.0"
}

type TickerSentimentData struct {
	Ticker               string  `json:"ticker"`
	RelevanceScore       float64 `json:"relevance_score,string"`        // "0.699089",
	TickerSentimentScore float64 `json:"ticker_sentiment_score,string"` // "0.116531",
	TickerSentimentLabel string  `json:"ticker_sentiment_label"`        //  "Neutral"
}

// var sentimentData SentimentData
// if err == json.Unmarshal([]byte(jsonData), &sentimentData); err != nil {}
// 	fmt.Println("Error unmarshaling JSON:", err)
// 	return
// }
