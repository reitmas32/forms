package dtos

import "fomrs/internal/api/v1/forms/domain/commands"

type CreateFormDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (dto CreateFormDTO) Validate() error {
	return nil
}

func (dto CreateFormDTO) ToCommand() commands.CreateFormCommand {
	return commands.CreateFormCommand{
		Title:       dto.Title,
		Description: dto.Description,
	}
}
