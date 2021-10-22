package connpass

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// Param is a function which set a value to url.Values.
type Param func(vals url.Values) error

// EventID sets value to url.Values with key "event_id".
// eventID must be positive integer.
func EventID(eventID int) Param {
	return func(vals url.Values) error {
		if eventID < 0 {
			return fmt.Errorf("invalid event id: %d", eventID)
		}
		vals.Add("event_id", strconv.Itoa(eventID))
		return nil
	}
}

// Keyword sets value to url.Values with key "keyword".
// keyword must not be empty string.
func Keyword(keyword string) Param {
	return func(vals url.Values) error {
		if keyword == "" {
			return errors.New("empty keyword")
		}
		vals.Add("keyword", keyword)
		return nil
	}
}

// KeywordOr sets value to url.Values with key "keyword_or".
// keyword must not be empty string.
func KeywordOr(keyword string) Param {
	return func(vals url.Values) error {
		if keyword == "" {
			return errors.New("empty keyword")
		}
		vals.Add("keyword_or", keyword)
		return nil
	}
}

// YearMonth sets value to url.Values with key "ym".
// year must be between 0 and 9999 and month must be between time.January and time.December.
func YearMonth(year int, month time.Month) Param {
	return func(vals url.Values) error {
		switch {
		case year < 0 || year > 9999:
			return fmt.Errorf("year must be between 0 and 9999: %d", year)
		case month < time.January || month > time.December:
			return fmt.Errorf("invalid month: %d", month)
		}
		vals.Add("ym", fmt.Sprintf("%04d%02d", year, month))
		return nil
	}
}

// YearMonth sets value to url.Values with key "ymd".
// year must be between 0 and 9999 and month must be between time.January and time.December.
// day must be valid day at the corresponded month.
func YearMonthDay(year int, month time.Month, day int) Param {
	return func(vals url.Values) error {
		switch {
		case year < 0 || year > 9999:
			return fmt.Errorf("year must be between 0 and 9999: %d", year)
		case month < time.January || month > time.December:
			return fmt.Errorf("invalid month: %d", month)
		}

		tm := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		if tm.Month() != month {
			return fmt.Errorf("invalid day: %d", day)
		}

		vals.Add("ymd", fmt.Sprintf("%04d%02d%02d", year, month, day))
		return nil
	}
}

// Nickname sets value to url.Values with key "nickname".
// nickname must not be empty string.
func Nickname(nickname string) Param {
	return func(vals url.Values) error {
		if nickname == "" {
			return fmt.Errorf("empty nickname")
		}
		vals.Add("nickname", nickname)
		return nil
	}
}

// OwnerNickname sets value to url.Values with key "owner_nickname".
// nickname must not be empty string.
func OwnerNickname(nickname string) Param {
	return func(vals url.Values) error {
		if nickname == "" {
			return errors.New("empty nickname")
		}
		vals.Add("owner_nickname", nickname)
		return nil
	}
}

// SeriesID sets value to url.Values with key "series_id".
// seriesID must be positive integer.
func SeriesID(seriesID int) Param {
	return func(vals url.Values) error {
		if seriesID < 0 {
			return fmt.Errorf("invalid series id: %d", seriesID)
		}
		vals.Add("series_id", strconv.Itoa(seriesID))
		return nil
	}
}

// Start sets value to url.Values with key "start".
// start must be positive integer.
func Start(start int) Param {
	return func(vals url.Values) error {
		if start < 0 {
			return fmt.Errorf("invalid start: %d", start)
		}
		vals.Add("start", strconv.Itoa(start))
		return nil
	}
}

// OrderBy represents order of events.
type OrderBy int

const (
	OrderByUpdate OrderBy = 1
	OrderByDate   OrderBy = 2
	OrderByNewest OrderBy = 3
)

// Order sets value to url.Values with key "order".
// order must be OrderByUpdate, OrderByDate or OrderByNewest.
func Order(by OrderBy) Param {
	return func(vals url.Values) error {
		switch by {
		case OrderByUpdate, OrderByDate, OrderByNewest:
		default:
			return fmt.Errorf("invalid start: %d", by)
		}
		vals.Add("order", strconv.Itoa(int(by)))
		return nil
	}
}

// Count sets value to url.Values with key "count".
// count must be positive integer.
func Count(count int) Param {
	return func(vals url.Values) error {
		if count < 1 || count > 100 {
			return fmt.Errorf("count must be between 1 and 100: %d", count)
		}
		vals.Add("count", strconv.Itoa(count))
		return nil
	}
}
