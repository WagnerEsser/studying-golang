package utils

import (
	"fmt"
	"log/slog"
	restError "studying-go/types"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s any) []restError.Cause {
	validate := validator.New()
	if err := validate.Struct(s); err != nil {
		var causes []restError.Cause
		for _, err := range err.(validator.ValidationErrors) {
			causes = append(causes, restError.Cause{
				Field:   err.Field(),
				Message: fmt.Sprintf("Validation error: %s", err.Tag()),
			})
		}
		slog.Error("Validation errors", "causes", causes)
		return causes
	}
	return nil
}
