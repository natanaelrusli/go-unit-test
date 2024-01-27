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
	// Initialize the mock repository
	repo := &mocks.CategoryRepository{}
	categoryService := CategoryService{Repository: repo}

	// Create a mock category
	mockCategory := entity.Category{Id: "1", Name: "Abc"}

	// Set up the mock expectation with the correct return types
	repo.On("FindById", "1").Return(&mockCategory)  // Use "1" for the argument, and return a pointer

	// Call the method under test
	res, err := categoryService.Get("1")

	// Assert that there is no error and that the result is not nil
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, &mockCategory, res) // Optionally check if the returned category is the same as the mock
}
