package user

import (
	"context"

	"bookstore/domain/user/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) CreateUser(ctx context.Context, req model.UserCreateReq) (model.UserCreateRes, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.UserCreateRes{}, err
	}
	id := uuid.NewString()

	err = s.repo.CreateUser(ctx, model.CreateUserInput{
		ID:             id,
		Email:          req.Email,
		Name:           req.Name,
		HashedPassword: string(hash),
		PhoneNumber:    req.PhoneNumber,
		Address:        req.Address,
		City:           req.City,
		State:          req.State,
		ZipCode:        req.ZipCode,
		Country:        req.Country,
	})
	if err != nil {
		return model.UserCreateRes{}, err
	}

	return model.UserCreateRes{
		ID:    id,
		Email: req.Email,
		Name:  req.Name,
	}, nil
}
