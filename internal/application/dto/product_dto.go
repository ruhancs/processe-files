package dto

type CreateProductInputDto struct {
	Title        string  `json:"title"`
	ProducerName string  `json:"producer_id"`
	Value        float64 `json:"value"`
}

type CreateProductOutputDto struct {
	Title        string  `json:"title"`
	ProducerName string  `json:"producer_name"`
	Value        float64 `json:"value"`
}

type ProductOutputDto struct {
	ID           string  `json:"id"`
	Title        string  `json:"title"`
	ProducerName string  `json:"producer_name"`
	Value        float64 `json:"value"`
}

type ListProductsOutputDto struct {
	Products []ProductOutputDto `json:"products"`
}
