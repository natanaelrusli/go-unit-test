package repository

import "github.com/natanaelrusli/go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
