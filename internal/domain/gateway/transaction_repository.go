package gateway

import (
	"context"

	"github.com/ruhancs/hubla-test/internal/domain/entity"
)

type TransactionRepositoryInterface interface {
	Create(ctx context.Context,transaction *entity.Transaction) error
	List(ctx context.Context) ([]*entity.Transaction,error)
	Get(ctx context.Context,id string) (*entity.Transaction,error)
}