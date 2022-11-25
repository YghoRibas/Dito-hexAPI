package app

import "github.com/YghoRibas/Dito-hexAPI/src/domain"

func (a Application) CreatePerson(input domain.PersonInput) error {
	if err := domain.ValidadeInput(input); err != nil {
		return err
	}
	return nil
}
