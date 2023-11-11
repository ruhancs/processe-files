package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ruhancs/hubla-test/internal/application/dto"
)

func(app *Application) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto.CreateUserInputDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}
	
	output,err := app.CreateUserUseCase.Execute(r.Context(),inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}
	
	app.writeJson(w,http.StatusOK,output)
}

func(app *Application) ListUserHandler(w http.ResponseWriter, r *http.Request) {
	output,err := app.ListUserUseCase.Execute(r.Context())
	if err != nil {
		app.errorJson(w,err,http.StatusInternalServerError)
		return
	}
	
	app.writeJson(w,http.StatusOK,output)
}

func(app *Application) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "name")
	output,err := app.GetUserByNameUseCase.Execute(r.Context(),userName)
	if err != nil {
		app.errorJson(w,err,http.StatusNotFound)
		return
	}

	app.writeJson(w,http.StatusOK,output)
}