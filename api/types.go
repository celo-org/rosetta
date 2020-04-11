package api

import (
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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

func GetOperationStatus(success bool) OperationResult {
	if success {
		return OperationSuccess
	} else {
		return OperationFailed
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

// TODO: source constants from ABI
const (
	TransferMethod      Method = "transfer"
	CreateAccountMethod Method = "createAccount"
)

var (
	DummyAddress = common.HexToAddress("abc")
)

//go:generate gencodec -type TransactionMetadata -out gen_transaction_metadata_json.go

type TransactionMetadata struct {
	ABIMethod           *abi.Method     `json:"abiMethod,omitempty"`    // nil means value transfer only
	To                  *common.Address `json:"to,omitempty" rlp:"nil"` // nil means contract creation
	Nonce               uint64          `json:"nonce"    gencodec:"required"`
	GasPrice            *big.Int        `json:"gasPrice" gencodec:"required"`
	GasLimit            uint64          `json:"gasLimit"      gencodec:"required"`
	GatewayFeeRecipient *common.Address `json:"gatewayFeeRecipient" rlp:"nil"` // nil means no gateway fee is paid
	GatewayFee          *big.Int        `json:"gatewayFee" rlp:"nil"`          // nil means no gateway fee is paid
}

func (txm *TransactionMetadata) asMessage() *ethereum.CallMsg {
	return &ethereum.CallMsg{
		GasPrice:            txm.GasPrice,
		GatewayFee:          txm.GatewayFee,
		GatewayFeeRecipient: txm.GatewayFeeRecipient,
		To:                  txm.To,
	}
}
