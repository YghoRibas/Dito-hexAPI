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

func ValidadeInput(input PersonInput) error {
	if input.Age < 0 {
		return errors.New("Invalid age")
	}

	if input.Address.Number < 0 {
		return errors.New("Invalid person address number")
	}

	if input.Address.PostalCode < 0 {
		return errors.New("Invalid person postal code")
	}

	return nil
}
