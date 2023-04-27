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

func (s *Store) GetOrders(ctx context.Context, userID, state int) ([]entity.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, ctxTimeout)
	defer cancel()

	query := `SELECT orders.id, orders.UserID, orders.FromCity, orders.ToCity, orders.Date, driver.id, driver.name,driver.phone, car.id, car.model, car.lic_plate 
				FROM orders 	
				INNER JOIN driver ON orders.DriverID = driver.id 
				INNER JOIN car ON orders.CarID = car.id
				WHERE orders.UserID = ? AND orders.State = ?`

	result, err := s.db.Query(query, userID, state)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var orders []entity.Order

	for result.Next() {
		driver := entity.Driver{}
		car := entity.Car{}
		order := entity.Order{}

		err := result.Scan(
			&order.ID,
			&order.UserID,
			&order.From,
			&order.To,
			&order.Date,
			&driver.ID,
			&driver.Name,
			&driver.Phone,
			&car.ID,
			&car.Model,
			&car.LicPlate)
		if err != nil {
			return nil, err
		}
		order.Driver = driver
		order.Car = car
		orders = append(orders, order)
	}

	return orders, nil
}
