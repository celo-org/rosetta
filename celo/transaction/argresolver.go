package transaction

import (
	"context"
	"fmt"
	"math/big"

	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ArgResolver func(context.Context, []interface{}) ([]interface{}, error)

type ResolverMap struct {
	resolvers map[CeloMethod]ArgResolver
}

func NewResolverMap(registry *wrapper.RegistryWrapper) *ResolverMap {
	resolvers := make(map[CeloMethod]ArgResolver)

	resolvers[Vote] = VoteResolver(registry)
	resolvers[RevokePendingVotes] = RevokeVoteResolver(registry)
	resolvers[RevokeActiveVotes] = RevokeVoteResolver(registry)
	resolvers[AuthorizeVoteSigner] = AuthorizeSignerResolver(registry)

	return &ResolverMap{
		resolvers: resolvers,
	}
}

func (rm ResolverMap) FindResolverFor(method CeloMethod) ArgResolver {
	resolver, ok := rm.resolvers[method]
	if !ok {
		return NoopResolver
	}
	return resolver
}

////////////////////////////////////////////////////////////////////////////////////////
// Resolvers
////////////////////////////////////////////////////////////////////////////////////////

func NoopResolver(ctx context.Context, args []interface{}) ([]interface{}, error) {
	return args, nil
}

func VoteResolver(registry *wrapper.RegistryWrapper) ArgResolver {
	return func(ctx context.Context, args []interface{}) ([]interface{}, error) {
		err := validateArgLength(args, 2)
		if err != nil {
			return nil, err
		}

		group, err := parseAddress(args, 0)
		if err != nil {
			return nil, err
		}

		value, err := parseBigInt(args, 1)
		if err != nil {
			return nil, err
		}

		election, err := registry.GetElectionWrapper(ctx, nil)
		if err != nil {
			return nil, err
		}

		keys, err := election.VoteMetadata(&bind.CallOpts{
			Context: ctx,
		}, group, value)
		if err != nil {
			return nil, err
		}

		return []interface{}{group, value, keys.Lesser, keys.Greater}, nil
	}
}

func RevokeVoteResolver(registry *wrapper.RegistryWrapper) ArgResolver {
	return func(ctx context.Context, args []interface{}) ([]interface{}, error) {
		err := validateArgLength(args, 3)
		if err != nil {
			return nil, err
		}

		account, err := parseAddress(args, 0)
		if err != nil {
			return nil, err
		}

		group, err := parseAddress(args, 1)
		if err != nil {
			return nil, err
		}

		value, err := parseBigInt(args, 2)
		if err != nil {
			return nil, err
		}

		election, err := registry.GetElectionWrapper(ctx, nil)
		if err != nil {
			return nil, err
		}
		keys, err := election.RevokeMetadata(&bind.CallOpts{
			Context: ctx,
		}, account, group, value)
		if err != nil {
			return nil, err
		}

		return []interface{}{group, value, keys.Lesser, keys.Greater}, nil
	}
}

func AuthorizeSignerResolver(registry *wrapper.RegistryWrapper) ArgResolver {
	return func(ctx context.Context, args []interface{}) ([]interface{}, error) {
		err := validateArgLength(args, 2)
		if err != nil {
			return nil, err
		}

		signer, err := parseAddress(args, 0)
		if err != nil {
			return nil, err
		}

		pop, err := parseBytes(args, 1)
		if err != nil {
			return nil, err
		}

		accounts, err := registry.GetAccountsWrapper(ctx, nil)
		if err != nil {
			return nil, err
		}
		encodedSig, err := accounts.AuthorizeMetadata(pop)
		if err != nil {
			return nil, err
		}

		return []interface{}{signer, encodedSig.V, encodedSig.R, encodedSig.S}, nil
	}
}

////////////////////////////////////////////////////////////////////////////////////////
// Argument Parsers
////////////////////////////////////////////////////////////////////////////////////////

func argTypeErr(arg interface{}, index int, expected string) error {
	return fmt.Errorf("Argument %v provided at index %d needs type %s", arg, index, expected)
}

func parseAddress(args []interface{}, index int) (common.Address, error) {
	addr, ok := args[index].(common.Address)
	if !ok {
		return addr, argTypeErr(args, index, "common.Address")
	}
	return addr, nil
}

func parseBigInt(args []interface{}, index int) (*big.Int, error) {
	bigInt, ok := args[index].(*big.Int)
	if !ok {
		return nil, argTypeErr(args, index, "*big.Int")
	}
	return bigInt, nil
}

func parseBytes(args []interface{}, index int) ([]byte, error) {
	bytes, ok := args[index].([]byte)
	if !ok {
		return nil, argTypeErr(args, index, "[]byte")
	}
	return bytes, nil
}

func validateArgLength(args []interface{}, length int) error {
	if len(args) != length {
		return fmt.Errorf("Received %d args; expected %d", len(args), length)
	}
	return nil
}
