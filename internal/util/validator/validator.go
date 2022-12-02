package validator

import (
	"errors"

	"github.com/go-playground/locales/ja_JP"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
)

var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
)

func init() {
	ja := ja_JP.New()
	uni = ut.New(ja, ja)

	trans, _ = uni.GetTranslator("ja")

	validate = validator.New()
	ja_translations.RegisterDefaultTranslations(validate, trans)
}

func Validate(s interface{}) []error {
	err := validate.Struct(s)
	if err != nil {
		return translate(err)
	}
	return nil
}

func translate(err error) []error {
	var errs []error
	for _, m := range err.(validator.ValidationErrors).Translate(trans) {
		errs = append(errs, errors.New(m))
	}
	return errs
}
