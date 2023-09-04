package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_BR_translations "github.com/go-playground/validator/v10/translations/pt_BR"
	"github.com/walber-vaz/crud-go/configs/rest_error"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		pt_BR := pt_BR.New()
		unt := ut.New(pt_BR, pt_BR)
		transl, _ = unt.GetTranslator("pt_BR")
		pt_BR_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_error.RestError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_error.NewBadRequestError("Tipo de dado inválido")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsTypes := []rest_error.Error_type{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			listError := rest_error.Error_type{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsTypes = append(errorsTypes, listError)
		}

		return rest_error.NewBadRequestValidationError("Erro de validação", errorsTypes)
	} else {
		return rest_error.NewBadRequestError("Erro de conversão de tipos de erro")
	}
}
