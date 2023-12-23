package store

import (
	"context"
	"taxi/internal/entity"
)

func (s *Store) Order(ctx context.Context, order *entity.Order) error {
	ctx, cancel := context.WithTimeout(ctx, ctxTimeout)
	defer cancel()
	//
	query := `CALL AddOrder(?, ?, ?, ?, ?, ?);`
	_, err := s.db.ExecContext(ctx, query,
		order.User.ID,
		order.From,
		order.To,
		order.Date,
		order.Time,
		order.Tickets)

	return err
}

func (s *Store) CancelOrder(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, ctxTimeout)
	defer cancel()

	query := `CALL CancelOrder(?)`
	_, err := s.db.ExecContext(ctx, query, id)

	return err
}
