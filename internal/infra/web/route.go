package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Heartbeat("/health"))

	mux.Post("/user", app.CreateUserHandler)
	mux.Get("/user", app.ListUserHandler)
	mux.Get("/user/{name}", app.GetUserHandler)
	
	mux.Post("/product", app.CreateProductHandler)
	mux.Get("/product", app.ListProductHandler)
	mux.Get("/product/{name}", app.GetProductHandler)

	mux.Post("/proccess-file",app.ProccesFileTransactioHandler)

	return mux
}
