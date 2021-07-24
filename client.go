package frankfurtergo

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

type ClientOptions struct {
	Host string `json:"host"`
}

type Client struct {
	options    *ClientOptions
	httpClient *fasthttp.Client
}

func New(options ...ClientOptions) *Client {
	var opt ClientOptions

	if len(options) == 1 {
		opt = options[0]
	}

	if opt.Host == "" {
		opt.Host = "api.frankfurter.app"
	}

	return &Client{
		options: &opt,
		httpClient: &fasthttp.Client{
			Name: "frankfurtergo",
		},
	}
}

func (c *Client) FetchHistorical(d time.Time, options ...RequestOptions) (historical *Historical, err error) {
	historical = &Historical{}

	opt := defaultRequestOptions(options)

	t := fmt.Sprintf("%02d-%02d-%02d", d.Year(), int(d.Month()), d.Day())
	body, err := c.jsonRequest(t, opt)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &historical)
	if err != nil {
		return
	}

	return
}

func (c *Client) FetchLatest(options ...RequestOptions) (latest *Latest, err error) {
	latest = &Latest{}

	opt := defaultRequestOptions(options)

	body, err := c.jsonRequest("latest", opt)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &latest)
	if err != nil {
		return
	}

	return
}

func (c *Client) FetchCurrencies() (currencies map[string]string, err error) {
	currencies = map[string]string{}

	body, err := c.jsonRequest("currencies")
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(body), &currencies)
	if err != nil {
		return
	}

	return
}
