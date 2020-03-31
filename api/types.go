package api

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

var (
	CeloGold = Currency{
		Symbol:   "cGLD",
		Decimals: 18,
	}
	CeloDollar = Currency{
		Symbol:   "cUSD",
		Decimals: 18,
	}
)

type SubmissionResult string

const (
	SubmissionSuccess SubmissionResult = "pending"
	SubmissionFailed  SubmissionResult = "rejected"
)

func (sr SubmissionResult) String() string { return string(sr) }

func (sr SubmissionResult) ToSubmissionStatus() SubmissionStatus {
	return SubmissionStatus{
		Status:     sr.String(),
		Successful: sr == SubmissionSuccess,
	}
}

type OperationResult string

const (
	OperationSuccess OperationResult = "success"
	OperationFailed  OperationResult = "failed"
)

func (or OperationResult) String() string { return string(or) }

func (or OperationResult) ToOperationStatus() OperationStatus {
	return OperationStatus{
		Successful: or == OperationSuccess,
		Status:     string(or),
	}
}

type OperationKind string

const (
	OpKindTransfer OperationKind = "transfer"
	OpKindFee      OperationKind = "fee"
	OpKindMint     OperationKind = "mint"
)

func (ok OperationKind) String() string { return string(ok) }

type Method = string

const (
	TransferMethod Method = "transfer"
)

//go:generate gencodec -type TransferMetadata -out gen_transfer_json.go

type TransferMetadata struct {
	Balance *big.Int      `json:"balance" gencodec:"required"`
	Txdata  *types.Txdata `json:"txdata"`
}

var (
	GasUpperBound = map[Method]string{
		TransferMethod: "40000",
	}
)
