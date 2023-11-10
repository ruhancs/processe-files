package entity

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// User pode produtor ou filiado, valor de balance é em centavos
type User struct {
	ID      string `json:"id" valid:"required,uuid"`
	Name    string `json:"name" valid:"required"`
	Balance int    `json:"balance" valid:"float"`
}

func NewUser(name string) (*User, error) {
	user := &User{
		ID:      uuid.NewV4().String(),
		Name:    name,
		Balance: 0,
	}

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) isValid() error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	return nil
}
