package dtos

import "errors"

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
