package client

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/rpc"
)

var (
	ErrMissingTrieNode     = errors.New("Block too Old: Missing Trie Node")
	ErrContractNotDeployed = errors.New("Contract Not Deployed")
)

func RpcErrorCode(err error) (int, bool) {
	rpcError, ok := err.(rpc.Error)
	if ok {
		return rpcError.ErrorCode(), true
	}
	return 0, false
}

func WrapRpcError(err error) error {
	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), "No Implementation set") {
		return ErrContractNotDeployed
	}

	code, isRpc := RpcErrorCode(err)
	if !isRpc {
		return err
	}

	if code == -32000 && strings.HasPrefix(err.Error(), "missing trie node") {
		return ErrMissingTrieNode
	}

	return err
}
