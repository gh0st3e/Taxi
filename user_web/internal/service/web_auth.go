package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"taxi/internal/entity"
	"time"
)

const (
	accessTokenTTL  = 12
	refreshTokenTTL = 24
	signingKey      = "ds1jlg3h245wp5oih432v"
)

type UserAuth interface {
	SignIn(ctx context.Context, phone, password string) (*entity.User, error)
	SignUp(ctx context.Context, user entity.User) (int, error)
}

func (s *Service) SignIn(ctx context.Context, phone, password string) (accessToken, refreshToken string, err error) {
	s.logger.Info("SignIn started")

	user, err := s.store.SignIn(ctx, phone, password)
	if err != nil {
		s.logger.Errorf("Error, while SignIn user: %s", err)
		return "", "", fmt.Errorf("error, while SignIn user: %w", err)
	}

	s.logger.Info(user.ID)

	sID := strconv.Itoa(user.ID)

	accessToken, err = NewToken(sID, signingKey, accessTokenTTL*time.Hour)
	if err != nil {
		s.logger.Errorf("Error while creating Access Token: %s", err)
		return "", "", errors.Wrap(err, "service.Login.NewRefreshToken: Couldn't generate refresh token")
	}

	refreshToken, err = NewToken(sID, signingKey, refreshTokenTTL*time.Hour)
	if err != nil {
		s.logger.Errorf("Error while creating Refresh Token: %s", err)
		return "", "", errors.Wrap(err, "service.Login.NewRefreshToken: Couldn't generate refresh token")
	}

	s.logger.Info(accessToken)
	s.logger.Info(refreshToken)

	s.logger.Info("service.user.LoginUser ended")

	return accessToken, refreshToken, nil
}

func (s *Service) SignUp(ctx context.Context, user entity.User) (int, error) {
	s.logger.Info("SignUp started")

	id, err := s.store.SignUp(ctx, user)
	if err != nil {
		s.logger.Errorf("Error, while SignUp: %s", err)
		return 0, err
	}

	s.logger.Info("SignUp ended")

	return id, nil
}
