package user

//go:generate mockgen -source ./service.go -destination ../../test/mock/domain/user/service.go -package=user_mock

import (
	"context"

	"bookstore/domain/user/model"
	"bookstore/domain/user/repository"
)

type Service interface {
	CreateUser(ctx context.Context, req model.CreateUserReq) (model.CreateUserRes, error)
}

type service struct {
	repo repository.Repository
}

func NewService(
	repo repository.Repository,
) Service {
	return &service{
		repo: repo,
	}
}
