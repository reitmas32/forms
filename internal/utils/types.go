package utils

import (
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

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

// Interfaz genérica
type Validator interface {
	Name() string
	IsValid(value string) bool
}

// ========== Implementaciones ==========

// Texto genérico (no vacío)
type TextValidator struct{}

func (t TextValidator) Name() string { return string(QuestionTypeText) }
func (t TextValidator) IsValid(value string) bool {
	return strings.TrimSpace(value) != ""
}

// Texto largo (> 20 chars por ejemplo)
type TextLongValidator struct{}

func (t TextLongValidator) Name() string { return string(QuestionTypeTextLong) }
func (t TextLongValidator) IsValid(value string) bool {
	return len(strings.TrimSpace(value)) > 20
}

// Texto corto (< 50 chars por ejemplo)
type TextShortValidator struct{}

func (t TextShortValidator) Name() string { return string(QuestionTypeTextShort) }
func (t TextShortValidator) IsValid(value string) bool {
	return len(strings.TrimSpace(value)) > 0 && len(value) <= 50
}

// Email
type EmailValidator struct{}

func (e EmailValidator) Name() string { return string(QuestionTypeTextEmail) }
func (e EmailValidator) IsValid(value string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(value)
}

// Radio (valor debe estar en un set predefinido, aquí simulamos)
type RadioValidator struct{}

func (r RadioValidator) Name() string { return string(QuestionTypeRadio) }
func (r RadioValidator) IsValid(value string) bool {
	options := []string{"op1", "op2", "op3"}
	for _, o := range options {
		if value == o {
			return true
		}
	}
	return false
}

// File (validamos extensión simple)
type FileValidator struct{}

func (f FileValidator) Name() string { return string(QuestionTypeFile) }
func (f FileValidator) IsValid(value string) bool {
	ext := strings.ToLower(filepath.Ext(value))
	allowed := []string{".jpg", ".png", ".pdf", ".txt"}
	for _, a := range allowed {
		if ext == a {
			return true
		}
	}
	return false
}

// Boolean
type BooleanValidator struct{}

func (b BooleanValidator) Name() string { return string(QuestionTypeBoolean) }
func (b BooleanValidator) IsValid(value string) bool {
	return value == "true" || value == "false"
}

// Select (igual que radio pero puede ser otra lista)
type SelectValidator struct{}

func (s SelectValidator) Name() string { return string(QuestionTypeSelect) }
func (s SelectValidator) IsValid(value string) bool {
	options := []string{"a", "b", "c"}
	for _, o := range options {
		if value == o {
			return true
		}
	}
	return false
}

// Checkbox (múltiples valores separados por coma, todos deben ser válidos)
type CheckboxValidator struct{}

func (c CheckboxValidator) Name() string { return string(QuestionTypeCheckbox) }
func (c CheckboxValidator) IsValid(value string) bool {
	options := []string{"x", "y", "z"}
	values := strings.Split(value, ",")
	if len(values) == 0 {
		return false
	}
	for _, v := range values {
		v = strings.TrimSpace(v)
		found := false
		for _, o := range options {
			if v == o {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Dropdown (igual que select pero distinta lista de opciones)
type DropdownValidator struct{}

func (d DropdownValidator) Name() string { return string(QuestionTypeDropdown) }
func (d DropdownValidator) IsValid(value string) bool {
	options := []string{"uno", "dos", "tres"}
	for _, o := range options {
		if value == o {
			return true
		}
	}
	return false
}

// Fecha (YYYY-MM-DD)
type DateValidator struct{}

func (d DateValidator) Name() string { return string(QuestionTypeDate) }
func (d DateValidator) IsValid(value string) bool {
	_, err := time.Parse("2006-01-02", value)
	return err == nil
}

// ========== Mapa de validadores ==========
var Validators = map[QuestionType]Validator{
	QuestionTypeText:      TextValidator{},
	QuestionTypeTextLong:  TextLongValidator{},
	QuestionTypeTextShort: TextShortValidator{},
	QuestionTypeTextEmail: EmailValidator{},
	QuestionTypeRadio:     RadioValidator{},
	QuestionTypeFile:      FileValidator{},
	QuestionTypeBoolean:   BooleanValidator{},
	QuestionTypeSelect:    SelectValidator{},
	QuestionTypeCheckbox:  CheckboxValidator{},
	QuestionTypeDropdown:  DropdownValidator{},
	QuestionTypeDate:      DateValidator{},
}
