package models

import "todo-list/internal/shared/entities"

type Todo struct {
	Id              int64
	ActivityGroupId int64
	Title           string
	IsActive        entities.Bool
	Priority        string
	CreatedAt       int64
	UpdatedAt       int64
	DeletedAt       entities.Int64

	Email string
}
