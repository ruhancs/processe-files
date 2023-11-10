package usecase_test

import (
	"testing"

	"github.com/ruhancs/hubla-test/internal/application/dto"
	"github.com/ruhancs/hubla-test/internal/application/usecase"
	"github.com/ruhancs/hubla-test/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

var product1,_ = entity.NewProduct("P1","U1",10)
var product2,_ = entity.NewProduct("P1","U1",10)

var listProductsEntities = []*entity.Product{product1,product2}

func(repo *ProductRepositoryMock) Create(user *entity.Product) error {
	//args := repo.Called(user)
	return nil
} 

func(repo *ProductRepositoryMock) List() ([]*entity.Product, error) {
	//args := repo.Called()
	return listProductsEntities,nil
} 

func(repo *ProductRepositoryMock) Get(id string) (*entity.Product, error) {
	//args := repo.Called(name)
	return product1,nil
} 

func(repo *ProductRepositoryMock) Delete(id string) error {
	//args := repo.Called(name)
	return nil
}

var productRepository = new(ProductRepositoryMock)

func TestCreateProductUseCase(t *testing.T) {
	createProductUseCase := usecase.NewCreateProductUseCase(userRepository,productRepository)

	input := dto.CreateProductInputDto{
		Title: "P1",
		ProducerName: "U1",
		Value: 10,
	}
	output,err := createProductUseCase.Execute(input)

	assert.Nil(t,err)
	assert.NotNil(t,output)
	assert.Equal(t,output.Title,"P1")
	assert.Equal(t,user1.Name,output.ProducerName)
	assert.Equal(t,output.Value,10.0)
}

func TestListProductUseCase(t *testing.T) {
	listProductUseCase := usecase.NewListProductUseCase(productRepository)

	output,err := listProductUseCase.Execute()

	assert.Nil(t,err)
	assert.NotNil(t,output)
	assert.Equal(t,output.Products[0].ID,product1.ID)
	assert.Equal(t,output.Products[0].ProducerName,product1.ProducerName)
	assert.Equal(t,output.Products[0].Title,product1.Title)
	assert.Equal(t,output.Products[0].Value,product1.Value)
	assert.Equal(t,output.Products[1].ID,product2.ID)
	assert.Equal(t,output.Products[1].ProducerName,product2.ProducerName)
	assert.Equal(t,output.Products[1].Title,product2.Title)
	assert.Equal(t,output.Products[1].Value,product2.Value)
}

func TestGetProductUseCase(t *testing.T) {
	getProductUseCase := usecase.NewGetProductUseCase(productRepository)

	output,err := getProductUseCase.Execute(product1.Title)

	assert.Nil(t,err)
	assert.NotNil(t,output)
	assert.Equal(t,output.ID,product1.ID)
	assert.Equal(t,output.Title,product1.Title)
	assert.Equal(t,output.ProducerName,product1.ProducerName)
	assert.Equal(t,output.Value,product1.Value)
}
