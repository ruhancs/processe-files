package gateway

import (
	"context"

	"github.com/ruhancs/hubla-test/internal/domain/entity"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context,user *entity.User) error
	List(ctx context.Context) ([]*entity.User, error)
	FindByName(ctx context.Context,name string) (*entity.User, error)
	UpdateBalance(ctx context.Context,username string,newBalance int) error
}
