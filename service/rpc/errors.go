package rpc

import (
	"errors"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/ethereum/go-ethereum/log"
)

var logger = log.Root().New("srv", "rpc")

func newErrorResponse(code int32, msg string, retriable bool) *types.Error {
	return &types.Error{
		Code:      code,
		Message:   msg,
		Retriable: retriable,
	}
}

// [Caution]: all expected error responses must be enumerated in /network/options
func NewErrorResponse(code int32, msg string) *types.Error {
	return newErrorResponse(code, msg, false)
}

// [Caution]: all expected error responses must be enumerated in /network/options
func NewRetriableErrorResponse(code int32, msg string) *types.Error {
	return newErrorResponse(code, msg, true)
}

var (
	ErrValidation    = NewErrorResponse(400, "Request body invalid")
	ErrUnimplemented = NewErrorResponse(405, "Unimplemented rosetta endpoint")
	ErrInternal      = NewErrorResponse(500, "Internal server error")
	ErrCeloClient    = NewErrorResponse(502, "Celo node rpc request failed")
)

func LogErrValidation(err error) *types.Error {
	logger.Error("ValidatorError", "err", err)
	return ErrValidation
}

func LogErrUnimplemented(rosettaEndpoint string) *types.Error {
	logger.Error("NotImplementedError", "endpoint", rosettaEndpoint)
	return ErrUnimplemented
}

func LogErrInternal(err error, params ...interface{}) *types.Error {
	params = append([]interface{}{"err", err}, params...)
	logger.Error("InternalError", params...)
	return ErrInternal
}

func LogErrCeloClient(rpcEndpoint string, err error) *types.Error {
	cause := client.WrapRpcError(err)
	logger.Error("CeloClientError", "endpoint", rpcEndpoint, "err", cause)
	return ErrCeloClient
}

func LogErrFetchBlockHeader(err error) *types.Error {
	return LogErrCeloClient("HeaderAndTxnHashesByNumber", err)
}

var (
	ErrBadBlockIdentifier = errors.New("Bad block identifier")
	ErrFetchBlockHeader   = errors.New("Failed to fetch block header")
	ErrMissingTxInBlock   = errors.New("Transaction doesn't belong to block")
)
