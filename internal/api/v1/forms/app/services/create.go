package services

import (
	"common/domain/customctx"
	"common/utils"
	"fomrs/internal/api/v1/forms/domain/commands"
	"fomrs/internal/db/mongo/forms"
	"net/http"
)

func (s *FormsService) CreateForm(cc *customctx.CustomContext, command commands.CreateFormCommand) utils.Response[forms.FormModel] {

	form := forms.FormModel{
		Title:       command.Title,
		Description: command.Description,
	}

	model := s.formsRepository.Save(cc.Context(), form)

	if model.Err != nil {
		return utils.Response[forms.FormModel]{
			Data:       form,
			StatusCode: http.StatusInternalServerError,
			Success:    false,
		}
	}

	form.ID = model.Data

	return utils.Response[forms.FormModel]{
		Data:       form,
		StatusCode: http.StatusCreated,
		Success:    true,
	}
}
