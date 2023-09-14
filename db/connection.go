package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBstring = "host=127.0.1.1 port=5432 user=postgres password=kevs dbname=api_go_auth "
var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DBstring), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("conexion exitosa")
	}
}
