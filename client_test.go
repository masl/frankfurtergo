package frankfurtergo

import (
	"testing"
)

func TestClient(*testing.T) {
	s := New(ClientOptions{
		Method: "GET",
		Host:   "api.frankfurter.app",
	})

	s.FetchCurrencies()
}
