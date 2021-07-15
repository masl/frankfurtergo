package frankfurtergo

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func (c *Client) jsonRequest() (err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	req.Header.SetMethod("GET")
	req.SetRequestURI("https://api.frankfurter.app/currencies")
	req.Header.SetContentType("application/json")

	if err = c.client.Do(req, res); err != nil {
		return
	}
	fmt.Println(req, res)

	return
}
