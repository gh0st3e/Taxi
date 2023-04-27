package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"taxi/internal/db"
	"taxi/internal/handler"
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
	taxiHandler := handler.NewHandler(logger, taxiService)
	server := gin.New()
	taxiHandler.Mount(server)

	err = server.Run(":9000")
	if err != nil {
		logger.Fatalf("Couldn't run server: %s", err)
	}
}
