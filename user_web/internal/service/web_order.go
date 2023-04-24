package service

import (
	"context"
	"taxi/internal/entity"
)

type UserOrders interface {
	Order(ctx context.Context, order *entity.Order) error
	CancelOrder(ctx context.Context, id int) error
}

func (s *Service) Order(ctx context.Context, order *entity.Order) error {
	s.logger.Info("Order started")

	err := s.store.Order(ctx, order)
	if err != nil {
		s.logger.Errorf("Error, while Order: %s", err)
		return err
	}

	s.logger.Info("Order ended")

	return nil
}

func (s *Service) CancelOrder(ctx context.Context, id int) error {
	s.logger.Info("CancelOrder started")

	err := s.store.CancelOrder(ctx, id)
	if err != nil {
		s.logger.Errorf("Error, while CancelOrder: %s", err)
		return err
	}

	s.logger.Info("CancelOrder ended")

	return nil
}
