package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kushal-png/golang-CRUD-postgres-gorm/initializers"
)

var server *gin.Engine

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load environment variables")
	}

	initializers.ConnectDB(&config)
	server= gin.Default()
}

func main(){
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load environment variables")
	}

	router:= server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context){
		message:="welcome"
		ctx.JSON(http.StatusOK,gin.H{"status":"success", "message":message})
	})

	log.Fatal(server.Run(":"+config.ServerPort))
}