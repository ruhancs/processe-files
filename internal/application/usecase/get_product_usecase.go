package usecase

import (
	"github.com/ruhancs/hubla-test/internal/application/dto"
	"github.com/ruhancs/hubla-test/internal/domain/gateway"
)

type GetProductUseCase struct {
	ProductRepository gateway.ProductRepositoryInterface
}

func NewGetProductUseCase(productRepo gateway.ProductRepositoryInterface) *GetProductUseCase{
	return &GetProductUseCase{
		ProductRepository: productRepo,
	}
}

func(usecase *GetProductUseCase) Execute(name string) (dto.ProductOutputDto,error) {
	product,err := usecase.ProductRepository.Get(name)
	if err != nil {
		return dto.ProductOutputDto{},err
	}

	output := dto.ProductOutputDto{
		ID: product.ID,
		Title: product.Title,
		ProducerName: product.ProducerName,
		Value: product.Value,
	}

	return output,nil
}