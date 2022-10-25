package dto

import "todo-list/internal/shared/entities"

type CreateActivityRequest struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type CreateActivityResponse struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type FindAllActivitiesResponse struct {
	Id        int64  `json:"id"`
	Email     string `json:"email"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type DetailActivityRequest struct {
	Id int64 `params:"id"`
}

type DetailActivityResponse struct {
	Id        int64           `json:"id"`
	Email     string          `json:"email"`
	Title     string          `json:"title"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
	DeletedAt entities.String `json:"deleted_at"`
}

type DeleteActivityRequest struct {
	Id int64 `params:"id"`
}

type UpdateActivityRequest struct {
	Id    int64  `params:"id"`
	Email string `json:"email"`
	Title string `json:"title"`
}

type UpdateActivityResponse struct {
	Id        int64           `json:"id"`
	Title     string          `json:"title"`
	Email     string          `json:"email"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
	DeletedAt entities.String `json:"deleted_at"`
}
