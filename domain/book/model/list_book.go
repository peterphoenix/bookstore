package model

import (
	commonmodel "bookstore/domain/common/model"
)

type ListBooksRes struct {
	Books []commonmodel.Book `json:"books"`
}
