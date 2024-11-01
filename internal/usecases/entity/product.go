package entity

type CreateProductRespone struct {
	Data    CreateProductDataRespone
	Message string `json:"message"`
}

type CreateProductDataRespone struct {
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
}
