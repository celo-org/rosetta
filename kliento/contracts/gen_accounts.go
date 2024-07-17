// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/celo-org/celo-blockchain"
	"github.com/celo-org/celo-blockchain/accounts/abi"
	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AccountsMetaData contains all meta data concerning the Accounts contract.
var AccountsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"test\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"dataEncryptionKey\",\"type\":\"bytes\"}],\"name\":\"AccountDataEncryptionKeySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataURL\",\"type\":\"string\"}],\"name\":\"AccountMetadataURLSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"AccountNameSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"}],\"name\":\"AccountWalletAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"AttestationSignerAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldSigner\",\"type\":\"address\"}],\"name\":\"AttestationSignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldSigner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"DefaultSignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"DefaultSignerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldSigner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"IndexedSignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"IndexedSignerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldSigner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"LegacySignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"LegacySignerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"OffchainStorageRootAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"OffchainStorageRootRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fraction\",\"type\":\"uint256\"}],\"name\":\"PaymentDelegationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"SignerAuthorizationCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"SignerAuthorizationStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"SignerAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldSigner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"SignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"ValidatorSignerAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldSigner\",\"type\":\"address\"}],\"name\":\"ValidatorSignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"VoteSignerAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldSigner\",\"type\":\"address\"}],\"name\":\"VoteSignerRemoved\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP712_AUTHORIZE_SIGNER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedBy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eip712DomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"offchainStorageRoots\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersionNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setEip712DomainSeparator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"dataEncryptionKey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"setAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"setWalletAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"dataEncryptionKey\",\"type\":\"bytes\"}],\"name\":\"setAccountDataEncryptionKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"metadataURL\",\"type\":\"string\"}],\"name\":\"setMetadataURL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"addStorageRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeStorageRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getOffchainStorageRoots\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fraction\",\"type\":\"uint256\"}],\"name\":\"setPaymentDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deletePaymentDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPaymentDelegation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"setIndexedSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"authorizeSignerWithSignature\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"authorizeVoteSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"authorizeValidatorSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"ecdsaPublicKey\",\"type\":\"bytes\"}],\"name\":\"authorizeValidatorSignerWithPublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"ecdsaPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"authorizeValidatorSignerWithKeys\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"authorizeAttestationSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"authorizeSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"completeSignerAuthorization\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"isLegacySigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"isDefaultSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"isIndexedSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"removeDefaultSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"removeIndexedSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"removeSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeVoteSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeValidatorSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeAttestationSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"attestationSignerToAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"validatorSignerToAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"voteSignerToAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"signerToAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"isLegacyRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getLegacySigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getDefaultSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getIndexedSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getVoteSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getValidatorSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAttestationSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"hasLegacySigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"hasDefaultSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"hasIndexedSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"}],\"name\":\"hasAuthorizedSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasAuthorizedVoteSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasAuthorizedValidatorSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasAuthorizedAttestationSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getMetadataURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accountsToQuery\",\"type\":\"address[]\"}],\"name\":\"batchGetMetadataURL\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDataEncryptionKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getWalletAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isAuthorizedSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"getRoleAuthorizationSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AccountsABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountsMetaData.ABI instead.
var AccountsABI = AccountsMetaData.ABI

// Accounts is an auto generated Go binding around an Ethereum contract.
type Accounts struct {
	AccountsCaller     // Read-only binding to the contract
	AccountsTransactor // Write-only binding to the contract
	AccountsFilterer   // Log filterer for contract events
}

// AccountsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountsSession struct {
	Contract     *Accounts         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountsCallerSession struct {
	Contract *AccountsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AccountsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountsTransactorSession struct {
	Contract     *AccountsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AccountsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountsRaw struct {
	Contract *Accounts // Generic contract binding to access the raw methods on
}

// AccountsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountsCallerRaw struct {
	Contract *AccountsCaller // Generic read-only contract binding to access the raw methods on
}

// AccountsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountsTransactorRaw struct {
	Contract *AccountsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccounts creates a new instance of Accounts, bound to a specific deployed contract.
func NewAccounts(address common.Address, backend bind.ContractBackend) (*Accounts, error) {
	contract, err := bindAccounts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Accounts{AccountsCaller: AccountsCaller{contract: contract}, AccountsTransactor: AccountsTransactor{contract: contract}, AccountsFilterer: AccountsFilterer{contract: contract}}, nil
}

// NewAccountsCaller creates a new read-only instance of Accounts, bound to a specific deployed contract.
func NewAccountsCaller(address common.Address, caller bind.ContractCaller) (*AccountsCaller, error) {
	contract, err := bindAccounts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsCaller{contract: contract}, nil
}

// NewAccountsTransactor creates a new write-only instance of Accounts, bound to a specific deployed contract.
func NewAccountsTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountsTransactor, error) {
	contract, err := bindAccounts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsTransactor{contract: contract}, nil
}

// NewAccountsFilterer creates a new log filterer instance of Accounts, bound to a specific deployed contract.
func NewAccountsFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountsFilterer, error) {
	contract, err := bindAccounts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountsFilterer{contract: contract}, nil
}

// bindAccounts binds a generic wrapper to an already deployed contract.
func bindAccounts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseAccountsABI parses the ABI
func ParseAccountsABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountsABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accounts *AccountsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accounts.Contract.AccountsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accounts *AccountsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.Contract.AccountsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accounts *AccountsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accounts.Contract.AccountsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accounts *AccountsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accounts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accounts *AccountsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accounts *AccountsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accounts.Contract.contract.Transact(opts, method, params...)
}

// EIP712AUTHORIZESIGNERTYPEHASH is a free data retrieval call binding the contract method 0xf0c56584.
//
// Solidity: function EIP712_AUTHORIZE_SIGNER_TYPEHASH() view returns(bytes32)
func (_Accounts *AccountsCaller) EIP712AUTHORIZESIGNERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "EIP712_AUTHORIZE_SIGNER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712AUTHORIZESIGNERTYPEHASH is a free data retrieval call binding the contract method 0xf0c56584.
//
// Solidity: function EIP712_AUTHORIZE_SIGNER_TYPEHASH() view returns(bytes32)
func (_Accounts *AccountsSession) EIP712AUTHORIZESIGNERTYPEHASH() ([32]byte, error) {
	return _Accounts.Contract.EIP712AUTHORIZESIGNERTYPEHASH(&_Accounts.CallOpts)
}

// EIP712AUTHORIZESIGNERTYPEHASH is a free data retrieval call binding the contract method 0xf0c56584.
//
// Solidity: function EIP712_AUTHORIZE_SIGNER_TYPEHASH() view returns(bytes32)
func (_Accounts *AccountsCallerSession) EIP712AUTHORIZESIGNERTYPEHASH() ([32]byte, error) {
	return _Accounts.Contract.EIP712AUTHORIZESIGNERTYPEHASH(&_Accounts.CallOpts)
}

// AttestationSignerToAccount is a free data retrieval call binding the contract method 0x7b2434cb.
//
// Solidity: function attestationSignerToAccount(address signer) view returns(address)
func (_Accounts *AccountsCaller) AttestationSignerToAccount(opts *bind.CallOpts, signer common.Address) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "attestationSignerToAccount", signer)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AttestationSignerToAccount is a free data retrieval call binding the contract method 0x7b2434cb.
//
// Solidity: function attestationSignerToAccount(address signer) view returns(address)
func (_Accounts *AccountsSession) AttestationSignerToAccount(signer common.Address) (common.Address, error) {
	return _Accounts.Contract.AttestationSignerToAccount(&_Accounts.CallOpts, signer)
}

// AttestationSignerToAccount is a free data retrieval call binding the contract method 0x7b2434cb.
//
// Solidity: function attestationSignerToAccount(address signer) view returns(address)
func (_Accounts *AccountsCallerSession) AttestationSignerToAccount(signer common.Address) (common.Address, error) {
	return _Accounts.Contract.AttestationSignerToAccount(&_Accounts.CallOpts, signer)
}

// AuthorizedBy is a free data retrieval call binding the contract method 0xb5a664c2.
//
// Solidity: function authorizedBy(address ) view returns(address)
func (_Accounts *AccountsCaller) AuthorizedBy(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "authorizedBy", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuthorizedBy is a free data retrieval call binding the contract method 0xb5a664c2.
//
// Solidity: function authorizedBy(address ) view returns(address)
func (_Accounts *AccountsSession) AuthorizedBy(arg0 common.Address) (common.Address, error) {
	return _Accounts.Contract.AuthorizedBy(&_Accounts.CallOpts, arg0)
}

// AuthorizedBy is a free data retrieval call binding the contract method 0xb5a664c2.
//
// Solidity: function authorizedBy(address ) view returns(address)
func (_Accounts *AccountsCallerSession) AuthorizedBy(arg0 common.Address) (common.Address, error) {
	return _Accounts.Contract.AuthorizedBy(&_Accounts.CallOpts, arg0)
}

// BatchGetMetadataURL is a free data retrieval call binding the contract method 0x8adaf96f.
//
// Solidity: function batchGetMetadataURL(address[] accountsToQuery) view returns(uint256[], bytes)
func (_Accounts *AccountsCaller) BatchGetMetadataURL(opts *bind.CallOpts, accountsToQuery []common.Address) ([]*big.Int, []byte, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "batchGetMetadataURL", accountsToQuery)

	if err != nil {
		return *new([]*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// BatchGetMetadataURL is a free data retrieval call binding the contract method 0x8adaf96f.
//
// Solidity: function batchGetMetadataURL(address[] accountsToQuery) view returns(uint256[], bytes)
func (_Accounts *AccountsSession) BatchGetMetadataURL(accountsToQuery []common.Address) ([]*big.Int, []byte, error) {
	return _Accounts.Contract.BatchGetMetadataURL(&_Accounts.CallOpts, accountsToQuery)
}

// BatchGetMetadataURL is a free data retrieval call binding the contract method 0x8adaf96f.
//
// Solidity: function batchGetMetadataURL(address[] accountsToQuery) view returns(uint256[], bytes)
func (_Accounts *AccountsCallerSession) BatchGetMetadataURL(accountsToQuery []common.Address) ([]*big.Int, []byte, error) {
	return _Accounts.Contract.BatchGetMetadataURL(&_Accounts.CallOpts, accountsToQuery)
}

// Eip712DomainSeparator is a free data retrieval call binding the contract method 0x5b07fdd8.
//
// Solidity: function eip712DomainSeparator() view returns(bytes32)
func (_Accounts *AccountsCaller) Eip712DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "eip712DomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Eip712DomainSeparator is a free data retrieval call binding the contract method 0x5b07fdd8.
//
// Solidity: function eip712DomainSeparator() view returns(bytes32)
func (_Accounts *AccountsSession) Eip712DomainSeparator() ([32]byte, error) {
	return _Accounts.Contract.Eip712DomainSeparator(&_Accounts.CallOpts)
}

// Eip712DomainSeparator is a free data retrieval call binding the contract method 0x5b07fdd8.
//
// Solidity: function eip712DomainSeparator() view returns(bytes32)
func (_Accounts *AccountsCallerSession) Eip712DomainSeparator() ([32]byte, error) {
	return _Accounts.Contract.Eip712DomainSeparator(&_Accounts.CallOpts)
}

// GetAttestationSigner is a free data retrieval call binding the contract method 0x61bab1ae.
//
// Solidity: function getAttestationSigner(address account) view returns(address)
func (_Accounts *AccountsCaller) GetAttestationSigner(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getAttestationSigner", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAttestationSigner is a free data retrieval call binding the contract method 0x61bab1ae.
//
// Solidity: function getAttestationSigner(address account) view returns(address)
func (_Accounts *AccountsSession) GetAttestationSigner(account common.Address) (common.Address, error) {
	return _Accounts.Contract.GetAttestationSigner(&_Accounts.CallOpts, account)
}

// GetAttestationSigner is a free data retrieval call binding the contract method 0x61bab1ae.
//
// Solidity: function getAttestationSigner(address account) view returns(address)
func (_Accounts *AccountsCallerSession) GetAttestationSigner(account common.Address) (common.Address, error) {
	return _Accounts.Contract.GetAttestationSigner(&_Accounts.CallOpts, account)
}

// GetDataEncryptionKey is a free data retrieval call binding the contract method 0xae32fa0e.
//
// Solidity: function getDataEncryptionKey(address account) view returns(bytes)
func (_Accounts *AccountsCaller) GetDataEncryptionKey(opts *bind.CallOpts, account common.Address) ([]byte, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getDataEncryptionKey", account)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDataEncryptionKey is a free data retrieval call binding the contract method 0xae32fa0e.
//
// Solidity: function getDataEncryptionKey(address account) view returns(bytes)
func (_Accounts *AccountsSession) GetDataEncryptionKey(account common.Address) ([]byte, error) {
	return _Accounts.Contract.GetDataEncryptionKey(&_Accounts.CallOpts, account)
}

// GetDataEncryptionKey is a free data retrieval call binding the contract method 0xae32fa0e.
//
// Solidity: function getDataEncryptionKey(address account) view returns(bytes)
func (_Accounts *AccountsCallerSession) GetDataEncryptionKey(account common.Address) ([]byte, error) {
	return _Accounts.Contract.GetDataEncryptionKey(&_Accounts.CallOpts, account)
}

// GetDefaultSigner is a free data retrieval call binding the contract method 0x5b6d9004.
//
// Solidity: function getDefaultSigner(address account, bytes32 role) view returns(address)
func (_Accounts *AccountsCaller) GetDefaultSigner(opts *bind.CallOpts, account common.Address, role [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getDefaultSigner", account, role)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDefaultSigner is a free data retrieval call binding the contract method 0x5b6d9004.
//
// Solidity: function getDefaultSigner(address account, bytes32 role) view returns(address)
func (_Accounts *AccountsSession) GetDefaultSigner(account common.Address, role [32]byte) (common.Address, error) {
	return _Accounts.Contract.GetDefaultSigner(&_Accounts.CallOpts, account, role)
}

// GetDefaultSigner is a free data retrieval call binding the contract method 0x5b6d9004.
//
// Solidity: function getDefaultSigner(address account, bytes32 role) view returns(address)
func (_Accounts *AccountsCallerSession) GetDefaultSigner(account common.Address, role [32]byte) (common.Address, error) {
	return _Accounts.Contract.GetDefaultSigner(&_Accounts.CallOpts, account, role)
}

// GetIndexedSigner is a free data retrieval call binding the contract method 0x87affe68.
//
// Solidity: function getIndexedSigner(address account, bytes32 role) view returns(address)
func (_Accounts *AccountsCaller) GetIndexedSigner(opts *bind.CallOpts, account common.Address, role [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getIndexedSigner", account, role)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIndexedSigner is a free data retrieval call binding the contract method 0x87affe68.
//
// Solidity: function getIndexedSigner(address account, bytes32 role) view returns(address)
func (_Accounts *AccountsSession) GetIndexedSigner(account common.Address, role [32]byte) (common.Address, error) {
	return _Accounts.Contract.GetIndexedSigner(&_Accounts.CallOpts, account, role)
}

// GetIndexedSigner is a free data retrieval call binding the contract method 0x87affe68.
//
// Solidity: function getIndexedSigner(address account, bytes32 role) view returns(address)
func (_Accounts *AccountsCallerSession) GetIndexedSigner(account common.Address, role [32]byte) (common.Address, error) {
	return _Accounts.Contract.GetIndexedSigner(&_Accounts.CallOpts, account, role)
}

// GetLegacySigner is a free data retrieval call binding the contract method 0x8bceca58.
//
// Solidity: function getLegacySigner(address _account, bytes32 role) view returns(address)
func (_Accounts *AccountsCaller) GetLegacySigner(opts *bind.CallOpts, _account common.Address, role [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getLegacySigner", _account, role)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLegacySigner is a free data retrieval call binding the contract method 0x8bceca58.
//
// Solidity: function getLegacySigner(address _account, bytes32 role) view returns(address)
func (_Accounts *AccountsSession) GetLegacySigner(_account common.Address, role [32]byte) (common.Address, error) {
	return _Accounts.Contract.GetLegacySigner(&_Accounts.CallOpts, _account, role)
}

// GetLegacySigner is a free data retrieval call binding the contract method 0x8bceca58.
//
// Solidity: function getLegacySigner(address _account, bytes32 role) view returns(address)
func (_Accounts *AccountsCallerSession) GetLegacySigner(_account common.Address, role [32]byte) (common.Address, error) {
	return _Accounts.Contract.GetLegacySigner(&_Accounts.CallOpts, _account, role)
}

// GetMetadataURL is a free data retrieval call binding the contract method 0xa8ae1a3d.
//
// Solidity: function getMetadataURL(address account) view returns(string)
func (_Accounts *AccountsCaller) GetMetadataURL(opts *bind.CallOpts, account common.Address) (string, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getMetadataURL", account)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMetadataURL is a free data retrieval call binding the contract method 0xa8ae1a3d.
//
// Solidity: function getMetadataURL(address account) view returns(string)
func (_Accounts *AccountsSession) GetMetadataURL(account common.Address) (string, error) {
	return _Accounts.Contract.GetMetadataURL(&_Accounts.CallOpts, account)
}

// GetMetadataURL is a free data retrieval call binding the contract method 0xa8ae1a3d.
//
// Solidity: function getMetadataURL(address account) view returns(string)
func (_Accounts *AccountsCallerSession) GetMetadataURL(account common.Address) (string, error) {
	return _Accounts.Contract.GetMetadataURL(&_Accounts.CallOpts, account)
}

// GetName is a free data retrieval call binding the contract method 0x5fd4b08a.
//
// Solidity: function getName(address account) view returns(string)
func (_Accounts *AccountsCaller) GetName(opts *bind.CallOpts, account common.Address) (string, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getName", account)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x5fd4b08a.
//
// Solidity: function getName(address account) view returns(string)
func (_Accounts *AccountsSession) GetName(account common.Address) (string, error) {
	return _Accounts.Contract.GetName(&_Accounts.CallOpts, account)
}

// GetName is a free data retrieval call binding the contract method 0x5fd4b08a.
//
// Solidity: function getName(address account) view returns(string)
func (_Accounts *AccountsCallerSession) GetName(account common.Address) (string, error) {
	return _Accounts.Contract.GetName(&_Accounts.CallOpts, account)
}

// GetOffchainStorageRoots is a free data retrieval call binding the contract method 0x289a1318.
//
// Solidity: function getOffchainStorageRoots(address account) view returns(bytes, uint256[])
func (_Accounts *AccountsCaller) GetOffchainStorageRoots(opts *bind.CallOpts, account common.Address) ([]byte, []*big.Int, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getOffchainStorageRoots", account)

	if err != nil {
		return *new([]byte), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetOffchainStorageRoots is a free data retrieval call binding the contract method 0x289a1318.
//
// Solidity: function getOffchainStorageRoots(address account) view returns(bytes, uint256[])
func (_Accounts *AccountsSession) GetOffchainStorageRoots(account common.Address) ([]byte, []*big.Int, error) {
	return _Accounts.Contract.GetOffchainStorageRoots(&_Accounts.CallOpts, account)
}

// GetOffchainStorageRoots is a free data retrieval call binding the contract method 0x289a1318.
//
// Solidity: function getOffchainStorageRoots(address account) view returns(bytes, uint256[])
func (_Accounts *AccountsCallerSession) GetOffchainStorageRoots(account common.Address) ([]byte, []*big.Int, error) {
	return _Accounts.Contract.GetOffchainStorageRoots(&_Accounts.CallOpts, account)
}

// GetPaymentDelegation is a free data retrieval call binding the contract method 0x9f024f4b.
//
// Solidity: function getPaymentDelegation(address account) view returns(address, uint256)
func (_Accounts *AccountsCaller) GetPaymentDelegation(opts *bind.CallOpts, account common.Address) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getPaymentDelegation", account)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPaymentDelegation is a free data retrieval call binding the contract method 0x9f024f4b.
//
// Solidity: function getPaymentDelegation(address account) view returns(address, uint256)
func (_Accounts *AccountsSession) GetPaymentDelegation(account common.Address) (common.Address, *big.Int, error) {
	return _Accounts.Contract.GetPaymentDelegation(&_Accounts.CallOpts, account)
}

// GetPaymentDelegation is a free data retrieval call binding the contract method 0x9f024f4b.
//
// Solidity: function getPaymentDelegation(address account) view returns(address, uint256)
func (_Accounts *AccountsCallerSession) GetPaymentDelegation(account common.Address) (common.Address, *big.Int, error) {
	return _Accounts.Contract.GetPaymentDelegation(&_Accounts.CallOpts, account)
}

// GetRoleAuthorizationSigner is a free data retrieval call binding the contract method 0xb6c66625.
//
// Solidity: function getRoleAuthorizationSigner(address account, address signer, bytes32 role, uint8 v, bytes32 r, bytes32 s) view returns(address)
func (_Accounts *AccountsCaller) GetRoleAuthorizationSigner(opts *bind.CallOpts, account common.Address, signer common.Address, role [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getRoleAuthorizationSigner", account, signer, role, v, r, s)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleAuthorizationSigner is a free data retrieval call binding the contract method 0xb6c66625.
//
// Solidity: function getRoleAuthorizationSigner(address account, address signer, bytes32 role, uint8 v, bytes32 r, bytes32 s) view returns(address)
func (_Accounts *AccountsSession) GetRoleAuthorizationSigner(account common.Address, signer common.Address, role [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Accounts.Contract.GetRoleAuthorizationSigner(&_Accounts.CallOpts, account, signer, role, v, r, s)
}

// GetRoleAuthorizationSigner is a free data retrieval call binding the contract method 0xb6c66625.
//
// Solidity: function getRoleAuthorizationSigner(address account, address signer, bytes32 role, uint8 v, bytes32 r, bytes32 s) view returns(address)
func (_Accounts *AccountsCallerSession) GetRoleAuthorizationSigner(account common.Address, signer common.Address, role [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Accounts.Contract.GetRoleAuthorizationSigner(&_Accounts.CallOpts, account, signer, role, v, r, s)
}

// GetValidatorSigner is a free data retrieval call binding the contract method 0x4ce38b5f.
//
// Solidity: function getValidatorSigner(address account) view returns(address)
func (_Accounts *AccountsCaller) GetValidatorSigner(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getValidatorSigner", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetValidatorSigner is a free data retrieval call binding the contract method 0x4ce38b5f.
//
// Solidity: function getValidatorSigner(address account) view returns(address)
func (_Accounts *AccountsSession) GetValidatorSigner(account common.Address) (common.Address, error) {
	return _Accounts.Contract.GetValidatorSigner(&_Accounts.CallOpts, account)
}

// GetValidatorSigner is a free data retrieval call binding the contract method 0x4ce38b5f.
//
// Solidity: function getValidatorSigner(address account) view returns(address)
func (_Accounts *AccountsCallerSession) GetValidatorSigner(account common.Address) (common.Address, error) {
	return _Accounts.Contract.GetValidatorSigner(&_Accounts.CallOpts, account)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Accounts *AccountsCaller) GetVersionNumber(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getVersionNumber")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Accounts *AccountsSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Accounts.Contract.GetVersionNumber(&_Accounts.CallOpts)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Accounts *AccountsCallerSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Accounts.Contract.GetVersionNumber(&_Accounts.CallOpts)
}

// GetVoteSigner is a free data retrieval call binding the contract method 0x41ddd880.
//
// Solidity: function getVoteSigner(address account) view returns(address)
func (_Accounts *AccountsCaller) GetVoteSigner(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getVoteSigner", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVoteSigner is a free data retrieval call binding the contract method 0x41ddd880.
//
// Solidity: function getVoteSigner(address account) view returns(address)
func (_Accounts *AccountsSession) GetVoteSigner(account common.Address) (common.Address, error) {
	return _Accounts.Contract.GetVoteSigner(&_Accounts.CallOpts, account)
}

// GetVoteSigner is a free data retrieval call binding the contract method 0x41ddd880.
//
// Solidity: function getVoteSigner(address account) view returns(address)
func (_Accounts *AccountsCallerSession) GetVoteSigner(account common.Address) (common.Address, error) {
	return _Accounts.Contract.GetVoteSigner(&_Accounts.CallOpts, account)
}

// GetWalletAddress is a free data retrieval call binding the contract method 0x1fd9afa5.
//
// Solidity: function getWalletAddress(address account) view returns(address)
func (_Accounts *AccountsCaller) GetWalletAddress(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getWalletAddress", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWalletAddress is a free data retrieval call binding the contract method 0x1fd9afa5.
//
// Solidity: function getWalletAddress(address account) view returns(address)
func (_Accounts *AccountsSession) GetWalletAddress(account common.Address) (common.Address, error) {
	return _Accounts.Contract.GetWalletAddress(&_Accounts.CallOpts, account)
}

// GetWalletAddress is a free data retrieval call binding the contract method 0x1fd9afa5.
//
// Solidity: function getWalletAddress(address account) view returns(address)
func (_Accounts *AccountsCallerSession) GetWalletAddress(account common.Address) (common.Address, error) {
	return _Accounts.Contract.GetWalletAddress(&_Accounts.CallOpts, account)
}

// HasAuthorizedAttestationSigner is a free data retrieval call binding the contract method 0xc2e0ee20.
//
// Solidity: function hasAuthorizedAttestationSigner(address account) view returns(bool)
func (_Accounts *AccountsCaller) HasAuthorizedAttestationSigner(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "hasAuthorizedAttestationSigner", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAuthorizedAttestationSigner is a free data retrieval call binding the contract method 0xc2e0ee20.
//
// Solidity: function hasAuthorizedAttestationSigner(address account) view returns(bool)
func (_Accounts *AccountsSession) HasAuthorizedAttestationSigner(account common.Address) (bool, error) {
	return _Accounts.Contract.HasAuthorizedAttestationSigner(&_Accounts.CallOpts, account)
}

// HasAuthorizedAttestationSigner is a free data retrieval call binding the contract method 0xc2e0ee20.
//
// Solidity: function hasAuthorizedAttestationSigner(address account) view returns(bool)
func (_Accounts *AccountsCallerSession) HasAuthorizedAttestationSigner(account common.Address) (bool, error) {
	return _Accounts.Contract.HasAuthorizedAttestationSigner(&_Accounts.CallOpts, account)
}

// HasAuthorizedSigner is a free data retrieval call binding the contract method 0xba40e4f6.
//
// Solidity: function hasAuthorizedSigner(address account, string role) view returns(bool)
func (_Accounts *AccountsCaller) HasAuthorizedSigner(opts *bind.CallOpts, account common.Address, role string) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "hasAuthorizedSigner", account, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAuthorizedSigner is a free data retrieval call binding the contract method 0xba40e4f6.
//
// Solidity: function hasAuthorizedSigner(address account, string role) view returns(bool)
func (_Accounts *AccountsSession) HasAuthorizedSigner(account common.Address, role string) (bool, error) {
	return _Accounts.Contract.HasAuthorizedSigner(&_Accounts.CallOpts, account, role)
}

// HasAuthorizedSigner is a free data retrieval call binding the contract method 0xba40e4f6.
//
// Solidity: function hasAuthorizedSigner(address account, string role) view returns(bool)
func (_Accounts *AccountsCallerSession) HasAuthorizedSigner(account common.Address, role string) (bool, error) {
	return _Accounts.Contract.HasAuthorizedSigner(&_Accounts.CallOpts, account, role)
}

// HasAuthorizedValidatorSigner is a free data retrieval call binding the contract method 0x0127dbed.
//
// Solidity: function hasAuthorizedValidatorSigner(address account) view returns(bool)
func (_Accounts *AccountsCaller) HasAuthorizedValidatorSigner(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "hasAuthorizedValidatorSigner", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAuthorizedValidatorSigner is a free data retrieval call binding the contract method 0x0127dbed.
//
// Solidity: function hasAuthorizedValidatorSigner(address account) view returns(bool)
func (_Accounts *AccountsSession) HasAuthorizedValidatorSigner(account common.Address) (bool, error) {
	return _Accounts.Contract.HasAuthorizedValidatorSigner(&_Accounts.CallOpts, account)
}

// HasAuthorizedValidatorSigner is a free data retrieval call binding the contract method 0x0127dbed.
//
// Solidity: function hasAuthorizedValidatorSigner(address account) view returns(bool)
func (_Accounts *AccountsCallerSession) HasAuthorizedValidatorSigner(account common.Address) (bool, error) {
	return _Accounts.Contract.HasAuthorizedValidatorSigner(&_Accounts.CallOpts, account)
}

// HasAuthorizedVoteSigner is a free data retrieval call binding the contract method 0x614ed493.
//
// Solidity: function hasAuthorizedVoteSigner(address account) view returns(bool)
func (_Accounts *AccountsCaller) HasAuthorizedVoteSigner(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "hasAuthorizedVoteSigner", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAuthorizedVoteSigner is a free data retrieval call binding the contract method 0x614ed493.
//
// Solidity: function hasAuthorizedVoteSigner(address account) view returns(bool)
func (_Accounts *AccountsSession) HasAuthorizedVoteSigner(account common.Address) (bool, error) {
	return _Accounts.Contract.HasAuthorizedVoteSigner(&_Accounts.CallOpts, account)
}

// HasAuthorizedVoteSigner is a free data retrieval call binding the contract method 0x614ed493.
//
// Solidity: function hasAuthorizedVoteSigner(address account) view returns(bool)
func (_Accounts *AccountsCallerSession) HasAuthorizedVoteSigner(account common.Address) (bool, error) {
	return _Accounts.Contract.HasAuthorizedVoteSigner(&_Accounts.CallOpts, account)
}

// HasDefaultSigner is a free data retrieval call binding the contract method 0x05be6229.
//
// Solidity: function hasDefaultSigner(address account, bytes32 role) view returns(bool)
func (_Accounts *AccountsCaller) HasDefaultSigner(opts *bind.CallOpts, account common.Address, role [32]byte) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "hasDefaultSigner", account, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDefaultSigner is a free data retrieval call binding the contract method 0x05be6229.
//
// Solidity: function hasDefaultSigner(address account, bytes32 role) view returns(bool)
func (_Accounts *AccountsSession) HasDefaultSigner(account common.Address, role [32]byte) (bool, error) {
	return _Accounts.Contract.HasDefaultSigner(&_Accounts.CallOpts, account, role)
}

// HasDefaultSigner is a free data retrieval call binding the contract method 0x05be6229.
//
// Solidity: function hasDefaultSigner(address account, bytes32 role) view returns(bool)
func (_Accounts *AccountsCallerSession) HasDefaultSigner(account common.Address, role [32]byte) (bool, error) {
	return _Accounts.Contract.HasDefaultSigner(&_Accounts.CallOpts, account, role)
}

// HasIndexedSigner is a free data retrieval call binding the contract method 0xff836d93.
//
// Solidity: function hasIndexedSigner(address account, bytes32 role) view returns(bool)
func (_Accounts *AccountsCaller) HasIndexedSigner(opts *bind.CallOpts, account common.Address, role [32]byte) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "hasIndexedSigner", account, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasIndexedSigner is a free data retrieval call binding the contract method 0xff836d93.
//
// Solidity: function hasIndexedSigner(address account, bytes32 role) view returns(bool)
func (_Accounts *AccountsSession) HasIndexedSigner(account common.Address, role [32]byte) (bool, error) {
	return _Accounts.Contract.HasIndexedSigner(&_Accounts.CallOpts, account, role)
}

// HasIndexedSigner is a free data retrieval call binding the contract method 0xff836d93.
//
// Solidity: function hasIndexedSigner(address account, bytes32 role) view returns(bool)
func (_Accounts *AccountsCallerSession) HasIndexedSigner(account common.Address, role [32]byte) (bool, error) {
	return _Accounts.Contract.HasIndexedSigner(&_Accounts.CallOpts, account, role)
}

// HasLegacySigner is a free data retrieval call binding the contract method 0x1441ece7.
//
// Solidity: function hasLegacySigner(address account, bytes32 role) view returns(bool)
func (_Accounts *AccountsCaller) HasLegacySigner(opts *bind.CallOpts, account common.Address, role [32]byte) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "hasLegacySigner", account, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasLegacySigner is a free data retrieval call binding the contract method 0x1441ece7.
//
// Solidity: function hasLegacySigner(address account, bytes32 role) view returns(bool)
func (_Accounts *AccountsSession) HasLegacySigner(account common.Address, role [32]byte) (bool, error) {
	return _Accounts.Contract.HasLegacySigner(&_Accounts.CallOpts, account, role)
}

// HasLegacySigner is a free data retrieval call binding the contract method 0x1441ece7.
//
// Solidity: function hasLegacySigner(address account, bytes32 role) view returns(bool)
func (_Accounts *AccountsCallerSession) HasLegacySigner(account common.Address, role [32]byte) (bool, error) {
	return _Accounts.Contract.HasLegacySigner(&_Accounts.CallOpts, account, role)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Accounts *AccountsCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Accounts *AccountsSession) Initialized() (bool, error) {
	return _Accounts.Contract.Initialized(&_Accounts.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Accounts *AccountsCallerSession) Initialized() (bool, error) {
	return _Accounts.Contract.Initialized(&_Accounts.CallOpts)
}

// IsAccount is a free data retrieval call binding the contract method 0x25ca4c9c.
//
// Solidity: function isAccount(address account) view returns(bool)
func (_Accounts *AccountsCaller) IsAccount(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "isAccount", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAccount is a free data retrieval call binding the contract method 0x25ca4c9c.
//
// Solidity: function isAccount(address account) view returns(bool)
func (_Accounts *AccountsSession) IsAccount(account common.Address) (bool, error) {
	return _Accounts.Contract.IsAccount(&_Accounts.CallOpts, account)
}

// IsAccount is a free data retrieval call binding the contract method 0x25ca4c9c.
//
// Solidity: function isAccount(address account) view returns(bool)
func (_Accounts *AccountsCallerSession) IsAccount(account common.Address) (bool, error) {
	return _Accounts.Contract.IsAccount(&_Accounts.CallOpts, account)
}

// IsAuthorizedSigner is a free data retrieval call binding the contract method 0x49045e16.
//
// Solidity: function isAuthorizedSigner(address signer) view returns(bool)
func (_Accounts *AccountsCaller) IsAuthorizedSigner(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "isAuthorizedSigner", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorizedSigner is a free data retrieval call binding the contract method 0x49045e16.
//
// Solidity: function isAuthorizedSigner(address signer) view returns(bool)
func (_Accounts *AccountsSession) IsAuthorizedSigner(signer common.Address) (bool, error) {
	return _Accounts.Contract.IsAuthorizedSigner(&_Accounts.CallOpts, signer)
}

// IsAuthorizedSigner is a free data retrieval call binding the contract method 0x49045e16.
//
// Solidity: function isAuthorizedSigner(address signer) view returns(bool)
func (_Accounts *AccountsCallerSession) IsAuthorizedSigner(signer common.Address) (bool, error) {
	return _Accounts.Contract.IsAuthorizedSigner(&_Accounts.CallOpts, signer)
}

// IsDefaultSigner is a free data retrieval call binding the contract method 0x0b8e0562.
//
// Solidity: function isDefaultSigner(address account, address signer, bytes32 role) view returns(bool)
func (_Accounts *AccountsCaller) IsDefaultSigner(opts *bind.CallOpts, account common.Address, signer common.Address, role [32]byte) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "isDefaultSigner", account, signer, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDefaultSigner is a free data retrieval call binding the contract method 0x0b8e0562.
//
// Solidity: function isDefaultSigner(address account, address signer, bytes32 role) view returns(bool)
func (_Accounts *AccountsSession) IsDefaultSigner(account common.Address, signer common.Address, role [32]byte) (bool, error) {
	return _Accounts.Contract.IsDefaultSigner(&_Accounts.CallOpts, account, signer, role)
}

// IsDefaultSigner is a free data retrieval call binding the contract method 0x0b8e0562.
//
// Solidity: function isDefaultSigner(address account, address signer, bytes32 role) view returns(bool)
func (_Accounts *AccountsCallerSession) IsDefaultSigner(account common.Address, signer common.Address, role [32]byte) (bool, error) {
	return _Accounts.Contract.IsDefaultSigner(&_Accounts.CallOpts, account, signer, role)
}

// IsIndexedSigner is a free data retrieval call binding the contract method 0xc48830d9.
//
// Solidity: function isIndexedSigner(address account, address signer, bytes32 role) view returns(bool)
func (_Accounts *AccountsCaller) IsIndexedSigner(opts *bind.CallOpts, account common.Address, signer common.Address, role [32]byte) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "isIndexedSigner", account, signer, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsIndexedSigner is a free data retrieval call binding the contract method 0xc48830d9.
//
// Solidity: function isIndexedSigner(address account, address signer, bytes32 role) view returns(bool)
func (_Accounts *AccountsSession) IsIndexedSigner(account common.Address, signer common.Address, role [32]byte) (bool, error) {
	return _Accounts.Contract.IsIndexedSigner(&_Accounts.CallOpts, account, signer, role)
}

// IsIndexedSigner is a free data retrieval call binding the contract method 0xc48830d9.
//
// Solidity: function isIndexedSigner(address account, address signer, bytes32 role) view returns(bool)
func (_Accounts *AccountsCallerSession) IsIndexedSigner(account common.Address, signer common.Address, role [32]byte) (bool, error) {
	return _Accounts.Contract.IsIndexedSigner(&_Accounts.CallOpts, account, signer, role)
}

// IsLegacyRole is a free data retrieval call binding the contract method 0xf333d836.
//
// Solidity: function isLegacyRole(bytes32 role) pure returns(bool)
func (_Accounts *AccountsCaller) IsLegacyRole(opts *bind.CallOpts, role [32]byte) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "isLegacyRole", role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLegacyRole is a free data retrieval call binding the contract method 0xf333d836.
//
// Solidity: function isLegacyRole(bytes32 role) pure returns(bool)
func (_Accounts *AccountsSession) IsLegacyRole(role [32]byte) (bool, error) {
	return _Accounts.Contract.IsLegacyRole(&_Accounts.CallOpts, role)
}

// IsLegacyRole is a free data retrieval call binding the contract method 0xf333d836.
//
// Solidity: function isLegacyRole(bytes32 role) pure returns(bool)
func (_Accounts *AccountsCallerSession) IsLegacyRole(role [32]byte) (bool, error) {
	return _Accounts.Contract.IsLegacyRole(&_Accounts.CallOpts, role)
}

// IsLegacySigner is a free data retrieval call binding the contract method 0x91cd074b.
//
// Solidity: function isLegacySigner(address _account, address signer, bytes32 role) view returns(bool)
func (_Accounts *AccountsCaller) IsLegacySigner(opts *bind.CallOpts, _account common.Address, signer common.Address, role [32]byte) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "isLegacySigner", _account, signer, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLegacySigner is a free data retrieval call binding the contract method 0x91cd074b.
//
// Solidity: function isLegacySigner(address _account, address signer, bytes32 role) view returns(bool)
func (_Accounts *AccountsSession) IsLegacySigner(_account common.Address, signer common.Address, role [32]byte) (bool, error) {
	return _Accounts.Contract.IsLegacySigner(&_Accounts.CallOpts, _account, signer, role)
}

// IsLegacySigner is a free data retrieval call binding the contract method 0x91cd074b.
//
// Solidity: function isLegacySigner(address _account, address signer, bytes32 role) view returns(bool)
func (_Accounts *AccountsCallerSession) IsLegacySigner(_account common.Address, signer common.Address, role [32]byte) (bool, error) {
	return _Accounts.Contract.IsLegacySigner(&_Accounts.CallOpts, _account, signer, role)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Accounts *AccountsCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Accounts *AccountsSession) IsOwner() (bool, error) {
	return _Accounts.Contract.IsOwner(&_Accounts.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Accounts *AccountsCallerSession) IsOwner() (bool, error) {
	return _Accounts.Contract.IsOwner(&_Accounts.CallOpts)
}

// IsSigner is a free data retrieval call binding the contract method 0x485a46d1.
//
// Solidity: function isSigner(address account, address signer, bytes32 role) view returns(bool)
func (_Accounts *AccountsCaller) IsSigner(opts *bind.CallOpts, account common.Address, signer common.Address, role [32]byte) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "isSigner", account, signer, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x485a46d1.
//
// Solidity: function isSigner(address account, address signer, bytes32 role) view returns(bool)
func (_Accounts *AccountsSession) IsSigner(account common.Address, signer common.Address, role [32]byte) (bool, error) {
	return _Accounts.Contract.IsSigner(&_Accounts.CallOpts, account, signer, role)
}

// IsSigner is a free data retrieval call binding the contract method 0x485a46d1.
//
// Solidity: function isSigner(address account, address signer, bytes32 role) view returns(bool)
func (_Accounts *AccountsCallerSession) IsSigner(account common.Address, signer common.Address, role [32]byte) (bool, error) {
	return _Accounts.Contract.IsSigner(&_Accounts.CallOpts, account, signer, role)
}

// OffchainStorageRoots is a free data retrieval call binding the contract method 0x76082c1f.
//
// Solidity: function offchainStorageRoots(address , uint256 ) view returns(bytes)
func (_Accounts *AccountsCaller) OffchainStorageRoots(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "offchainStorageRoots", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// OffchainStorageRoots is a free data retrieval call binding the contract method 0x76082c1f.
//
// Solidity: function offchainStorageRoots(address , uint256 ) view returns(bytes)
func (_Accounts *AccountsSession) OffchainStorageRoots(arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	return _Accounts.Contract.OffchainStorageRoots(&_Accounts.CallOpts, arg0, arg1)
}

// OffchainStorageRoots is a free data retrieval call binding the contract method 0x76082c1f.
//
// Solidity: function offchainStorageRoots(address , uint256 ) view returns(bytes)
func (_Accounts *AccountsCallerSession) OffchainStorageRoots(arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	return _Accounts.Contract.OffchainStorageRoots(&_Accounts.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Accounts *AccountsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Accounts *AccountsSession) Owner() (common.Address, error) {
	return _Accounts.Contract.Owner(&_Accounts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Accounts *AccountsCallerSession) Owner() (common.Address, error) {
	return _Accounts.Contract.Owner(&_Accounts.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Accounts *AccountsCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Accounts *AccountsSession) Registry() (common.Address, error) {
	return _Accounts.Contract.Registry(&_Accounts.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Accounts *AccountsCallerSession) Registry() (common.Address, error) {
	return _Accounts.Contract.Registry(&_Accounts.CallOpts)
}

// SignerToAccount is a free data retrieval call binding the contract method 0x93c5c487.
//
// Solidity: function signerToAccount(address signer) view returns(address)
func (_Accounts *AccountsCaller) SignerToAccount(opts *bind.CallOpts, signer common.Address) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "signerToAccount", signer)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerToAccount is a free data retrieval call binding the contract method 0x93c5c487.
//
// Solidity: function signerToAccount(address signer) view returns(address)
func (_Accounts *AccountsSession) SignerToAccount(signer common.Address) (common.Address, error) {
	return _Accounts.Contract.SignerToAccount(&_Accounts.CallOpts, signer)
}

// SignerToAccount is a free data retrieval call binding the contract method 0x93c5c487.
//
// Solidity: function signerToAccount(address signer) view returns(address)
func (_Accounts *AccountsCallerSession) SignerToAccount(signer common.Address) (common.Address, error) {
	return _Accounts.Contract.SignerToAccount(&_Accounts.CallOpts, signer)
}

// ValidatorSignerToAccount is a free data retrieval call binding the contract method 0x64439b43.
//
// Solidity: function validatorSignerToAccount(address signer) view returns(address)
func (_Accounts *AccountsCaller) ValidatorSignerToAccount(opts *bind.CallOpts, signer common.Address) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "validatorSignerToAccount", signer)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerToAccount is a free data retrieval call binding the contract method 0x64439b43.
//
// Solidity: function validatorSignerToAccount(address signer) view returns(address)
func (_Accounts *AccountsSession) ValidatorSignerToAccount(signer common.Address) (common.Address, error) {
	return _Accounts.Contract.ValidatorSignerToAccount(&_Accounts.CallOpts, signer)
}

// ValidatorSignerToAccount is a free data retrieval call binding the contract method 0x64439b43.
//
// Solidity: function validatorSignerToAccount(address signer) view returns(address)
func (_Accounts *AccountsCallerSession) ValidatorSignerToAccount(signer common.Address) (common.Address, error) {
	return _Accounts.Contract.ValidatorSignerToAccount(&_Accounts.CallOpts, signer)
}

// VoteSignerToAccount is a free data retrieval call binding the contract method 0x6642d594.
//
// Solidity: function voteSignerToAccount(address signer) view returns(address)
func (_Accounts *AccountsCaller) VoteSignerToAccount(opts *bind.CallOpts, signer common.Address) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "voteSignerToAccount", signer)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteSignerToAccount is a free data retrieval call binding the contract method 0x6642d594.
//
// Solidity: function voteSignerToAccount(address signer) view returns(address)
func (_Accounts *AccountsSession) VoteSignerToAccount(signer common.Address) (common.Address, error) {
	return _Accounts.Contract.VoteSignerToAccount(&_Accounts.CallOpts, signer)
}

// VoteSignerToAccount is a free data retrieval call binding the contract method 0x6642d594.
//
// Solidity: function voteSignerToAccount(address signer) view returns(address)
func (_Accounts *AccountsCallerSession) VoteSignerToAccount(signer common.Address) (common.Address, error) {
	return _Accounts.Contract.VoteSignerToAccount(&_Accounts.CallOpts, signer)
}

// AddStorageRoot is a paid mutator transaction binding the contract method 0xb062c843.
//
// Solidity: function addStorageRoot(bytes url) returns()
func (_Accounts *AccountsTransactor) AddStorageRoot(opts *bind.TransactOpts, url []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "addStorageRoot", url)
}

// AddStorageRoot is a paid mutator transaction binding the contract method 0xb062c843.
//
// Solidity: function addStorageRoot(bytes url) returns()
func (_Accounts *AccountsSession) AddStorageRoot(url []byte) (*types.Transaction, error) {
	return _Accounts.Contract.AddStorageRoot(&_Accounts.TransactOpts, url)
}

// AddStorageRoot is a paid mutator transaction binding the contract method 0xb062c843.
//
// Solidity: function addStorageRoot(bytes url) returns()
func (_Accounts *AccountsTransactorSession) AddStorageRoot(url []byte) (*types.Transaction, error) {
	return _Accounts.Contract.AddStorageRoot(&_Accounts.TransactOpts, url)
}

// AuthorizeAttestationSigner is a paid mutator transaction binding the contract method 0x76afa04c.
//
// Solidity: function authorizeAttestationSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactor) AuthorizeAttestationSigner(opts *bind.TransactOpts, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "authorizeAttestationSigner", signer, v, r, s)
}

// AuthorizeAttestationSigner is a paid mutator transaction binding the contract method 0x76afa04c.
//
// Solidity: function authorizeAttestationSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsSession) AuthorizeAttestationSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeAttestationSigner(&_Accounts.TransactOpts, signer, v, r, s)
}

// AuthorizeAttestationSigner is a paid mutator transaction binding the contract method 0x76afa04c.
//
// Solidity: function authorizeAttestationSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactorSession) AuthorizeAttestationSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeAttestationSigner(&_Accounts.TransactOpts, signer, v, r, s)
}

// AuthorizeSigner is a paid mutator transaction binding the contract method 0x58b81ea8.
//
// Solidity: function authorizeSigner(address signer, bytes32 role) returns()
func (_Accounts *AccountsTransactor) AuthorizeSigner(opts *bind.TransactOpts, signer common.Address, role [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "authorizeSigner", signer, role)
}

// AuthorizeSigner is a paid mutator transaction binding the contract method 0x58b81ea8.
//
// Solidity: function authorizeSigner(address signer, bytes32 role) returns()
func (_Accounts *AccountsSession) AuthorizeSigner(signer common.Address, role [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeSigner(&_Accounts.TransactOpts, signer, role)
}

// AuthorizeSigner is a paid mutator transaction binding the contract method 0x58b81ea8.
//
// Solidity: function authorizeSigner(address signer, bytes32 role) returns()
func (_Accounts *AccountsTransactorSession) AuthorizeSigner(signer common.Address, role [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeSigner(&_Accounts.TransactOpts, signer, role)
}

// AuthorizeSignerWithSignature is a paid mutator transaction binding the contract method 0x92f90fbf.
//
// Solidity: function authorizeSignerWithSignature(address signer, bytes32 role, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactor) AuthorizeSignerWithSignature(opts *bind.TransactOpts, signer common.Address, role [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "authorizeSignerWithSignature", signer, role, v, r, s)
}

// AuthorizeSignerWithSignature is a paid mutator transaction binding the contract method 0x92f90fbf.
//
// Solidity: function authorizeSignerWithSignature(address signer, bytes32 role, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsSession) AuthorizeSignerWithSignature(signer common.Address, role [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeSignerWithSignature(&_Accounts.TransactOpts, signer, role, v, r, s)
}

// AuthorizeSignerWithSignature is a paid mutator transaction binding the contract method 0x92f90fbf.
//
// Solidity: function authorizeSignerWithSignature(address signer, bytes32 role, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactorSession) AuthorizeSignerWithSignature(signer common.Address, role [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeSignerWithSignature(&_Accounts.TransactOpts, signer, role, v, r, s)
}

// AuthorizeValidatorSigner is a paid mutator transaction binding the contract method 0xbaf7ef0f.
//
// Solidity: function authorizeValidatorSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactor) AuthorizeValidatorSigner(opts *bind.TransactOpts, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "authorizeValidatorSigner", signer, v, r, s)
}

// AuthorizeValidatorSigner is a paid mutator transaction binding the contract method 0xbaf7ef0f.
//
// Solidity: function authorizeValidatorSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsSession) AuthorizeValidatorSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeValidatorSigner(&_Accounts.TransactOpts, signer, v, r, s)
}

// AuthorizeValidatorSigner is a paid mutator transaction binding the contract method 0xbaf7ef0f.
//
// Solidity: function authorizeValidatorSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactorSession) AuthorizeValidatorSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeValidatorSigner(&_Accounts.TransactOpts, signer, v, r, s)
}

// AuthorizeValidatorSignerWithKeys is a paid mutator transaction binding the contract method 0x1465b923.
//
// Solidity: function authorizeValidatorSignerWithKeys(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns()
func (_Accounts *AccountsTransactor) AuthorizeValidatorSignerWithKeys(opts *bind.TransactOpts, signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "authorizeValidatorSignerWithKeys", signer, v, r, s, ecdsaPublicKey, blsPublicKey, blsPop)
}

// AuthorizeValidatorSignerWithKeys is a paid mutator transaction binding the contract method 0x1465b923.
//
// Solidity: function authorizeValidatorSignerWithKeys(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns()
func (_Accounts *AccountsSession) AuthorizeValidatorSignerWithKeys(signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeValidatorSignerWithKeys(&_Accounts.TransactOpts, signer, v, r, s, ecdsaPublicKey, blsPublicKey, blsPop)
}

// AuthorizeValidatorSignerWithKeys is a paid mutator transaction binding the contract method 0x1465b923.
//
// Solidity: function authorizeValidatorSignerWithKeys(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns()
func (_Accounts *AccountsTransactorSession) AuthorizeValidatorSignerWithKeys(signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeValidatorSignerWithKeys(&_Accounts.TransactOpts, signer, v, r, s, ecdsaPublicKey, blsPublicKey, blsPop)
}

// AuthorizeValidatorSignerWithPublicKey is a paid mutator transaction binding the contract method 0x0fa750d2.
//
// Solidity: function authorizeValidatorSignerWithPublicKey(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey) returns()
func (_Accounts *AccountsTransactor) AuthorizeValidatorSignerWithPublicKey(opts *bind.TransactOpts, signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "authorizeValidatorSignerWithPublicKey", signer, v, r, s, ecdsaPublicKey)
}

// AuthorizeValidatorSignerWithPublicKey is a paid mutator transaction binding the contract method 0x0fa750d2.
//
// Solidity: function authorizeValidatorSignerWithPublicKey(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey) returns()
func (_Accounts *AccountsSession) AuthorizeValidatorSignerWithPublicKey(signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeValidatorSignerWithPublicKey(&_Accounts.TransactOpts, signer, v, r, s, ecdsaPublicKey)
}

// AuthorizeValidatorSignerWithPublicKey is a paid mutator transaction binding the contract method 0x0fa750d2.
//
// Solidity: function authorizeValidatorSignerWithPublicKey(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey) returns()
func (_Accounts *AccountsTransactorSession) AuthorizeValidatorSignerWithPublicKey(signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeValidatorSignerWithPublicKey(&_Accounts.TransactOpts, signer, v, r, s, ecdsaPublicKey)
}

// AuthorizeVoteSigner is a paid mutator transaction binding the contract method 0x4282ee6d.
//
// Solidity: function authorizeVoteSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactor) AuthorizeVoteSigner(opts *bind.TransactOpts, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "authorizeVoteSigner", signer, v, r, s)
}

// AuthorizeVoteSigner is a paid mutator transaction binding the contract method 0x4282ee6d.
//
// Solidity: function authorizeVoteSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsSession) AuthorizeVoteSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeVoteSigner(&_Accounts.TransactOpts, signer, v, r, s)
}

// AuthorizeVoteSigner is a paid mutator transaction binding the contract method 0x4282ee6d.
//
// Solidity: function authorizeVoteSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactorSession) AuthorizeVoteSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeVoteSigner(&_Accounts.TransactOpts, signer, v, r, s)
}

// CompleteSignerAuthorization is a paid mutator transaction binding the contract method 0x9f682976.
//
// Solidity: function completeSignerAuthorization(address account, bytes32 role) returns()
func (_Accounts *AccountsTransactor) CompleteSignerAuthorization(opts *bind.TransactOpts, account common.Address, role [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "completeSignerAuthorization", account, role)
}

// CompleteSignerAuthorization is a paid mutator transaction binding the contract method 0x9f682976.
//
// Solidity: function completeSignerAuthorization(address account, bytes32 role) returns()
func (_Accounts *AccountsSession) CompleteSignerAuthorization(account common.Address, role [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.CompleteSignerAuthorization(&_Accounts.TransactOpts, account, role)
}

// CompleteSignerAuthorization is a paid mutator transaction binding the contract method 0x9f682976.
//
// Solidity: function completeSignerAuthorization(address account, bytes32 role) returns()
func (_Accounts *AccountsTransactorSession) CompleteSignerAuthorization(account common.Address, role [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.CompleteSignerAuthorization(&_Accounts.TransactOpts, account, role)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bool)
func (_Accounts *AccountsTransactor) CreateAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "createAccount")
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bool)
func (_Accounts *AccountsSession) CreateAccount() (*types.Transaction, error) {
	return _Accounts.Contract.CreateAccount(&_Accounts.TransactOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bool)
func (_Accounts *AccountsTransactorSession) CreateAccount() (*types.Transaction, error) {
	return _Accounts.Contract.CreateAccount(&_Accounts.TransactOpts)
}

// DeletePaymentDelegation is a paid mutator transaction binding the contract method 0xbce2b8d6.
//
// Solidity: function deletePaymentDelegation() returns()
func (_Accounts *AccountsTransactor) DeletePaymentDelegation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "deletePaymentDelegation")
}

// DeletePaymentDelegation is a paid mutator transaction binding the contract method 0xbce2b8d6.
//
// Solidity: function deletePaymentDelegation() returns()
func (_Accounts *AccountsSession) DeletePaymentDelegation() (*types.Transaction, error) {
	return _Accounts.Contract.DeletePaymentDelegation(&_Accounts.TransactOpts)
}

// DeletePaymentDelegation is a paid mutator transaction binding the contract method 0xbce2b8d6.
//
// Solidity: function deletePaymentDelegation() returns()
func (_Accounts *AccountsTransactorSession) DeletePaymentDelegation() (*types.Transaction, error) {
	return _Accounts.Contract.DeletePaymentDelegation(&_Accounts.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registryAddress) returns()
func (_Accounts *AccountsTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "initialize", registryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registryAddress) returns()
func (_Accounts *AccountsSession) Initialize(registryAddress common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.Initialize(&_Accounts.TransactOpts, registryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registryAddress) returns()
func (_Accounts *AccountsTransactorSession) Initialize(registryAddress common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.Initialize(&_Accounts.TransactOpts, registryAddress)
}

// RemoveAttestationSigner is a paid mutator transaction binding the contract method 0x760fbbb2.
//
// Solidity: function removeAttestationSigner() returns()
func (_Accounts *AccountsTransactor) RemoveAttestationSigner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "removeAttestationSigner")
}

// RemoveAttestationSigner is a paid mutator transaction binding the contract method 0x760fbbb2.
//
// Solidity: function removeAttestationSigner() returns()
func (_Accounts *AccountsSession) RemoveAttestationSigner() (*types.Transaction, error) {
	return _Accounts.Contract.RemoveAttestationSigner(&_Accounts.TransactOpts)
}

// RemoveAttestationSigner is a paid mutator transaction binding the contract method 0x760fbbb2.
//
// Solidity: function removeAttestationSigner() returns()
func (_Accounts *AccountsTransactorSession) RemoveAttestationSigner() (*types.Transaction, error) {
	return _Accounts.Contract.RemoveAttestationSigner(&_Accounts.TransactOpts)
}

// RemoveDefaultSigner is a paid mutator transaction binding the contract method 0xe7a16e6b.
//
// Solidity: function removeDefaultSigner(bytes32 role) returns()
func (_Accounts *AccountsTransactor) RemoveDefaultSigner(opts *bind.TransactOpts, role [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "removeDefaultSigner", role)
}

// RemoveDefaultSigner is a paid mutator transaction binding the contract method 0xe7a16e6b.
//
// Solidity: function removeDefaultSigner(bytes32 role) returns()
func (_Accounts *AccountsSession) RemoveDefaultSigner(role [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.RemoveDefaultSigner(&_Accounts.TransactOpts, role)
}

// RemoveDefaultSigner is a paid mutator transaction binding the contract method 0xe7a16e6b.
//
// Solidity: function removeDefaultSigner(bytes32 role) returns()
func (_Accounts *AccountsTransactorSession) RemoveDefaultSigner(role [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.RemoveDefaultSigner(&_Accounts.TransactOpts, role)
}

// RemoveIndexedSigner is a paid mutator transaction binding the contract method 0x0185a232.
//
// Solidity: function removeIndexedSigner(bytes32 role) returns()
func (_Accounts *AccountsTransactor) RemoveIndexedSigner(opts *bind.TransactOpts, role [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "removeIndexedSigner", role)
}

// RemoveIndexedSigner is a paid mutator transaction binding the contract method 0x0185a232.
//
// Solidity: function removeIndexedSigner(bytes32 role) returns()
func (_Accounts *AccountsSession) RemoveIndexedSigner(role [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.RemoveIndexedSigner(&_Accounts.TransactOpts, role)
}

// RemoveIndexedSigner is a paid mutator transaction binding the contract method 0x0185a232.
//
// Solidity: function removeIndexedSigner(bytes32 role) returns()
func (_Accounts *AccountsTransactorSession) RemoveIndexedSigner(role [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.RemoveIndexedSigner(&_Accounts.TransactOpts, role)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0xfbe3c373.
//
// Solidity: function removeSigner(address signer, bytes32 role) returns()
func (_Accounts *AccountsTransactor) RemoveSigner(opts *bind.TransactOpts, signer common.Address, role [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "removeSigner", signer, role)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0xfbe3c373.
//
// Solidity: function removeSigner(address signer, bytes32 role) returns()
func (_Accounts *AccountsSession) RemoveSigner(signer common.Address, role [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.RemoveSigner(&_Accounts.TransactOpts, signer, role)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0xfbe3c373.
//
// Solidity: function removeSigner(address signer, bytes32 role) returns()
func (_Accounts *AccountsTransactorSession) RemoveSigner(signer common.Address, role [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.RemoveSigner(&_Accounts.TransactOpts, signer, role)
}

// RemoveStorageRoot is a paid mutator transaction binding the contract method 0xe8d787cb.
//
// Solidity: function removeStorageRoot(uint256 index) returns()
func (_Accounts *AccountsTransactor) RemoveStorageRoot(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "removeStorageRoot", index)
}

// RemoveStorageRoot is a paid mutator transaction binding the contract method 0xe8d787cb.
//
// Solidity: function removeStorageRoot(uint256 index) returns()
func (_Accounts *AccountsSession) RemoveStorageRoot(index *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.RemoveStorageRoot(&_Accounts.TransactOpts, index)
}

// RemoveStorageRoot is a paid mutator transaction binding the contract method 0xe8d787cb.
//
// Solidity: function removeStorageRoot(uint256 index) returns()
func (_Accounts *AccountsTransactorSession) RemoveStorageRoot(index *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.RemoveStorageRoot(&_Accounts.TransactOpts, index)
}

// RemoveValidatorSigner is a paid mutator transaction binding the contract method 0xa5ec94f9.
//
// Solidity: function removeValidatorSigner() returns()
func (_Accounts *AccountsTransactor) RemoveValidatorSigner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "removeValidatorSigner")
}

// RemoveValidatorSigner is a paid mutator transaction binding the contract method 0xa5ec94f9.
//
// Solidity: function removeValidatorSigner() returns()
func (_Accounts *AccountsSession) RemoveValidatorSigner() (*types.Transaction, error) {
	return _Accounts.Contract.RemoveValidatorSigner(&_Accounts.TransactOpts)
}

// RemoveValidatorSigner is a paid mutator transaction binding the contract method 0xa5ec94f9.
//
// Solidity: function removeValidatorSigner() returns()
func (_Accounts *AccountsTransactorSession) RemoveValidatorSigner() (*types.Transaction, error) {
	return _Accounts.Contract.RemoveValidatorSigner(&_Accounts.TransactOpts)
}

// RemoveVoteSigner is a paid mutator transaction binding the contract method 0x10c504b5.
//
// Solidity: function removeVoteSigner() returns()
func (_Accounts *AccountsTransactor) RemoveVoteSigner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "removeVoteSigner")
}

// RemoveVoteSigner is a paid mutator transaction binding the contract method 0x10c504b5.
//
// Solidity: function removeVoteSigner() returns()
func (_Accounts *AccountsSession) RemoveVoteSigner() (*types.Transaction, error) {
	return _Accounts.Contract.RemoveVoteSigner(&_Accounts.TransactOpts)
}

// RemoveVoteSigner is a paid mutator transaction binding the contract method 0x10c504b5.
//
// Solidity: function removeVoteSigner() returns()
func (_Accounts *AccountsTransactorSession) RemoveVoteSigner() (*types.Transaction, error) {
	return _Accounts.Contract.RemoveVoteSigner(&_Accounts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accounts *AccountsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accounts *AccountsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Accounts.Contract.RenounceOwnership(&_Accounts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accounts *AccountsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Accounts.Contract.RenounceOwnership(&_Accounts.TransactOpts)
}

// SetAccount is a paid mutator transaction binding the contract method 0x90b12b47.
//
// Solidity: function setAccount(string name, bytes dataEncryptionKey, address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactor) SetAccount(opts *bind.TransactOpts, name string, dataEncryptionKey []byte, walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setAccount", name, dataEncryptionKey, walletAddress, v, r, s)
}

// SetAccount is a paid mutator transaction binding the contract method 0x90b12b47.
//
// Solidity: function setAccount(string name, bytes dataEncryptionKey, address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsSession) SetAccount(name string, dataEncryptionKey []byte, walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.SetAccount(&_Accounts.TransactOpts, name, dataEncryptionKey, walletAddress, v, r, s)
}

// SetAccount is a paid mutator transaction binding the contract method 0x90b12b47.
//
// Solidity: function setAccount(string name, bytes dataEncryptionKey, address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactorSession) SetAccount(name string, dataEncryptionKey []byte, walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.SetAccount(&_Accounts.TransactOpts, name, dataEncryptionKey, walletAddress, v, r, s)
}

// SetAccountDataEncryptionKey is a paid mutator transaction binding the contract method 0x0fe7abab.
//
// Solidity: function setAccountDataEncryptionKey(bytes dataEncryptionKey) returns()
func (_Accounts *AccountsTransactor) SetAccountDataEncryptionKey(opts *bind.TransactOpts, dataEncryptionKey []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setAccountDataEncryptionKey", dataEncryptionKey)
}

// SetAccountDataEncryptionKey is a paid mutator transaction binding the contract method 0x0fe7abab.
//
// Solidity: function setAccountDataEncryptionKey(bytes dataEncryptionKey) returns()
func (_Accounts *AccountsSession) SetAccountDataEncryptionKey(dataEncryptionKey []byte) (*types.Transaction, error) {
	return _Accounts.Contract.SetAccountDataEncryptionKey(&_Accounts.TransactOpts, dataEncryptionKey)
}

// SetAccountDataEncryptionKey is a paid mutator transaction binding the contract method 0x0fe7abab.
//
// Solidity: function setAccountDataEncryptionKey(bytes dataEncryptionKey) returns()
func (_Accounts *AccountsTransactorSession) SetAccountDataEncryptionKey(dataEncryptionKey []byte) (*types.Transaction, error) {
	return _Accounts.Contract.SetAccountDataEncryptionKey(&_Accounts.TransactOpts, dataEncryptionKey)
}

// SetEip712DomainSeparator is a paid mutator transaction binding the contract method 0x3184b3c5.
//
// Solidity: function setEip712DomainSeparator() returns()
func (_Accounts *AccountsTransactor) SetEip712DomainSeparator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setEip712DomainSeparator")
}

// SetEip712DomainSeparator is a paid mutator transaction binding the contract method 0x3184b3c5.
//
// Solidity: function setEip712DomainSeparator() returns()
func (_Accounts *AccountsSession) SetEip712DomainSeparator() (*types.Transaction, error) {
	return _Accounts.Contract.SetEip712DomainSeparator(&_Accounts.TransactOpts)
}

// SetEip712DomainSeparator is a paid mutator transaction binding the contract method 0x3184b3c5.
//
// Solidity: function setEip712DomainSeparator() returns()
func (_Accounts *AccountsTransactorSession) SetEip712DomainSeparator() (*types.Transaction, error) {
	return _Accounts.Contract.SetEip712DomainSeparator(&_Accounts.TransactOpts)
}

// SetIndexedSigner is a paid mutator transaction binding the contract method 0x727d079c.
//
// Solidity: function setIndexedSigner(address signer, bytes32 role) returns()
func (_Accounts *AccountsTransactor) SetIndexedSigner(opts *bind.TransactOpts, signer common.Address, role [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setIndexedSigner", signer, role)
}

// SetIndexedSigner is a paid mutator transaction binding the contract method 0x727d079c.
//
// Solidity: function setIndexedSigner(address signer, bytes32 role) returns()
func (_Accounts *AccountsSession) SetIndexedSigner(signer common.Address, role [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.SetIndexedSigner(&_Accounts.TransactOpts, signer, role)
}

// SetIndexedSigner is a paid mutator transaction binding the contract method 0x727d079c.
//
// Solidity: function setIndexedSigner(address signer, bytes32 role) returns()
func (_Accounts *AccountsTransactorSession) SetIndexedSigner(signer common.Address, role [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.SetIndexedSigner(&_Accounts.TransactOpts, signer, role)
}

// SetMetadataURL is a paid mutator transaction binding the contract method 0x747daec5.
//
// Solidity: function setMetadataURL(string metadataURL) returns()
func (_Accounts *AccountsTransactor) SetMetadataURL(opts *bind.TransactOpts, metadataURL string) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setMetadataURL", metadataURL)
}

// SetMetadataURL is a paid mutator transaction binding the contract method 0x747daec5.
//
// Solidity: function setMetadataURL(string metadataURL) returns()
func (_Accounts *AccountsSession) SetMetadataURL(metadataURL string) (*types.Transaction, error) {
	return _Accounts.Contract.SetMetadataURL(&_Accounts.TransactOpts, metadataURL)
}

// SetMetadataURL is a paid mutator transaction binding the contract method 0x747daec5.
//
// Solidity: function setMetadataURL(string metadataURL) returns()
func (_Accounts *AccountsTransactorSession) SetMetadataURL(metadataURL string) (*types.Transaction, error) {
	return _Accounts.Contract.SetMetadataURL(&_Accounts.TransactOpts, metadataURL)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_Accounts *AccountsTransactor) SetName(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setName", name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_Accounts *AccountsSession) SetName(name string) (*types.Transaction, error) {
	return _Accounts.Contract.SetName(&_Accounts.TransactOpts, name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_Accounts *AccountsTransactorSession) SetName(name string) (*types.Transaction, error) {
	return _Accounts.Contract.SetName(&_Accounts.TransactOpts, name)
}

// SetPaymentDelegation is a paid mutator transaction binding the contract method 0x8f9ae6dc.
//
// Solidity: function setPaymentDelegation(address beneficiary, uint256 fraction) returns()
func (_Accounts *AccountsTransactor) SetPaymentDelegation(opts *bind.TransactOpts, beneficiary common.Address, fraction *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setPaymentDelegation", beneficiary, fraction)
}

// SetPaymentDelegation is a paid mutator transaction binding the contract method 0x8f9ae6dc.
//
// Solidity: function setPaymentDelegation(address beneficiary, uint256 fraction) returns()
func (_Accounts *AccountsSession) SetPaymentDelegation(beneficiary common.Address, fraction *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.SetPaymentDelegation(&_Accounts.TransactOpts, beneficiary, fraction)
}

// SetPaymentDelegation is a paid mutator transaction binding the contract method 0x8f9ae6dc.
//
// Solidity: function setPaymentDelegation(address beneficiary, uint256 fraction) returns()
func (_Accounts *AccountsTransactorSession) SetPaymentDelegation(beneficiary common.Address, fraction *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.SetPaymentDelegation(&_Accounts.TransactOpts, beneficiary, fraction)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Accounts *AccountsTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Accounts *AccountsSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.SetRegistry(&_Accounts.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Accounts *AccountsTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.SetRegistry(&_Accounts.TransactOpts, registryAddress)
}

// SetWalletAddress is a paid mutator transaction binding the contract method 0x9cafb2a1.
//
// Solidity: function setWalletAddress(address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactor) SetWalletAddress(opts *bind.TransactOpts, walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setWalletAddress", walletAddress, v, r, s)
}

// SetWalletAddress is a paid mutator transaction binding the contract method 0x9cafb2a1.
//
// Solidity: function setWalletAddress(address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsSession) SetWalletAddress(walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.SetWalletAddress(&_Accounts.TransactOpts, walletAddress, v, r, s)
}

// SetWalletAddress is a paid mutator transaction binding the contract method 0x9cafb2a1.
//
// Solidity: function setWalletAddress(address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactorSession) SetWalletAddress(walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.SetWalletAddress(&_Accounts.TransactOpts, walletAddress, v, r, s)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Accounts *AccountsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Accounts *AccountsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.TransferOwnership(&_Accounts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Accounts *AccountsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.TransferOwnership(&_Accounts.TransactOpts, newOwner)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_Accounts *AccountsFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _Accounts.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "AccountCreated":
		event, err = _Accounts.ParseAccountCreated(log)
	case "AccountDataEncryptionKeySet":
		event, err = _Accounts.ParseAccountDataEncryptionKeySet(log)
	case "AccountMetadataURLSet":
		event, err = _Accounts.ParseAccountMetadataURLSet(log)
	case "AccountNameSet":
		event, err = _Accounts.ParseAccountNameSet(log)
	case "AccountWalletAddressSet":
		event, err = _Accounts.ParseAccountWalletAddressSet(log)
	case "AttestationSignerAuthorized":
		event, err = _Accounts.ParseAttestationSignerAuthorized(log)
	case "AttestationSignerRemoved":
		event, err = _Accounts.ParseAttestationSignerRemoved(log)
	case "DefaultSignerRemoved":
		event, err = _Accounts.ParseDefaultSignerRemoved(log)
	case "DefaultSignerSet":
		event, err = _Accounts.ParseDefaultSignerSet(log)
	case "IndexedSignerRemoved":
		event, err = _Accounts.ParseIndexedSignerRemoved(log)
	case "IndexedSignerSet":
		event, err = _Accounts.ParseIndexedSignerSet(log)
	case "LegacySignerRemoved":
		event, err = _Accounts.ParseLegacySignerRemoved(log)
	case "LegacySignerSet":
		event, err = _Accounts.ParseLegacySignerSet(log)
	case "OffchainStorageRootAdded":
		event, err = _Accounts.ParseOffchainStorageRootAdded(log)
	case "OffchainStorageRootRemoved":
		event, err = _Accounts.ParseOffchainStorageRootRemoved(log)
	case "OwnershipTransferred":
		event, err = _Accounts.ParseOwnershipTransferred(log)
	case "PaymentDelegationSet":
		event, err = _Accounts.ParsePaymentDelegationSet(log)
	case "RegistrySet":
		event, err = _Accounts.ParseRegistrySet(log)
	case "SignerAuthorizationCompleted":
		event, err = _Accounts.ParseSignerAuthorizationCompleted(log)
	case "SignerAuthorizationStarted":
		event, err = _Accounts.ParseSignerAuthorizationStarted(log)
	case "SignerAuthorized":
		event, err = _Accounts.ParseSignerAuthorized(log)
	case "SignerRemoved":
		event, err = _Accounts.ParseSignerRemoved(log)
	case "ValidatorSignerAuthorized":
		event, err = _Accounts.ParseValidatorSignerAuthorized(log)
	case "ValidatorSignerRemoved":
		event, err = _Accounts.ParseValidatorSignerRemoved(log)
	case "VoteSignerAuthorized":
		event, err = _Accounts.ParseVoteSignerAuthorized(log)
	case "VoteSignerRemoved":
		event, err = _Accounts.ParseVoteSignerRemoved(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// AccountsAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the Accounts contract.
type AccountsAccountCreatedIterator struct {
	Event *AccountsAccountCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAccountCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsAccountCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAccountCreated represents a AccountCreated event raised by the Accounts contract.
type AccountsAccountCreated struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0x805996f252884581e2f74cf3d2b03564d5ec26ccc90850ae12653dc1b72d1fa2.
//
// Solidity: event AccountCreated(address indexed account)
func (_Accounts *AccountsFilterer) FilterAccountCreated(opts *bind.FilterOpts, account []common.Address) (*AccountsAccountCreatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AccountCreated", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsAccountCreatedIterator{contract: _Accounts.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0x805996f252884581e2f74cf3d2b03564d5ec26ccc90850ae12653dc1b72d1fa2.
//
// Solidity: event AccountCreated(address indexed account)
func (_Accounts *AccountsFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *AccountsAccountCreated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AccountCreated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAccountCreated)
				if err := _Accounts.contract.UnpackLog(event, "AccountCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountCreated is a log parse operation binding the contract event 0x805996f252884581e2f74cf3d2b03564d5ec26ccc90850ae12653dc1b72d1fa2.
//
// Solidity: event AccountCreated(address indexed account)
func (_Accounts *AccountsFilterer) ParseAccountCreated(log types.Log) (*AccountsAccountCreated, error) {
	event := new(AccountsAccountCreated)
	if err := _Accounts.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsAccountDataEncryptionKeySetIterator is returned from FilterAccountDataEncryptionKeySet and is used to iterate over the raw logs and unpacked data for AccountDataEncryptionKeySet events raised by the Accounts contract.
type AccountsAccountDataEncryptionKeySetIterator struct {
	Event *AccountsAccountDataEncryptionKeySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsAccountDataEncryptionKeySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAccountDataEncryptionKeySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsAccountDataEncryptionKeySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsAccountDataEncryptionKeySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAccountDataEncryptionKeySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAccountDataEncryptionKeySet represents a AccountDataEncryptionKeySet event raised by the Accounts contract.
type AccountsAccountDataEncryptionKeySet struct {
	Account           common.Address
	DataEncryptionKey []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAccountDataEncryptionKeySet is a free log retrieval operation binding the contract event 0x43fdefe0a824cb0e3bbaf9c4bc97669187996136fe9282382baf10787f0d808d.
//
// Solidity: event AccountDataEncryptionKeySet(address indexed account, bytes dataEncryptionKey)
func (_Accounts *AccountsFilterer) FilterAccountDataEncryptionKeySet(opts *bind.FilterOpts, account []common.Address) (*AccountsAccountDataEncryptionKeySetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AccountDataEncryptionKeySet", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsAccountDataEncryptionKeySetIterator{contract: _Accounts.contract, event: "AccountDataEncryptionKeySet", logs: logs, sub: sub}, nil
}

// WatchAccountDataEncryptionKeySet is a free log subscription operation binding the contract event 0x43fdefe0a824cb0e3bbaf9c4bc97669187996136fe9282382baf10787f0d808d.
//
// Solidity: event AccountDataEncryptionKeySet(address indexed account, bytes dataEncryptionKey)
func (_Accounts *AccountsFilterer) WatchAccountDataEncryptionKeySet(opts *bind.WatchOpts, sink chan<- *AccountsAccountDataEncryptionKeySet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AccountDataEncryptionKeySet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAccountDataEncryptionKeySet)
				if err := _Accounts.contract.UnpackLog(event, "AccountDataEncryptionKeySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountDataEncryptionKeySet is a log parse operation binding the contract event 0x43fdefe0a824cb0e3bbaf9c4bc97669187996136fe9282382baf10787f0d808d.
//
// Solidity: event AccountDataEncryptionKeySet(address indexed account, bytes dataEncryptionKey)
func (_Accounts *AccountsFilterer) ParseAccountDataEncryptionKeySet(log types.Log) (*AccountsAccountDataEncryptionKeySet, error) {
	event := new(AccountsAccountDataEncryptionKeySet)
	if err := _Accounts.contract.UnpackLog(event, "AccountDataEncryptionKeySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsAccountMetadataURLSetIterator is returned from FilterAccountMetadataURLSet and is used to iterate over the raw logs and unpacked data for AccountMetadataURLSet events raised by the Accounts contract.
type AccountsAccountMetadataURLSetIterator struct {
	Event *AccountsAccountMetadataURLSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsAccountMetadataURLSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAccountMetadataURLSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsAccountMetadataURLSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsAccountMetadataURLSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAccountMetadataURLSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAccountMetadataURLSet represents a AccountMetadataURLSet event raised by the Accounts contract.
type AccountsAccountMetadataURLSet struct {
	Account     common.Address
	MetadataURL string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAccountMetadataURLSet is a free log retrieval operation binding the contract event 0x0b5629fec5b6b5a1c2cfe0de7495111627a8cf297dced72e0669527425d3f01b.
//
// Solidity: event AccountMetadataURLSet(address indexed account, string metadataURL)
func (_Accounts *AccountsFilterer) FilterAccountMetadataURLSet(opts *bind.FilterOpts, account []common.Address) (*AccountsAccountMetadataURLSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AccountMetadataURLSet", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsAccountMetadataURLSetIterator{contract: _Accounts.contract, event: "AccountMetadataURLSet", logs: logs, sub: sub}, nil
}

// WatchAccountMetadataURLSet is a free log subscription operation binding the contract event 0x0b5629fec5b6b5a1c2cfe0de7495111627a8cf297dced72e0669527425d3f01b.
//
// Solidity: event AccountMetadataURLSet(address indexed account, string metadataURL)
func (_Accounts *AccountsFilterer) WatchAccountMetadataURLSet(opts *bind.WatchOpts, sink chan<- *AccountsAccountMetadataURLSet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AccountMetadataURLSet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAccountMetadataURLSet)
				if err := _Accounts.contract.UnpackLog(event, "AccountMetadataURLSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountMetadataURLSet is a log parse operation binding the contract event 0x0b5629fec5b6b5a1c2cfe0de7495111627a8cf297dced72e0669527425d3f01b.
//
// Solidity: event AccountMetadataURLSet(address indexed account, string metadataURL)
func (_Accounts *AccountsFilterer) ParseAccountMetadataURLSet(log types.Log) (*AccountsAccountMetadataURLSet, error) {
	event := new(AccountsAccountMetadataURLSet)
	if err := _Accounts.contract.UnpackLog(event, "AccountMetadataURLSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsAccountNameSetIterator is returned from FilterAccountNameSet and is used to iterate over the raw logs and unpacked data for AccountNameSet events raised by the Accounts contract.
type AccountsAccountNameSetIterator struct {
	Event *AccountsAccountNameSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsAccountNameSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAccountNameSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsAccountNameSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsAccountNameSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAccountNameSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAccountNameSet represents a AccountNameSet event raised by the Accounts contract.
type AccountsAccountNameSet struct {
	Account common.Address
	Name    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAccountNameSet is a free log retrieval operation binding the contract event 0xa6e2c5a23bb917ba0a584c4b250257ddad698685829b66a8813c004b39934fe4.
//
// Solidity: event AccountNameSet(address indexed account, string name)
func (_Accounts *AccountsFilterer) FilterAccountNameSet(opts *bind.FilterOpts, account []common.Address) (*AccountsAccountNameSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AccountNameSet", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsAccountNameSetIterator{contract: _Accounts.contract, event: "AccountNameSet", logs: logs, sub: sub}, nil
}

// WatchAccountNameSet is a free log subscription operation binding the contract event 0xa6e2c5a23bb917ba0a584c4b250257ddad698685829b66a8813c004b39934fe4.
//
// Solidity: event AccountNameSet(address indexed account, string name)
func (_Accounts *AccountsFilterer) WatchAccountNameSet(opts *bind.WatchOpts, sink chan<- *AccountsAccountNameSet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AccountNameSet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAccountNameSet)
				if err := _Accounts.contract.UnpackLog(event, "AccountNameSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountNameSet is a log parse operation binding the contract event 0xa6e2c5a23bb917ba0a584c4b250257ddad698685829b66a8813c004b39934fe4.
//
// Solidity: event AccountNameSet(address indexed account, string name)
func (_Accounts *AccountsFilterer) ParseAccountNameSet(log types.Log) (*AccountsAccountNameSet, error) {
	event := new(AccountsAccountNameSet)
	if err := _Accounts.contract.UnpackLog(event, "AccountNameSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsAccountWalletAddressSetIterator is returned from FilterAccountWalletAddressSet and is used to iterate over the raw logs and unpacked data for AccountWalletAddressSet events raised by the Accounts contract.
type AccountsAccountWalletAddressSetIterator struct {
	Event *AccountsAccountWalletAddressSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsAccountWalletAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAccountWalletAddressSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsAccountWalletAddressSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsAccountWalletAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAccountWalletAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAccountWalletAddressSet represents a AccountWalletAddressSet event raised by the Accounts contract.
type AccountsAccountWalletAddressSet struct {
	Account       common.Address
	WalletAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAccountWalletAddressSet is a free log retrieval operation binding the contract event 0xf81d74398fd47e35c36b714019df15f200f623dde569b5b531d6a0b4da5c5f26.
//
// Solidity: event AccountWalletAddressSet(address indexed account, address walletAddress)
func (_Accounts *AccountsFilterer) FilterAccountWalletAddressSet(opts *bind.FilterOpts, account []common.Address) (*AccountsAccountWalletAddressSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AccountWalletAddressSet", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsAccountWalletAddressSetIterator{contract: _Accounts.contract, event: "AccountWalletAddressSet", logs: logs, sub: sub}, nil
}

// WatchAccountWalletAddressSet is a free log subscription operation binding the contract event 0xf81d74398fd47e35c36b714019df15f200f623dde569b5b531d6a0b4da5c5f26.
//
// Solidity: event AccountWalletAddressSet(address indexed account, address walletAddress)
func (_Accounts *AccountsFilterer) WatchAccountWalletAddressSet(opts *bind.WatchOpts, sink chan<- *AccountsAccountWalletAddressSet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AccountWalletAddressSet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAccountWalletAddressSet)
				if err := _Accounts.contract.UnpackLog(event, "AccountWalletAddressSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountWalletAddressSet is a log parse operation binding the contract event 0xf81d74398fd47e35c36b714019df15f200f623dde569b5b531d6a0b4da5c5f26.
//
// Solidity: event AccountWalletAddressSet(address indexed account, address walletAddress)
func (_Accounts *AccountsFilterer) ParseAccountWalletAddressSet(log types.Log) (*AccountsAccountWalletAddressSet, error) {
	event := new(AccountsAccountWalletAddressSet)
	if err := _Accounts.contract.UnpackLog(event, "AccountWalletAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsAttestationSignerAuthorizedIterator is returned from FilterAttestationSignerAuthorized and is used to iterate over the raw logs and unpacked data for AttestationSignerAuthorized events raised by the Accounts contract.
type AccountsAttestationSignerAuthorizedIterator struct {
	Event *AccountsAttestationSignerAuthorized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsAttestationSignerAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAttestationSignerAuthorized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsAttestationSignerAuthorized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsAttestationSignerAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAttestationSignerAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAttestationSignerAuthorized represents a AttestationSignerAuthorized event raised by the Accounts contract.
type AccountsAttestationSignerAuthorized struct {
	Account common.Address
	Signer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAttestationSignerAuthorized is a free log retrieval operation binding the contract event 0x9dfbc5a621c3e2d0d83beee687a17dfc796bbce2118793e5e254409bb265ca0b.
//
// Solidity: event AttestationSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) FilterAttestationSignerAuthorized(opts *bind.FilterOpts, account []common.Address) (*AccountsAttestationSignerAuthorizedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AttestationSignerAuthorized", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsAttestationSignerAuthorizedIterator{contract: _Accounts.contract, event: "AttestationSignerAuthorized", logs: logs, sub: sub}, nil
}

// WatchAttestationSignerAuthorized is a free log subscription operation binding the contract event 0x9dfbc5a621c3e2d0d83beee687a17dfc796bbce2118793e5e254409bb265ca0b.
//
// Solidity: event AttestationSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) WatchAttestationSignerAuthorized(opts *bind.WatchOpts, sink chan<- *AccountsAttestationSignerAuthorized, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AttestationSignerAuthorized", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAttestationSignerAuthorized)
				if err := _Accounts.contract.UnpackLog(event, "AttestationSignerAuthorized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAttestationSignerAuthorized is a log parse operation binding the contract event 0x9dfbc5a621c3e2d0d83beee687a17dfc796bbce2118793e5e254409bb265ca0b.
//
// Solidity: event AttestationSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) ParseAttestationSignerAuthorized(log types.Log) (*AccountsAttestationSignerAuthorized, error) {
	event := new(AccountsAttestationSignerAuthorized)
	if err := _Accounts.contract.UnpackLog(event, "AttestationSignerAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsAttestationSignerRemovedIterator is returned from FilterAttestationSignerRemoved and is used to iterate over the raw logs and unpacked data for AttestationSignerRemoved events raised by the Accounts contract.
type AccountsAttestationSignerRemovedIterator struct {
	Event *AccountsAttestationSignerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsAttestationSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAttestationSignerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsAttestationSignerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsAttestationSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAttestationSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAttestationSignerRemoved represents a AttestationSignerRemoved event raised by the Accounts contract.
type AccountsAttestationSignerRemoved struct {
	Account   common.Address
	OldSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAttestationSignerRemoved is a free log retrieval operation binding the contract event 0x14670729407debb6ed03d885f8ba57155de89ce39bf17127ae4900ec7c2ad103.
//
// Solidity: event AttestationSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) FilterAttestationSignerRemoved(opts *bind.FilterOpts, account []common.Address) (*AccountsAttestationSignerRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AttestationSignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsAttestationSignerRemovedIterator{contract: _Accounts.contract, event: "AttestationSignerRemoved", logs: logs, sub: sub}, nil
}

// WatchAttestationSignerRemoved is a free log subscription operation binding the contract event 0x14670729407debb6ed03d885f8ba57155de89ce39bf17127ae4900ec7c2ad103.
//
// Solidity: event AttestationSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) WatchAttestationSignerRemoved(opts *bind.WatchOpts, sink chan<- *AccountsAttestationSignerRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AttestationSignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAttestationSignerRemoved)
				if err := _Accounts.contract.UnpackLog(event, "AttestationSignerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAttestationSignerRemoved is a log parse operation binding the contract event 0x14670729407debb6ed03d885f8ba57155de89ce39bf17127ae4900ec7c2ad103.
//
// Solidity: event AttestationSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) ParseAttestationSignerRemoved(log types.Log) (*AccountsAttestationSignerRemoved, error) {
	event := new(AccountsAttestationSignerRemoved)
	if err := _Accounts.contract.UnpackLog(event, "AttestationSignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsDefaultSignerRemovedIterator is returned from FilterDefaultSignerRemoved and is used to iterate over the raw logs and unpacked data for DefaultSignerRemoved events raised by the Accounts contract.
type AccountsDefaultSignerRemovedIterator struct {
	Event *AccountsDefaultSignerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsDefaultSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsDefaultSignerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsDefaultSignerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsDefaultSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsDefaultSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsDefaultSignerRemoved represents a DefaultSignerRemoved event raised by the Accounts contract.
type AccountsDefaultSignerRemoved struct {
	Account   common.Address
	OldSigner common.Address
	Role      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDefaultSignerRemoved is a free log retrieval operation binding the contract event 0xe553a3065d5a77d4ec2a0e0c31d49be4bf4d9f4c45883b2d67f61ba9c1b89c5d.
//
// Solidity: event DefaultSignerRemoved(address indexed account, address oldSigner, bytes32 role)
func (_Accounts *AccountsFilterer) FilterDefaultSignerRemoved(opts *bind.FilterOpts, account []common.Address) (*AccountsDefaultSignerRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "DefaultSignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsDefaultSignerRemovedIterator{contract: _Accounts.contract, event: "DefaultSignerRemoved", logs: logs, sub: sub}, nil
}

// WatchDefaultSignerRemoved is a free log subscription operation binding the contract event 0xe553a3065d5a77d4ec2a0e0c31d49be4bf4d9f4c45883b2d67f61ba9c1b89c5d.
//
// Solidity: event DefaultSignerRemoved(address indexed account, address oldSigner, bytes32 role)
func (_Accounts *AccountsFilterer) WatchDefaultSignerRemoved(opts *bind.WatchOpts, sink chan<- *AccountsDefaultSignerRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "DefaultSignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsDefaultSignerRemoved)
				if err := _Accounts.contract.UnpackLog(event, "DefaultSignerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultSignerRemoved is a log parse operation binding the contract event 0xe553a3065d5a77d4ec2a0e0c31d49be4bf4d9f4c45883b2d67f61ba9c1b89c5d.
//
// Solidity: event DefaultSignerRemoved(address indexed account, address oldSigner, bytes32 role)
func (_Accounts *AccountsFilterer) ParseDefaultSignerRemoved(log types.Log) (*AccountsDefaultSignerRemoved, error) {
	event := new(AccountsDefaultSignerRemoved)
	if err := _Accounts.contract.UnpackLog(event, "DefaultSignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsDefaultSignerSetIterator is returned from FilterDefaultSignerSet and is used to iterate over the raw logs and unpacked data for DefaultSignerSet events raised by the Accounts contract.
type AccountsDefaultSignerSetIterator struct {
	Event *AccountsDefaultSignerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsDefaultSignerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsDefaultSignerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsDefaultSignerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsDefaultSignerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsDefaultSignerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsDefaultSignerSet represents a DefaultSignerSet event raised by the Accounts contract.
type AccountsDefaultSignerSet struct {
	Account common.Address
	Signer  common.Address
	Role    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDefaultSignerSet is a free log retrieval operation binding the contract event 0x2613ed414d18d8152e86c896c04ccce344b75a2f06141f04d39ad069a3872523.
//
// Solidity: event DefaultSignerSet(address indexed account, address signer, bytes32 role)
func (_Accounts *AccountsFilterer) FilterDefaultSignerSet(opts *bind.FilterOpts, account []common.Address) (*AccountsDefaultSignerSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "DefaultSignerSet", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsDefaultSignerSetIterator{contract: _Accounts.contract, event: "DefaultSignerSet", logs: logs, sub: sub}, nil
}

// WatchDefaultSignerSet is a free log subscription operation binding the contract event 0x2613ed414d18d8152e86c896c04ccce344b75a2f06141f04d39ad069a3872523.
//
// Solidity: event DefaultSignerSet(address indexed account, address signer, bytes32 role)
func (_Accounts *AccountsFilterer) WatchDefaultSignerSet(opts *bind.WatchOpts, sink chan<- *AccountsDefaultSignerSet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "DefaultSignerSet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsDefaultSignerSet)
				if err := _Accounts.contract.UnpackLog(event, "DefaultSignerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultSignerSet is a log parse operation binding the contract event 0x2613ed414d18d8152e86c896c04ccce344b75a2f06141f04d39ad069a3872523.
//
// Solidity: event DefaultSignerSet(address indexed account, address signer, bytes32 role)
func (_Accounts *AccountsFilterer) ParseDefaultSignerSet(log types.Log) (*AccountsDefaultSignerSet, error) {
	event := new(AccountsDefaultSignerSet)
	if err := _Accounts.contract.UnpackLog(event, "DefaultSignerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsIndexedSignerRemovedIterator is returned from FilterIndexedSignerRemoved and is used to iterate over the raw logs and unpacked data for IndexedSignerRemoved events raised by the Accounts contract.
type AccountsIndexedSignerRemovedIterator struct {
	Event *AccountsIndexedSignerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsIndexedSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsIndexedSignerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsIndexedSignerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsIndexedSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsIndexedSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsIndexedSignerRemoved represents a IndexedSignerRemoved event raised by the Accounts contract.
type AccountsIndexedSignerRemoved struct {
	Account   common.Address
	OldSigner common.Address
	Role      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIndexedSignerRemoved is a free log retrieval operation binding the contract event 0xccc97b55d227538f487c521e1218ba74768b73d088310238027c2ae3b43e9c91.
//
// Solidity: event IndexedSignerRemoved(address indexed account, address oldSigner, bytes32 role)
func (_Accounts *AccountsFilterer) FilterIndexedSignerRemoved(opts *bind.FilterOpts, account []common.Address) (*AccountsIndexedSignerRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "IndexedSignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsIndexedSignerRemovedIterator{contract: _Accounts.contract, event: "IndexedSignerRemoved", logs: logs, sub: sub}, nil
}

// WatchIndexedSignerRemoved is a free log subscription operation binding the contract event 0xccc97b55d227538f487c521e1218ba74768b73d088310238027c2ae3b43e9c91.
//
// Solidity: event IndexedSignerRemoved(address indexed account, address oldSigner, bytes32 role)
func (_Accounts *AccountsFilterer) WatchIndexedSignerRemoved(opts *bind.WatchOpts, sink chan<- *AccountsIndexedSignerRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "IndexedSignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsIndexedSignerRemoved)
				if err := _Accounts.contract.UnpackLog(event, "IndexedSignerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIndexedSignerRemoved is a log parse operation binding the contract event 0xccc97b55d227538f487c521e1218ba74768b73d088310238027c2ae3b43e9c91.
//
// Solidity: event IndexedSignerRemoved(address indexed account, address oldSigner, bytes32 role)
func (_Accounts *AccountsFilterer) ParseIndexedSignerRemoved(log types.Log) (*AccountsIndexedSignerRemoved, error) {
	event := new(AccountsIndexedSignerRemoved)
	if err := _Accounts.contract.UnpackLog(event, "IndexedSignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsIndexedSignerSetIterator is returned from FilterIndexedSignerSet and is used to iterate over the raw logs and unpacked data for IndexedSignerSet events raised by the Accounts contract.
type AccountsIndexedSignerSetIterator struct {
	Event *AccountsIndexedSignerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsIndexedSignerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsIndexedSignerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsIndexedSignerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsIndexedSignerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsIndexedSignerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsIndexedSignerSet represents a IndexedSignerSet event raised by the Accounts contract.
type AccountsIndexedSignerSet struct {
	Account common.Address
	Signer  common.Address
	Role    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIndexedSignerSet is a free log retrieval operation binding the contract event 0x8a00ae3e0722558391733230bfc96d425df2dd7b44f7ce506580785adcf171a2.
//
// Solidity: event IndexedSignerSet(address indexed account, address signer, bytes32 role)
func (_Accounts *AccountsFilterer) FilterIndexedSignerSet(opts *bind.FilterOpts, account []common.Address) (*AccountsIndexedSignerSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "IndexedSignerSet", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsIndexedSignerSetIterator{contract: _Accounts.contract, event: "IndexedSignerSet", logs: logs, sub: sub}, nil
}

// WatchIndexedSignerSet is a free log subscription operation binding the contract event 0x8a00ae3e0722558391733230bfc96d425df2dd7b44f7ce506580785adcf171a2.
//
// Solidity: event IndexedSignerSet(address indexed account, address signer, bytes32 role)
func (_Accounts *AccountsFilterer) WatchIndexedSignerSet(opts *bind.WatchOpts, sink chan<- *AccountsIndexedSignerSet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "IndexedSignerSet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsIndexedSignerSet)
				if err := _Accounts.contract.UnpackLog(event, "IndexedSignerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIndexedSignerSet is a log parse operation binding the contract event 0x8a00ae3e0722558391733230bfc96d425df2dd7b44f7ce506580785adcf171a2.
//
// Solidity: event IndexedSignerSet(address indexed account, address signer, bytes32 role)
func (_Accounts *AccountsFilterer) ParseIndexedSignerSet(log types.Log) (*AccountsIndexedSignerSet, error) {
	event := new(AccountsIndexedSignerSet)
	if err := _Accounts.contract.UnpackLog(event, "IndexedSignerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsLegacySignerRemovedIterator is returned from FilterLegacySignerRemoved and is used to iterate over the raw logs and unpacked data for LegacySignerRemoved events raised by the Accounts contract.
type AccountsLegacySignerRemovedIterator struct {
	Event *AccountsLegacySignerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsLegacySignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsLegacySignerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsLegacySignerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsLegacySignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsLegacySignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsLegacySignerRemoved represents a LegacySignerRemoved event raised by the Accounts contract.
type AccountsLegacySignerRemoved struct {
	Account   common.Address
	OldSigner common.Address
	Role      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLegacySignerRemoved is a free log retrieval operation binding the contract event 0xdd0b0d959c29750e7bfabbb7543a56957699d07edc512d2523ffe7502901ac19.
//
// Solidity: event LegacySignerRemoved(address indexed account, address oldSigner, bytes32 role)
func (_Accounts *AccountsFilterer) FilterLegacySignerRemoved(opts *bind.FilterOpts, account []common.Address) (*AccountsLegacySignerRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "LegacySignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsLegacySignerRemovedIterator{contract: _Accounts.contract, event: "LegacySignerRemoved", logs: logs, sub: sub}, nil
}

// WatchLegacySignerRemoved is a free log subscription operation binding the contract event 0xdd0b0d959c29750e7bfabbb7543a56957699d07edc512d2523ffe7502901ac19.
//
// Solidity: event LegacySignerRemoved(address indexed account, address oldSigner, bytes32 role)
func (_Accounts *AccountsFilterer) WatchLegacySignerRemoved(opts *bind.WatchOpts, sink chan<- *AccountsLegacySignerRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "LegacySignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsLegacySignerRemoved)
				if err := _Accounts.contract.UnpackLog(event, "LegacySignerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLegacySignerRemoved is a log parse operation binding the contract event 0xdd0b0d959c29750e7bfabbb7543a56957699d07edc512d2523ffe7502901ac19.
//
// Solidity: event LegacySignerRemoved(address indexed account, address oldSigner, bytes32 role)
func (_Accounts *AccountsFilterer) ParseLegacySignerRemoved(log types.Log) (*AccountsLegacySignerRemoved, error) {
	event := new(AccountsLegacySignerRemoved)
	if err := _Accounts.contract.UnpackLog(event, "LegacySignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsLegacySignerSetIterator is returned from FilterLegacySignerSet and is used to iterate over the raw logs and unpacked data for LegacySignerSet events raised by the Accounts contract.
type AccountsLegacySignerSetIterator struct {
	Event *AccountsLegacySignerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsLegacySignerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsLegacySignerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsLegacySignerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsLegacySignerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsLegacySignerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsLegacySignerSet represents a LegacySignerSet event raised by the Accounts contract.
type AccountsLegacySignerSet struct {
	Account common.Address
	Signer  common.Address
	Role    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLegacySignerSet is a free log retrieval operation binding the contract event 0xc5cd67202a8095484f559b338b2b6fee0dd81af9f70c4814c6517fcf9a09564c.
//
// Solidity: event LegacySignerSet(address indexed account, address signer, bytes32 role)
func (_Accounts *AccountsFilterer) FilterLegacySignerSet(opts *bind.FilterOpts, account []common.Address) (*AccountsLegacySignerSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "LegacySignerSet", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsLegacySignerSetIterator{contract: _Accounts.contract, event: "LegacySignerSet", logs: logs, sub: sub}, nil
}

// WatchLegacySignerSet is a free log subscription operation binding the contract event 0xc5cd67202a8095484f559b338b2b6fee0dd81af9f70c4814c6517fcf9a09564c.
//
// Solidity: event LegacySignerSet(address indexed account, address signer, bytes32 role)
func (_Accounts *AccountsFilterer) WatchLegacySignerSet(opts *bind.WatchOpts, sink chan<- *AccountsLegacySignerSet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "LegacySignerSet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsLegacySignerSet)
				if err := _Accounts.contract.UnpackLog(event, "LegacySignerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLegacySignerSet is a log parse operation binding the contract event 0xc5cd67202a8095484f559b338b2b6fee0dd81af9f70c4814c6517fcf9a09564c.
//
// Solidity: event LegacySignerSet(address indexed account, address signer, bytes32 role)
func (_Accounts *AccountsFilterer) ParseLegacySignerSet(log types.Log) (*AccountsLegacySignerSet, error) {
	event := new(AccountsLegacySignerSet)
	if err := _Accounts.contract.UnpackLog(event, "LegacySignerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsOffchainStorageRootAddedIterator is returned from FilterOffchainStorageRootAdded and is used to iterate over the raw logs and unpacked data for OffchainStorageRootAdded events raised by the Accounts contract.
type AccountsOffchainStorageRootAddedIterator struct {
	Event *AccountsOffchainStorageRootAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsOffchainStorageRootAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsOffchainStorageRootAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsOffchainStorageRootAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsOffchainStorageRootAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsOffchainStorageRootAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsOffchainStorageRootAdded represents a OffchainStorageRootAdded event raised by the Accounts contract.
type AccountsOffchainStorageRootAdded struct {
	Account common.Address
	Url     []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOffchainStorageRootAdded is a free log retrieval operation binding the contract event 0x15dfb3066a1bbbdaf9a7f62c47db990114058ae1739fd70a90361ea715bbf3c8.
//
// Solidity: event OffchainStorageRootAdded(address indexed account, bytes url)
func (_Accounts *AccountsFilterer) FilterOffchainStorageRootAdded(opts *bind.FilterOpts, account []common.Address) (*AccountsOffchainStorageRootAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "OffchainStorageRootAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsOffchainStorageRootAddedIterator{contract: _Accounts.contract, event: "OffchainStorageRootAdded", logs: logs, sub: sub}, nil
}

// WatchOffchainStorageRootAdded is a free log subscription operation binding the contract event 0x15dfb3066a1bbbdaf9a7f62c47db990114058ae1739fd70a90361ea715bbf3c8.
//
// Solidity: event OffchainStorageRootAdded(address indexed account, bytes url)
func (_Accounts *AccountsFilterer) WatchOffchainStorageRootAdded(opts *bind.WatchOpts, sink chan<- *AccountsOffchainStorageRootAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "OffchainStorageRootAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsOffchainStorageRootAdded)
				if err := _Accounts.contract.UnpackLog(event, "OffchainStorageRootAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOffchainStorageRootAdded is a log parse operation binding the contract event 0x15dfb3066a1bbbdaf9a7f62c47db990114058ae1739fd70a90361ea715bbf3c8.
//
// Solidity: event OffchainStorageRootAdded(address indexed account, bytes url)
func (_Accounts *AccountsFilterer) ParseOffchainStorageRootAdded(log types.Log) (*AccountsOffchainStorageRootAdded, error) {
	event := new(AccountsOffchainStorageRootAdded)
	if err := _Accounts.contract.UnpackLog(event, "OffchainStorageRootAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsOffchainStorageRootRemovedIterator is returned from FilterOffchainStorageRootRemoved and is used to iterate over the raw logs and unpacked data for OffchainStorageRootRemoved events raised by the Accounts contract.
type AccountsOffchainStorageRootRemovedIterator struct {
	Event *AccountsOffchainStorageRootRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsOffchainStorageRootRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsOffchainStorageRootRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsOffchainStorageRootRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsOffchainStorageRootRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsOffchainStorageRootRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsOffchainStorageRootRemoved represents a OffchainStorageRootRemoved event raised by the Accounts contract.
type AccountsOffchainStorageRootRemoved struct {
	Account common.Address
	Url     []byte
	Index   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOffchainStorageRootRemoved is a free log retrieval operation binding the contract event 0xae0f2fa495a3eb65d46fe97b0baea8b6fd7edb243175c70f2455e6e83bc6fbaf.
//
// Solidity: event OffchainStorageRootRemoved(address indexed account, bytes url, uint256 index)
func (_Accounts *AccountsFilterer) FilterOffchainStorageRootRemoved(opts *bind.FilterOpts, account []common.Address) (*AccountsOffchainStorageRootRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "OffchainStorageRootRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsOffchainStorageRootRemovedIterator{contract: _Accounts.contract, event: "OffchainStorageRootRemoved", logs: logs, sub: sub}, nil
}

// WatchOffchainStorageRootRemoved is a free log subscription operation binding the contract event 0xae0f2fa495a3eb65d46fe97b0baea8b6fd7edb243175c70f2455e6e83bc6fbaf.
//
// Solidity: event OffchainStorageRootRemoved(address indexed account, bytes url, uint256 index)
func (_Accounts *AccountsFilterer) WatchOffchainStorageRootRemoved(opts *bind.WatchOpts, sink chan<- *AccountsOffchainStorageRootRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "OffchainStorageRootRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsOffchainStorageRootRemoved)
				if err := _Accounts.contract.UnpackLog(event, "OffchainStorageRootRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOffchainStorageRootRemoved is a log parse operation binding the contract event 0xae0f2fa495a3eb65d46fe97b0baea8b6fd7edb243175c70f2455e6e83bc6fbaf.
//
// Solidity: event OffchainStorageRootRemoved(address indexed account, bytes url, uint256 index)
func (_Accounts *AccountsFilterer) ParseOffchainStorageRootRemoved(log types.Log) (*AccountsOffchainStorageRootRemoved, error) {
	event := new(AccountsOffchainStorageRootRemoved)
	if err := _Accounts.contract.UnpackLog(event, "OffchainStorageRootRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Accounts contract.
type AccountsOwnershipTransferredIterator struct {
	Event *AccountsOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsOwnershipTransferred represents a OwnershipTransferred event raised by the Accounts contract.
type AccountsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Accounts *AccountsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccountsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountsOwnershipTransferredIterator{contract: _Accounts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Accounts *AccountsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsOwnershipTransferred)
				if err := _Accounts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Accounts *AccountsFilterer) ParseOwnershipTransferred(log types.Log) (*AccountsOwnershipTransferred, error) {
	event := new(AccountsOwnershipTransferred)
	if err := _Accounts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsPaymentDelegationSetIterator is returned from FilterPaymentDelegationSet and is used to iterate over the raw logs and unpacked data for PaymentDelegationSet events raised by the Accounts contract.
type AccountsPaymentDelegationSetIterator struct {
	Event *AccountsPaymentDelegationSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsPaymentDelegationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsPaymentDelegationSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsPaymentDelegationSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsPaymentDelegationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsPaymentDelegationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsPaymentDelegationSet represents a PaymentDelegationSet event raised by the Accounts contract.
type AccountsPaymentDelegationSet struct {
	Beneficiary common.Address
	Fraction    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPaymentDelegationSet is a free log retrieval operation binding the contract event 0x3bff8b126c8f283f709ae37dc0d3fc03cae85ca4772cfb25b601f4b0b49ca6df.
//
// Solidity: event PaymentDelegationSet(address indexed beneficiary, uint256 fraction)
func (_Accounts *AccountsFilterer) FilterPaymentDelegationSet(opts *bind.FilterOpts, beneficiary []common.Address) (*AccountsPaymentDelegationSetIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "PaymentDelegationSet", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &AccountsPaymentDelegationSetIterator{contract: _Accounts.contract, event: "PaymentDelegationSet", logs: logs, sub: sub}, nil
}

// WatchPaymentDelegationSet is a free log subscription operation binding the contract event 0x3bff8b126c8f283f709ae37dc0d3fc03cae85ca4772cfb25b601f4b0b49ca6df.
//
// Solidity: event PaymentDelegationSet(address indexed beneficiary, uint256 fraction)
func (_Accounts *AccountsFilterer) WatchPaymentDelegationSet(opts *bind.WatchOpts, sink chan<- *AccountsPaymentDelegationSet, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "PaymentDelegationSet", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsPaymentDelegationSet)
				if err := _Accounts.contract.UnpackLog(event, "PaymentDelegationSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentDelegationSet is a log parse operation binding the contract event 0x3bff8b126c8f283f709ae37dc0d3fc03cae85ca4772cfb25b601f4b0b49ca6df.
//
// Solidity: event PaymentDelegationSet(address indexed beneficiary, uint256 fraction)
func (_Accounts *AccountsFilterer) ParsePaymentDelegationSet(log types.Log) (*AccountsPaymentDelegationSet, error) {
	event := new(AccountsPaymentDelegationSet)
	if err := _Accounts.contract.UnpackLog(event, "PaymentDelegationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the Accounts contract.
type AccountsRegistrySetIterator struct {
	Event *AccountsRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsRegistrySet represents a RegistrySet event raised by the Accounts contract.
type AccountsRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Accounts *AccountsFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*AccountsRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &AccountsRegistrySetIterator{contract: _Accounts.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Accounts *AccountsFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *AccountsRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsRegistrySet)
				if err := _Accounts.contract.UnpackLog(event, "RegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegistrySet is a log parse operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Accounts *AccountsFilterer) ParseRegistrySet(log types.Log) (*AccountsRegistrySet, error) {
	event := new(AccountsRegistrySet)
	if err := _Accounts.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsSignerAuthorizationCompletedIterator is returned from FilterSignerAuthorizationCompleted and is used to iterate over the raw logs and unpacked data for SignerAuthorizationCompleted events raised by the Accounts contract.
type AccountsSignerAuthorizationCompletedIterator struct {
	Event *AccountsSignerAuthorizationCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsSignerAuthorizationCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsSignerAuthorizationCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsSignerAuthorizationCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsSignerAuthorizationCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsSignerAuthorizationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsSignerAuthorizationCompleted represents a SignerAuthorizationCompleted event raised by the Accounts contract.
type AccountsSignerAuthorizationCompleted struct {
	Account common.Address
	Signer  common.Address
	Role    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignerAuthorizationCompleted is a free log retrieval operation binding the contract event 0x9eeca140dda0bdb74fc9acfda0f1c0324e188a732bd48e078a96b16d97eb54e5.
//
// Solidity: event SignerAuthorizationCompleted(address indexed account, address signer, bytes32 indexed role)
func (_Accounts *AccountsFilterer) FilterSignerAuthorizationCompleted(opts *bind.FilterOpts, account []common.Address, role [][32]byte) (*AccountsSignerAuthorizationCompletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "SignerAuthorizationCompleted", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &AccountsSignerAuthorizationCompletedIterator{contract: _Accounts.contract, event: "SignerAuthorizationCompleted", logs: logs, sub: sub}, nil
}

// WatchSignerAuthorizationCompleted is a free log subscription operation binding the contract event 0x9eeca140dda0bdb74fc9acfda0f1c0324e188a732bd48e078a96b16d97eb54e5.
//
// Solidity: event SignerAuthorizationCompleted(address indexed account, address signer, bytes32 indexed role)
func (_Accounts *AccountsFilterer) WatchSignerAuthorizationCompleted(opts *bind.WatchOpts, sink chan<- *AccountsSignerAuthorizationCompleted, account []common.Address, role [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "SignerAuthorizationCompleted", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsSignerAuthorizationCompleted)
				if err := _Accounts.contract.UnpackLog(event, "SignerAuthorizationCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignerAuthorizationCompleted is a log parse operation binding the contract event 0x9eeca140dda0bdb74fc9acfda0f1c0324e188a732bd48e078a96b16d97eb54e5.
//
// Solidity: event SignerAuthorizationCompleted(address indexed account, address signer, bytes32 indexed role)
func (_Accounts *AccountsFilterer) ParseSignerAuthorizationCompleted(log types.Log) (*AccountsSignerAuthorizationCompleted, error) {
	event := new(AccountsSignerAuthorizationCompleted)
	if err := _Accounts.contract.UnpackLog(event, "SignerAuthorizationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsSignerAuthorizationStartedIterator is returned from FilterSignerAuthorizationStarted and is used to iterate over the raw logs and unpacked data for SignerAuthorizationStarted events raised by the Accounts contract.
type AccountsSignerAuthorizationStartedIterator struct {
	Event *AccountsSignerAuthorizationStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsSignerAuthorizationStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsSignerAuthorizationStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsSignerAuthorizationStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsSignerAuthorizationStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsSignerAuthorizationStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsSignerAuthorizationStarted represents a SignerAuthorizationStarted event raised by the Accounts contract.
type AccountsSignerAuthorizationStarted struct {
	Account common.Address
	Signer  common.Address
	Role    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignerAuthorizationStarted is a free log retrieval operation binding the contract event 0x7a162218a1b7bec7fb15b4bb95220fbf423fa3a49718133f5c50361ff1c85376.
//
// Solidity: event SignerAuthorizationStarted(address indexed account, address signer, bytes32 indexed role)
func (_Accounts *AccountsFilterer) FilterSignerAuthorizationStarted(opts *bind.FilterOpts, account []common.Address, role [][32]byte) (*AccountsSignerAuthorizationStartedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "SignerAuthorizationStarted", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &AccountsSignerAuthorizationStartedIterator{contract: _Accounts.contract, event: "SignerAuthorizationStarted", logs: logs, sub: sub}, nil
}

// WatchSignerAuthorizationStarted is a free log subscription operation binding the contract event 0x7a162218a1b7bec7fb15b4bb95220fbf423fa3a49718133f5c50361ff1c85376.
//
// Solidity: event SignerAuthorizationStarted(address indexed account, address signer, bytes32 indexed role)
func (_Accounts *AccountsFilterer) WatchSignerAuthorizationStarted(opts *bind.WatchOpts, sink chan<- *AccountsSignerAuthorizationStarted, account []common.Address, role [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "SignerAuthorizationStarted", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsSignerAuthorizationStarted)
				if err := _Accounts.contract.UnpackLog(event, "SignerAuthorizationStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignerAuthorizationStarted is a log parse operation binding the contract event 0x7a162218a1b7bec7fb15b4bb95220fbf423fa3a49718133f5c50361ff1c85376.
//
// Solidity: event SignerAuthorizationStarted(address indexed account, address signer, bytes32 indexed role)
func (_Accounts *AccountsFilterer) ParseSignerAuthorizationStarted(log types.Log) (*AccountsSignerAuthorizationStarted, error) {
	event := new(AccountsSignerAuthorizationStarted)
	if err := _Accounts.contract.UnpackLog(event, "SignerAuthorizationStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsSignerAuthorizedIterator is returned from FilterSignerAuthorized and is used to iterate over the raw logs and unpacked data for SignerAuthorized events raised by the Accounts contract.
type AccountsSignerAuthorizedIterator struct {
	Event *AccountsSignerAuthorized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsSignerAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsSignerAuthorized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsSignerAuthorized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsSignerAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsSignerAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsSignerAuthorized represents a SignerAuthorized event raised by the Accounts contract.
type AccountsSignerAuthorized struct {
	Account common.Address
	Signer  common.Address
	Role    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignerAuthorized is a free log retrieval operation binding the contract event 0x6cc56bd06daacce5b10fdf5ad1dc781941e14d7a71d29d33e7001e2956d14e07.
//
// Solidity: event SignerAuthorized(address indexed account, address signer, bytes32 indexed role)
func (_Accounts *AccountsFilterer) FilterSignerAuthorized(opts *bind.FilterOpts, account []common.Address, role [][32]byte) (*AccountsSignerAuthorizedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "SignerAuthorized", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &AccountsSignerAuthorizedIterator{contract: _Accounts.contract, event: "SignerAuthorized", logs: logs, sub: sub}, nil
}

// WatchSignerAuthorized is a free log subscription operation binding the contract event 0x6cc56bd06daacce5b10fdf5ad1dc781941e14d7a71d29d33e7001e2956d14e07.
//
// Solidity: event SignerAuthorized(address indexed account, address signer, bytes32 indexed role)
func (_Accounts *AccountsFilterer) WatchSignerAuthorized(opts *bind.WatchOpts, sink chan<- *AccountsSignerAuthorized, account []common.Address, role [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "SignerAuthorized", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsSignerAuthorized)
				if err := _Accounts.contract.UnpackLog(event, "SignerAuthorized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignerAuthorized is a log parse operation binding the contract event 0x6cc56bd06daacce5b10fdf5ad1dc781941e14d7a71d29d33e7001e2956d14e07.
//
// Solidity: event SignerAuthorized(address indexed account, address signer, bytes32 indexed role)
func (_Accounts *AccountsFilterer) ParseSignerAuthorized(log types.Log) (*AccountsSignerAuthorized, error) {
	event := new(AccountsSignerAuthorized)
	if err := _Accounts.contract.UnpackLog(event, "SignerAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsSignerRemovedIterator is returned from FilterSignerRemoved and is used to iterate over the raw logs and unpacked data for SignerRemoved events raised by the Accounts contract.
type AccountsSignerRemovedIterator struct {
	Event *AccountsSignerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsSignerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsSignerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsSignerRemoved represents a SignerRemoved event raised by the Accounts contract.
type AccountsSignerRemoved struct {
	Account   common.Address
	OldSigner common.Address
	Role      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignerRemoved is a free log retrieval operation binding the contract event 0xde9ce22cf1f8631ae2b668300f0493971114f40edd305173bd099ce7100fbe0b.
//
// Solidity: event SignerRemoved(address indexed account, address oldSigner, bytes32 indexed role)
func (_Accounts *AccountsFilterer) FilterSignerRemoved(opts *bind.FilterOpts, account []common.Address, role [][32]byte) (*AccountsSignerRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "SignerRemoved", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &AccountsSignerRemovedIterator{contract: _Accounts.contract, event: "SignerRemoved", logs: logs, sub: sub}, nil
}

// WatchSignerRemoved is a free log subscription operation binding the contract event 0xde9ce22cf1f8631ae2b668300f0493971114f40edd305173bd099ce7100fbe0b.
//
// Solidity: event SignerRemoved(address indexed account, address oldSigner, bytes32 indexed role)
func (_Accounts *AccountsFilterer) WatchSignerRemoved(opts *bind.WatchOpts, sink chan<- *AccountsSignerRemoved, account []common.Address, role [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "SignerRemoved", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsSignerRemoved)
				if err := _Accounts.contract.UnpackLog(event, "SignerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignerRemoved is a log parse operation binding the contract event 0xde9ce22cf1f8631ae2b668300f0493971114f40edd305173bd099ce7100fbe0b.
//
// Solidity: event SignerRemoved(address indexed account, address oldSigner, bytes32 indexed role)
func (_Accounts *AccountsFilterer) ParseSignerRemoved(log types.Log) (*AccountsSignerRemoved, error) {
	event := new(AccountsSignerRemoved)
	if err := _Accounts.contract.UnpackLog(event, "SignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsValidatorSignerAuthorizedIterator is returned from FilterValidatorSignerAuthorized and is used to iterate over the raw logs and unpacked data for ValidatorSignerAuthorized events raised by the Accounts contract.
type AccountsValidatorSignerAuthorizedIterator struct {
	Event *AccountsValidatorSignerAuthorized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsValidatorSignerAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsValidatorSignerAuthorized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsValidatorSignerAuthorized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsValidatorSignerAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsValidatorSignerAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsValidatorSignerAuthorized represents a ValidatorSignerAuthorized event raised by the Accounts contract.
type AccountsValidatorSignerAuthorized struct {
	Account common.Address
	Signer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorSignerAuthorized is a free log retrieval operation binding the contract event 0x16e382723fb40543364faf68863212ba253a099607bf6d3a5b47e50a8bf94943.
//
// Solidity: event ValidatorSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) FilterValidatorSignerAuthorized(opts *bind.FilterOpts, account []common.Address) (*AccountsValidatorSignerAuthorizedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "ValidatorSignerAuthorized", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsValidatorSignerAuthorizedIterator{contract: _Accounts.contract, event: "ValidatorSignerAuthorized", logs: logs, sub: sub}, nil
}

// WatchValidatorSignerAuthorized is a free log subscription operation binding the contract event 0x16e382723fb40543364faf68863212ba253a099607bf6d3a5b47e50a8bf94943.
//
// Solidity: event ValidatorSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) WatchValidatorSignerAuthorized(opts *bind.WatchOpts, sink chan<- *AccountsValidatorSignerAuthorized, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "ValidatorSignerAuthorized", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsValidatorSignerAuthorized)
				if err := _Accounts.contract.UnpackLog(event, "ValidatorSignerAuthorized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorSignerAuthorized is a log parse operation binding the contract event 0x16e382723fb40543364faf68863212ba253a099607bf6d3a5b47e50a8bf94943.
//
// Solidity: event ValidatorSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) ParseValidatorSignerAuthorized(log types.Log) (*AccountsValidatorSignerAuthorized, error) {
	event := new(AccountsValidatorSignerAuthorized)
	if err := _Accounts.contract.UnpackLog(event, "ValidatorSignerAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsValidatorSignerRemovedIterator is returned from FilterValidatorSignerRemoved and is used to iterate over the raw logs and unpacked data for ValidatorSignerRemoved events raised by the Accounts contract.
type AccountsValidatorSignerRemovedIterator struct {
	Event *AccountsValidatorSignerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsValidatorSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsValidatorSignerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsValidatorSignerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsValidatorSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsValidatorSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsValidatorSignerRemoved represents a ValidatorSignerRemoved event raised by the Accounts contract.
type AccountsValidatorSignerRemoved struct {
	Account   common.Address
	OldSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorSignerRemoved is a free log retrieval operation binding the contract event 0xa54764c62865ff0cd3f271fb1d4635662bff10f0878694f1654fb7fbdecb830d.
//
// Solidity: event ValidatorSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) FilterValidatorSignerRemoved(opts *bind.FilterOpts, account []common.Address) (*AccountsValidatorSignerRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "ValidatorSignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsValidatorSignerRemovedIterator{contract: _Accounts.contract, event: "ValidatorSignerRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorSignerRemoved is a free log subscription operation binding the contract event 0xa54764c62865ff0cd3f271fb1d4635662bff10f0878694f1654fb7fbdecb830d.
//
// Solidity: event ValidatorSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) WatchValidatorSignerRemoved(opts *bind.WatchOpts, sink chan<- *AccountsValidatorSignerRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "ValidatorSignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsValidatorSignerRemoved)
				if err := _Accounts.contract.UnpackLog(event, "ValidatorSignerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorSignerRemoved is a log parse operation binding the contract event 0xa54764c62865ff0cd3f271fb1d4635662bff10f0878694f1654fb7fbdecb830d.
//
// Solidity: event ValidatorSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) ParseValidatorSignerRemoved(log types.Log) (*AccountsValidatorSignerRemoved, error) {
	event := new(AccountsValidatorSignerRemoved)
	if err := _Accounts.contract.UnpackLog(event, "ValidatorSignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsVoteSignerAuthorizedIterator is returned from FilterVoteSignerAuthorized and is used to iterate over the raw logs and unpacked data for VoteSignerAuthorized events raised by the Accounts contract.
type AccountsVoteSignerAuthorizedIterator struct {
	Event *AccountsVoteSignerAuthorized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsVoteSignerAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsVoteSignerAuthorized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsVoteSignerAuthorized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsVoteSignerAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsVoteSignerAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsVoteSignerAuthorized represents a VoteSignerAuthorized event raised by the Accounts contract.
type AccountsVoteSignerAuthorized struct {
	Account common.Address
	Signer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVoteSignerAuthorized is a free log retrieval operation binding the contract event 0xaab5f8a189373aaa290f42ae65ea5d7971b732366ca5bf66556e76263944af28.
//
// Solidity: event VoteSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) FilterVoteSignerAuthorized(opts *bind.FilterOpts, account []common.Address) (*AccountsVoteSignerAuthorizedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "VoteSignerAuthorized", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsVoteSignerAuthorizedIterator{contract: _Accounts.contract, event: "VoteSignerAuthorized", logs: logs, sub: sub}, nil
}

// WatchVoteSignerAuthorized is a free log subscription operation binding the contract event 0xaab5f8a189373aaa290f42ae65ea5d7971b732366ca5bf66556e76263944af28.
//
// Solidity: event VoteSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) WatchVoteSignerAuthorized(opts *bind.WatchOpts, sink chan<- *AccountsVoteSignerAuthorized, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "VoteSignerAuthorized", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsVoteSignerAuthorized)
				if err := _Accounts.contract.UnpackLog(event, "VoteSignerAuthorized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteSignerAuthorized is a log parse operation binding the contract event 0xaab5f8a189373aaa290f42ae65ea5d7971b732366ca5bf66556e76263944af28.
//
// Solidity: event VoteSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) ParseVoteSignerAuthorized(log types.Log) (*AccountsVoteSignerAuthorized, error) {
	event := new(AccountsVoteSignerAuthorized)
	if err := _Accounts.contract.UnpackLog(event, "VoteSignerAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsVoteSignerRemovedIterator is returned from FilterVoteSignerRemoved and is used to iterate over the raw logs and unpacked data for VoteSignerRemoved events raised by the Accounts contract.
type AccountsVoteSignerRemovedIterator struct {
	Event *AccountsVoteSignerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccountsVoteSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsVoteSignerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccountsVoteSignerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccountsVoteSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsVoteSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsVoteSignerRemoved represents a VoteSignerRemoved event raised by the Accounts contract.
type AccountsVoteSignerRemoved struct {
	Account   common.Address
	OldSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoteSignerRemoved is a free log retrieval operation binding the contract event 0xa197481f404d8a8082368ad7445380f01e75f27dea6b7aef234a4ce071127fae.
//
// Solidity: event VoteSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) FilterVoteSignerRemoved(opts *bind.FilterOpts, account []common.Address) (*AccountsVoteSignerRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "VoteSignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsVoteSignerRemovedIterator{contract: _Accounts.contract, event: "VoteSignerRemoved", logs: logs, sub: sub}, nil
}

// WatchVoteSignerRemoved is a free log subscription operation binding the contract event 0xa197481f404d8a8082368ad7445380f01e75f27dea6b7aef234a4ce071127fae.
//
// Solidity: event VoteSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) WatchVoteSignerRemoved(opts *bind.WatchOpts, sink chan<- *AccountsVoteSignerRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "VoteSignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsVoteSignerRemoved)
				if err := _Accounts.contract.UnpackLog(event, "VoteSignerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteSignerRemoved is a log parse operation binding the contract event 0xa197481f404d8a8082368ad7445380f01e75f27dea6b7aef234a4ce071127fae.
//
// Solidity: event VoteSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) ParseVoteSignerRemoved(log types.Log) (*AccountsVoteSignerRemoved, error) {
	event := new(AccountsVoteSignerRemoved)
	if err := _Accounts.contract.UnpackLog(event, "VoteSignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
