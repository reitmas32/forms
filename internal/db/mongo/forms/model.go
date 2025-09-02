package forms

// Geolocalization es una implementación de Entity.
type FormModel struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}

func (g FormModel) GetID() string {
	return g.ID
}
