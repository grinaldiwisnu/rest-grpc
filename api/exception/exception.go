package exception

import (
	"github.com/gofiber/fiber/v2"
	"rest-grpc/api/model"
)

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return ctx.JSON(model.Response{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	return ctx.JSON(model.Response{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}

type ValidationError struct {
	Message string
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}
