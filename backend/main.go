package main

import (
	"context"
	"log"

	constant "github.com/2110366-2566-2/Mai-Roi-Ra/backend/constant"
	db "github.com/2110366-2566-2/Mai-Roi-Ra/backend/pkg/db"
	services "github.com/2110366-2566-2/Mai-Roi-Ra/backend/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	dbMongoClient, err := db.InitMongoDB()
	if err != nil {
		defer dbMongoClient.Disconnect(context.Background())
		log.Fatal("Error initializing MongoDB:", err)
	}

	r.GET("/", services.GetTest)

	r.Run(":" + constant.PORT)
}
