package connpass

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://connpass.com/api/v1"
)

var (
	//go:embed version.txt
	version          string
	defaultUserAgent = "tenntenn-connpass/" + version + " (+https://github.com/tenntenn/connpass)"
)

// APIError represents an error of connpass API.
type APIError struct {
	StatusCode int
}

// Error implements error.Error.
func (err *APIError) Error() string {
	return fmt.Sprintf("StatusCode: %d", err.StatusCode)
}

// Client is a client of connpass API.
type Client struct {
	SearchService
	HTTPClient *http.Client
	BaseURL    string
	UserAgent  string
}

// NewClient creates a new connpass client.
func NewClient() *Client {
	var cli Client
	cli.SearchService = &searchService{cli: &cli}
	cli.HTTPClient = http.DefaultClient
	cli.BaseURL = defaultBaseURL
	cli.UserAgent = defaultUserAgent
	return &cli
}

func (cli *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", cli.UserAgent)
	return cli.HTTPClient.Do(req)
}

func (cli *Client) get(ctx context.Context, path string, params url.Values, v interface{}) error {
	reqURL := cli.BaseURL + "/" + path
	if params != nil {
		reqURL += "?" + params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return fmt.Errorf("cannot create HTTP request: %w", err)
	}

	resp, err := cli.do(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		return &APIError{StatusCode: resp.StatusCode}
	}

	var body io.Reader = resp.Body

	// For debug
	//var buf bytes.Buffer
	//io.Copy(&buf, body)
	//fmt.Println(buf.String())
	//body = &buf

	if err := json.NewDecoder(body).Decode(v); err != nil {
		return fmt.Errorf("cannot parse HTTP body: %w", err)
	}

	return nil
}
