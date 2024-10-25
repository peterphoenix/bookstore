package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"bookstore/common/respond"
	"bookstore/domain/user/model"
	"github.com/pkg/errors"
)

// CreateUser godoc
// @Summary		Create new user
// @Description	Endpoint used to create new user
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		request body model.CreateUserReq true "body"
// @Success		200 {object} 	respond.APIModel[model.CreateUserRes]
// @Failure		400	{object}	respond.APIModel[string]
// @Failure		500 {object} 	respond.APIModel[string]
// @Router		/user [post]
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var input model.CreateUserReq
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		respond.Error(w, ctx, respond.APIError{
			Code: respond.CodeInvalidRequest,
			Desc: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	err = input.Validate()
	if err != nil {
		respond.Error(w, ctx, respond.APIError{
			Code: respond.CodeInvalidRequest,
			Desc: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	res, err := h.UserSvc.CreateUser(ctx, input)
	if err != nil {
		if strings.Contains(err.Error(), "user_email_key") {
			err = errors.New("email already exist")
			respond.Error(w, ctx, respond.APIError{
				Code: respond.CodeInvalidRequest,
				Desc: err.Error(),
			}, http.StatusBadRequest)
			return
		}

		respond.Error(w, ctx, respond.APIError{
			Code: respond.CodeInternalServerError,
			Desc: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	respond.Success(w, res, http.StatusCreated)
}
