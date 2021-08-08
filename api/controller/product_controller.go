package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"rest-grpc/api/model"
	"rest-grpc/api/service"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService *service.ProductService) ProductController {
	return ProductController{
		ProductService: *productService,
	}
}

func (p *ProductController) Route(app *fiber.App) {
	app.Post("/api/product", p.Create)
	app.Get("/api/product", p.Get)
	app.Get("/api/product/:id", p.GetByID)
}

func (p *ProductController) Create(ctx *fiber.Ctx) error {
	var req model.CreateProductRequest
	err := ctx.BodyParser(&req)
	req.Id = uuid.New().String()

	if err != nil {
		return ctx.JSON(model.Response{
			Code:   401,
			Status: "Error parsing request",
			Data:   nil,
		})
	}

	res, errs := p.ProductService.Create(req)
	if errs != nil {
		return ctx.JSON(model.Response{
			Code:   500,
			Status: errs.Error(),
			Data:   nil,
		})
	}
	return ctx.JSON(model.Response{
		Code:   201,
		Status: "OK",
		Data:   res,
	})
}

func (p *ProductController) Get(ctx *fiber.Ctx) error {
	res, err := p.ProductService.List()
	if err != nil {
		return ctx.JSON(model.Response{
			Code:   500,
			Status: err.Error(),
			Data:   nil,
		})
	}
	return ctx.JSON(model.Response{
		Code:   200,
		Status: "OK",
		Data:   res,
	})
}

func (p *ProductController) GetByID(ctx *fiber.Ctx) error {
	res, err := p.ProductService.Get(ctx.Params("id"))
	if err != nil {
		return ctx.JSON(model.Response{
			Code:   500,
			Status: err.Error(),
			Data:   nil,
		})
	}
	return ctx.JSON(model.Response{
		Code:   200,
		Status: "OK",
		Data:   res,
	})
}
