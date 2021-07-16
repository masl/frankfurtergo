package frankfurtergo

import (
	"testing"
)

func TestClient(*testing.T) {
	s := New(ClientOptions{
		Host: "api.frankfurter.app",
	})

	s.FetchCurrencies()
	s.FetchLatest()
}
