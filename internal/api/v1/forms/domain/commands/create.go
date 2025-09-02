package commands

import "fomrs/internal/api/v1/forms/domain/entities"

type QuestionCommand struct {
	Title       string         `json:"title" binding:"required"`
	Description string         `json:"description" binding:"required"`
	Type        string         `json:"type" binding:"required"`
	Required    bool           `json:"required" binding:"required"`
	Section     string         `json:"section"`
	Metadata    map[string]any `json:"metadata"`
}

func (c QuestionCommand) ToEntity() entities.QuestionEntity {
	return entities.QuestionEntity{
		Title:       c.Title,
		Description: c.Description,
		Type:        c.Type,
		Required:    c.Required,
		Section:     c.Section,
		Metadata:    c.Metadata,
	}
}

type CreateFormCommand struct {
	Title       string            `json:"title" binding:"required"`
	Description string            `json:"description" binding:"required"`
	Questions   []QuestionCommand `json:"questions" binding:"required"`
}

func (c CreateFormCommand) Validate() error {
	return nil
}
