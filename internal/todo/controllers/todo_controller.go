package controllers

import (
	"todo-list/internal/shared/presentation"
	"todo-list/internal/shared/utils"
	"todo-list/internal/todo/dto"

	"github.com/gofiber/fiber/v2"
)

type ITodoService interface {
	Create(dto.CreateTodoRequest) utils.BaseResponse
	FindAll(dto.FindAllTodoRequest) utils.BaseResponseArray[dto.FindAllTodoResponse]
	Detail(dto.DetailTodoRequest) utils.BaseResponse
	Update(dto.UpdateTodoRequest) utils.BaseResponse
	Delete(dto.DeleteTodoRequest) utils.BaseResponse
}

type todo struct {
	service ITodoService
}

func NewTodo(service ITodoService) *todo {
	return &todo{
		service: service,
	}
}

func (t *todo) CreateTodo(c *fiber.Ctx) error {
	in := dto.CreateTodoRequest{}

	// bind body
	if err := presentation.ReadRestInJSON(c, &in); err != nil {
		out := utils.BaseResponse{}
		out.Data = struct{}{}
		out.SetResponse(400, err)
		return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
	}

	out := t.service.Create(in)

	return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
}

func (t *todo) FindAllTodo(c *fiber.Ctx) error {
	in := dto.FindAllTodoRequest{}

	// bind query
	if err := presentation.ReadRestInQuery(c, &in); err != nil {
		out := utils.BaseResponse{}
		out.Data = struct{}{}
		out.SetResponse(400, err)
		return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
	}

	out := t.service.FindAll(in)

	return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
}

func (t *todo) DetailTodo(c *fiber.Ctx) error {
	in := dto.DetailTodoRequest{}

	// bind params
	if err := presentation.ReadRestInParams(c, &in); err != nil {
		out := utils.BaseResponse{}
		out.Data = struct{}{}
		out.SetResponse(400, err)
		return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
	}

	out := t.service.Detail(in)

	return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
}

func (t *todo) UpdateTodo(c *fiber.Ctx) error {
	in := dto.UpdateTodoRequest{}

	// bind params
	if err := presentation.ReadRestInParams(c, &in); err != nil {
		out := utils.BaseResponse{}
		out.Data = struct{}{}
		out.SetResponse(400, err)
		return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
	}

	// bind body
	if err := presentation.ReadRestInJSON(c, &in); err != nil {
		out := utils.BaseResponse{}
		out.Data = struct{}{}
		out.SetResponse(400, err)
		return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
	}

	out := t.service.Update(in)

	return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
}

func (t *todo) DeleteTodo(c *fiber.Ctx) error {
	in := dto.DeleteTodoRequest{}

	// bind params
	if err := presentation.ReadRestInParams(c, &in); err != nil {
		out := utils.BaseResponse{}
		out.Data = struct{}{}
		out.SetResponse(400, err)
		return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
	}

	out := t.service.Delete(in)

	return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
}
