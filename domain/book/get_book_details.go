package book

import (
	"context"

	"bookstore/domain/book/model"
)

func (s service) GetBookDetailsMap(ctx context.Context, req model.GetBookDetailsReq) (model.BookDetailsMap, error) {
	books, err := s.repo.GetBookDetails(ctx, req.IDs)
	if err != nil {
		return nil, err
	}

	res := make(model.BookDetailsMap)
	for _, b := range books {
		res[b.ID] = b
	}

	return res, nil
}
