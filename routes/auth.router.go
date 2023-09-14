package routes

import (
	"api/auth/db"
	"api/auth/models"
	"encoding/json"
	"net/http"
)

// la funcion es como una funcion flecha y recibe 2 parametros
// request y response, del modulo http, el de toda la vida
// ahora mandamos por el response un .write
func Test(reponse http.ResponseWriter, request *http.Request) {
	reponse.Write([]byte("Hello World"))
}

func VerfyEmail(reponse http.ResponseWriter, request *http.Request) {}

func Register(reponse http.ResponseWriter, request *http.Request) {

	var user models.UserAuth
	json.NewDecoder(request.Body).Decode(&user)

	if user.Email == "" || user.Password == "" {
		reponse.WriteHeader(http.StatusBadRequest)
		reponse.Write([]byte("ingrese email y contraseña"))
		return
	}

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		reponse.WriteHeader(http.StatusBadRequest)
		reponse.Write([]byte(err.Error()))
	} else {
		reponse.WriteHeader(http.StatusCreated)
		reponse.Write([]byte("Creado correctamente"))
	}

	res := struct {
		UserID  int    `json:"user_id"`
		Message string `json:"message"`
	}{
		UserID:  int(user.ID),
		Message: "Registro exitoso",
	}

	json.NewEncoder(reponse).Encode(&res)
}

func Loggin(reponse http.ResponseWriter, request *http.Request) {}

func Auth(reponse http.ResponseWriter, request *http.Request) {}

//func RestorePassword() {}
