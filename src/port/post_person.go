package port

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type PersonDataInput struct {
	Name string `json:"name"`
	Age	int		 `json:"age"`
	Gender	*string `json:"gender"`
	Address PersonDataInputAddress `json:"address"`
	Dependents []string  `json:"dependents"`
}

type PersonDataInputAddress struct {
	Street string `json:"street"`
	Number int `json:"number"`
	Country string `json:"country"`
	PostalCode int `json:"postal_code"`
}

//PostPerson...
func PostPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var inputPerson PersonDataInput

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &inputPerson); err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	json.Unmarshal(body, &inputPerson)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(inputPerson); err != nil {
		panic(err)
	}
}