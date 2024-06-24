package request

type ProductRequest struct {
	Title       string  `json:"title" binding:"required,max=255"`
	Amount      int     `json:"amount" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `json:"description" binding:"max=255"`
}
