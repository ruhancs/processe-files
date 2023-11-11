package web

import (
	"log"
	"net/http"
	"time"

	"github.com/ruhancs/hubla-test/internal/application/usecase"
)

type Application struct {
	CreateUserUseCase    *usecase.CreateUserUseCase
	ListUserUseCase      *usecase.ListUserUseCase
	GetUserByNameUseCase *usecase.GetUserByNameUseCase
	CreateProductUseCase *usecase.CreateProductUseCase
	ListProductsUseCase  *usecase.ListProductUseCase
	GetProductUseCase    *usecase.GetProductUseCase
	ProccessFileUseCase  *usecase.ProccessFileUseCase
}

func NewApplication(
	createUserUseCase *usecase.CreateUserUseCase,
	listUserUseCase *usecase.ListUserUseCase,
	getUserByNameUseCase *usecase.GetUserByNameUseCase,
	createProductUseCase *usecase.CreateProductUseCase,
	listProductUseCase *usecase.ListProductUseCase,
	getProductUseCase *usecase.GetProductUseCase,
	proccessFileUseCase *usecase.ProccessFileUseCase,
) *Application {
	return &Application{
		CreateUserUseCase:    createUserUseCase,
		ListUserUseCase:      listUserUseCase,
		GetUserByNameUseCase: getUserByNameUseCase,
		CreateProductUseCase: createProductUseCase,
		ListProductsUseCase:  listProductUseCase,
		GetProductUseCase:    getProductUseCase,
		ProccessFileUseCase:  proccessFileUseCase,
	}
}

func (app *Application) Server() error {
	srv := &http.Server{
		Addr:              ":8000",
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Println("Runing server on port 8000...")
	return srv.ListenAndServe()
}
