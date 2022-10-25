package controllers

import (
	"todo-list/internal/activity/dto"
	"todo-list/internal/shared/presentation"
	"todo-list/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
)

type IActivityService interface {
	Create(dto.CreateActivityRequest) utils.BaseResponse
	FindAll() utils.BaseResponseArray[dto.FindAllActivitiesResponse]
	Detail(dto.DetailActivityRequest) utils.BaseResponse
	Delete(dto.DeleteActivityRequest) utils.BaseResponse
	Update(dto.UpdateActivityRequest) utils.BaseResponse
}

type activity struct {
	service IActivityService
}

func NewActivity(
	service IActivityService,
) *activity {
	return &activity{
		service: service,
	}
}

// CreateActivity bind the create activity request data
func (a *activity) CreateActivity(c *fiber.Ctx) error {
	in := dto.CreateActivityRequest{}

	// bind request
	if err := presentation.ReadRestInJSON(c, &in); err != nil {
		out := utils.BaseResponse{}
		out.Data = struct{}{}

		out.SetResponse(400, err)
		return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
	}

	// call services
	out := a.service.Create(in)

	return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
}

func (a *activity) FindAllActivity(c *fiber.Ctx) error {
	out := a.service.FindAll()
	return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
}

func (a *activity) DetailActivity(c *fiber.Ctx) error {
	in := dto.DetailActivityRequest{}

	// bind params
	if err := presentation.ReadRestInParams(c, &in); err != nil {
		out := utils.BaseResponse{}
		out.Data = struct{}{}
		out.SetResponse(400, err)
		return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
	}

	out := a.service.Detail(in)

	return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
}

func (a *activity) UpdateActivity(c *fiber.Ctx) error {
	in := dto.UpdateActivityRequest{}

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

	out := a.service.Update(in)

	return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
}

func (a *activity) DeleteActivity(c *fiber.Ctx) error {
	in := dto.DeleteActivityRequest{}

	// bind params
	if err := presentation.ReadRestInParams(c, &in); err != nil {
		out := utils.BaseResponse{}
		out.Data = struct{}{}
		out.SetResponse(400, err)
		return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
	}

	out := a.service.Delete(in)

	return presentation.WriteRestOutJSON(c, &out, &out.CommonResult)
}
