package service

import (
	"context"
	"taxi/internal/entity"
)

type UserAuth interface {
	SignIn(ctx context.Context, phone, password string) (*entity.User, error)
	SignUp(ctx context.Context, user entity.User) (int, error)
}

func (s *Service) SignIn(ctx context.Context, phone, password string) error {
	return nil
}

func (s *Service) SignUp(ctx context.Context, user entity.User) error {
	return nil
}
