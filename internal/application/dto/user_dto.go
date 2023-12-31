package dto

type CreateUserInputDto struct {
	Name string `json:"name"`
}

type CreateUserOutputDto struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type ListUserOutputDto struct {
	Users []UserOutput
}

type UserOutput struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
