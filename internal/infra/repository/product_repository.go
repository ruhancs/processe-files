package repository

import (
	"context"
	"database/sql"

	"github.com/ruhancs/hubla-test/internal/domain/entity"
	"github.com/ruhancs/hubla-test/internal/infra/db"
)

type ProductRepository struct {
	DB *sql.DB
	Queries *db.Queries
}

func NewProductRepository(database *sql.DB) *ProductRepository{
	return &ProductRepository{
		DB: database,
		Queries: db.New(database),
	}
}

func(repo *ProductRepository) Create(ctx context.Context, product *entity.Product) error{
	err := repo.Queries.CreateProduct(ctx,db.CreateProductParams{
		ID: product.ID,
		Title: product.Title,
		ProducerName: product.ProducerName,
		Value: int32(product.Value),
	})
	if err != nil {
		return err
	}
	return nil
}

func(repo *ProductRepository) List(ctx context.Context) ([]*entity.Product, error){
	productsModel,err := repo.Queries.ListProducts(ctx)
	if err != nil {
		return nil,err
	}
	
	var productsEntity []*entity.Product
	for _,productM := range productsModel {
		product := &entity.Product{ID: productM.ID,Title: productM.Title,Value: int(productM.Value), ProducerName: productM.ProducerName}
		productsEntity = append(productsEntity, product)
	}

	return productsEntity,nil
}

func(repo *ProductRepository) Get(ctx context.Context, title string) (*entity.Product, error){
	productModel,err := repo.Queries.GetProductByName(ctx,title)
	if err != nil {
		return nil,err
	}
	productEntity := &entity.Product{ID: productModel.ID,Title: productModel.Title,Value: int(productModel.Value), ProducerName: productModel.ProducerName}
	return productEntity,nil
}

func(repo *ProductRepository) Delete(ctx context.Context, id string) error{
	err := repo.Queries.DeleteProduct(ctx,id)
	if err != nil {
		return err
	}
	return nil
}

