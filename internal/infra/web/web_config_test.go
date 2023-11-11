package web_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ruhancs/hubla-test/internal/application/usecase"
	"github.com/ruhancs/hubla-test/internal/domain/entity"
	"github.com/ruhancs/hubla-test/internal/infra/repository"
	"github.com/ruhancs/hubla-test/internal/infra/web"
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

func initProductRepository() *repository.ProductRepository {
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

	repository := repository.NewProductRepository(testDB)
	return repository
}

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

func initCreateUserUseCase(repo *repository.UserRepository) *usecase.CreateUserUseCase{
	usecase := usecase.NewCreateUserUseCase(repo)

	return usecase
}

func initListUserUseCase(repo *repository.UserRepository) *usecase.ListUserUseCase{
	usecase := usecase.NewListUserUseCase(repo)

	return usecase
}

func initGetUserUseCase(repo *repository.UserRepository) *usecase.GetUserByNameUseCase{
	usecase := usecase.NewGetUserByNameUseCase(repo)

	return usecase
}

func initCreateProductUseCase(userRepo *repository.UserRepository,prodRepo *repository.ProductRepository) *usecase.CreateProductUseCase{
	usecase := usecase.NewCreateProductUseCase(userRepo,prodRepo)

	return usecase
}

func initListProductUseCase(prodRepo *repository.ProductRepository) *usecase.ListProductUseCase{
	usecase := usecase.NewListProductUseCase(prodRepo)
	
	return usecase
}

func initGetProductUseCase(prodRepo *repository.ProductRepository) *usecase.GetProductUseCase{
	usecase := usecase.NewGetProductUseCase(prodRepo)
	
	return usecase
}

func initProccessFileUseCase(
	userRepo *repository.UserRepository,
	prodRepo *repository.ProductRepository,
	transRepo *repository.TransactionRepository,
) *usecase.ProccessFileUseCase{
	usecase := usecase.NewProccessFileUseCase(transRepo,prodRepo,userRepo)

	return usecase
}

func initApplication(
	creteUser *usecase.CreateUserUseCase,
	listUser *usecase.ListUserUseCase,
	getUser *usecase.GetUserByNameUseCase,
	createProduct *usecase.CreateProductUseCase,
	listProduct *usecase.ListProductUseCase,
	getProduct *usecase.GetProductUseCase,
	proccessFile *usecase.ProccessFileUseCase,
) *web.Application {
	app := web.NewApplication(creteUser,listUser,getUser,createProduct,listProduct,getProduct,proccessFile)

	return app
}

func createUser() *entity.User {
	user, _ := entity.NewUser("U1")
	return user
}