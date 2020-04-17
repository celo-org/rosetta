package rpc

import (
	"errors"

	"log"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/coinbase/rosetta-sdk-go/types"
)

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
	log.Print(err.Error())
	return ErrValidation
}

func LogErrUnimplemented(rosettaEndpoint string) *types.Error {
	log.Printf("'%s' endpoint not implemented", rosettaEndpoint)
	return ErrUnimplemented
}

func LogErrInternal(err error) *types.Error {
	log.Print(err.Error())
	return ErrInternal
}

func LogErrCeloClient(rpcEndpoint string, err error) *types.Error {
	cause := client.WrapRpcError(err)
	log.Printf("rpc %s failed with %s", rpcEndpoint, cause.Error())
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
