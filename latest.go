package frankfurtergo

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Latest struct {
	Amount float32            `json:"amount"`
	Base   string             `json:"base"`
	Date   Date               `json:"date"`
	Rates  map[string]float64 `json:"rates"`
}

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

func (d *Date) UnmarshalJSON(data []byte) error {
	currencies := new(string)
	err := json.Unmarshal([]byte(data), &currencies)
	if err != nil {
		return err
	}

	dateP := strings.Split(*currencies, "-")

	d.Year, _ = strconv.Atoi(dateP[0])
	d.Month, _ = strconv.Atoi(dateP[1])
	d.Day, _ = strconv.Atoi(dateP[2])

	return nil
}
