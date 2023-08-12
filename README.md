# StockDataSDK

- in early alpha
- will eventually collect and normalize data from different sources
- Currently only covers Alpha Vantage

#### Alpha Vantage

Need:
- save all ticker data in single file, instead of split by dates etc.?
- testing

Possible:
- output csv?
- output into posgres?
- make interfaces for structs, to enable marshalling?
  - or better?
- reformat functions to accept interfaces, which each API can satisfy

Not:
- for doing further analytics, this is just for data cleaning/accessing

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


### What to do with Fundemental Data?

- cf. alphaVantageAPISummary.md

- 3/4 main documents
    - balance sheet
    - income statement
    - cashflow statement