package repository

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
