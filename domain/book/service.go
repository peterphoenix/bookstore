package book

import (
	"context"

	"bookstore/common/pagination"
	"bookstore/domain/book/model"
	"bookstore/domain/book/repository"
	"bookstore/infra/database"
)

type Service interface {
	ListBooks(ctx context.Context, pagination pagination.Pagination) (model.ListBooksRes, error)
	GetBookDetailsMap(ctx context.Context, req model.GetBookDetailsReq) (model.BookDetailsMap, error)
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
