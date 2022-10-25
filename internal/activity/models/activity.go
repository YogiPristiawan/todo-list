package models

import "todo-list/internal/shared/entities"

type Activity struct {
	Id        int64
	Title     string
	Email     string
	CreatedAt int64
	UpdatedAt int64
	DeletedAt entities.Int64
}
