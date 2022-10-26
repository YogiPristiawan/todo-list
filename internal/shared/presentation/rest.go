package presentation

import (
	"todo-list/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
)

func ReadRestInQuery[T interface{}](c *fiber.Ctx, in *T) error {
	return c.QueryParser(in)
}

func ReadRestInParams[T interface{}](c *fiber.Ctx, in *T) error {
	return c.ParamsParser(in)
}

func ReadRestInJSON[T interface{}](c *fiber.Ctx, in *T) error {
	return c.BodyParser(in)
}

func WriteRestOutJSON[T interface{}](c *fiber.Ctx, out *T, cr *utils.CommonResult) error {
	if cr.GetCode() == 0 {
		cr.SetResponse(200, nil)
		return c.Status(cr.GetCode()).JSON(out)
	}
	if cr.GetCode() >= 200 && cr.GetCode() < 500 {
		return c.Status(cr.GetCode()).JSON(out)
	}

	return c.Status(cr.GetCode()).JSON(struct {
		Message string `json:"message"`
	}{
		Message: "internal server error",
	})
}
