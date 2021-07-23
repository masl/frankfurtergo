package frankfurtergo

type Historical struct {
	Amount float32            `json:"amount"`
	Base   string             `json:"base"`
	Date   Date               `json:"date"`
	Rates  map[string]float64 `json:"rates"`
}
