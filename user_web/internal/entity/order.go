package entity

type Order struct {
	ID       int
	UserID   int
	DriverID int
	CarID    int
	From     string
	To       string
	Date     string
	State    int
}
