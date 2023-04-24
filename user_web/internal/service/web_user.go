package service

import (
	"context"
	"taxi/internal/entity"
)

type UserActions interface {
	RetrieveByID(ctx context.Context, id int) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id int) error
	GetOrders(ctx context.Context, userID, state int) ([]entity.Order, error)
}

func (s *Service) RetrieveByID(ctx context.Context, id int) (*entity.User, error) {
	s.logger.Info("RetrieveByID started")

	user, err := s.store.RetrieveByID(ctx, id)
	if err != nil {
		s.logger.Errorf("Error, while RetrieveByID, %s", err)
		return nil, err
	}

	s.logger.Info(user)
	s.logger.Info("RetrieveByID ended")

	return user, nil
}

func (s *Service) UpdateUser(ctx context.Context, user *entity.User) error {
	s.logger.Info("Update User started")

	err := s.store.UpdateUser(ctx, user)
	if err != nil {
		s.logger.Errorf("Error, while UpdateUser, %s", err)
		return err
	}

	s.logger.Info("Update User ended")
	return nil
}

func (s *Service) DeleteUser(ctx context.Context, id int) error {
	s.logger.Info("Delete User started")

	err := s.store.DeleteUser(ctx, id)
	if err != nil {
		s.logger.Errorf("Error, while DeleteUser, %s", err)
		return err
	}

	s.logger.Info("Delete User ended")
	return nil
}

func (s *Service) GetOrders(ctx context.Context, userID, state int) ([]entity.Order, error) {
	s.logger.Info("Get Orders started")

	orders, err := s.store.GetOrders(ctx, userID, state)
	if err != nil {
		s.logger.Errorf("Error, while GetOrders, %s", err)
		return nil, err
	}

	s.logger.Info(orders)
	s.logger.Info("GetOrders ended")

	return orders, nil
}
