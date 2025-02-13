package initializers

import (
	"log"
	"os"

	"github.com/DevAthhh/trainer-ai/server/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func SyncDatabase() {
	DB.AutoMigrate(&models.UserTrainers{})
}
