package graylog

import (
	"context"
	"net/http"
	"time"
)

const searchBasePath = "search"
const searchRelativeBasePath = "search/universal/relative"
const searchAbsoluteBasePath = "search/universal/absolute"

type usedIndices struct {
	IndexName    string    `json:"index_name"`
	Begin        time.Time `json:"begin"`
	End          time.Time `json:"end"`
	CalculatedAt time.Time `json:"calculated_at"`
	TookMs       int       `json:"took_ms"`
}

type highlightRanges struct {
	Source []struct {
		Start  int `json:"start"`
		Length int `json:"length"`
	} `json:"source"`
}

type searchRoot struct {
	Query       string        `json:"query"`
	BuiltQuery  string        `json:"built_query"`
	UsedIndices []usedIndices `json:"used_indices"`
	Messages    []struct {
		HighlightRanges highlightRanges        `json:"highlight_ranges"`
		Message         map[string]interface{} `json:"message"`
		Index           string                 `json:"index"`
		DecorationStats interface{}            `json:"decoration_stats"`
	} `json:"messages"`
	Fields          []string    `json:"fields"`
	Time            int         `json:"time"`
	TotalResults    int         `json:"total_results"`
	From            time.Time   `json:"from"`
	To              time.Time   `json:"to"`
	DecorationStats interface{} `json:"decoration_stats"`
}

// SearchService is an interface for interfacing with the Absolute Search endpoints of the Graylog API
type SearchService interface {
	Absolute(context.Context, *AbsoluteSearchOptions) ([]map[string]interface{}, *Response, error)
	Relative(context.Context, *RelativeSearchOptions) ([]map[string]interface{}, *Response, error)
}

// SearchServiceOp handles communication with the Absolute Search related methods of the API.
type SearchServiceOp struct {
	client *Client
}

// Performs a relative search request given a path.
func (s *SearchServiceOp) universalSearch(ctx context.Context, path string) ([]map[string]interface{}, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(searchRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}
	messages := []map[string]interface{}{}
	for _, message := range root.Messages {
		messages = append(messages, message.Message)
	}

	return messages, resp, err
}

// Relative does a relative search of messages.
func (s *SearchServiceOp) Relative(ctx context.Context, opt *RelativeSearchOptions) ([]map[string]interface{}, *Response, error) {
	path := searchRelativeBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.universalSearch(ctx, path)
}

// Absolute does an absolute search of messages.
func (s *SearchServiceOp) Absolute(ctx context.Context, opt *AbsoluteSearchOptions) ([]map[string]interface{}, *Response, error) {
	path := searchAbsoluteBasePath
	opt.ParseTime()
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.universalSearch(ctx, path)
}
