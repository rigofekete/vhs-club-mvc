package model

type Tape struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Director string  `json:"director"`
	Genre    string  `json:"genre"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
