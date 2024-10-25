package repository

import (
	"context"
	"database/sql"

	"bookstore/domain/book/model"
	sq "github.com/Masterminds/squirrel"
)

func (r *repository) GetBookDetails(ctx context.Context, id []string) ([]model.BookDetails, error) {
	var bookDetails []model.BookDetails
	res, err := sq.
		Select(
			bookTable.ID,
			bookTable.Title,
			bookTable.Price,
			bookTable.ImageURL,
			bookTable.ISBN,
			bookTable.Description,
		).
		From(bookTableName).
		RunWith(r.psql.DB()).
		QueryContext(ctx)

	if err != nil {
		return nil, err
	}

	defer res.Close()

	for res.Next() {
		var bookDetail model.BookDetails
		var imageURL, description sql.NullString
		err = res.Scan(
			&bookDetail.ID,
			&bookDetail.Title,
			&bookDetail.Price,
			&imageURL,
			&bookDetail.ISBN,
			&description,
		)
		if err != nil {
			return nil, err
		}

		bookDetail.ImageURL = imageURL.String
		bookDetail.Description = description.String

		bookDetails = append(bookDetails, bookDetail)
	}

	return bookDetails, nil
}
