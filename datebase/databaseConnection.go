package datebase

import (
	"fmt"
	_ "github.com/jackc/pgx/v5"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

var (
	host     = os.Getenv("HOST")
	port     = os.Getenv("ADDR")
	user     = os.Getenv("USER")
	password = os.Getenv("PASSWORD")
	dbName   = os.Getenv("DBNAME")
)

func DBConnection() *gorm.DB {
	db, err := gorm.Open("pgx", "host=%s port=%s dbname=%s user=%s password=%s", host, port, user, password, dbName)
	if err != nil {
		log.Fatalf("Unable to connect Database %v \n", err)
	}
	defer db.Close()
	fmt.Println("Connected to postgres..!!!!")
	return db
}
