package user

import (
	"context"

	"bookstore/domain/user/model"
	"bookstore/domain/user/repository"
	"bookstore/infra/database"
)

type Service interface {
	CreateUser(ctx context.Context, req model.UserCreateReq) (model.UserCreateRes, error)
}

type service struct {
	repo repository.Repository
}

func NewService(
	db *database.PostgreSQL,
) Service {
	return &service{
		repo: repository.NewRepository(db),
	}
}
