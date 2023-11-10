package usecase

import (
	"context"

	"github.com/ruhancs/hubla-test/internal/application/dto"
	"github.com/ruhancs/hubla-test/internal/domain/entity"
	"github.com/ruhancs/hubla-test/internal/domain/gateway"
)

type CreateUserUseCase struct {
	UserRepository gateway.UserRepositoryInterface
}

func NewCreateUserUseCase(repository gateway.UserRepositoryInterface) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: repository,
	}
}

func(usecase *CreateUserUseCase) Execute(ctx context.Context,input dto.CreateUserInputDto) (dto.CreateUserOutputDto,error) {
	user,err := entity.NewUser(input.Name)
	if err != nil {
		return dto.CreateUserOutputDto{},nil
	}

	err = usecase.UserRepository.Create(ctx,user)
	if err != nil {
		return dto.CreateUserOutputDto{},err
	}

	output := dto.CreateUserOutputDto{
		ID: user.ID,
		Name: user.Name,
		Balance: user.Balance,
	}

	return output,nil
}