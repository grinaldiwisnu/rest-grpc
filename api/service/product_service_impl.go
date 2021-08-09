package service

import (
	"errors"
	"rest-grpc/api/entities"
	"rest-grpc/api/model"
	"rest-grpc/api/repository"
)

func NewProductService(repository *repository.ProductRepository) ProductService {
	return &productServiceImpl{ProductRepository: *repository}
}

type productServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func (service *productServiceImpl) Create(req model.CreateProductRequest) (*model.CreateProductResponse, error) {
	product := entities.Product{
		Id:       req.Id,
		Name:     req.Name,
		Price:    req.Price,
		Quantity: req.Quantity,
	}

	result := service.ProductRepository.Insert(product)
	if result == false {
		return nil, errors.New("failed to create product")
	}

	res := model.CreateProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}

	return &res, nil
}

func (service *productServiceImpl) List() (res []model.GetProductResponse, err error) {
	products := service.ProductRepository.All()

	if len(products) == 0 {
		return res, errors.New("products empty")
	}

	for _, product := range products {
		res = append(res, model.GetProductResponse{
			Id:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}

	return res, nil
}

func (service *productServiceImpl) Get(req string) (*model.GetProductResponse, error) {
	product := service.ProductRepository.Get(req)

	if product.Id == "" {
		return nil, errors.New("product not found")
	}

	res := model.GetProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}

	return &res, nil
}

func (service *productServiceImpl) Delete(req string) (string, error) {
	p := entities.Product{Id: req}
	product := service.ProductRepository.Delete(p)

	return product, nil
}
