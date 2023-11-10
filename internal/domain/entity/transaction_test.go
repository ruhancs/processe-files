package entity_test

import (
	"testing"

	"github.com/ruhancs/hubla-test/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	
	transaction,err := entity.NewTransaction(1,10,"2022-01-15T19:20","P1", "S1")
	
	assert.Nil(t,err)
	assert.NotNil(t,transaction)
	assert.Equal(t,transaction.ProductName,"P1")
	assert.Equal(t,transaction.SellerName,"S1")
	assert.Equal(t,transaction.Type,1)
}
