package web_test

import (
	"bytes"
	//"encoding/json"
	"io"
	"net/http"
	"testing"

	_ "github.com/lib/pq"
	//"github.com/ruhancs/hubla-test/internal/application/dto"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserHandler(t *testing.T) {
	userRepo := initUserRepository()
	prodRepo := initProductRepository()
	transRepo := initTransactionRepository()

	createUserUsecase := initCreateUserUseCase(userRepo)
	listUser := initListUserUseCase(userRepo)
	getUser := initGetUserUseCase(userRepo)
	createProd := initCreateProductUseCase(userRepo, prodRepo)
	listProd := initListProductUseCase(prodRepo)
	getProd := initGetProductUseCase(prodRepo)
	proccessFile := initProccessFileUseCase(userRepo,prodRepo,transRepo)

	app := initApplication(createUserUsecase,listUser,getUser,createProd,listProd,getProd,proccessFile)

	go app.Server()

	input := []byte(`{"name": "U1"}`)
	req,err := http.NewRequest("POST","http://localhost:8000/user",bytes.NewBuffer(input))
	assert.Nil(t,err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp,err := client.Do(req)
	assert.Nil(t,err)
	defer resp.Body.Close()
	
	outputBytes,err := io.ReadAll(req.Body)
	assert.Nil(t,err)
	assert.Nil(t,string(outputBytes))
	//var output dto.CreateUserOutputDto
	//err = json.Unmarshal(outputBytes,&output)
	//assert.Nil(t,err)

	//assert.Nil(t,output)
}
