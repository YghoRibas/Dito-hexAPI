package port

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/YghoRibas/Dito-hexAPI/src/domain"
)

type PersonDataInput struct {
	Name       string                 `json:"name"`
	Age        int                    `json:"age"`
	Gender     *string                `json:"gender"`
	Address    PersonDataInputAddress `json:"address"`
	Dependents []string               `json:"dependents"`
}

type PersonDataInputAddress struct {
	Street     string `json:"street"`
	Number     int    `json:"number"`
	Country    string `json:"country"`
	PostalCode int    `json:"postal_code"`
}

// PostPerson...
func (s Server) PostPerson(w http.ResponseWriter, r *http.Request) {
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

	err = s.App.CreatePerson(domain.PersonInput{
		Name:   inputPerson.Name,
		Age:    inputPerson.Age,
		Gender: inputPerson.Gender,
		Address: domain.PersonInputAddress{
			Street:     inputPerson.Address.Street,
			Country:    inputPerson.Address.Country,
			PostalCode: inputPerson.Address.PostalCode,
			Number:     inputPerson.Address.Number,
		},
		Dependents: inputPerson.Dependents,
	})

	if err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(inputPerson); err != nil {
		panic(err)
	}
}
