package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "rian:Cash6666@/go_products")
	fmt.Print("Db connected!!!!! ")

	if err != nil {
		panic(err)
	}

	DB = db

	log.Println("DB Connected!!")

}
