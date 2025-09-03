package commands

import "fomrs/internal/api/v1/answers/domain/entities"

type ResponseCommand struct {
	FormID    string                  `json:"form_id" binding:"required"`
	UserID    string                  `json:"user_id"`
	Responses []entities.AnswerEntity `json:"responses" binding:"required"`
}
