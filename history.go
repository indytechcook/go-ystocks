// Package ystocks implements a library for using the Yahoo! Finance stock
// market API.
package ystocks

import (
  "fmt"
  "encoding/csv"
  "net/http"
  "strconv"
  "time"
)

// Inverted index for months; given a Month it will return the integer
// 'index' of the month minus 1 (I don't know... Can't seem to find a way to
// go backwards from Month->int in the Go API. And even more I don't know...
// Yahoo! Finance API seems to take the month minus 1... Maybe I'm missing
// something?... Odd. Whatever...).
var monthsInverted = map[time.Month]int{
  time.January: 0,
  time.February: 1,
  time.March: 2,
  time.April: 3,
  time.May: 4,
  time.June: 5,
  time.July: 6,
  time.August: 7,
  time.September: 8,
  time.October: 9,
  time.November: 10,
  time.December: 11,
}

// Constants with URL parts for API call building.
const (
  historyBaseUrl = "http://ichart.yahoo.com/table.csv?s="
  fromMonth = "&a="
  fromDay = "&b="
  fromYear = "&c="
  toMonth = "&d="
  toDay = "&e="
  toYear = "&f"
  interval = "&g="
  endUrl = "&ignore=.csv"
)

// Time interval type, see interval constants later.
type TimeInterval string

// Interval constants for daily, monthly, and weekly stock histories (thanks to
// http://code.google.com/p/yahoo-finance-managed/wiki/enumHistQuotesInterval).
const (
  DailyInterval   = "d"
  WeeklyInterval  = "w"
  MonthlyInterval = "m"
)

// Stock history data type.
type StockHistory struct {
  from time.Time
  to time.Time
  interval TimeInterval
  prices []int
}

func (s *Stock) getHistory(from time.Time, to time.Time, i TimeInterval) (StockHistory, error) {
  url := historyBaseUrl + s.id +
         fromMonth + strconv.Itoa(monthsInverted[from.Month()]) +
         fromDay + strconv.Itoa(from.Day()) +
         fromYear + strconv.Itoa(from.Year()) +
         toMonth + strconv.Itoa(monthsInverted[to.Month()]) +
         toDay + strconv.Itoa(to.Day()) +
         toYear + strconv.Itoa(to.Year()) +
         interval + string(i) +
         endUrl

	// HTTP GET the CSV data
	resp, httpErr := http.Get(url)
  defer resp.Body.Close()
	if httpErr != nil {
		return nil, httpErr
	}

	// Convert string CSV data to a usable data structure
	reader := csv.NewReader(resp.Body)
	records, parseErr := reader.ReadAll()
	if parseErr != nil {
		return nil, parseErr
	}

  fmt.Println(records)

  // TODO
  prices := []int{0}
  hist := StockHistory{from, to, i, prices}
  return hist, nil
}
