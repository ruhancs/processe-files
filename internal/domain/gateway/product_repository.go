package gateway

import "github.com/ruhancs/hubla-test/internal/domain/entity"

type ProductRepositoryInterface interface {
	Create(product *entity.Product) error
	Get(title string) (*entity.Product,error)
	List() ([]*entity.Product,error)
	Delete(id string) error
}