package frankfurtergo

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
)

type RequestOptions struct {
	From   string   `json:"from"`
	To     []string `json:"to"`
	Amount float64  `json:"amount"`
}

func (c *Client) jsonRequest(path string, options ...RequestOptions) (body []byte, err error) {
	uri := url.URL{
		Scheme: "https",
		Host:   c.options.Host,
		Path:   path,
	}

	if len(options) > 0 {
		opt := options[0]
		queryValues := url.Values{}
		queryValues.Add("from", opt.From)
		if len(opt.To) > 0 {
			queryValues.Add("to", strings.Join(opt.To, ","))
		}
		queryValues.Add("amount", fmt.Sprintf("%.2f", opt.Amount))
		uri.RawQuery = queryValues.Encode()
	}

	status, body, err := c.httpClient.Get(nil, uri.String())
	if err != nil {
		return
	}

	if status != fasthttp.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}

	return
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
