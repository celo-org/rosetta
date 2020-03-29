package api

type OperationResult string

const (
	OperationSuccess OperationResult = "success"
	OperationFailed  OperationResult = "failed"
)

func (or OperationResult) ToOperationStatus() OperationStatus {
	return OperationStatus{
		Successful: or == OperationSuccess,
		Status:     string(or),
	}
}

func (or OperationResult) String() string {
	return string(or)
}

type OperationKind string

const (
	OpKindTransfer OperationKind = "transfer"
	OpKindFee      OperationKind = "fee"
	OpKindMint     OperationKind = "mint"
)

func (ok OperationKind) String() string {
	return string(ok)
}
