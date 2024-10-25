package user

import (
	"context"
	"fmt"
	"testing"

	"bookstore/domain/user/model"
	user_mock "bookstore/test/mock/domain/user"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestService_CreateUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := user_mock.NewMockRepository(mockCtrl)

	svc := NewService(mockRepo)

	ctx := context.TODO()
	dummyError := "dummy error"

	t.Run("error hash password", func(t *testing.T) {
		res, err := svc.CreateUser(ctx, model.CreateUserReq{
			Password: "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
		})

		assert.Empty(nil, res)
		assert.EqualError(t, err, bcrypt.ErrPasswordTooLong.Error())
	})

	t.Run("error create user", func(t *testing.T) {
		mockRepo.EXPECT().CreateUser(ctx, gomock.Any()).
			Return(fmt.Errorf(dummyError))

		res, err := svc.CreateUser(ctx, model.CreateUserReq{})

		assert.Empty(nil, res)
		assert.EqualError(t, err, dummyError)
	})

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().CreateUser(ctx, gomock.Any()).
			Return(nil)

		res, err := svc.CreateUser(ctx, model.CreateUserReq{
			Email: "test@test.com",
			Name:  "name",
		})

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		assert.Equal(t, "test@test.com", res.Email)
		assert.Equal(t, "name", res.Name)
	})
}
