package migration

import (
	"fmt"

	"github.com/AdityaAWP/go-crud/database"
	"github.com/AdityaAWP/go-crud/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		panic("Failed to migrate database!")
	}
	fmt.Println("Database migrated!")
}
