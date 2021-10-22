package connpass

import (
	"context"
	"fmt"
	"net/url"

	"go.uber.org/multierr"
)

// SearchService provides the Search method which searches events of connpass.
type SearchService interface {
	Search(ctx context.Context, params url.Values) (*SearchResult, error)
}

// SearchParam sets parameter to url.Values.
// If params have validation errors, they will be returned as the second result.
// If you want to get individual errors, you can use "go.uber.org/multierr".Errors.
func SearchParam(params ...Param) (url.Values, error) {
	vals := make(url.Values)
	var err error
	for _, p := range params {
		err = multierr.Append(err, p(vals))
	}
	if err != nil {
		return nil, err
	}
	return vals, nil
}

// SearchResult represents a result of Search.
type SearchResult struct {
	Returned  int      `json:"results_returned"`
	Available int      `json:"results_available"`
	Start     int      `json:"results_start"`
	Events    []*Event `json:"events"`
}

type searchService struct {
	cli *Client
}

func (s *searchService) Search(ctx context.Context, params url.Values) (*SearchResult, error) {
	var r SearchResult
	if err := s.cli.get(ctx, "event", params, &r); err != nil {
		return nil, fmt.Errorf("connpass.Search: %w", err)
	}
	return &r, nil
}
