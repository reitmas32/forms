package dtos

import (
	"common/utils/ctypes"
	"errors"
	"fomrs/internal/api/v1/answers/domain/commands"
	"fomrs/internal/api/v1/answers/domain/entities"
)

type AnswerDTO struct {
	QuestionID string   `json:"question_id" binding:"required"`
	Answer     string   `json:"answer"`
	Values     []string `json:"values"`
}

func (a AnswerDTO) Validate() error {
	if a.Answer == "" && len(a.Values) == 0 {
		return errors.New("answer or values are required")
	}
	return nil
}

func (a AnswerDTO) ToEntity() entities.AnswerEntity {

	return entities.AnswerEntity{
		QuestionID: a.QuestionID,
		Answer:     a.Answer,
		Values:     a.Values,
	}
}

type CreateAnswerDTO struct {
	FormID    string      `json:"form_id" binding:"required"`
	UserID    string      `json:"user_id"`
	Responses []AnswerDTO `json:"responses" binding:"required"`
}

func (c CreateAnswerDTO) Validate() error {
	for _, answer := range c.Responses {
		if err := answer.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (dto CreateAnswerDTO) ToCommand() commands.ResponseCommand {

	return commands.ResponseCommand{
		FormID: dto.FormID,
		UserID: dto.UserID,
		Responses: ctypes.Map(
			dto.Responses,
			func(answer AnswerDTO) entities.AnswerEntity {
				return answer.ToEntity()
			},
		),
	}
}
