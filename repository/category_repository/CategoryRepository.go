package category_repository

import (
	"database/sql"
	"src/go-mocks-learn/entity"
)

type ICategory interface {
	FindByID(*sql.DB, entity.CategoryEntity) (entity.CategoryEntity, error)
}

var Category = categoryRepository{}.New()

type categoryRepository struct {
	TableName string
}

func (i categoryRepository) New() *categoryRepository {
	return &categoryRepository{TableName: "Category_DB"}
}

func (i categoryRepository) FindByID(_ *sql.DB, entity entity.CategoryEntity) (entity.CategoryEntity, error) {
	//-- Simulate with real db connection
	return entity, nil
}
