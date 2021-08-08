package service

import "rest-grpc/api/model"

type ProductService interface {
	Create(req model.CreateProductRequest) (res model.CreateProductResponse, err error)
	List() (res []model.GetProductResponse, err error)
	Delete(req string) (res string, err error)
	Get(req string) (res model.GetProductResponse, err error)
}
