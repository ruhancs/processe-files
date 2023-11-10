package usecase

import (
	"context"

	"github.com/ruhancs/hubla-test/internal/application/dto"
	"github.com/ruhancs/hubla-test/internal/domain/entity"
	"github.com/ruhancs/hubla-test/internal/domain/gateway"
)

type CreateProductUseCase struct {
	UserRepository gateway.UserRepositoryInterface
	ProductRepository gateway.ProductRepositoryInterface
}

func NewCreateProductUseCase(userRepo gateway.UserRepositoryInterface, productRepo gateway.ProductRepositoryInterface) *CreateProductUseCase{
	return &CreateProductUseCase{
		UserRepository: userRepo,
		ProductRepository: productRepo,
	}
}

func(usecase *CreateProductUseCase) Execute(ctx context.Context,input dto.CreateProductInputDto) (dto.CreateProductOutputDto,error) {
	producer,err := usecase.UserRepository.FindByName(ctx,input.ProducerName)
	if err != nil {
		return dto.CreateProductOutputDto{},err
	}

	product,err := entity.NewProduct(input.Title,producer.Name,input.Value)
	if err != nil {
		return dto.CreateProductOutputDto{},err
	}
	
	err = usecase.ProductRepository.Create(ctx,product)
	if err != nil {
		return dto.CreateProductOutputDto{},err
	}

	output := dto.CreateProductOutputDto{
		Title: product.Title,
		ProducerName: product.ProducerName,
		Value: product.Value,
	}

	return output,nil
}