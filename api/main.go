package main

import (
	"github.com/gofiber/fiber/v2"
	"rest-grpc/api/config"
	"rest-grpc/api/controller"
	"rest-grpc/api/exception"
	"rest-grpc/api/repository"
	"rest-grpc/api/service"
)

func main() {
	db := config.NewDBConn()

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(&productRepo)
	productController := controller.NewProductController(&productService)

	app := fiber.New(config.NewFiberConfig())

	productController.Route(app)

	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
