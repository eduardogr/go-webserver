package usecases

import (
	"errors"

	"github.com/eduardogr/webser-go/internal/adapters/interfaces/repositories"
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/usecases"
	"github.com/eduardogr/webser-go/internal/domain"
)

type CreateNewNumber struct {
	Repository repositories.CreateNumberRepository
}

/*
In 'NewCreateNumberUsecase' is the key to force the implementation of the interface

'CreateNewNumber' must implement this interface:

	type CreateNewNumberUsecase interface {
		Execute(n domain.NewNumberRequest) error
	}
*/
func NewCreateNumberUsecase(repository repositories.NumberRepository) usecases.CreateNewNumberUsecase {
	return &CreateNewNumber{
		// we received 'repositories.NumberRepository'
		// but we will have visibility only for 'repositories.CreateNumberRepository'
		Repository: repository,
	}
}

func (u *CreateNewNumber) Execute(n domain.NewNumberRequest) error {
	// validations
	if n.ID < 0 {
		return errors.New("numbers service - negative numbers not allowed")
	}

	err := u.Repository.Create(n)
	if err != nil {
		return err
	}

	return nil
}
