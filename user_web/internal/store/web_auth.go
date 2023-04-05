package store

import (
	"context"
	"taxi/internal/entity"
)

func (s *Store) SignIn(ctx context.Context, phone, password string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, ctxTimeout)
	defer cancel()

	user := entity.User{}
	query := `SELECT id,name,phone,password,email FROM user WHERE phone=? AND password=?`
	err := s.db.QueryRowContext(ctx, query, phone, password).
		Scan(&user.ID,
			&user.Name,
			&user.Phone,
			&user.Password,
			&user.Email)

	return &user, err
}

func (s *Store) SignUp(ctx context.Context, user entity.User) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, ctxTimeout)
	defer cancel()

	var id int

	insertQuery := `INSERT INTO user(name,phone,password,email) VALUES(?,?,?,?)`
	_, err := s.db.ExecContext(ctx, insertQuery,
		user.Name,
		user.Phone,
		user.Password,
		user.Email)
	if err != nil {
		return 0, err
	}

	selectQuery := `SELECT LAST_INSERT_ID() AS id`
	err = s.db.QueryRowContext(ctx, selectQuery).Scan(&id)

	return id, err
}
