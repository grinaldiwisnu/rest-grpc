package repository

import (
	"github.com/go-pg/pg/v10"
	"rest-grpc/api/entities"
)

func NewProductRepository(db *pg.DB) ProductRepository {
	return &productRepositoryImpl{
		Database: db,
	}
}

type productRepositoryImpl struct {
	Database *pg.DB
}

func (repo *productRepositoryImpl) Insert(product entities.Product) (status bool) {
	_, err := repo.Database.Model(&product).Returning("id").Insert()

	if err != nil {
		return false
	}

	return true
}

func (repo *productRepositoryImpl) All() (products []entities.Product) {
	err := repo.Database.Model(&products).Select()

	if err != nil {
		return []entities.Product{}
	}

	return products
}

func (repo *productRepositoryImpl) DeleteAll() (status bool) {
	_, err := repo.Database.Model(&entities.Product{}).Delete()
	if err != nil {
		return false
	}
	return true
}

func (repo *productRepositoryImpl) Delete(product entities.Product) (pid string) {

	_, err := repo.Database.Model(&product).Where("id = ?", product.Id).Delete()
	if err != nil {
		return ""
	}

	return product.Id
}

func (repo *productRepositoryImpl) Get(pid string) (product entities.Product) {
	err := repo.Database.Model(&product).Where("id = ?", pid).Select()
	if err != nil {
		return entities.Product{}
	}

	return product
}
