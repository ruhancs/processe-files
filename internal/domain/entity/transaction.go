package entity

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	ID          string `json:"id" valid:"uuid"`
	Type        int    `json:"type" valid:"required"`
	Date        string `json:"date" valid:"required"`
	Value       int    `json:"value" valid:"required"`
	ProductName string `json:"product_name" valid:"required"`
	SellerName  string `json:"seller_name" valid:"required"`
}

func NewTransaction(transactionType, value int, date, productName, sellerName string) (*Transaction, error) {
	transaction := &Transaction{
		ID:          uuid.NewV4().String(),
		Type:        transactionType,
		Date:        date,
		Value:       value,
		ProductName: productName,
		SellerName:  sellerName,
	}

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return err
	}
	return nil
}
