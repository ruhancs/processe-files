package usecase_test

import (
	"context"
	"testing"

	"github.com/ruhancs/hubla-test/internal/application/dto"
	"github.com/ruhancs/hubla-test/internal/application/usecase"
	"github.com/ruhancs/hubla-test/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

var user1,_ = entity.NewUser("U1")
var user2,_ = entity.NewUser("U2")

var listUserEntities = []*entity.User{user1,user2}

func(repo *UserRepositoryMock) Create(ctx context.Context,user *entity.User) error {
	//args := repo.Called(user)
	return nil
} 

func(repo *UserRepositoryMock) List(ctx context.Context) ([]*entity.User, error) {
	//args := repo.Called()
	return listUserEntities,nil
} 

func(repo *UserRepositoryMock) FindByName(ctx context.Context,name string) (*entity.User, error) {
	//args := repo.Called(name)
	return user1,nil
} 

func(repo *UserRepositoryMock) UpdateBalance(ctx context.Context,username string, newBalance int) error {
	//args := repo.Called(name)
	return nil
}

var userRepository = new(UserRepositoryMock)

func TestCreateUserUseCase(t *testing.T) {
	createUserUseCase := usecase.NewCreateUserUseCase(userRepository)

	input := dto.CreateUserInputDto{
		Name: "U1",
	}
	output,err := createUserUseCase.Execute(context.Background(),input)

	assert.Nil(t,err)
	assert.NotNil(t,output)
	assert.Equal(t,output.Name,"U1")
	assert.Equal(t,output.Balance,0.0)
	assert.NotNil(t,output.ID)
}

func TestListUserUseCase(t *testing.T) {
	listUserUseCase := usecase.NewListUserUseCase(userRepository)

	output,err := listUserUseCase.Execute(context.Background())

	assert.Nil(t,err)
	assert.NotNil(t,output)
	assert.Equal(t,output.Users[0],dto.UserOutput{ID: user1.ID,Name: user1.Name,Balance: user1.Balance})
	assert.Equal(t,output.Users[1],dto.UserOutput{ID: user2.ID,Name: user2.Name,Balance: user2.Balance})
}

func TestGetUserByNameUseCase(t *testing.T) {
	getUserUseCase := usecase.NewGetUserByNameUseCase(userRepository)

	output,err := getUserUseCase.Execute(context.Background(),"U1")

	assert.Nil(t,err)
	assert.NotNil(t,output)
	assert.Equal(t,output.ID,user1.ID)
	assert.Equal(t,output.Name,"U1")
	assert.Equal(t,output.Balance,user1.Balance)
}