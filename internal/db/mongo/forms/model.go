package forms

import "fomrs/internal/api/v1/forms/domain/entities"

// Geolocalization es una implementación de Entity.
type FormModel struct {
	ID          string                    `json:"id" bson:"_id,omitempty"`
	Title       string                    `json:"title" bson:"title"`
	Description string                    `json:"description" bson:"description"`
	Questions   []entities.QuestionEntity `json:"questions" bson:"questions"`
}

func (g FormModel) GetID() string {
	return g.ID
}
