package model

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image string  `json:"description"`
}

type Shipping struct {
	ID       int     `json:"id"`
	From     string  `json:"from"`
	To       string  `json:"to"`
	Distance float64 `json:"distance"`
	Price    float64 `json:"price"`
}

type Order struct {
	ID       int      `json:"id"`
	Product  Product  `json:"product"`
	Shipping Shipping `json:"shipping"`
	Quantity int      `json:"quantity"`
	Total    float64  `json:"total"`
}
