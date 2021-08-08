package repository

import "rest-grpc/api/entities"

type ProductRepository interface {
	Insert(p entities.Product) (status bool)
	All() (p []entities.Product)
	Delete(p entities.Product) (pid string)
	DeleteAll() (status bool)
	Get(pid string) (p entities.Product)
}
