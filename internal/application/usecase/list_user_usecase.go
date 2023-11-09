package usecase

import (
	"github.com/ruhancs/hubla-test/internal/application/dto"
	"github.com/ruhancs/hubla-test/internal/domain/gateway"
)

type ListUserUseCase struct {
	UserRepository gateway.UserRepositoryInterface
}

func NewListUserUseCase(repository gateway.UserRepositoryInterface) *ListUserUseCase {
	return &ListUserUseCase{
		UserRepository: repository,
	}
}

func (usecase *ListUserUseCase) Execute() (dto.ListUserOutputDto,error) {
	users,err := usecase.UserRepository.List()
	if err != nil {
		return dto.ListUserOutputDto{},err
	}

	var usersOut []dto.UserOutput
	for _,user := range users{
		userOut := dto.UserOutput{
			ID: user.ID,
			Name: user.Name,
			Balance: user.Balance,
		}
		usersOut = append(usersOut, userOut)
	}

	output := dto.ListUserOutputDto{
		Users: usersOut,
	}

	return output,nil
}