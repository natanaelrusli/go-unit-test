package service

import (
	"testing"

	"github.com/natanaelrusli/go-unit-test/entity"
	mocks "github.com/natanaelrusli/go-unit-test/mock"
	"github.com/natanaelrusli/go-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: &categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")

	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestCategoryService_GetFound(t *testing.T) {
	repo := &mocks.CategoryRepository{}
	categoryService := CategoryService{Repository: repo}

	mockCategory := entity.Category{Id: "1", Name: "Abc"}

	repo.On("FindById", "1").Return(&mockCategory)

	res, err := categoryService.Get("1")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, &mockCategory, res)
}
