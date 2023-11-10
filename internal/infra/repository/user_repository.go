package repository

import (
	"context"
	"database/sql"

	"github.com/ruhancs/hubla-test/internal/domain/entity"
	"github.com/ruhancs/hubla-test/internal/infra/db"
)

type UserRepository struct {
	DB *sql.DB
	Queries *db.Queries
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{
		DB: database,
		Queries: db.New(database),
	}
}

func(repo *UserRepository) Create(ctx context.Context,userEntity entity.User) error {
	err := repo.Queries.CreateUser(ctx,db.CreateUserParams{
		ID: userEntity.ID,
		Name: userEntity.Name,
		Balance: int32(userEntity.Balance),
	})
	if err != nil {
		return err
	}
	return nil
}

func(repo *UserRepository) List(ctx context.Context) ([]*entity.User, error) {
	usersModel,err := repo.Queries.ListUsers(ctx)
	if err != nil {
		return nil,err
	}
	var usersEntity []*entity.User
	for _,user := range usersModel {
		userEntity := entity.User{ID: user.ID,Name: user.Name,Balance: int(user.Balance)}
		usersEntity = append(usersEntity, &userEntity)
	}
	return usersEntity,nil
}

func(repo *UserRepository) FindByName(ctx context.Context,name string) (*entity.User, error) {
	userModel,err := repo.Queries.GetUserByName(ctx,name)
	if err != nil {
		return nil,err
	}

	userEntity := entity.User{ID: userModel.ID,Name: userModel.Name,Balance: int(userModel.Balance)}
	
	return &userEntity,nil
}

func(repo *UserRepository) UpdateBalance(ctx context.Context,username string,newBalance int) error {
	_,err := repo.Queries.UpdateUserBalance(ctx,db.UpdateUserBalanceParams{
		Name: username,
		Balance: int32(newBalance),
	})
	if err != nil {
		return err
	}	
	return nil
}