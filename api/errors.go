package api

import "fmt"

type ValidationError struct {
	field string
	err   error
}

func NewValidationError(field string, err error) *ValidationError {
	return &ValidationError{
		field: field,
		err:   err,
	}
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("validation error: field: %s error: %s", ve.field, ve.err)
}
