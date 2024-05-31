package main

import (
	"fmt"

	"github.com/kushal-png/golang-CRUD-postgres-gorm/initializers"
	"github.com/kushal-png/golang-CRUD-postgres-gorm/models"
)

func main() {
	err := initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Error in Migration")
		return
	}
	fmt.Println("Migration Completed")
}
