package validator

import (
	"fmt"
	"log/slog"
	"strings"
	restError "studying-go/types/resterror"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	pt_translations "github.com/go-playground/validator/v10/translations/pt"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	fieldMap = map[string]map[string]string{
		"en": {
			"PhoneNumber": "Phone number",
			"Email":       "Email",
			"Age":         "Age",
			"Address":     "Address",
			"Password":    "Password",
			"Street":      "Street",
			"Number":      "Number",
			"City":        "City",
			"State":       "State",
			"Country":     "Country",
		},
		"pt": {
			"PhoneNumber": "Número de telefone",
			"Email":       "E-mail",
			"Age":         "Idade",
			"Address":     "Endereço",
			"Password":    "Senha",
			"Street":      "Rua",
			"Number":      "Número",
			"City":        "Cidade",
			"State":       "Estado",
			"Country":     "País",
		},
	}
)

func init() {
	enLocale := en.New()
	ptLocale := pt.New()
	uni = ut.New(enLocale, ptLocale)

	validate = validator.New()
}

func ValidateStruct(s any, lang string) []restError.Cause {
	translator, found := uni.GetTranslator(lang)
	if !found {
		translator, _ = uni.GetTranslator("en")
	}

	switch lang {
	case "pt":
		pt_translations.RegisterDefaultTranslations(validate, translator)
	default:
		en_translations.RegisterDefaultTranslations(validate, translator)
	}

	registerFieldTranslations(translator, lang)

	if err := validate.Struct(s); err != nil {
		var causes []restError.Cause
		for _, err := range err.(validator.ValidationErrors) {
			translatedFieldName := translateFieldName(err.Field(), lang)

			message := err.Translate(translator)
			message = replaceFieldNameInMessage(message, err.Field(), translatedFieldName)

			causes = append(causes, restError.Cause{
				Field:   err.Field(),
				Message: message,
			})
		}
		slog.Error("Validation errors", "causes", causes)
		return causes
	}
	return nil
}

func registerFieldTranslations(translator ut.Translator, lang string) {
	for field, translation := range fieldMap[lang] {
		_ = validate.RegisterTranslation(field, translator, func(ut ut.Translator) error {
			return ut.Add(field, fmt.Sprintf("%s is invalid", translation), true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(field, translation)
			return t
		})
	}
}

func translateFieldName(fieldName, lang string) string {
	if translation, ok := fieldMap[lang][fieldName]; ok {
		return translation
	}
	return fieldName
}

func replaceFieldNameInMessage(message, originalField, translatedField string) string {
	return strings.ReplaceAll(message, originalField, translatedField)
}
