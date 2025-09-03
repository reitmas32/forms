package services

import (
	"common/domain/criteria"
	"common/domain/customctx"
	"common/domain/logger"
	"common/utils"
	"fomrs/internal/db/mongo/answers"
	"net/http"
)

func (s *FormsService) Answers(cc *customctx.CustomContext, id string) utils.Response[answers.AnswerListModel] {

	entry := logger.FromContext(cc.Context())

	entry.Info("Retrieving answers of form: ", id)

	cri := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "form_id",
					Operator: criteria.OperatorEqual,
					Value:    id,
				},
			},
		),
	}

	answersResults := s.answersRepository.Matching(cri, "answers", 0, 10)

	if answersResults.Err != nil {
		entry.Error("Error retrieving answers", answersResults.Err)
		return utils.Response[answers.AnswerListModel]{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Error:      answersResults.Err,
		}
	}

	return utils.Response[answers.AnswerListModel]{
		StatusCode: http.StatusOK,
		Success:    true,
		Results:    answersResults.Data,
	}
}
