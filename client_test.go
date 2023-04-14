package frankfurtergo

import (
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	s := New(ClientOptions{
		Host: "api.frankfurter.app",
	})

	s.FetchCurrencies()
	s.FetchLatest()
	t.Log(s.FetchLatest())
	t.Log(s.FetchSeries(time.Now().AddDate(0, 0, -3), time.Now(), RequestOptions{
		From: "USD",
		To:   []string{"EUR", "PLN"},
	}))
}
