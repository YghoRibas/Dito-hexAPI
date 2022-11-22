package domain

type Repository interface {
	CreatePerson(input PersonData) error
}