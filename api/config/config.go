package config

import (
	"github.com/gofiber/fiber/v2"
	"rest-grpc/api/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
