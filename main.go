package main

import (
	"fmt"
	"log"
	"src/go-mocks-learn/entity"
	"src/go-mocks-learn/service/category_service"
)

func main() {
	result, err := category_service.CategoryService.GetData(entity.CategoryEntity{
		ID:       1,
		Name:     "Real",
		Gender:   "L",
		IsActive: true,
	})

	if err != nil {
		log.Fatalf(`error is exist %s`, err.Error())
	}

	fmt.Println(fmt.Sprintf(`result [ID] -> %d`, result.ID))
	fmt.Println(fmt.Sprintf(`result [Name] -> %s`, result.Name))
	fmt.Println(fmt.Sprintf(`result [Gender] -> %s`, result.Gender))
}
