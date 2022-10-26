package services

import (
	"fmt"
	"todo-list/internal/shared/utils"
	"todo-list/internal/todo/dto"
	"todo-list/internal/todo/models"
)

type ITodoValidator interface {
	ValidateCreate(dto.CreateTodoRequest) error
	ValidateUpdate(dto.UpdateTodoRequest) error
}

type ITodoRepo interface {
	Create(*models.Todo) error
	FindAll(activityGroupId int64) ([]models.Todo, error)
	Detail(id int64) (models.Todo, error)
	Update(id int64, todo *models.Todo) (affected int64, err error)
	Delete(id int64) (affected int64, err error)
}

type todo struct {
	validator ITodoValidator
	todoRepo  ITodoRepo
}

func NewTodo(
	validator ITodoValidator,
	todoRepo ITodoRepo,
) *todo {
	return &todo{
		validator: validator,
		todoRepo:  todoRepo,
	}
}

func (t *todo) Create(in dto.CreateTodoRequest) (out utils.BaseResponse) {
	// validate payload
	if err := t.validator.ValidateCreate(in); err != nil {
		out.Data = struct{}{}
		out.SetResponse(400, err)
		return
	}

	// create todo
	todo := models.Todo{
		Title:           in.Title,
		ActivityGroupId: in.ActivityGroupId,
		Priority:        in.Priority,
		IsActive:        in.IsActive,
	}
	err := t.todoRepo.Create(&todo)
	switch wrapDBError(err) {
	case 404:
		out.Data = struct{}{}
		out.SetResponse(404, fmt.Errorf("activity not found"))
		return
	case 500:
		out.SetResponse(500, err)
		return
	}

	// map to response
	res := dto.CreateTodoResponse{}
	mapCreateTodoToResponse(&res, todo)

	out.Message = "Success"
	out.Data = res
	out.SetResponse(201, nil)
	return
}

func (t *todo) FindAll(in dto.FindAllTodoRequest) (out utils.BaseResponseArray[dto.FindAllTodoResponse]) {
	// find all todos
	todos, err := t.todoRepo.FindAll(in.ActivityGroupId)
	switch wrapDBError(err) {
	case 500:
		out.SetResponse(500, err)
		return
	}

	if len(todos) == 0 {
		out.Message = "Success"
		out.Data = []dto.FindAllTodoResponse{}
		out.SetResponse(200, nil)
		return
	}

	// map to response
	mapFindAllTodoToResponse(&out.Data, todos)

	out.Message = "Success"
	return
}

func (t *todo) Detail(in dto.DetailTodoRequest) (out utils.BaseResponse) {
	// get detail todo
	todo, err := t.todoRepo.Detail(in.Id)
	switch wrapDBError(err) {
	case 404:
		out.Data = struct{}{}
		out.SetResponse(404, fmt.Errorf("Todo with ID %d Not Found", in.Id))
		return
	case 500:
		out.SetResponse(500, err)
		return
	}

	// map to response
	res := dto.DetailTodoResponse{}
	mapDetailTodoToRespose(&res, todo)

	out.Message = "Success"
	out.Data = res
	out.SetResponse(200, nil)
	return
}

func (t *todo) Update(in dto.UpdateTodoRequest) (out utils.BaseResponse) {
	// validate request
	if err := t.validator.ValidateUpdate(in); err != nil {
		out.SetResponse(400, err)
		return
	}

	// update todo
	todo := models.Todo{
		ActivityGroupId: in.ActivityGroupId,
		Title:           in.Title,
		Priority:        in.Priority,
		IsActive:        in.IsActive,
	}
	affected, err := t.todoRepo.Update(in.Id, &todo)
	switch wrapDBError(err) {
	case 404:
		out.Data = struct{}{}
		out.SetResponse(404, fmt.Errorf("Todo with ID %d Not Found", in.Id))
		return
	case 500:
		out.SetResponse(500, err)
		return
	}

	if affected == 0 {
		out.SetResponse(404, fmt.Errorf("Todo with ID %d Not Found", in.Id))
		return
	}

	// map response
	res := dto.UpdateTodoResponse{}
	mapUpdateTodoToResponse(&res, todo)

	out.Message = "Success"
	out.Data = res
	out.SetResponse(200, nil)
	return
}

func (t *todo) Delete(in dto.DeleteTodoRequest) (out utils.BaseResponse) {
	// delete todo
	affected, err := t.todoRepo.Delete(in.Id)
	switch wrapDBError(err) {
	case 500:
		out.SetResponse(500, err)
		return
	}

	if affected == 0 {
		out.Data = struct{}{}
		out.SetResponse(404, fmt.Errorf("Todo with ID %d Not Found", in.Id))
		return
	}

	out.Message = "Success"
	out.Data = struct{}{}
	out.SetResponse(200, nil)
	return
}
