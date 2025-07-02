package usecases

import (
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/repositories"
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/usecases"
	"github.com/eduardogr/webser-go/internal/domain"
)

type GetAllNumbers struct {
	Repository repositories.GetAllNumbersRepository
}

func NewGetAllNumbersUsecase(repository repositories.NumberRepository) usecases.GetAllNumbersUsecase {
	return &GetAllNumbers{
		// we received 'repositories.NumberRepository'
		// but we will have visibility only for 'repositories.GetAllNumbersRepository'
		Repository: repository,
	}
}

func (u *GetAllNumbers) Execute() ([]domain.Number, error) {
	return u.Repository.GetAll()
}
