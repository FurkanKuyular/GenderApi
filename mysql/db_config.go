package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "gorm.io/driver/mysql"
	"log"
)

func main() {
	_ = godotenv.Load(".env")
	var myEnv map[string]string
	myEnv, _ = godotenv.Read()

	db, err := sql.Open("mysql", fmt.Sprintf("%s/", myEnv["DATABASE_DSN"]))

	defer func(db *sql.DB) {
		err := db.Close()

		if err != nil {
			log.Fatal(fmt.Sprintf("%s: %s", "Something went wrong", err))
		}
	}(db)

	_, err = db.Exec(fmt.Sprintf("%s %s", "CREATE DATABASE IF NOT EXISTS", myEnv["DATABASE_NAME"]))

	err = db.Close()
	if err != nil {
		return
	}

	fmt.Println("Database created successfully.")
}
