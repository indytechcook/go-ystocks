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

const (
    AfterHoursChangeRealtime = "c8"
    AnnualizedGain           = "g3"
    Ask                      = "a0"
    // ...TODO...
    Symbol                   = "s0"
    // ...TODO...
    Volume                   = "v0"
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
