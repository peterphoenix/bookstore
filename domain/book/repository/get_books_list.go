package repository

import (
	"context"
	"database/sql"
	"fmt"

	"bookstore/common/pagination"
	commonmodel "bookstore/domain/common/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func (r *repository) GetBooksList(ctx context.Context, pagination pagination.Pagination) ([]commonmodel.Book, error) {
	var books []commonmodel.Book
	query := sq.
		Select(
			bookTableName+"."+bookTable.ID,
			r.psql.ArrayAggregate(authorTable.Name, "authors"),
			bookTable.Title,
			bookTable.Price,
			bookTable.ImageURL,
			bookTable.Description,
		).
		From(bookTableName).
		Join(
			fmt.Sprintf("%s ON %s = %s",
				bookAuthorTableName,
				bookTableName+"."+bookTable.ID,
				bookAuthorTableName+"."+bookAuthorTable.BookID),
		).
		Join(
			fmt.Sprintf("%s ON %s = %s",
				authorTableName,
				bookAuthorTableName+"."+bookAuthorTable.AuthorID,
				authorTableName+"."+authorTable.ID),
		).
		Offset((pagination.Page - 1) * pagination.Limit).
		Limit(pagination.Limit).
		GroupBy(bookTableName + "." + bookTable.ID)

	res, err := query.RunWith(r.psql.DB()).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var book commonmodel.Book
		var imageURL, description sql.NullString
		err = res.Scan(
			&book.ID,
			pq.Array(&book.Authors),
			&book.Title,
			&book.Price,
			&imageURL,
			&description,
		)
		if err != nil {
			return nil, err
		}

		book.ImageURL = imageURL.String
		book.Description = description.String

		books = append(books, book)
	}

	return books, nil
}
