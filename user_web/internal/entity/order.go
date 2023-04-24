package entity

type Order struct {
	ID     int
	UserID int
	Driver Driver
	Car    Car
	From   string
	To     string
	Date   string
	State  int
}
