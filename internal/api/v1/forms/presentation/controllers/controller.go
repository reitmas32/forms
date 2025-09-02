package controllers

import "fomrs/internal/api/v1/forms/app/services"

type FormsController struct {
	formsService *services.FormsService
}

func NewFormsController(formsService *services.FormsService) *FormsController {
	return &FormsController{formsService: formsService}
}
