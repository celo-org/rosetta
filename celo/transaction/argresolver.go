package transaction

import (
	"context"
	"fmt"
	"math/big"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ArgResolver func([]interface{}, context.Context) ([]interface{}, error)

type ResolverLocator struct {
	registry   *wrapper.RegistryWrapper
	election   *wrapper.ElectionWrapper
	lockedGold *wrapper.LockedGoldWrapper
}

func NewResolverLocator(registry *wrapper.RegistryWrapper, celoClient *client.CeloClient) (*ResolverLocator, error) {
	election, err := wrapper.NewElection(celoClient, registry)
	if err != nil {
		return nil, err
	}

	lockedGold, err := wrapper.NewLockedGold(celoClient, registry)
	if err != nil {
		return nil, err
	}

	return &ResolverLocator{
		registry:   registry,
		election:   election,
		lockedGold: lockedGold,
	}, nil
}

func (rl *ResolverLocator) GetResolver(method *CeloMethod) ArgResolver {
	switch method {
	case &Vote:
		return rl.voteResolver
	case &RevokePendingVotes:
		fallthrough
	case &RevokeActiveVotes:
		return rl.revokeVoteResolver
	default:
		return rl.defaultResolver
	}
}

func (*ResolverLocator) defaultResolver(args []interface{}, _ context.Context) ([]interface{}, error) {
	return args, nil
}

type argType string

var (
	addressType argType = "common.Address"
	bigIntType  argType = "*big.Int"
)

func (at argType) String() string { return string(at) }

func argTypeErr(arg interface{}, index int, expected argType) error {
	return fmt.Errorf("Argument %v provided at index %d needs type %s", arg, index, expected.String())
}

func assertAddressType(args []interface{}, index int) (common.Address, error) {
	addr, ok := args[index].(common.Address)
	if !ok {
		return addr, argTypeErr(args, index, addressType)
	}
	return addr, nil
}

func assertBigIntType(args []interface{}, index int) (*big.Int, error) {
	bigInt, ok := args[index].(*big.Int)
	if !ok {
		return nil, argTypeErr(args, index, bigIntType)
	}
	return bigInt, nil
}

func assertLength(args []interface{}, length int) error {
	if len(args) != length {
		return fmt.Errorf("Received %d args; expected %d", len(args), length)
	}
	return nil
}

func (rl *ResolverLocator) voteResolver(args []interface{}, ctx context.Context) ([]interface{}, error) {
	err := assertLength(args, 2)
	if err != nil {
		return nil, err
	}

	group, err := assertAddressType(args, 0)
	if err != nil {
		return nil, err
	}

	value, err := assertBigIntType(args, 1)
	if err != nil {
		return nil, err
	}

	keys, err := rl.election.VoteMetadata(&bind.CallOpts{
		Context: ctx,
	}, group, value)
	if err != nil {
		return nil, err
	}

	args = append(args, keys.Lesser, keys.Greater)
	return args, nil
}

func (rl *ResolverLocator) revokeVoteResolver(args []interface{}, ctx context.Context) ([]interface{}, error) {
	err := assertLength(args, 3)
	if err != nil {
		return nil, err
	}

	account, err := assertAddressType(args, 0)
	if err != nil {
		return nil, err
	}

	group, err := assertAddressType(args, 1)
	if err != nil {
		return nil, err
	}

	value, err := assertBigIntType(args, 2)
	if err != nil {
		return nil, err
	}

	keys, err := rl.election.RevokeMetadata(&bind.CallOpts{
		Context: ctx,
	}, account, group, value)
	if err != nil {
		return nil, err
	}

	args = append(args, keys.Lesser, keys.Greater)
	return args, nil
}
