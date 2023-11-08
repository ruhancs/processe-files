package entity

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID         string  `json:"id" valid:"uuid"`
	Title      string  `json:"title" valid:"required"`
	ProducerID string  `json:"producer_id" valid:"uuid"`
	Value      float64 `json:"value" valid:"float"`
}

func NewProduct(title, producerID string, value float64) (*Product, error) {
	product := &Product{
		ID:         uuid.NewV4().String(),
		Title:      title,
		ProducerID: producerID,
		Value:      float64(value),
	}

	err := product.IsValid()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) IsValid() error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}
	return nil
}
