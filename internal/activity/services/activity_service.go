package services

import (
	"fmt"
	"todo-list/internal/activity/dto"
	"todo-list/internal/activity/models"
	"todo-list/internal/shared/utils"
)

type IActivityValidator interface {
	ValidateCreate(dto.CreateActivityRequest) error
	ValidateUpdate(dto.UpdateActivityRequest) error
}

type IActivityRepo interface {
	Create(*models.Activity) error
	FindAll() ([]models.Activity, error)
	GetById(id int64) (models.Activity, error)
	Delete(id int64) (affected int64, err error)
	Update(id int64, activity *models.Activity) (affected int64, err error)
}

type activity struct {
	validator    IActivityValidator
	activityRepo IActivityRepo
}

func NewActivity(
	validator IActivityValidator,
	activityRepo IActivityRepo,
) *activity {
	return &activity{
		validator:    validator,
		activityRepo: activityRepo,
	}
}

func (a *activity) Create(in dto.CreateActivityRequest) (out utils.BaseResponse) {
	// validate request
	if err := a.validator.ValidateCreate(in); err != nil {
		out.Data = struct{}{}
		out.SetResponse(400, err)
		return
	}

	// create activity
	activity := models.Activity{
		Title: in.Title,
		Email: in.Email,
	}
	err := a.activityRepo.Create(&activity)
	switch wrapDBError(err) {
	case 500:
		out.Data = struct{}{}
		out.SetResponse(500, err)
		return
	}

	// map to response
	res := dto.CreateActivityResponse{}
	mapCreateActivityToResponse(&res, activity)

	out.Message = "Success"
	out.Data = res
	out.SetResponse(201, nil)
	return
}

func (a *activity) FindAll() (out utils.BaseResponseArray[dto.FindAllActivitiesResponse]) {
	// find activities
	activities, err := a.activityRepo.FindAll()
	switch wrapDBError(err) {
	case 404:
		out.Data = append(out.Data, dto.FindAllActivitiesResponse{})
		out.SetResponse(200, nil)
		return
	case 500:
		out.SetResponse(500, err)
		return
	}

	if len(activities) == 0 {
		out.Message = "Success"
		out.Data = []dto.FindAllActivitiesResponse{}
		out.SetResponse(200, nil)
		return
	}

	// map to response
	mapFindAllActivityToResponse(&out.Data, activities)

	out.Message = "Success"
	out.SetResponse(200, nil)
	return
}

func (a *activity) Detail(in dto.DetailActivityRequest) (out utils.BaseResponse) {
	// get activity
	activiy, err := a.activityRepo.GetById(in.Id)
	switch wrapDBError(err) {
	case 404:
		out.Data = struct{}{}
		out.SetResponse(404, fmt.Errorf("Activity with ID %d Not Found", in.Id))
		return
	case 500:
		out.SetResponse(500, err)
		return
	}

	// map to response
	res := dto.DetailActivityResponse{}
	mapDetailActivityToResponse(&res, activiy)

	out.Message = "Success"
	out.Data = res
	return
}

func (a *activity) Delete(in dto.DeleteActivityRequest) (out utils.BaseResponse) {
	// delete activity
	affected, err := a.activityRepo.Delete(in.Id)
	switch wrapDBError(err) {
	case 500:
		out.SetResponse(500, err)
		return
	}

	if affected == 0 {
		out.Data = struct{}{}
		out.SetResponse(404, fmt.Errorf("Activity with ID %d Not Found", in.Id))
		return
	}

	out.Message = "Success"
	out.Data = struct{}{}
	return
}

func (a *activity) Update(in dto.UpdateActivityRequest) (out utils.BaseResponse) {
	// validate request
	if err := a.validator.ValidateUpdate(in); err != nil {
		out.Data = struct{}{}
		out.SetResponse(400, err)
		return
	}

	// perform update
	activity := models.Activity{}
	if in.Title != "" {
		activity.Title = in.Title
	}
	if in.Email != "" {
		activity.Email = in.Email
	}
	affected, err := a.activityRepo.Update(in.Id, &activity)
	switch wrapDBError(err) {
	case 500:
		out.Data = struct{}{}
		out.SetResponse(500, err)
		return
	}

	if affected == 0 {
		out.Data = struct{}{}
		out.SetResponse(404, fmt.Errorf("Activity with ID %d Not Found", in.Id))
		return
	}

	// map to response
	res := dto.UpdateActivityResponse{}
	mapUpdateActivityToResponse(&res, activity)

	out.Message = "Success"
	out.Data = res
	return
}
