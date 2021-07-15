package frankfurtergo

import (
	"github.com/valyala/fasthttp"
)

type ClientOptions struct {
	Endpoint string `json:"endpoint"`
}

type Client struct {
	options    *ClientOptions
	httpClient *fasthttp.Client
}

func New(options ...ClientOptions) *Client {
	var opt ClientOptions

	if len(options) > 0 {
		opt = options[0]
	}

	if opt.Endpoint == "" {
		opt.Endpoint = "GET"
	}

	return &Client{
		options: &opt,
		httpClient: &fasthttp.Client{
			Name: "frankfurtergo",
		},
	}
}

func (c *Client) FetchCurrencies() {
	err := c.jsonRequest()
	if err != nil {
		panic(err)
	}
}
