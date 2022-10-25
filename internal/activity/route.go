package activity

import "github.com/gofiber/fiber/v2"

type IActivity interface {
	IActivityController
}

type IActivityController interface {
	CreateActivity(c *fiber.Ctx) error
	FindAllActivity(c *fiber.Ctx) error
	DetailActivity(c *fiber.Ctx) error
	UpdateActivity(c *fiber.Ctx) error
	DeleteActivity(c *fiber.Ctx) error
}

func NewRoute(app *fiber.App, controller IActivity) {
	app.Post("/activity-groups", controller.CreateActivity)
	app.Get("/activity-groups", controller.FindAllActivity)
	app.Get("/activity-groups/:id", controller.DetailActivity)
	app.Delete("/activity-groups/:id", controller.DeleteActivity)
	app.Patch("/activity-groups/:id", controller.UpdateActivity)
}
