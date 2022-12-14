package validators

import (
	"fmt"
	"todo-list/internal/todo/dto"
)

type todo struct {
}

func NewTodo() *todo {
	return &todo{}
}

func (t *todo) ValidateCreate(in dto.CreateTodoRequest) error {
	if in.ActivityGroupId == 0 {
		return fmt.Errorf("activity_group_id cannot be null")
	}

	if in.Title == "" {
		return fmt.Errorf("title cannot be null")
	}

	if len(in.Title) > 255 {
		return fmt.Errorf("title too long")
	}

	// optional priority
	if in.Priority != "" {
		var allowedPriority = []string{"very-low", "low", "high", "very-high"}

		for _, val := range allowedPriority {
			if val == in.Priority {
				return nil
			}
		}
		return fmt.Errorf("invalid priority value")
	}

	return nil
}

func (t *todo) ValidateUpdate(in dto.UpdateTodoRequest) error {
	if len(in.Title) > 255 {
		return fmt.Errorf("title too long")
	}

	// optional priority
	if in.Priority != "" {
		var allowedPriority = []string{"very-low", "low", "high", "very-high"}

		for _, val := range allowedPriority {
			if val == in.Priority {
				return nil
			}
		}
		return fmt.Errorf("invalid priority value")
	}
	return nil
}
