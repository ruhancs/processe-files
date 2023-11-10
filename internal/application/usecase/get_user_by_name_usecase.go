package usecase

import (
	"context"

	"github.com/ruhancs/hubla-test/internal/application/dto"
	"github.com/ruhancs/hubla-test/internal/domain/gateway"
)

type GetUserByNameUseCase struct {
	UserRepository gateway.UserRepositoryInterface
}

func NewGetUserByNameUseCase(repository gateway.UserRepositoryInterface) *GetUserByNameUseCase{
	return &GetUserByNameUseCase{
		UserRepository: repository,
	}
}

func (usecase *GetUserByNameUseCase) Execute(ctx context.Context,name string) (dto.UserOutput,error) {
	user,err := usecase.UserRepository.FindByName(ctx,name)
	if err != nil {
		return dto.UserOutput{},err
	}

	output := dto.UserOutput{
		ID: user.ID,
		Name: user.Name,
		Balance: user.Balance,
	}

	return output,nil
} 