package domain

type Repository interface {
	CreatePerson(input PersonData) error
}

type Application interface {
	CreatePerson(input PersonInput) error
}
