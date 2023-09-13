package main

import (
	"api/auth/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//acá se cre un objeto ruta del modulo mux
	router := mux.NewRouter()

	// se crean las primeras rutas
	//la funcion handlefunc lo que hace es recibir dos parametros
	//el primero es la ruta a la cual se va a dirigir
	//el segundo recibe la funcion de lo que va a responder
	//responde con una funcion
	//la funcion es como una funcion flecha y recibe 2 parametros
	//request y response, del modulo http, el de toda la vida
	//ahora mandamos por el response un .write
	router.HandleFunc("/", routes.HelloHome)

	//inicializamos el servidor
	//recibe el puerto y el router inicializador
	http.ListenAndServe(":3000", router)
}
