package services

import (
	"fomrs/internal/db/mongo/answers"
	"fomrs/internal/db/mongo/forms"
)

type FormsService struct {
	formsRepository   *forms.FormsMongoRepository
	answersRepository *answers.AnswersMongoRepository
}

func NewFormsService(formsRepository *forms.FormsMongoRepository, answersRepository *answers.AnswersMongoRepository) *FormsService {
	return &FormsService{formsRepository: formsRepository, answersRepository: answersRepository}
}
