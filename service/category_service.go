package service

import (
	"errors"

	"github.com/danoeandaru/sample-go/v2/entity"
	"github.com/danoeandaru/sample-go/v2/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (categoryService CategoryService) Get(Id string) (*entity.Category, error) {
	category := categoryService.Repository.FindById(Id)
	if category == nil {
		return nil, errors.New("Category not found.")
	} else {
		return category, nil
	}
}
