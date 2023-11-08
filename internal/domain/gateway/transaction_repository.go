package gateway

import "github.com/ruhancs/hubla-test/internal/domain/entity"

type TransactionRepositoryInterface interface {
	Create(transaction *entity.Transaction) error
	List() ([]*entity.Transaction,error)
	Get(id string) (*entity.Transaction,error)
}