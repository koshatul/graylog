package graylog

import (
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/google/go-querystring/query"
)

// TimeFormat is the recognised time format for the Graylog API
const TimeFormat = "2006-01-02T15:04:05.000Z"

// Response is a Graylog API response. This wraps the standard http.Response returned from Graylog.
type Response struct {
	*http.Response
}

// AbsoluteSearchOptions is the options available for use when using an absolute search.
type AbsoluteSearchOptions struct {
	Query    string    `url:"query"`
	From     string    `url:"from"`
	To       string    `url:"to"`
	Limit    int       `url:"limit,omitempty"`
	Offset   int       `url:"offset,omitempty"`
	Filter   string    `url:"filter,omitempty"`
	Fields   string    `url:"fields,omitempty"`
	Sort     string    `url:"sort,omitempty"`
	Decorate bool      `url:"decorate"`
	FromTime time.Time `url:"-"`
	ToTime   time.Time `url:"-"`
}

// ParseTime formats the From and To times into the fromTime and toTime strings for the query
func (a *AbsoluteSearchOptions) ParseTime() {
	if a.FromTime.Nanosecond() == 0 {
		a.From = a.FromTime.Add(time.Duration(-1) * time.Nanosecond).UTC().Format(TimeFormat)
	} else {
		a.From = a.FromTime.UTC().Format(TimeFormat)
	}
	if a.ToTime.Nanosecond() == 0 {
		a.To = a.ToTime.Add(1 * time.Nanosecond).UTC().Format(TimeFormat)
	} else {
		a.To = a.ToTime.UTC().Format(TimeFormat)
	}
}

// RelativeSearchOptions is the options available when using a relative search.
type RelativeSearchOptions struct {
	Query    string `url:"query"`
	Range    int    `url:"range"`
	Limit    int    `url:"limit,omitempty"`
	Offset   int    `url:"offset,omitempty"`
	Filter   string `url:"filter,omitempty"`
	Fields   string `url:"fields,omitempty"`
	Sort     string `url:"sort,omitempty"`
	Decorate bool   `url:"decorate"`
}

func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	origURL, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	origValues := origURL.Query()

	newValues, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	for k, v := range newValues {
		origValues[k] = v
	}

	origURL.RawQuery = origValues.Encode()
	return origURL.String(), nil
}
