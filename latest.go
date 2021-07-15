package frankfurtergo

type Latest struct {
	Amount int      `json:"amount"`
	Base   string   `json:"base"`
	Date   Date     `json:"date"`
	Rates  []string `json:"string"`
}

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}
