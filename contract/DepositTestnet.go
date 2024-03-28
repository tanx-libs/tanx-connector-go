// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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
	_ = abi.ConvertType
)

// DepositTestnetMetaData contains all meta data concerning the DepositTestnet contract.
var DepositTestnetMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"acceptedGovernor\",\"type\":\"address\"}],\"name\":\"LogNewGovernorAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nominatedGovernor\",\"type\":\"address\"}],\"name\":\"LogNominatedGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogNominationCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entryId\",\"type\":\"string\"}],\"name\":\"LogRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entryId\",\"type\":\"string\"}],\"name\":\"LogRemovalIntent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entryId\",\"type\":\"string\"}],\"name\":\"LogRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"removedGovernor\",\"type\":\"address\"}],\"name\":\"LogRemovedGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogUnFrozen\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSIT_CANCEL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FREEZE_GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAIN_GOVERNANCE_INFO_TAG\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FORCED_ACTIONS_REQS_PER_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_VERIFIER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNFREEZE_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_REMOVAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"announceAvailabilityVerifierRemovalIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"announceVerifierRemovalIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredAvailabilityVerifiers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_verifers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredVerifiers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_verifers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"name\":\"isAvailabilityVerifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"name\":\"isVerifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainAcceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainCancelNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"testGovernor\",\"type\":\"address\"}],\"name\":\"mainIsGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"mainNominateNewGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governorForRemoval\",\"type\":\"address\"}],\"name\":\"mainRemoveGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"name\":\"registerAvailabilityVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"name\":\"registerVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"removeAvailabilityVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"removeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unFreeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositorEthKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogDepositCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDepositCancelReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogDepositNftCancelReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogMintWithdrawalPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogMintableWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositorEthKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogNftDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogNftWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"LogNftWithdrawalPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAdmin\",\"type\":\"address\"}],\"name\":\"LogTokenAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAdmin\",\"type\":\"address\"}],\"name\":\"LogTokenAdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantum\",\"type\":\"uint256\"}],\"name\":\"LogTokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"LogUserRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"LogWithdrawalPerformed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"defaultVaultWithdrawalLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositCancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositNftReclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositReclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actionIndex\",\"type\":\"uint256\"}],\"name\":\"getActionHashByIndex\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"getAssetInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getCancellationRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"request\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getDepositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"}],\"name\":\"getEthKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getFullWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getQuantizedDepositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"presumedAssetType\",\"type\":\"uint256\"}],\"name\":\"getQuantum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quantum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"isAssetRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"testedAdmin\",\"type\":\"address\"}],\"name\":\"isTokenAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"registerAndDepositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"registerAndDepositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"starkSignature\",\"type\":\"bytes\"}],\"name\":\"registerEthAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"starkSignature\",\"type\":\"bytes\"}],\"name\":\"registerSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"quantum\",\"type\":\"uint256\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"registerTokenAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"}],\"name\":\"unregisterTokenAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"mintingBlob\",\"type\":\"bytes\"}],\"name\":\"withdrawAndMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"LogOperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"LogOperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderRoot\",\"type\":\"uint256\"}],\"name\":\"LogRootUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateTransitionFact\",\"type\":\"bytes32\"}],\"name\":\"LogStateTransitionFact\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"quantizedAmountChange\",\"type\":\"int256\"}],\"name\":\"LogVaultBalanceChangeApplied\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"STARKEX_MAX_DEFAULT_VAULT_LOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"escape\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBatchId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrderRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrderTreeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultTreeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"testedOperator\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"removedOperator\",\"type\":\"address\"}],\"name\":\"unregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"publicInput\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"applicationData\",\"type\":\"uint256[]\"}],\"name\":\"updateState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"LogFullWithdrawalRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"freezeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"fullWithdrawalRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDefaultLockTime\",\"type\":\"uint256\"}],\"name\":\"LogDefaultVaultWithdrawalLockSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDepositToVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeRelease\",\"type\":\"uint256\"}],\"name\":\"LogVaultWithdrawalLockSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawalFromVault\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"depositERC20ToVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositEthToVault\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getQuantizedVaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getVaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getVaultWithdrawalLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isStrictVaultBalancePolicy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"isVaultLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"}],\"name\":\"lockVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDefaultTime\",\"type\":\"uint256\"}],\"name\":\"setDefaultVaultWithdrawalLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedActivationTime\",\"type\":\"uint256\"}],\"name\":\"ImplementationActivationRescheduled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"finalize\",\"type\":\"bool\"}],\"name\":\"updateImplementationActivationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DepositTestnetABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositTestnetMetaData.ABI instead.
var DepositTestnetABI = DepositTestnetMetaData.ABI

// DepositTestnet is an auto generated Go binding around an Ethereum contract.
type DepositTestnet struct {
	DepositTestnetCaller     // Read-only binding to the contract
	DepositTestnetTransactor // Write-only binding to the contract
	DepositTestnetFilterer   // Log filterer for contract events
}

// DepositTestnetCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositTestnetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositTestnetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositTestnetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositTestnetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositTestnetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositTestnetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositTestnetSession struct {
	Contract     *DepositTestnet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositTestnetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositTestnetCallerSession struct {
	Contract *DepositTestnetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DepositTestnetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositTestnetTransactorSession struct {
	Contract     *DepositTestnetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DepositTestnetRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositTestnetRaw struct {
	Contract *DepositTestnet // Generic contract binding to access the raw methods on
}

// DepositTestnetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositTestnetCallerRaw struct {
	Contract *DepositTestnetCaller // Generic read-only contract binding to access the raw methods on
}

// DepositTestnetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositTestnetTransactorRaw struct {
	Contract *DepositTestnetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositTestnet creates a new instance of DepositTestnet, bound to a specific deployed contract.
func NewDepositTestnet(address common.Address, backend bind.ContractBackend) (*DepositTestnet, error) {
	contract, err := bindDepositTestnet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DepositTestnet{DepositTestnetCaller: DepositTestnetCaller{contract: contract}, DepositTestnetTransactor: DepositTestnetTransactor{contract: contract}, DepositTestnetFilterer: DepositTestnetFilterer{contract: contract}}, nil
}

// NewDepositTestnetCaller creates a new read-only instance of DepositTestnet, bound to a specific deployed contract.
func NewDepositTestnetCaller(address common.Address, caller bind.ContractCaller) (*DepositTestnetCaller, error) {
	contract, err := bindDepositTestnet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositTestnetCaller{contract: contract}, nil
}

// NewDepositTestnetTransactor creates a new write-only instance of DepositTestnet, bound to a specific deployed contract.
func NewDepositTestnetTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositTestnetTransactor, error) {
	contract, err := bindDepositTestnet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositTestnetTransactor{contract: contract}, nil
}

// NewDepositTestnetFilterer creates a new log filterer instance of DepositTestnet, bound to a specific deployed contract.
func NewDepositTestnetFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositTestnetFilterer, error) {
	contract, err := bindDepositTestnet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositTestnetFilterer{contract: contract}, nil
}

// bindDepositTestnet binds a generic wrapper to an already deployed contract.
func bindDepositTestnet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DepositTestnetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositTestnet *DepositTestnetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositTestnet.Contract.DepositTestnetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositTestnet *DepositTestnetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositTestnetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositTestnet *DepositTestnetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositTestnetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositTestnet *DepositTestnetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositTestnet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositTestnet *DepositTestnetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositTestnet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositTestnet *DepositTestnetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositTestnet.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_DepositTestnet *DepositTestnetCaller) DEPOSITCANCELDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "DEPOSIT_CANCEL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_DepositTestnet *DepositTestnetSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _DepositTestnet.Contract.DEPOSITCANCELDELAY(&_DepositTestnet.CallOpts)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_DepositTestnet *DepositTestnetCallerSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _DepositTestnet.Contract.DEPOSITCANCELDELAY(&_DepositTestnet.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_DepositTestnet *DepositTestnetCaller) FREEZEGRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "FREEZE_GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_DepositTestnet *DepositTestnetSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _DepositTestnet.Contract.FREEZEGRACEPERIOD(&_DepositTestnet.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_DepositTestnet *DepositTestnetCallerSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _DepositTestnet.Contract.FREEZEGRACEPERIOD(&_DepositTestnet.CallOpts)
}

// MAINGOVERNANCEINFOTAG is a free data retrieval call binding the contract method 0xc23b60ef.
//
// Solidity: function MAIN_GOVERNANCE_INFO_TAG() view returns(string)
func (_DepositTestnet *DepositTestnetCaller) MAINGOVERNANCEINFOTAG(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "MAIN_GOVERNANCE_INFO_TAG")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MAINGOVERNANCEINFOTAG is a free data retrieval call binding the contract method 0xc23b60ef.
//
// Solidity: function MAIN_GOVERNANCE_INFO_TAG() view returns(string)
func (_DepositTestnet *DepositTestnetSession) MAINGOVERNANCEINFOTAG() (string, error) {
	return _DepositTestnet.Contract.MAINGOVERNANCEINFOTAG(&_DepositTestnet.CallOpts)
}

// MAINGOVERNANCEINFOTAG is a free data retrieval call binding the contract method 0xc23b60ef.
//
// Solidity: function MAIN_GOVERNANCE_INFO_TAG() view returns(string)
func (_DepositTestnet *DepositTestnetCallerSession) MAINGOVERNANCEINFOTAG() (string, error) {
	return _DepositTestnet.Contract.MAINGOVERNANCEINFOTAG(&_DepositTestnet.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_DepositTestnet *DepositTestnetCaller) MAXFORCEDACTIONSREQSPERBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "MAX_FORCED_ACTIONS_REQS_PER_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_DepositTestnet *DepositTestnetSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _DepositTestnet.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_DepositTestnet.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_DepositTestnet *DepositTestnetCallerSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _DepositTestnet.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_DepositTestnet.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_DepositTestnet *DepositTestnetCaller) MAXVERIFIERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "MAX_VERIFIER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_DepositTestnet *DepositTestnetSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _DepositTestnet.Contract.MAXVERIFIERCOUNT(&_DepositTestnet.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_DepositTestnet *DepositTestnetCallerSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _DepositTestnet.Contract.MAXVERIFIERCOUNT(&_DepositTestnet.CallOpts)
}

// STARKEXMAXDEFAULTVAULTLOCK is a free data retrieval call binding the contract method 0x835a71c3.
//
// Solidity: function STARKEX_MAX_DEFAULT_VAULT_LOCK() view returns(uint256)
func (_DepositTestnet *DepositTestnetCaller) STARKEXMAXDEFAULTVAULTLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "STARKEX_MAX_DEFAULT_VAULT_LOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STARKEXMAXDEFAULTVAULTLOCK is a free data retrieval call binding the contract method 0x835a71c3.
//
// Solidity: function STARKEX_MAX_DEFAULT_VAULT_LOCK() view returns(uint256)
func (_DepositTestnet *DepositTestnetSession) STARKEXMAXDEFAULTVAULTLOCK() (*big.Int, error) {
	return _DepositTestnet.Contract.STARKEXMAXDEFAULTVAULTLOCK(&_DepositTestnet.CallOpts)
}

// STARKEXMAXDEFAULTVAULTLOCK is a free data retrieval call binding the contract method 0x835a71c3.
//
// Solidity: function STARKEX_MAX_DEFAULT_VAULT_LOCK() view returns(uint256)
func (_DepositTestnet *DepositTestnetCallerSession) STARKEXMAXDEFAULTVAULTLOCK() (*big.Int, error) {
	return _DepositTestnet.Contract.STARKEXMAXDEFAULTVAULTLOCK(&_DepositTestnet.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_DepositTestnet *DepositTestnetCaller) UNFREEZEDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "UNFREEZE_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_DepositTestnet *DepositTestnetSession) UNFREEZEDELAY() (*big.Int, error) {
	return _DepositTestnet.Contract.UNFREEZEDELAY(&_DepositTestnet.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_DepositTestnet *DepositTestnetCallerSession) UNFREEZEDELAY() (*big.Int, error) {
	return _DepositTestnet.Contract.UNFREEZEDELAY(&_DepositTestnet.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_DepositTestnet *DepositTestnetCaller) VERIFIERREMOVALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "VERIFIER_REMOVAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_DepositTestnet *DepositTestnetSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _DepositTestnet.Contract.VERIFIERREMOVALDELAY(&_DepositTestnet.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_DepositTestnet *DepositTestnetCallerSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _DepositTestnet.Contract.VERIFIERREMOVALDELAY(&_DepositTestnet.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_DepositTestnet *DepositTestnetCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_DepositTestnet *DepositTestnetSession) VERSION() (string, error) {
	return _DepositTestnet.Contract.VERSION(&_DepositTestnet.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_DepositTestnet *DepositTestnetCallerSession) VERSION() (string, error) {
	return _DepositTestnet.Contract.VERSION(&_DepositTestnet.CallOpts)
}

// DefaultVaultWithdrawalLock is a free data retrieval call binding the contract method 0xa45d7841.
//
// Solidity: function defaultVaultWithdrawalLock() view returns(uint256)
func (_DepositTestnet *DepositTestnetCaller) DefaultVaultWithdrawalLock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "defaultVaultWithdrawalLock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultVaultWithdrawalLock is a free data retrieval call binding the contract method 0xa45d7841.
//
// Solidity: function defaultVaultWithdrawalLock() view returns(uint256)
func (_DepositTestnet *DepositTestnetSession) DefaultVaultWithdrawalLock() (*big.Int, error) {
	return _DepositTestnet.Contract.DefaultVaultWithdrawalLock(&_DepositTestnet.CallOpts)
}

// DefaultVaultWithdrawalLock is a free data retrieval call binding the contract method 0xa45d7841.
//
// Solidity: function defaultVaultWithdrawalLock() view returns(uint256)
func (_DepositTestnet *DepositTestnetCallerSession) DefaultVaultWithdrawalLock() (*big.Int, error) {
	return _DepositTestnet.Contract.DefaultVaultWithdrawalLock(&_DepositTestnet.CallOpts)
}

// GetActionCount is a free data retrieval call binding the contract method 0x5eecd218.
//
// Solidity: function getActionCount() view returns(uint256)
func (_DepositTestnet *DepositTestnetCaller) GetActionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getActionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActionCount is a free data retrieval call binding the contract method 0x5eecd218.
//
// Solidity: function getActionCount() view returns(uint256)
func (_DepositTestnet *DepositTestnetSession) GetActionCount() (*big.Int, error) {
	return _DepositTestnet.Contract.GetActionCount(&_DepositTestnet.CallOpts)
}

// GetActionCount is a free data retrieval call binding the contract method 0x5eecd218.
//
// Solidity: function getActionCount() view returns(uint256)
func (_DepositTestnet *DepositTestnetCallerSession) GetActionCount() (*big.Int, error) {
	return _DepositTestnet.Contract.GetActionCount(&_DepositTestnet.CallOpts)
}

// GetActionHashByIndex is a free data retrieval call binding the contract method 0x5e586cd1.
//
// Solidity: function getActionHashByIndex(uint256 actionIndex) view returns(bytes32)
func (_DepositTestnet *DepositTestnetCaller) GetActionHashByIndex(opts *bind.CallOpts, actionIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getActionHashByIndex", actionIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetActionHashByIndex is a free data retrieval call binding the contract method 0x5e586cd1.
//
// Solidity: function getActionHashByIndex(uint256 actionIndex) view returns(bytes32)
func (_DepositTestnet *DepositTestnetSession) GetActionHashByIndex(actionIndex *big.Int) ([32]byte, error) {
	return _DepositTestnet.Contract.GetActionHashByIndex(&_DepositTestnet.CallOpts, actionIndex)
}

// GetActionHashByIndex is a free data retrieval call binding the contract method 0x5e586cd1.
//
// Solidity: function getActionHashByIndex(uint256 actionIndex) view returns(bytes32)
func (_DepositTestnet *DepositTestnetCallerSession) GetActionHashByIndex(actionIndex *big.Int) ([32]byte, error) {
	return _DepositTestnet.Contract.GetActionHashByIndex(&_DepositTestnet.CallOpts, actionIndex)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_DepositTestnet *DepositTestnetCaller) GetAssetInfo(opts *bind.CallOpts, assetType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getAssetInfo", assetType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_DepositTestnet *DepositTestnetSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _DepositTestnet.Contract.GetAssetInfo(&_DepositTestnet.CallOpts, assetType)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_DepositTestnet *DepositTestnetCallerSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _DepositTestnet.Contract.GetAssetInfo(&_DepositTestnet.CallOpts, assetType)
}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_DepositTestnet *DepositTestnetCaller) GetCancellationRequest(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getCancellationRequest", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_DepositTestnet *DepositTestnetSession) GetCancellationRequest(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetCancellationRequest(&_DepositTestnet.CallOpts, starkKey, assetId, vaultId)
}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_DepositTestnet *DepositTestnetCallerSession) GetCancellationRequest(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetCancellationRequest(&_DepositTestnet.CallOpts, starkKey, assetId, vaultId)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_DepositTestnet *DepositTestnetCaller) GetDepositBalance(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getDepositBalance", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_DepositTestnet *DepositTestnetSession) GetDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetDepositBalance(&_DepositTestnet.CallOpts, starkKey, assetId, vaultId)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_DepositTestnet *DepositTestnetCallerSession) GetDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetDepositBalance(&_DepositTestnet.CallOpts, starkKey, assetId, vaultId)
}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 ownerKey) view returns(address)
func (_DepositTestnet *DepositTestnetCaller) GetEthKey(opts *bind.CallOpts, ownerKey *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getEthKey", ownerKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 ownerKey) view returns(address)
func (_DepositTestnet *DepositTestnetSession) GetEthKey(ownerKey *big.Int) (common.Address, error) {
	return _DepositTestnet.Contract.GetEthKey(&_DepositTestnet.CallOpts, ownerKey)
}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 ownerKey) view returns(address)
func (_DepositTestnet *DepositTestnetCallerSession) GetEthKey(ownerKey *big.Int) (common.Address, error) {
	return _DepositTestnet.Contract.GetEthKey(&_DepositTestnet.CallOpts, ownerKey)
}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 starkKey, uint256 vaultId) view returns(uint256 res)
func (_DepositTestnet *DepositTestnetCaller) GetFullWithdrawalRequest(opts *bind.CallOpts, starkKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getFullWithdrawalRequest", starkKey, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 starkKey, uint256 vaultId) view returns(uint256 res)
func (_DepositTestnet *DepositTestnetSession) GetFullWithdrawalRequest(starkKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetFullWithdrawalRequest(&_DepositTestnet.CallOpts, starkKey, vaultId)
}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 starkKey, uint256 vaultId) view returns(uint256 res)
func (_DepositTestnet *DepositTestnetCallerSession) GetFullWithdrawalRequest(starkKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetFullWithdrawalRequest(&_DepositTestnet.CallOpts, starkKey, vaultId)
}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256 batchId)
func (_DepositTestnet *DepositTestnetCaller) GetLastBatchId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getLastBatchId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256 batchId)
func (_DepositTestnet *DepositTestnetSession) GetLastBatchId() (*big.Int, error) {
	return _DepositTestnet.Contract.GetLastBatchId(&_DepositTestnet.CallOpts)
}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256 batchId)
func (_DepositTestnet *DepositTestnetCallerSession) GetLastBatchId() (*big.Int, error) {
	return _DepositTestnet.Contract.GetLastBatchId(&_DepositTestnet.CallOpts)
}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256 root)
func (_DepositTestnet *DepositTestnetCaller) GetOrderRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getOrderRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256 root)
func (_DepositTestnet *DepositTestnetSession) GetOrderRoot() (*big.Int, error) {
	return _DepositTestnet.Contract.GetOrderRoot(&_DepositTestnet.CallOpts)
}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256 root)
func (_DepositTestnet *DepositTestnetCallerSession) GetOrderRoot() (*big.Int, error) {
	return _DepositTestnet.Contract.GetOrderRoot(&_DepositTestnet.CallOpts)
}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256 height)
func (_DepositTestnet *DepositTestnetCaller) GetOrderTreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getOrderTreeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256 height)
func (_DepositTestnet *DepositTestnetSession) GetOrderTreeHeight() (*big.Int, error) {
	return _DepositTestnet.Contract.GetOrderTreeHeight(&_DepositTestnet.CallOpts)
}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256 height)
func (_DepositTestnet *DepositTestnetCallerSession) GetOrderTreeHeight() (*big.Int, error) {
	return _DepositTestnet.Contract.GetOrderTreeHeight(&_DepositTestnet.CallOpts)
}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_DepositTestnet *DepositTestnetCaller) GetQuantizedDepositBalance(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getQuantizedDepositBalance", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_DepositTestnet *DepositTestnetSession) GetQuantizedDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetQuantizedDepositBalance(&_DepositTestnet.CallOpts, starkKey, assetId, vaultId)
}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_DepositTestnet *DepositTestnetCallerSession) GetQuantizedDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetQuantizedDepositBalance(&_DepositTestnet.CallOpts, starkKey, assetId, vaultId)
}

// GetQuantizedVaultBalance is a free data retrieval call binding the contract method 0xe69a9de7.
//
// Solidity: function getQuantizedVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositTestnet *DepositTestnetCaller) GetQuantizedVaultBalance(opts *bind.CallOpts, ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getQuantizedVaultBalance", ethKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantizedVaultBalance is a free data retrieval call binding the contract method 0xe69a9de7.
//
// Solidity: function getQuantizedVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositTestnet *DepositTestnetSession) GetQuantizedVaultBalance(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetQuantizedVaultBalance(&_DepositTestnet.CallOpts, ethKey, assetId, vaultId)
}

// GetQuantizedVaultBalance is a free data retrieval call binding the contract method 0xe69a9de7.
//
// Solidity: function getQuantizedVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositTestnet *DepositTestnetCallerSession) GetQuantizedVaultBalance(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetQuantizedVaultBalance(&_DepositTestnet.CallOpts, ethKey, assetId, vaultId)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_DepositTestnet *DepositTestnetCaller) GetQuantum(opts *bind.CallOpts, presumedAssetType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getQuantum", presumedAssetType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_DepositTestnet *DepositTestnetSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetQuantum(&_DepositTestnet.CallOpts, presumedAssetType)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_DepositTestnet *DepositTestnetCallerSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetQuantum(&_DepositTestnet.CallOpts, presumedAssetType)
}

// GetRegisteredAvailabilityVerifiers is a free data retrieval call binding the contract method 0x1ac347f2.
//
// Solidity: function getRegisteredAvailabilityVerifiers() view returns(address[] _verifers)
func (_DepositTestnet *DepositTestnetCaller) GetRegisteredAvailabilityVerifiers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getRegisteredAvailabilityVerifiers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredAvailabilityVerifiers is a free data retrieval call binding the contract method 0x1ac347f2.
//
// Solidity: function getRegisteredAvailabilityVerifiers() view returns(address[] _verifers)
func (_DepositTestnet *DepositTestnetSession) GetRegisteredAvailabilityVerifiers() ([]common.Address, error) {
	return _DepositTestnet.Contract.GetRegisteredAvailabilityVerifiers(&_DepositTestnet.CallOpts)
}

// GetRegisteredAvailabilityVerifiers is a free data retrieval call binding the contract method 0x1ac347f2.
//
// Solidity: function getRegisteredAvailabilityVerifiers() view returns(address[] _verifers)
func (_DepositTestnet *DepositTestnetCallerSession) GetRegisteredAvailabilityVerifiers() ([]common.Address, error) {
	return _DepositTestnet.Contract.GetRegisteredAvailabilityVerifiers(&_DepositTestnet.CallOpts)
}

// GetRegisteredVerifiers is a free data retrieval call binding the contract method 0x4eab9ed3.
//
// Solidity: function getRegisteredVerifiers() view returns(address[] _verifers)
func (_DepositTestnet *DepositTestnetCaller) GetRegisteredVerifiers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getRegisteredVerifiers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredVerifiers is a free data retrieval call binding the contract method 0x4eab9ed3.
//
// Solidity: function getRegisteredVerifiers() view returns(address[] _verifers)
func (_DepositTestnet *DepositTestnetSession) GetRegisteredVerifiers() ([]common.Address, error) {
	return _DepositTestnet.Contract.GetRegisteredVerifiers(&_DepositTestnet.CallOpts)
}

// GetRegisteredVerifiers is a free data retrieval call binding the contract method 0x4eab9ed3.
//
// Solidity: function getRegisteredVerifiers() view returns(address[] _verifers)
func (_DepositTestnet *DepositTestnetCallerSession) GetRegisteredVerifiers() ([]common.Address, error) {
	return _DepositTestnet.Contract.GetRegisteredVerifiers(&_DepositTestnet.CallOpts)
}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256 seq)
func (_DepositTestnet *DepositTestnetCaller) GetSequenceNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getSequenceNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256 seq)
func (_DepositTestnet *DepositTestnetSession) GetSequenceNumber() (*big.Int, error) {
	return _DepositTestnet.Contract.GetSequenceNumber(&_DepositTestnet.CallOpts)
}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256 seq)
func (_DepositTestnet *DepositTestnetCallerSession) GetSequenceNumber() (*big.Int, error) {
	return _DepositTestnet.Contract.GetSequenceNumber(&_DepositTestnet.CallOpts)
}

// GetVaultBalance is a free data retrieval call binding the contract method 0x1c6bd544.
//
// Solidity: function getVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositTestnet *DepositTestnetCaller) GetVaultBalance(opts *bind.CallOpts, ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getVaultBalance", ethKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultBalance is a free data retrieval call binding the contract method 0x1c6bd544.
//
// Solidity: function getVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositTestnet *DepositTestnetSession) GetVaultBalance(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetVaultBalance(&_DepositTestnet.CallOpts, ethKey, assetId, vaultId)
}

// GetVaultBalance is a free data retrieval call binding the contract method 0x1c6bd544.
//
// Solidity: function getVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositTestnet *DepositTestnetCallerSession) GetVaultBalance(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetVaultBalance(&_DepositTestnet.CallOpts, ethKey, assetId, vaultId)
}

// GetVaultRoot is a free data retrieval call binding the contract method 0x64da5dfe.
//
// Solidity: function getVaultRoot() view returns(uint256 root)
func (_DepositTestnet *DepositTestnetCaller) GetVaultRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getVaultRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultRoot is a free data retrieval call binding the contract method 0x64da5dfe.
//
// Solidity: function getVaultRoot() view returns(uint256 root)
func (_DepositTestnet *DepositTestnetSession) GetVaultRoot() (*big.Int, error) {
	return _DepositTestnet.Contract.GetVaultRoot(&_DepositTestnet.CallOpts)
}

// GetVaultRoot is a free data retrieval call binding the contract method 0x64da5dfe.
//
// Solidity: function getVaultRoot() view returns(uint256 root)
func (_DepositTestnet *DepositTestnetCallerSession) GetVaultRoot() (*big.Int, error) {
	return _DepositTestnet.Contract.GetVaultRoot(&_DepositTestnet.CallOpts)
}

// GetVaultTreeHeight is a free data retrieval call binding the contract method 0xf288a3ff.
//
// Solidity: function getVaultTreeHeight() view returns(uint256 height)
func (_DepositTestnet *DepositTestnetCaller) GetVaultTreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getVaultTreeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultTreeHeight is a free data retrieval call binding the contract method 0xf288a3ff.
//
// Solidity: function getVaultTreeHeight() view returns(uint256 height)
func (_DepositTestnet *DepositTestnetSession) GetVaultTreeHeight() (*big.Int, error) {
	return _DepositTestnet.Contract.GetVaultTreeHeight(&_DepositTestnet.CallOpts)
}

// GetVaultTreeHeight is a free data retrieval call binding the contract method 0xf288a3ff.
//
// Solidity: function getVaultTreeHeight() view returns(uint256 height)
func (_DepositTestnet *DepositTestnetCallerSession) GetVaultTreeHeight() (*big.Int, error) {
	return _DepositTestnet.Contract.GetVaultTreeHeight(&_DepositTestnet.CallOpts)
}

// GetVaultWithdrawalLock is a free data retrieval call binding the contract method 0x9de5c4e5.
//
// Solidity: function getVaultWithdrawalLock(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositTestnet *DepositTestnetCaller) GetVaultWithdrawalLock(opts *bind.CallOpts, ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getVaultWithdrawalLock", ethKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultWithdrawalLock is a free data retrieval call binding the contract method 0x9de5c4e5.
//
// Solidity: function getVaultWithdrawalLock(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositTestnet *DepositTestnetSession) GetVaultWithdrawalLock(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetVaultWithdrawalLock(&_DepositTestnet.CallOpts, ethKey, assetId, vaultId)
}

// GetVaultWithdrawalLock is a free data retrieval call binding the contract method 0x9de5c4e5.
//
// Solidity: function getVaultWithdrawalLock(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositTestnet *DepositTestnetCallerSession) GetVaultWithdrawalLock(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetVaultWithdrawalLock(&_DepositTestnet.CallOpts, ethKey, assetId, vaultId)
}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 ownerKey, uint256 assetId) view returns(uint256 balance)
func (_DepositTestnet *DepositTestnetCaller) GetWithdrawalBalance(opts *bind.CallOpts, ownerKey *big.Int, assetId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "getWithdrawalBalance", ownerKey, assetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 ownerKey, uint256 assetId) view returns(uint256 balance)
func (_DepositTestnet *DepositTestnetSession) GetWithdrawalBalance(ownerKey *big.Int, assetId *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetWithdrawalBalance(&_DepositTestnet.CallOpts, ownerKey, assetId)
}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 ownerKey, uint256 assetId) view returns(uint256 balance)
func (_DepositTestnet *DepositTestnetCallerSession) GetWithdrawalBalance(ownerKey *big.Int, assetId *big.Int) (*big.Int, error) {
	return _DepositTestnet.Contract.GetWithdrawalBalance(&_DepositTestnet.CallOpts, ownerKey, assetId)
}

// IsAssetRegistered is a free data retrieval call binding the contract method 0x049f5ade.
//
// Solidity: function isAssetRegistered(uint256 assetType) view returns(bool)
func (_DepositTestnet *DepositTestnetCaller) IsAssetRegistered(opts *bind.CallOpts, assetType *big.Int) (bool, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "isAssetRegistered", assetType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssetRegistered is a free data retrieval call binding the contract method 0x049f5ade.
//
// Solidity: function isAssetRegistered(uint256 assetType) view returns(bool)
func (_DepositTestnet *DepositTestnetSession) IsAssetRegistered(assetType *big.Int) (bool, error) {
	return _DepositTestnet.Contract.IsAssetRegistered(&_DepositTestnet.CallOpts, assetType)
}

// IsAssetRegistered is a free data retrieval call binding the contract method 0x049f5ade.
//
// Solidity: function isAssetRegistered(uint256 assetType) view returns(bool)
func (_DepositTestnet *DepositTestnetCallerSession) IsAssetRegistered(assetType *big.Int) (bool, error) {
	return _DepositTestnet.Contract.IsAssetRegistered(&_DepositTestnet.CallOpts, assetType)
}

// IsAvailabilityVerifier is a free data retrieval call binding the contract method 0xbd1279ae.
//
// Solidity: function isAvailabilityVerifier(address verifierAddress) view returns(bool)
func (_DepositTestnet *DepositTestnetCaller) IsAvailabilityVerifier(opts *bind.CallOpts, verifierAddress common.Address) (bool, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "isAvailabilityVerifier", verifierAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAvailabilityVerifier is a free data retrieval call binding the contract method 0xbd1279ae.
//
// Solidity: function isAvailabilityVerifier(address verifierAddress) view returns(bool)
func (_DepositTestnet *DepositTestnetSession) IsAvailabilityVerifier(verifierAddress common.Address) (bool, error) {
	return _DepositTestnet.Contract.IsAvailabilityVerifier(&_DepositTestnet.CallOpts, verifierAddress)
}

// IsAvailabilityVerifier is a free data retrieval call binding the contract method 0xbd1279ae.
//
// Solidity: function isAvailabilityVerifier(address verifierAddress) view returns(bool)
func (_DepositTestnet *DepositTestnetCallerSession) IsAvailabilityVerifier(verifierAddress common.Address) (bool, error) {
	return _DepositTestnet.Contract.IsAvailabilityVerifier(&_DepositTestnet.CallOpts, verifierAddress)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_DepositTestnet *DepositTestnetCaller) IsFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "isFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_DepositTestnet *DepositTestnetSession) IsFrozen() (bool, error) {
	return _DepositTestnet.Contract.IsFrozen(&_DepositTestnet.CallOpts)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_DepositTestnet *DepositTestnetCallerSession) IsFrozen() (bool, error) {
	return _DepositTestnet.Contract.IsFrozen(&_DepositTestnet.CallOpts)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address testedOperator) view returns(bool)
func (_DepositTestnet *DepositTestnetCaller) IsOperator(opts *bind.CallOpts, testedOperator common.Address) (bool, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "isOperator", testedOperator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address testedOperator) view returns(bool)
func (_DepositTestnet *DepositTestnetSession) IsOperator(testedOperator common.Address) (bool, error) {
	return _DepositTestnet.Contract.IsOperator(&_DepositTestnet.CallOpts, testedOperator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address testedOperator) view returns(bool)
func (_DepositTestnet *DepositTestnetCallerSession) IsOperator(testedOperator common.Address) (bool, error) {
	return _DepositTestnet.Contract.IsOperator(&_DepositTestnet.CallOpts, testedOperator)
}

// IsStrictVaultBalancePolicy is a free data retrieval call binding the contract method 0x5adf1639.
//
// Solidity: function isStrictVaultBalancePolicy() view returns(bool)
func (_DepositTestnet *DepositTestnetCaller) IsStrictVaultBalancePolicy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "isStrictVaultBalancePolicy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStrictVaultBalancePolicy is a free data retrieval call binding the contract method 0x5adf1639.
//
// Solidity: function isStrictVaultBalancePolicy() view returns(bool)
func (_DepositTestnet *DepositTestnetSession) IsStrictVaultBalancePolicy() (bool, error) {
	return _DepositTestnet.Contract.IsStrictVaultBalancePolicy(&_DepositTestnet.CallOpts)
}

// IsStrictVaultBalancePolicy is a free data retrieval call binding the contract method 0x5adf1639.
//
// Solidity: function isStrictVaultBalancePolicy() view returns(bool)
func (_DepositTestnet *DepositTestnetCallerSession) IsStrictVaultBalancePolicy() (bool, error) {
	return _DepositTestnet.Contract.IsStrictVaultBalancePolicy(&_DepositTestnet.CallOpts)
}

// IsTokenAdmin is a free data retrieval call binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address testedAdmin) view returns(bool)
func (_DepositTestnet *DepositTestnetCaller) IsTokenAdmin(opts *bind.CallOpts, testedAdmin common.Address) (bool, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "isTokenAdmin", testedAdmin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenAdmin is a free data retrieval call binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address testedAdmin) view returns(bool)
func (_DepositTestnet *DepositTestnetSession) IsTokenAdmin(testedAdmin common.Address) (bool, error) {
	return _DepositTestnet.Contract.IsTokenAdmin(&_DepositTestnet.CallOpts, testedAdmin)
}

// IsTokenAdmin is a free data retrieval call binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address testedAdmin) view returns(bool)
func (_DepositTestnet *DepositTestnetCallerSession) IsTokenAdmin(testedAdmin common.Address) (bool, error) {
	return _DepositTestnet.Contract.IsTokenAdmin(&_DepositTestnet.CallOpts, testedAdmin)
}

// IsVaultLocked is a free data retrieval call binding the contract method 0x27c2b36d.
//
// Solidity: function isVaultLocked(address ethKey, uint256 assetId, uint256 vaultId) view returns(bool)
func (_DepositTestnet *DepositTestnetCaller) IsVaultLocked(opts *bind.CallOpts, ethKey common.Address, assetId *big.Int, vaultId *big.Int) (bool, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "isVaultLocked", ethKey, assetId, vaultId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVaultLocked is a free data retrieval call binding the contract method 0x27c2b36d.
//
// Solidity: function isVaultLocked(address ethKey, uint256 assetId, uint256 vaultId) view returns(bool)
func (_DepositTestnet *DepositTestnetSession) IsVaultLocked(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (bool, error) {
	return _DepositTestnet.Contract.IsVaultLocked(&_DepositTestnet.CallOpts, ethKey, assetId, vaultId)
}

// IsVaultLocked is a free data retrieval call binding the contract method 0x27c2b36d.
//
// Solidity: function isVaultLocked(address ethKey, uint256 assetId, uint256 vaultId) view returns(bool)
func (_DepositTestnet *DepositTestnetCallerSession) IsVaultLocked(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (bool, error) {
	return _DepositTestnet.Contract.IsVaultLocked(&_DepositTestnet.CallOpts, ethKey, assetId, vaultId)
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address verifierAddress) view returns(bool)
func (_DepositTestnet *DepositTestnetCaller) IsVerifier(opts *bind.CallOpts, verifierAddress common.Address) (bool, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "isVerifier", verifierAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address verifierAddress) view returns(bool)
func (_DepositTestnet *DepositTestnetSession) IsVerifier(verifierAddress common.Address) (bool, error) {
	return _DepositTestnet.Contract.IsVerifier(&_DepositTestnet.CallOpts, verifierAddress)
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address verifierAddress) view returns(bool)
func (_DepositTestnet *DepositTestnetCallerSession) IsVerifier(verifierAddress common.Address) (bool, error) {
	return _DepositTestnet.Contract.IsVerifier(&_DepositTestnet.CallOpts, verifierAddress)
}

// MainIsGovernor is a free data retrieval call binding the contract method 0x45f5cd97.
//
// Solidity: function mainIsGovernor(address testGovernor) view returns(bool)
func (_DepositTestnet *DepositTestnetCaller) MainIsGovernor(opts *bind.CallOpts, testGovernor common.Address) (bool, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "mainIsGovernor", testGovernor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MainIsGovernor is a free data retrieval call binding the contract method 0x45f5cd97.
//
// Solidity: function mainIsGovernor(address testGovernor) view returns(bool)
func (_DepositTestnet *DepositTestnetSession) MainIsGovernor(testGovernor common.Address) (bool, error) {
	return _DepositTestnet.Contract.MainIsGovernor(&_DepositTestnet.CallOpts, testGovernor)
}

// MainIsGovernor is a free data retrieval call binding the contract method 0x45f5cd97.
//
// Solidity: function mainIsGovernor(address testGovernor) view returns(bool)
func (_DepositTestnet *DepositTestnetCallerSession) MainIsGovernor(testGovernor common.Address) (bool, error) {
	return _DepositTestnet.Contract.MainIsGovernor(&_DepositTestnet.CallOpts, testGovernor)
}

// OrderRegistryAddress is a free data retrieval call binding the contract method 0x9c6a2837.
//
// Solidity: function orderRegistryAddress() view returns(address)
func (_DepositTestnet *DepositTestnetCaller) OrderRegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DepositTestnet.contract.Call(opts, &out, "orderRegistryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrderRegistryAddress is a free data retrieval call binding the contract method 0x9c6a2837.
//
// Solidity: function orderRegistryAddress() view returns(address)
func (_DepositTestnet *DepositTestnetSession) OrderRegistryAddress() (common.Address, error) {
	return _DepositTestnet.Contract.OrderRegistryAddress(&_DepositTestnet.CallOpts)
}

// OrderRegistryAddress is a free data retrieval call binding the contract method 0x9c6a2837.
//
// Solidity: function orderRegistryAddress() view returns(address)
func (_DepositTestnet *DepositTestnetCallerSession) OrderRegistryAddress() (common.Address, error) {
	return _DepositTestnet.Contract.OrderRegistryAddress(&_DepositTestnet.CallOpts)
}

// AnnounceAvailabilityVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x1d078bbb.
//
// Solidity: function announceAvailabilityVerifierRemovalIntent(address verifier) returns()
func (_DepositTestnet *DepositTestnetTransactor) AnnounceAvailabilityVerifierRemovalIntent(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "announceAvailabilityVerifierRemovalIntent", verifier)
}

// AnnounceAvailabilityVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x1d078bbb.
//
// Solidity: function announceAvailabilityVerifierRemovalIntent(address verifier) returns()
func (_DepositTestnet *DepositTestnetSession) AnnounceAvailabilityVerifierRemovalIntent(verifier common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.AnnounceAvailabilityVerifierRemovalIntent(&_DepositTestnet.TransactOpts, verifier)
}

// AnnounceAvailabilityVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x1d078bbb.
//
// Solidity: function announceAvailabilityVerifierRemovalIntent(address verifier) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) AnnounceAvailabilityVerifierRemovalIntent(verifier common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.AnnounceAvailabilityVerifierRemovalIntent(&_DepositTestnet.TransactOpts, verifier)
}

// AnnounceVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x418573b7.
//
// Solidity: function announceVerifierRemovalIntent(address verifier) returns()
func (_DepositTestnet *DepositTestnetTransactor) AnnounceVerifierRemovalIntent(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "announceVerifierRemovalIntent", verifier)
}

// AnnounceVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x418573b7.
//
// Solidity: function announceVerifierRemovalIntent(address verifier) returns()
func (_DepositTestnet *DepositTestnetSession) AnnounceVerifierRemovalIntent(verifier common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.AnnounceVerifierRemovalIntent(&_DepositTestnet.TransactOpts, verifier)
}

// AnnounceVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x418573b7.
//
// Solidity: function announceVerifierRemovalIntent(address verifier) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) AnnounceVerifierRemovalIntent(verifier common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.AnnounceVerifierRemovalIntent(&_DepositTestnet.TransactOpts, verifier)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_DepositTestnet *DepositTestnetTransactor) Deposit(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "deposit", starkKey, assetType, vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_DepositTestnet *DepositTestnetSession) Deposit(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.Deposit(&_DepositTestnet.TransactOpts, starkKey, assetType, vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_DepositTestnet *DepositTestnetTransactorSession) Deposit(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.Deposit(&_DepositTestnet.TransactOpts, starkKey, assetType, vaultId)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetTransactor) Deposit0(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "deposit0", starkKey, assetType, vaultId, quantizedAmount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetSession) Deposit0(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.Deposit0(&_DepositTestnet.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) Deposit0(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.Deposit0(&_DepositTestnet.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_DepositTestnet *DepositTestnetTransactor) DepositCancel(opts *bind.TransactOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "depositCancel", starkKey, assetId, vaultId)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_DepositTestnet *DepositTestnetSession) DepositCancel(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositCancel(&_DepositTestnet.TransactOpts, starkKey, assetId, vaultId)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) DepositCancel(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositCancel(&_DepositTestnet.TransactOpts, starkKey, assetId, vaultId)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x9ed17084.
//
// Solidity: function depositERC20(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetTransactor) DepositERC20(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "depositERC20", starkKey, assetType, vaultId, quantizedAmount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x9ed17084.
//
// Solidity: function depositERC20(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetSession) DepositERC20(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositERC20(&_DepositTestnet.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x9ed17084.
//
// Solidity: function depositERC20(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) DepositERC20(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositERC20(&_DepositTestnet.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// DepositERC20ToVault is a paid mutator transaction binding the contract method 0x1d133002.
//
// Solidity: function depositERC20ToVault(uint256 assetId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetTransactor) DepositERC20ToVault(opts *bind.TransactOpts, assetId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "depositERC20ToVault", assetId, vaultId, quantizedAmount)
}

// DepositERC20ToVault is a paid mutator transaction binding the contract method 0x1d133002.
//
// Solidity: function depositERC20ToVault(uint256 assetId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetSession) DepositERC20ToVault(assetId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositERC20ToVault(&_DepositTestnet.TransactOpts, assetId, vaultId, quantizedAmount)
}

// DepositERC20ToVault is a paid mutator transaction binding the contract method 0x1d133002.
//
// Solidity: function depositERC20ToVault(uint256 assetId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) DepositERC20ToVault(assetId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositERC20ToVault(&_DepositTestnet.TransactOpts, assetId, vaultId, quantizedAmount)
}

// DepositEth is a paid mutator transaction binding the contract method 0x6ce5d957.
//
// Solidity: function depositEth(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_DepositTestnet *DepositTestnetTransactor) DepositEth(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "depositEth", starkKey, assetType, vaultId)
}

// DepositEth is a paid mutator transaction binding the contract method 0x6ce5d957.
//
// Solidity: function depositEth(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_DepositTestnet *DepositTestnetSession) DepositEth(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositEth(&_DepositTestnet.TransactOpts, starkKey, assetType, vaultId)
}

// DepositEth is a paid mutator transaction binding the contract method 0x6ce5d957.
//
// Solidity: function depositEth(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_DepositTestnet *DepositTestnetTransactorSession) DepositEth(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositEth(&_DepositTestnet.TransactOpts, starkKey, assetType, vaultId)
}

// DepositEthToVault is a paid mutator transaction binding the contract method 0xa3c71fde.
//
// Solidity: function depositEthToVault(uint256 assetId, uint256 vaultId) payable returns()
func (_DepositTestnet *DepositTestnetTransactor) DepositEthToVault(opts *bind.TransactOpts, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "depositEthToVault", assetId, vaultId)
}

// DepositEthToVault is a paid mutator transaction binding the contract method 0xa3c71fde.
//
// Solidity: function depositEthToVault(uint256 assetId, uint256 vaultId) payable returns()
func (_DepositTestnet *DepositTestnetSession) DepositEthToVault(assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositEthToVault(&_DepositTestnet.TransactOpts, assetId, vaultId)
}

// DepositEthToVault is a paid mutator transaction binding the contract method 0xa3c71fde.
//
// Solidity: function depositEthToVault(uint256 assetId, uint256 vaultId) payable returns()
func (_DepositTestnet *DepositTestnetTransactorSession) DepositEthToVault(assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositEthToVault(&_DepositTestnet.TransactOpts, assetId, vaultId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_DepositTestnet *DepositTestnetTransactor) DepositNft(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "depositNft", starkKey, assetType, vaultId, tokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_DepositTestnet *DepositTestnetSession) DepositNft(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositNft(&_DepositTestnet.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) DepositNft(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositNft(&_DepositTestnet.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_DepositTestnet *DepositTestnetTransactor) DepositNftReclaim(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "depositNftReclaim", starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_DepositTestnet *DepositTestnetSession) DepositNftReclaim(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositNftReclaim(&_DepositTestnet.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) DepositNftReclaim(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositNftReclaim(&_DepositTestnet.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_DepositTestnet *DepositTestnetTransactor) DepositReclaim(opts *bind.TransactOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "depositReclaim", starkKey, assetId, vaultId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_DepositTestnet *DepositTestnetSession) DepositReclaim(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositReclaim(&_DepositTestnet.TransactOpts, starkKey, assetId, vaultId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) DepositReclaim(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.DepositReclaim(&_DepositTestnet.TransactOpts, starkKey, assetId, vaultId)
}

// Escape is a paid mutator transaction binding the contract method 0x9e3adac4.
//
// Solidity: function escape(uint256 starkKey, uint256 vaultId, uint256 assetId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetTransactor) Escape(opts *bind.TransactOpts, starkKey *big.Int, vaultId *big.Int, assetId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "escape", starkKey, vaultId, assetId, quantizedAmount)
}

// Escape is a paid mutator transaction binding the contract method 0x9e3adac4.
//
// Solidity: function escape(uint256 starkKey, uint256 vaultId, uint256 assetId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetSession) Escape(starkKey *big.Int, vaultId *big.Int, assetId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.Escape(&_DepositTestnet.TransactOpts, starkKey, vaultId, assetId, quantizedAmount)
}

// Escape is a paid mutator transaction binding the contract method 0x9e3adac4.
//
// Solidity: function escape(uint256 starkKey, uint256 vaultId, uint256 assetId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) Escape(starkKey *big.Int, vaultId *big.Int, assetId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.Escape(&_DepositTestnet.TransactOpts, starkKey, vaultId, assetId, quantizedAmount)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x93c1e466.
//
// Solidity: function freezeRequest(uint256 starkKey, uint256 vaultId) returns()
func (_DepositTestnet *DepositTestnetTransactor) FreezeRequest(opts *bind.TransactOpts, starkKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "freezeRequest", starkKey, vaultId)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x93c1e466.
//
// Solidity: function freezeRequest(uint256 starkKey, uint256 vaultId) returns()
func (_DepositTestnet *DepositTestnetSession) FreezeRequest(starkKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.FreezeRequest(&_DepositTestnet.TransactOpts, starkKey, vaultId)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x93c1e466.
//
// Solidity: function freezeRequest(uint256 starkKey, uint256 vaultId) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) FreezeRequest(starkKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.FreezeRequest(&_DepositTestnet.TransactOpts, starkKey, vaultId)
}

// FullWithdrawalRequest is a paid mutator transaction binding the contract method 0xa93310c4.
//
// Solidity: function fullWithdrawalRequest(uint256 starkKey, uint256 vaultId) returns()
func (_DepositTestnet *DepositTestnetTransactor) FullWithdrawalRequest(opts *bind.TransactOpts, starkKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "fullWithdrawalRequest", starkKey, vaultId)
}

// FullWithdrawalRequest is a paid mutator transaction binding the contract method 0xa93310c4.
//
// Solidity: function fullWithdrawalRequest(uint256 starkKey, uint256 vaultId) returns()
func (_DepositTestnet *DepositTestnetSession) FullWithdrawalRequest(starkKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.FullWithdrawalRequest(&_DepositTestnet.TransactOpts, starkKey, vaultId)
}

// FullWithdrawalRequest is a paid mutator transaction binding the contract method 0xa93310c4.
//
// Solidity: function fullWithdrawalRequest(uint256 starkKey, uint256 vaultId) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) FullWithdrawalRequest(starkKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.FullWithdrawalRequest(&_DepositTestnet.TransactOpts, starkKey, vaultId)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_DepositTestnet *DepositTestnetTransactor) Initialize(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "initialize", data)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_DepositTestnet *DepositTestnetSession) Initialize(data []byte) (*types.Transaction, error) {
	return _DepositTestnet.Contract.Initialize(&_DepositTestnet.TransactOpts, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) Initialize(data []byte) (*types.Transaction, error) {
	return _DepositTestnet.Contract.Initialize(&_DepositTestnet.TransactOpts, data)
}

// LockVault is a paid mutator transaction binding the contract method 0xe8d28a9d.
//
// Solidity: function lockVault(uint256 assetId, uint256 vaultId, uint256 lockTime) returns()
func (_DepositTestnet *DepositTestnetTransactor) LockVault(opts *bind.TransactOpts, assetId *big.Int, vaultId *big.Int, lockTime *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "lockVault", assetId, vaultId, lockTime)
}

// LockVault is a paid mutator transaction binding the contract method 0xe8d28a9d.
//
// Solidity: function lockVault(uint256 assetId, uint256 vaultId, uint256 lockTime) returns()
func (_DepositTestnet *DepositTestnetSession) LockVault(assetId *big.Int, vaultId *big.Int, lockTime *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.LockVault(&_DepositTestnet.TransactOpts, assetId, vaultId, lockTime)
}

// LockVault is a paid mutator transaction binding the contract method 0xe8d28a9d.
//
// Solidity: function lockVault(uint256 assetId, uint256 vaultId, uint256 lockTime) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) LockVault(assetId *big.Int, vaultId *big.Int, lockTime *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.LockVault(&_DepositTestnet.TransactOpts, assetId, vaultId, lockTime)
}

// MainAcceptGovernance is a paid mutator transaction binding the contract method 0x28700a15.
//
// Solidity: function mainAcceptGovernance() returns()
func (_DepositTestnet *DepositTestnetTransactor) MainAcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "mainAcceptGovernance")
}

// MainAcceptGovernance is a paid mutator transaction binding the contract method 0x28700a15.
//
// Solidity: function mainAcceptGovernance() returns()
func (_DepositTestnet *DepositTestnetSession) MainAcceptGovernance() (*types.Transaction, error) {
	return _DepositTestnet.Contract.MainAcceptGovernance(&_DepositTestnet.TransactOpts)
}

// MainAcceptGovernance is a paid mutator transaction binding the contract method 0x28700a15.
//
// Solidity: function mainAcceptGovernance() returns()
func (_DepositTestnet *DepositTestnetTransactorSession) MainAcceptGovernance() (*types.Transaction, error) {
	return _DepositTestnet.Contract.MainAcceptGovernance(&_DepositTestnet.TransactOpts)
}

// MainCancelNomination is a paid mutator transaction binding the contract method 0x72eb3688.
//
// Solidity: function mainCancelNomination() returns()
func (_DepositTestnet *DepositTestnetTransactor) MainCancelNomination(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "mainCancelNomination")
}

// MainCancelNomination is a paid mutator transaction binding the contract method 0x72eb3688.
//
// Solidity: function mainCancelNomination() returns()
func (_DepositTestnet *DepositTestnetSession) MainCancelNomination() (*types.Transaction, error) {
	return _DepositTestnet.Contract.MainCancelNomination(&_DepositTestnet.TransactOpts)
}

// MainCancelNomination is a paid mutator transaction binding the contract method 0x72eb3688.
//
// Solidity: function mainCancelNomination() returns()
func (_DepositTestnet *DepositTestnetTransactorSession) MainCancelNomination() (*types.Transaction, error) {
	return _DepositTestnet.Contract.MainCancelNomination(&_DepositTestnet.TransactOpts)
}

// MainNominateNewGovernor is a paid mutator transaction binding the contract method 0x8c4bce1c.
//
// Solidity: function mainNominateNewGovernor(address newGovernor) returns()
func (_DepositTestnet *DepositTestnetTransactor) MainNominateNewGovernor(opts *bind.TransactOpts, newGovernor common.Address) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "mainNominateNewGovernor", newGovernor)
}

// MainNominateNewGovernor is a paid mutator transaction binding the contract method 0x8c4bce1c.
//
// Solidity: function mainNominateNewGovernor(address newGovernor) returns()
func (_DepositTestnet *DepositTestnetSession) MainNominateNewGovernor(newGovernor common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.MainNominateNewGovernor(&_DepositTestnet.TransactOpts, newGovernor)
}

// MainNominateNewGovernor is a paid mutator transaction binding the contract method 0x8c4bce1c.
//
// Solidity: function mainNominateNewGovernor(address newGovernor) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) MainNominateNewGovernor(newGovernor common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.MainNominateNewGovernor(&_DepositTestnet.TransactOpts, newGovernor)
}

// MainRemoveGovernor is a paid mutator transaction binding the contract method 0xa1cc921e.
//
// Solidity: function mainRemoveGovernor(address governorForRemoval) returns()
func (_DepositTestnet *DepositTestnetTransactor) MainRemoveGovernor(opts *bind.TransactOpts, governorForRemoval common.Address) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "mainRemoveGovernor", governorForRemoval)
}

// MainRemoveGovernor is a paid mutator transaction binding the contract method 0xa1cc921e.
//
// Solidity: function mainRemoveGovernor(address governorForRemoval) returns()
func (_DepositTestnet *DepositTestnetSession) MainRemoveGovernor(governorForRemoval common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.MainRemoveGovernor(&_DepositTestnet.TransactOpts, governorForRemoval)
}

// MainRemoveGovernor is a paid mutator transaction binding the contract method 0xa1cc921e.
//
// Solidity: function mainRemoveGovernor(address governorForRemoval) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) MainRemoveGovernor(governorForRemoval common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.MainRemoveGovernor(&_DepositTestnet.TransactOpts, governorForRemoval)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DepositTestnet *DepositTestnetTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DepositTestnet *DepositTestnetSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DepositTestnet.Contract.OnERC721Received(&_DepositTestnet.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DepositTestnet *DepositTestnetTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DepositTestnet.Contract.OnERC721Received(&_DepositTestnet.TransactOpts, arg0, arg1, arg2, arg3)
}

// RegisterAndDepositERC20 is a paid mutator transaction binding the contract method 0xd5280589.
//
// Solidity: function registerAndDepositERC20(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetTransactor) RegisterAndDepositERC20(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "registerAndDepositERC20", ethKey, starkKey, signature, assetType, vaultId, quantizedAmount)
}

// RegisterAndDepositERC20 is a paid mutator transaction binding the contract method 0xd5280589.
//
// Solidity: function registerAndDepositERC20(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetSession) RegisterAndDepositERC20(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterAndDepositERC20(&_DepositTestnet.TransactOpts, ethKey, starkKey, signature, assetType, vaultId, quantizedAmount)
}

// RegisterAndDepositERC20 is a paid mutator transaction binding the contract method 0xd5280589.
//
// Solidity: function registerAndDepositERC20(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) RegisterAndDepositERC20(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterAndDepositERC20(&_DepositTestnet.TransactOpts, ethKey, starkKey, signature, assetType, vaultId, quantizedAmount)
}

// RegisterAndDepositEth is a paid mutator transaction binding the contract method 0x3ccfc8ed.
//
// Solidity: function registerAndDepositEth(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId) payable returns()
func (_DepositTestnet *DepositTestnetTransactor) RegisterAndDepositEth(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "registerAndDepositEth", ethKey, starkKey, signature, assetType, vaultId)
}

// RegisterAndDepositEth is a paid mutator transaction binding the contract method 0x3ccfc8ed.
//
// Solidity: function registerAndDepositEth(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId) payable returns()
func (_DepositTestnet *DepositTestnetSession) RegisterAndDepositEth(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterAndDepositEth(&_DepositTestnet.TransactOpts, ethKey, starkKey, signature, assetType, vaultId)
}

// RegisterAndDepositEth is a paid mutator transaction binding the contract method 0x3ccfc8ed.
//
// Solidity: function registerAndDepositEth(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId) payable returns()
func (_DepositTestnet *DepositTestnetTransactorSession) RegisterAndDepositEth(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterAndDepositEth(&_DepositTestnet.TransactOpts, ethKey, starkKey, signature, assetType, vaultId)
}

// RegisterAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbdb75785.
//
// Solidity: function registerAvailabilityVerifier(address verifier, string identifier) returns()
func (_DepositTestnet *DepositTestnetTransactor) RegisterAvailabilityVerifier(opts *bind.TransactOpts, verifier common.Address, identifier string) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "registerAvailabilityVerifier", verifier, identifier)
}

// RegisterAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbdb75785.
//
// Solidity: function registerAvailabilityVerifier(address verifier, string identifier) returns()
func (_DepositTestnet *DepositTestnetSession) RegisterAvailabilityVerifier(verifier common.Address, identifier string) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterAvailabilityVerifier(&_DepositTestnet.TransactOpts, verifier, identifier)
}

// RegisterAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbdb75785.
//
// Solidity: function registerAvailabilityVerifier(address verifier, string identifier) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) RegisterAvailabilityVerifier(verifier common.Address, identifier string) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterAvailabilityVerifier(&_DepositTestnet.TransactOpts, verifier, identifier)
}

// RegisterEthAddress is a paid mutator transaction binding the contract method 0xbea84187.
//
// Solidity: function registerEthAddress(address ethKey, uint256 starkKey, bytes starkSignature) returns()
func (_DepositTestnet *DepositTestnetTransactor) RegisterEthAddress(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "registerEthAddress", ethKey, starkKey, starkSignature)
}

// RegisterEthAddress is a paid mutator transaction binding the contract method 0xbea84187.
//
// Solidity: function registerEthAddress(address ethKey, uint256 starkKey, bytes starkSignature) returns()
func (_DepositTestnet *DepositTestnetSession) RegisterEthAddress(ethKey common.Address, starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterEthAddress(&_DepositTestnet.TransactOpts, ethKey, starkKey, starkSignature)
}

// RegisterEthAddress is a paid mutator transaction binding the contract method 0xbea84187.
//
// Solidity: function registerEthAddress(address ethKey, uint256 starkKey, bytes starkSignature) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) RegisterEthAddress(ethKey common.Address, starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterEthAddress(&_DepositTestnet.TransactOpts, ethKey, starkKey, starkSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_DepositTestnet *DepositTestnetTransactor) RegisterOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "registerOperator", newOperator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_DepositTestnet *DepositTestnetSession) RegisterOperator(newOperator common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterOperator(&_DepositTestnet.TransactOpts, newOperator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) RegisterOperator(newOperator common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterOperator(&_DepositTestnet.TransactOpts, newOperator)
}

// RegisterSender is a paid mutator transaction binding the contract method 0x86aeb445.
//
// Solidity: function registerSender(uint256 starkKey, bytes starkSignature) returns()
func (_DepositTestnet *DepositTestnetTransactor) RegisterSender(opts *bind.TransactOpts, starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "registerSender", starkKey, starkSignature)
}

// RegisterSender is a paid mutator transaction binding the contract method 0x86aeb445.
//
// Solidity: function registerSender(uint256 starkKey, bytes starkSignature) returns()
func (_DepositTestnet *DepositTestnetSession) RegisterSender(starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterSender(&_DepositTestnet.TransactOpts, starkKey, starkSignature)
}

// RegisterSender is a paid mutator transaction binding the contract method 0x86aeb445.
//
// Solidity: function registerSender(uint256 starkKey, bytes starkSignature) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) RegisterSender(starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterSender(&_DepositTestnet.TransactOpts, starkKey, starkSignature)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo) returns()
func (_DepositTestnet *DepositTestnetTransactor) RegisterToken(opts *bind.TransactOpts, assetType *big.Int, assetInfo []byte) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "registerToken", assetType, assetInfo)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo) returns()
func (_DepositTestnet *DepositTestnetSession) RegisterToken(assetType *big.Int, assetInfo []byte) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterToken(&_DepositTestnet.TransactOpts, assetType, assetInfo)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) RegisterToken(assetType *big.Int, assetInfo []byte) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterToken(&_DepositTestnet.TransactOpts, assetType, assetInfo)
}

// RegisterToken0 is a paid mutator transaction binding the contract method 0xd88d8b38.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo, uint256 quantum) returns()
func (_DepositTestnet *DepositTestnetTransactor) RegisterToken0(opts *bind.TransactOpts, assetType *big.Int, assetInfo []byte, quantum *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "registerToken0", assetType, assetInfo, quantum)
}

// RegisterToken0 is a paid mutator transaction binding the contract method 0xd88d8b38.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo, uint256 quantum) returns()
func (_DepositTestnet *DepositTestnetSession) RegisterToken0(assetType *big.Int, assetInfo []byte, quantum *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterToken0(&_DepositTestnet.TransactOpts, assetType, assetInfo, quantum)
}

// RegisterToken0 is a paid mutator transaction binding the contract method 0xd88d8b38.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo, uint256 quantum) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) RegisterToken0(assetType *big.Int, assetInfo []byte, quantum *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterToken0(&_DepositTestnet.TransactOpts, assetType, assetInfo, quantum)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address newAdmin) returns()
func (_DepositTestnet *DepositTestnetTransactor) RegisterTokenAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "registerTokenAdmin", newAdmin)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address newAdmin) returns()
func (_DepositTestnet *DepositTestnetSession) RegisterTokenAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterTokenAdmin(&_DepositTestnet.TransactOpts, newAdmin)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address newAdmin) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) RegisterTokenAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterTokenAdmin(&_DepositTestnet.TransactOpts, newAdmin)
}

// RegisterVerifier is a paid mutator transaction binding the contract method 0x3776fe2a.
//
// Solidity: function registerVerifier(address verifier, string identifier) returns()
func (_DepositTestnet *DepositTestnetTransactor) RegisterVerifier(opts *bind.TransactOpts, verifier common.Address, identifier string) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "registerVerifier", verifier, identifier)
}

// RegisterVerifier is a paid mutator transaction binding the contract method 0x3776fe2a.
//
// Solidity: function registerVerifier(address verifier, string identifier) returns()
func (_DepositTestnet *DepositTestnetSession) RegisterVerifier(verifier common.Address, identifier string) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterVerifier(&_DepositTestnet.TransactOpts, verifier, identifier)
}

// RegisterVerifier is a paid mutator transaction binding the contract method 0x3776fe2a.
//
// Solidity: function registerVerifier(address verifier, string identifier) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) RegisterVerifier(verifier common.Address, identifier string) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RegisterVerifier(&_DepositTestnet.TransactOpts, verifier, identifier)
}

// RemoveAvailabilityVerifier is a paid mutator transaction binding the contract method 0xb1e640bf.
//
// Solidity: function removeAvailabilityVerifier(address verifier) returns()
func (_DepositTestnet *DepositTestnetTransactor) RemoveAvailabilityVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "removeAvailabilityVerifier", verifier)
}

// RemoveAvailabilityVerifier is a paid mutator transaction binding the contract method 0xb1e640bf.
//
// Solidity: function removeAvailabilityVerifier(address verifier) returns()
func (_DepositTestnet *DepositTestnetSession) RemoveAvailabilityVerifier(verifier common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RemoveAvailabilityVerifier(&_DepositTestnet.TransactOpts, verifier)
}

// RemoveAvailabilityVerifier is a paid mutator transaction binding the contract method 0xb1e640bf.
//
// Solidity: function removeAvailabilityVerifier(address verifier) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) RemoveAvailabilityVerifier(verifier common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RemoveAvailabilityVerifier(&_DepositTestnet.TransactOpts, verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_DepositTestnet *DepositTestnetTransactor) RemoveVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "removeVerifier", verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_DepositTestnet *DepositTestnetSession) RemoveVerifier(verifier common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RemoveVerifier(&_DepositTestnet.TransactOpts, verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) RemoveVerifier(verifier common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.RemoveVerifier(&_DepositTestnet.TransactOpts, verifier)
}

// SetDefaultVaultWithdrawalLock is a paid mutator transaction binding the contract method 0xd4181dea.
//
// Solidity: function setDefaultVaultWithdrawalLock(uint256 newDefaultTime) returns()
func (_DepositTestnet *DepositTestnetTransactor) SetDefaultVaultWithdrawalLock(opts *bind.TransactOpts, newDefaultTime *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "setDefaultVaultWithdrawalLock", newDefaultTime)
}

// SetDefaultVaultWithdrawalLock is a paid mutator transaction binding the contract method 0xd4181dea.
//
// Solidity: function setDefaultVaultWithdrawalLock(uint256 newDefaultTime) returns()
func (_DepositTestnet *DepositTestnetSession) SetDefaultVaultWithdrawalLock(newDefaultTime *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.SetDefaultVaultWithdrawalLock(&_DepositTestnet.TransactOpts, newDefaultTime)
}

// SetDefaultVaultWithdrawalLock is a paid mutator transaction binding the contract method 0xd4181dea.
//
// Solidity: function setDefaultVaultWithdrawalLock(uint256 newDefaultTime) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) SetDefaultVaultWithdrawalLock(newDefaultTime *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.SetDefaultVaultWithdrawalLock(&_DepositTestnet.TransactOpts, newDefaultTime)
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_DepositTestnet *DepositTestnetTransactor) UnFreeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "unFreeze")
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_DepositTestnet *DepositTestnetSession) UnFreeze() (*types.Transaction, error) {
	return _DepositTestnet.Contract.UnFreeze(&_DepositTestnet.TransactOpts)
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_DepositTestnet *DepositTestnetTransactorSession) UnFreeze() (*types.Transaction, error) {
	return _DepositTestnet.Contract.UnFreeze(&_DepositTestnet.TransactOpts)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_DepositTestnet *DepositTestnetTransactor) UnregisterOperator(opts *bind.TransactOpts, removedOperator common.Address) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "unregisterOperator", removedOperator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_DepositTestnet *DepositTestnetSession) UnregisterOperator(removedOperator common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.UnregisterOperator(&_DepositTestnet.TransactOpts, removedOperator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) UnregisterOperator(removedOperator common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.UnregisterOperator(&_DepositTestnet.TransactOpts, removedOperator)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address oldAdmin) returns()
func (_DepositTestnet *DepositTestnetTransactor) UnregisterTokenAdmin(opts *bind.TransactOpts, oldAdmin common.Address) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "unregisterTokenAdmin", oldAdmin)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address oldAdmin) returns()
func (_DepositTestnet *DepositTestnetSession) UnregisterTokenAdmin(oldAdmin common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.UnregisterTokenAdmin(&_DepositTestnet.TransactOpts, oldAdmin)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address oldAdmin) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) UnregisterTokenAdmin(oldAdmin common.Address) (*types.Transaction, error) {
	return _DepositTestnet.Contract.UnregisterTokenAdmin(&_DepositTestnet.TransactOpts, oldAdmin)
}

// UpdateImplementationActivationTime is a paid mutator transaction binding the contract method 0x02a93fae.
//
// Solidity: function updateImplementationActivationTime(address implementation, bytes data, bool finalize) returns()
func (_DepositTestnet *DepositTestnetTransactor) UpdateImplementationActivationTime(opts *bind.TransactOpts, implementation common.Address, data []byte, finalize bool) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "updateImplementationActivationTime", implementation, data, finalize)
}

// UpdateImplementationActivationTime is a paid mutator transaction binding the contract method 0x02a93fae.
//
// Solidity: function updateImplementationActivationTime(address implementation, bytes data, bool finalize) returns()
func (_DepositTestnet *DepositTestnetSession) UpdateImplementationActivationTime(implementation common.Address, data []byte, finalize bool) (*types.Transaction, error) {
	return _DepositTestnet.Contract.UpdateImplementationActivationTime(&_DepositTestnet.TransactOpts, implementation, data, finalize)
}

// UpdateImplementationActivationTime is a paid mutator transaction binding the contract method 0x02a93fae.
//
// Solidity: function updateImplementationActivationTime(address implementation, bytes data, bool finalize) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) UpdateImplementationActivationTime(implementation common.Address, data []byte, finalize bool) (*types.Transaction, error) {
	return _DepositTestnet.Contract.UpdateImplementationActivationTime(&_DepositTestnet.TransactOpts, implementation, data, finalize)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] publicInput, uint256[] applicationData) returns()
func (_DepositTestnet *DepositTestnetTransactor) UpdateState(opts *bind.TransactOpts, publicInput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "updateState", publicInput, applicationData)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] publicInput, uint256[] applicationData) returns()
func (_DepositTestnet *DepositTestnetSession) UpdateState(publicInput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.UpdateState(&_DepositTestnet.TransactOpts, publicInput, applicationData)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] publicInput, uint256[] applicationData) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) UpdateState(publicInput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.UpdateState(&_DepositTestnet.TransactOpts, publicInput, applicationData)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 ownerKey, uint256 assetType) returns()
func (_DepositTestnet *DepositTestnetTransactor) Withdraw(opts *bind.TransactOpts, ownerKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "withdraw", ownerKey, assetType)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 ownerKey, uint256 assetType) returns()
func (_DepositTestnet *DepositTestnetSession) Withdraw(ownerKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.Withdraw(&_DepositTestnet.TransactOpts, ownerKey, assetType)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 ownerKey, uint256 assetType) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) Withdraw(ownerKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.Withdraw(&_DepositTestnet.TransactOpts, ownerKey, assetType)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 ownerKey, uint256 assetType, bytes mintingBlob) returns()
func (_DepositTestnet *DepositTestnetTransactor) WithdrawAndMint(opts *bind.TransactOpts, ownerKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "withdrawAndMint", ownerKey, assetType, mintingBlob)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 ownerKey, uint256 assetType, bytes mintingBlob) returns()
func (_DepositTestnet *DepositTestnetSession) WithdrawAndMint(ownerKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _DepositTestnet.Contract.WithdrawAndMint(&_DepositTestnet.TransactOpts, ownerKey, assetType, mintingBlob)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 ownerKey, uint256 assetType, bytes mintingBlob) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) WithdrawAndMint(ownerKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _DepositTestnet.Contract.WithdrawAndMint(&_DepositTestnet.TransactOpts, ownerKey, assetType, mintingBlob)
}

// WithdrawFromVault is a paid mutator transaction binding the contract method 0xbf422b2e.
//
// Solidity: function withdrawFromVault(uint256 assetId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetTransactor) WithdrawFromVault(opts *bind.TransactOpts, assetId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "withdrawFromVault", assetId, vaultId, quantizedAmount)
}

// WithdrawFromVault is a paid mutator transaction binding the contract method 0xbf422b2e.
//
// Solidity: function withdrawFromVault(uint256 assetId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetSession) WithdrawFromVault(assetId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.WithdrawFromVault(&_DepositTestnet.TransactOpts, assetId, vaultId, quantizedAmount)
}

// WithdrawFromVault is a paid mutator transaction binding the contract method 0xbf422b2e.
//
// Solidity: function withdrawFromVault(uint256 assetId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) WithdrawFromVault(assetId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.WithdrawFromVault(&_DepositTestnet.TransactOpts, assetId, vaultId, quantizedAmount)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 ownerKey, uint256 assetType, uint256 tokenId) returns()
func (_DepositTestnet *DepositTestnetTransactor) WithdrawNft(opts *bind.TransactOpts, ownerKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.contract.Transact(opts, "withdrawNft", ownerKey, assetType, tokenId)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 ownerKey, uint256 assetType, uint256 tokenId) returns()
func (_DepositTestnet *DepositTestnetSession) WithdrawNft(ownerKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.WithdrawNft(&_DepositTestnet.TransactOpts, ownerKey, assetType, tokenId)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 ownerKey, uint256 assetType, uint256 tokenId) returns()
func (_DepositTestnet *DepositTestnetTransactorSession) WithdrawNft(ownerKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositTestnet.Contract.WithdrawNft(&_DepositTestnet.TransactOpts, ownerKey, assetType, tokenId)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DepositTestnet *DepositTestnetTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _DepositTestnet.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DepositTestnet *DepositTestnetSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DepositTestnet.Contract.Fallback(&_DepositTestnet.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DepositTestnet *DepositTestnetTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DepositTestnet.Contract.Fallback(&_DepositTestnet.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DepositTestnet *DepositTestnetTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositTestnet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DepositTestnet *DepositTestnetSession) Receive() (*types.Transaction, error) {
	return _DepositTestnet.Contract.Receive(&_DepositTestnet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DepositTestnet *DepositTestnetTransactorSession) Receive() (*types.Transaction, error) {
	return _DepositTestnet.Contract.Receive(&_DepositTestnet.TransactOpts)
}

// DepositTestnetImplementationActivationRescheduledIterator is returned from FilterImplementationActivationRescheduled and is used to iterate over the raw logs and unpacked data for ImplementationActivationRescheduled events raised by the DepositTestnet contract.
type DepositTestnetImplementationActivationRescheduledIterator struct {
	Event *DepositTestnetImplementationActivationRescheduled // Event containing the contract specifics and raw log

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
func (it *DepositTestnetImplementationActivationRescheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetImplementationActivationRescheduled)
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
		it.Event = new(DepositTestnetImplementationActivationRescheduled)
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
func (it *DepositTestnetImplementationActivationRescheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetImplementationActivationRescheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetImplementationActivationRescheduled represents a ImplementationActivationRescheduled event raised by the DepositTestnet contract.
type DepositTestnetImplementationActivationRescheduled struct {
	Implementation        common.Address
	UpdatedActivationTime *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterImplementationActivationRescheduled is a free log retrieval operation binding the contract event 0xdda7b7d1f8141bd98b4378ee60e6231f89598ca02949a9d0550904dc96efeeb7.
//
// Solidity: event ImplementationActivationRescheduled(address indexed implementation, uint256 updatedActivationTime)
func (_DepositTestnet *DepositTestnetFilterer) FilterImplementationActivationRescheduled(opts *bind.FilterOpts, implementation []common.Address) (*DepositTestnetImplementationActivationRescheduledIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "ImplementationActivationRescheduled", implementationRule)
	if err != nil {
		return nil, err
	}
	return &DepositTestnetImplementationActivationRescheduledIterator{contract: _DepositTestnet.contract, event: "ImplementationActivationRescheduled", logs: logs, sub: sub}, nil
}

// WatchImplementationActivationRescheduled is a free log subscription operation binding the contract event 0xdda7b7d1f8141bd98b4378ee60e6231f89598ca02949a9d0550904dc96efeeb7.
//
// Solidity: event ImplementationActivationRescheduled(address indexed implementation, uint256 updatedActivationTime)
func (_DepositTestnet *DepositTestnetFilterer) WatchImplementationActivationRescheduled(opts *bind.WatchOpts, sink chan<- *DepositTestnetImplementationActivationRescheduled, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "ImplementationActivationRescheduled", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetImplementationActivationRescheduled)
				if err := _DepositTestnet.contract.UnpackLog(event, "ImplementationActivationRescheduled", log); err != nil {
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

// ParseImplementationActivationRescheduled is a log parse operation binding the contract event 0xdda7b7d1f8141bd98b4378ee60e6231f89598ca02949a9d0550904dc96efeeb7.
//
// Solidity: event ImplementationActivationRescheduled(address indexed implementation, uint256 updatedActivationTime)
func (_DepositTestnet *DepositTestnetFilterer) ParseImplementationActivationRescheduled(log types.Log) (*DepositTestnetImplementationActivationRescheduled, error) {
	event := new(DepositTestnetImplementationActivationRescheduled)
	if err := _DepositTestnet.contract.UnpackLog(event, "ImplementationActivationRescheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogDefaultVaultWithdrawalLockSetIterator is returned from FilterLogDefaultVaultWithdrawalLockSet and is used to iterate over the raw logs and unpacked data for LogDefaultVaultWithdrawalLockSet events raised by the DepositTestnet contract.
type DepositTestnetLogDefaultVaultWithdrawalLockSetIterator struct {
	Event *DepositTestnetLogDefaultVaultWithdrawalLockSet // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogDefaultVaultWithdrawalLockSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogDefaultVaultWithdrawalLockSet)
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
		it.Event = new(DepositTestnetLogDefaultVaultWithdrawalLockSet)
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
func (it *DepositTestnetLogDefaultVaultWithdrawalLockSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogDefaultVaultWithdrawalLockSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogDefaultVaultWithdrawalLockSet represents a LogDefaultVaultWithdrawalLockSet event raised by the DepositTestnet contract.
type DepositTestnetLogDefaultVaultWithdrawalLockSet struct {
	NewDefaultLockTime *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDefaultVaultWithdrawalLockSet is a free log retrieval operation binding the contract event 0x832169a4c3cea413f0041437fd118afbc4b33edbf6783da382128bca1e56b2d2.
//
// Solidity: event LogDefaultVaultWithdrawalLockSet(uint256 newDefaultLockTime)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogDefaultVaultWithdrawalLockSet(opts *bind.FilterOpts) (*DepositTestnetLogDefaultVaultWithdrawalLockSetIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogDefaultVaultWithdrawalLockSet")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogDefaultVaultWithdrawalLockSetIterator{contract: _DepositTestnet.contract, event: "LogDefaultVaultWithdrawalLockSet", logs: logs, sub: sub}, nil
}

// WatchLogDefaultVaultWithdrawalLockSet is a free log subscription operation binding the contract event 0x832169a4c3cea413f0041437fd118afbc4b33edbf6783da382128bca1e56b2d2.
//
// Solidity: event LogDefaultVaultWithdrawalLockSet(uint256 newDefaultLockTime)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogDefaultVaultWithdrawalLockSet(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogDefaultVaultWithdrawalLockSet) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogDefaultVaultWithdrawalLockSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogDefaultVaultWithdrawalLockSet)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogDefaultVaultWithdrawalLockSet", log); err != nil {
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

// ParseLogDefaultVaultWithdrawalLockSet is a log parse operation binding the contract event 0x832169a4c3cea413f0041437fd118afbc4b33edbf6783da382128bca1e56b2d2.
//
// Solidity: event LogDefaultVaultWithdrawalLockSet(uint256 newDefaultLockTime)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogDefaultVaultWithdrawalLockSet(log types.Log) (*DepositTestnetLogDefaultVaultWithdrawalLockSet, error) {
	event := new(DepositTestnetLogDefaultVaultWithdrawalLockSet)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogDefaultVaultWithdrawalLockSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogDepositIterator is returned from FilterLogDeposit and is used to iterate over the raw logs and unpacked data for LogDeposit events raised by the DepositTestnet contract.
type DepositTestnetLogDepositIterator struct {
	Event *DepositTestnetLogDeposit // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogDeposit)
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
		it.Event = new(DepositTestnetLogDeposit)
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
func (it *DepositTestnetLogDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogDeposit represents a LogDeposit event raised by the DepositTestnet contract.
type DepositTestnetLogDeposit struct {
	DepositorEthKey    common.Address
	StarkKey           *big.Int
	VaultId            *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDeposit is a free log retrieval operation binding the contract event 0x06724742ccc8c330a39a641ef02a0b419bd09248360680bb38159b0a8c2635d6.
//
// Solidity: event LogDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogDeposit(opts *bind.FilterOpts) (*DepositTestnetLogDepositIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogDeposit")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogDepositIterator{contract: _DepositTestnet.contract, event: "LogDeposit", logs: logs, sub: sub}, nil
}

// WatchLogDeposit is a free log subscription operation binding the contract event 0x06724742ccc8c330a39a641ef02a0b419bd09248360680bb38159b0a8c2635d6.
//
// Solidity: event LogDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogDeposit(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogDeposit) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogDeposit)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogDeposit", log); err != nil {
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

// ParseLogDeposit is a log parse operation binding the contract event 0x06724742ccc8c330a39a641ef02a0b419bd09248360680bb38159b0a8c2635d6.
//
// Solidity: event LogDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogDeposit(log types.Log) (*DepositTestnetLogDeposit, error) {
	event := new(DepositTestnetLogDeposit)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogDepositCancelIterator is returned from FilterLogDepositCancel and is used to iterate over the raw logs and unpacked data for LogDepositCancel events raised by the DepositTestnet contract.
type DepositTestnetLogDepositCancelIterator struct {
	Event *DepositTestnetLogDepositCancel // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogDepositCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogDepositCancel)
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
		it.Event = new(DepositTestnetLogDepositCancel)
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
func (it *DepositTestnetLogDepositCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogDepositCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogDepositCancel represents a LogDepositCancel event raised by the DepositTestnet contract.
type DepositTestnetLogDepositCancel struct {
	StarkKey *big.Int
	VaultId  *big.Int
	AssetId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogDepositCancel is a free log retrieval operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogDepositCancel(opts *bind.FilterOpts) (*DepositTestnetLogDepositCancelIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogDepositCancel")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogDepositCancelIterator{contract: _DepositTestnet.contract, event: "LogDepositCancel", logs: logs, sub: sub}, nil
}

// WatchLogDepositCancel is a free log subscription operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogDepositCancel(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogDepositCancel) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogDepositCancel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogDepositCancel)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogDepositCancel", log); err != nil {
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

// ParseLogDepositCancel is a log parse operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogDepositCancel(log types.Log) (*DepositTestnetLogDepositCancel, error) {
	event := new(DepositTestnetLogDepositCancel)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogDepositCancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogDepositCancelReclaimedIterator is returned from FilterLogDepositCancelReclaimed and is used to iterate over the raw logs and unpacked data for LogDepositCancelReclaimed events raised by the DepositTestnet contract.
type DepositTestnetLogDepositCancelReclaimedIterator struct {
	Event *DepositTestnetLogDepositCancelReclaimed // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogDepositCancelReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogDepositCancelReclaimed)
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
		it.Event = new(DepositTestnetLogDepositCancelReclaimed)
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
func (it *DepositTestnetLogDepositCancelReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogDepositCancelReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogDepositCancelReclaimed represents a LogDepositCancelReclaimed event raised by the DepositTestnet contract.
type DepositTestnetLogDepositCancelReclaimed struct {
	StarkKey           *big.Int
	VaultId            *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDepositCancelReclaimed is a free log retrieval operation binding the contract event 0xe3e46ecf1138180bf93cba62a0b7e661d976a8ab3d40243f7b082667d8f500af.
//
// Solidity: event LogDepositCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogDepositCancelReclaimed(opts *bind.FilterOpts) (*DepositTestnetLogDepositCancelReclaimedIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogDepositCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogDepositCancelReclaimedIterator{contract: _DepositTestnet.contract, event: "LogDepositCancelReclaimed", logs: logs, sub: sub}, nil
}

// WatchLogDepositCancelReclaimed is a free log subscription operation binding the contract event 0xe3e46ecf1138180bf93cba62a0b7e661d976a8ab3d40243f7b082667d8f500af.
//
// Solidity: event LogDepositCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogDepositCancelReclaimed(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogDepositCancelReclaimed) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogDepositCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogDepositCancelReclaimed)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogDepositCancelReclaimed", log); err != nil {
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

// ParseLogDepositCancelReclaimed is a log parse operation binding the contract event 0xe3e46ecf1138180bf93cba62a0b7e661d976a8ab3d40243f7b082667d8f500af.
//
// Solidity: event LogDepositCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogDepositCancelReclaimed(log types.Log) (*DepositTestnetLogDepositCancelReclaimed, error) {
	event := new(DepositTestnetLogDepositCancelReclaimed)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogDepositCancelReclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogDepositNftCancelReclaimedIterator is returned from FilterLogDepositNftCancelReclaimed and is used to iterate over the raw logs and unpacked data for LogDepositNftCancelReclaimed events raised by the DepositTestnet contract.
type DepositTestnetLogDepositNftCancelReclaimedIterator struct {
	Event *DepositTestnetLogDepositNftCancelReclaimed // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogDepositNftCancelReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogDepositNftCancelReclaimed)
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
		it.Event = new(DepositTestnetLogDepositNftCancelReclaimed)
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
func (it *DepositTestnetLogDepositNftCancelReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogDepositNftCancelReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogDepositNftCancelReclaimed represents a LogDepositNftCancelReclaimed event raised by the DepositTestnet contract.
type DepositTestnetLogDepositNftCancelReclaimed struct {
	StarkKey  *big.Int
	VaultId   *big.Int
	AssetType *big.Int
	TokenId   *big.Int
	AssetId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogDepositNftCancelReclaimed is a free log retrieval operation binding the contract event 0xf00c0c1a754f6df7545d96a7e12aad552728b94ca6aa94f81e297bdbcf1dab9c.
//
// Solidity: event LogDepositNftCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogDepositNftCancelReclaimed(opts *bind.FilterOpts) (*DepositTestnetLogDepositNftCancelReclaimedIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogDepositNftCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogDepositNftCancelReclaimedIterator{contract: _DepositTestnet.contract, event: "LogDepositNftCancelReclaimed", logs: logs, sub: sub}, nil
}

// WatchLogDepositNftCancelReclaimed is a free log subscription operation binding the contract event 0xf00c0c1a754f6df7545d96a7e12aad552728b94ca6aa94f81e297bdbcf1dab9c.
//
// Solidity: event LogDepositNftCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogDepositNftCancelReclaimed(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogDepositNftCancelReclaimed) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogDepositNftCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogDepositNftCancelReclaimed)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogDepositNftCancelReclaimed", log); err != nil {
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

// ParseLogDepositNftCancelReclaimed is a log parse operation binding the contract event 0xf00c0c1a754f6df7545d96a7e12aad552728b94ca6aa94f81e297bdbcf1dab9c.
//
// Solidity: event LogDepositNftCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogDepositNftCancelReclaimed(log types.Log) (*DepositTestnetLogDepositNftCancelReclaimed, error) {
	event := new(DepositTestnetLogDepositNftCancelReclaimed)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogDepositNftCancelReclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogDepositToVaultIterator is returned from FilterLogDepositToVault and is used to iterate over the raw logs and unpacked data for LogDepositToVault events raised by the DepositTestnet contract.
type DepositTestnetLogDepositToVaultIterator struct {
	Event *DepositTestnetLogDepositToVault // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogDepositToVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogDepositToVault)
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
		it.Event = new(DepositTestnetLogDepositToVault)
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
func (it *DepositTestnetLogDepositToVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogDepositToVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogDepositToVault represents a LogDepositToVault event raised by the DepositTestnet contract.
type DepositTestnetLogDepositToVault struct {
	EthKey             common.Address
	AssetId            *big.Int
	VaultId            *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDepositToVault is a free log retrieval operation binding the contract event 0x31f9107e2d3910ab3e05f1950b34a1983eceaed380fbcdc68ae159c4cf8cee75.
//
// Solidity: event LogDepositToVault(address ethKey, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogDepositToVault(opts *bind.FilterOpts) (*DepositTestnetLogDepositToVaultIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogDepositToVault")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogDepositToVaultIterator{contract: _DepositTestnet.contract, event: "LogDepositToVault", logs: logs, sub: sub}, nil
}

// WatchLogDepositToVault is a free log subscription operation binding the contract event 0x31f9107e2d3910ab3e05f1950b34a1983eceaed380fbcdc68ae159c4cf8cee75.
//
// Solidity: event LogDepositToVault(address ethKey, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogDepositToVault(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogDepositToVault) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogDepositToVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogDepositToVault)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogDepositToVault", log); err != nil {
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

// ParseLogDepositToVault is a log parse operation binding the contract event 0x31f9107e2d3910ab3e05f1950b34a1983eceaed380fbcdc68ae159c4cf8cee75.
//
// Solidity: event LogDepositToVault(address ethKey, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogDepositToVault(log types.Log) (*DepositTestnetLogDepositToVault, error) {
	event := new(DepositTestnetLogDepositToVault)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogDepositToVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogFrozenIterator is returned from FilterLogFrozen and is used to iterate over the raw logs and unpacked data for LogFrozen events raised by the DepositTestnet contract.
type DepositTestnetLogFrozenIterator struct {
	Event *DepositTestnetLogFrozen // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogFrozen)
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
		it.Event = new(DepositTestnetLogFrozen)
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
func (it *DepositTestnetLogFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogFrozen represents a LogFrozen event raised by the DepositTestnet contract.
type DepositTestnetLogFrozen struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogFrozen is a free log retrieval operation binding the contract event 0xf5b8e6419478ab140eb98026ab5bd607038cb0ac4d4dad5b1fc0848dfd203d1f.
//
// Solidity: event LogFrozen()
func (_DepositTestnet *DepositTestnetFilterer) FilterLogFrozen(opts *bind.FilterOpts) (*DepositTestnetLogFrozenIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogFrozen")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogFrozenIterator{contract: _DepositTestnet.contract, event: "LogFrozen", logs: logs, sub: sub}, nil
}

// WatchLogFrozen is a free log subscription operation binding the contract event 0xf5b8e6419478ab140eb98026ab5bd607038cb0ac4d4dad5b1fc0848dfd203d1f.
//
// Solidity: event LogFrozen()
func (_DepositTestnet *DepositTestnetFilterer) WatchLogFrozen(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogFrozen) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogFrozen)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogFrozen", log); err != nil {
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

// ParseLogFrozen is a log parse operation binding the contract event 0xf5b8e6419478ab140eb98026ab5bd607038cb0ac4d4dad5b1fc0848dfd203d1f.
//
// Solidity: event LogFrozen()
func (_DepositTestnet *DepositTestnetFilterer) ParseLogFrozen(log types.Log) (*DepositTestnetLogFrozen, error) {
	event := new(DepositTestnetLogFrozen)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogFullWithdrawalRequestIterator is returned from FilterLogFullWithdrawalRequest and is used to iterate over the raw logs and unpacked data for LogFullWithdrawalRequest events raised by the DepositTestnet contract.
type DepositTestnetLogFullWithdrawalRequestIterator struct {
	Event *DepositTestnetLogFullWithdrawalRequest // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogFullWithdrawalRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogFullWithdrawalRequest)
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
		it.Event = new(DepositTestnetLogFullWithdrawalRequest)
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
func (it *DepositTestnetLogFullWithdrawalRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogFullWithdrawalRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogFullWithdrawalRequest represents a LogFullWithdrawalRequest event raised by the DepositTestnet contract.
type DepositTestnetLogFullWithdrawalRequest struct {
	StarkKey *big.Int
	VaultId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogFullWithdrawalRequest is a free log retrieval operation binding the contract event 0x08eb46dbb87dcfe92d4846e5766802051525fba08a9b48318f5e0fe41186d298.
//
// Solidity: event LogFullWithdrawalRequest(uint256 starkKey, uint256 vaultId)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogFullWithdrawalRequest(opts *bind.FilterOpts) (*DepositTestnetLogFullWithdrawalRequestIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogFullWithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogFullWithdrawalRequestIterator{contract: _DepositTestnet.contract, event: "LogFullWithdrawalRequest", logs: logs, sub: sub}, nil
}

// WatchLogFullWithdrawalRequest is a free log subscription operation binding the contract event 0x08eb46dbb87dcfe92d4846e5766802051525fba08a9b48318f5e0fe41186d298.
//
// Solidity: event LogFullWithdrawalRequest(uint256 starkKey, uint256 vaultId)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogFullWithdrawalRequest(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogFullWithdrawalRequest) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogFullWithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogFullWithdrawalRequest)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogFullWithdrawalRequest", log); err != nil {
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

// ParseLogFullWithdrawalRequest is a log parse operation binding the contract event 0x08eb46dbb87dcfe92d4846e5766802051525fba08a9b48318f5e0fe41186d298.
//
// Solidity: event LogFullWithdrawalRequest(uint256 starkKey, uint256 vaultId)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogFullWithdrawalRequest(log types.Log) (*DepositTestnetLogFullWithdrawalRequest, error) {
	event := new(DepositTestnetLogFullWithdrawalRequest)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogFullWithdrawalRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogMintWithdrawalPerformedIterator is returned from FilterLogMintWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogMintWithdrawalPerformed events raised by the DepositTestnet contract.
type DepositTestnetLogMintWithdrawalPerformedIterator struct {
	Event *DepositTestnetLogMintWithdrawalPerformed // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogMintWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogMintWithdrawalPerformed)
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
		it.Event = new(DepositTestnetLogMintWithdrawalPerformed)
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
func (it *DepositTestnetLogMintWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogMintWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogMintWithdrawalPerformed represents a LogMintWithdrawalPerformed event raised by the DepositTestnet contract.
type DepositTestnetLogMintWithdrawalPerformed struct {
	OwnerKey           *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	AssetId            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogMintWithdrawalPerformed is a free log retrieval operation binding the contract event 0x7e6e15df814c1a309a57686de672b2bedd128eacde35c5370c36d6840d4e9a92.
//
// Solidity: event LogMintWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, uint256 assetId)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogMintWithdrawalPerformed(opts *bind.FilterOpts) (*DepositTestnetLogMintWithdrawalPerformedIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogMintWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogMintWithdrawalPerformedIterator{contract: _DepositTestnet.contract, event: "LogMintWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogMintWithdrawalPerformed is a free log subscription operation binding the contract event 0x7e6e15df814c1a309a57686de672b2bedd128eacde35c5370c36d6840d4e9a92.
//
// Solidity: event LogMintWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, uint256 assetId)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogMintWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogMintWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogMintWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogMintWithdrawalPerformed)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogMintWithdrawalPerformed", log); err != nil {
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

// ParseLogMintWithdrawalPerformed is a log parse operation binding the contract event 0x7e6e15df814c1a309a57686de672b2bedd128eacde35c5370c36d6840d4e9a92.
//
// Solidity: event LogMintWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, uint256 assetId)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogMintWithdrawalPerformed(log types.Log) (*DepositTestnetLogMintWithdrawalPerformed, error) {
	event := new(DepositTestnetLogMintWithdrawalPerformed)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogMintWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogMintableWithdrawalAllowedIterator is returned from FilterLogMintableWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogMintableWithdrawalAllowed events raised by the DepositTestnet contract.
type DepositTestnetLogMintableWithdrawalAllowedIterator struct {
	Event *DepositTestnetLogMintableWithdrawalAllowed // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogMintableWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogMintableWithdrawalAllowed)
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
		it.Event = new(DepositTestnetLogMintableWithdrawalAllowed)
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
func (it *DepositTestnetLogMintableWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogMintableWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogMintableWithdrawalAllowed represents a LogMintableWithdrawalAllowed event raised by the DepositTestnet contract.
type DepositTestnetLogMintableWithdrawalAllowed struct {
	OwnerKey        *big.Int
	AssetId         *big.Int
	QuantizedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogMintableWithdrawalAllowed is a free log retrieval operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogMintableWithdrawalAllowed(opts *bind.FilterOpts) (*DepositTestnetLogMintableWithdrawalAllowedIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogMintableWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogMintableWithdrawalAllowedIterator{contract: _DepositTestnet.contract, event: "LogMintableWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogMintableWithdrawalAllowed is a free log subscription operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogMintableWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogMintableWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogMintableWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogMintableWithdrawalAllowed)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogMintableWithdrawalAllowed", log); err != nil {
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

// ParseLogMintableWithdrawalAllowed is a log parse operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogMintableWithdrawalAllowed(log types.Log) (*DepositTestnetLogMintableWithdrawalAllowed, error) {
	event := new(DepositTestnetLogMintableWithdrawalAllowed)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogMintableWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogNewGovernorAcceptedIterator is returned from FilterLogNewGovernorAccepted and is used to iterate over the raw logs and unpacked data for LogNewGovernorAccepted events raised by the DepositTestnet contract.
type DepositTestnetLogNewGovernorAcceptedIterator struct {
	Event *DepositTestnetLogNewGovernorAccepted // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogNewGovernorAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogNewGovernorAccepted)
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
		it.Event = new(DepositTestnetLogNewGovernorAccepted)
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
func (it *DepositTestnetLogNewGovernorAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogNewGovernorAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogNewGovernorAccepted represents a LogNewGovernorAccepted event raised by the DepositTestnet contract.
type DepositTestnetLogNewGovernorAccepted struct {
	AcceptedGovernor common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogNewGovernorAccepted is a free log retrieval operation binding the contract event 0xcfb473e6c03f9a29ddaf990e736fa3de5188a0bd85d684f5b6e164ebfbfff5d2.
//
// Solidity: event LogNewGovernorAccepted(address acceptedGovernor)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogNewGovernorAccepted(opts *bind.FilterOpts) (*DepositTestnetLogNewGovernorAcceptedIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogNewGovernorAccepted")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogNewGovernorAcceptedIterator{contract: _DepositTestnet.contract, event: "LogNewGovernorAccepted", logs: logs, sub: sub}, nil
}

// WatchLogNewGovernorAccepted is a free log subscription operation binding the contract event 0xcfb473e6c03f9a29ddaf990e736fa3de5188a0bd85d684f5b6e164ebfbfff5d2.
//
// Solidity: event LogNewGovernorAccepted(address acceptedGovernor)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogNewGovernorAccepted(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogNewGovernorAccepted) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogNewGovernorAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogNewGovernorAccepted)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogNewGovernorAccepted", log); err != nil {
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

// ParseLogNewGovernorAccepted is a log parse operation binding the contract event 0xcfb473e6c03f9a29ddaf990e736fa3de5188a0bd85d684f5b6e164ebfbfff5d2.
//
// Solidity: event LogNewGovernorAccepted(address acceptedGovernor)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogNewGovernorAccepted(log types.Log) (*DepositTestnetLogNewGovernorAccepted, error) {
	event := new(DepositTestnetLogNewGovernorAccepted)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogNewGovernorAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogNftDepositIterator is returned from FilterLogNftDeposit and is used to iterate over the raw logs and unpacked data for LogNftDeposit events raised by the DepositTestnet contract.
type DepositTestnetLogNftDepositIterator struct {
	Event *DepositTestnetLogNftDeposit // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogNftDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogNftDeposit)
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
		it.Event = new(DepositTestnetLogNftDeposit)
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
func (it *DepositTestnetLogNftDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogNftDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogNftDeposit represents a LogNftDeposit event raised by the DepositTestnet contract.
type DepositTestnetLogNftDeposit struct {
	DepositorEthKey common.Address
	StarkKey        *big.Int
	VaultId         *big.Int
	AssetType       *big.Int
	TokenId         *big.Int
	AssetId         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogNftDeposit is a free log retrieval operation binding the contract event 0x0fcf2162832b2d6033d4d34d2f45a28d9cfee523f1899945bbdd32529cfda67b.
//
// Solidity: event LogNftDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogNftDeposit(opts *bind.FilterOpts) (*DepositTestnetLogNftDepositIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogNftDeposit")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogNftDepositIterator{contract: _DepositTestnet.contract, event: "LogNftDeposit", logs: logs, sub: sub}, nil
}

// WatchLogNftDeposit is a free log subscription operation binding the contract event 0x0fcf2162832b2d6033d4d34d2f45a28d9cfee523f1899945bbdd32529cfda67b.
//
// Solidity: event LogNftDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogNftDeposit(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogNftDeposit) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogNftDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogNftDeposit)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogNftDeposit", log); err != nil {
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

// ParseLogNftDeposit is a log parse operation binding the contract event 0x0fcf2162832b2d6033d4d34d2f45a28d9cfee523f1899945bbdd32529cfda67b.
//
// Solidity: event LogNftDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogNftDeposit(log types.Log) (*DepositTestnetLogNftDeposit, error) {
	event := new(DepositTestnetLogNftDeposit)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogNftDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogNftWithdrawalAllowedIterator is returned from FilterLogNftWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogNftWithdrawalAllowed events raised by the DepositTestnet contract.
type DepositTestnetLogNftWithdrawalAllowedIterator struct {
	Event *DepositTestnetLogNftWithdrawalAllowed // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogNftWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogNftWithdrawalAllowed)
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
		it.Event = new(DepositTestnetLogNftWithdrawalAllowed)
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
func (it *DepositTestnetLogNftWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogNftWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogNftWithdrawalAllowed represents a LogNftWithdrawalAllowed event raised by the DepositTestnet contract.
type DepositTestnetLogNftWithdrawalAllowed struct {
	OwnerKey *big.Int
	AssetId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNftWithdrawalAllowed is a free log retrieval operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 ownerKey, uint256 assetId)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogNftWithdrawalAllowed(opts *bind.FilterOpts) (*DepositTestnetLogNftWithdrawalAllowedIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogNftWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogNftWithdrawalAllowedIterator{contract: _DepositTestnet.contract, event: "LogNftWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogNftWithdrawalAllowed is a free log subscription operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 ownerKey, uint256 assetId)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogNftWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogNftWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogNftWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogNftWithdrawalAllowed)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogNftWithdrawalAllowed", log); err != nil {
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

// ParseLogNftWithdrawalAllowed is a log parse operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 ownerKey, uint256 assetId)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogNftWithdrawalAllowed(log types.Log) (*DepositTestnetLogNftWithdrawalAllowed, error) {
	event := new(DepositTestnetLogNftWithdrawalAllowed)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogNftWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogNftWithdrawalPerformedIterator is returned from FilterLogNftWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogNftWithdrawalPerformed events raised by the DepositTestnet contract.
type DepositTestnetLogNftWithdrawalPerformedIterator struct {
	Event *DepositTestnetLogNftWithdrawalPerformed // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogNftWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogNftWithdrawalPerformed)
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
		it.Event = new(DepositTestnetLogNftWithdrawalPerformed)
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
func (it *DepositTestnetLogNftWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogNftWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogNftWithdrawalPerformed represents a LogNftWithdrawalPerformed event raised by the DepositTestnet contract.
type DepositTestnetLogNftWithdrawalPerformed struct {
	OwnerKey  *big.Int
	AssetType *big.Int
	TokenId   *big.Int
	AssetId   *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogNftWithdrawalPerformed is a free log retrieval operation binding the contract event 0xa5cfa8e2199ec5b8ca319288bcab72734207d30569756ee594a74b4df7abbf41.
//
// Solidity: event LogNftWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 tokenId, uint256 assetId, address recipient)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogNftWithdrawalPerformed(opts *bind.FilterOpts) (*DepositTestnetLogNftWithdrawalPerformedIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogNftWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogNftWithdrawalPerformedIterator{contract: _DepositTestnet.contract, event: "LogNftWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogNftWithdrawalPerformed is a free log subscription operation binding the contract event 0xa5cfa8e2199ec5b8ca319288bcab72734207d30569756ee594a74b4df7abbf41.
//
// Solidity: event LogNftWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 tokenId, uint256 assetId, address recipient)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogNftWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogNftWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogNftWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogNftWithdrawalPerformed)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogNftWithdrawalPerformed", log); err != nil {
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

// ParseLogNftWithdrawalPerformed is a log parse operation binding the contract event 0xa5cfa8e2199ec5b8ca319288bcab72734207d30569756ee594a74b4df7abbf41.
//
// Solidity: event LogNftWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 tokenId, uint256 assetId, address recipient)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogNftWithdrawalPerformed(log types.Log) (*DepositTestnetLogNftWithdrawalPerformed, error) {
	event := new(DepositTestnetLogNftWithdrawalPerformed)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogNftWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogNominatedGovernorIterator is returned from FilterLogNominatedGovernor and is used to iterate over the raw logs and unpacked data for LogNominatedGovernor events raised by the DepositTestnet contract.
type DepositTestnetLogNominatedGovernorIterator struct {
	Event *DepositTestnetLogNominatedGovernor // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogNominatedGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogNominatedGovernor)
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
		it.Event = new(DepositTestnetLogNominatedGovernor)
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
func (it *DepositTestnetLogNominatedGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogNominatedGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogNominatedGovernor represents a LogNominatedGovernor event raised by the DepositTestnet contract.
type DepositTestnetLogNominatedGovernor struct {
	NominatedGovernor common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogNominatedGovernor is a free log retrieval operation binding the contract event 0x6166272c8d3f5f579082f2827532732f97195007983bb5b83ac12c56700b01a6.
//
// Solidity: event LogNominatedGovernor(address nominatedGovernor)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogNominatedGovernor(opts *bind.FilterOpts) (*DepositTestnetLogNominatedGovernorIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogNominatedGovernor")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogNominatedGovernorIterator{contract: _DepositTestnet.contract, event: "LogNominatedGovernor", logs: logs, sub: sub}, nil
}

// WatchLogNominatedGovernor is a free log subscription operation binding the contract event 0x6166272c8d3f5f579082f2827532732f97195007983bb5b83ac12c56700b01a6.
//
// Solidity: event LogNominatedGovernor(address nominatedGovernor)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogNominatedGovernor(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogNominatedGovernor) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogNominatedGovernor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogNominatedGovernor)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogNominatedGovernor", log); err != nil {
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

// ParseLogNominatedGovernor is a log parse operation binding the contract event 0x6166272c8d3f5f579082f2827532732f97195007983bb5b83ac12c56700b01a6.
//
// Solidity: event LogNominatedGovernor(address nominatedGovernor)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogNominatedGovernor(log types.Log) (*DepositTestnetLogNominatedGovernor, error) {
	event := new(DepositTestnetLogNominatedGovernor)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogNominatedGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogNominationCancelledIterator is returned from FilterLogNominationCancelled and is used to iterate over the raw logs and unpacked data for LogNominationCancelled events raised by the DepositTestnet contract.
type DepositTestnetLogNominationCancelledIterator struct {
	Event *DepositTestnetLogNominationCancelled // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogNominationCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogNominationCancelled)
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
		it.Event = new(DepositTestnetLogNominationCancelled)
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
func (it *DepositTestnetLogNominationCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogNominationCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogNominationCancelled represents a LogNominationCancelled event raised by the DepositTestnet contract.
type DepositTestnetLogNominationCancelled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNominationCancelled is a free log retrieval operation binding the contract event 0x7a8dc7dd7fffb43c4807438fa62729225156941e641fd877938f4edade3429f5.
//
// Solidity: event LogNominationCancelled()
func (_DepositTestnet *DepositTestnetFilterer) FilterLogNominationCancelled(opts *bind.FilterOpts) (*DepositTestnetLogNominationCancelledIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogNominationCancelled")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogNominationCancelledIterator{contract: _DepositTestnet.contract, event: "LogNominationCancelled", logs: logs, sub: sub}, nil
}

// WatchLogNominationCancelled is a free log subscription operation binding the contract event 0x7a8dc7dd7fffb43c4807438fa62729225156941e641fd877938f4edade3429f5.
//
// Solidity: event LogNominationCancelled()
func (_DepositTestnet *DepositTestnetFilterer) WatchLogNominationCancelled(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogNominationCancelled) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogNominationCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogNominationCancelled)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogNominationCancelled", log); err != nil {
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

// ParseLogNominationCancelled is a log parse operation binding the contract event 0x7a8dc7dd7fffb43c4807438fa62729225156941e641fd877938f4edade3429f5.
//
// Solidity: event LogNominationCancelled()
func (_DepositTestnet *DepositTestnetFilterer) ParseLogNominationCancelled(log types.Log) (*DepositTestnetLogNominationCancelled, error) {
	event := new(DepositTestnetLogNominationCancelled)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogNominationCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogOperatorAddedIterator is returned from FilterLogOperatorAdded and is used to iterate over the raw logs and unpacked data for LogOperatorAdded events raised by the DepositTestnet contract.
type DepositTestnetLogOperatorAddedIterator struct {
	Event *DepositTestnetLogOperatorAdded // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogOperatorAdded)
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
		it.Event = new(DepositTestnetLogOperatorAdded)
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
func (it *DepositTestnetLogOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogOperatorAdded represents a LogOperatorAdded event raised by the DepositTestnet contract.
type DepositTestnetLogOperatorAdded struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogOperatorAdded is a free log retrieval operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogOperatorAdded(opts *bind.FilterOpts) (*DepositTestnetLogOperatorAddedIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogOperatorAdded")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogOperatorAddedIterator{contract: _DepositTestnet.contract, event: "LogOperatorAdded", logs: logs, sub: sub}, nil
}

// WatchLogOperatorAdded is a free log subscription operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogOperatorAdded(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogOperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogOperatorAdded)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogOperatorAdded", log); err != nil {
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

// ParseLogOperatorAdded is a log parse operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogOperatorAdded(log types.Log) (*DepositTestnetLogOperatorAdded, error) {
	event := new(DepositTestnetLogOperatorAdded)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogOperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogOperatorRemovedIterator is returned from FilterLogOperatorRemoved and is used to iterate over the raw logs and unpacked data for LogOperatorRemoved events raised by the DepositTestnet contract.
type DepositTestnetLogOperatorRemovedIterator struct {
	Event *DepositTestnetLogOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogOperatorRemoved)
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
		it.Event = new(DepositTestnetLogOperatorRemoved)
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
func (it *DepositTestnetLogOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogOperatorRemoved represents a LogOperatorRemoved event raised by the DepositTestnet contract.
type DepositTestnetLogOperatorRemoved struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogOperatorRemoved is a free log retrieval operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogOperatorRemoved(opts *bind.FilterOpts) (*DepositTestnetLogOperatorRemovedIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogOperatorRemoved")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogOperatorRemovedIterator{contract: _DepositTestnet.contract, event: "LogOperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchLogOperatorRemoved is a free log subscription operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogOperatorRemoved(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogOperatorRemoved) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogOperatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogOperatorRemoved)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogOperatorRemoved", log); err != nil {
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

// ParseLogOperatorRemoved is a log parse operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogOperatorRemoved(log types.Log) (*DepositTestnetLogOperatorRemoved, error) {
	event := new(DepositTestnetLogOperatorRemoved)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogOperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogRegisteredIterator is returned from FilterLogRegistered and is used to iterate over the raw logs and unpacked data for LogRegistered events raised by the DepositTestnet contract.
type DepositTestnetLogRegisteredIterator struct {
	Event *DepositTestnetLogRegistered // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogRegistered)
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
		it.Event = new(DepositTestnetLogRegistered)
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
func (it *DepositTestnetLogRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogRegistered represents a LogRegistered event raised by the DepositTestnet contract.
type DepositTestnetLogRegistered struct {
	Entry   common.Address
	EntryId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogRegistered is a free log retrieval operation binding the contract event 0xaa7f29c97c6763919ef56006f07fbf5c1ac734b0414665df43cecdbce9010c9b.
//
// Solidity: event LogRegistered(address entry, string entryId)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogRegistered(opts *bind.FilterOpts) (*DepositTestnetLogRegisteredIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogRegistered")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogRegisteredIterator{contract: _DepositTestnet.contract, event: "LogRegistered", logs: logs, sub: sub}, nil
}

// WatchLogRegistered is a free log subscription operation binding the contract event 0xaa7f29c97c6763919ef56006f07fbf5c1ac734b0414665df43cecdbce9010c9b.
//
// Solidity: event LogRegistered(address entry, string entryId)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogRegistered(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogRegistered) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogRegistered)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogRegistered", log); err != nil {
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

// ParseLogRegistered is a log parse operation binding the contract event 0xaa7f29c97c6763919ef56006f07fbf5c1ac734b0414665df43cecdbce9010c9b.
//
// Solidity: event LogRegistered(address entry, string entryId)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogRegistered(log types.Log) (*DepositTestnetLogRegistered, error) {
	event := new(DepositTestnetLogRegistered)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogRemovalIntentIterator is returned from FilterLogRemovalIntent and is used to iterate over the raw logs and unpacked data for LogRemovalIntent events raised by the DepositTestnet contract.
type DepositTestnetLogRemovalIntentIterator struct {
	Event *DepositTestnetLogRemovalIntent // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogRemovalIntentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogRemovalIntent)
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
		it.Event = new(DepositTestnetLogRemovalIntent)
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
func (it *DepositTestnetLogRemovalIntentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogRemovalIntentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogRemovalIntent represents a LogRemovalIntent event raised by the DepositTestnet contract.
type DepositTestnetLogRemovalIntent struct {
	Entry   common.Address
	EntryId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogRemovalIntent is a free log retrieval operation binding the contract event 0xa98ac1f696177f16ca482653307aa4a02b68cf03b14409a546de5adf946484fc.
//
// Solidity: event LogRemovalIntent(address entry, string entryId)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogRemovalIntent(opts *bind.FilterOpts) (*DepositTestnetLogRemovalIntentIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogRemovalIntent")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogRemovalIntentIterator{contract: _DepositTestnet.contract, event: "LogRemovalIntent", logs: logs, sub: sub}, nil
}

// WatchLogRemovalIntent is a free log subscription operation binding the contract event 0xa98ac1f696177f16ca482653307aa4a02b68cf03b14409a546de5adf946484fc.
//
// Solidity: event LogRemovalIntent(address entry, string entryId)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogRemovalIntent(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogRemovalIntent) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogRemovalIntent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogRemovalIntent)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogRemovalIntent", log); err != nil {
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

// ParseLogRemovalIntent is a log parse operation binding the contract event 0xa98ac1f696177f16ca482653307aa4a02b68cf03b14409a546de5adf946484fc.
//
// Solidity: event LogRemovalIntent(address entry, string entryId)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogRemovalIntent(log types.Log) (*DepositTestnetLogRemovalIntent, error) {
	event := new(DepositTestnetLogRemovalIntent)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogRemovalIntent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogRemovedIterator is returned from FilterLogRemoved and is used to iterate over the raw logs and unpacked data for LogRemoved events raised by the DepositTestnet contract.
type DepositTestnetLogRemovedIterator struct {
	Event *DepositTestnetLogRemoved // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogRemoved)
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
		it.Event = new(DepositTestnetLogRemoved)
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
func (it *DepositTestnetLogRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogRemoved represents a LogRemoved event raised by the DepositTestnet contract.
type DepositTestnetLogRemoved struct {
	Entry   common.Address
	EntryId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogRemoved is a free log retrieval operation binding the contract event 0x35b176cf4f09df896aa55e10eec90bb4c4689ea1d076135a8de3138d829d0142.
//
// Solidity: event LogRemoved(address entry, string entryId)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogRemoved(opts *bind.FilterOpts) (*DepositTestnetLogRemovedIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogRemoved")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogRemovedIterator{contract: _DepositTestnet.contract, event: "LogRemoved", logs: logs, sub: sub}, nil
}

// WatchLogRemoved is a free log subscription operation binding the contract event 0x35b176cf4f09df896aa55e10eec90bb4c4689ea1d076135a8de3138d829d0142.
//
// Solidity: event LogRemoved(address entry, string entryId)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogRemoved(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogRemoved) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogRemoved)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogRemoved", log); err != nil {
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

// ParseLogRemoved is a log parse operation binding the contract event 0x35b176cf4f09df896aa55e10eec90bb4c4689ea1d076135a8de3138d829d0142.
//
// Solidity: event LogRemoved(address entry, string entryId)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogRemoved(log types.Log) (*DepositTestnetLogRemoved, error) {
	event := new(DepositTestnetLogRemoved)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogRemovedGovernorIterator is returned from FilterLogRemovedGovernor and is used to iterate over the raw logs and unpacked data for LogRemovedGovernor events raised by the DepositTestnet contract.
type DepositTestnetLogRemovedGovernorIterator struct {
	Event *DepositTestnetLogRemovedGovernor // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogRemovedGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogRemovedGovernor)
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
		it.Event = new(DepositTestnetLogRemovedGovernor)
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
func (it *DepositTestnetLogRemovedGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogRemovedGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogRemovedGovernor represents a LogRemovedGovernor event raised by the DepositTestnet contract.
type DepositTestnetLogRemovedGovernor struct {
	RemovedGovernor common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogRemovedGovernor is a free log retrieval operation binding the contract event 0xd75f94825e770b8b512be8e74759e252ad00e102e38f50cce2f7c6f868a29599.
//
// Solidity: event LogRemovedGovernor(address removedGovernor)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogRemovedGovernor(opts *bind.FilterOpts) (*DepositTestnetLogRemovedGovernorIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogRemovedGovernor")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogRemovedGovernorIterator{contract: _DepositTestnet.contract, event: "LogRemovedGovernor", logs: logs, sub: sub}, nil
}

// WatchLogRemovedGovernor is a free log subscription operation binding the contract event 0xd75f94825e770b8b512be8e74759e252ad00e102e38f50cce2f7c6f868a29599.
//
// Solidity: event LogRemovedGovernor(address removedGovernor)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogRemovedGovernor(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogRemovedGovernor) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogRemovedGovernor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogRemovedGovernor)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogRemovedGovernor", log); err != nil {
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

// ParseLogRemovedGovernor is a log parse operation binding the contract event 0xd75f94825e770b8b512be8e74759e252ad00e102e38f50cce2f7c6f868a29599.
//
// Solidity: event LogRemovedGovernor(address removedGovernor)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogRemovedGovernor(log types.Log) (*DepositTestnetLogRemovedGovernor, error) {
	event := new(DepositTestnetLogRemovedGovernor)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogRemovedGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogRootUpdateIterator is returned from FilterLogRootUpdate and is used to iterate over the raw logs and unpacked data for LogRootUpdate events raised by the DepositTestnet contract.
type DepositTestnetLogRootUpdateIterator struct {
	Event *DepositTestnetLogRootUpdate // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogRootUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogRootUpdate)
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
		it.Event = new(DepositTestnetLogRootUpdate)
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
func (it *DepositTestnetLogRootUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogRootUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogRootUpdate represents a LogRootUpdate event raised by the DepositTestnet contract.
type DepositTestnetLogRootUpdate struct {
	SequenceNumber *big.Int
	BatchId        *big.Int
	VaultRoot      *big.Int
	OrderRoot      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogRootUpdate is a free log retrieval operation binding the contract event 0xd606ef105963a7e789d927c1d21df5111915b832996b92648138f59eb9763a20.
//
// Solidity: event LogRootUpdate(uint256 sequenceNumber, uint256 batchId, uint256 vaultRoot, uint256 orderRoot)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogRootUpdate(opts *bind.FilterOpts) (*DepositTestnetLogRootUpdateIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogRootUpdate")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogRootUpdateIterator{contract: _DepositTestnet.contract, event: "LogRootUpdate", logs: logs, sub: sub}, nil
}

// WatchLogRootUpdate is a free log subscription operation binding the contract event 0xd606ef105963a7e789d927c1d21df5111915b832996b92648138f59eb9763a20.
//
// Solidity: event LogRootUpdate(uint256 sequenceNumber, uint256 batchId, uint256 vaultRoot, uint256 orderRoot)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogRootUpdate(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogRootUpdate) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogRootUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogRootUpdate)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogRootUpdate", log); err != nil {
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

// ParseLogRootUpdate is a log parse operation binding the contract event 0xd606ef105963a7e789d927c1d21df5111915b832996b92648138f59eb9763a20.
//
// Solidity: event LogRootUpdate(uint256 sequenceNumber, uint256 batchId, uint256 vaultRoot, uint256 orderRoot)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogRootUpdate(log types.Log) (*DepositTestnetLogRootUpdate, error) {
	event := new(DepositTestnetLogRootUpdate)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogRootUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogStateTransitionFactIterator is returned from FilterLogStateTransitionFact and is used to iterate over the raw logs and unpacked data for LogStateTransitionFact events raised by the DepositTestnet contract.
type DepositTestnetLogStateTransitionFactIterator struct {
	Event *DepositTestnetLogStateTransitionFact // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogStateTransitionFactIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogStateTransitionFact)
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
		it.Event = new(DepositTestnetLogStateTransitionFact)
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
func (it *DepositTestnetLogStateTransitionFactIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogStateTransitionFactIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogStateTransitionFact represents a LogStateTransitionFact event raised by the DepositTestnet contract.
type DepositTestnetLogStateTransitionFact struct {
	StateTransitionFact [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLogStateTransitionFact is a free log retrieval operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogStateTransitionFact(opts *bind.FilterOpts) (*DepositTestnetLogStateTransitionFactIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogStateTransitionFact")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogStateTransitionFactIterator{contract: _DepositTestnet.contract, event: "LogStateTransitionFact", logs: logs, sub: sub}, nil
}

// WatchLogStateTransitionFact is a free log subscription operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogStateTransitionFact(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogStateTransitionFact) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogStateTransitionFact")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogStateTransitionFact)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogStateTransitionFact", log); err != nil {
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

// ParseLogStateTransitionFact is a log parse operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogStateTransitionFact(log types.Log) (*DepositTestnetLogStateTransitionFact, error) {
	event := new(DepositTestnetLogStateTransitionFact)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogStateTransitionFact", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogTokenAdminAddedIterator is returned from FilterLogTokenAdminAdded and is used to iterate over the raw logs and unpacked data for LogTokenAdminAdded events raised by the DepositTestnet contract.
type DepositTestnetLogTokenAdminAddedIterator struct {
	Event *DepositTestnetLogTokenAdminAdded // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogTokenAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogTokenAdminAdded)
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
		it.Event = new(DepositTestnetLogTokenAdminAdded)
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
func (it *DepositTestnetLogTokenAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogTokenAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogTokenAdminAdded represents a LogTokenAdminAdded event raised by the DepositTestnet contract.
type DepositTestnetLogTokenAdminAdded struct {
	TokenAdmin common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogTokenAdminAdded is a free log retrieval operation binding the contract event 0x9085a9044aeb6daeeb5b4bf84af42b1a1613d4056f503c4e992b6396c16bd52f.
//
// Solidity: event LogTokenAdminAdded(address tokenAdmin)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogTokenAdminAdded(opts *bind.FilterOpts) (*DepositTestnetLogTokenAdminAddedIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogTokenAdminAdded")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogTokenAdminAddedIterator{contract: _DepositTestnet.contract, event: "LogTokenAdminAdded", logs: logs, sub: sub}, nil
}

// WatchLogTokenAdminAdded is a free log subscription operation binding the contract event 0x9085a9044aeb6daeeb5b4bf84af42b1a1613d4056f503c4e992b6396c16bd52f.
//
// Solidity: event LogTokenAdminAdded(address tokenAdmin)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogTokenAdminAdded(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogTokenAdminAdded) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogTokenAdminAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogTokenAdminAdded)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogTokenAdminAdded", log); err != nil {
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

// ParseLogTokenAdminAdded is a log parse operation binding the contract event 0x9085a9044aeb6daeeb5b4bf84af42b1a1613d4056f503c4e992b6396c16bd52f.
//
// Solidity: event LogTokenAdminAdded(address tokenAdmin)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogTokenAdminAdded(log types.Log) (*DepositTestnetLogTokenAdminAdded, error) {
	event := new(DepositTestnetLogTokenAdminAdded)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogTokenAdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogTokenAdminRemovedIterator is returned from FilterLogTokenAdminRemoved and is used to iterate over the raw logs and unpacked data for LogTokenAdminRemoved events raised by the DepositTestnet contract.
type DepositTestnetLogTokenAdminRemovedIterator struct {
	Event *DepositTestnetLogTokenAdminRemoved // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogTokenAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogTokenAdminRemoved)
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
		it.Event = new(DepositTestnetLogTokenAdminRemoved)
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
func (it *DepositTestnetLogTokenAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogTokenAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogTokenAdminRemoved represents a LogTokenAdminRemoved event raised by the DepositTestnet contract.
type DepositTestnetLogTokenAdminRemoved struct {
	TokenAdmin common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogTokenAdminRemoved is a free log retrieval operation binding the contract event 0xfa49aecb996ea8d99950bb051552dfcc0b5460a0bb209867a1ed8067c32c2177.
//
// Solidity: event LogTokenAdminRemoved(address tokenAdmin)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogTokenAdminRemoved(opts *bind.FilterOpts) (*DepositTestnetLogTokenAdminRemovedIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogTokenAdminRemoved")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogTokenAdminRemovedIterator{contract: _DepositTestnet.contract, event: "LogTokenAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchLogTokenAdminRemoved is a free log subscription operation binding the contract event 0xfa49aecb996ea8d99950bb051552dfcc0b5460a0bb209867a1ed8067c32c2177.
//
// Solidity: event LogTokenAdminRemoved(address tokenAdmin)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogTokenAdminRemoved(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogTokenAdminRemoved) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogTokenAdminRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogTokenAdminRemoved)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogTokenAdminRemoved", log); err != nil {
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

// ParseLogTokenAdminRemoved is a log parse operation binding the contract event 0xfa49aecb996ea8d99950bb051552dfcc0b5460a0bb209867a1ed8067c32c2177.
//
// Solidity: event LogTokenAdminRemoved(address tokenAdmin)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogTokenAdminRemoved(log types.Log) (*DepositTestnetLogTokenAdminRemoved, error) {
	event := new(DepositTestnetLogTokenAdminRemoved)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogTokenAdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogTokenRegisteredIterator is returned from FilterLogTokenRegistered and is used to iterate over the raw logs and unpacked data for LogTokenRegistered events raised by the DepositTestnet contract.
type DepositTestnetLogTokenRegisteredIterator struct {
	Event *DepositTestnetLogTokenRegistered // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogTokenRegistered)
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
		it.Event = new(DepositTestnetLogTokenRegistered)
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
func (it *DepositTestnetLogTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogTokenRegistered represents a LogTokenRegistered event raised by the DepositTestnet contract.
type DepositTestnetLogTokenRegistered struct {
	AssetType *big.Int
	AssetInfo []byte
	Quantum   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogTokenRegistered is a free log retrieval operation binding the contract event 0x7a0efbc885500f3b4a895231945be4520e4c0ba5ef7274a225a0272c81ccbcb7.
//
// Solidity: event LogTokenRegistered(uint256 assetType, bytes assetInfo, uint256 quantum)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogTokenRegistered(opts *bind.FilterOpts) (*DepositTestnetLogTokenRegisteredIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogTokenRegistered")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogTokenRegisteredIterator{contract: _DepositTestnet.contract, event: "LogTokenRegistered", logs: logs, sub: sub}, nil
}

// WatchLogTokenRegistered is a free log subscription operation binding the contract event 0x7a0efbc885500f3b4a895231945be4520e4c0ba5ef7274a225a0272c81ccbcb7.
//
// Solidity: event LogTokenRegistered(uint256 assetType, bytes assetInfo, uint256 quantum)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogTokenRegistered(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogTokenRegistered) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogTokenRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogTokenRegistered)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogTokenRegistered", log); err != nil {
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

// ParseLogTokenRegistered is a log parse operation binding the contract event 0x7a0efbc885500f3b4a895231945be4520e4c0ba5ef7274a225a0272c81ccbcb7.
//
// Solidity: event LogTokenRegistered(uint256 assetType, bytes assetInfo, uint256 quantum)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogTokenRegistered(log types.Log) (*DepositTestnetLogTokenRegistered, error) {
	event := new(DepositTestnetLogTokenRegistered)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogTokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogUnFrozenIterator is returned from FilterLogUnFrozen and is used to iterate over the raw logs and unpacked data for LogUnFrozen events raised by the DepositTestnet contract.
type DepositTestnetLogUnFrozenIterator struct {
	Event *DepositTestnetLogUnFrozen // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogUnFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogUnFrozen)
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
		it.Event = new(DepositTestnetLogUnFrozen)
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
func (it *DepositTestnetLogUnFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogUnFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogUnFrozen represents a LogUnFrozen event raised by the DepositTestnet contract.
type DepositTestnetLogUnFrozen struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogUnFrozen is a free log retrieval operation binding the contract event 0x07017fe9180629cfffba412f65a9affcf9a121de02294179f5c058f881dcc9f8.
//
// Solidity: event LogUnFrozen()
func (_DepositTestnet *DepositTestnetFilterer) FilterLogUnFrozen(opts *bind.FilterOpts) (*DepositTestnetLogUnFrozenIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogUnFrozen")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogUnFrozenIterator{contract: _DepositTestnet.contract, event: "LogUnFrozen", logs: logs, sub: sub}, nil
}

// WatchLogUnFrozen is a free log subscription operation binding the contract event 0x07017fe9180629cfffba412f65a9affcf9a121de02294179f5c058f881dcc9f8.
//
// Solidity: event LogUnFrozen()
func (_DepositTestnet *DepositTestnetFilterer) WatchLogUnFrozen(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogUnFrozen) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogUnFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogUnFrozen)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogUnFrozen", log); err != nil {
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

// ParseLogUnFrozen is a log parse operation binding the contract event 0x07017fe9180629cfffba412f65a9affcf9a121de02294179f5c058f881dcc9f8.
//
// Solidity: event LogUnFrozen()
func (_DepositTestnet *DepositTestnetFilterer) ParseLogUnFrozen(log types.Log) (*DepositTestnetLogUnFrozen, error) {
	event := new(DepositTestnetLogUnFrozen)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogUnFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogUserRegisteredIterator is returned from FilterLogUserRegistered and is used to iterate over the raw logs and unpacked data for LogUserRegistered events raised by the DepositTestnet contract.
type DepositTestnetLogUserRegisteredIterator struct {
	Event *DepositTestnetLogUserRegistered // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogUserRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogUserRegistered)
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
		it.Event = new(DepositTestnetLogUserRegistered)
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
func (it *DepositTestnetLogUserRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogUserRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogUserRegistered represents a LogUserRegistered event raised by the DepositTestnet contract.
type DepositTestnetLogUserRegistered struct {
	EthKey   common.Address
	StarkKey *big.Int
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogUserRegistered is a free log retrieval operation binding the contract event 0xcab1cf17c190e4e2195a7b8f7b362023246fa774390432b4704ab6b29d56b07b.
//
// Solidity: event LogUserRegistered(address ethKey, uint256 starkKey, address sender)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogUserRegistered(opts *bind.FilterOpts) (*DepositTestnetLogUserRegisteredIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogUserRegistered")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogUserRegisteredIterator{contract: _DepositTestnet.contract, event: "LogUserRegistered", logs: logs, sub: sub}, nil
}

// WatchLogUserRegistered is a free log subscription operation binding the contract event 0xcab1cf17c190e4e2195a7b8f7b362023246fa774390432b4704ab6b29d56b07b.
//
// Solidity: event LogUserRegistered(address ethKey, uint256 starkKey, address sender)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogUserRegistered(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogUserRegistered) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogUserRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogUserRegistered)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogUserRegistered", log); err != nil {
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

// ParseLogUserRegistered is a log parse operation binding the contract event 0xcab1cf17c190e4e2195a7b8f7b362023246fa774390432b4704ab6b29d56b07b.
//
// Solidity: event LogUserRegistered(address ethKey, uint256 starkKey, address sender)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogUserRegistered(log types.Log) (*DepositTestnetLogUserRegistered, error) {
	event := new(DepositTestnetLogUserRegistered)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogUserRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogVaultBalanceChangeAppliedIterator is returned from FilterLogVaultBalanceChangeApplied and is used to iterate over the raw logs and unpacked data for LogVaultBalanceChangeApplied events raised by the DepositTestnet contract.
type DepositTestnetLogVaultBalanceChangeAppliedIterator struct {
	Event *DepositTestnetLogVaultBalanceChangeApplied // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogVaultBalanceChangeAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogVaultBalanceChangeApplied)
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
		it.Event = new(DepositTestnetLogVaultBalanceChangeApplied)
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
func (it *DepositTestnetLogVaultBalanceChangeAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogVaultBalanceChangeAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogVaultBalanceChangeApplied represents a LogVaultBalanceChangeApplied event raised by the DepositTestnet contract.
type DepositTestnetLogVaultBalanceChangeApplied struct {
	EthKey                common.Address
	AssetId               *big.Int
	VaultId               *big.Int
	QuantizedAmountChange *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterLogVaultBalanceChangeApplied is a free log retrieval operation binding the contract event 0x2b2b583bb5166d03b87a8e7c2785e61530ab776400fb69e1bcb13b33afef2b58.
//
// Solidity: event LogVaultBalanceChangeApplied(address ethKey, uint256 assetId, uint256 vaultId, int256 quantizedAmountChange)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogVaultBalanceChangeApplied(opts *bind.FilterOpts) (*DepositTestnetLogVaultBalanceChangeAppliedIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogVaultBalanceChangeApplied")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogVaultBalanceChangeAppliedIterator{contract: _DepositTestnet.contract, event: "LogVaultBalanceChangeApplied", logs: logs, sub: sub}, nil
}

// WatchLogVaultBalanceChangeApplied is a free log subscription operation binding the contract event 0x2b2b583bb5166d03b87a8e7c2785e61530ab776400fb69e1bcb13b33afef2b58.
//
// Solidity: event LogVaultBalanceChangeApplied(address ethKey, uint256 assetId, uint256 vaultId, int256 quantizedAmountChange)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogVaultBalanceChangeApplied(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogVaultBalanceChangeApplied) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogVaultBalanceChangeApplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogVaultBalanceChangeApplied)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogVaultBalanceChangeApplied", log); err != nil {
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

// ParseLogVaultBalanceChangeApplied is a log parse operation binding the contract event 0x2b2b583bb5166d03b87a8e7c2785e61530ab776400fb69e1bcb13b33afef2b58.
//
// Solidity: event LogVaultBalanceChangeApplied(address ethKey, uint256 assetId, uint256 vaultId, int256 quantizedAmountChange)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogVaultBalanceChangeApplied(log types.Log) (*DepositTestnetLogVaultBalanceChangeApplied, error) {
	event := new(DepositTestnetLogVaultBalanceChangeApplied)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogVaultBalanceChangeApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogVaultWithdrawalLockSetIterator is returned from FilterLogVaultWithdrawalLockSet and is used to iterate over the raw logs and unpacked data for LogVaultWithdrawalLockSet events raised by the DepositTestnet contract.
type DepositTestnetLogVaultWithdrawalLockSetIterator struct {
	Event *DepositTestnetLogVaultWithdrawalLockSet // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogVaultWithdrawalLockSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogVaultWithdrawalLockSet)
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
		it.Event = new(DepositTestnetLogVaultWithdrawalLockSet)
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
func (it *DepositTestnetLogVaultWithdrawalLockSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogVaultWithdrawalLockSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogVaultWithdrawalLockSet represents a LogVaultWithdrawalLockSet event raised by the DepositTestnet contract.
type DepositTestnetLogVaultWithdrawalLockSet struct {
	EthKey      common.Address
	AssetId     *big.Int
	VaultId     *big.Int
	TimeRelease *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogVaultWithdrawalLockSet is a free log retrieval operation binding the contract event 0xcd08744f6e50d6b64d69f8a5eddc7f66e34db287f51c1da789ff770ebc7cf73e.
//
// Solidity: event LogVaultWithdrawalLockSet(address ethKey, uint256 assetId, uint256 vaultId, uint256 timeRelease)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogVaultWithdrawalLockSet(opts *bind.FilterOpts) (*DepositTestnetLogVaultWithdrawalLockSetIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogVaultWithdrawalLockSet")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogVaultWithdrawalLockSetIterator{contract: _DepositTestnet.contract, event: "LogVaultWithdrawalLockSet", logs: logs, sub: sub}, nil
}

// WatchLogVaultWithdrawalLockSet is a free log subscription operation binding the contract event 0xcd08744f6e50d6b64d69f8a5eddc7f66e34db287f51c1da789ff770ebc7cf73e.
//
// Solidity: event LogVaultWithdrawalLockSet(address ethKey, uint256 assetId, uint256 vaultId, uint256 timeRelease)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogVaultWithdrawalLockSet(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogVaultWithdrawalLockSet) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogVaultWithdrawalLockSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogVaultWithdrawalLockSet)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogVaultWithdrawalLockSet", log); err != nil {
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

// ParseLogVaultWithdrawalLockSet is a log parse operation binding the contract event 0xcd08744f6e50d6b64d69f8a5eddc7f66e34db287f51c1da789ff770ebc7cf73e.
//
// Solidity: event LogVaultWithdrawalLockSet(address ethKey, uint256 assetId, uint256 vaultId, uint256 timeRelease)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogVaultWithdrawalLockSet(log types.Log) (*DepositTestnetLogVaultWithdrawalLockSet, error) {
	event := new(DepositTestnetLogVaultWithdrawalLockSet)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogVaultWithdrawalLockSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogWithdrawalAllowedIterator is returned from FilterLogWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogWithdrawalAllowed events raised by the DepositTestnet contract.
type DepositTestnetLogWithdrawalAllowedIterator struct {
	Event *DepositTestnetLogWithdrawalAllowed // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogWithdrawalAllowed)
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
		it.Event = new(DepositTestnetLogWithdrawalAllowed)
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
func (it *DepositTestnetLogWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogWithdrawalAllowed represents a LogWithdrawalAllowed event raised by the DepositTestnet contract.
type DepositTestnetLogWithdrawalAllowed struct {
	OwnerKey           *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawalAllowed is a free log retrieval operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogWithdrawalAllowed(opts *bind.FilterOpts) (*DepositTestnetLogWithdrawalAllowedIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogWithdrawalAllowedIterator{contract: _DepositTestnet.contract, event: "LogWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawalAllowed is a free log subscription operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogWithdrawalAllowed)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogWithdrawalAllowed", log); err != nil {
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

// ParseLogWithdrawalAllowed is a log parse operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogWithdrawalAllowed(log types.Log) (*DepositTestnetLogWithdrawalAllowed, error) {
	event := new(DepositTestnetLogWithdrawalAllowed)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogWithdrawalFromVaultIterator is returned from FilterLogWithdrawalFromVault and is used to iterate over the raw logs and unpacked data for LogWithdrawalFromVault events raised by the DepositTestnet contract.
type DepositTestnetLogWithdrawalFromVaultIterator struct {
	Event *DepositTestnetLogWithdrawalFromVault // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogWithdrawalFromVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogWithdrawalFromVault)
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
		it.Event = new(DepositTestnetLogWithdrawalFromVault)
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
func (it *DepositTestnetLogWithdrawalFromVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogWithdrawalFromVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogWithdrawalFromVault represents a LogWithdrawalFromVault event raised by the DepositTestnet contract.
type DepositTestnetLogWithdrawalFromVault struct {
	EthKey             common.Address
	AssetId            *big.Int
	VaultId            *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawalFromVault is a free log retrieval operation binding the contract event 0xa866e320a5fe47b394ae4f8cd2c180329ac2a143cc50f61d7b26a7af1f1c05fc.
//
// Solidity: event LogWithdrawalFromVault(address ethKey, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogWithdrawalFromVault(opts *bind.FilterOpts) (*DepositTestnetLogWithdrawalFromVaultIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogWithdrawalFromVault")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogWithdrawalFromVaultIterator{contract: _DepositTestnet.contract, event: "LogWithdrawalFromVault", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawalFromVault is a free log subscription operation binding the contract event 0xa866e320a5fe47b394ae4f8cd2c180329ac2a143cc50f61d7b26a7af1f1c05fc.
//
// Solidity: event LogWithdrawalFromVault(address ethKey, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogWithdrawalFromVault(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogWithdrawalFromVault) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogWithdrawalFromVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogWithdrawalFromVault)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogWithdrawalFromVault", log); err != nil {
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

// ParseLogWithdrawalFromVault is a log parse operation binding the contract event 0xa866e320a5fe47b394ae4f8cd2c180329ac2a143cc50f61d7b26a7af1f1c05fc.
//
// Solidity: event LogWithdrawalFromVault(address ethKey, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogWithdrawalFromVault(log types.Log) (*DepositTestnetLogWithdrawalFromVault, error) {
	event := new(DepositTestnetLogWithdrawalFromVault)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogWithdrawalFromVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTestnetLogWithdrawalPerformedIterator is returned from FilterLogWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogWithdrawalPerformed events raised by the DepositTestnet contract.
type DepositTestnetLogWithdrawalPerformedIterator struct {
	Event *DepositTestnetLogWithdrawalPerformed // Event containing the contract specifics and raw log

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
func (it *DepositTestnetLogWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTestnetLogWithdrawalPerformed)
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
		it.Event = new(DepositTestnetLogWithdrawalPerformed)
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
func (it *DepositTestnetLogWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTestnetLogWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTestnetLogWithdrawalPerformed represents a LogWithdrawalPerformed event raised by the DepositTestnet contract.
type DepositTestnetLogWithdrawalPerformed struct {
	OwnerKey           *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Recipient          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawalPerformed is a free log retrieval operation binding the contract event 0xb7477a7b93b2addc5272bbd7ad0986ef1c0d0bd265f26c3dc4bbe42727c2ac0c.
//
// Solidity: event LogWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_DepositTestnet *DepositTestnetFilterer) FilterLogWithdrawalPerformed(opts *bind.FilterOpts) (*DepositTestnetLogWithdrawalPerformedIterator, error) {

	logs, sub, err := _DepositTestnet.contract.FilterLogs(opts, "LogWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &DepositTestnetLogWithdrawalPerformedIterator{contract: _DepositTestnet.contract, event: "LogWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawalPerformed is a free log subscription operation binding the contract event 0xb7477a7b93b2addc5272bbd7ad0986ef1c0d0bd265f26c3dc4bbe42727c2ac0c.
//
// Solidity: event LogWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_DepositTestnet *DepositTestnetFilterer) WatchLogWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *DepositTestnetLogWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _DepositTestnet.contract.WatchLogs(opts, "LogWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTestnetLogWithdrawalPerformed)
				if err := _DepositTestnet.contract.UnpackLog(event, "LogWithdrawalPerformed", log); err != nil {
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

// ParseLogWithdrawalPerformed is a log parse operation binding the contract event 0xb7477a7b93b2addc5272bbd7ad0986ef1c0d0bd265f26c3dc4bbe42727c2ac0c.
//
// Solidity: event LogWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_DepositTestnet *DepositTestnetFilterer) ParseLogWithdrawalPerformed(log types.Log) (*DepositTestnetLogWithdrawalPerformed, error) {
	event := new(DepositTestnetLogWithdrawalPerformed)
	if err := _DepositTestnet.contract.UnpackLog(event, "LogWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
