package book

//go:generate mockgen -source ./service.go -destination ../../test/mock/domain/book/service.go -package=book_mock

import (
	"context"

	"bookstore/common/pagination"
	"bookstore/domain/book/model"
	"bookstore/domain/book/repository"
)

type Service interface {
	ListBooks(ctx context.Context, pagination pagination.Pagination) (model.ListBooksRes, error)
	GetBookDetailsMap(ctx context.Context, req model.GetBookDetailsReq) (model.BookDetailsMap, error)
}

type service struct {
	repo repository.Repository
}

func NewService(
	repo repository.Repository,
) Service {
	return &service{
		repo: repo,
	}
}
