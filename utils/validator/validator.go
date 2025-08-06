package validator

import (
	"log/slog"
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
)

func init() {
	// Configura os idiomas suportados
	enLocale := en.New()
	ptLocale := pt.New()
	uni = ut.New(enLocale, ptLocale)

	// Inicializa o validador
	validate = validator.New()
}

func ValidateStruct(s any, lang string) []restError.Cause {
	// Seleciona o tradutor com base no idioma
	translator, found := uni.GetTranslator(lang)
	if !found {
		translator, _ = uni.GetTranslator("en") // Fallback para inglês
	}

	// Registra as traduções no validador
	switch lang {
	case "pt":
		pt_translations.RegisterDefaultTranslations(validate, translator)
	default:
		en_translations.RegisterDefaultTranslations(validate, translator)
	}

	if err := validate.Struct(s); err != nil {
		var causes []restError.Cause
		for _, err := range err.(validator.ValidationErrors) {
			causes = append(causes, restError.Cause{
				Field:   err.Field(),
				Message: err.Translate(translator),
			})
		}
		slog.Error("Validation errors", "causes", causes)
		return causes
	}
	return nil
}
