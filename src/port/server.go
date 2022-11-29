package port

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YghoRibas/Dito-hexAPI/src/domain"
	"github.com/gorilla/mux"
)

type Server struct {
	App domain.Application
}

func (s Server) Run() {
	rotas := mux.NewRouter().StrictSlash(true)

	rotas.HandleFunc("/", GetPerson).Methods("GET")
	rotas.HandleFunc("/persons", s.PostPerson).Methods("POST")
	var port = ":3000"
	fmt.Println("Server running in port", port)
	log.Fatal(http.ListenAndServe(port, rotas))
}
