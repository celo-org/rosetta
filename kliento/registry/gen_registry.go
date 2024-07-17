// Code generated - DO NOT EDIT.

package registry

import (
	"context"
	"fmt"
	"math/big"

	"github.com/celo-org/rosetta/kliento/contracts"
)

// AccountsContractID is the registry identifier for 'Accounts' contract
var AccountsContractID ContractID = "Accounts"

// AttestationsContractID is the registry identifier for 'Attestations' contract
var AttestationsContractID ContractID = "Attestations"

// BlockchainParametersContractID is the registry identifier for 'BlockchainParameters' contract
var BlockchainParametersContractID ContractID = "BlockchainParameters"

// DoubleSigningSlasherContractID is the registry identifier for 'DoubleSigningSlasher' contract
var DoubleSigningSlasherContractID ContractID = "DoubleSigningSlasher"

// DowntimeSlasherContractID is the registry identifier for 'DowntimeSlasher' contract
var DowntimeSlasherContractID ContractID = "DowntimeSlasher"

// ElectionContractID is the registry identifier for 'Election' contract
var ElectionContractID ContractID = "Election"

// EpochRewardsContractID is the registry identifier for 'EpochRewards' contract
var EpochRewardsContractID ContractID = "EpochRewards"

// EscrowContractID is the registry identifier for 'Escrow' contract
var EscrowContractID ContractID = "Escrow"

// ExchangeContractID is the registry identifier for 'Exchange' contract
var ExchangeContractID ContractID = "Exchange"

// ExchangeBRLContractID is the registry identifier for 'ExchangeBRL' contract
var ExchangeBRLContractID ContractID = "ExchangeBRL"

// ExchangeEURContractID is the registry identifier for 'ExchangeEUR' contract
var ExchangeEURContractID ContractID = "ExchangeEUR"

// FeeHandlerContractID is the registry identifier for 'FeeHandler' contract
var FeeHandlerContractID ContractID = "FeeHandler"

// GasPriceMinimumContractID is the registry identifier for 'GasPriceMinimum' contract
var GasPriceMinimumContractID ContractID = "GasPriceMinimum"

// GoldTokenContractID is the registry identifier for 'GoldToken' contract
var GoldTokenContractID ContractID = "GoldToken"

// GovernanceContractID is the registry identifier for 'Governance' contract
var GovernanceContractID ContractID = "Governance"

// GovernanceSlasherContractID is the registry identifier for 'GovernanceSlasher' contract
var GovernanceSlasherContractID ContractID = "GovernanceSlasher"

// LockedGoldContractID is the registry identifier for 'LockedGold' contract
var LockedGoldContractID ContractID = "LockedGold"

// RandomContractID is the registry identifier for 'Random' contract
var RandomContractID ContractID = "Random"

// ReserveContractID is the registry identifier for 'Reserve' contract
var ReserveContractID ContractID = "Reserve"

// SortedOraclesContractID is the registry identifier for 'SortedOracles' contract
var SortedOraclesContractID ContractID = "SortedOracles"

// StableTokenContractID is the registry identifier for 'StableToken' contract
var StableTokenContractID ContractID = "StableToken"

// StableTokenBRLContractID is the registry identifier for 'StableTokenBRL' contract
var StableTokenBRLContractID ContractID = "StableTokenBRL"

// StableTokenEURContractID is the registry identifier for 'StableTokenEUR' contract
var StableTokenEURContractID ContractID = "StableTokenEUR"

// ValidatorsContractID is the registry identifier for 'Validators' contract
var ValidatorsContractID ContractID = "Validators"

// RegisteredContractIDs are all (known) registered contract identifiers
var RegisteredContractIDs = []ContractID{

	AccountsContractID,

	AttestationsContractID,

	BlockchainParametersContractID,

	DoubleSigningSlasherContractID,

	DowntimeSlasherContractID,

	ElectionContractID,

	EpochRewardsContractID,

	EscrowContractID,

	ExchangeContractID,

	ExchangeBRLContractID,

	ExchangeEURContractID,

	FeeHandlerContractID,

	GasPriceMinimumContractID,

	GoldTokenContractID,

	GovernanceContractID,

	GovernanceSlasherContractID,

	LockedGoldContractID,

	RandomContractID,

	ReserveContractID,

	SortedOraclesContractID,

	StableTokenContractID,

	StableTokenBRLContractID,

	StableTokenEURContractID,

	ValidatorsContractID,
}

type boundContracts struct {
	AccountsContract *contracts.Accounts

	AttestationsContract *contracts.Attestations

	BlockchainParametersContract *contracts.BlockchainParameters

	DoubleSigningSlasherContract *contracts.DoubleSigningSlasher

	DowntimeSlasherContract *contracts.DowntimeSlasher

	ElectionContract *contracts.Election

	EpochRewardsContract *contracts.EpochRewards

	EscrowContract *contracts.Escrow

	ExchangeContract *contracts.Exchange

	ExchangeBRLContract *contracts.Exchange

	ExchangeEURContract *contracts.Exchange

	FeeHandlerContract *contracts.FeeHandler

	GasPriceMinimumContract *contracts.GasPriceMinimum

	GoldTokenContract *contracts.GoldToken

	GovernanceContract *contracts.Governance

	GovernanceSlasherContract *contracts.GovernanceSlasher

	LockedGoldContract *contracts.LockedGold

	RandomContract *contracts.Random

	ReserveContract *contracts.Reserve

	SortedOraclesContract *contracts.SortedOracles

	StableTokenContract *contracts.StableToken

	StableTokenBRLContract *contracts.StableToken

	StableTokenEURContract *contracts.StableToken

	ValidatorsContract *contracts.Validators
}

type generatedRegistry interface {
	GetContractByID(ctx context.Context, identifier string, blockNumber *big.Int) (interface{}, error)

	GetAccountsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Accounts, error)

	GetAttestationsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Attestations, error)

	GetBlockchainParametersContract(ctx context.Context, blockNumber *big.Int) (*contracts.BlockchainParameters, error)

	GetDoubleSigningSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.DoubleSigningSlasher, error)

	GetDowntimeSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.DowntimeSlasher, error)

	GetElectionContract(ctx context.Context, blockNumber *big.Int) (*contracts.Election, error)

	GetEpochRewardsContract(ctx context.Context, blockNumber *big.Int) (*contracts.EpochRewards, error)

	GetEscrowContract(ctx context.Context, blockNumber *big.Int) (*contracts.Escrow, error)

	GetExchangeContract(ctx context.Context, blockNumber *big.Int) (*contracts.Exchange, error)

	GetExchangeBRLContract(ctx context.Context, blockNumber *big.Int) (*contracts.Exchange, error)

	GetExchangeEURContract(ctx context.Context, blockNumber *big.Int) (*contracts.Exchange, error)

	GetFeeHandlerContract(ctx context.Context, blockNumber *big.Int) (*contracts.FeeHandler, error)

	GetGasPriceMinimumContract(ctx context.Context, blockNumber *big.Int) (*contracts.GasPriceMinimum, error)

	GetGoldTokenContract(ctx context.Context, blockNumber *big.Int) (*contracts.GoldToken, error)

	GetGovernanceContract(ctx context.Context, blockNumber *big.Int) (*contracts.Governance, error)

	GetGovernanceSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.GovernanceSlasher, error)

	GetLockedGoldContract(ctx context.Context, blockNumber *big.Int) (*contracts.LockedGold, error)

	GetRandomContract(ctx context.Context, blockNumber *big.Int) (*contracts.Random, error)

	GetReserveContract(ctx context.Context, blockNumber *big.Int) (*contracts.Reserve, error)

	GetSortedOraclesContract(ctx context.Context, blockNumber *big.Int) (*contracts.SortedOracles, error)

	GetStableTokenContract(ctx context.Context, blockNumber *big.Int) (*contracts.StableToken, error)

	GetStableTokenBRLContract(ctx context.Context, blockNumber *big.Int) (*contracts.StableToken, error)

	GetStableTokenEURContract(ctx context.Context, blockNumber *big.Int) (*contracts.StableToken, error)

	GetValidatorsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Validators, error)
}

func (r *registryImpl) GetContractByID(ctx context.Context, identifier string, blockNumber *big.Int) (interface{}, error) {
	switch identifier {

	case AccountsContractID.String():
		return r.GetAccountsContract(ctx, blockNumber)

	case AttestationsContractID.String():
		return r.GetAttestationsContract(ctx, blockNumber)

	case BlockchainParametersContractID.String():
		return r.GetBlockchainParametersContract(ctx, blockNumber)

	case DoubleSigningSlasherContractID.String():
		return r.GetDoubleSigningSlasherContract(ctx, blockNumber)

	case DowntimeSlasherContractID.String():
		return r.GetDowntimeSlasherContract(ctx, blockNumber)

	case ElectionContractID.String():
		return r.GetElectionContract(ctx, blockNumber)

	case EpochRewardsContractID.String():
		return r.GetEpochRewardsContract(ctx, blockNumber)

	case EscrowContractID.String():
		return r.GetEscrowContract(ctx, blockNumber)

	case ExchangeContractID.String():
		return r.GetExchangeContract(ctx, blockNumber)

	case ExchangeBRLContractID.String():
		return r.GetExchangeBRLContract(ctx, blockNumber)

	case ExchangeEURContractID.String():
		return r.GetExchangeEURContract(ctx, blockNumber)

	case FeeHandlerContractID.String():
		return r.GetFeeHandlerContract(ctx, blockNumber)

	case GasPriceMinimumContractID.String():
		return r.GetGasPriceMinimumContract(ctx, blockNumber)

	case GoldTokenContractID.String():
		return r.GetGoldTokenContract(ctx, blockNumber)

	case GovernanceContractID.String():
		return r.GetGovernanceContract(ctx, blockNumber)

	case GovernanceSlasherContractID.String():
		return r.GetGovernanceSlasherContract(ctx, blockNumber)

	case LockedGoldContractID.String():
		return r.GetLockedGoldContract(ctx, blockNumber)

	case RandomContractID.String():
		return r.GetRandomContract(ctx, blockNumber)

	case ReserveContractID.String():
		return r.GetReserveContract(ctx, blockNumber)

	case SortedOraclesContractID.String():
		return r.GetSortedOraclesContract(ctx, blockNumber)

	case StableTokenContractID.String():
		return r.GetStableTokenContract(ctx, blockNumber)

	case StableTokenBRLContractID.String():
		return r.GetStableTokenBRLContract(ctx, blockNumber)

	case StableTokenEURContractID.String():
		return r.GetStableTokenEURContract(ctx, blockNumber)

	case ValidatorsContractID.String():
		return r.GetValidatorsContract(ctx, blockNumber)

	}
	return nil, fmt.Errorf("identifier '%s' not found", identifier)
}

func (r *registryImpl) GetAccountsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Accounts, error) {
	identifier := AccountsContractID.String()
	if r.AccountsContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, AccountsContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewAccounts(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.AccountsContract = contract
	}
	return r.AccountsContract, nil
}

func (r *registryImpl) GetAttestationsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Attestations, error) {
	identifier := AttestationsContractID.String()
	if r.AttestationsContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, AttestationsContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewAttestations(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.AttestationsContract = contract
	}
	return r.AttestationsContract, nil
}

func (r *registryImpl) GetBlockchainParametersContract(ctx context.Context, blockNumber *big.Int) (*contracts.BlockchainParameters, error) {
	identifier := BlockchainParametersContractID.String()
	if r.BlockchainParametersContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, BlockchainParametersContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewBlockchainParameters(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.BlockchainParametersContract = contract
	}
	return r.BlockchainParametersContract, nil
}

func (r *registryImpl) GetDoubleSigningSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.DoubleSigningSlasher, error) {
	identifier := DoubleSigningSlasherContractID.String()
	if r.DoubleSigningSlasherContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, DoubleSigningSlasherContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewDoubleSigningSlasher(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.DoubleSigningSlasherContract = contract
	}
	return r.DoubleSigningSlasherContract, nil
}

func (r *registryImpl) GetDowntimeSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.DowntimeSlasher, error) {
	identifier := DowntimeSlasherContractID.String()
	if r.DowntimeSlasherContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, DowntimeSlasherContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewDowntimeSlasher(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.DowntimeSlasherContract = contract
	}
	return r.DowntimeSlasherContract, nil
}

func (r *registryImpl) GetElectionContract(ctx context.Context, blockNumber *big.Int) (*contracts.Election, error) {
	identifier := ElectionContractID.String()
	if r.ElectionContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, ElectionContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewElection(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.ElectionContract = contract
	}
	return r.ElectionContract, nil
}

func (r *registryImpl) GetEpochRewardsContract(ctx context.Context, blockNumber *big.Int) (*contracts.EpochRewards, error) {
	identifier := EpochRewardsContractID.String()
	if r.EpochRewardsContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, EpochRewardsContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewEpochRewards(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.EpochRewardsContract = contract
	}
	return r.EpochRewardsContract, nil
}

func (r *registryImpl) GetEscrowContract(ctx context.Context, blockNumber *big.Int) (*contracts.Escrow, error) {
	identifier := EscrowContractID.String()
	if r.EscrowContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, EscrowContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewEscrow(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.EscrowContract = contract
	}
	return r.EscrowContract, nil
}

func (r *registryImpl) GetExchangeContract(ctx context.Context, blockNumber *big.Int) (*contracts.Exchange, error) {
	identifier := ExchangeContractID.String()
	if r.ExchangeContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, ExchangeContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewExchange(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.ExchangeContract = contract
	}
	return r.ExchangeContract, nil
}

func (r *registryImpl) GetExchangeBRLContract(ctx context.Context, blockNumber *big.Int) (*contracts.Exchange, error) {
	identifier := ExchangeBRLContractID.String()
	if r.ExchangeBRLContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, ExchangeBRLContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewExchange(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.ExchangeBRLContract = contract
	}
	return r.ExchangeBRLContract, nil
}

func (r *registryImpl) GetExchangeEURContract(ctx context.Context, blockNumber *big.Int) (*contracts.Exchange, error) {
	identifier := ExchangeEURContractID.String()
	if r.ExchangeEURContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, ExchangeEURContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewExchange(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.ExchangeEURContract = contract
	}
	return r.ExchangeEURContract, nil
}

func (r *registryImpl) GetFeeHandlerContract(ctx context.Context, blockNumber *big.Int) (*contracts.FeeHandler, error) {
	identifier := FeeHandlerContractID.String()
	if r.FeeHandlerContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, FeeHandlerContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewFeeHandler(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.FeeHandlerContract = contract
	}
	return r.FeeHandlerContract, nil
}

func (r *registryImpl) GetGasPriceMinimumContract(ctx context.Context, blockNumber *big.Int) (*contracts.GasPriceMinimum, error) {
	identifier := GasPriceMinimumContractID.String()
	if r.GasPriceMinimumContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, GasPriceMinimumContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewGasPriceMinimum(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.GasPriceMinimumContract = contract
	}
	return r.GasPriceMinimumContract, nil
}

func (r *registryImpl) GetGoldTokenContract(ctx context.Context, blockNumber *big.Int) (*contracts.GoldToken, error) {
	identifier := GoldTokenContractID.String()
	if r.GoldTokenContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, GoldTokenContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewGoldToken(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.GoldTokenContract = contract
	}
	return r.GoldTokenContract, nil
}

func (r *registryImpl) GetGovernanceContract(ctx context.Context, blockNumber *big.Int) (*contracts.Governance, error) {
	identifier := GovernanceContractID.String()
	if r.GovernanceContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, GovernanceContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewGovernance(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.GovernanceContract = contract
	}
	return r.GovernanceContract, nil
}

func (r *registryImpl) GetGovernanceSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.GovernanceSlasher, error) {
	identifier := GovernanceSlasherContractID.String()
	if r.GovernanceSlasherContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, GovernanceSlasherContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewGovernanceSlasher(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.GovernanceSlasherContract = contract
	}
	return r.GovernanceSlasherContract, nil
}

func (r *registryImpl) GetLockedGoldContract(ctx context.Context, blockNumber *big.Int) (*contracts.LockedGold, error) {
	identifier := LockedGoldContractID.String()
	if r.LockedGoldContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, LockedGoldContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewLockedGold(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.LockedGoldContract = contract
	}
	return r.LockedGoldContract, nil
}

func (r *registryImpl) GetRandomContract(ctx context.Context, blockNumber *big.Int) (*contracts.Random, error) {
	identifier := RandomContractID.String()
	if r.RandomContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, RandomContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewRandom(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.RandomContract = contract
	}
	return r.RandomContract, nil
}

func (r *registryImpl) GetReserveContract(ctx context.Context, blockNumber *big.Int) (*contracts.Reserve, error) {
	identifier := ReserveContractID.String()
	if r.ReserveContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, ReserveContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewReserve(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.ReserveContract = contract
	}
	return r.ReserveContract, nil
}

func (r *registryImpl) GetSortedOraclesContract(ctx context.Context, blockNumber *big.Int) (*contracts.SortedOracles, error) {
	identifier := SortedOraclesContractID.String()
	if r.SortedOraclesContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, SortedOraclesContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewSortedOracles(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.SortedOraclesContract = contract
	}
	return r.SortedOraclesContract, nil
}

func (r *registryImpl) GetStableTokenContract(ctx context.Context, blockNumber *big.Int) (*contracts.StableToken, error) {
	identifier := StableTokenContractID.String()
	if r.StableTokenContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, StableTokenContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewStableToken(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.StableTokenContract = contract
	}
	return r.StableTokenContract, nil
}

func (r *registryImpl) GetStableTokenBRLContract(ctx context.Context, blockNumber *big.Int) (*contracts.StableToken, error) {
	identifier := StableTokenBRLContractID.String()
	if r.StableTokenBRLContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, StableTokenBRLContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewStableToken(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.StableTokenBRLContract = contract
	}
	return r.StableTokenBRLContract, nil
}

func (r *registryImpl) GetStableTokenEURContract(ctx context.Context, blockNumber *big.Int) (*contracts.StableToken, error) {
	identifier := StableTokenEURContractID.String()
	if r.StableTokenEURContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, StableTokenEURContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewStableToken(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.StableTokenEURContract = contract
	}
	return r.StableTokenEURContract, nil
}

func (r *registryImpl) GetValidatorsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Validators, error) {
	identifier := ValidatorsContractID.String()
	if r.ValidatorsContract == nil || r.cache.isDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, ValidatorsContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewValidators(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.ValidatorsContract = contract
	}
	return r.ValidatorsContract, nil
}
