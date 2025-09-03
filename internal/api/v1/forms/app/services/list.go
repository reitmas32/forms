package services

import (
	"common/domain/customctx"
	"common/domain/logger"
	"common/utils"
	"fomrs/internal/db/mongo/forms"
	"net/http"
)

func (s *FormsService) List(cc *customctx.CustomContext) utils.Response[forms.FormListModel] {

	entry := logger.FromContext(cc.Context())

	entry.Info("Listing forms")

	formsResult := s.formsRepository.FindAll(cc.Context())

	if formsResult.Err != nil {
		entry.Error("Error listing forms", formsResult.Err)
		return utils.Response[forms.FormListModel]{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Error:      formsResult.Err,
		}
	}

	return utils.Response[forms.FormListModel]{
		StatusCode: http.StatusOK,
		Success:    true,
		Results:    formsResult.Data,
	}
}
