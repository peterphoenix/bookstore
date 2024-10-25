package repository

//go:generate mockgen -source ./repository.go -destination ../../../test/mock/domain/user/repository.go -package=user_mock

import (
	"context"

	"bookstore/domain/user/model"
	"bookstore/infra/database"
)

type Repository interface {
	CreateUser(ctx context.Context, input model.CreateUserInput) error
}

type repository struct {
	psql *database.PostgreSQL
}

func NewRepository(psql *database.PostgreSQL) Repository {
	return &repository{psql}
}
