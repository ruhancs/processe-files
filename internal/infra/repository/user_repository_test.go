package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/ruhancs/hubla-test/internal/domain/entity"
	"github.com/ruhancs/hubla-test/internal/infra/db"
	"github.com/ruhancs/hubla-test/internal/infra/repository"
	"github.com/stretchr/testify/assert"
)

func initUserRepository() *repository.UserRepository {
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Println("Error loading .env")
	}

	dbDriver := os.Getenv("DB_DRIVER_TEST")
	dbSource := os.Getenv("DB_SOURCE_TEST")

	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Println(err)
		log.Fatal("cannot connect to db")
	}

	repository := repository.NewUserRepository(testDB)
	return repository
}

func createUser() *entity.User {
	user, _ := entity.NewUser("U1")
	return user
}

func TestCreateUser(t *testing.T) {
	var userRepository = initUserRepository()
	userEntity := createUser()

	err := userRepository.Create(context.Background(), *userEntity)
	assert.Nil(t, err)

	userRepository.Queries.DeleteUser(context.Background(),userEntity.ID)
}

func TestListUser(t *testing.T) {
	var userRepository = initUserRepository()
	userEntity := createUser()

	err := userRepository.Queries.CreateUser(context.Background(),db.CreateUserParams{
		ID: userEntity.ID,
		Name: userEntity.Name,
		Balance: int32(userEntity.Balance),
	})
	assert.Nil(t, err)

	users,err := userRepository.List(context.Background())

	assert.Nil(t,err)
	assert.NotNil(t,users)
	assert.Equal(t,userEntity.ID,users[0].ID)
	assert.Equal(t,userEntity.Name,users[0].Name)
	assert.Equal(t,userEntity.Balance,users[0].Balance)

	userRepository.Queries.DeleteUser(context.Background(),userEntity.ID)
}

func TestFindUserByName(t *testing.T) {
	var userRepository = initUserRepository()
	userEntity := createUser()

	err := userRepository.Queries.CreateUser(context.Background(),db.CreateUserParams{
		ID: userEntity.ID,
		Name: userEntity.Name,
		Balance: int32(userEntity.Balance),
	})
	assert.Nil(t, err)

	user,err := userRepository.FindByName(context.Background(),userEntity.Name)

	assert.Nil(t,err)
	assert.NotNil(t,user)
	assert.Equal(t,userEntity.ID,user.ID)
	assert.Equal(t,userEntity.Name,user.Name)
	assert.Equal(t,userEntity.Balance,user.Balance)

	userRepository.Queries.DeleteUser(context.Background(),userEntity.ID)
}

func TestUserNotFound(t *testing.T) {
	var userRepository = initUserRepository()
	
	user,err := userRepository.FindByName(context.Background(),"not found")
	
	assert.Nil(t,user)
	assert.NotNil(t,err)
	assert.Equal(t,"sql: no rows in result set",err.Error())
}

func TestUpdateUserBalance(t *testing.T) {
	var userRepository = initUserRepository()
	userEntity := createUser()

	err := userRepository.Queries.CreateUser(context.Background(),db.CreateUserParams{
		ID: userEntity.ID,
		Name: userEntity.Name,
		Balance: int32(userEntity.Balance),
	})
	assert.Nil(t, err)

	err = userRepository.UpdateBalance(context.Background(),userEntity.Name,10)
	assert.Nil(t,err)
	
	user,_ := userRepository.Queries.GetUserByName(context.Background(),userEntity.Name)

	assert.NotNil(t,user)
	assert.Equal(t,userEntity.ID,user.ID)
	assert.Equal(t,userEntity.Name,user.Name)
	assert.Equal(t,int32(10),user.Balance)

	userRepository.Queries.DeleteUser(context.Background(),userEntity.ID)
}
