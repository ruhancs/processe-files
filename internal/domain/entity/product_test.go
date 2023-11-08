package entity_test

import (
	"testing"

	"github.com/ruhancs/hubla-test/internal/domain/entity"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidNewProduct(t *testing.T) {
	producerId := uuid.NewV4().String()
	product,err := entity.NewProduct("P1",producerId,10)
	
	assert.Nil(t,err)
	assert.NotNil(t,product)
	assert.Equal(t,producerId,product.ProducerID)
	assert.Equal(t,10.0,product.Value)
	assert.Equal(t,"P1",product.Title)
}

func TestInvalidProductTitle(t *testing.T) {
	producerId := uuid.NewV4().String()
	_,err := entity.NewProduct("",producerId,10)

	assert.NotNil(t,err)
	assert.Equal(t,err.Error(),"title: non zero value required")
}

func TestInvalidProductProducerID(t *testing.T) {
	producerId := "id invalido"
	_,err := entity.NewProduct("P1",producerId,10)

	assert.NotNil(t,err.Error())
	assert.Equal(t,err.Error(),"producer_id: id invalido does not validate as uuid")
}