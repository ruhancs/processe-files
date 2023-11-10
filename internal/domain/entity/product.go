package entity

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID           string `json:"id" valid:"uuid"`
	Title        string `json:"title" valid:"required"`
	ProducerName string `json:"producer_name" valid:"required"`
	Value        int    `json:"value" valid:"float"`
}

func NewProduct(title, producerName string, value int) (*Product, error) {
	product := &Product{
		ID:           uuid.NewV4().String(),
		Title:        title,
		ProducerName: producerName,
		Value:        value,
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
