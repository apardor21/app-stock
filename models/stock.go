package models

type Stock struct {
	Symbol        string  `json:"symbol"`
	Price         float64 `json:"price"`
	ChangePercent float64 `json:"change_percent"`
	Volume        int64   `json:"volume"`
	Source        string  `json:"source"`
	CreatedAt     string  `json:"created_at"`
}
