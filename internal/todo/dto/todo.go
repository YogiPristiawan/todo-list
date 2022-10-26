package dto

import "todo-list/internal/shared/entities"

type CreateTodoRequest struct {
	ActivityGroupId int64         `json:"activity_group_id"`
	Title           string        `json:"title"`
	Priority        string        `json:"priority"`
	IsActive        entities.Bool `json:"is_active"`
}

type CreateTodoResponse struct {
	Id              int64         `json:"id"`
	Title           string        `json:"title"`
	ActivityGroupId int64         `json:"activity_group_id"`
	IsActive        entities.Bool `json:"is_active"`
	Priority        string        `json:"priority"`
	CreatedAt       string        `json:"created_at"`
	UpdatedAt       string        `json:"updated_at"`
}

type FindAllTodoRequest struct {
	ActivityGroupId int64 `query:"activity_group_id"`
}

type FindAllTodoResponse struct {
	Id              int64           `json:"id"`
	ActivityGroupId int64           `json:"activity_group_id"`
	Title           string          `json:"title"`
	IsActive        string          `json:"is_active"`
	Priority        string          `json:"priority"`
	CreatedAt       string          `json:"created_at"`
	UpdatedAt       string          `json:"updated_at"`
	DeletedAt       entities.String `json:"deleted_at"`
}

type DetailTodoRequest struct {
	Id int64 `params:"id"`
}

type DetailTodoResponse struct {
	Id              int64           `json:"id"`
	ActivityGroupId int64           `json:"activity_group_id"`
	Title           string          `json:"title"`
	IsActive        string          `json:"is_active"`
	Priority        string          `json:"priority"`
	CreatedAt       string          `json:"created_at"`
	UpdatedAt       string          `json:"updated_at"`
	DeletedAt       entities.String `json:"deleted_at"`
}

type UpdateTodoRequest struct {
	Id              int64         `params:"id"`
	ActivityGroupId int64         `json:"activity_group_id"`
	Title           string        `json:"title"`
	Priority        string        `json:"priority"`
	IsActive        entities.Bool `json:"is_active"`
}

type UpdateTodoResponse struct {
	Id              int64           `json:"id"`
	ActivityGroupId int64           `json:"activity_group_id"`
	Title           string          `json:"title"`
	IsActive        string          `json:"is_active"`
	Priority        string          `json:"priority"`
	CreatedAt       string          `json:"created_at"`
	UpdatedAt       string          `json:"udpated_at"`
	DeletedAt       entities.String `json:"deleted_at"`
}

type DeleteTodoRequest struct {
	Id int64 `params:"id"`
}
