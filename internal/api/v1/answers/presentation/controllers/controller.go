package controllers

import "fomrs/internal/api/v1/answers/app/services"

type AnswerController struct {
	service *services.AnswerService
}

func NewAnswerController(service *services.AnswerService) *AnswerController {
	return &AnswerController{service: service}
}
