package app

func (a Application) CreatePerson() error {
	a.Repo.CreatePerson(input)
}

