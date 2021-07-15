package frankfurtergo

import (
	"testing"
)

func TestClient(*testing.T) {
	s := New(ClientOptions{
		Endpoint: "Test",
	})

	s.FetchCurrencies()
}
