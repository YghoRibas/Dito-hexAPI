package port

import (
	"net/http"
	"encoding/json"
)

type PersonDataOutput struct {
	Name string `json:"name"`
	Age	int		 `json:"age"`
	Gender	*string `json:"gender"`
	Address PersonDataOutputAddress `json:"address"`
	Dependents []string  `json:"dependents"`
}

type PersonDataOutputAddress struct {
	Street string `json:"street"`
	Number int `json:"number"`
	Country string `json:"country"`
	PostalCode int `json:"postal_code"`
}

// GetPerson...
func GetPerson(w http.ResponseWriter, r *http.Request) {
	var outputPerson PersonDataOutput

	json.NewEncoder(w).Encode(outputPerson)
}