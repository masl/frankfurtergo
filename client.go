package frankfurtergo

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type ClientOptions struct {
	Method string `json:"method"`
	Host   string `json:"host"`
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

	if opt.Method == "" {
		opt.Method = "GET"
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

func (c *Client) FetchCurrencies() (currencies map[string]string) {
	body, err := c.jsonRequest(c.options.Method, c.options.Host, "currencies")
	if err != nil {
		panic(err)
	}

	currencies = map[string]string{}
	err = json.Unmarshal([]byte(body), &currencies)
	if err != nil {
		panic(err)
	}

	return
}
