package frankfurtergo

import (
	"fmt"
	"net/url"

	"github.com/valyala/fasthttp"
)

func (c *Client) jsonRequest(path string, queries ...map[string]string) (body []byte, err error) {

	uri := url.URL{
		Scheme: "https",
		Host:   c.options.Host,
		Path:   path,
	}

	if len(queries) > 0 {
		queryValues := url.Values{}
		for k, v := range queries[0] {
			queryValues.Add(k, v)
		}
		uri.RawQuery = queryValues.Encode()
	}

	fmt.Println(uri.String())
	status, body, err := c.httpClient.Get(nil, uri.String())
	if err != nil {
		return nil, err
	}

	if status != fasthttp.StatusOK {
		return nil, err
	}

	return body, err
}
