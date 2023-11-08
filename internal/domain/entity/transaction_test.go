package entity_test

import (
	"testing"

	"github.com/ruhancs/hubla-test/internal/domain/entity"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	productID := uuid.NewV4().String()
	sellerID := uuid.NewV4().String()
	transaction,err := entity.NewTransaction(1,"2022-01-15T19:20",productID,sellerID, "S1")
	
	assert.Nil(t,err)
	assert.NotNil(t,transaction)
	assert.Equal(t,transaction.ProductID,productID)
	assert.Equal(t,transaction.SellerID,sellerID)
	assert.Equal(t,transaction.SellerName,"S1")
	assert.Equal(t,transaction.Type,1)
}
