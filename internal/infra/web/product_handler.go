package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ruhancs/hubla-test/internal/application/dto"
)

func(app *Application) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto.CreateProductInputDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}

	output,err := app.CreateProductUseCase.Execute(r.Context(),inputDto)
	if err != nil {
		fmt.Println(inputDto)
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}

	app.writeJson(w,http.StatusCreated,output)
}

func(app *Application) ListProductHandler(w http.ResponseWriter, r *http.Request) {
	output,err := app.ListProductsUseCase.Execute(r.Context())
	if err != nil {
		app.errorJson(w,err,http.StatusInternalServerError)
		return
	}

	app.writeJson(w,http.StatusOK,output)
}

func(app *Application) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r,"name")
	output,err := app.GetProductUseCase.Execute(r.Context(),name)
	if err != nil {
		app.errorJson(w,err,http.StatusNotFound)
		return
	}

	app.writeJson(w,http.StatusOK,output)
}