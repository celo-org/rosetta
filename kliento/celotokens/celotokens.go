package celotokens

import (
	"context"
	"fmt"
	"math/big"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/kliento/contracts"
	"github.com/celo-org/kliento/registry"
)

// CeloToken is a token native to Celo, eg CELO and stable tokens
type CeloToken string

const (
	// CELO - previously known as cGLD
	CELO CeloToken = "CELO"
	// CUSD - Celo Dollar
	CUSD CeloToken = "cUSD"
	// CEUR - Celo Euro
	CEUR CeloToken = "cEUR"
	// CREAL - Celo Real
	CREAL CeloToken = "cREAL"
)

// CeloTokenInfo provides basic info for a Celo token
type CeloTokenInfo struct {
	registryID         registry.ContractID
	exchangeRegistryID registry.ContractID
	isStableToken      bool
}

// CELOInfo contains info on the CELO token
var CELOInfo = CeloTokenInfo{
	registryID:    registry.GoldTokenContractID,
	isStableToken: false,
}

// CUSDInfo contains info on the cUSD token
var CUSDInfo = CeloTokenInfo{
	registryID:         registry.StableTokenContractID,
	exchangeRegistryID: registry.ExchangeContractID,
	isStableToken:      true,
}

// CEURInfo contains info on the cEUR token
var CEURInfo = CeloTokenInfo{
	registryID:         registry.StableTokenEURContractID,
	exchangeRegistryID: registry.ExchangeEURContractID,
	isStableToken:      true,
}

// CREALInfo contains info on the cREAL token
var CREALInfo = CeloTokenInfo{
	registryID:         registry.StableTokenBRLContractID,
	exchangeRegistryID: registry.ExchangeBRLContractID,
	isStableToken:      true,
}

// CeloTokenInfos contains a CeloTokenInfo entry for each CeloToken key.
// This is used for functions that iterate on all tokens or stable tokens.
var CeloTokenInfos = map[CeloToken]CeloTokenInfo{
	CELO:  CELOInfo,
	CUSD:  CUSDInfo,
	CEUR:  CEURInfo,
	CREAL: CREALInfo,
}

// CeloTokens provides a friendly interface for interacting with Celo tokens
type CeloTokens struct {
	registry registry.Registry
}

// New creates a new pointer to a CeloTokens
func New(reg registry.Registry) *CeloTokens {
	return &CeloTokens{
		registry: reg,
	}
}

// GetExchangeContract gets the exchange contract for a provided stable token.
// Note that if the contract has not been registered with the Registry as of
// blockNumber, an error is returned.
func (ct *CeloTokens) GetExchangeContract(ctx context.Context, token CeloToken, blockNumber *big.Int) (*contracts.Exchange, error) {
	tokenInfo, ok := CeloTokenInfos[token]
	if !ok {
		return nil, fmt.Errorf("Token %s not found", token)
	}
	if !tokenInfo.isStableToken {
		return nil, fmt.Errorf("Token %s not a stable token", token)
	}
	contract, err := ct.registry.GetContractByID(ctx, string(tokenInfo.exchangeRegistryID), blockNumber)
	if err != nil {
		return nil, err
	}
	return contract.(*contracts.Exchange), nil
}

// GetStableTokenContract gets the stabletoken contract for a provided stable token.
// Note that if the contract has not been registered with the Registry as of
// blockNumber, an error is returned.
func (ct *CeloTokens) GetStableTokenContract(ctx context.Context, token CeloToken, blockNumber *big.Int) (*contracts.StableToken, error) {
	tokenInfo, ok := CeloTokenInfos[token]
	if !ok {
		return nil, fmt.Errorf("Token %s not found", token)
	}
	if !tokenInfo.isStableToken {
		return nil, fmt.Errorf("Token %s not a stable token", token)
	}
	contract, err := ct.registry.GetContractByID(ctx, string(tokenInfo.registryID), blockNumber)
	if err != nil {
		return nil, err
	}
	return contract.(*contracts.StableToken), nil
}

// GetContract gets the contract for a provided celo token.
// Note that if the contract has not been registered with the Registry as of
// blockNumber, an error is returned.
func (ct *CeloTokens) GetContract(ctx context.Context, token CeloToken, blockNumber *big.Int) (contracts.CeloTokenContract, error) {
	tokenInfo, ok := CeloTokenInfos[token]
	if !ok {
		return nil, fmt.Errorf("Token %s not found", token)
	}
	contract, err := ct.registry.GetContractByID(ctx, string(tokenInfo.registryID), blockNumber)
	if err != nil {
		return nil, err
	}
	return contract.(contracts.CeloTokenContract), nil
}

// GetExchangeContracts gets the exchange contracts for all stable tokens.
// If an exchange contract has not been registered with the Registry as of blockNumber,
// the contract is set to nil and no error is returned.
func (ct *CeloTokens) GetExchangeContracts(ctx context.Context, blockNumber *big.Int) (map[CeloToken]*contracts.Exchange, error) {
	exchangeContracts := make(map[CeloToken]*contracts.Exchange)
	for token, tokenInfo := range CeloTokenInfos {
		// Only stable tokens have a corresponding exchange contract
		if !tokenInfo.isStableToken {
			continue
		}
		exchangeContract, err := ct.GetExchangeContract(ctx, token, blockNumber)
		if err != nil {
			if !registry.IsExpectedBeforeContractsDeployed(err) {
				return nil, err
			}
			exchangeContract = nil
		}
		exchangeContracts[token] = exchangeContract
	}
	return exchangeContracts, nil
}

// GetContracts gets the contracts for all tokens. If onlyStables is true, contracts
// for only stable tokens are provided.
// If a contract has not been registered with the Registry as of blockNumber,
// the contract is set to nil and no error is returned.
func (ct *CeloTokens) GetContracts(ctx context.Context, blockNumber *big.Int, onlyStables bool) (map[CeloToken]contracts.CeloTokenContract, error) {
	contracts := make(map[CeloToken]contracts.CeloTokenContract)
	for token, tokenInfo := range CeloTokenInfos {
		if onlyStables && !tokenInfo.isStableToken {
			continue
		}
		contract, err := ct.GetContract(ctx, token, blockNumber)
		if err != nil {
			if !registry.IsExpectedBeforeContractsDeployed(err) {
				return nil, err
			}
			contract = nil
		}
		contracts[token] = contract
	}
	return contracts, nil
}

// GetAddresses gets the addresses for all token contracts. If onlyStables is true, addresses
// for only stable tokens are provided.
// If a contract has not been registered with the Registry as of blockNumber,
// the address is set to the zero address and no error is returned.
func (ct *CeloTokens) GetAddresses(ctx context.Context, blockNumber *big.Int, onlyStables bool) (map[CeloToken]common.Address, error) {
	addresses := make(map[CeloToken]common.Address)
	for token, tokenInfo := range CeloTokenInfos {
		if onlyStables && !tokenInfo.isStableToken {
			continue
		}
		address, err := ct.registry.GetAddressFor(ctx, blockNumber, tokenInfo.registryID)
		if err != nil {
			if !registry.IsExpectedBeforeContractsDeployed(err) {
				return nil, err
			}
			address = common.ZeroAddress
		}
		addresses[token] = address
	}
	return addresses, nil
}

// Helpers ---

// IsStableToken returns whether the provided token is a stable token
func IsStableToken(token CeloToken) bool {
	tokenInfo, ok := CeloTokenInfos[token]
	if !ok {
		return false
	}
	return tokenInfo.isStableToken
}

// GetRegistryID gets the contract's registry ID for a given token
func GetRegistryID(token CeloToken) (registry.ContractID, error) {
	tokenInfo, ok := CeloTokenInfos[token]
	if !ok {
		return "", fmt.Errorf("Token %s not found", token)
	}
	return tokenInfo.registryID, nil
}

// GetExchangeRegistryID gets the exchange contract's registry ID for a given stable token
func GetExchangeRegistryID(token CeloToken) (registry.ContractID, error) {
	tokenInfo, ok := CeloTokenInfos[token]
	if !ok {
		return "", fmt.Errorf("Token %s not found", token)
	}
	if !tokenInfo.isStableToken {
		return "", fmt.Errorf("Token %s not a stable token", token)
	}
	return tokenInfo.exchangeRegistryID, nil
}
