package frankfurtergo

import (
	"net/url"

	"github.com/valyala/fasthttp"
)

func (c *Client) jsonRequest(path string, queries ...[]string) (body []byte, err error) {

	uri := url.URL{
		Scheme: "https",
		Host:   c.options.Host,
		Path:   path,
	}

	status, body, err := c.httpClient.Get(nil, uri.String())
	if err != nil {
		return nil, err
	}

	if status != fasthttp.StatusOK {
		return nil, err
	}

	return body, err
}
