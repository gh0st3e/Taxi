package store

import (
	"context"
	"taxi/internal/entity"
)

func (s *Store) RetrieveByID(ctx context.Context, id int) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, ctxTimeout)
	defer cancel()

	user := entity.User{}

	query := `SELECT id,name,phone,password,email FROM user WHERE id=?`
	err := s.db.QueryRowContext(ctx, query, id).
		Scan(&user.ID,
			&user.Name,
			&user.Phone,
			&user.Password,
			&user.Email)

	return &user, err
}

func (s *Store) UpdateUser(ctx context.Context, user *entity.User) error {
	ctx, cancel := context.WithTimeout(ctx, ctxTimeout)
	defer cancel()

	query := `UPDATE user SET name=?,password=?,email=? WHERE id = ?`
	_, err := s.db.ExecContext(ctx, query, user.Name, user.Password, user.Email, user.ID)

	return err
}

func (s *Store) DeleteUser(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, ctxTimeout)
	defer cancel()

	query := `DELETE FROM user WHERE id=?`
	_, err := s.db.ExecContext(ctx, query, id)

	return err
}
