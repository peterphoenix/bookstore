package book

import (
	"context"

	"bookstore/common/pagination"
	"bookstore/domain/book/model"
)

func (s service) ListBooks(ctx context.Context, pagination pagination.Pagination) (model.ListBooksRes, error) {
	books, err := s.repo.GetBooksList(ctx, pagination)
	if err != nil {
		return model.ListBooksRes{}, err
	}

	return model.ListBooksRes{Books: books}, nil
}
