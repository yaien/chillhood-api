package validation

import (
	"sync"

	"github.com/go-playground/locales/en"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate   *validator.Validate
	translator ut.Translator
}

func (v *Validator) Validate(s interface{}) (map[string]string, error) {
	err := v.validate.Struct(s)
	if err == nil {
		return make(map[string]string), nil
	}
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return nil, err
	}
	errors := err.(validator.ValidationErrors)
	return errors.Translate(v.translator), nil
}

var v *Validator
var once sync.Once

func GetValidator() *Validator {
	once.Do(func() {
		validate := validator.New()
		translator := en.New()
		universalTranslator := ut.New(translator, translator)
		v = &Validator{validate, universalTranslator.GetFallback()}
	})
	return v
}
