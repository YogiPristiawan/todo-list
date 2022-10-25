package services

import (
	"time"
	"todo-list/internal/activity/dto"
	"todo-list/internal/activity/models"
	"todo-list/internal/shared/databases"
	"todo-list/internal/shared/entities"
)

// This variables store functions
// required by service
// It store into variable for easier testing purpose

var wrapDBError = databases.WrapDBError

// mapCreateActivityToResponse map the activity model
// into create activity response
func mapCreateActivityToResponse(res *dto.CreateActivityResponse, activity models.Activity) {
	res.Id = activity.Id
	res.Title = activity.Title
	res.Email = activity.Email
	res.CreatedAt = time.Unix(activity.CreatedAt, 0).Format(time.RFC3339)
	res.UpdatedAt = time.Unix(activity.UpdatedAt, 0).Format(time.RFC3339)
}

// mapFindAllToResponse map the activity model
// into find all activities response
func mapFindAllActivityToResponse(res *[]dto.FindAllActivitiesResponse, activities []models.Activity) {
	for _, val := range activities {
		var r dto.FindAllActivitiesResponse

		r.Id = val.Id
		r.Email = val.Email
		r.Title = val.Title
		r.CreatedAt = time.Unix(val.CreatedAt, 0).Format(time.RFC3339)
		r.UpdatedAt = time.Unix(val.UpdatedAt, 0).Format(time.RFC3339)
		if val.DeletedAt.Valid {
			r.DeletedAt = time.Unix(val.DeletedAt.Int64, 0).Format(time.RFC3339)
		}

		*res = append(*res, r)
	}
}

func mapDetailActivityToResponse(res *dto.DetailActivityResponse, activity models.Activity) {
	res.Id = activity.Id
	res.Email = activity.Email
	res.Title = activity.Title
	res.CreatedAt = time.Unix(activity.CreatedAt, 0).Format(time.RFC3339)
	res.UpdatedAt = time.Unix(activity.UpdatedAt, 0).Format(time.RFC3339)
	if activity.DeletedAt.Valid {
		res.DeletedAt = entities.String{
			String: time.Unix(activity.DeletedAt.Int64, 0).Format(time.RFC3339),
			Valid:  true,
		}
	}
}

func mapUpdateActivityToResponse(res *dto.UpdateActivityResponse, activity models.Activity) {
	res.Id = activity.Id
	res.Title = activity.Title
	res.Email = activity.Email
	res.CreatedAt = time.Unix(activity.CreatedAt, 0).Format(time.RFC3339)
	res.UpdatedAt = time.Unix(activity.UpdatedAt, 0).Format(time.RFC3339)
	if activity.DeletedAt.Valid {
		res.DeletedAt = entities.String{
			String: time.Unix(activity.DeletedAt.Int64, 0).Format(time.RFC3339),
			Valid:  true,
		}
	}
}
