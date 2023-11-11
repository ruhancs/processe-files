package main

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/ruhancs/hubla-test/internal/application/factory"
	"github.com/ruhancs/hubla-test/internal/infra/web"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db,err := sql.Open(os.Getenv("DB_DRIVER"),os.Getenv("DB_SOURCE"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	createUserUseCase := factory.CreateUserUseCaseFactory(db)
	listUserUseCase := factory.ListUserUseCaseFactory(db)
	getUserUseCase := factory.GetUserUseCaseFactory(db)
	createProductUseCase := factory.CreateProductUseCaseFactory(db)
	listProductUseCase := factory.ListProductUseCaseFactory(db)
	getProductUseCase := factory.GetProductUseCaseFactory(db)
	proccessFileUseCase := factory.ProccessFileUseCaseFactory(db)

	app := web.NewApplication(
		createUserUseCase,
		listUserUseCase,
		getUserUseCase,
		createProductUseCase,
		listProductUseCase,
		getProductUseCase,
		proccessFileUseCase,
	)

	app.Server()
}