package answers

import "fomrs/internal/api/v1/answers/domain/entities"

// Geolocalization es una implementación de Entity.
type AnswerModel struct {
	ID      string                  `json:"id" bson:"_id,omitempty"`
	FormID  string                  `json:"form_id" bson:"form_id"`
	Answers []entities.AnswerEntity `json:"answers" bson:"answers"`
}

func (g AnswerModel) GetID() string {
	return g.ID
}
