package domain

import "errors"

type PersonData struct {
	Name       string
	Age        int
	Gender     *string
	Address    PersonDataAddress
	Dependents []string
}

type PersonDataAddress struct {
	Street     string
	Number     int
	Country    string
	PostalCode int
}

type PersonInput struct {
	Name       string
	Age        int
	Gender     *string
	Address    PersonInputAddress
	Dependents []string
}

type PersonInputAddress struct {
	Street     string
	Number     int
	Country    string
	PostalCode int
}

var ErrInvalidInput = errors.New("Invalid input")

func ValidadeInput(input PersonInput) error {
	if input.Age < 0 {
		panic("Person age can not be negative")
	}

	if input.Address.Number < 0 {
		panic("Address number can not be negative")
	}

	if input.Address.PostalCode < 0 {
		panic("Postal code can not be negative")
	}

	return nil
}
