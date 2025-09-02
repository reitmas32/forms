package utils

type QuestionType string

const (
	QuestionTypeText      QuestionType = "text"
	QuestionTypeTextLong  QuestionType = "text-long"
	QuestionTypeTextShort QuestionType = "text-short"
	QuestionTypeTextEmail QuestionType = "text-email"

	QuestionTypeRadio    QuestionType = "radio"
	QuestionTypeFile     QuestionType = "file"
	QuestionTypeBoolean  QuestionType = "boolean"
	QuestionTypeSelect   QuestionType = "select"
	QuestionTypeCheckbox QuestionType = "checkbox"
	QuestionTypeDropdown QuestionType = "dropdown"
	QuestionTypeDate     QuestionType = "date"
)

var QuestionTypes = []QuestionType{
	QuestionTypeText,
	QuestionTypeTextLong,
	QuestionTypeTextShort,
	QuestionTypeTextEmail,
	QuestionTypeRadio,
	QuestionTypeFile,
	QuestionTypeBoolean,
	QuestionTypeSelect,
	QuestionTypeCheckbox,
	QuestionTypeDropdown,
	QuestionTypeDate,
}
