package airgap

import "testing"

func Test_validateMethodSignature(t *testing.T) {
	tests := []struct {
		name      string
		methodSig string
		wantErr   bool
	}{
		{name: "valid signature with no args", methodSig: "noArgs()", wantErr: false},
		{name: "valid signature with one arg", methodSig: "deploy(address)", wantErr: false},
		{name: "valid signature with multiple args", methodSig: "deploy(address,uint8,bytes16,address)", wantErr: false},
		{name: "signature with invalid arg type", methodSig: "batchTransfer(DepositWalletTransfer[])", wantErr: true},
		{name: "closing parenthesis only", methodSig: "noArgs)", wantErr: true},
		{name: "open parenthesis only", methodSig: "noArgs(", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateMethodSignature(tt.methodSig); (err != nil) != tt.wantErr {
				t.Errorf("validateMethodSignature() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
