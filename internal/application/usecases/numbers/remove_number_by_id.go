package usecases

import (
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/repositories"
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/usecases"
)

type RemoveNumberById struct {
	Repository repositories.RemoveNumberById
}

func NewRemoveNumberByIdUsecase(repository repositories.NumberRepository) usecases.RemoveNumberByIdUsecase {
	return &RemoveNumberById{
		// we received 'repositories.NumberRepository'
		// but we will have visibility only for 'repositories.RemoveNumberById'
		Repository: repository,
	}
}

func (u *RemoveNumberById) Execute(id int) error {
	return u.Repository.Remove(id)
}
