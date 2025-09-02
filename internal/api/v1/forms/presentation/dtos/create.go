package dtos

import (
	"errors"
	"fomrs/internal/api/v1/forms/domain/commands"
	"fomrs/internal/utils"
	"slices"

	"common/utils/ctypes"
)

type QuestionDTO struct {
	Title       string         `json:"title" binding:"required"`
	Description string         `json:"description" binding:"required"`
	Type        string         `json:"type" binding:"required"`
	Required    bool           `json:"required" binding:"required"`
	Section     string         `json:"section"`
	Metadata    map[string]any `json:"metadata"`
}

func (question QuestionDTO) Validate() error {

	if !slices.Contains(utils.QuestionTypes, utils.QuestionType(question.Type)) {
		return errors.New("invalid question type: " + question.Type)
	}

	return nil
}

func (question QuestionDTO) ToCommand() commands.QuestionCommand {

	return commands.QuestionCommand{
		Title:       question.Title,
		Description: question.Description,
		Type:        question.Type,
		Required:    question.Required,
		Section:     question.Section,
		Metadata:    question.Metadata,
	}
}

type CreateFormDTO struct {
	Title       string        `json:"title" binding:"required"`
	Description string        `json:"description" binding:"required"`
	Questions   []QuestionDTO `json:"questions" binding:"required"`
}

func (dto CreateFormDTO) Validate() error {

	for _, question := range dto.Questions {
		if err := question.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (dto CreateFormDTO) ToCommand() commands.CreateFormCommand {
	return commands.CreateFormCommand{
		Title:       dto.Title,
		Description: dto.Description,
		Questions: ctypes.Map(
			dto.Questions,
			func(question QuestionDTO) commands.QuestionCommand {
				return question.ToCommand()
			},
		),
	}
}
