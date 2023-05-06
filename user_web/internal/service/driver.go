package service

import (
	"context"
	"taxi/internal/entity"
)

type DriverActions interface {
	DriverAuth(ctx context.Context, phone, password string) (entity.Driver, error)
	DriverOrders(ctx context.Context, driverID int, date string) ([]entity.Order, error)
	ChangeStatus(ctx context.Context, id int) error
}

func (s *Service) DriverAuth(ctx context.Context, phone, password string) (entity.Driver, error) {
	s.logger.Info("Driver Auth started")

	driver, err := s.store.DriverAuth(ctx, phone, password)
	if err != nil {
		s.logger.Errorf("Error, while DriverAuth: %s", err)
		return entity.Driver{}, err
	}

	s.logger.Info(driver)
	s.logger.Info("Driver Auth ended")

	return driver, nil
}

func (s *Service) DriverOrders(ctx context.Context, driverID int, date string) ([]entity.Order, error) {
	s.logger.Info("DriverOrders started")

	orders, err := s.store.DriverOrders(ctx, driverID, date)
	if err != nil {
		s.logger.Errorf("Error, while DriverOrders: %s", err)
		return nil, err
	}

	s.logger.Info(orders)
	s.logger.Info("DriverOrders ended")

	return orders, nil
}

func (s *Service) ChangeStatus(ctx context.Context, id int) error {
	s.logger.Info("ChangeStatus started")

	err := s.store.ChangeStatus(ctx, id)
	if err != nil {
		s.logger.Errorf("Error, while ChangeStatus: %s", err)
		return err
	}
	s.logger.Info("ChangeStatus ended")

	return nil
}
