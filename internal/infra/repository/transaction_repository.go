package repository

import (
	"context"
	"database/sql"

	"github.com/ruhancs/hubla-test/internal/domain/entity"
	"github.com/ruhancs/hubla-test/internal/infra/db"
)

type TransactionRepository struct {
	DB *sql.DB
	Queries *db.Queries
}

func NewTransactionRepository(database *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: database,
		Queries: db.New(database),
	}
}

func(repo *TransactionRepository) Create(ctx context.Context,transaction *entity.Transaction) error{
	err := repo.Queries.CreateTransaction(ctx,db.CreateTransactionParams{
		ID: transaction.ID,
		Type: int32(transaction.Type),
		Value: int32(transaction.Value),
		Date: transaction.Date,
		ProductName: transaction.ProductName,
		SellerName: transaction.SellerName,
	})
	if err != nil {
		return err
	}
	return nil
}

func(repo *TransactionRepository) List(ctx context.Context) ([]*entity.Transaction,error){
	transactionsModel,err := repo.Queries.ListTransactions(ctx)
	if err != nil {
		return nil,err
	}
	
	var transactionsEntity []*entity.Transaction
	for _,transactionM := range transactionsModel {
		transaction := entity.Transaction{
			ID: transactionM.ID,
			Type: int(transactionM.Type),
			Date: transactionM.Date,
			Value: int(transactionM.Value),
			ProductName: transactionM.ProductName,
			SellerName: transactionM.SellerName,
		}
		transactionsEntity = append(transactionsEntity, &transaction)
	}

	return transactionsEntity,nil
}

func(repo *TransactionRepository) Get(ctx context.Context,id string) (*entity.Transaction,error){
	transactionModel,err := repo.Queries.GetTransactionById(ctx,id)
	if err != nil {
		return nil,err
	}

	transactionEntity := &entity.Transaction{
		ID: transactionModel.ID,
		Type: int(transactionModel.Type),
		Date: transactionModel.Date,
		Value: int(transactionModel.Value),
		ProductName: transactionModel.ProductName,
		SellerName: transactionModel.SellerName,
	}

	return transactionEntity,nil
}