package entities

type QuestionEntity struct {
	ID          string         `json:"id" bson:"_id,omitempty"`
	Title       string         `json:"title" bson:"title"`
	Description string         `json:"description" bson:"description"`
	Type        string         `json:"type" bson:"type"`
	Required    bool           `json:"required" bson:"required"`
	Section     string         `json:"section" bson:"section"`
	Metadata    map[string]any `json:"metadata" bson:"metadata"`
}
