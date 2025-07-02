package usecases

import (
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/repositories"
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/usecases"
	"github.com/eduardogr/webser-go/internal/domain"
)

type GetNumberById struct {
	Repository repositories.GetNumberById
}

func NewGetNumberByIdUsecase(repository repositories.NumberRepository) usecases.GetNumberByIdUsecase {
	return &GetNumberById{
		// we received 'repositories.NumberRepository'
		// but we will have visibility only for 'repositories.GetNumberById'
		Repository: repository,
	}
}

func (u *GetNumberById) Execute(id int) (domain.Number, error) {
	return u.Repository.Get(id)
}
