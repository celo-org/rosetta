package rpc

import (
	"errors"
	"fmt"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/coinbase/rosetta-sdk-go/types"
)

func newErrorResponse(code int32, err error, retriable bool) *types.Error {
	return &types.Error{
		Code:      code,
		Message:   err.Error(),
		Retriable: retriable,
	}
}

func NewErrorResponse(code int32, err error) *types.Error {
	return newErrorResponse(code, err, false)
}

func NewRetriableErrorResponse(code int32, err error) *types.Error {
	return newErrorResponse(code, err, true)
}

func NewValidationError(err error) *types.Error {
	return NewErrorResponse(400, err)
}

func NewInternalError(err error) *types.Error {
	return NewErrorResponse(500, err)
}

func NewCeloClientError(rpcEndpoint string, err error) *types.Error {
	cause := client.WrapRpcError(err)
	formattedErr := fmt.Errorf("rpc %s failed with %s", rpcEndpoint, cause.Error())
	return NewErrorResponse(502, formattedErr)
}

func ErrUnimplementedEndpoint(rosettaEndpoint string) *types.Error {
	return NewErrorResponse(404, fmt.Errorf("'%s' endpoint not implemented", rosettaEndpoint))
}

var (
	ErrBadBlockIdentifier = errors.New("Invalid Block Identifier")
	ErrMissingTxInBlock   = errors.New("Transaction doesn't belong to block")
)

func ErrCantFetchBlockHeader(cause error) *types.Error {
	return NewCeloClientError("Can't fetch Block Header", cause)
}
