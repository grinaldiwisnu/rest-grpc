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

func (service *productServiceImpl) Create(req model.CreateProductRequest) (res model.CreateProductResponse, err error) {
	product := entities.Product{
		Id:       req.Id,
		Name:     req.Name,
		Price:    req.Price,
		Quantity: req.Quantity,
	}

	result := service.ProductRepository.Insert(product)
	if result == false {
		return res, errors.New("failed to create product")
	}

	res = model.CreateProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}

	return res, nil
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

func (service *productServiceImpl) Get(req string) (res model.GetProductResponse, err error) {
	product := service.ProductRepository.Get(req)

	if product.Id == "" {
		return res, errors.New("product not found")
	}

	res = model.GetProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}

	return res, nil
}

func (service *productServiceImpl) Delete(req string) (res string, err error) {
	p := entities.Product{Id: req}
	product := service.ProductRepository.Delete(p)

	return product, nil
}
