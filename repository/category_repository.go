package repository

import "github.com/danoeandaru/sample-go/v2/entity"

type CategoryRepository interface {
	FindById(Id string) *entity.Category
}
