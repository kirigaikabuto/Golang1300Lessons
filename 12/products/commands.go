package products

type CreateProductCommand struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}