package entity_test

import (
	"testing"

	"github.com/ruhancs/hubla-test/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestValidNewProduct(t *testing.T) {
	product,err := entity.NewProduct("P1","P1",10)
	
	assert.Nil(t,err)
	assert.NotNil(t,product)
	assert.Equal(t,"P1",product.ProducerName)
	assert.Equal(t,10.0,product.Value)
	assert.Equal(t,"P1",product.Title)
}

func TestInvalidProductTitle(t *testing.T) {
	_,err := entity.NewProduct("","P1",10)

	assert.NotNil(t,err)
	assert.Equal(t,err.Error(),"title: non zero value required")
}

func TestInvalidProductProducerName(t *testing.T) {
	_,err := entity.NewProduct("P1","",10)

	assert.NotNil(t,err.Error())
	assert.Equal(t,err.Error(),"producer_name: non zero value required")
}