package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // don't forget to add it. It doesn't be added automatically
)

type User struct {
	Username string
	Password string
}

var Db *sql.DB //created outside to make it global.

// make sure your function start with uppercase to call outside of the directory.

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error occured on .env file please check")
	}

	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("password")

	psqlsetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, pass)

	db, errSql := sql.Open("postgres", psqlsetup)
	if errSql != nil {
		fmt.Println("Error ocuured in the database connection")
		panic(errSql)
	} else {
		Db = db
		fmt.Println("Successfully connected to the database")

	}

}
