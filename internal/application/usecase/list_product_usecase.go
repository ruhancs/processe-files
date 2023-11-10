package usecase

import (
	"context"

	"github.com/ruhancs/hubla-test/internal/application/dto"
	"github.com/ruhancs/hubla-test/internal/domain/gateway"
)

type ListProductUseCase struct {
	ProductRepository gateway.ProductRepositoryInterface
}

func NewListProductUseCase(productRepo gateway.ProductRepositoryInterface) *ListProductUseCase{
	return &ListProductUseCase{
		ProductRepository: productRepo,
	}
}

func(usecase *ListProductUseCase) Execute(ctx context.Context) (dto.ListProductsOutputDto,error) {
	products,err := usecase.ProductRepository.List(ctx)
	if err != nil {
		return dto.ListProductsOutputDto{},err
	}

	var outProducts []dto.ProductOutputDto
	for _,product := range products {
		outProd := dto.ProductOutputDto{
			ID: product.ID,
			Title: product.Title,
			ProducerName: product.ProducerName,
			Value: product.Value,
		}
		outProducts = append(outProducts, outProd)
	}
	
	output := dto.ListProductsOutputDto{
		Products: outProducts,
	}

	return output,nil
}