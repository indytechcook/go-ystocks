// Package ystocks implements a library for using the Yahoo! Finance stock
// market API.
package ystocks

import (
    "encoding/csv"
    "net/http"
    "strings"
)

// Constants with URL parts for API URL-building.
const (
    quotesBaseUrl = "http://download.finance.yahoo.com/d/quotes.csv?s="
    staticUrl     = "&e=.csv"
)

// Stock symbol type; id must be a valid stock market symbol.
type Stock struct {
    id string
}

type StockProperties [][]string

// Named constants for various stock properties.
const (
    AfterHoursChangeRealtime    = "c8"
    AnnualizedGain              = "g3"
    Ask                         = "a0"
    AskRealtime                 = "b2"
    AskSize                     = "a5"
    AverageDailyVolume          = "a2"
    Bid                         = "b0"
    BidRealtime                 = "b3"
    BidSize                     = "b6"
    BookValuePerShare           = "b4"
    Change                      = "c1"
    ChangeChangeInPercent       = "c0"
    ChangeFromFiftyDayMovingAverage = "m7"
    ChangeFromTwoHundredDayMovingAverage = "m5"
    ChangeFromYearHigh          = "k4"
    ChangeFromYearLow           = "j5"
    ChangeInPercent             = "p2"
    ChangeInPercentRealtime     = "k2"
    ChangeRealtime              = "c6"
    Commission                  = "c3"
    Currency                    = "c4"
    DaysHigh                    = "h0"
    DaysLow                     = "g0"
    DaysRange                   = "m0"
    DaysRangeRealtime           = "m2"
    DaysValueChange             = "w1"
    DaysValueChangeRealtime     = "w4"
    DividendPayDate             = "r1"
    TrailingAnnualDividendYield = "d0"
    TrailingAnnualDividendYieldInPercent = "y0"
    DilutedEPS                  = "e0"
    EBITDA                      = "j4"
    EPSEstimateCurrentYear      = "e7"
    EPSEstimateNextQuarter      = "e9"
    EPSEstimateNextYear         = "e8"
    ExDividendDate              = "q0"
    FiftyDayMovingAverage       = "m3"
    SharesFloat                 = "f6"
    HighLimit                   = "l2"
    HoldingsGain                = "g4"
    HoldingsGainPercent         = "g1"
    HoldingsGainPercentRealtime = "g5"
    HoldingsGainRealtime        = "g6"
    HoldingsValue               = "v1"
    HoldingsValueRealtime       = "v7"
    LastTradeDate               = "d1"
    LastTradePriceOnly          = "l1"
    LastTradeRealtimeWithTime   = "k1"
    LastTradeSize               = "k3"
    LastTradeTime               = "t1"
    LastTradeWithTime           = "l0"
    LowLimit                    = "l3"
    MarketCapitalization        = "j1"
    MarketCapRealtime           = "j3"
    MoreInfo                    = "i0"
    Name                        = "n0"
    Notes                       = "n4"
    OneYearTargetPrice          = "t8"
    Open                        = "o0"
    OrderBookRealtime           = "i5"
    PEGRatio                    = "r5"
    PERatio                     = "r0"
    PERatioRealtime             = "r2"
    PercentChangeFromFiftyDayMovingAverage = "m8"
    PercentChangeFromTwoHundredDayMovingAverage = "m6"
    ChangeInPercentFromYearHigh = "k5"
    PercentChangeFromYearLow    = "j6"
    PreviousClose               = "p0"
    PriceBook                   = "p6"
    PriceEPSEstimateCurrentYear = "r6"
    PriceEPSEstimateNextYear    = "r7"
    PricePaid                   = "p1"
    PriceSales                  = "p5"
    Revenue                     = "s6"
    SharesOwned                 = "s1"
    SharesOutstanding           = "j2"
    ShortRatio                  = "s7"
    StockExchange               = "x0"
    Symbol                      = "s0"
    TickerTrend                 = "t7"
    TradeDate                   = "d2"
    TradeLinks                  = "t6"
    TradeLinksAdditional        = "f0"
    TwoHundredDayMovingAverage  = "m4"
    Volume                      = "v0"
    YearHigh                    = "k0"
    YearLow                     = "j0"
    YearRange                   = "w0"
)

func (s *Stock) getProperty(prop string) (string, error) {
    props, err := s.getProperties([]string{prop})

    // Flatten properties to a single string if no error was found
    if err == nil && len(props) == 1 && len(props[0]) == 1 {
        return props[0][0], nil
    }

    return "", err
}

func (s *Stock) getProperties(props []string) (StockProperties, error) {
    // Build up the Y! Finance API URL
    propsStr := strings.Join(props, "")
    url := quotesBaseUrl + s.id +
           "&f=" + propsStr +
           staticUrl

    // HTTP GET the CSV data
    resp, httpErr := http.Get(url)
    if httpErr != nil {
        return nil, httpErr
    }

    // Convert string CSV data to a usable data structure
    reader := csv.NewReader(resp.Body)
    records, parseErr := reader.ReadAll()
    if parseErr != nil {
        return nil, parseErr
    }

    return records, nil
}
