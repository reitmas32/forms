package entities

type AnswerEntity struct {
	QuestionID string   `json:"question_id" binding:"required"`
	Answer     string   `json:"answer"`
	Values     []string `json:"values"`
}

type ResponseEntity struct {
	ID        string         `json:"id"`
	FormID    string         `json:"form_id"`
	UserID    string         `json:"user_id"`
	Responses []AnswerEntity `json:"responses"`
}
