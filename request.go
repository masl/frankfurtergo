package frankfurtergo

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

type Request struct {
	Body string `json:"body"`
}

func (c *Client) jsonRequest(method, host, path string, queries ...[]string) (body []byte, err error) {

	uri := fmt.Sprintf("https://%v/%v", host, path)

	status, body, err := c.httpClient.Get(nil, uri)
	if err != nil {
		return nil, err
	}

	if status != fasthttp.StatusOK {
		return nil, err
	}

	return body, err
}
