package validators

import (
	"fmt"
	"regexp"
	"todo-list/internal/activity/dto"
)

type activity struct {
}

func NewActivity() *activity {
	return &activity{}
}

func (a *activity) ValidateCreate(in dto.CreateActivityRequest) error {
	if in.Email == "" {
		return fmt.Errorf("email required")
	}

	r, _ := regexp.Compile(`.+@.+\..+`)
	if !r.Match([]byte(in.Email)) {
		return fmt.Errorf("email tidak valid")
	}

	if len(in.Email) > 255 {
		return fmt.Errorf("email too long")
	}

	if in.Title == "" {
		return fmt.Errorf("title required")
	}

	if len(in.Title) > 255 {
		return fmt.Errorf("title too long")
	}

	return nil
}

func (a *activity) ValidateUpdate(in dto.UpdateActivityRequest) error {
	if in.Title == "" {
		return fmt.Errorf("title cannot be null")
	}

	if in.Email != "" {
		r, _ := regexp.Compile(`.+@.+\..+`)
		if !r.Match([]byte(in.Email)) {
			return fmt.Errorf("email tidak valid")
		}
	}

	if len(in.Email) > 255 {
		return fmt.Errorf("email too long")
	}

	if len(in.Title) > 255 {
		return fmt.Errorf("title too long")
	}

	return nil
}
