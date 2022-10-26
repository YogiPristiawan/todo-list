package services

import (
	"time"
	"todo-list/internal/shared/databases"
	"todo-list/internal/todo/dto"
	"todo-list/internal/todo/models"
)

var wrapDBError = databases.WrapDBError

func mapCreateTodoToResponse(res *dto.CreateTodoResponse, todo models.Todo) {
	res.Id = todo.Id
	res.Title = todo.Title
	res.ActivityGroupId = todo.ActivityGroupId
	res.IsActive = todo.IsActive
	res.Priority = todo.Priority
	res.CreatedAt = time.Unix(todo.CreatedAt, 0).Format(time.RFC3339)
	res.UpdatedAt = time.Unix(todo.UpdatedAt, 0).Format(time.RFC3339)
}

func mapFindAllTodoToResponse(res *[]dto.FindAllTodoResponse, todos []models.Todo) {
	for _, val := range todos {
		var r dto.FindAllTodoResponse
		r.Id = val.Id
		r.ActivityGroupId = val.ActivityGroupId
		r.Title = val.Title
		if val.IsActive.Bool {
			r.IsActive = "1"
		} else {
			r.IsActive = "0"
		}
		r.Priority = val.Priority
		r.CreatedAt = time.Unix(val.CreatedAt, 0).Format(time.RFC3339)
		r.UpdatedAt = time.Unix(val.UpdatedAt, 0).Format(time.RFC3339)
		if val.DeletedAt.Valid {
			r.DeletedAt.String = time.Unix(val.DeletedAt.Int64, 0).Format(time.RFC3339)
			r.DeletedAt.Valid = true
		}

		*res = append(*res, r)
	}
}

func mapDetailTodoToRespose(res *dto.DetailTodoResponse, todo models.Todo) {
	res.Id = todo.Id
	res.Email = todo.Email
	res.Title = todo.Title
	res.CreatedAt = time.Unix(todo.CreatedAt, 0).Format(time.RFC3339)
	res.UpdatedAt = time.Unix(todo.UpdatedAt, 0).Format(time.RFC3339)
	if todo.DeletedAt.Valid {
		res.DeletedAt.Valid = true
		res.DeletedAt.String = time.Unix(todo.DeletedAt.Int64, 0).Format(time.RFC3339)
	}
}

func mapUpdateTodoToResponse(res *dto.UpdateTodoResponse, todo models.Todo) {
	res.Id = todo.Id
	res.Email = todo.Email
	res.Title = todo.Title
	res.CreatedAt = time.Unix(todo.CreatedAt, 0).Format(time.RFC3339)
	res.UpdatedAt = time.Unix(todo.UpdatedAt, 0).Format(time.RFC3339)
	if todo.DeletedAt.Valid {
		res.DeletedAt.Valid = true
		res.DeletedAt.String = time.Unix(todo.DeletedAt.Int64, 0).Format(time.RFC3339)
	}
}
