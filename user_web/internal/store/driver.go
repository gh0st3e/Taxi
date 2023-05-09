package store

import (
	"context"
	"taxi/internal/entity"
)

func (s *Store) DriverAuth(ctx context.Context, phone, password string) (entity.Driver, error) {
	localCtx, cancel := context.WithTimeout(ctx, ctxTimeout)
	defer cancel()

	query := `CALL DriverAuth(?,?)`

	driver := entity.Driver{}

	err := s.db.QueryRowContext(localCtx, query, phone, password).
		Scan(
			&driver.ID,
			&driver.Name,
			&driver.Phone,
			&driver.Password,
			&driver.WorkExp)

	return driver, err
}

func (s *Store) DriverOrders(ctx context.Context, driverID int, date string) ([]entity.Order, error) {
	localCtx, cancel := context.WithTimeout(ctx, ctxTimeout)
	defer cancel()

	query := `CALL GetDriverOrders(?,?)`

	result, err := s.db.QueryContext(localCtx, query, driverID, date)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var orders []entity.Order

	for result.Next() {
		order := entity.Order{}
		user := entity.User{}

		err := result.Scan(
			&order.ID,
			&order.From,
			&order.To,
			&order.Date,
			&order.Time,
			&order.Tickets,
			&order.State,
			&user.ID,
			&user.Name,
			&user.Phone)
		if err != nil {
			return nil, err
		}
		order.User = user
		orders = append(orders, order)
	}

	return orders, nil
}

func (s *Store) ChangeStatus(ctx context.Context, id, state int) error {
	localCtx, cancel := context.WithTimeout(ctx, ctxTimeout)
	defer cancel()

	query := `CALL UpdatePassengerStatus(?,?)`

	_, err := s.db.ExecContext(localCtx, query, state, id)

	return err

}
