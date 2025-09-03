package services

import (
	"common/domain/customctx"
	"common/domain/logger"
	"common/utils"
	"fomrs/internal/db/mongo/answers"
	"net/http"
)

func (s *AnswerService) Retrieve(cc *customctx.CustomContext, id string) utils.Response[answers.AnswerModel] {

	entry := logger.FromContext(cc.Context())

	entry.Info("Retrieving answer: ", id)

	answer := s.answersRepository.Find(cc.Context(), id)

	if answer.Err != nil {
		entry.Error("Error retrieving answer", answer.Err)
		return utils.Response[answers.AnswerModel]{
			StatusCode: http.StatusNotFound,
			Success:    false,
			Error:      answer.Err,
		}
	}

	return utils.Response[answers.AnswerModel]{
		StatusCode: http.StatusOK,
		Success:    true,
		Data:       answer.Data,
	}
}
