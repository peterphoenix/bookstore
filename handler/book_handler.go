package handler

import (
	"net/http"

	"bookstore/common/pagination"
	"bookstore/common/respond"
	_ "bookstore/domain/book/model"
)

// ListBooks godoc
//
// @Summary		Get list books
// @Description	Endpoint used to get list of available books
// @Tags		Book
// @Accept		json
// @Produce		json
// @Param 		limit query string false "limit per page"
// @Param 		page query string false "page number"
// @Success		200		{object}	respond.APIModel[model.ListBooksRes]
// @Failure		500		{object}	respond.APIModel[string]
// @Router		/book [get]
func (h *handler) ListBooks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	res, err := h.BookSvc.ListBooks(ctx, pagination.GetPagination(r))
	if err != nil {
		respond.Error(w, ctx, respond.APIError{
			Code: respond.CodeInternalServerError,
			Desc: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	respond.Success(w, res, http.StatusOK)
}
