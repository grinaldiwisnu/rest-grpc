package repository

import (
	"github.com/stretchr/testify/mock"
	"rest-grpc/api/entities"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (p2 *ProductRepositoryMock) Insert(p entities.Product) (status bool) {
	args := p2.Mock.Called(p)
	if args.Get(0) == nil {
		return false
	}

	_ = args.Get(0).(entities.Product)
	return true
}

func (p2 *ProductRepositoryMock) All() (p []entities.Product) {
	panic("implement me")
}

func (p2 *ProductRepositoryMock) Delete(p entities.Product) (pid string) {
	panic("implement me")
}

func (p2 *ProductRepositoryMock) DeleteAll() (status bool) {
	panic("implement me")
}

func (p2 *ProductRepositoryMock) Get(pid string) (p entities.Product) {
	panic("implement me")
}
