package frankfurtergo

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func (c *Client) jsonRequest(path string, queries ...[]string) (body []byte, err error) {

	uri := fmt.Sprintf("https://%v/%v", c.options.Host, path)

	status, body, err := c.httpClient.Get(nil, uri)
	if err != nil {
		return nil, err
	}

	if status != fasthttp.StatusOK {
		return nil, err
	}

	return body, err
}
