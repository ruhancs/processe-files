package gateway

import (
	"context"

	"github.com/ruhancs/hubla-test/internal/domain/entity"
)

type ProductRepositoryInterface interface {
	Create(ctx context.Context, product *entity.Product) error
	Get(ctx context.Context, title string) (*entity.Product, error)
	List(ctx context.Context) ([]*entity.Product, error)
	Delete(ctx context.Context, id string) error
}
