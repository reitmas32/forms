package services

import (
	"common/domain/customctx"
	"common/domain/logger"
	"common/utils"
	"common/utils/cerrs"
	"fmt"
	"fomrs/internal/api/v1/answers/domain/commands"
	"fomrs/internal/db/mongo/answers"
	"net/http"

	utils_internal "fomrs/internal/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *AnswerService) Create(cc *customctx.CustomContext, command *commands.ResponseCommand) utils.Response[answers.AnswerModel] {

	entry := logger.FromContext(cc.Context())

	// Get Form
	form := s.formsRepository.Find(cc.Context(), command.FormID)

	if form.Err != nil {
		entry.Error("Error getting form", form.Err)
		return utils.Response[answers.AnswerModel]{
			StatusCode: http.StatusNotFound,
			Success:    false,
			Error:      form.Err,
		}
	}

	// Validate

	questions := form.Data.Questions

	for _, question := range questions {
		isRequired := question.Required
		questionType := question.Type
		questionID := question.ID

		var responseAnswer string

		for _, response := range command.Responses {
			if questionID == response.QuestionID {
				responseAnswer = response.Answer
			}
		}

		if isRequired && responseAnswer == "" {
			entry.Error("Question is required", cerrs.NewCustomError(http.StatusBadRequest, "Question is required", "create.answer"))
			return utils.Response[answers.AnswerModel]{
				StatusCode: http.StatusBadRequest,
				Success:    false,
				Error: cc.NewError(
					cerrs.NewCustomError(
						http.StatusBadRequest,
						"Question is required: "+question.Title,
						"forms.create.answer.required",
					),
				),
			}
		}

		if responseAnswer == "" {
			continue
		}

		validator := utils_internal.Validators[utils_internal.QuestionType(questionType)]

		if questionType == string(utils_internal.QuestionTypeRadio) {

			metadata := question.Metadata

			options := metadata["options"].(primitive.A)

			optionsString := make([]string, len(options))
			for i, option := range options {
				optionsString[i] = option.(string)
			}

			validator = utils_internal.RadioValidator{
				Options: optionsString,
			}

			fmt.Println("options", options)
		}

		if questionType == string(utils_internal.QuestionTypeSelect) {

			metadata := question.Metadata

			options := metadata["options"].(primitive.A)

			optionsString := make([]string, len(options))
			for i, option := range options {
				optionsString[i] = option.(string)
			}

			validator = utils_internal.SelectValidator{
				Options: optionsString,
			}

			fmt.Println("options", options)
		}

		if questionType == string(utils_internal.QuestionTypeCheckbox) {

			metadata := question.Metadata

			options := metadata["options"].(primitive.A)

			optionsString := make([]string, len(options))
			for i, option := range options {
				optionsString[i] = option.(string)
			}

			validator = utils_internal.CheckboxValidator{
				Options: optionsString,
			}

			fmt.Println("options", options)
		}

		isValid := validator.IsValid(responseAnswer)

		if !isValid {
			entry.Error("Invalid answer", cerrs.NewCustomError(http.StatusBadRequest, "Invalid answer", "create.answer"))
			return utils.Response[answers.AnswerModel]{
				StatusCode: http.StatusBadRequest,
				Success:    false,
				Error: cc.NewError(
					cerrs.NewCustomError(
						http.StatusBadRequest,
						"Invalid answer: ["+question.Title+"] "+validator.Description(),
						"forms.create.answer.invalid",
					),
				),
			}
		}

	}

	// Insert Response
	answer := answers.AnswerModel{
		FormID:  command.FormID,
		Answers: command.Responses,
	}

	res := s.answersRepository.Save(cc.Context(), answer)

	if res.Err != nil {
		entry.Error("Error saving answer", res.Err)
		return utils.Response[answers.AnswerModel]{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Error:      res.Err,
		}
	}

	answer.ID = res.Data

	return utils.Response[answers.AnswerModel]{
		Data:       answer,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}
