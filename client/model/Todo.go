package model

import (
	"html"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
)

type Todo struct {
	Id          string
	Title       string
	Description string
	Completed   bool
}

var uni *ut.UniversalTranslator

func (t *Todo) Prepare() {
	t.Title = html.EscapeString(strings.TrimSpace(t.Title))
	t.Description = html.EscapeString(strings.TrimSpace(t.Description))
}

func ValidateInput(data interface{}) ([]string, bool) {
	var validation []string
	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()
	enTranslations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(data)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, l := range errs.Translate(trans) {
			validation = append(validation, l)
		}
		return validation, true
	}
	return nil, false
}
