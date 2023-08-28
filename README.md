# StockDataSDK

- in early alpha
- will eventually collect and normalize data from different sources
- Currently only covers Alpha Vantage
- one of 3 sister projects


### Goal
- persist financial data to DB from different vendors (for vendor agnostic)
    - orchestrate calls to update data vis a vis vendor limits (different tiers)
- not for analytics, this is just for data collecting/cleaning
- not for free key abuse
- not for scraping yahoo finance etc.
- not for live data/updates, just end of day or weekly constructions for backtesting

Need:
- postgres
- testing
- decouple from alphavantage api, refactor with interfaces, which APIs satisfy
    - clarify API surface (similar to usage), add it here

Possible:
- output csv or json
- change fileout to single file or e.g. /ticker/daily.txt
- detect discongruities between sources
- reconsider use of floats? (imprecise, money etc. but most likely irrelevant here)

### implementation
- in Go
- use vendor APIs to collect data (not scraping sites)
- deployment via [MarketModel](https://github.com/veqqq/MarketModel)
    - containerized
    - persist data elsewhere

### alternatives
- existing go alpha vantage api - found wanting, basically just curl, little implemented
- julia api
- various bash scripts
- python stuff?
    - Easier to build scratch container agents with go than build whole linux instances to run python, julia or bash scripts

### system context
- StockDataSDK - data fetching and cleaning
- price to val comparer
- [CompanyModels](https://github.com/veqqq/CompanyModels), using financial statements
- [MarketModel](https://github.com/veqqq/MarketModel) - unified project gluing these together with workflow, build scripts etc. (builds DB to persist data)

### API

-----------

## Usage:

- the query builder normalizes input, so lower case and any order of arguments can be used. E.g. `yield 5yr` or `bond 5` will both be normalized to `TREASURY_YIELD 5year`.
- the internal representation require the "function" call first (e.g. an income statement, bond etc.)

- explanation of internal logic so far:
- "function" comes first. then ticker 
- need 3 args:
    - TOP_GAINERS_LOSERS
	- baseUrl + "FX_DAILY" + "&outputsize=full" + "&from_symbol=" + from + "&to_symbol=" + to
    - baseUrl + "TIME_SERIES_INTRADAY" + "&month" + tickerFirst + "&interval=1min" + "&symbol=" + tickerNext + "&outputsize=full"
        - needs a 2004-12 formated month

- need 2 args
    - baseUrl + "OVERVIEW" + "&symbol=" + tickerNext
    - baseUrl + "INCOME_STATEMENT" + "&symbol=" + tickerNext
    - baseUrl + "BALANCE_SHEET" + "&symbol=" + tickerNext
    - baseUrl + "CASH_FLOW" + "&symbol=" + tickerNext
    - baseUrl + "EARNINGS" + "&symbol=" + tickerNext
    - baseUrl + "TREASURY_YIELD" + "&interval=daily" + "&maturity=" + maturity
        - maturity is 3month, 2year, 5 year, 7year, 10year, 30year

- only needs a stock ticker?
    - baseUrl + "TIME_SERIES_DAILY" + "&symbol=" + ticker // + "&outputsize=full"

- 1 arg:
    - baseUrl + "WTI" + "&interval=daily"
    - baseUrl + "BRENT" + "&interval=daily"
    - baseUrl + "NATURAL_GAS" + "&interval=daily"
    - baseUrl + "COPPER" + "&interval=quarterly"
    - baseUrl + "ALUMINUM" + "&interval=quarterly"
    - baseUrl + "WHEAT" + "&interval=quarterly"
    - baseUrl + "CORN" + "&interval=quarterly"
    - baseUrl + "COTTON" + "&interval=quarterly"
    - baseUrl + "SUGAR" + "&interval=quarterly"
    - baseUrl + "COFFEE" + "&interval=quarterly"
    - baseUrl + "ALL_COMMODITIES" + "&interval=quarterly"
    - baseUrl + "REAL_GDP" + "&interval=quarterly"
    - baseUrl + "REAL_GDP_PER_CAPITA" + "&interval=quarterly"
    - baseUrl + "FEDERAL_FUNDS_RATE" + "&interval=daily"
    - baseUrl + "CPI" + "&interval=monthly"
    - baseUrl + "INFLATION"
    - baseUrl + "RETAIL_SALES"
    - baseUrl + "DURABLES"
    - baseUrl + "UNEMPLOYMENT"
    - baseUrl + "NONFARM_PAYROLL"


    ---------

### API Targets

- Alpha Vantage
- iexcloud
- financial modelling prep
- boursorama
- Tehran market testmc.com
- SEC's EDGAR for corporate filings
- stock shark - seems tedious for many things e.g. getCompanyFinancials

cf. https://rapidapi.com/category/Finance

