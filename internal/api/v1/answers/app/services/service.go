package services

import (
	"fomrs/internal/db/mongo/answers"
	"fomrs/internal/db/mongo/forms"
)

type AnswerService struct {
	formsRepository   *forms.FormsMongoRepository
	answersRepository *answers.AnswersMongoRepository
}

func NewAnswerService(formsRepository *forms.FormsMongoRepository, answersRepository *answers.AnswersMongoRepository) *AnswerService {
	return &AnswerService{formsRepository: formsRepository, answersRepository: answersRepository}
}
