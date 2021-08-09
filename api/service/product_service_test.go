package service

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"rest-grpc/api/entities"
	"rest-grpc/api/model"
	"rest-grpc/api/repository"
	"testing"
)

var repo = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var service = productServiceImpl{ProductRepository: repo}

func TestCreateProductSuccess(t *testing.T) {
	param := entities.Product{
		Id:       uuid.New().String(),
		Name:     "Product Testing",
		Price:    10000,
		Quantity: 100,
	}

	repo.Mock.On("Insert", param).Return(param)

	prd, err := service.Create(model.CreateProductRequest{
		Id:       param.Id,
		Name:     param.Name,
		Price:    param.Price,
		Quantity: param.Quantity,
	})

	assert.Nil(t, err)
	assert.NotNil(t, prd)
}

func TestCreateProductFailed(t *testing.T) {
	param := entities.Product{
		Id:       uuid.New().String(),
		Name:     "Product Testing",
		Price:    10000,
		Quantity: 100,
	}

	repo.Mock.On("Insert", param).Return(nil)

	prd, err := service.Create(model.CreateProductRequest{
		Id:       param.Id,
		Name:     param.Name,
		Price:    param.Price,
		Quantity: param.Quantity,
	})

	assert.Nil(t, prd)
	assert.NotNil(t, err)
}

func TestGetAllProductEmpty(t *testing.T) {
	repo.Mock.On("All").Return([]entities.Product{})
	prds, err := service.List()

	assert.Nil(t, err)
	assert.NotNil(t, prds)
	assert.Len(t, prds, 0)
}

func TestGetAllProductSuccess(t *testing.T) {
	list := []entities.Product{
		{
			Id:       uuid.New().String(),
			Name:     "Product Testing 1",
			Price:    2000,
			Quantity: 20,
		},
	}
	repo.Mock.On("All").Return(list)
	prds, err := service.List()

	assert.Nil(t, err)
	assert.NotNil(t, prds)
	assert.Len(t, prds, 1)
	assert.ElementsMatch(t, prds, list)
}
