package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	var err error
	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("DB_URL"),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed!")
	}

	fmt.Println(DB)
}
