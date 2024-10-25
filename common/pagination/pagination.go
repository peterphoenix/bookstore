package pagination

import (
	"net/http"
	"strconv"
)

const (
	DefaultLimit = 10
	DefaultPage  = 1
)

type Pagination struct {
	Limit uint64 `json:"limit"`
	Page  uint64 `json:"page"`
}

func GetPagination(r *http.Request) Pagination {
	page := DefaultPage
	limit := DefaultLimit

	if p := r.URL.Query().Get("page"); p != "" {
		parsedPage, err := strconv.Atoi(p)
		if err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	if l := r.URL.Query().Get("limit"); l != "" {
		parsedLimit, err := strconv.Atoi(l)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	return Pagination{
		Limit: uint64(limit),
		Page:  uint64(page),
	}
}
