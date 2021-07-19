package infrastructure

import (
	"os"

	"github.com/jmoiron/sqlx"
	// To enable sqlx to connect MySQL.
	_ "github.com/go-sql-driver/mysql"
)

// DB exposes database connection.
var DB *sqlx.DB

func init() {
	DB = connectDB()
}

func connectDB() *sqlx.DB {
	username := os.Getenv("AMADEUS_DATABASE_USERNAME")
	password := os.Getenv("AMADEUS_DATABASE_PASSWORD")
	host := os.Getenv("AMADEUS_DATABASE_HOST")
	database := os.Getenv("AMADEUS_DATABASE")

	dbConfigStr := username + ":" + password + "@tcp(" + host + ":3306)/" + database + "?parseTime=true"
	db, err := sqlx.Connect("mysql", dbConfigStr)
	if err != nil {
		panic(err)
	}

	return db
}
