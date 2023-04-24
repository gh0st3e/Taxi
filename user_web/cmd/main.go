package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"taxi/internal/db"
	"taxi/internal/entity"
	"taxi/internal/service"
	"taxi/internal/store"
)

func main() {

	logger := logrus.New()

	db, err := db.MySqlConnect()
	if err != nil {
		panic(err)
	}

	taxiStore := store.NewStore(db)
	taxiService := service.NewService(logger, taxiStore)

	user := entity.User{
		ID:       1,
		Name:     "Denis",
		Phone:    "",
		Password: "",
		Email:    "",
	}

	fmt.Println(taxiService.UpdateUser(context.TODO(), &user))
}
