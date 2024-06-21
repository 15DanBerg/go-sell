package request

type ProductRequest struct {
	Title       string  `json:"title"`
	Amount      int     `json:"amount"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
