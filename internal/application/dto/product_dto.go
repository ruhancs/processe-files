package dto

type CreateProductInputDto struct {
	Title        string  `json:"title"`
	ProducerName string  `json:"producer_id"`
	Value        int `json:"value"`
}

type CreateProductOutputDto struct {
	Title        string  `json:"title"`
	ProducerName string  `json:"producer_name"`
	Value        int `json:"value"`
}

type ProductOutputDto struct {
	ID           string  `json:"id"`
	Title        string  `json:"title"`
	ProducerName string  `json:"producer_name"`
	Value        int `json:"value"`
}

type ListProductsOutputDto struct {
	Products []ProductOutputDto `json:"products"`
}
