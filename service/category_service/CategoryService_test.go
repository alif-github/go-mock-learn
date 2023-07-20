package category_service

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"src/go-mocks-learn/entity"
	"src/go-mocks-learn/repository/category_repository"
	"testing"
)

func TestCategoryService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks := category_repository.NewMockICategory(ctrl)
	service := &categoryServiceStruct{Repository: mocks}

	cases := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "Case 1 Name Different",
			test: func(t *testing.T) {
				// Input DAO
				categoryModel := entity.CategoryEntity{
					ID:       99,
					Name:     "Mock",
					Gender:   "P",
					IsActive: true,
				}

				// Output DAO
				categoryModelResult := entity.CategoryEntity{
					ID:       99,
					Name:     "Mock Changed",
					Gender:   "P",
					IsActive: true,
				}

				mocks.EXPECT().FindByID(nil, categoryModel).Return(categoryModelResult, nil).Times(1)
				result, err := service.GetData(categoryModel)
				assert.Nil(t, err, "Error must nil")
				assert.NotEqual(t, categoryModel.Name, result.Name, "Result Name must not equal")
			},
		},
		{
			name: "Case 2 Gender Different",
			test: func(t *testing.T) {
				// Input DAO
				categoryModel := entity.CategoryEntity{
					ID:       99,
					Name:     "Mock",
					Gender:   "P",
					IsActive: true,
				}

				// Output DAO
				categoryModelResult := entity.CategoryEntity{
					ID:       99,
					Name:     "Mock",
					Gender:   "L",
					IsActive: true,
				}

				mocks.EXPECT().FindByID(nil, categoryModel).Return(categoryModelResult, nil).Times(1)
				result, err := service.GetData(categoryModel)
				assert.Nil(t, err, "Error must nil")
				assert.NotEqual(t, categoryModel.Gender, result.Gender, "Result Gender must not equal")
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, tc.test)
	}
}
