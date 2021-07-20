package frankfurtergo

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
)

type RequestOptions struct {
	From   string   `json:"from"`
	To     []string `json:"to"`
	Amount int      `json:"amount"`
}

func (c *Client) jsonRequest(path string, options ...RequestOptions) (body []byte, err error) {
	opt := options[0]
	uri := url.URL{
		Scheme: "https",
		Host:   c.options.Host,
		Path:   path,
	}

	queryValues := url.Values{}
	queryValues.Add("from", opt.From)
	if len(opt.To) > 0 {
		queryValues.Add("to", strings.Join(opt.To, ","))
	}
	queryValues.Add("amount", strconv.Itoa(opt.Amount))
	uri.RawQuery = queryValues.Encode()

	status, body, err := c.httpClient.Get(nil, uri.String())
	if err != nil {
		return nil, err
	}

	if status != fasthttp.StatusOK {
		return nil, err
	}

	return body, err
}

func defaultRequestOptions(options []RequestOptions) (opt RequestOptions) {
	if len(options) > 0 {
		opt = options[0]
	}

	if opt.From == "" {
		opt.From = "EUR"
	}

	if len(opt.To) == 0 {
		opt.To = make([]string, 0)
	}

	if opt.Amount <= 0 {
		opt.Amount = 1
	}

	return
}
