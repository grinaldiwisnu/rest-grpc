package service

import "rest-grpc/api/model"

type ProductService interface {
	Create(req model.CreateProductRequest) (*model.CreateProductResponse, error)
	List() ([]model.GetProductResponse, error)
	Delete(req string) (string, error)
	Get(req string) (*model.GetProductResponse, error)
}
