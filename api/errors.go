package api

import (
	"errors"
	"fmt"
)

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

var (
	ErrBadBlockIdentifier = errors.New("Invalid Block Identifier")
	ErrMissingTxInBlock   = errors.New("Transaction doesn't belong to block")
)

func WrapError(prefix string, cause error) error { return fmt.Errorf("%s: %w", prefix, cause) }
func ErrCantFetchBlockHeader(cause error) error  { return WrapError("Can't fetch Block Header", cause) }
func ErrBadNetworkIdentifier(cause error) error  { return WrapError("Invalid Network Identifier", cause) }

func ErrRpcError(operation string, cause error) error {
	return WrapError(fmt.Sprintf("Celo rpc error: %s", operation), cause)
}
