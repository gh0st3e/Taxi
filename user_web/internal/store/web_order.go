package store

import (
	"context"
	"taxi/internal/entity"
)

func (s *Store) Order(ctx context.Context, order *entity.Order) error {
	ctx, cancel := context.WithTimeout(ctx, ctxTimeout)
	defer cancel()

	query := `INSERT INTO orders(UserID,FromCity,ToCity,Date) VALUES (?,?,?,?)`
	_, err := s.db.ExecContext(ctx, query,
		order.UserID,
		order.From,
		order.To,
		order.Date)

	return err
}

func (s *Store) CancelOrder(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, ctxTimeout)
	defer cancel()

	query := `UPDATE orders SET State=-1 WHERE id=?`
	_, err := s.db.ExecContext(ctx, query, id)

	return err
}
