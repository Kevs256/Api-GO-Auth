package routes

import "net/http"

func HelloHome(reponse http.ResponseWriter, request *http.Request) {
	reponse.Write([]byte("Hello World hola"))
}
