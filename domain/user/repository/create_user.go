package repository

import (
	"context"

	"bookstore/domain/user/model"
	sq "github.com/Masterminds/squirrel"
)

func (r repository) CreateUser(ctx context.Context, input model.CreateUserInput) error {
	_, err := sq.Insert(userTableName).
		Columns(
			userTable.ID,
			userTable.Name,
			userTable.Email,
			userTable.Password,
			userTable.PhoneNumber,
			userTable.Address,
			userTable.City,
			userTable.State,
			userTable.ZipCode,
			userTable.Country,
		).
		Values(
			input.ID,
			input.Name,
			input.Email,
			input.HashedPassword,
			input.PhoneNumber,
			input.Address,
			input.City,
			input.State,
			input.ZipCode,
			input.Country,
		).
		PlaceholderFormat(sq.Dollar).
		RunWith(r.psql.DB()).
		ExecContext(ctx)

	return err
}
