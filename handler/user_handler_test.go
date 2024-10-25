package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"bookstore/common/respond"
	"bookstore/domain/user/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_CreateUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testHandler, mocks := newTestHandler(mockCtrl)

	dummyError := "dummy error"
	getInput := func() model.CreateUserReq {
		return model.CreateUserReq{
			Email:       "email@example.com",
			Name:        "name",
			Password:    "password",
			PhoneNumber: "1234567890",
		}
	}

	res := model.CreateUserRes{
		ID:    "123",
		Email: "email@example.com",
		Name:  "name",
	}

	t.Run("error decode input", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/user", nil)
		testHandler.CreateUser(w, r)

		var resp respond.APIModel[model.CreateUserRes]
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, respond.CodeInvalidRequest, resp.Errors[0].Code)
		assert.NotEmpty(t, resp.Errors[0].Message)
	})

	t.Run("error invalid email", func(t *testing.T) {
		input := getInput()
		input.Email = "invalid email"
		inputJSON, err := json.Marshal(input)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(inputJSON))
		testHandler.CreateUser(w, r)

		var resp respond.APIModel[model.CreateUserRes]
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, respond.CodeInvalidRequest, resp.Errors[0].Code)
		assert.Equal(t, model.InvalidEmailErr, resp.Errors[0].Message)
	})

	t.Run("error invalid password", func(t *testing.T) {
		input := getInput()
		input.Password = ""
		inputJSON, err := json.Marshal(input)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(inputJSON))
		testHandler.CreateUser(w, r)

		var resp respond.APIModel[model.CreateUserRes]
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, respond.CodeInvalidRequest, resp.Errors[0].Code)
		assert.Equal(t, model.InvalidPasswordErr, resp.Errors[0].Message)
	})

	t.Run("error invalid phone", func(t *testing.T) {
		input := getInput()
		input.PhoneNumber = "123"
		inputJSON, err := json.Marshal(input)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(inputJSON))
		testHandler.CreateUser(w, r)

		var resp respond.APIModel[model.CreateUserRes]
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, respond.CodeInvalidRequest, resp.Errors[0].Code)
		assert.Equal(t, model.InvalidPhoneErr, resp.Errors[0].Message)
	})

	t.Run("error create user", func(t *testing.T) {
		input := getInput()
		inputJSON, err := json.Marshal(input)
		assert.NoError(t, err)

		mocks.userMock.EXPECT().CreateUser(gomock.Any(), gomock.Any()).
			Return(model.CreateUserRes{}, fmt.Errorf(dummyError))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(inputJSON))
		testHandler.CreateUser(w, r)

		var resp respond.APIModel[model.CreateUserRes]
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, respond.CodeInternalServerError, resp.Errors[0].Code)
		assert.Equal(t, dummyError, resp.Errors[0].Message)
	})

	t.Run("error email already exists", func(t *testing.T) {
		input := getInput()
		inputJSON, err := json.Marshal(input)
		assert.NoError(t, err)

		mocks.userMock.EXPECT().CreateUser(gomock.Any(), gomock.Any()).
			Return(model.CreateUserRes{}, fmt.Errorf("violate fkey constraint user_email_key"))

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(inputJSON))
		testHandler.CreateUser(w, r)

		var resp respond.APIModel[model.CreateUserRes]
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, respond.CodeInvalidRequest, resp.Errors[0].Code)
		assert.Equal(t, "email already exist", resp.Errors[0].Message)
	})

	t.Run("success", func(t *testing.T) {
		input := getInput()
		inputJSON, err := json.Marshal(input)
		assert.NoError(t, err)

		mocks.userMock.EXPECT().CreateUser(gomock.Any(), gomock.Any()).
			Return(res, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(inputJSON))
		testHandler.CreateUser(w, r)

		var resp respond.APIModel[model.CreateUserRes]
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Empty(t, resp.Errors)
		assert.Equal(t, res, resp.Data)
	})
}
