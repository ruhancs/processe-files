package entity

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	ID        string `json:"id" valid:"uuid"`
	Type      int    `json:"type" valid:"required"`
	Date      string `json:"date" valid:"required"`
	ProductID string `json:"product_id" valid:"required,uuid"`
	SellerID  string `json:"seler_id" valid:"required,uuid"`
	SellerName  string `json:"seller_name" valid:"required"`
}

func NewTransaction(transactionType int, date, productID, sellerID, sellerName string)(*Transaction,error) {
	transaction := &Transaction{
		ID:        uuid.NewV4().String(),
		Type:      transactionType,
		Date:      date,
		ProductID: productID,
		SellerID:  sellerID,
		SellerName: sellerName,
	}

	err := transaction.isValid()
	if err != nil {
		return nil,err
	}

	return transaction,nil
}

func (t *Transaction) isValid() error {
	_,err := govalidator.ValidateStruct(t)
	if err != nil {
		return err
	}
	return nil
}
