package response

type ProductRequest struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Amount      int     `json:"amount"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
