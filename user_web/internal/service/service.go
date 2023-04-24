package service

import "github.com/sirupsen/logrus"

type UserService interface {
	UserAuth
	UserActions
	UserOrders
}

type Service struct {
	logger *logrus.Logger
	store  UserService
}

func NewService(logger *logrus.Logger, store UserService) *Service {
	return &Service{
		logger: logger,
		store:  store,
	}
}
