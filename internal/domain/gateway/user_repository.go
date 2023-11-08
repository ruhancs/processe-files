package gateway

import "github.com/ruhancs/hubla-test/internal/domain/entity"

type UserRepositoryInterface interface {
	Create(user *entity.User) error
	List() ([]*entity.User, error)
	FindByName(name string) (*entity.User, error)
	UpdateBalance(newBalance float64) error
}
