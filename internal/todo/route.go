package todo

import "github.com/gofiber/fiber/v2"

type ITodo interface {
	ITodoController
}

type ITodoController interface {
	CreateTodo(c *fiber.Ctx) error
	FindAllTodo(c *fiber.Ctx) error
	DetailTodo(c *fiber.Ctx) error
	UpdateTodo(c *fiber.Ctx) error
	DeleteTodo(c *fiber.Ctx) error
}

func NewRoute(app *fiber.App, controller ITodo) {
	app.Post("/todo-items", controller.CreateTodo)
	app.Get("/todo-items", controller.FindAllTodo)
	app.Get("/todo-items/:id", controller.DetailTodo)
	app.Patch("/todo-items/:id", controller.UpdateTodo)
	app.Delete("/todo-items/:id", controller.DeleteTodo)
}
