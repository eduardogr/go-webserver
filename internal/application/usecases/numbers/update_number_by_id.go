package usecases

import (
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/repositories"
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/usecases"
	"github.com/eduardogr/webser-go/internal/domain"
)

type UpdateNumberById struct {
	Repository repositories.UpdateNumberById
}

func NewUpdateNumberByIdUsecase(repository repositories.NumberRepository) usecases.UpdateNumberByIdUsecase {
	return &UpdateNumberById{
		// we received 'repositories.NumberRepository'
		// but we will have visibility only for 'repositories.UpdateNumberById'
		Repository: repository,
	}
}

func (u *UpdateNumberById) Execute(n domain.Number, id int) error {
	return u.Repository.Update(n, id)
}
