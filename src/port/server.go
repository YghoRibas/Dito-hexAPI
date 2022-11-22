package port

import (
	"fmt"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

type Server struct {

}

func (s Server) Run() {
	rotas := mux.NewRouter().StrictSlash(true)

	rotas.HandleFunc("/", GetPerson).Methods("GET")
	rotas.HandleFunc("/persons", PostPerson).Methods("POST")
	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, rotas))
}
