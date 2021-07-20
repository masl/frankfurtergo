package frankfurtergo

import (
	"encoding/json"

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

func (c *Client) FetchLatest(options ...RequestOptions) (latest *Latest, err error) {
	latest = &Latest{}

	opt := defaultRequestOptions(options)

	body, err := c.jsonRequest("latest", opt)
	if err != nil {
		return latest, err
	}

	err = json.Unmarshal(body, &latest)
	if err != nil {
		return latest, err
	}

	return
}

func (c *Client) FetchCurrencies() (currencies map[string]string, err error) {
	body, err := c.jsonRequest("currencies")
	if err != nil {
		return nil, err
	}

	currencies = map[string]string{}
	err = json.Unmarshal([]byte(body), &currencies)
	if err != nil {
		return nil, err
	}

	return
}
