package repository

import (
	"context"

	"bookstore/common/pagination"
	"bookstore/domain/book/model"
	commonmodel "bookstore/domain/common/model"
	"bookstore/infra/database"
)

type Repository interface {
	GetBooksList(ctx context.Context, pagination pagination.Pagination) ([]commonmodel.Book, error)
	GetBookDetails(ctx context.Context, id []string) ([]model.BookDetails, error)
}

type repository struct {
	psql *database.PostgreSQL
}

func NewRepository(psql *database.PostgreSQL) Repository {
	return &repository{psql}
}
