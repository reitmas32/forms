package services

import (
	"common/domain/customctx"
	"common/domain/logger"
	"common/utils"
	"fomrs/internal/db/mongo/forms"
	"net/http"
)

func (s *FormsService) Retrieve(cc *customctx.CustomContext, id string) utils.Response[forms.FormModel] {

	entry := logger.FromContext(cc.Context())

	entry.Info("Retrieving form id: ", id)

	form := s.formsRepository.Find(cc.Context(), id)

	if form.Err != nil {
		entry.Error("Error retrieving form", form.Err)
		return utils.Response[forms.FormModel]{
			StatusCode: http.StatusNotFound,
			Success:    false,
			Error:      form.Err,
		}
	}

	return utils.Response[forms.FormModel]{
		StatusCode: http.StatusOK,
		Success:    true,
		Data:       form.Data,
	}
}
