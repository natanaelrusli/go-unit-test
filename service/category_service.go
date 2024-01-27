package service

import (
	"errors"

	"github.com/natanaelrusli/go-unit-test/entity"
	"github.com/natanaelrusli/go-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service *CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
			return nil, errors.New("category not found")
	}
	return category, nil
}
