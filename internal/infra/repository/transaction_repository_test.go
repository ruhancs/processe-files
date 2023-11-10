package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/ruhancs/hubla-test/internal/domain/entity"
	"github.com/ruhancs/hubla-test/internal/infra/db"
	"github.com/ruhancs/hubla-test/internal/infra/repository"
	"github.com/stretchr/testify/assert"
)

func initTransactionRepository() *repository.TransactionRepository {
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

	repository := repository.NewTransactionRepository(testDB)
	return repository
}

func createProduct(user *entity.User) *entity.Product{
	product,_ := entity.NewProduct("P1",user.Name,10)
	return product
}

func clearDB(repo *repository.TransactionRepository,tranID,userID,prodID string) {
	repo.Queries.DeleteTransaction(context.Background(),tranID)
	repo.Queries.DeleteProduct(context.Background(),prodID)
	repo.Queries.DeleteUser(context.Background(),userID)
}

func TestCreateTransaction(t *testing.T) {
	transactionRepository := initTransactionRepository()

	userEntity := createUser()
	err := transactionRepository.Queries.CreateUser(context.Background(),db.CreateUserParams{
		ID: userEntity.ID,
		Name: userEntity.Name,
		Balance: int32(userEntity.Balance),
	})
	assert.Nil(t,err)
	productEntity := createProduct(userEntity)
	err = transactionRepository.Queries.CreateProduct(context.Background(),db.CreateProductParams{
		ID: productEntity.ID,
		Title: productEntity.Title,
		ProducerName: productEntity.ProducerName,
		Value: int32(productEntity.Value),
	})
	assert.Nil(t,err)

	transactionEntity,_ := entity.NewTransaction(1,productEntity.Value,"2022-06-12",productEntity.Title,userEntity.Name)
	err = transactionRepository.Create(context.Background(),transactionEntity)

	assert.Nil(t,err)
	
	clearDB(transactionRepository,transactionEntity.ID,userEntity.ID,productEntity.ID)
}