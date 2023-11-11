package factory

import (
	"database/sql"

	"github.com/ruhancs/hubla-test/internal/application/usecase"
	"github.com/ruhancs/hubla-test/internal/infra/repository"
)

func CreateUserUseCaseFactory(db *sql.DB) *usecase.CreateUserUseCase {
	userRepo := repository.NewUserRepository(db)
	usecase := usecase.NewCreateUserUseCase(userRepo)
	return usecase
}

func ListUserUseCaseFactory(db *sql.DB) *usecase.ListUserUseCase {
	userRepo := repository.NewUserRepository(db)
	usecase := usecase.NewListUserUseCase(userRepo)
	return usecase
}

func GetUserUseCaseFactory(db *sql.DB) *usecase.GetUserByNameUseCase {
	userRepo := repository.NewUserRepository(db)
	usecase := usecase.NewGetUserByNameUseCase(userRepo)
	return usecase
}

func CreateProductUseCaseFactory(db *sql.DB) *usecase.CreateProductUseCase {
	userRepo := repository.NewUserRepository(db)
	prodRepo := repository.NewProductRepository(db)
	usecase := usecase.NewCreateProductUseCase(userRepo,prodRepo)
	return usecase
}

func ListProductUseCaseFactory(db *sql.DB) *usecase.ListProductUseCase {
	prodRepo := repository.NewProductRepository(db)
	usecase := usecase.NewListProductUseCase(prodRepo)
	return usecase
}

func GetProductUseCaseFactory(db *sql.DB) *usecase.GetProductUseCase {
	prodRepo := repository.NewProductRepository(db)
	usecase := usecase.NewGetProductUseCase(prodRepo)
	return usecase
}

func ProccessFileUseCaseFactory(db *sql.DB) *usecase.ProccessFileUseCase {
	userRepo := repository.NewUserRepository(db)
	prodRepo := repository.NewProductRepository(db)
	tranRepo := repository.NewTransactionRepository(db)
	usecase := usecase.NewProccessFileUseCase(tranRepo,prodRepo,userRepo)
	return usecase
}