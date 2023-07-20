package category_service

import (
	"src/go-mocks-learn/entity"
	"src/go-mocks-learn/repository/category_repository"
)

type categoryServiceStruct struct {
	Repository category_repository.ICategory
}

var CategoryService = categoryServiceStruct{
	Repository: category_repository.Category,
}

func (input categoryServiceStruct) GetData(data entity.CategoryEntity) (result entity.CategoryEntity, err error) {
	result, err = input.Repository.FindByID(nil, data)
	return
}
