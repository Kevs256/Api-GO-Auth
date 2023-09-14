package main

import (
	"api/auth/db"
	"api/auth/models"
	"api/auth/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//vamos a conectar con la base de datos
	db.DBConnection()

	//vamos a crear las tablas
	db.DB.AutoMigrate(models.UserAuth{})

	//acá se cre un objeto ruta del modulo mux
	router := mux.NewRouter()

	// se crean las primeras rutas
	//la funcion handlefunc lo que hace es recibir dos parametros
	//el primero es la ruta a la cual se va a dirigir
	//el segundo recibe la funcion de lo que va a responder
	//responde con una funcion
	router.HandleFunc("/", routes.Test).Methods("GET")
	router.HandleFunc("/Email", routes.VerfyEmail).Methods("GET")
	router.HandleFunc("/User", routes.Register).Methods("POST")
	router.HandleFunc("/User", routes.Loggin).Methods("GET")
	router.HandleFunc("/Auth", routes.Auth).Methods("GET")

	//inicializamos el servidor
	//recibe el puerto y el router inicializador
	http.ListenAndServe(":3000", router)
}
