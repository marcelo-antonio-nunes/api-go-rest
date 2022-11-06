package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcelo-antonio-nunes/api-go-rest.git/controllers"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidade", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades/x", controllers.CriaUmaNovaPersonalidade).Methods("Get")

	fmt.Println("Iniciando Servidor Rest com Go")
	log.Fatal((http.ListenAndServe(":8080", r)))
}
