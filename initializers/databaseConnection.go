package initializers

import (
	"fmt"
	"github.com/Prameesh-P/E-COMMERCE/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func init() {
	Getenv()
}

var DB *gorm.DB

func DBConnection() {
	var err error
	host := os.Getenv("HOST")
	port := os.Getenv("ADDR")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbName, password)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error to connecting to database..!!")
	}
	DB.AutoMigrate(
		&models.User{},
	)

	fmt.Println("Connected to postgres..!!!!")
}
