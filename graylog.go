package graylog

import (
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

// Response is a Graylog API response. This wraps the standard http.Response returned from Graylog.
type Response struct {
	*http.Response

	// // Links that were returned with the response. These are parsed from
	// // request body and not the header.
	// Links *Links

	// // Monitoring URI
	// Monitor string

	// Rate
}

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	PerPage int `url:"per_page,omitempty"`
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
