package services

import "fomrs/internal/db/mongo/forms"

type FormsService struct {
	formsRepository *forms.FormsMongoRepository
}

func NewFormsService(formsRepository *forms.FormsMongoRepository) *FormsService {
	return &FormsService{formsRepository: formsRepository}
}
