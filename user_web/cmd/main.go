package main

import (
	"context"
	"taxi/internal/db"
	"taxi/internal/store"
)

func main() {
	db, err := db.MySqlConnect()
	if err != nil {
		panic(err)
	}

	taxiStore := store.NewStore(db)

	err = taxiStore.CancelOrder(context.TODO(), 1)
	if err != nil {
		panic(err)
	}

	//fmt.Println(user)
}
