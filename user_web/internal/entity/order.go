package entity

type Order struct {
	ID      int
	User    User
	Driver  Driver
	Car     Car
	From    string
	To      string
	Date    string
	Time    string
	Tickets int
	State   int
}

type ChangeStateRequest struct {
	ID    int `json:"ID"`
	State int `json:"State"`
}

type ChangeManyStateRequest struct {
	Orders []ChangeStateRequest `json:"orders"`
}
