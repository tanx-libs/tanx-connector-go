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

// DepositMainnetMetaData contains all meta data concerning the DepositMainnet contract.
var DepositMainnetMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"getSubContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"handlingContractId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"acceptedGovernor\",\"type\":\"address\"}],\"name\":\"LogNewGovernorAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nominatedGovernor\",\"type\":\"address\"}],\"name\":\"LogNominatedGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogNominationCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entryId\",\"type\":\"string\"}],\"name\":\"LogRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entryId\",\"type\":\"string\"}],\"name\":\"LogRemovalIntent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entryId\",\"type\":\"string\"}],\"name\":\"LogRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"removedGovernor\",\"type\":\"address\"}],\"name\":\"LogRemovedGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogUnFrozen\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSIT_CANCEL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FREEZE_GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAIN_GOVERNANCE_INFO_TAG\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FORCED_ACTIONS_REQS_PER_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_VERIFIER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNFREEZE_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_REMOVAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"announceAvailabilityVerifierRemovalIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"announceVerifierRemovalIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredAvailabilityVerifiers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_verifers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredVerifiers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_verifers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"name\":\"isAvailabilityVerifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"name\":\"isVerifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainAcceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainCancelNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"testGovernor\",\"type\":\"address\"}],\"name\":\"mainIsGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"mainNominateNewGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governorForRemoval\",\"type\":\"address\"}],\"name\":\"mainRemoveGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"name\":\"registerAvailabilityVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"name\":\"registerVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"removeAvailabilityVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"removeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unFreeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogAssetWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositorEthKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogDepositCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDepositCancelReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogDepositNftCancelReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositorEthKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDepositWithTokenId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDepositWithTokenIdCancelReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogMintWithdrawalPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogMintableWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositorEthKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogNftDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogNftWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"LogNftWithdrawalPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAdmin\",\"type\":\"address\"}],\"name\":\"LogTokenAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAdmin\",\"type\":\"address\"}],\"name\":\"LogTokenAdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantum\",\"type\":\"uint256\"}],\"name\":\"LogTokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"LogWithdrawalPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"LogWithdrawalWithTokenIdPerformed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"calculateAssetIdWithTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"mintingBlob\",\"type\":\"bytes\"}],\"name\":\"calculateMintableAssetId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultVaultWithdrawalLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositCancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"depositERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositNftReclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositReclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"depositWithTokenId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositWithTokenIdReclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actionIndex\",\"type\":\"uint256\"}],\"name\":\"getActionHashByIndex\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"getAssetInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getCancellationRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"request\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getDepositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"}],\"name\":\"getEthKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getFullWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getQuantizedDepositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"presumedAssetType\",\"type\":\"uint256\"}],\"name\":\"getQuantum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quantum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"isAssetRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"testedAdmin\",\"type\":\"address\"}],\"name\":\"isTokenAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"quantum\",\"type\":\"uint256\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"registerTokenAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"}],\"name\":\"unregisterTokenAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"mintingBlob\",\"type\":\"bytes\"}],\"name\":\"withdrawAndMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawWithTokenId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"LogOperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"LogOperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validiumVaultRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rollupVaultRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderRoot\",\"type\":\"uint256\"}],\"name\":\"LogRootUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateTransitionFact\",\"type\":\"bytes32\"}],\"name\":\"LogStateTransitionFact\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"quantizedAmountChange\",\"type\":\"int256\"}],\"name\":\"LogVaultBalanceChangeApplied\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"STARKEX_MAX_DEFAULT_VAULT_LOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"escape\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalConfigCode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBatchId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrderRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrderTreeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupTreeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupVaultRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidiumTreeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidiumVaultRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"testedOperator\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"removedOperator\",\"type\":\"address\"}],\"name\":\"unregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"publicInput\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"applicationData\",\"type\":\"uint256[]\"}],\"name\":\"updateState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"LogFullWithdrawalRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"LogUserRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"freezeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"fullWithdrawalRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"starkSignature\",\"type\":\"bytes\"}],\"name\":\"registerEthAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"starkSignature\",\"type\":\"bytes\"}],\"name\":\"registerSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDefaultLockTime\",\"type\":\"uint256\"}],\"name\":\"LogDefaultVaultWithdrawalLockSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDepositToVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeRelease\",\"type\":\"uint256\"}],\"name\":\"LogVaultWithdrawalLockSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawalFromVault\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"depositERC1155ToVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"depositERC20ToVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositEthToVault\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getQuantizedErc1155VaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getQuantizedVaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getVaultBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getVaultWithdrawalLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isStrictVaultBalancePolicy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"isVaultLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"}],\"name\":\"lockVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDefaultTime\",\"type\":\"uint256\"}],\"name\":\"setDefaultVaultWithdrawalLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawErc1155FromVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedActivationTime\",\"type\":\"uint256\"}],\"name\":\"ImplementationActivationRescheduled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"finalize\",\"type\":\"bool\"}],\"name\":\"updateImplementationActivationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DepositMainnetABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositMainnetMetaData.ABI instead.
var DepositMainnetABI = DepositMainnetMetaData.ABI

// DepositMainnet is an auto generated Go binding around an Ethereum contract.
type DepositMainnet struct {
	DepositMainnetCaller     // Read-only binding to the contract
	DepositMainnetTransactor // Write-only binding to the contract
	DepositMainnetFilterer   // Log filterer for contract events
}

// DepositMainnetCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositMainnetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositMainnetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositMainnetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositMainnetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositMainnetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositMainnetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositMainnetSession struct {
	Contract     *DepositMainnet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositMainnetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositMainnetCallerSession struct {
	Contract *DepositMainnetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DepositMainnetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositMainnetTransactorSession struct {
	Contract     *DepositMainnetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DepositMainnetRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositMainnetRaw struct {
	Contract *DepositMainnet // Generic contract binding to access the raw methods on
}

// DepositMainnetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositMainnetCallerRaw struct {
	Contract *DepositMainnetCaller // Generic read-only contract binding to access the raw methods on
}

// DepositMainnetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositMainnetTransactorRaw struct {
	Contract *DepositMainnetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositMainnet creates a new instance of DepositMainnet, bound to a specific deployed contract.
func NewDepositMainnet(address common.Address, backend bind.ContractBackend) (*DepositMainnet, error) {
	contract, err := bindDepositMainnet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DepositMainnet{DepositMainnetCaller: DepositMainnetCaller{contract: contract}, DepositMainnetTransactor: DepositMainnetTransactor{contract: contract}, DepositMainnetFilterer: DepositMainnetFilterer{contract: contract}}, nil
}

// NewDepositMainnetCaller creates a new read-only instance of DepositMainnet, bound to a specific deployed contract.
func NewDepositMainnetCaller(address common.Address, caller bind.ContractCaller) (*DepositMainnetCaller, error) {
	contract, err := bindDepositMainnet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositMainnetCaller{contract: contract}, nil
}

// NewDepositMainnetTransactor creates a new write-only instance of DepositMainnet, bound to a specific deployed contract.
func NewDepositMainnetTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositMainnetTransactor, error) {
	contract, err := bindDepositMainnet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositMainnetTransactor{contract: contract}, nil
}

// NewDepositMainnetFilterer creates a new log filterer instance of DepositMainnet, bound to a specific deployed contract.
func NewDepositMainnetFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositMainnetFilterer, error) {
	contract, err := bindDepositMainnet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositMainnetFilterer{contract: contract}, nil
}

// bindDepositMainnet binds a generic wrapper to an already deployed contract.
func bindDepositMainnet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DepositMainnetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositMainnet *DepositMainnetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositMainnet.Contract.DepositMainnetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositMainnet *DepositMainnetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositMainnetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositMainnet *DepositMainnetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositMainnetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositMainnet *DepositMainnetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositMainnet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositMainnet *DepositMainnetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositMainnet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositMainnet *DepositMainnetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositMainnet.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) DEPOSITCANCELDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "DEPOSIT_CANCEL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _DepositMainnet.Contract.DEPOSITCANCELDELAY(&_DepositMainnet.CallOpts)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _DepositMainnet.Contract.DEPOSITCANCELDELAY(&_DepositMainnet.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) FREEZEGRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "FREEZE_GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _DepositMainnet.Contract.FREEZEGRACEPERIOD(&_DepositMainnet.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _DepositMainnet.Contract.FREEZEGRACEPERIOD(&_DepositMainnet.CallOpts)
}

// MAINGOVERNANCEINFOTAG is a free data retrieval call binding the contract method 0xc23b60ef.
//
// Solidity: function MAIN_GOVERNANCE_INFO_TAG() view returns(string)
func (_DepositMainnet *DepositMainnetCaller) MAINGOVERNANCEINFOTAG(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "MAIN_GOVERNANCE_INFO_TAG")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MAINGOVERNANCEINFOTAG is a free data retrieval call binding the contract method 0xc23b60ef.
//
// Solidity: function MAIN_GOVERNANCE_INFO_TAG() view returns(string)
func (_DepositMainnet *DepositMainnetSession) MAINGOVERNANCEINFOTAG() (string, error) {
	return _DepositMainnet.Contract.MAINGOVERNANCEINFOTAG(&_DepositMainnet.CallOpts)
}

// MAINGOVERNANCEINFOTAG is a free data retrieval call binding the contract method 0xc23b60ef.
//
// Solidity: function MAIN_GOVERNANCE_INFO_TAG() view returns(string)
func (_DepositMainnet *DepositMainnetCallerSession) MAINGOVERNANCEINFOTAG() (string, error) {
	return _DepositMainnet.Contract.MAINGOVERNANCEINFOTAG(&_DepositMainnet.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) MAXFORCEDACTIONSREQSPERBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "MAX_FORCED_ACTIONS_REQS_PER_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _DepositMainnet.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_DepositMainnet.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _DepositMainnet.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_DepositMainnet.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) MAXVERIFIERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "MAX_VERIFIER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _DepositMainnet.Contract.MAXVERIFIERCOUNT(&_DepositMainnet.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _DepositMainnet.Contract.MAXVERIFIERCOUNT(&_DepositMainnet.CallOpts)
}

// STARKEXMAXDEFAULTVAULTLOCK is a free data retrieval call binding the contract method 0x835a71c3.
//
// Solidity: function STARKEX_MAX_DEFAULT_VAULT_LOCK() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) STARKEXMAXDEFAULTVAULTLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "STARKEX_MAX_DEFAULT_VAULT_LOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STARKEXMAXDEFAULTVAULTLOCK is a free data retrieval call binding the contract method 0x835a71c3.
//
// Solidity: function STARKEX_MAX_DEFAULT_VAULT_LOCK() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) STARKEXMAXDEFAULTVAULTLOCK() (*big.Int, error) {
	return _DepositMainnet.Contract.STARKEXMAXDEFAULTVAULTLOCK(&_DepositMainnet.CallOpts)
}

// STARKEXMAXDEFAULTVAULTLOCK is a free data retrieval call binding the contract method 0x835a71c3.
//
// Solidity: function STARKEX_MAX_DEFAULT_VAULT_LOCK() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) STARKEXMAXDEFAULTVAULTLOCK() (*big.Int, error) {
	return _DepositMainnet.Contract.STARKEXMAXDEFAULTVAULTLOCK(&_DepositMainnet.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) UNFREEZEDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "UNFREEZE_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) UNFREEZEDELAY() (*big.Int, error) {
	return _DepositMainnet.Contract.UNFREEZEDELAY(&_DepositMainnet.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) UNFREEZEDELAY() (*big.Int, error) {
	return _DepositMainnet.Contract.UNFREEZEDELAY(&_DepositMainnet.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) VERIFIERREMOVALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "VERIFIER_REMOVAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _DepositMainnet.Contract.VERIFIERREMOVALDELAY(&_DepositMainnet.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _DepositMainnet.Contract.VERIFIERREMOVALDELAY(&_DepositMainnet.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_DepositMainnet *DepositMainnetCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_DepositMainnet *DepositMainnetSession) VERSION() (string, error) {
	return _DepositMainnet.Contract.VERSION(&_DepositMainnet.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_DepositMainnet *DepositMainnetCallerSession) VERSION() (string, error) {
	return _DepositMainnet.Contract.VERSION(&_DepositMainnet.CallOpts)
}

// CalculateAssetIdWithTokenId is a free data retrieval call binding the contract method 0xa1cc5e13.
//
// Solidity: function calculateAssetIdWithTokenId(uint256 assetType, uint256 tokenId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) CalculateAssetIdWithTokenId(opts *bind.CallOpts, assetType *big.Int, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "calculateAssetIdWithTokenId", assetType, tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateAssetIdWithTokenId is a free data retrieval call binding the contract method 0xa1cc5e13.
//
// Solidity: function calculateAssetIdWithTokenId(uint256 assetType, uint256 tokenId) view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) CalculateAssetIdWithTokenId(assetType *big.Int, tokenId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.CalculateAssetIdWithTokenId(&_DepositMainnet.CallOpts, assetType, tokenId)
}

// CalculateAssetIdWithTokenId is a free data retrieval call binding the contract method 0xa1cc5e13.
//
// Solidity: function calculateAssetIdWithTokenId(uint256 assetType, uint256 tokenId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) CalculateAssetIdWithTokenId(assetType *big.Int, tokenId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.CalculateAssetIdWithTokenId(&_DepositMainnet.CallOpts, assetType, tokenId)
}

// CalculateMintableAssetId is a free data retrieval call binding the contract method 0xb12773fb.
//
// Solidity: function calculateMintableAssetId(uint256 assetType, bytes mintingBlob) pure returns(uint256 assetId)
func (_DepositMainnet *DepositMainnetCaller) CalculateMintableAssetId(opts *bind.CallOpts, assetType *big.Int, mintingBlob []byte) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "calculateMintableAssetId", assetType, mintingBlob)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateMintableAssetId is a free data retrieval call binding the contract method 0xb12773fb.
//
// Solidity: function calculateMintableAssetId(uint256 assetType, bytes mintingBlob) pure returns(uint256 assetId)
func (_DepositMainnet *DepositMainnetSession) CalculateMintableAssetId(assetType *big.Int, mintingBlob []byte) (*big.Int, error) {
	return _DepositMainnet.Contract.CalculateMintableAssetId(&_DepositMainnet.CallOpts, assetType, mintingBlob)
}

// CalculateMintableAssetId is a free data retrieval call binding the contract method 0xb12773fb.
//
// Solidity: function calculateMintableAssetId(uint256 assetType, bytes mintingBlob) pure returns(uint256 assetId)
func (_DepositMainnet *DepositMainnetCallerSession) CalculateMintableAssetId(assetType *big.Int, mintingBlob []byte) (*big.Int, error) {
	return _DepositMainnet.Contract.CalculateMintableAssetId(&_DepositMainnet.CallOpts, assetType, mintingBlob)
}

// DefaultVaultWithdrawalLock is a free data retrieval call binding the contract method 0xa45d7841.
//
// Solidity: function defaultVaultWithdrawalLock() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) DefaultVaultWithdrawalLock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "defaultVaultWithdrawalLock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultVaultWithdrawalLock is a free data retrieval call binding the contract method 0xa45d7841.
//
// Solidity: function defaultVaultWithdrawalLock() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) DefaultVaultWithdrawalLock() (*big.Int, error) {
	return _DepositMainnet.Contract.DefaultVaultWithdrawalLock(&_DepositMainnet.CallOpts)
}

// DefaultVaultWithdrawalLock is a free data retrieval call binding the contract method 0xa45d7841.
//
// Solidity: function defaultVaultWithdrawalLock() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) DefaultVaultWithdrawalLock() (*big.Int, error) {
	return _DepositMainnet.Contract.DefaultVaultWithdrawalLock(&_DepositMainnet.CallOpts)
}

// GetActionCount is a free data retrieval call binding the contract method 0x5eecd218.
//
// Solidity: function getActionCount() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetActionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getActionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActionCount is a free data retrieval call binding the contract method 0x5eecd218.
//
// Solidity: function getActionCount() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetActionCount() (*big.Int, error) {
	return _DepositMainnet.Contract.GetActionCount(&_DepositMainnet.CallOpts)
}

// GetActionCount is a free data retrieval call binding the contract method 0x5eecd218.
//
// Solidity: function getActionCount() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetActionCount() (*big.Int, error) {
	return _DepositMainnet.Contract.GetActionCount(&_DepositMainnet.CallOpts)
}

// GetActionHashByIndex is a free data retrieval call binding the contract method 0x5e586cd1.
//
// Solidity: function getActionHashByIndex(uint256 actionIndex) view returns(bytes32)
func (_DepositMainnet *DepositMainnetCaller) GetActionHashByIndex(opts *bind.CallOpts, actionIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getActionHashByIndex", actionIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetActionHashByIndex is a free data retrieval call binding the contract method 0x5e586cd1.
//
// Solidity: function getActionHashByIndex(uint256 actionIndex) view returns(bytes32)
func (_DepositMainnet *DepositMainnetSession) GetActionHashByIndex(actionIndex *big.Int) ([32]byte, error) {
	return _DepositMainnet.Contract.GetActionHashByIndex(&_DepositMainnet.CallOpts, actionIndex)
}

// GetActionHashByIndex is a free data retrieval call binding the contract method 0x5e586cd1.
//
// Solidity: function getActionHashByIndex(uint256 actionIndex) view returns(bytes32)
func (_DepositMainnet *DepositMainnetCallerSession) GetActionHashByIndex(actionIndex *big.Int) ([32]byte, error) {
	return _DepositMainnet.Contract.GetActionHashByIndex(&_DepositMainnet.CallOpts, actionIndex)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_DepositMainnet *DepositMainnetCaller) GetAssetInfo(opts *bind.CallOpts, assetType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getAssetInfo", assetType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_DepositMainnet *DepositMainnetSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _DepositMainnet.Contract.GetAssetInfo(&_DepositMainnet.CallOpts, assetType)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_DepositMainnet *DepositMainnetCallerSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _DepositMainnet.Contract.GetAssetInfo(&_DepositMainnet.CallOpts, assetType)
}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_DepositMainnet *DepositMainnetCaller) GetCancellationRequest(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getCancellationRequest", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_DepositMainnet *DepositMainnetSession) GetCancellationRequest(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetCancellationRequest(&_DepositMainnet.CallOpts, starkKey, assetId, vaultId)
}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_DepositMainnet *DepositMainnetCallerSession) GetCancellationRequest(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetCancellationRequest(&_DepositMainnet.CallOpts, starkKey, assetId, vaultId)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetDepositBalance(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getDepositBalance", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetDepositBalance(&_DepositMainnet.CallOpts, starkKey, assetId, vaultId)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetDepositBalance(&_DepositMainnet.CallOpts, starkKey, assetId, vaultId)
}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 ownerKey) view returns(address)
func (_DepositMainnet *DepositMainnetCaller) GetEthKey(opts *bind.CallOpts, ownerKey *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getEthKey", ownerKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 ownerKey) view returns(address)
func (_DepositMainnet *DepositMainnetSession) GetEthKey(ownerKey *big.Int) (common.Address, error) {
	return _DepositMainnet.Contract.GetEthKey(&_DepositMainnet.CallOpts, ownerKey)
}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 ownerKey) view returns(address)
func (_DepositMainnet *DepositMainnetCallerSession) GetEthKey(ownerKey *big.Int) (common.Address, error) {
	return _DepositMainnet.Contract.GetEthKey(&_DepositMainnet.CallOpts, ownerKey)
}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetFullWithdrawalRequest(opts *bind.CallOpts, ownerKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getFullWithdrawalRequest", ownerKey, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetFullWithdrawalRequest(ownerKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetFullWithdrawalRequest(&_DepositMainnet.CallOpts, ownerKey, vaultId)
}

// GetFullWithdrawalRequest is a free data retrieval call binding the contract method 0x296e2f37.
//
// Solidity: function getFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetFullWithdrawalRequest(ownerKey *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetFullWithdrawalRequest(&_DepositMainnet.CallOpts, ownerKey, vaultId)
}

// GetGlobalConfigCode is a free data retrieval call binding the contract method 0xe9aa2d6b.
//
// Solidity: function getGlobalConfigCode() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetGlobalConfigCode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getGlobalConfigCode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGlobalConfigCode is a free data retrieval call binding the contract method 0xe9aa2d6b.
//
// Solidity: function getGlobalConfigCode() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetGlobalConfigCode() (*big.Int, error) {
	return _DepositMainnet.Contract.GetGlobalConfigCode(&_DepositMainnet.CallOpts)
}

// GetGlobalConfigCode is a free data retrieval call binding the contract method 0xe9aa2d6b.
//
// Solidity: function getGlobalConfigCode() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetGlobalConfigCode() (*big.Int, error) {
	return _DepositMainnet.Contract.GetGlobalConfigCode(&_DepositMainnet.CallOpts)
}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetLastBatchId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getLastBatchId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetLastBatchId() (*big.Int, error) {
	return _DepositMainnet.Contract.GetLastBatchId(&_DepositMainnet.CallOpts)
}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetLastBatchId() (*big.Int, error) {
	return _DepositMainnet.Contract.GetLastBatchId(&_DepositMainnet.CallOpts)
}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetOrderRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getOrderRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetOrderRoot() (*big.Int, error) {
	return _DepositMainnet.Contract.GetOrderRoot(&_DepositMainnet.CallOpts)
}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetOrderRoot() (*big.Int, error) {
	return _DepositMainnet.Contract.GetOrderRoot(&_DepositMainnet.CallOpts)
}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetOrderTreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getOrderTreeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetOrderTreeHeight() (*big.Int, error) {
	return _DepositMainnet.Contract.GetOrderTreeHeight(&_DepositMainnet.CallOpts)
}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetOrderTreeHeight() (*big.Int, error) {
	return _DepositMainnet.Contract.GetOrderTreeHeight(&_DepositMainnet.CallOpts)
}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetQuantizedDepositBalance(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getQuantizedDepositBalance", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetQuantizedDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetQuantizedDepositBalance(&_DepositMainnet.CallOpts, starkKey, assetId, vaultId)
}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetQuantizedDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetQuantizedDepositBalance(&_DepositMainnet.CallOpts, starkKey, assetId, vaultId)
}

// GetQuantizedErc1155VaultBalance is a free data retrieval call binding the contract method 0xf7ac8587.
//
// Solidity: function getQuantizedErc1155VaultBalance(address ethKey, uint256 assetType, uint256 tokenId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetQuantizedErc1155VaultBalance(opts *bind.CallOpts, ethKey common.Address, assetType *big.Int, tokenId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getQuantizedErc1155VaultBalance", ethKey, assetType, tokenId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantizedErc1155VaultBalance is a free data retrieval call binding the contract method 0xf7ac8587.
//
// Solidity: function getQuantizedErc1155VaultBalance(address ethKey, uint256 assetType, uint256 tokenId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetQuantizedErc1155VaultBalance(ethKey common.Address, assetType *big.Int, tokenId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetQuantizedErc1155VaultBalance(&_DepositMainnet.CallOpts, ethKey, assetType, tokenId, vaultId)
}

// GetQuantizedErc1155VaultBalance is a free data retrieval call binding the contract method 0xf7ac8587.
//
// Solidity: function getQuantizedErc1155VaultBalance(address ethKey, uint256 assetType, uint256 tokenId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetQuantizedErc1155VaultBalance(ethKey common.Address, assetType *big.Int, tokenId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetQuantizedErc1155VaultBalance(&_DepositMainnet.CallOpts, ethKey, assetType, tokenId, vaultId)
}

// GetQuantizedVaultBalance is a free data retrieval call binding the contract method 0xe69a9de7.
//
// Solidity: function getQuantizedVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetQuantizedVaultBalance(opts *bind.CallOpts, ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getQuantizedVaultBalance", ethKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantizedVaultBalance is a free data retrieval call binding the contract method 0xe69a9de7.
//
// Solidity: function getQuantizedVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetQuantizedVaultBalance(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetQuantizedVaultBalance(&_DepositMainnet.CallOpts, ethKey, assetId, vaultId)
}

// GetQuantizedVaultBalance is a free data retrieval call binding the contract method 0xe69a9de7.
//
// Solidity: function getQuantizedVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetQuantizedVaultBalance(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetQuantizedVaultBalance(&_DepositMainnet.CallOpts, ethKey, assetId, vaultId)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_DepositMainnet *DepositMainnetCaller) GetQuantum(opts *bind.CallOpts, presumedAssetType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getQuantum", presumedAssetType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_DepositMainnet *DepositMainnetSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetQuantum(&_DepositMainnet.CallOpts, presumedAssetType)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_DepositMainnet *DepositMainnetCallerSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetQuantum(&_DepositMainnet.CallOpts, presumedAssetType)
}

// GetRegisteredAvailabilityVerifiers is a free data retrieval call binding the contract method 0x1ac347f2.
//
// Solidity: function getRegisteredAvailabilityVerifiers() view returns(address[] _verifers)
func (_DepositMainnet *DepositMainnetCaller) GetRegisteredAvailabilityVerifiers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getRegisteredAvailabilityVerifiers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredAvailabilityVerifiers is a free data retrieval call binding the contract method 0x1ac347f2.
//
// Solidity: function getRegisteredAvailabilityVerifiers() view returns(address[] _verifers)
func (_DepositMainnet *DepositMainnetSession) GetRegisteredAvailabilityVerifiers() ([]common.Address, error) {
	return _DepositMainnet.Contract.GetRegisteredAvailabilityVerifiers(&_DepositMainnet.CallOpts)
}

// GetRegisteredAvailabilityVerifiers is a free data retrieval call binding the contract method 0x1ac347f2.
//
// Solidity: function getRegisteredAvailabilityVerifiers() view returns(address[] _verifers)
func (_DepositMainnet *DepositMainnetCallerSession) GetRegisteredAvailabilityVerifiers() ([]common.Address, error) {
	return _DepositMainnet.Contract.GetRegisteredAvailabilityVerifiers(&_DepositMainnet.CallOpts)
}

// GetRegisteredVerifiers is a free data retrieval call binding the contract method 0x4eab9ed3.
//
// Solidity: function getRegisteredVerifiers() view returns(address[] _verifers)
func (_DepositMainnet *DepositMainnetCaller) GetRegisteredVerifiers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getRegisteredVerifiers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredVerifiers is a free data retrieval call binding the contract method 0x4eab9ed3.
//
// Solidity: function getRegisteredVerifiers() view returns(address[] _verifers)
func (_DepositMainnet *DepositMainnetSession) GetRegisteredVerifiers() ([]common.Address, error) {
	return _DepositMainnet.Contract.GetRegisteredVerifiers(&_DepositMainnet.CallOpts)
}

// GetRegisteredVerifiers is a free data retrieval call binding the contract method 0x4eab9ed3.
//
// Solidity: function getRegisteredVerifiers() view returns(address[] _verifers)
func (_DepositMainnet *DepositMainnetCallerSession) GetRegisteredVerifiers() ([]common.Address, error) {
	return _DepositMainnet.Contract.GetRegisteredVerifiers(&_DepositMainnet.CallOpts)
}

// GetRollupTreeHeight is a free data retrieval call binding the contract method 0x81b47796.
//
// Solidity: function getRollupTreeHeight() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetRollupTreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getRollupTreeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRollupTreeHeight is a free data retrieval call binding the contract method 0x81b47796.
//
// Solidity: function getRollupTreeHeight() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetRollupTreeHeight() (*big.Int, error) {
	return _DepositMainnet.Contract.GetRollupTreeHeight(&_DepositMainnet.CallOpts)
}

// GetRollupTreeHeight is a free data retrieval call binding the contract method 0x81b47796.
//
// Solidity: function getRollupTreeHeight() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetRollupTreeHeight() (*big.Int, error) {
	return _DepositMainnet.Contract.GetRollupTreeHeight(&_DepositMainnet.CallOpts)
}

// GetRollupVaultRoot is a free data retrieval call binding the contract method 0x8ed31439.
//
// Solidity: function getRollupVaultRoot() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetRollupVaultRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getRollupVaultRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRollupVaultRoot is a free data retrieval call binding the contract method 0x8ed31439.
//
// Solidity: function getRollupVaultRoot() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetRollupVaultRoot() (*big.Int, error) {
	return _DepositMainnet.Contract.GetRollupVaultRoot(&_DepositMainnet.CallOpts)
}

// GetRollupVaultRoot is a free data retrieval call binding the contract method 0x8ed31439.
//
// Solidity: function getRollupVaultRoot() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetRollupVaultRoot() (*big.Int, error) {
	return _DepositMainnet.Contract.GetRollupVaultRoot(&_DepositMainnet.CallOpts)
}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetSequenceNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getSequenceNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetSequenceNumber() (*big.Int, error) {
	return _DepositMainnet.Contract.GetSequenceNumber(&_DepositMainnet.CallOpts)
}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetSequenceNumber() (*big.Int, error) {
	return _DepositMainnet.Contract.GetSequenceNumber(&_DepositMainnet.CallOpts)
}

// GetSubContract is a free data retrieval call binding the contract method 0x2f9014b4.
//
// Solidity: function getSubContract(bytes4 selector) view returns(address)
func (_DepositMainnet *DepositMainnetCaller) GetSubContract(opts *bind.CallOpts, selector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getSubContract", selector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSubContract is a free data retrieval call binding the contract method 0x2f9014b4.
//
// Solidity: function getSubContract(bytes4 selector) view returns(address)
func (_DepositMainnet *DepositMainnetSession) GetSubContract(selector [4]byte) (common.Address, error) {
	return _DepositMainnet.Contract.GetSubContract(&_DepositMainnet.CallOpts, selector)
}

// GetSubContract is a free data retrieval call binding the contract method 0x2f9014b4.
//
// Solidity: function getSubContract(bytes4 selector) view returns(address)
func (_DepositMainnet *DepositMainnetCallerSession) GetSubContract(selector [4]byte) (common.Address, error) {
	return _DepositMainnet.Contract.GetSubContract(&_DepositMainnet.CallOpts, selector)
}

// GetValidiumTreeHeight is a free data retrieval call binding the contract method 0xbb7c3d32.
//
// Solidity: function getValidiumTreeHeight() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetValidiumTreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getValidiumTreeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidiumTreeHeight is a free data retrieval call binding the contract method 0xbb7c3d32.
//
// Solidity: function getValidiumTreeHeight() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetValidiumTreeHeight() (*big.Int, error) {
	return _DepositMainnet.Contract.GetValidiumTreeHeight(&_DepositMainnet.CallOpts)
}

// GetValidiumTreeHeight is a free data retrieval call binding the contract method 0xbb7c3d32.
//
// Solidity: function getValidiumTreeHeight() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetValidiumTreeHeight() (*big.Int, error) {
	return _DepositMainnet.Contract.GetValidiumTreeHeight(&_DepositMainnet.CallOpts)
}

// GetValidiumVaultRoot is a free data retrieval call binding the contract method 0x3a8cc4fc.
//
// Solidity: function getValidiumVaultRoot() view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetValidiumVaultRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getValidiumVaultRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidiumVaultRoot is a free data retrieval call binding the contract method 0x3a8cc4fc.
//
// Solidity: function getValidiumVaultRoot() view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetValidiumVaultRoot() (*big.Int, error) {
	return _DepositMainnet.Contract.GetValidiumVaultRoot(&_DepositMainnet.CallOpts)
}

// GetValidiumVaultRoot is a free data retrieval call binding the contract method 0x3a8cc4fc.
//
// Solidity: function getValidiumVaultRoot() view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetValidiumVaultRoot() (*big.Int, error) {
	return _DepositMainnet.Contract.GetValidiumVaultRoot(&_DepositMainnet.CallOpts)
}

// GetVaultBalance is a free data retrieval call binding the contract method 0x1c6bd544.
//
// Solidity: function getVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetVaultBalance(opts *bind.CallOpts, ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getVaultBalance", ethKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultBalance is a free data retrieval call binding the contract method 0x1c6bd544.
//
// Solidity: function getVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetVaultBalance(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetVaultBalance(&_DepositMainnet.CallOpts, ethKey, assetId, vaultId)
}

// GetVaultBalance is a free data retrieval call binding the contract method 0x1c6bd544.
//
// Solidity: function getVaultBalance(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetVaultBalance(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetVaultBalance(&_DepositMainnet.CallOpts, ethKey, assetId, vaultId)
}

// GetVaultWithdrawalLock is a free data retrieval call binding the contract method 0x9de5c4e5.
//
// Solidity: function getVaultWithdrawalLock(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetVaultWithdrawalLock(opts *bind.CallOpts, ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getVaultWithdrawalLock", ethKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultWithdrawalLock is a free data retrieval call binding the contract method 0x9de5c4e5.
//
// Solidity: function getVaultWithdrawalLock(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetVaultWithdrawalLock(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetVaultWithdrawalLock(&_DepositMainnet.CallOpts, ethKey, assetId, vaultId)
}

// GetVaultWithdrawalLock is a free data retrieval call binding the contract method 0x9de5c4e5.
//
// Solidity: function getVaultWithdrawalLock(address ethKey, uint256 assetId, uint256 vaultId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetVaultWithdrawalLock(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetVaultWithdrawalLock(&_DepositMainnet.CallOpts, ethKey, assetId, vaultId)
}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 ownerKey, uint256 assetId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCaller) GetWithdrawalBalance(opts *bind.CallOpts, ownerKey *big.Int, assetId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "getWithdrawalBalance", ownerKey, assetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 ownerKey, uint256 assetId) view returns(uint256)
func (_DepositMainnet *DepositMainnetSession) GetWithdrawalBalance(ownerKey *big.Int, assetId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetWithdrawalBalance(&_DepositMainnet.CallOpts, ownerKey, assetId)
}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 ownerKey, uint256 assetId) view returns(uint256)
func (_DepositMainnet *DepositMainnetCallerSession) GetWithdrawalBalance(ownerKey *big.Int, assetId *big.Int) (*big.Int, error) {
	return _DepositMainnet.Contract.GetWithdrawalBalance(&_DepositMainnet.CallOpts, ownerKey, assetId)
}

// HandlingContractId is a free data retrieval call binding the contract method 0xd4e878e8.
//
// Solidity: function handlingContractId(bytes4 selector) pure returns(string id)
func (_DepositMainnet *DepositMainnetCaller) HandlingContractId(opts *bind.CallOpts, selector [4]byte) (string, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "handlingContractId", selector)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// HandlingContractId is a free data retrieval call binding the contract method 0xd4e878e8.
//
// Solidity: function handlingContractId(bytes4 selector) pure returns(string id)
func (_DepositMainnet *DepositMainnetSession) HandlingContractId(selector [4]byte) (string, error) {
	return _DepositMainnet.Contract.HandlingContractId(&_DepositMainnet.CallOpts, selector)
}

// HandlingContractId is a free data retrieval call binding the contract method 0xd4e878e8.
//
// Solidity: function handlingContractId(bytes4 selector) pure returns(string id)
func (_DepositMainnet *DepositMainnetCallerSession) HandlingContractId(selector [4]byte) (string, error) {
	return _DepositMainnet.Contract.HandlingContractId(&_DepositMainnet.CallOpts, selector)
}

// IsAssetRegistered is a free data retrieval call binding the contract method 0x049f5ade.
//
// Solidity: function isAssetRegistered(uint256 assetType) view returns(bool)
func (_DepositMainnet *DepositMainnetCaller) IsAssetRegistered(opts *bind.CallOpts, assetType *big.Int) (bool, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "isAssetRegistered", assetType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssetRegistered is a free data retrieval call binding the contract method 0x049f5ade.
//
// Solidity: function isAssetRegistered(uint256 assetType) view returns(bool)
func (_DepositMainnet *DepositMainnetSession) IsAssetRegistered(assetType *big.Int) (bool, error) {
	return _DepositMainnet.Contract.IsAssetRegistered(&_DepositMainnet.CallOpts, assetType)
}

// IsAssetRegistered is a free data retrieval call binding the contract method 0x049f5ade.
//
// Solidity: function isAssetRegistered(uint256 assetType) view returns(bool)
func (_DepositMainnet *DepositMainnetCallerSession) IsAssetRegistered(assetType *big.Int) (bool, error) {
	return _DepositMainnet.Contract.IsAssetRegistered(&_DepositMainnet.CallOpts, assetType)
}

// IsAvailabilityVerifier is a free data retrieval call binding the contract method 0xbd1279ae.
//
// Solidity: function isAvailabilityVerifier(address verifierAddress) view returns(bool)
func (_DepositMainnet *DepositMainnetCaller) IsAvailabilityVerifier(opts *bind.CallOpts, verifierAddress common.Address) (bool, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "isAvailabilityVerifier", verifierAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAvailabilityVerifier is a free data retrieval call binding the contract method 0xbd1279ae.
//
// Solidity: function isAvailabilityVerifier(address verifierAddress) view returns(bool)
func (_DepositMainnet *DepositMainnetSession) IsAvailabilityVerifier(verifierAddress common.Address) (bool, error) {
	return _DepositMainnet.Contract.IsAvailabilityVerifier(&_DepositMainnet.CallOpts, verifierAddress)
}

// IsAvailabilityVerifier is a free data retrieval call binding the contract method 0xbd1279ae.
//
// Solidity: function isAvailabilityVerifier(address verifierAddress) view returns(bool)
func (_DepositMainnet *DepositMainnetCallerSession) IsAvailabilityVerifier(verifierAddress common.Address) (bool, error) {
	return _DepositMainnet.Contract.IsAvailabilityVerifier(&_DepositMainnet.CallOpts, verifierAddress)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_DepositMainnet *DepositMainnetCaller) IsFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "isFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_DepositMainnet *DepositMainnetSession) IsFrozen() (bool, error) {
	return _DepositMainnet.Contract.IsFrozen(&_DepositMainnet.CallOpts)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_DepositMainnet *DepositMainnetCallerSession) IsFrozen() (bool, error) {
	return _DepositMainnet.Contract.IsFrozen(&_DepositMainnet.CallOpts)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address testedOperator) view returns(bool)
func (_DepositMainnet *DepositMainnetCaller) IsOperator(opts *bind.CallOpts, testedOperator common.Address) (bool, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "isOperator", testedOperator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address testedOperator) view returns(bool)
func (_DepositMainnet *DepositMainnetSession) IsOperator(testedOperator common.Address) (bool, error) {
	return _DepositMainnet.Contract.IsOperator(&_DepositMainnet.CallOpts, testedOperator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address testedOperator) view returns(bool)
func (_DepositMainnet *DepositMainnetCallerSession) IsOperator(testedOperator common.Address) (bool, error) {
	return _DepositMainnet.Contract.IsOperator(&_DepositMainnet.CallOpts, testedOperator)
}

// IsStrictVaultBalancePolicy is a free data retrieval call binding the contract method 0x5adf1639.
//
// Solidity: function isStrictVaultBalancePolicy() view returns(bool)
func (_DepositMainnet *DepositMainnetCaller) IsStrictVaultBalancePolicy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "isStrictVaultBalancePolicy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStrictVaultBalancePolicy is a free data retrieval call binding the contract method 0x5adf1639.
//
// Solidity: function isStrictVaultBalancePolicy() view returns(bool)
func (_DepositMainnet *DepositMainnetSession) IsStrictVaultBalancePolicy() (bool, error) {
	return _DepositMainnet.Contract.IsStrictVaultBalancePolicy(&_DepositMainnet.CallOpts)
}

// IsStrictVaultBalancePolicy is a free data retrieval call binding the contract method 0x5adf1639.
//
// Solidity: function isStrictVaultBalancePolicy() view returns(bool)
func (_DepositMainnet *DepositMainnetCallerSession) IsStrictVaultBalancePolicy() (bool, error) {
	return _DepositMainnet.Contract.IsStrictVaultBalancePolicy(&_DepositMainnet.CallOpts)
}

// IsTokenAdmin is a free data retrieval call binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address testedAdmin) view returns(bool)
func (_DepositMainnet *DepositMainnetCaller) IsTokenAdmin(opts *bind.CallOpts, testedAdmin common.Address) (bool, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "isTokenAdmin", testedAdmin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenAdmin is a free data retrieval call binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address testedAdmin) view returns(bool)
func (_DepositMainnet *DepositMainnetSession) IsTokenAdmin(testedAdmin common.Address) (bool, error) {
	return _DepositMainnet.Contract.IsTokenAdmin(&_DepositMainnet.CallOpts, testedAdmin)
}

// IsTokenAdmin is a free data retrieval call binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address testedAdmin) view returns(bool)
func (_DepositMainnet *DepositMainnetCallerSession) IsTokenAdmin(testedAdmin common.Address) (bool, error) {
	return _DepositMainnet.Contract.IsTokenAdmin(&_DepositMainnet.CallOpts, testedAdmin)
}

// IsVaultLocked is a free data retrieval call binding the contract method 0x27c2b36d.
//
// Solidity: function isVaultLocked(address ethKey, uint256 assetId, uint256 vaultId) view returns(bool)
func (_DepositMainnet *DepositMainnetCaller) IsVaultLocked(opts *bind.CallOpts, ethKey common.Address, assetId *big.Int, vaultId *big.Int) (bool, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "isVaultLocked", ethKey, assetId, vaultId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVaultLocked is a free data retrieval call binding the contract method 0x27c2b36d.
//
// Solidity: function isVaultLocked(address ethKey, uint256 assetId, uint256 vaultId) view returns(bool)
func (_DepositMainnet *DepositMainnetSession) IsVaultLocked(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (bool, error) {
	return _DepositMainnet.Contract.IsVaultLocked(&_DepositMainnet.CallOpts, ethKey, assetId, vaultId)
}

// IsVaultLocked is a free data retrieval call binding the contract method 0x27c2b36d.
//
// Solidity: function isVaultLocked(address ethKey, uint256 assetId, uint256 vaultId) view returns(bool)
func (_DepositMainnet *DepositMainnetCallerSession) IsVaultLocked(ethKey common.Address, assetId *big.Int, vaultId *big.Int) (bool, error) {
	return _DepositMainnet.Contract.IsVaultLocked(&_DepositMainnet.CallOpts, ethKey, assetId, vaultId)
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address verifierAddress) view returns(bool)
func (_DepositMainnet *DepositMainnetCaller) IsVerifier(opts *bind.CallOpts, verifierAddress common.Address) (bool, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "isVerifier", verifierAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address verifierAddress) view returns(bool)
func (_DepositMainnet *DepositMainnetSession) IsVerifier(verifierAddress common.Address) (bool, error) {
	return _DepositMainnet.Contract.IsVerifier(&_DepositMainnet.CallOpts, verifierAddress)
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address verifierAddress) view returns(bool)
func (_DepositMainnet *DepositMainnetCallerSession) IsVerifier(verifierAddress common.Address) (bool, error) {
	return _DepositMainnet.Contract.IsVerifier(&_DepositMainnet.CallOpts, verifierAddress)
}

// MainIsGovernor is a free data retrieval call binding the contract method 0x45f5cd97.
//
// Solidity: function mainIsGovernor(address testGovernor) view returns(bool)
func (_DepositMainnet *DepositMainnetCaller) MainIsGovernor(opts *bind.CallOpts, testGovernor common.Address) (bool, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "mainIsGovernor", testGovernor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MainIsGovernor is a free data retrieval call binding the contract method 0x45f5cd97.
//
// Solidity: function mainIsGovernor(address testGovernor) view returns(bool)
func (_DepositMainnet *DepositMainnetSession) MainIsGovernor(testGovernor common.Address) (bool, error) {
	return _DepositMainnet.Contract.MainIsGovernor(&_DepositMainnet.CallOpts, testGovernor)
}

// MainIsGovernor is a free data retrieval call binding the contract method 0x45f5cd97.
//
// Solidity: function mainIsGovernor(address testGovernor) view returns(bool)
func (_DepositMainnet *DepositMainnetCallerSession) MainIsGovernor(testGovernor common.Address) (bool, error) {
	return _DepositMainnet.Contract.MainIsGovernor(&_DepositMainnet.CallOpts, testGovernor)
}

// OrderRegistryAddress is a free data retrieval call binding the contract method 0x9c6a2837.
//
// Solidity: function orderRegistryAddress() view returns(address)
func (_DepositMainnet *DepositMainnetCaller) OrderRegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DepositMainnet.contract.Call(opts, &out, "orderRegistryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrderRegistryAddress is a free data retrieval call binding the contract method 0x9c6a2837.
//
// Solidity: function orderRegistryAddress() view returns(address)
func (_DepositMainnet *DepositMainnetSession) OrderRegistryAddress() (common.Address, error) {
	return _DepositMainnet.Contract.OrderRegistryAddress(&_DepositMainnet.CallOpts)
}

// OrderRegistryAddress is a free data retrieval call binding the contract method 0x9c6a2837.
//
// Solidity: function orderRegistryAddress() view returns(address)
func (_DepositMainnet *DepositMainnetCallerSession) OrderRegistryAddress() (common.Address, error) {
	return _DepositMainnet.Contract.OrderRegistryAddress(&_DepositMainnet.CallOpts)
}

// AnnounceAvailabilityVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x1d078bbb.
//
// Solidity: function announceAvailabilityVerifierRemovalIntent(address verifier) returns()
func (_DepositMainnet *DepositMainnetTransactor) AnnounceAvailabilityVerifierRemovalIntent(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "announceAvailabilityVerifierRemovalIntent", verifier)
}

// AnnounceAvailabilityVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x1d078bbb.
//
// Solidity: function announceAvailabilityVerifierRemovalIntent(address verifier) returns()
func (_DepositMainnet *DepositMainnetSession) AnnounceAvailabilityVerifierRemovalIntent(verifier common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.AnnounceAvailabilityVerifierRemovalIntent(&_DepositMainnet.TransactOpts, verifier)
}

// AnnounceAvailabilityVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x1d078bbb.
//
// Solidity: function announceAvailabilityVerifierRemovalIntent(address verifier) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) AnnounceAvailabilityVerifierRemovalIntent(verifier common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.AnnounceAvailabilityVerifierRemovalIntent(&_DepositMainnet.TransactOpts, verifier)
}

// AnnounceVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x418573b7.
//
// Solidity: function announceVerifierRemovalIntent(address verifier) returns()
func (_DepositMainnet *DepositMainnetTransactor) AnnounceVerifierRemovalIntent(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "announceVerifierRemovalIntent", verifier)
}

// AnnounceVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x418573b7.
//
// Solidity: function announceVerifierRemovalIntent(address verifier) returns()
func (_DepositMainnet *DepositMainnetSession) AnnounceVerifierRemovalIntent(verifier common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.AnnounceVerifierRemovalIntent(&_DepositMainnet.TransactOpts, verifier)
}

// AnnounceVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x418573b7.
//
// Solidity: function announceVerifierRemovalIntent(address verifier) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) AnnounceVerifierRemovalIntent(verifier common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.AnnounceVerifierRemovalIntent(&_DepositMainnet.TransactOpts, verifier)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_DepositMainnet *DepositMainnetTransactor) Deposit(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "deposit", starkKey, assetType, vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_DepositMainnet *DepositMainnetSession) Deposit(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.Deposit(&_DepositMainnet.TransactOpts, starkKey, assetType, vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_DepositMainnet *DepositMainnetTransactorSession) Deposit(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.Deposit(&_DepositMainnet.TransactOpts, starkKey, assetType, vaultId)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactor) Deposit0(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "deposit0", starkKey, assetType, vaultId, quantizedAmount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetSession) Deposit0(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.Deposit0(&_DepositMainnet.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) Deposit0(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.Deposit0(&_DepositMainnet.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_DepositMainnet *DepositMainnetTransactor) DepositCancel(opts *bind.TransactOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "depositCancel", starkKey, assetId, vaultId)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_DepositMainnet *DepositMainnetSession) DepositCancel(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositCancel(&_DepositMainnet.TransactOpts, starkKey, assetId, vaultId)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) DepositCancel(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositCancel(&_DepositMainnet.TransactOpts, starkKey, assetId, vaultId)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x49325bac.
//
// Solidity: function depositERC1155(uint256 starkKey, uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactor) DepositERC1155(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "depositERC1155", starkKey, assetType, tokenId, vaultId, quantizedAmount)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x49325bac.
//
// Solidity: function depositERC1155(uint256 starkKey, uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetSession) DepositERC1155(starkKey *big.Int, assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositERC1155(&_DepositMainnet.TransactOpts, starkKey, assetType, tokenId, vaultId, quantizedAmount)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x49325bac.
//
// Solidity: function depositERC1155(uint256 starkKey, uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) DepositERC1155(starkKey *big.Int, assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositERC1155(&_DepositMainnet.TransactOpts, starkKey, assetType, tokenId, vaultId, quantizedAmount)
}

// DepositERC1155ToVault is a paid mutator transaction binding the contract method 0xa70c4b3b.
//
// Solidity: function depositERC1155ToVault(uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactor) DepositERC1155ToVault(opts *bind.TransactOpts, assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "depositERC1155ToVault", assetType, tokenId, vaultId, quantizedAmount)
}

// DepositERC1155ToVault is a paid mutator transaction binding the contract method 0xa70c4b3b.
//
// Solidity: function depositERC1155ToVault(uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetSession) DepositERC1155ToVault(assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositERC1155ToVault(&_DepositMainnet.TransactOpts, assetType, tokenId, vaultId, quantizedAmount)
}

// DepositERC1155ToVault is a paid mutator transaction binding the contract method 0xa70c4b3b.
//
// Solidity: function depositERC1155ToVault(uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) DepositERC1155ToVault(assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositERC1155ToVault(&_DepositMainnet.TransactOpts, assetType, tokenId, vaultId, quantizedAmount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x9ed17084.
//
// Solidity: function depositERC20(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactor) DepositERC20(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "depositERC20", starkKey, assetType, vaultId, quantizedAmount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x9ed17084.
//
// Solidity: function depositERC20(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetSession) DepositERC20(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositERC20(&_DepositMainnet.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x9ed17084.
//
// Solidity: function depositERC20(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) DepositERC20(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositERC20(&_DepositMainnet.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// DepositERC20ToVault is a paid mutator transaction binding the contract method 0x1d133002.
//
// Solidity: function depositERC20ToVault(uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactor) DepositERC20ToVault(opts *bind.TransactOpts, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "depositERC20ToVault", assetType, vaultId, quantizedAmount)
}

// DepositERC20ToVault is a paid mutator transaction binding the contract method 0x1d133002.
//
// Solidity: function depositERC20ToVault(uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetSession) DepositERC20ToVault(assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositERC20ToVault(&_DepositMainnet.TransactOpts, assetType, vaultId, quantizedAmount)
}

// DepositERC20ToVault is a paid mutator transaction binding the contract method 0x1d133002.
//
// Solidity: function depositERC20ToVault(uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) DepositERC20ToVault(assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositERC20ToVault(&_DepositMainnet.TransactOpts, assetType, vaultId, quantizedAmount)
}

// DepositEth is a paid mutator transaction binding the contract method 0x6ce5d957.
//
// Solidity: function depositEth(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_DepositMainnet *DepositMainnetTransactor) DepositEth(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "depositEth", starkKey, assetType, vaultId)
}

// DepositEth is a paid mutator transaction binding the contract method 0x6ce5d957.
//
// Solidity: function depositEth(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_DepositMainnet *DepositMainnetSession) DepositEth(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositEth(&_DepositMainnet.TransactOpts, starkKey, assetType, vaultId)
}

// DepositEth is a paid mutator transaction binding the contract method 0x6ce5d957.
//
// Solidity: function depositEth(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_DepositMainnet *DepositMainnetTransactorSession) DepositEth(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositEth(&_DepositMainnet.TransactOpts, starkKey, assetType, vaultId)
}

// DepositEthToVault is a paid mutator transaction binding the contract method 0xa3c71fde.
//
// Solidity: function depositEthToVault(uint256 assetType, uint256 vaultId) payable returns()
func (_DepositMainnet *DepositMainnetTransactor) DepositEthToVault(opts *bind.TransactOpts, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "depositEthToVault", assetType, vaultId)
}

// DepositEthToVault is a paid mutator transaction binding the contract method 0xa3c71fde.
//
// Solidity: function depositEthToVault(uint256 assetType, uint256 vaultId) payable returns()
func (_DepositMainnet *DepositMainnetSession) DepositEthToVault(assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositEthToVault(&_DepositMainnet.TransactOpts, assetType, vaultId)
}

// DepositEthToVault is a paid mutator transaction binding the contract method 0xa3c71fde.
//
// Solidity: function depositEthToVault(uint256 assetType, uint256 vaultId) payable returns()
func (_DepositMainnet *DepositMainnetTransactorSession) DepositEthToVault(assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositEthToVault(&_DepositMainnet.TransactOpts, assetType, vaultId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_DepositMainnet *DepositMainnetTransactor) DepositNft(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "depositNft", starkKey, assetType, vaultId, tokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_DepositMainnet *DepositMainnetSession) DepositNft(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositNft(&_DepositMainnet.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) DepositNft(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositNft(&_DepositMainnet.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_DepositMainnet *DepositMainnetTransactor) DepositNftReclaim(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "depositNftReclaim", starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_DepositMainnet *DepositMainnetSession) DepositNftReclaim(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositNftReclaim(&_DepositMainnet.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) DepositNftReclaim(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositNftReclaim(&_DepositMainnet.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId) returns()
func (_DepositMainnet *DepositMainnetTransactor) DepositReclaim(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "depositReclaim", starkKey, assetType, vaultId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId) returns()
func (_DepositMainnet *DepositMainnetSession) DepositReclaim(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositReclaim(&_DepositMainnet.TransactOpts, starkKey, assetType, vaultId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) DepositReclaim(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositReclaim(&_DepositMainnet.TransactOpts, starkKey, assetType, vaultId)
}

// DepositWithTokenId is a paid mutator transaction binding the contract method 0x29e411ac.
//
// Solidity: function depositWithTokenId(uint256 starkKey, uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactor) DepositWithTokenId(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "depositWithTokenId", starkKey, assetType, tokenId, vaultId, quantizedAmount)
}

// DepositWithTokenId is a paid mutator transaction binding the contract method 0x29e411ac.
//
// Solidity: function depositWithTokenId(uint256 starkKey, uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetSession) DepositWithTokenId(starkKey *big.Int, assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositWithTokenId(&_DepositMainnet.TransactOpts, starkKey, assetType, tokenId, vaultId, quantizedAmount)
}

// DepositWithTokenId is a paid mutator transaction binding the contract method 0x29e411ac.
//
// Solidity: function depositWithTokenId(uint256 starkKey, uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) DepositWithTokenId(starkKey *big.Int, assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositWithTokenId(&_DepositMainnet.TransactOpts, starkKey, assetType, tokenId, vaultId, quantizedAmount)
}

// DepositWithTokenIdReclaim is a paid mutator transaction binding the contract method 0xbe2b1105.
//
// Solidity: function depositWithTokenIdReclaim(uint256 starkKey, uint256 assetType, uint256 tokenId, uint256 vaultId) returns()
func (_DepositMainnet *DepositMainnetTransactor) DepositWithTokenIdReclaim(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, tokenId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "depositWithTokenIdReclaim", starkKey, assetType, tokenId, vaultId)
}

// DepositWithTokenIdReclaim is a paid mutator transaction binding the contract method 0xbe2b1105.
//
// Solidity: function depositWithTokenIdReclaim(uint256 starkKey, uint256 assetType, uint256 tokenId, uint256 vaultId) returns()
func (_DepositMainnet *DepositMainnetSession) DepositWithTokenIdReclaim(starkKey *big.Int, assetType *big.Int, tokenId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositWithTokenIdReclaim(&_DepositMainnet.TransactOpts, starkKey, assetType, tokenId, vaultId)
}

// DepositWithTokenIdReclaim is a paid mutator transaction binding the contract method 0xbe2b1105.
//
// Solidity: function depositWithTokenIdReclaim(uint256 starkKey, uint256 assetType, uint256 tokenId, uint256 vaultId) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) DepositWithTokenIdReclaim(starkKey *big.Int, assetType *big.Int, tokenId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.DepositWithTokenIdReclaim(&_DepositMainnet.TransactOpts, starkKey, assetType, tokenId, vaultId)
}

// Escape is a paid mutator transaction binding the contract method 0x9e3adac4.
//
// Solidity: function escape(uint256 ownerKey, uint256 vaultId, uint256 assetId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactor) Escape(opts *bind.TransactOpts, ownerKey *big.Int, vaultId *big.Int, assetId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "escape", ownerKey, vaultId, assetId, quantizedAmount)
}

// Escape is a paid mutator transaction binding the contract method 0x9e3adac4.
//
// Solidity: function escape(uint256 ownerKey, uint256 vaultId, uint256 assetId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetSession) Escape(ownerKey *big.Int, vaultId *big.Int, assetId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.Escape(&_DepositMainnet.TransactOpts, ownerKey, vaultId, assetId, quantizedAmount)
}

// Escape is a paid mutator transaction binding the contract method 0x9e3adac4.
//
// Solidity: function escape(uint256 ownerKey, uint256 vaultId, uint256 assetId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) Escape(ownerKey *big.Int, vaultId *big.Int, assetId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.Escape(&_DepositMainnet.TransactOpts, ownerKey, vaultId, assetId, quantizedAmount)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x93c1e466.
//
// Solidity: function freezeRequest(uint256 ownerKey, uint256 vaultId) returns()
func (_DepositMainnet *DepositMainnetTransactor) FreezeRequest(opts *bind.TransactOpts, ownerKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "freezeRequest", ownerKey, vaultId)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x93c1e466.
//
// Solidity: function freezeRequest(uint256 ownerKey, uint256 vaultId) returns()
func (_DepositMainnet *DepositMainnetSession) FreezeRequest(ownerKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.FreezeRequest(&_DepositMainnet.TransactOpts, ownerKey, vaultId)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x93c1e466.
//
// Solidity: function freezeRequest(uint256 ownerKey, uint256 vaultId) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) FreezeRequest(ownerKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.FreezeRequest(&_DepositMainnet.TransactOpts, ownerKey, vaultId)
}

// FullWithdrawalRequest is a paid mutator transaction binding the contract method 0xa93310c4.
//
// Solidity: function fullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) returns()
func (_DepositMainnet *DepositMainnetTransactor) FullWithdrawalRequest(opts *bind.TransactOpts, ownerKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "fullWithdrawalRequest", ownerKey, vaultId)
}

// FullWithdrawalRequest is a paid mutator transaction binding the contract method 0xa93310c4.
//
// Solidity: function fullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) returns()
func (_DepositMainnet *DepositMainnetSession) FullWithdrawalRequest(ownerKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.FullWithdrawalRequest(&_DepositMainnet.TransactOpts, ownerKey, vaultId)
}

// FullWithdrawalRequest is a paid mutator transaction binding the contract method 0xa93310c4.
//
// Solidity: function fullWithdrawalRequest(uint256 ownerKey, uint256 vaultId) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) FullWithdrawalRequest(ownerKey *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.FullWithdrawalRequest(&_DepositMainnet.TransactOpts, ownerKey, vaultId)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_DepositMainnet *DepositMainnetTransactor) Initialize(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "initialize", data)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_DepositMainnet *DepositMainnetSession) Initialize(data []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.Initialize(&_DepositMainnet.TransactOpts, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) Initialize(data []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.Initialize(&_DepositMainnet.TransactOpts, data)
}

// LockVault is a paid mutator transaction binding the contract method 0xe8d28a9d.
//
// Solidity: function lockVault(uint256 assetId, uint256 vaultId, uint256 lockTime) returns()
func (_DepositMainnet *DepositMainnetTransactor) LockVault(opts *bind.TransactOpts, assetId *big.Int, vaultId *big.Int, lockTime *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "lockVault", assetId, vaultId, lockTime)
}

// LockVault is a paid mutator transaction binding the contract method 0xe8d28a9d.
//
// Solidity: function lockVault(uint256 assetId, uint256 vaultId, uint256 lockTime) returns()
func (_DepositMainnet *DepositMainnetSession) LockVault(assetId *big.Int, vaultId *big.Int, lockTime *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.LockVault(&_DepositMainnet.TransactOpts, assetId, vaultId, lockTime)
}

// LockVault is a paid mutator transaction binding the contract method 0xe8d28a9d.
//
// Solidity: function lockVault(uint256 assetId, uint256 vaultId, uint256 lockTime) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) LockVault(assetId *big.Int, vaultId *big.Int, lockTime *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.LockVault(&_DepositMainnet.TransactOpts, assetId, vaultId, lockTime)
}

// MainAcceptGovernance is a paid mutator transaction binding the contract method 0x28700a15.
//
// Solidity: function mainAcceptGovernance() returns()
func (_DepositMainnet *DepositMainnetTransactor) MainAcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "mainAcceptGovernance")
}

// MainAcceptGovernance is a paid mutator transaction binding the contract method 0x28700a15.
//
// Solidity: function mainAcceptGovernance() returns()
func (_DepositMainnet *DepositMainnetSession) MainAcceptGovernance() (*types.Transaction, error) {
	return _DepositMainnet.Contract.MainAcceptGovernance(&_DepositMainnet.TransactOpts)
}

// MainAcceptGovernance is a paid mutator transaction binding the contract method 0x28700a15.
//
// Solidity: function mainAcceptGovernance() returns()
func (_DepositMainnet *DepositMainnetTransactorSession) MainAcceptGovernance() (*types.Transaction, error) {
	return _DepositMainnet.Contract.MainAcceptGovernance(&_DepositMainnet.TransactOpts)
}

// MainCancelNomination is a paid mutator transaction binding the contract method 0x72eb3688.
//
// Solidity: function mainCancelNomination() returns()
func (_DepositMainnet *DepositMainnetTransactor) MainCancelNomination(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "mainCancelNomination")
}

// MainCancelNomination is a paid mutator transaction binding the contract method 0x72eb3688.
//
// Solidity: function mainCancelNomination() returns()
func (_DepositMainnet *DepositMainnetSession) MainCancelNomination() (*types.Transaction, error) {
	return _DepositMainnet.Contract.MainCancelNomination(&_DepositMainnet.TransactOpts)
}

// MainCancelNomination is a paid mutator transaction binding the contract method 0x72eb3688.
//
// Solidity: function mainCancelNomination() returns()
func (_DepositMainnet *DepositMainnetTransactorSession) MainCancelNomination() (*types.Transaction, error) {
	return _DepositMainnet.Contract.MainCancelNomination(&_DepositMainnet.TransactOpts)
}

// MainNominateNewGovernor is a paid mutator transaction binding the contract method 0x8c4bce1c.
//
// Solidity: function mainNominateNewGovernor(address newGovernor) returns()
func (_DepositMainnet *DepositMainnetTransactor) MainNominateNewGovernor(opts *bind.TransactOpts, newGovernor common.Address) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "mainNominateNewGovernor", newGovernor)
}

// MainNominateNewGovernor is a paid mutator transaction binding the contract method 0x8c4bce1c.
//
// Solidity: function mainNominateNewGovernor(address newGovernor) returns()
func (_DepositMainnet *DepositMainnetSession) MainNominateNewGovernor(newGovernor common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.MainNominateNewGovernor(&_DepositMainnet.TransactOpts, newGovernor)
}

// MainNominateNewGovernor is a paid mutator transaction binding the contract method 0x8c4bce1c.
//
// Solidity: function mainNominateNewGovernor(address newGovernor) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) MainNominateNewGovernor(newGovernor common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.MainNominateNewGovernor(&_DepositMainnet.TransactOpts, newGovernor)
}

// MainRemoveGovernor is a paid mutator transaction binding the contract method 0xa1cc921e.
//
// Solidity: function mainRemoveGovernor(address governorForRemoval) returns()
func (_DepositMainnet *DepositMainnetTransactor) MainRemoveGovernor(opts *bind.TransactOpts, governorForRemoval common.Address) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "mainRemoveGovernor", governorForRemoval)
}

// MainRemoveGovernor is a paid mutator transaction binding the contract method 0xa1cc921e.
//
// Solidity: function mainRemoveGovernor(address governorForRemoval) returns()
func (_DepositMainnet *DepositMainnetSession) MainRemoveGovernor(governorForRemoval common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.MainRemoveGovernor(&_DepositMainnet.TransactOpts, governorForRemoval)
}

// MainRemoveGovernor is a paid mutator transaction binding the contract method 0xa1cc921e.
//
// Solidity: function mainRemoveGovernor(address governorForRemoval) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) MainRemoveGovernor(governorForRemoval common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.MainRemoveGovernor(&_DepositMainnet.TransactOpts, governorForRemoval)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_DepositMainnet *DepositMainnetTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, operator common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "onERC1155BatchReceived", operator, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_DepositMainnet *DepositMainnetSession) OnERC1155BatchReceived(operator common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.OnERC1155BatchReceived(&_DepositMainnet.TransactOpts, operator, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_DepositMainnet *DepositMainnetTransactorSession) OnERC1155BatchReceived(operator common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.OnERC1155BatchReceived(&_DepositMainnet.TransactOpts, operator, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address , uint256 , uint256 , bytes ) returns(bytes4)
func (_DepositMainnet *DepositMainnetTransactor) OnERC1155Received(opts *bind.TransactOpts, operator common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "onERC1155Received", operator, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address , uint256 , uint256 , bytes ) returns(bytes4)
func (_DepositMainnet *DepositMainnetSession) OnERC1155Received(operator common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.OnERC1155Received(&_DepositMainnet.TransactOpts, operator, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address , uint256 , uint256 , bytes ) returns(bytes4)
func (_DepositMainnet *DepositMainnetTransactorSession) OnERC1155Received(operator common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.OnERC1155Received(&_DepositMainnet.TransactOpts, operator, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address , uint256 , bytes ) returns(bytes4)
func (_DepositMainnet *DepositMainnetTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "onERC721Received", operator, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address , uint256 , bytes ) returns(bytes4)
func (_DepositMainnet *DepositMainnetSession) OnERC721Received(operator common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.OnERC721Received(&_DepositMainnet.TransactOpts, operator, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address , uint256 , bytes ) returns(bytes4)
func (_DepositMainnet *DepositMainnetTransactorSession) OnERC721Received(operator common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.OnERC721Received(&_DepositMainnet.TransactOpts, operator, arg1, arg2, arg3)
}

// RegisterAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbdb75785.
//
// Solidity: function registerAvailabilityVerifier(address verifier, string identifier) returns()
func (_DepositMainnet *DepositMainnetTransactor) RegisterAvailabilityVerifier(opts *bind.TransactOpts, verifier common.Address, identifier string) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "registerAvailabilityVerifier", verifier, identifier)
}

// RegisterAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbdb75785.
//
// Solidity: function registerAvailabilityVerifier(address verifier, string identifier) returns()
func (_DepositMainnet *DepositMainnetSession) RegisterAvailabilityVerifier(verifier common.Address, identifier string) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RegisterAvailabilityVerifier(&_DepositMainnet.TransactOpts, verifier, identifier)
}

// RegisterAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbdb75785.
//
// Solidity: function registerAvailabilityVerifier(address verifier, string identifier) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) RegisterAvailabilityVerifier(verifier common.Address, identifier string) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RegisterAvailabilityVerifier(&_DepositMainnet.TransactOpts, verifier, identifier)
}

// RegisterEthAddress is a paid mutator transaction binding the contract method 0xbea84187.
//
// Solidity: function registerEthAddress(address ethKey, uint256 starkKey, bytes starkSignature) returns()
func (_DepositMainnet *DepositMainnetTransactor) RegisterEthAddress(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "registerEthAddress", ethKey, starkKey, starkSignature)
}

// RegisterEthAddress is a paid mutator transaction binding the contract method 0xbea84187.
//
// Solidity: function registerEthAddress(address ethKey, uint256 starkKey, bytes starkSignature) returns()
func (_DepositMainnet *DepositMainnetSession) RegisterEthAddress(ethKey common.Address, starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RegisterEthAddress(&_DepositMainnet.TransactOpts, ethKey, starkKey, starkSignature)
}

// RegisterEthAddress is a paid mutator transaction binding the contract method 0xbea84187.
//
// Solidity: function registerEthAddress(address ethKey, uint256 starkKey, bytes starkSignature) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) RegisterEthAddress(ethKey common.Address, starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RegisterEthAddress(&_DepositMainnet.TransactOpts, ethKey, starkKey, starkSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_DepositMainnet *DepositMainnetTransactor) RegisterOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "registerOperator", newOperator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_DepositMainnet *DepositMainnetSession) RegisterOperator(newOperator common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RegisterOperator(&_DepositMainnet.TransactOpts, newOperator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) RegisterOperator(newOperator common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RegisterOperator(&_DepositMainnet.TransactOpts, newOperator)
}

// RegisterSender is a paid mutator transaction binding the contract method 0x86aeb445.
//
// Solidity: function registerSender(uint256 starkKey, bytes starkSignature) returns()
func (_DepositMainnet *DepositMainnetTransactor) RegisterSender(opts *bind.TransactOpts, starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "registerSender", starkKey, starkSignature)
}

// RegisterSender is a paid mutator transaction binding the contract method 0x86aeb445.
//
// Solidity: function registerSender(uint256 starkKey, bytes starkSignature) returns()
func (_DepositMainnet *DepositMainnetSession) RegisterSender(starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RegisterSender(&_DepositMainnet.TransactOpts, starkKey, starkSignature)
}

// RegisterSender is a paid mutator transaction binding the contract method 0x86aeb445.
//
// Solidity: function registerSender(uint256 starkKey, bytes starkSignature) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) RegisterSender(starkKey *big.Int, starkSignature []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RegisterSender(&_DepositMainnet.TransactOpts, starkKey, starkSignature)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo) returns()
func (_DepositMainnet *DepositMainnetTransactor) RegisterToken(opts *bind.TransactOpts, assetType *big.Int, assetInfo []byte) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "registerToken", assetType, assetInfo)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo) returns()
func (_DepositMainnet *DepositMainnetSession) RegisterToken(assetType *big.Int, assetInfo []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RegisterToken(&_DepositMainnet.TransactOpts, assetType, assetInfo)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) RegisterToken(assetType *big.Int, assetInfo []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RegisterToken(&_DepositMainnet.TransactOpts, assetType, assetInfo)
}

// RegisterToken0 is a paid mutator transaction binding the contract method 0xd88d8b38.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo, uint256 quantum) returns()
func (_DepositMainnet *DepositMainnetTransactor) RegisterToken0(opts *bind.TransactOpts, assetType *big.Int, assetInfo []byte, quantum *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "registerToken0", assetType, assetInfo, quantum)
}

// RegisterToken0 is a paid mutator transaction binding the contract method 0xd88d8b38.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo, uint256 quantum) returns()
func (_DepositMainnet *DepositMainnetSession) RegisterToken0(assetType *big.Int, assetInfo []byte, quantum *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RegisterToken0(&_DepositMainnet.TransactOpts, assetType, assetInfo, quantum)
}

// RegisterToken0 is a paid mutator transaction binding the contract method 0xd88d8b38.
//
// Solidity: function registerToken(uint256 assetType, bytes assetInfo, uint256 quantum) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) RegisterToken0(assetType *big.Int, assetInfo []byte, quantum *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RegisterToken0(&_DepositMainnet.TransactOpts, assetType, assetInfo, quantum)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address newAdmin) returns()
func (_DepositMainnet *DepositMainnetTransactor) RegisterTokenAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "registerTokenAdmin", newAdmin)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address newAdmin) returns()
func (_DepositMainnet *DepositMainnetSession) RegisterTokenAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RegisterTokenAdmin(&_DepositMainnet.TransactOpts, newAdmin)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address newAdmin) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) RegisterTokenAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RegisterTokenAdmin(&_DepositMainnet.TransactOpts, newAdmin)
}

// RegisterVerifier is a paid mutator transaction binding the contract method 0x3776fe2a.
//
// Solidity: function registerVerifier(address verifier, string identifier) returns()
func (_DepositMainnet *DepositMainnetTransactor) RegisterVerifier(opts *bind.TransactOpts, verifier common.Address, identifier string) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "registerVerifier", verifier, identifier)
}

// RegisterVerifier is a paid mutator transaction binding the contract method 0x3776fe2a.
//
// Solidity: function registerVerifier(address verifier, string identifier) returns()
func (_DepositMainnet *DepositMainnetSession) RegisterVerifier(verifier common.Address, identifier string) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RegisterVerifier(&_DepositMainnet.TransactOpts, verifier, identifier)
}

// RegisterVerifier is a paid mutator transaction binding the contract method 0x3776fe2a.
//
// Solidity: function registerVerifier(address verifier, string identifier) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) RegisterVerifier(verifier common.Address, identifier string) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RegisterVerifier(&_DepositMainnet.TransactOpts, verifier, identifier)
}

// RemoveAvailabilityVerifier is a paid mutator transaction binding the contract method 0xb1e640bf.
//
// Solidity: function removeAvailabilityVerifier(address verifier) returns()
func (_DepositMainnet *DepositMainnetTransactor) RemoveAvailabilityVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "removeAvailabilityVerifier", verifier)
}

// RemoveAvailabilityVerifier is a paid mutator transaction binding the contract method 0xb1e640bf.
//
// Solidity: function removeAvailabilityVerifier(address verifier) returns()
func (_DepositMainnet *DepositMainnetSession) RemoveAvailabilityVerifier(verifier common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RemoveAvailabilityVerifier(&_DepositMainnet.TransactOpts, verifier)
}

// RemoveAvailabilityVerifier is a paid mutator transaction binding the contract method 0xb1e640bf.
//
// Solidity: function removeAvailabilityVerifier(address verifier) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) RemoveAvailabilityVerifier(verifier common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RemoveAvailabilityVerifier(&_DepositMainnet.TransactOpts, verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_DepositMainnet *DepositMainnetTransactor) RemoveVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "removeVerifier", verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_DepositMainnet *DepositMainnetSession) RemoveVerifier(verifier common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RemoveVerifier(&_DepositMainnet.TransactOpts, verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) RemoveVerifier(verifier common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.RemoveVerifier(&_DepositMainnet.TransactOpts, verifier)
}

// SetDefaultVaultWithdrawalLock is a paid mutator transaction binding the contract method 0xd4181dea.
//
// Solidity: function setDefaultVaultWithdrawalLock(uint256 newDefaultTime) returns()
func (_DepositMainnet *DepositMainnetTransactor) SetDefaultVaultWithdrawalLock(opts *bind.TransactOpts, newDefaultTime *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "setDefaultVaultWithdrawalLock", newDefaultTime)
}

// SetDefaultVaultWithdrawalLock is a paid mutator transaction binding the contract method 0xd4181dea.
//
// Solidity: function setDefaultVaultWithdrawalLock(uint256 newDefaultTime) returns()
func (_DepositMainnet *DepositMainnetSession) SetDefaultVaultWithdrawalLock(newDefaultTime *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.SetDefaultVaultWithdrawalLock(&_DepositMainnet.TransactOpts, newDefaultTime)
}

// SetDefaultVaultWithdrawalLock is a paid mutator transaction binding the contract method 0xd4181dea.
//
// Solidity: function setDefaultVaultWithdrawalLock(uint256 newDefaultTime) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) SetDefaultVaultWithdrawalLock(newDefaultTime *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.SetDefaultVaultWithdrawalLock(&_DepositMainnet.TransactOpts, newDefaultTime)
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_DepositMainnet *DepositMainnetTransactor) UnFreeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "unFreeze")
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_DepositMainnet *DepositMainnetSession) UnFreeze() (*types.Transaction, error) {
	return _DepositMainnet.Contract.UnFreeze(&_DepositMainnet.TransactOpts)
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_DepositMainnet *DepositMainnetTransactorSession) UnFreeze() (*types.Transaction, error) {
	return _DepositMainnet.Contract.UnFreeze(&_DepositMainnet.TransactOpts)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_DepositMainnet *DepositMainnetTransactor) UnregisterOperator(opts *bind.TransactOpts, removedOperator common.Address) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "unregisterOperator", removedOperator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_DepositMainnet *DepositMainnetSession) UnregisterOperator(removedOperator common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.UnregisterOperator(&_DepositMainnet.TransactOpts, removedOperator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) UnregisterOperator(removedOperator common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.UnregisterOperator(&_DepositMainnet.TransactOpts, removedOperator)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address oldAdmin) returns()
func (_DepositMainnet *DepositMainnetTransactor) UnregisterTokenAdmin(opts *bind.TransactOpts, oldAdmin common.Address) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "unregisterTokenAdmin", oldAdmin)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address oldAdmin) returns()
func (_DepositMainnet *DepositMainnetSession) UnregisterTokenAdmin(oldAdmin common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.UnregisterTokenAdmin(&_DepositMainnet.TransactOpts, oldAdmin)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address oldAdmin) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) UnregisterTokenAdmin(oldAdmin common.Address) (*types.Transaction, error) {
	return _DepositMainnet.Contract.UnregisterTokenAdmin(&_DepositMainnet.TransactOpts, oldAdmin)
}

// UpdateImplementationActivationTime is a paid mutator transaction binding the contract method 0x02a93fae.
//
// Solidity: function updateImplementationActivationTime(address implementation, bytes data, bool finalize) returns()
func (_DepositMainnet *DepositMainnetTransactor) UpdateImplementationActivationTime(opts *bind.TransactOpts, implementation common.Address, data []byte, finalize bool) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "updateImplementationActivationTime", implementation, data, finalize)
}

// UpdateImplementationActivationTime is a paid mutator transaction binding the contract method 0x02a93fae.
//
// Solidity: function updateImplementationActivationTime(address implementation, bytes data, bool finalize) returns()
func (_DepositMainnet *DepositMainnetSession) UpdateImplementationActivationTime(implementation common.Address, data []byte, finalize bool) (*types.Transaction, error) {
	return _DepositMainnet.Contract.UpdateImplementationActivationTime(&_DepositMainnet.TransactOpts, implementation, data, finalize)
}

// UpdateImplementationActivationTime is a paid mutator transaction binding the contract method 0x02a93fae.
//
// Solidity: function updateImplementationActivationTime(address implementation, bytes data, bool finalize) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) UpdateImplementationActivationTime(implementation common.Address, data []byte, finalize bool) (*types.Transaction, error) {
	return _DepositMainnet.Contract.UpdateImplementationActivationTime(&_DepositMainnet.TransactOpts, implementation, data, finalize)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] publicInput, uint256[] applicationData) returns()
func (_DepositMainnet *DepositMainnetTransactor) UpdateState(opts *bind.TransactOpts, publicInput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "updateState", publicInput, applicationData)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] publicInput, uint256[] applicationData) returns()
func (_DepositMainnet *DepositMainnetSession) UpdateState(publicInput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.UpdateState(&_DepositMainnet.TransactOpts, publicInput, applicationData)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] publicInput, uint256[] applicationData) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) UpdateState(publicInput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.UpdateState(&_DepositMainnet.TransactOpts, publicInput, applicationData)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 ownerKey, uint256 assetType) returns()
func (_DepositMainnet *DepositMainnetTransactor) Withdraw(opts *bind.TransactOpts, ownerKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "withdraw", ownerKey, assetType)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 ownerKey, uint256 assetType) returns()
func (_DepositMainnet *DepositMainnetSession) Withdraw(ownerKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.Withdraw(&_DepositMainnet.TransactOpts, ownerKey, assetType)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 ownerKey, uint256 assetType) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) Withdraw(ownerKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.Withdraw(&_DepositMainnet.TransactOpts, ownerKey, assetType)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 ownerKey, uint256 assetType, bytes mintingBlob) returns()
func (_DepositMainnet *DepositMainnetTransactor) WithdrawAndMint(opts *bind.TransactOpts, ownerKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "withdrawAndMint", ownerKey, assetType, mintingBlob)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 ownerKey, uint256 assetType, bytes mintingBlob) returns()
func (_DepositMainnet *DepositMainnetSession) WithdrawAndMint(ownerKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.WithdrawAndMint(&_DepositMainnet.TransactOpts, ownerKey, assetType, mintingBlob)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 ownerKey, uint256 assetType, bytes mintingBlob) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) WithdrawAndMint(ownerKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.WithdrawAndMint(&_DepositMainnet.TransactOpts, ownerKey, assetType, mintingBlob)
}

// WithdrawErc1155FromVault is a paid mutator transaction binding the contract method 0xa90f34a3.
//
// Solidity: function withdrawErc1155FromVault(uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactor) WithdrawErc1155FromVault(opts *bind.TransactOpts, assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "withdrawErc1155FromVault", assetType, tokenId, vaultId, quantizedAmount)
}

// WithdrawErc1155FromVault is a paid mutator transaction binding the contract method 0xa90f34a3.
//
// Solidity: function withdrawErc1155FromVault(uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetSession) WithdrawErc1155FromVault(assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.WithdrawErc1155FromVault(&_DepositMainnet.TransactOpts, assetType, tokenId, vaultId, quantizedAmount)
}

// WithdrawErc1155FromVault is a paid mutator transaction binding the contract method 0xa90f34a3.
//
// Solidity: function withdrawErc1155FromVault(uint256 assetType, uint256 tokenId, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) WithdrawErc1155FromVault(assetType *big.Int, tokenId *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.WithdrawErc1155FromVault(&_DepositMainnet.TransactOpts, assetType, tokenId, vaultId, quantizedAmount)
}

// WithdrawFromVault is a paid mutator transaction binding the contract method 0xbf422b2e.
//
// Solidity: function withdrawFromVault(uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactor) WithdrawFromVault(opts *bind.TransactOpts, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "withdrawFromVault", assetType, vaultId, quantizedAmount)
}

// WithdrawFromVault is a paid mutator transaction binding the contract method 0xbf422b2e.
//
// Solidity: function withdrawFromVault(uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetSession) WithdrawFromVault(assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.WithdrawFromVault(&_DepositMainnet.TransactOpts, assetType, vaultId, quantizedAmount)
}

// WithdrawFromVault is a paid mutator transaction binding the contract method 0xbf422b2e.
//
// Solidity: function withdrawFromVault(uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) WithdrawFromVault(assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.WithdrawFromVault(&_DepositMainnet.TransactOpts, assetType, vaultId, quantizedAmount)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 ownerKey, uint256 assetType, uint256 tokenId) returns()
func (_DepositMainnet *DepositMainnetTransactor) WithdrawNft(opts *bind.TransactOpts, ownerKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "withdrawNft", ownerKey, assetType, tokenId)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 ownerKey, uint256 assetType, uint256 tokenId) returns()
func (_DepositMainnet *DepositMainnetSession) WithdrawNft(ownerKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.WithdrawNft(&_DepositMainnet.TransactOpts, ownerKey, assetType, tokenId)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 ownerKey, uint256 assetType, uint256 tokenId) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) WithdrawNft(ownerKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.WithdrawNft(&_DepositMainnet.TransactOpts, ownerKey, assetType, tokenId)
}

// WithdrawWithTokenId is a paid mutator transaction binding the contract method 0x64d84842.
//
// Solidity: function withdrawWithTokenId(uint256 ownerKey, uint256 assetType, uint256 tokenId) returns()
func (_DepositMainnet *DepositMainnetTransactor) WithdrawWithTokenId(opts *bind.TransactOpts, ownerKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.contract.Transact(opts, "withdrawWithTokenId", ownerKey, assetType, tokenId)
}

// WithdrawWithTokenId is a paid mutator transaction binding the contract method 0x64d84842.
//
// Solidity: function withdrawWithTokenId(uint256 ownerKey, uint256 assetType, uint256 tokenId) returns()
func (_DepositMainnet *DepositMainnetSession) WithdrawWithTokenId(ownerKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.WithdrawWithTokenId(&_DepositMainnet.TransactOpts, ownerKey, assetType, tokenId)
}

// WithdrawWithTokenId is a paid mutator transaction binding the contract method 0x64d84842.
//
// Solidity: function withdrawWithTokenId(uint256 ownerKey, uint256 assetType, uint256 tokenId) returns()
func (_DepositMainnet *DepositMainnetTransactorSession) WithdrawWithTokenId(ownerKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _DepositMainnet.Contract.WithdrawWithTokenId(&_DepositMainnet.TransactOpts, ownerKey, assetType, tokenId)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DepositMainnet *DepositMainnetTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _DepositMainnet.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DepositMainnet *DepositMainnetSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.Fallback(&_DepositMainnet.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DepositMainnet *DepositMainnetTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DepositMainnet.Contract.Fallback(&_DepositMainnet.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DepositMainnet *DepositMainnetTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositMainnet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DepositMainnet *DepositMainnetSession) Receive() (*types.Transaction, error) {
	return _DepositMainnet.Contract.Receive(&_DepositMainnet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DepositMainnet *DepositMainnetTransactorSession) Receive() (*types.Transaction, error) {
	return _DepositMainnet.Contract.Receive(&_DepositMainnet.TransactOpts)
}

// DepositMainnetImplementationActivationRescheduledIterator is returned from FilterImplementationActivationRescheduled and is used to iterate over the raw logs and unpacked data for ImplementationActivationRescheduled events raised by the DepositMainnet contract.
type DepositMainnetImplementationActivationRescheduledIterator struct {
	Event *DepositMainnetImplementationActivationRescheduled // Event containing the contract specifics and raw log

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
func (it *DepositMainnetImplementationActivationRescheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetImplementationActivationRescheduled)
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
		it.Event = new(DepositMainnetImplementationActivationRescheduled)
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
func (it *DepositMainnetImplementationActivationRescheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetImplementationActivationRescheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetImplementationActivationRescheduled represents a ImplementationActivationRescheduled event raised by the DepositMainnet contract.
type DepositMainnetImplementationActivationRescheduled struct {
	Implementation        common.Address
	UpdatedActivationTime *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterImplementationActivationRescheduled is a free log retrieval operation binding the contract event 0xdda7b7d1f8141bd98b4378ee60e6231f89598ca02949a9d0550904dc96efeeb7.
//
// Solidity: event ImplementationActivationRescheduled(address indexed implementation, uint256 updatedActivationTime)
func (_DepositMainnet *DepositMainnetFilterer) FilterImplementationActivationRescheduled(opts *bind.FilterOpts, implementation []common.Address) (*DepositMainnetImplementationActivationRescheduledIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "ImplementationActivationRescheduled", implementationRule)
	if err != nil {
		return nil, err
	}
	return &DepositMainnetImplementationActivationRescheduledIterator{contract: _DepositMainnet.contract, event: "ImplementationActivationRescheduled", logs: logs, sub: sub}, nil
}

// WatchImplementationActivationRescheduled is a free log subscription operation binding the contract event 0xdda7b7d1f8141bd98b4378ee60e6231f89598ca02949a9d0550904dc96efeeb7.
//
// Solidity: event ImplementationActivationRescheduled(address indexed implementation, uint256 updatedActivationTime)
func (_DepositMainnet *DepositMainnetFilterer) WatchImplementationActivationRescheduled(opts *bind.WatchOpts, sink chan<- *DepositMainnetImplementationActivationRescheduled, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "ImplementationActivationRescheduled", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetImplementationActivationRescheduled)
				if err := _DepositMainnet.contract.UnpackLog(event, "ImplementationActivationRescheduled", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseImplementationActivationRescheduled(log types.Log) (*DepositMainnetImplementationActivationRescheduled, error) {
	event := new(DepositMainnetImplementationActivationRescheduled)
	if err := _DepositMainnet.contract.UnpackLog(event, "ImplementationActivationRescheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogAssetWithdrawalAllowedIterator is returned from FilterLogAssetWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogAssetWithdrawalAllowed events raised by the DepositMainnet contract.
type DepositMainnetLogAssetWithdrawalAllowedIterator struct {
	Event *DepositMainnetLogAssetWithdrawalAllowed // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogAssetWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogAssetWithdrawalAllowed)
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
		it.Event = new(DepositMainnetLogAssetWithdrawalAllowed)
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
func (it *DepositMainnetLogAssetWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogAssetWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogAssetWithdrawalAllowed represents a LogAssetWithdrawalAllowed event raised by the DepositMainnet contract.
type DepositMainnetLogAssetWithdrawalAllowed struct {
	OwnerKey        *big.Int
	AssetId         *big.Int
	QuantizedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogAssetWithdrawalAllowed is a free log retrieval operation binding the contract event 0xd31f59c968eb53320f624721416cf88605da9aadbaa723405e52affe9de4b07f.
//
// Solidity: event LogAssetWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogAssetWithdrawalAllowed(opts *bind.FilterOpts) (*DepositMainnetLogAssetWithdrawalAllowedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogAssetWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogAssetWithdrawalAllowedIterator{contract: _DepositMainnet.contract, event: "LogAssetWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogAssetWithdrawalAllowed is a free log subscription operation binding the contract event 0xd31f59c968eb53320f624721416cf88605da9aadbaa723405e52affe9de4b07f.
//
// Solidity: event LogAssetWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogAssetWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogAssetWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogAssetWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogAssetWithdrawalAllowed)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogAssetWithdrawalAllowed", log); err != nil {
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

// ParseLogAssetWithdrawalAllowed is a log parse operation binding the contract event 0xd31f59c968eb53320f624721416cf88605da9aadbaa723405e52affe9de4b07f.
//
// Solidity: event LogAssetWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) ParseLogAssetWithdrawalAllowed(log types.Log) (*DepositMainnetLogAssetWithdrawalAllowed, error) {
	event := new(DepositMainnetLogAssetWithdrawalAllowed)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogAssetWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogDefaultVaultWithdrawalLockSetIterator is returned from FilterLogDefaultVaultWithdrawalLockSet and is used to iterate over the raw logs and unpacked data for LogDefaultVaultWithdrawalLockSet events raised by the DepositMainnet contract.
type DepositMainnetLogDefaultVaultWithdrawalLockSetIterator struct {
	Event *DepositMainnetLogDefaultVaultWithdrawalLockSet // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogDefaultVaultWithdrawalLockSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogDefaultVaultWithdrawalLockSet)
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
		it.Event = new(DepositMainnetLogDefaultVaultWithdrawalLockSet)
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
func (it *DepositMainnetLogDefaultVaultWithdrawalLockSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogDefaultVaultWithdrawalLockSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogDefaultVaultWithdrawalLockSet represents a LogDefaultVaultWithdrawalLockSet event raised by the DepositMainnet contract.
type DepositMainnetLogDefaultVaultWithdrawalLockSet struct {
	NewDefaultLockTime *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDefaultVaultWithdrawalLockSet is a free log retrieval operation binding the contract event 0x832169a4c3cea413f0041437fd118afbc4b33edbf6783da382128bca1e56b2d2.
//
// Solidity: event LogDefaultVaultWithdrawalLockSet(uint256 newDefaultLockTime)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogDefaultVaultWithdrawalLockSet(opts *bind.FilterOpts) (*DepositMainnetLogDefaultVaultWithdrawalLockSetIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogDefaultVaultWithdrawalLockSet")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogDefaultVaultWithdrawalLockSetIterator{contract: _DepositMainnet.contract, event: "LogDefaultVaultWithdrawalLockSet", logs: logs, sub: sub}, nil
}

// WatchLogDefaultVaultWithdrawalLockSet is a free log subscription operation binding the contract event 0x832169a4c3cea413f0041437fd118afbc4b33edbf6783da382128bca1e56b2d2.
//
// Solidity: event LogDefaultVaultWithdrawalLockSet(uint256 newDefaultLockTime)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogDefaultVaultWithdrawalLockSet(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogDefaultVaultWithdrawalLockSet) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogDefaultVaultWithdrawalLockSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogDefaultVaultWithdrawalLockSet)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogDefaultVaultWithdrawalLockSet", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogDefaultVaultWithdrawalLockSet(log types.Log) (*DepositMainnetLogDefaultVaultWithdrawalLockSet, error) {
	event := new(DepositMainnetLogDefaultVaultWithdrawalLockSet)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogDefaultVaultWithdrawalLockSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogDepositIterator is returned from FilterLogDeposit and is used to iterate over the raw logs and unpacked data for LogDeposit events raised by the DepositMainnet contract.
type DepositMainnetLogDepositIterator struct {
	Event *DepositMainnetLogDeposit // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogDeposit)
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
		it.Event = new(DepositMainnetLogDeposit)
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
func (it *DepositMainnetLogDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogDeposit represents a LogDeposit event raised by the DepositMainnet contract.
type DepositMainnetLogDeposit struct {
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
func (_DepositMainnet *DepositMainnetFilterer) FilterLogDeposit(opts *bind.FilterOpts) (*DepositMainnetLogDepositIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogDeposit")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogDepositIterator{contract: _DepositMainnet.contract, event: "LogDeposit", logs: logs, sub: sub}, nil
}

// WatchLogDeposit is a free log subscription operation binding the contract event 0x06724742ccc8c330a39a641ef02a0b419bd09248360680bb38159b0a8c2635d6.
//
// Solidity: event LogDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogDeposit(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogDeposit) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogDeposit)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogDeposit", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogDeposit(log types.Log) (*DepositMainnetLogDeposit, error) {
	event := new(DepositMainnetLogDeposit)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogDepositCancelIterator is returned from FilterLogDepositCancel and is used to iterate over the raw logs and unpacked data for LogDepositCancel events raised by the DepositMainnet contract.
type DepositMainnetLogDepositCancelIterator struct {
	Event *DepositMainnetLogDepositCancel // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogDepositCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogDepositCancel)
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
		it.Event = new(DepositMainnetLogDepositCancel)
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
func (it *DepositMainnetLogDepositCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogDepositCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogDepositCancel represents a LogDepositCancel event raised by the DepositMainnet contract.
type DepositMainnetLogDepositCancel struct {
	StarkKey *big.Int
	VaultId  *big.Int
	AssetId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogDepositCancel is a free log retrieval operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogDepositCancel(opts *bind.FilterOpts) (*DepositMainnetLogDepositCancelIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogDepositCancel")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogDepositCancelIterator{contract: _DepositMainnet.contract, event: "LogDepositCancel", logs: logs, sub: sub}, nil
}

// WatchLogDepositCancel is a free log subscription operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogDepositCancel(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogDepositCancel) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogDepositCancel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogDepositCancel)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogDepositCancel", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogDepositCancel(log types.Log) (*DepositMainnetLogDepositCancel, error) {
	event := new(DepositMainnetLogDepositCancel)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogDepositCancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogDepositCancelReclaimedIterator is returned from FilterLogDepositCancelReclaimed and is used to iterate over the raw logs and unpacked data for LogDepositCancelReclaimed events raised by the DepositMainnet contract.
type DepositMainnetLogDepositCancelReclaimedIterator struct {
	Event *DepositMainnetLogDepositCancelReclaimed // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogDepositCancelReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogDepositCancelReclaimed)
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
		it.Event = new(DepositMainnetLogDepositCancelReclaimed)
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
func (it *DepositMainnetLogDepositCancelReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogDepositCancelReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogDepositCancelReclaimed represents a LogDepositCancelReclaimed event raised by the DepositMainnet contract.
type DepositMainnetLogDepositCancelReclaimed struct {
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
func (_DepositMainnet *DepositMainnetFilterer) FilterLogDepositCancelReclaimed(opts *bind.FilterOpts) (*DepositMainnetLogDepositCancelReclaimedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogDepositCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogDepositCancelReclaimedIterator{contract: _DepositMainnet.contract, event: "LogDepositCancelReclaimed", logs: logs, sub: sub}, nil
}

// WatchLogDepositCancelReclaimed is a free log subscription operation binding the contract event 0xe3e46ecf1138180bf93cba62a0b7e661d976a8ab3d40243f7b082667d8f500af.
//
// Solidity: event LogDepositCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogDepositCancelReclaimed(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogDepositCancelReclaimed) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogDepositCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogDepositCancelReclaimed)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogDepositCancelReclaimed", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogDepositCancelReclaimed(log types.Log) (*DepositMainnetLogDepositCancelReclaimed, error) {
	event := new(DepositMainnetLogDepositCancelReclaimed)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogDepositCancelReclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogDepositNftCancelReclaimedIterator is returned from FilterLogDepositNftCancelReclaimed and is used to iterate over the raw logs and unpacked data for LogDepositNftCancelReclaimed events raised by the DepositMainnet contract.
type DepositMainnetLogDepositNftCancelReclaimedIterator struct {
	Event *DepositMainnetLogDepositNftCancelReclaimed // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogDepositNftCancelReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogDepositNftCancelReclaimed)
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
		it.Event = new(DepositMainnetLogDepositNftCancelReclaimed)
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
func (it *DepositMainnetLogDepositNftCancelReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogDepositNftCancelReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogDepositNftCancelReclaimed represents a LogDepositNftCancelReclaimed event raised by the DepositMainnet contract.
type DepositMainnetLogDepositNftCancelReclaimed struct {
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
func (_DepositMainnet *DepositMainnetFilterer) FilterLogDepositNftCancelReclaimed(opts *bind.FilterOpts) (*DepositMainnetLogDepositNftCancelReclaimedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogDepositNftCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogDepositNftCancelReclaimedIterator{contract: _DepositMainnet.contract, event: "LogDepositNftCancelReclaimed", logs: logs, sub: sub}, nil
}

// WatchLogDepositNftCancelReclaimed is a free log subscription operation binding the contract event 0xf00c0c1a754f6df7545d96a7e12aad552728b94ca6aa94f81e297bdbcf1dab9c.
//
// Solidity: event LogDepositNftCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogDepositNftCancelReclaimed(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogDepositNftCancelReclaimed) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogDepositNftCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogDepositNftCancelReclaimed)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogDepositNftCancelReclaimed", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogDepositNftCancelReclaimed(log types.Log) (*DepositMainnetLogDepositNftCancelReclaimed, error) {
	event := new(DepositMainnetLogDepositNftCancelReclaimed)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogDepositNftCancelReclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogDepositToVaultIterator is returned from FilterLogDepositToVault and is used to iterate over the raw logs and unpacked data for LogDepositToVault events raised by the DepositMainnet contract.
type DepositMainnetLogDepositToVaultIterator struct {
	Event *DepositMainnetLogDepositToVault // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogDepositToVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogDepositToVault)
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
		it.Event = new(DepositMainnetLogDepositToVault)
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
func (it *DepositMainnetLogDepositToVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogDepositToVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogDepositToVault represents a LogDepositToVault event raised by the DepositMainnet contract.
type DepositMainnetLogDepositToVault struct {
	EthKey             common.Address
	AssetType          *big.Int
	AssetId            *big.Int
	VaultId            *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDepositToVault is a free log retrieval operation binding the contract event 0xeab1149e7dda57c040af22580ce50b38b8b3af38433be30a24b3938166cd45e9.
//
// Solidity: event LogDepositToVault(address ethKey, uint256 assetType, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogDepositToVault(opts *bind.FilterOpts) (*DepositMainnetLogDepositToVaultIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogDepositToVault")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogDepositToVaultIterator{contract: _DepositMainnet.contract, event: "LogDepositToVault", logs: logs, sub: sub}, nil
}

// WatchLogDepositToVault is a free log subscription operation binding the contract event 0xeab1149e7dda57c040af22580ce50b38b8b3af38433be30a24b3938166cd45e9.
//
// Solidity: event LogDepositToVault(address ethKey, uint256 assetType, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogDepositToVault(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogDepositToVault) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogDepositToVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogDepositToVault)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogDepositToVault", log); err != nil {
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

// ParseLogDepositToVault is a log parse operation binding the contract event 0xeab1149e7dda57c040af22580ce50b38b8b3af38433be30a24b3938166cd45e9.
//
// Solidity: event LogDepositToVault(address ethKey, uint256 assetType, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) ParseLogDepositToVault(log types.Log) (*DepositMainnetLogDepositToVault, error) {
	event := new(DepositMainnetLogDepositToVault)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogDepositToVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogDepositWithTokenIdIterator is returned from FilterLogDepositWithTokenId and is used to iterate over the raw logs and unpacked data for LogDepositWithTokenId events raised by the DepositMainnet contract.
type DepositMainnetLogDepositWithTokenIdIterator struct {
	Event *DepositMainnetLogDepositWithTokenId // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogDepositWithTokenIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogDepositWithTokenId)
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
		it.Event = new(DepositMainnetLogDepositWithTokenId)
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
func (it *DepositMainnetLogDepositWithTokenIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogDepositWithTokenIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogDepositWithTokenId represents a LogDepositWithTokenId event raised by the DepositMainnet contract.
type DepositMainnetLogDepositWithTokenId struct {
	DepositorEthKey    common.Address
	StarkKey           *big.Int
	VaultId            *big.Int
	AssetType          *big.Int
	TokenId            *big.Int
	AssetId            *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDepositWithTokenId is a free log retrieval operation binding the contract event 0xed94dc026fa9364c53bc0af51cde7f54f3109b3f31fceb26d01396d80e20453b.
//
// Solidity: event LogDepositWithTokenId(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogDepositWithTokenId(opts *bind.FilterOpts) (*DepositMainnetLogDepositWithTokenIdIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogDepositWithTokenId")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogDepositWithTokenIdIterator{contract: _DepositMainnet.contract, event: "LogDepositWithTokenId", logs: logs, sub: sub}, nil
}

// WatchLogDepositWithTokenId is a free log subscription operation binding the contract event 0xed94dc026fa9364c53bc0af51cde7f54f3109b3f31fceb26d01396d80e20453b.
//
// Solidity: event LogDepositWithTokenId(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogDepositWithTokenId(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogDepositWithTokenId) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogDepositWithTokenId")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogDepositWithTokenId)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogDepositWithTokenId", log); err != nil {
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

// ParseLogDepositWithTokenId is a log parse operation binding the contract event 0xed94dc026fa9364c53bc0af51cde7f54f3109b3f31fceb26d01396d80e20453b.
//
// Solidity: event LogDepositWithTokenId(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) ParseLogDepositWithTokenId(log types.Log) (*DepositMainnetLogDepositWithTokenId, error) {
	event := new(DepositMainnetLogDepositWithTokenId)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogDepositWithTokenId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogDepositWithTokenIdCancelReclaimedIterator is returned from FilterLogDepositWithTokenIdCancelReclaimed and is used to iterate over the raw logs and unpacked data for LogDepositWithTokenIdCancelReclaimed events raised by the DepositMainnet contract.
type DepositMainnetLogDepositWithTokenIdCancelReclaimedIterator struct {
	Event *DepositMainnetLogDepositWithTokenIdCancelReclaimed // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogDepositWithTokenIdCancelReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogDepositWithTokenIdCancelReclaimed)
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
		it.Event = new(DepositMainnetLogDepositWithTokenIdCancelReclaimed)
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
func (it *DepositMainnetLogDepositWithTokenIdCancelReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogDepositWithTokenIdCancelReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogDepositWithTokenIdCancelReclaimed represents a LogDepositWithTokenIdCancelReclaimed event raised by the DepositMainnet contract.
type DepositMainnetLogDepositWithTokenIdCancelReclaimed struct {
	StarkKey           *big.Int
	VaultId            *big.Int
	AssetType          *big.Int
	TokenId            *big.Int
	AssetId            *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDepositWithTokenIdCancelReclaimed is a free log retrieval operation binding the contract event 0xcc00f2179d127845242252f3c3b6b238c5ed33c2e933179f09653cfb1cdee7ca.
//
// Solidity: event LogDepositWithTokenIdCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogDepositWithTokenIdCancelReclaimed(opts *bind.FilterOpts) (*DepositMainnetLogDepositWithTokenIdCancelReclaimedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogDepositWithTokenIdCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogDepositWithTokenIdCancelReclaimedIterator{contract: _DepositMainnet.contract, event: "LogDepositWithTokenIdCancelReclaimed", logs: logs, sub: sub}, nil
}

// WatchLogDepositWithTokenIdCancelReclaimed is a free log subscription operation binding the contract event 0xcc00f2179d127845242252f3c3b6b238c5ed33c2e933179f09653cfb1cdee7ca.
//
// Solidity: event LogDepositWithTokenIdCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogDepositWithTokenIdCancelReclaimed(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogDepositWithTokenIdCancelReclaimed) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogDepositWithTokenIdCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogDepositWithTokenIdCancelReclaimed)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogDepositWithTokenIdCancelReclaimed", log); err != nil {
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

// ParseLogDepositWithTokenIdCancelReclaimed is a log parse operation binding the contract event 0xcc00f2179d127845242252f3c3b6b238c5ed33c2e933179f09653cfb1cdee7ca.
//
// Solidity: event LogDepositWithTokenIdCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) ParseLogDepositWithTokenIdCancelReclaimed(log types.Log) (*DepositMainnetLogDepositWithTokenIdCancelReclaimed, error) {
	event := new(DepositMainnetLogDepositWithTokenIdCancelReclaimed)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogDepositWithTokenIdCancelReclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogFrozenIterator is returned from FilterLogFrozen and is used to iterate over the raw logs and unpacked data for LogFrozen events raised by the DepositMainnet contract.
type DepositMainnetLogFrozenIterator struct {
	Event *DepositMainnetLogFrozen // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogFrozen)
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
		it.Event = new(DepositMainnetLogFrozen)
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
func (it *DepositMainnetLogFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogFrozen represents a LogFrozen event raised by the DepositMainnet contract.
type DepositMainnetLogFrozen struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogFrozen is a free log retrieval operation binding the contract event 0xf5b8e6419478ab140eb98026ab5bd607038cb0ac4d4dad5b1fc0848dfd203d1f.
//
// Solidity: event LogFrozen()
func (_DepositMainnet *DepositMainnetFilterer) FilterLogFrozen(opts *bind.FilterOpts) (*DepositMainnetLogFrozenIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogFrozen")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogFrozenIterator{contract: _DepositMainnet.contract, event: "LogFrozen", logs: logs, sub: sub}, nil
}

// WatchLogFrozen is a free log subscription operation binding the contract event 0xf5b8e6419478ab140eb98026ab5bd607038cb0ac4d4dad5b1fc0848dfd203d1f.
//
// Solidity: event LogFrozen()
func (_DepositMainnet *DepositMainnetFilterer) WatchLogFrozen(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogFrozen) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogFrozen)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogFrozen", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogFrozen(log types.Log) (*DepositMainnetLogFrozen, error) {
	event := new(DepositMainnetLogFrozen)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogFullWithdrawalRequestIterator is returned from FilterLogFullWithdrawalRequest and is used to iterate over the raw logs and unpacked data for LogFullWithdrawalRequest events raised by the DepositMainnet contract.
type DepositMainnetLogFullWithdrawalRequestIterator struct {
	Event *DepositMainnetLogFullWithdrawalRequest // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogFullWithdrawalRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogFullWithdrawalRequest)
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
		it.Event = new(DepositMainnetLogFullWithdrawalRequest)
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
func (it *DepositMainnetLogFullWithdrawalRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogFullWithdrawalRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogFullWithdrawalRequest represents a LogFullWithdrawalRequest event raised by the DepositMainnet contract.
type DepositMainnetLogFullWithdrawalRequest struct {
	OwnerKey *big.Int
	VaultId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogFullWithdrawalRequest is a free log retrieval operation binding the contract event 0x08eb46dbb87dcfe92d4846e5766802051525fba08a9b48318f5e0fe41186d298.
//
// Solidity: event LogFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogFullWithdrawalRequest(opts *bind.FilterOpts) (*DepositMainnetLogFullWithdrawalRequestIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogFullWithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogFullWithdrawalRequestIterator{contract: _DepositMainnet.contract, event: "LogFullWithdrawalRequest", logs: logs, sub: sub}, nil
}

// WatchLogFullWithdrawalRequest is a free log subscription operation binding the contract event 0x08eb46dbb87dcfe92d4846e5766802051525fba08a9b48318f5e0fe41186d298.
//
// Solidity: event LogFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogFullWithdrawalRequest(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogFullWithdrawalRequest) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogFullWithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogFullWithdrawalRequest)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogFullWithdrawalRequest", log); err != nil {
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
// Solidity: event LogFullWithdrawalRequest(uint256 ownerKey, uint256 vaultId)
func (_DepositMainnet *DepositMainnetFilterer) ParseLogFullWithdrawalRequest(log types.Log) (*DepositMainnetLogFullWithdrawalRequest, error) {
	event := new(DepositMainnetLogFullWithdrawalRequest)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogFullWithdrawalRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogMintWithdrawalPerformedIterator is returned from FilterLogMintWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogMintWithdrawalPerformed events raised by the DepositMainnet contract.
type DepositMainnetLogMintWithdrawalPerformedIterator struct {
	Event *DepositMainnetLogMintWithdrawalPerformed // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogMintWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogMintWithdrawalPerformed)
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
		it.Event = new(DepositMainnetLogMintWithdrawalPerformed)
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
func (it *DepositMainnetLogMintWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogMintWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogMintWithdrawalPerformed represents a LogMintWithdrawalPerformed event raised by the DepositMainnet contract.
type DepositMainnetLogMintWithdrawalPerformed struct {
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
func (_DepositMainnet *DepositMainnetFilterer) FilterLogMintWithdrawalPerformed(opts *bind.FilterOpts) (*DepositMainnetLogMintWithdrawalPerformedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogMintWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogMintWithdrawalPerformedIterator{contract: _DepositMainnet.contract, event: "LogMintWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogMintWithdrawalPerformed is a free log subscription operation binding the contract event 0x7e6e15df814c1a309a57686de672b2bedd128eacde35c5370c36d6840d4e9a92.
//
// Solidity: event LogMintWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, uint256 assetId)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogMintWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogMintWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogMintWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogMintWithdrawalPerformed)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogMintWithdrawalPerformed", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogMintWithdrawalPerformed(log types.Log) (*DepositMainnetLogMintWithdrawalPerformed, error) {
	event := new(DepositMainnetLogMintWithdrawalPerformed)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogMintWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogMintableWithdrawalAllowedIterator is returned from FilterLogMintableWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogMintableWithdrawalAllowed events raised by the DepositMainnet contract.
type DepositMainnetLogMintableWithdrawalAllowedIterator struct {
	Event *DepositMainnetLogMintableWithdrawalAllowed // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogMintableWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogMintableWithdrawalAllowed)
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
		it.Event = new(DepositMainnetLogMintableWithdrawalAllowed)
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
func (it *DepositMainnetLogMintableWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogMintableWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogMintableWithdrawalAllowed represents a LogMintableWithdrawalAllowed event raised by the DepositMainnet contract.
type DepositMainnetLogMintableWithdrawalAllowed struct {
	OwnerKey        *big.Int
	AssetId         *big.Int
	QuantizedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogMintableWithdrawalAllowed is a free log retrieval operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogMintableWithdrawalAllowed(opts *bind.FilterOpts) (*DepositMainnetLogMintableWithdrawalAllowedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogMintableWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogMintableWithdrawalAllowedIterator{contract: _DepositMainnet.contract, event: "LogMintableWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogMintableWithdrawalAllowed is a free log subscription operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 ownerKey, uint256 assetId, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogMintableWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogMintableWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogMintableWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogMintableWithdrawalAllowed)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogMintableWithdrawalAllowed", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogMintableWithdrawalAllowed(log types.Log) (*DepositMainnetLogMintableWithdrawalAllowed, error) {
	event := new(DepositMainnetLogMintableWithdrawalAllowed)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogMintableWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogNewGovernorAcceptedIterator is returned from FilterLogNewGovernorAccepted and is used to iterate over the raw logs and unpacked data for LogNewGovernorAccepted events raised by the DepositMainnet contract.
type DepositMainnetLogNewGovernorAcceptedIterator struct {
	Event *DepositMainnetLogNewGovernorAccepted // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogNewGovernorAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogNewGovernorAccepted)
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
		it.Event = new(DepositMainnetLogNewGovernorAccepted)
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
func (it *DepositMainnetLogNewGovernorAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogNewGovernorAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogNewGovernorAccepted represents a LogNewGovernorAccepted event raised by the DepositMainnet contract.
type DepositMainnetLogNewGovernorAccepted struct {
	AcceptedGovernor common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogNewGovernorAccepted is a free log retrieval operation binding the contract event 0xcfb473e6c03f9a29ddaf990e736fa3de5188a0bd85d684f5b6e164ebfbfff5d2.
//
// Solidity: event LogNewGovernorAccepted(address acceptedGovernor)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogNewGovernorAccepted(opts *bind.FilterOpts) (*DepositMainnetLogNewGovernorAcceptedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogNewGovernorAccepted")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogNewGovernorAcceptedIterator{contract: _DepositMainnet.contract, event: "LogNewGovernorAccepted", logs: logs, sub: sub}, nil
}

// WatchLogNewGovernorAccepted is a free log subscription operation binding the contract event 0xcfb473e6c03f9a29ddaf990e736fa3de5188a0bd85d684f5b6e164ebfbfff5d2.
//
// Solidity: event LogNewGovernorAccepted(address acceptedGovernor)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogNewGovernorAccepted(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogNewGovernorAccepted) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogNewGovernorAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogNewGovernorAccepted)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogNewGovernorAccepted", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogNewGovernorAccepted(log types.Log) (*DepositMainnetLogNewGovernorAccepted, error) {
	event := new(DepositMainnetLogNewGovernorAccepted)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogNewGovernorAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogNftDepositIterator is returned from FilterLogNftDeposit and is used to iterate over the raw logs and unpacked data for LogNftDeposit events raised by the DepositMainnet contract.
type DepositMainnetLogNftDepositIterator struct {
	Event *DepositMainnetLogNftDeposit // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogNftDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogNftDeposit)
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
		it.Event = new(DepositMainnetLogNftDeposit)
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
func (it *DepositMainnetLogNftDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogNftDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogNftDeposit represents a LogNftDeposit event raised by the DepositMainnet contract.
type DepositMainnetLogNftDeposit struct {
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
func (_DepositMainnet *DepositMainnetFilterer) FilterLogNftDeposit(opts *bind.FilterOpts) (*DepositMainnetLogNftDepositIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogNftDeposit")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogNftDepositIterator{contract: _DepositMainnet.contract, event: "LogNftDeposit", logs: logs, sub: sub}, nil
}

// WatchLogNftDeposit is a free log subscription operation binding the contract event 0x0fcf2162832b2d6033d4d34d2f45a28d9cfee523f1899945bbdd32529cfda67b.
//
// Solidity: event LogNftDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogNftDeposit(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogNftDeposit) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogNftDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogNftDeposit)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogNftDeposit", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogNftDeposit(log types.Log) (*DepositMainnetLogNftDeposit, error) {
	event := new(DepositMainnetLogNftDeposit)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogNftDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogNftWithdrawalAllowedIterator is returned from FilterLogNftWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogNftWithdrawalAllowed events raised by the DepositMainnet contract.
type DepositMainnetLogNftWithdrawalAllowedIterator struct {
	Event *DepositMainnetLogNftWithdrawalAllowed // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogNftWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogNftWithdrawalAllowed)
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
		it.Event = new(DepositMainnetLogNftWithdrawalAllowed)
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
func (it *DepositMainnetLogNftWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogNftWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogNftWithdrawalAllowed represents a LogNftWithdrawalAllowed event raised by the DepositMainnet contract.
type DepositMainnetLogNftWithdrawalAllowed struct {
	OwnerKey *big.Int
	AssetId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNftWithdrawalAllowed is a free log retrieval operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 ownerKey, uint256 assetId)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogNftWithdrawalAllowed(opts *bind.FilterOpts) (*DepositMainnetLogNftWithdrawalAllowedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogNftWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogNftWithdrawalAllowedIterator{contract: _DepositMainnet.contract, event: "LogNftWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogNftWithdrawalAllowed is a free log subscription operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 ownerKey, uint256 assetId)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogNftWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogNftWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogNftWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogNftWithdrawalAllowed)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogNftWithdrawalAllowed", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogNftWithdrawalAllowed(log types.Log) (*DepositMainnetLogNftWithdrawalAllowed, error) {
	event := new(DepositMainnetLogNftWithdrawalAllowed)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogNftWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogNftWithdrawalPerformedIterator is returned from FilterLogNftWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogNftWithdrawalPerformed events raised by the DepositMainnet contract.
type DepositMainnetLogNftWithdrawalPerformedIterator struct {
	Event *DepositMainnetLogNftWithdrawalPerformed // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogNftWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogNftWithdrawalPerformed)
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
		it.Event = new(DepositMainnetLogNftWithdrawalPerformed)
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
func (it *DepositMainnetLogNftWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogNftWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogNftWithdrawalPerformed represents a LogNftWithdrawalPerformed event raised by the DepositMainnet contract.
type DepositMainnetLogNftWithdrawalPerformed struct {
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
func (_DepositMainnet *DepositMainnetFilterer) FilterLogNftWithdrawalPerformed(opts *bind.FilterOpts) (*DepositMainnetLogNftWithdrawalPerformedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogNftWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogNftWithdrawalPerformedIterator{contract: _DepositMainnet.contract, event: "LogNftWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogNftWithdrawalPerformed is a free log subscription operation binding the contract event 0xa5cfa8e2199ec5b8ca319288bcab72734207d30569756ee594a74b4df7abbf41.
//
// Solidity: event LogNftWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 tokenId, uint256 assetId, address recipient)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogNftWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogNftWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogNftWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogNftWithdrawalPerformed)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogNftWithdrawalPerformed", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogNftWithdrawalPerformed(log types.Log) (*DepositMainnetLogNftWithdrawalPerformed, error) {
	event := new(DepositMainnetLogNftWithdrawalPerformed)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogNftWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogNominatedGovernorIterator is returned from FilterLogNominatedGovernor and is used to iterate over the raw logs and unpacked data for LogNominatedGovernor events raised by the DepositMainnet contract.
type DepositMainnetLogNominatedGovernorIterator struct {
	Event *DepositMainnetLogNominatedGovernor // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogNominatedGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogNominatedGovernor)
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
		it.Event = new(DepositMainnetLogNominatedGovernor)
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
func (it *DepositMainnetLogNominatedGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogNominatedGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogNominatedGovernor represents a LogNominatedGovernor event raised by the DepositMainnet contract.
type DepositMainnetLogNominatedGovernor struct {
	NominatedGovernor common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogNominatedGovernor is a free log retrieval operation binding the contract event 0x6166272c8d3f5f579082f2827532732f97195007983bb5b83ac12c56700b01a6.
//
// Solidity: event LogNominatedGovernor(address nominatedGovernor)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogNominatedGovernor(opts *bind.FilterOpts) (*DepositMainnetLogNominatedGovernorIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogNominatedGovernor")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogNominatedGovernorIterator{contract: _DepositMainnet.contract, event: "LogNominatedGovernor", logs: logs, sub: sub}, nil
}

// WatchLogNominatedGovernor is a free log subscription operation binding the contract event 0x6166272c8d3f5f579082f2827532732f97195007983bb5b83ac12c56700b01a6.
//
// Solidity: event LogNominatedGovernor(address nominatedGovernor)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogNominatedGovernor(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogNominatedGovernor) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogNominatedGovernor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogNominatedGovernor)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogNominatedGovernor", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogNominatedGovernor(log types.Log) (*DepositMainnetLogNominatedGovernor, error) {
	event := new(DepositMainnetLogNominatedGovernor)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogNominatedGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogNominationCancelledIterator is returned from FilterLogNominationCancelled and is used to iterate over the raw logs and unpacked data for LogNominationCancelled events raised by the DepositMainnet contract.
type DepositMainnetLogNominationCancelledIterator struct {
	Event *DepositMainnetLogNominationCancelled // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogNominationCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogNominationCancelled)
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
		it.Event = new(DepositMainnetLogNominationCancelled)
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
func (it *DepositMainnetLogNominationCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogNominationCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogNominationCancelled represents a LogNominationCancelled event raised by the DepositMainnet contract.
type DepositMainnetLogNominationCancelled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNominationCancelled is a free log retrieval operation binding the contract event 0x7a8dc7dd7fffb43c4807438fa62729225156941e641fd877938f4edade3429f5.
//
// Solidity: event LogNominationCancelled()
func (_DepositMainnet *DepositMainnetFilterer) FilterLogNominationCancelled(opts *bind.FilterOpts) (*DepositMainnetLogNominationCancelledIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogNominationCancelled")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogNominationCancelledIterator{contract: _DepositMainnet.contract, event: "LogNominationCancelled", logs: logs, sub: sub}, nil
}

// WatchLogNominationCancelled is a free log subscription operation binding the contract event 0x7a8dc7dd7fffb43c4807438fa62729225156941e641fd877938f4edade3429f5.
//
// Solidity: event LogNominationCancelled()
func (_DepositMainnet *DepositMainnetFilterer) WatchLogNominationCancelled(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogNominationCancelled) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogNominationCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogNominationCancelled)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogNominationCancelled", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogNominationCancelled(log types.Log) (*DepositMainnetLogNominationCancelled, error) {
	event := new(DepositMainnetLogNominationCancelled)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogNominationCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogOperatorAddedIterator is returned from FilterLogOperatorAdded and is used to iterate over the raw logs and unpacked data for LogOperatorAdded events raised by the DepositMainnet contract.
type DepositMainnetLogOperatorAddedIterator struct {
	Event *DepositMainnetLogOperatorAdded // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogOperatorAdded)
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
		it.Event = new(DepositMainnetLogOperatorAdded)
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
func (it *DepositMainnetLogOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogOperatorAdded represents a LogOperatorAdded event raised by the DepositMainnet contract.
type DepositMainnetLogOperatorAdded struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogOperatorAdded is a free log retrieval operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogOperatorAdded(opts *bind.FilterOpts) (*DepositMainnetLogOperatorAddedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogOperatorAdded")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogOperatorAddedIterator{contract: _DepositMainnet.contract, event: "LogOperatorAdded", logs: logs, sub: sub}, nil
}

// WatchLogOperatorAdded is a free log subscription operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogOperatorAdded(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogOperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogOperatorAdded)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogOperatorAdded", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogOperatorAdded(log types.Log) (*DepositMainnetLogOperatorAdded, error) {
	event := new(DepositMainnetLogOperatorAdded)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogOperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogOperatorRemovedIterator is returned from FilterLogOperatorRemoved and is used to iterate over the raw logs and unpacked data for LogOperatorRemoved events raised by the DepositMainnet contract.
type DepositMainnetLogOperatorRemovedIterator struct {
	Event *DepositMainnetLogOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogOperatorRemoved)
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
		it.Event = new(DepositMainnetLogOperatorRemoved)
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
func (it *DepositMainnetLogOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogOperatorRemoved represents a LogOperatorRemoved event raised by the DepositMainnet contract.
type DepositMainnetLogOperatorRemoved struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogOperatorRemoved is a free log retrieval operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogOperatorRemoved(opts *bind.FilterOpts) (*DepositMainnetLogOperatorRemovedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogOperatorRemoved")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogOperatorRemovedIterator{contract: _DepositMainnet.contract, event: "LogOperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchLogOperatorRemoved is a free log subscription operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogOperatorRemoved(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogOperatorRemoved) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogOperatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogOperatorRemoved)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogOperatorRemoved", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogOperatorRemoved(log types.Log) (*DepositMainnetLogOperatorRemoved, error) {
	event := new(DepositMainnetLogOperatorRemoved)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogOperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogRegisteredIterator is returned from FilterLogRegistered and is used to iterate over the raw logs and unpacked data for LogRegistered events raised by the DepositMainnet contract.
type DepositMainnetLogRegisteredIterator struct {
	Event *DepositMainnetLogRegistered // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogRegistered)
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
		it.Event = new(DepositMainnetLogRegistered)
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
func (it *DepositMainnetLogRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogRegistered represents a LogRegistered event raised by the DepositMainnet contract.
type DepositMainnetLogRegistered struct {
	Entry   common.Address
	EntryId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogRegistered is a free log retrieval operation binding the contract event 0xaa7f29c97c6763919ef56006f07fbf5c1ac734b0414665df43cecdbce9010c9b.
//
// Solidity: event LogRegistered(address entry, string entryId)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogRegistered(opts *bind.FilterOpts) (*DepositMainnetLogRegisteredIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogRegistered")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogRegisteredIterator{contract: _DepositMainnet.contract, event: "LogRegistered", logs: logs, sub: sub}, nil
}

// WatchLogRegistered is a free log subscription operation binding the contract event 0xaa7f29c97c6763919ef56006f07fbf5c1ac734b0414665df43cecdbce9010c9b.
//
// Solidity: event LogRegistered(address entry, string entryId)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogRegistered(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogRegistered) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogRegistered)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogRegistered", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogRegistered(log types.Log) (*DepositMainnetLogRegistered, error) {
	event := new(DepositMainnetLogRegistered)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogRemovalIntentIterator is returned from FilterLogRemovalIntent and is used to iterate over the raw logs and unpacked data for LogRemovalIntent events raised by the DepositMainnet contract.
type DepositMainnetLogRemovalIntentIterator struct {
	Event *DepositMainnetLogRemovalIntent // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogRemovalIntentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogRemovalIntent)
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
		it.Event = new(DepositMainnetLogRemovalIntent)
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
func (it *DepositMainnetLogRemovalIntentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogRemovalIntentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogRemovalIntent represents a LogRemovalIntent event raised by the DepositMainnet contract.
type DepositMainnetLogRemovalIntent struct {
	Entry   common.Address
	EntryId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogRemovalIntent is a free log retrieval operation binding the contract event 0xa98ac1f696177f16ca482653307aa4a02b68cf03b14409a546de5adf946484fc.
//
// Solidity: event LogRemovalIntent(address entry, string entryId)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogRemovalIntent(opts *bind.FilterOpts) (*DepositMainnetLogRemovalIntentIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogRemovalIntent")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogRemovalIntentIterator{contract: _DepositMainnet.contract, event: "LogRemovalIntent", logs: logs, sub: sub}, nil
}

// WatchLogRemovalIntent is a free log subscription operation binding the contract event 0xa98ac1f696177f16ca482653307aa4a02b68cf03b14409a546de5adf946484fc.
//
// Solidity: event LogRemovalIntent(address entry, string entryId)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogRemovalIntent(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogRemovalIntent) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogRemovalIntent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogRemovalIntent)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogRemovalIntent", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogRemovalIntent(log types.Log) (*DepositMainnetLogRemovalIntent, error) {
	event := new(DepositMainnetLogRemovalIntent)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogRemovalIntent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogRemovedIterator is returned from FilterLogRemoved and is used to iterate over the raw logs and unpacked data for LogRemoved events raised by the DepositMainnet contract.
type DepositMainnetLogRemovedIterator struct {
	Event *DepositMainnetLogRemoved // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogRemoved)
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
		it.Event = new(DepositMainnetLogRemoved)
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
func (it *DepositMainnetLogRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogRemoved represents a LogRemoved event raised by the DepositMainnet contract.
type DepositMainnetLogRemoved struct {
	Entry   common.Address
	EntryId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogRemoved is a free log retrieval operation binding the contract event 0x35b176cf4f09df896aa55e10eec90bb4c4689ea1d076135a8de3138d829d0142.
//
// Solidity: event LogRemoved(address entry, string entryId)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogRemoved(opts *bind.FilterOpts) (*DepositMainnetLogRemovedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogRemoved")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogRemovedIterator{contract: _DepositMainnet.contract, event: "LogRemoved", logs: logs, sub: sub}, nil
}

// WatchLogRemoved is a free log subscription operation binding the contract event 0x35b176cf4f09df896aa55e10eec90bb4c4689ea1d076135a8de3138d829d0142.
//
// Solidity: event LogRemoved(address entry, string entryId)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogRemoved(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogRemoved) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogRemoved)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogRemoved", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogRemoved(log types.Log) (*DepositMainnetLogRemoved, error) {
	event := new(DepositMainnetLogRemoved)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogRemovedGovernorIterator is returned from FilterLogRemovedGovernor and is used to iterate over the raw logs and unpacked data for LogRemovedGovernor events raised by the DepositMainnet contract.
type DepositMainnetLogRemovedGovernorIterator struct {
	Event *DepositMainnetLogRemovedGovernor // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogRemovedGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogRemovedGovernor)
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
		it.Event = new(DepositMainnetLogRemovedGovernor)
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
func (it *DepositMainnetLogRemovedGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogRemovedGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogRemovedGovernor represents a LogRemovedGovernor event raised by the DepositMainnet contract.
type DepositMainnetLogRemovedGovernor struct {
	RemovedGovernor common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogRemovedGovernor is a free log retrieval operation binding the contract event 0xd75f94825e770b8b512be8e74759e252ad00e102e38f50cce2f7c6f868a29599.
//
// Solidity: event LogRemovedGovernor(address removedGovernor)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogRemovedGovernor(opts *bind.FilterOpts) (*DepositMainnetLogRemovedGovernorIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogRemovedGovernor")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogRemovedGovernorIterator{contract: _DepositMainnet.contract, event: "LogRemovedGovernor", logs: logs, sub: sub}, nil
}

// WatchLogRemovedGovernor is a free log subscription operation binding the contract event 0xd75f94825e770b8b512be8e74759e252ad00e102e38f50cce2f7c6f868a29599.
//
// Solidity: event LogRemovedGovernor(address removedGovernor)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogRemovedGovernor(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogRemovedGovernor) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogRemovedGovernor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogRemovedGovernor)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogRemovedGovernor", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogRemovedGovernor(log types.Log) (*DepositMainnetLogRemovedGovernor, error) {
	event := new(DepositMainnetLogRemovedGovernor)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogRemovedGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogRootUpdateIterator is returned from FilterLogRootUpdate and is used to iterate over the raw logs and unpacked data for LogRootUpdate events raised by the DepositMainnet contract.
type DepositMainnetLogRootUpdateIterator struct {
	Event *DepositMainnetLogRootUpdate // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogRootUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogRootUpdate)
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
		it.Event = new(DepositMainnetLogRootUpdate)
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
func (it *DepositMainnetLogRootUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogRootUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogRootUpdate represents a LogRootUpdate event raised by the DepositMainnet contract.
type DepositMainnetLogRootUpdate struct {
	SequenceNumber    *big.Int
	BatchId           *big.Int
	ValidiumVaultRoot *big.Int
	RollupVaultRoot   *big.Int
	OrderRoot         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogRootUpdate is a free log retrieval operation binding the contract event 0x54fe7a67f8957a96919a0d1b81eeb25ea8c06f96ad528671da17a4a840040664.
//
// Solidity: event LogRootUpdate(uint256 sequenceNumber, uint256 batchId, uint256 validiumVaultRoot, uint256 rollupVaultRoot, uint256 orderRoot)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogRootUpdate(opts *bind.FilterOpts) (*DepositMainnetLogRootUpdateIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogRootUpdate")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogRootUpdateIterator{contract: _DepositMainnet.contract, event: "LogRootUpdate", logs: logs, sub: sub}, nil
}

// WatchLogRootUpdate is a free log subscription operation binding the contract event 0x54fe7a67f8957a96919a0d1b81eeb25ea8c06f96ad528671da17a4a840040664.
//
// Solidity: event LogRootUpdate(uint256 sequenceNumber, uint256 batchId, uint256 validiumVaultRoot, uint256 rollupVaultRoot, uint256 orderRoot)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogRootUpdate(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogRootUpdate) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogRootUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogRootUpdate)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogRootUpdate", log); err != nil {
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

// ParseLogRootUpdate is a log parse operation binding the contract event 0x54fe7a67f8957a96919a0d1b81eeb25ea8c06f96ad528671da17a4a840040664.
//
// Solidity: event LogRootUpdate(uint256 sequenceNumber, uint256 batchId, uint256 validiumVaultRoot, uint256 rollupVaultRoot, uint256 orderRoot)
func (_DepositMainnet *DepositMainnetFilterer) ParseLogRootUpdate(log types.Log) (*DepositMainnetLogRootUpdate, error) {
	event := new(DepositMainnetLogRootUpdate)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogRootUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogStateTransitionFactIterator is returned from FilterLogStateTransitionFact and is used to iterate over the raw logs and unpacked data for LogStateTransitionFact events raised by the DepositMainnet contract.
type DepositMainnetLogStateTransitionFactIterator struct {
	Event *DepositMainnetLogStateTransitionFact // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogStateTransitionFactIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogStateTransitionFact)
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
		it.Event = new(DepositMainnetLogStateTransitionFact)
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
func (it *DepositMainnetLogStateTransitionFactIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogStateTransitionFactIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogStateTransitionFact represents a LogStateTransitionFact event raised by the DepositMainnet contract.
type DepositMainnetLogStateTransitionFact struct {
	StateTransitionFact [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLogStateTransitionFact is a free log retrieval operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogStateTransitionFact(opts *bind.FilterOpts) (*DepositMainnetLogStateTransitionFactIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogStateTransitionFact")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogStateTransitionFactIterator{contract: _DepositMainnet.contract, event: "LogStateTransitionFact", logs: logs, sub: sub}, nil
}

// WatchLogStateTransitionFact is a free log subscription operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogStateTransitionFact(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogStateTransitionFact) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogStateTransitionFact")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogStateTransitionFact)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogStateTransitionFact", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogStateTransitionFact(log types.Log) (*DepositMainnetLogStateTransitionFact, error) {
	event := new(DepositMainnetLogStateTransitionFact)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogStateTransitionFact", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogTokenAdminAddedIterator is returned from FilterLogTokenAdminAdded and is used to iterate over the raw logs and unpacked data for LogTokenAdminAdded events raised by the DepositMainnet contract.
type DepositMainnetLogTokenAdminAddedIterator struct {
	Event *DepositMainnetLogTokenAdminAdded // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogTokenAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogTokenAdminAdded)
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
		it.Event = new(DepositMainnetLogTokenAdminAdded)
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
func (it *DepositMainnetLogTokenAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogTokenAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogTokenAdminAdded represents a LogTokenAdminAdded event raised by the DepositMainnet contract.
type DepositMainnetLogTokenAdminAdded struct {
	TokenAdmin common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogTokenAdminAdded is a free log retrieval operation binding the contract event 0x9085a9044aeb6daeeb5b4bf84af42b1a1613d4056f503c4e992b6396c16bd52f.
//
// Solidity: event LogTokenAdminAdded(address tokenAdmin)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogTokenAdminAdded(opts *bind.FilterOpts) (*DepositMainnetLogTokenAdminAddedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogTokenAdminAdded")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogTokenAdminAddedIterator{contract: _DepositMainnet.contract, event: "LogTokenAdminAdded", logs: logs, sub: sub}, nil
}

// WatchLogTokenAdminAdded is a free log subscription operation binding the contract event 0x9085a9044aeb6daeeb5b4bf84af42b1a1613d4056f503c4e992b6396c16bd52f.
//
// Solidity: event LogTokenAdminAdded(address tokenAdmin)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogTokenAdminAdded(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogTokenAdminAdded) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogTokenAdminAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogTokenAdminAdded)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogTokenAdminAdded", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogTokenAdminAdded(log types.Log) (*DepositMainnetLogTokenAdminAdded, error) {
	event := new(DepositMainnetLogTokenAdminAdded)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogTokenAdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogTokenAdminRemovedIterator is returned from FilterLogTokenAdminRemoved and is used to iterate over the raw logs and unpacked data for LogTokenAdminRemoved events raised by the DepositMainnet contract.
type DepositMainnetLogTokenAdminRemovedIterator struct {
	Event *DepositMainnetLogTokenAdminRemoved // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogTokenAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogTokenAdminRemoved)
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
		it.Event = new(DepositMainnetLogTokenAdminRemoved)
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
func (it *DepositMainnetLogTokenAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogTokenAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogTokenAdminRemoved represents a LogTokenAdminRemoved event raised by the DepositMainnet contract.
type DepositMainnetLogTokenAdminRemoved struct {
	TokenAdmin common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogTokenAdminRemoved is a free log retrieval operation binding the contract event 0xfa49aecb996ea8d99950bb051552dfcc0b5460a0bb209867a1ed8067c32c2177.
//
// Solidity: event LogTokenAdminRemoved(address tokenAdmin)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogTokenAdminRemoved(opts *bind.FilterOpts) (*DepositMainnetLogTokenAdminRemovedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogTokenAdminRemoved")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogTokenAdminRemovedIterator{contract: _DepositMainnet.contract, event: "LogTokenAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchLogTokenAdminRemoved is a free log subscription operation binding the contract event 0xfa49aecb996ea8d99950bb051552dfcc0b5460a0bb209867a1ed8067c32c2177.
//
// Solidity: event LogTokenAdminRemoved(address tokenAdmin)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogTokenAdminRemoved(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogTokenAdminRemoved) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogTokenAdminRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogTokenAdminRemoved)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogTokenAdminRemoved", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogTokenAdminRemoved(log types.Log) (*DepositMainnetLogTokenAdminRemoved, error) {
	event := new(DepositMainnetLogTokenAdminRemoved)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogTokenAdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogTokenRegisteredIterator is returned from FilterLogTokenRegistered and is used to iterate over the raw logs and unpacked data for LogTokenRegistered events raised by the DepositMainnet contract.
type DepositMainnetLogTokenRegisteredIterator struct {
	Event *DepositMainnetLogTokenRegistered // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogTokenRegistered)
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
		it.Event = new(DepositMainnetLogTokenRegistered)
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
func (it *DepositMainnetLogTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogTokenRegistered represents a LogTokenRegistered event raised by the DepositMainnet contract.
type DepositMainnetLogTokenRegistered struct {
	AssetType *big.Int
	AssetInfo []byte
	Quantum   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogTokenRegistered is a free log retrieval operation binding the contract event 0x7a0efbc885500f3b4a895231945be4520e4c0ba5ef7274a225a0272c81ccbcb7.
//
// Solidity: event LogTokenRegistered(uint256 assetType, bytes assetInfo, uint256 quantum)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogTokenRegistered(opts *bind.FilterOpts) (*DepositMainnetLogTokenRegisteredIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogTokenRegistered")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogTokenRegisteredIterator{contract: _DepositMainnet.contract, event: "LogTokenRegistered", logs: logs, sub: sub}, nil
}

// WatchLogTokenRegistered is a free log subscription operation binding the contract event 0x7a0efbc885500f3b4a895231945be4520e4c0ba5ef7274a225a0272c81ccbcb7.
//
// Solidity: event LogTokenRegistered(uint256 assetType, bytes assetInfo, uint256 quantum)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogTokenRegistered(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogTokenRegistered) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogTokenRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogTokenRegistered)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogTokenRegistered", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogTokenRegistered(log types.Log) (*DepositMainnetLogTokenRegistered, error) {
	event := new(DepositMainnetLogTokenRegistered)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogTokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogUnFrozenIterator is returned from FilterLogUnFrozen and is used to iterate over the raw logs and unpacked data for LogUnFrozen events raised by the DepositMainnet contract.
type DepositMainnetLogUnFrozenIterator struct {
	Event *DepositMainnetLogUnFrozen // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogUnFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogUnFrozen)
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
		it.Event = new(DepositMainnetLogUnFrozen)
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
func (it *DepositMainnetLogUnFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogUnFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogUnFrozen represents a LogUnFrozen event raised by the DepositMainnet contract.
type DepositMainnetLogUnFrozen struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogUnFrozen is a free log retrieval operation binding the contract event 0x07017fe9180629cfffba412f65a9affcf9a121de02294179f5c058f881dcc9f8.
//
// Solidity: event LogUnFrozen()
func (_DepositMainnet *DepositMainnetFilterer) FilterLogUnFrozen(opts *bind.FilterOpts) (*DepositMainnetLogUnFrozenIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogUnFrozen")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogUnFrozenIterator{contract: _DepositMainnet.contract, event: "LogUnFrozen", logs: logs, sub: sub}, nil
}

// WatchLogUnFrozen is a free log subscription operation binding the contract event 0x07017fe9180629cfffba412f65a9affcf9a121de02294179f5c058f881dcc9f8.
//
// Solidity: event LogUnFrozen()
func (_DepositMainnet *DepositMainnetFilterer) WatchLogUnFrozen(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogUnFrozen) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogUnFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogUnFrozen)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogUnFrozen", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogUnFrozen(log types.Log) (*DepositMainnetLogUnFrozen, error) {
	event := new(DepositMainnetLogUnFrozen)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogUnFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogUserRegisteredIterator is returned from FilterLogUserRegistered and is used to iterate over the raw logs and unpacked data for LogUserRegistered events raised by the DepositMainnet contract.
type DepositMainnetLogUserRegisteredIterator struct {
	Event *DepositMainnetLogUserRegistered // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogUserRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogUserRegistered)
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
		it.Event = new(DepositMainnetLogUserRegistered)
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
func (it *DepositMainnetLogUserRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogUserRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogUserRegistered represents a LogUserRegistered event raised by the DepositMainnet contract.
type DepositMainnetLogUserRegistered struct {
	EthKey   common.Address
	StarkKey *big.Int
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogUserRegistered is a free log retrieval operation binding the contract event 0xcab1cf17c190e4e2195a7b8f7b362023246fa774390432b4704ab6b29d56b07b.
//
// Solidity: event LogUserRegistered(address ethKey, uint256 starkKey, address sender)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogUserRegistered(opts *bind.FilterOpts) (*DepositMainnetLogUserRegisteredIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogUserRegistered")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogUserRegisteredIterator{contract: _DepositMainnet.contract, event: "LogUserRegistered", logs: logs, sub: sub}, nil
}

// WatchLogUserRegistered is a free log subscription operation binding the contract event 0xcab1cf17c190e4e2195a7b8f7b362023246fa774390432b4704ab6b29d56b07b.
//
// Solidity: event LogUserRegistered(address ethKey, uint256 starkKey, address sender)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogUserRegistered(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogUserRegistered) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogUserRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogUserRegistered)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogUserRegistered", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogUserRegistered(log types.Log) (*DepositMainnetLogUserRegistered, error) {
	event := new(DepositMainnetLogUserRegistered)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogUserRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogVaultBalanceChangeAppliedIterator is returned from FilterLogVaultBalanceChangeApplied and is used to iterate over the raw logs and unpacked data for LogVaultBalanceChangeApplied events raised by the DepositMainnet contract.
type DepositMainnetLogVaultBalanceChangeAppliedIterator struct {
	Event *DepositMainnetLogVaultBalanceChangeApplied // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogVaultBalanceChangeAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogVaultBalanceChangeApplied)
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
		it.Event = new(DepositMainnetLogVaultBalanceChangeApplied)
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
func (it *DepositMainnetLogVaultBalanceChangeAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogVaultBalanceChangeAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogVaultBalanceChangeApplied represents a LogVaultBalanceChangeApplied event raised by the DepositMainnet contract.
type DepositMainnetLogVaultBalanceChangeApplied struct {
	EthKey                common.Address
	AssetId               *big.Int
	VaultId               *big.Int
	QuantizedAmountChange *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterLogVaultBalanceChangeApplied is a free log retrieval operation binding the contract event 0x2b2b583bb5166d03b87a8e7c2785e61530ab776400fb69e1bcb13b33afef2b58.
//
// Solidity: event LogVaultBalanceChangeApplied(address ethKey, uint256 assetId, uint256 vaultId, int256 quantizedAmountChange)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogVaultBalanceChangeApplied(opts *bind.FilterOpts) (*DepositMainnetLogVaultBalanceChangeAppliedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogVaultBalanceChangeApplied")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogVaultBalanceChangeAppliedIterator{contract: _DepositMainnet.contract, event: "LogVaultBalanceChangeApplied", logs: logs, sub: sub}, nil
}

// WatchLogVaultBalanceChangeApplied is a free log subscription operation binding the contract event 0x2b2b583bb5166d03b87a8e7c2785e61530ab776400fb69e1bcb13b33afef2b58.
//
// Solidity: event LogVaultBalanceChangeApplied(address ethKey, uint256 assetId, uint256 vaultId, int256 quantizedAmountChange)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogVaultBalanceChangeApplied(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogVaultBalanceChangeApplied) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogVaultBalanceChangeApplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogVaultBalanceChangeApplied)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogVaultBalanceChangeApplied", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogVaultBalanceChangeApplied(log types.Log) (*DepositMainnetLogVaultBalanceChangeApplied, error) {
	event := new(DepositMainnetLogVaultBalanceChangeApplied)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogVaultBalanceChangeApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogVaultWithdrawalLockSetIterator is returned from FilterLogVaultWithdrawalLockSet and is used to iterate over the raw logs and unpacked data for LogVaultWithdrawalLockSet events raised by the DepositMainnet contract.
type DepositMainnetLogVaultWithdrawalLockSetIterator struct {
	Event *DepositMainnetLogVaultWithdrawalLockSet // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogVaultWithdrawalLockSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogVaultWithdrawalLockSet)
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
		it.Event = new(DepositMainnetLogVaultWithdrawalLockSet)
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
func (it *DepositMainnetLogVaultWithdrawalLockSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogVaultWithdrawalLockSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogVaultWithdrawalLockSet represents a LogVaultWithdrawalLockSet event raised by the DepositMainnet contract.
type DepositMainnetLogVaultWithdrawalLockSet struct {
	EthKey      common.Address
	AssetId     *big.Int
	VaultId     *big.Int
	TimeRelease *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogVaultWithdrawalLockSet is a free log retrieval operation binding the contract event 0xcd08744f6e50d6b64d69f8a5eddc7f66e34db287f51c1da789ff770ebc7cf73e.
//
// Solidity: event LogVaultWithdrawalLockSet(address ethKey, uint256 assetId, uint256 vaultId, uint256 timeRelease)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogVaultWithdrawalLockSet(opts *bind.FilterOpts) (*DepositMainnetLogVaultWithdrawalLockSetIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogVaultWithdrawalLockSet")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogVaultWithdrawalLockSetIterator{contract: _DepositMainnet.contract, event: "LogVaultWithdrawalLockSet", logs: logs, sub: sub}, nil
}

// WatchLogVaultWithdrawalLockSet is a free log subscription operation binding the contract event 0xcd08744f6e50d6b64d69f8a5eddc7f66e34db287f51c1da789ff770ebc7cf73e.
//
// Solidity: event LogVaultWithdrawalLockSet(address ethKey, uint256 assetId, uint256 vaultId, uint256 timeRelease)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogVaultWithdrawalLockSet(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogVaultWithdrawalLockSet) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogVaultWithdrawalLockSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogVaultWithdrawalLockSet)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogVaultWithdrawalLockSet", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogVaultWithdrawalLockSet(log types.Log) (*DepositMainnetLogVaultWithdrawalLockSet, error) {
	event := new(DepositMainnetLogVaultWithdrawalLockSet)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogVaultWithdrawalLockSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogWithdrawalAllowedIterator is returned from FilterLogWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogWithdrawalAllowed events raised by the DepositMainnet contract.
type DepositMainnetLogWithdrawalAllowedIterator struct {
	Event *DepositMainnetLogWithdrawalAllowed // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogWithdrawalAllowed)
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
		it.Event = new(DepositMainnetLogWithdrawalAllowed)
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
func (it *DepositMainnetLogWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogWithdrawalAllowed represents a LogWithdrawalAllowed event raised by the DepositMainnet contract.
type DepositMainnetLogWithdrawalAllowed struct {
	OwnerKey           *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawalAllowed is a free log retrieval operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogWithdrawalAllowed(opts *bind.FilterOpts) (*DepositMainnetLogWithdrawalAllowedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogWithdrawalAllowedIterator{contract: _DepositMainnet.contract, event: "LogWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawalAllowed is a free log subscription operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogWithdrawalAllowed)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogWithdrawalAllowed", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogWithdrawalAllowed(log types.Log) (*DepositMainnetLogWithdrawalAllowed, error) {
	event := new(DepositMainnetLogWithdrawalAllowed)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogWithdrawalFromVaultIterator is returned from FilterLogWithdrawalFromVault and is used to iterate over the raw logs and unpacked data for LogWithdrawalFromVault events raised by the DepositMainnet contract.
type DepositMainnetLogWithdrawalFromVaultIterator struct {
	Event *DepositMainnetLogWithdrawalFromVault // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogWithdrawalFromVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogWithdrawalFromVault)
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
		it.Event = new(DepositMainnetLogWithdrawalFromVault)
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
func (it *DepositMainnetLogWithdrawalFromVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogWithdrawalFromVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogWithdrawalFromVault represents a LogWithdrawalFromVault event raised by the DepositMainnet contract.
type DepositMainnetLogWithdrawalFromVault struct {
	EthKey             common.Address
	AssetType          *big.Int
	AssetId            *big.Int
	VaultId            *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawalFromVault is a free log retrieval operation binding the contract event 0x824058caef851aafcaa37b343756b98e0c9e3a86be89eaa6a326d66df04e5a43.
//
// Solidity: event LogWithdrawalFromVault(address ethKey, uint256 assetType, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogWithdrawalFromVault(opts *bind.FilterOpts) (*DepositMainnetLogWithdrawalFromVaultIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogWithdrawalFromVault")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogWithdrawalFromVaultIterator{contract: _DepositMainnet.contract, event: "LogWithdrawalFromVault", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawalFromVault is a free log subscription operation binding the contract event 0x824058caef851aafcaa37b343756b98e0c9e3a86be89eaa6a326d66df04e5a43.
//
// Solidity: event LogWithdrawalFromVault(address ethKey, uint256 assetType, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogWithdrawalFromVault(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogWithdrawalFromVault) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogWithdrawalFromVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogWithdrawalFromVault)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogWithdrawalFromVault", log); err != nil {
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

// ParseLogWithdrawalFromVault is a log parse operation binding the contract event 0x824058caef851aafcaa37b343756b98e0c9e3a86be89eaa6a326d66df04e5a43.
//
// Solidity: event LogWithdrawalFromVault(address ethKey, uint256 assetType, uint256 assetId, uint256 vaultId, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_DepositMainnet *DepositMainnetFilterer) ParseLogWithdrawalFromVault(log types.Log) (*DepositMainnetLogWithdrawalFromVault, error) {
	event := new(DepositMainnetLogWithdrawalFromVault)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogWithdrawalFromVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogWithdrawalPerformedIterator is returned from FilterLogWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogWithdrawalPerformed events raised by the DepositMainnet contract.
type DepositMainnetLogWithdrawalPerformedIterator struct {
	Event *DepositMainnetLogWithdrawalPerformed // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogWithdrawalPerformed)
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
		it.Event = new(DepositMainnetLogWithdrawalPerformed)
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
func (it *DepositMainnetLogWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogWithdrawalPerformed represents a LogWithdrawalPerformed event raised by the DepositMainnet contract.
type DepositMainnetLogWithdrawalPerformed struct {
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
func (_DepositMainnet *DepositMainnetFilterer) FilterLogWithdrawalPerformed(opts *bind.FilterOpts) (*DepositMainnetLogWithdrawalPerformedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogWithdrawalPerformedIterator{contract: _DepositMainnet.contract, event: "LogWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawalPerformed is a free log subscription operation binding the contract event 0xb7477a7b93b2addc5272bbd7ad0986ef1c0d0bd265f26c3dc4bbe42727c2ac0c.
//
// Solidity: event LogWithdrawalPerformed(uint256 ownerKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogWithdrawalPerformed)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogWithdrawalPerformed", log); err != nil {
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
func (_DepositMainnet *DepositMainnetFilterer) ParseLogWithdrawalPerformed(log types.Log) (*DepositMainnetLogWithdrawalPerformed, error) {
	event := new(DepositMainnetLogWithdrawalPerformed)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositMainnetLogWithdrawalWithTokenIdPerformedIterator is returned from FilterLogWithdrawalWithTokenIdPerformed and is used to iterate over the raw logs and unpacked data for LogWithdrawalWithTokenIdPerformed events raised by the DepositMainnet contract.
type DepositMainnetLogWithdrawalWithTokenIdPerformedIterator struct {
	Event *DepositMainnetLogWithdrawalWithTokenIdPerformed // Event containing the contract specifics and raw log

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
func (it *DepositMainnetLogWithdrawalWithTokenIdPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositMainnetLogWithdrawalWithTokenIdPerformed)
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
		it.Event = new(DepositMainnetLogWithdrawalWithTokenIdPerformed)
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
func (it *DepositMainnetLogWithdrawalWithTokenIdPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositMainnetLogWithdrawalWithTokenIdPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositMainnetLogWithdrawalWithTokenIdPerformed represents a LogWithdrawalWithTokenIdPerformed event raised by the DepositMainnet contract.
type DepositMainnetLogWithdrawalWithTokenIdPerformed struct {
	OwnerKey           *big.Int
	AssetType          *big.Int
	TokenId            *big.Int
	AssetId            *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Recipient          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawalWithTokenIdPerformed is a free log retrieval operation binding the contract event 0xc6ba68235f3229e53f3a95cda25543ad54c0f6df2493a06c05fb930bea7966fe.
//
// Solidity: event LogWithdrawalWithTokenIdPerformed(uint256 ownerKey, uint256 assetType, uint256 tokenId, uint256 assetId, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_DepositMainnet *DepositMainnetFilterer) FilterLogWithdrawalWithTokenIdPerformed(opts *bind.FilterOpts) (*DepositMainnetLogWithdrawalWithTokenIdPerformedIterator, error) {

	logs, sub, err := _DepositMainnet.contract.FilterLogs(opts, "LogWithdrawalWithTokenIdPerformed")
	if err != nil {
		return nil, err
	}
	return &DepositMainnetLogWithdrawalWithTokenIdPerformedIterator{contract: _DepositMainnet.contract, event: "LogWithdrawalWithTokenIdPerformed", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawalWithTokenIdPerformed is a free log subscription operation binding the contract event 0xc6ba68235f3229e53f3a95cda25543ad54c0f6df2493a06c05fb930bea7966fe.
//
// Solidity: event LogWithdrawalWithTokenIdPerformed(uint256 ownerKey, uint256 assetType, uint256 tokenId, uint256 assetId, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_DepositMainnet *DepositMainnetFilterer) WatchLogWithdrawalWithTokenIdPerformed(opts *bind.WatchOpts, sink chan<- *DepositMainnetLogWithdrawalWithTokenIdPerformed) (event.Subscription, error) {

	logs, sub, err := _DepositMainnet.contract.WatchLogs(opts, "LogWithdrawalWithTokenIdPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositMainnetLogWithdrawalWithTokenIdPerformed)
				if err := _DepositMainnet.contract.UnpackLog(event, "LogWithdrawalWithTokenIdPerformed", log); err != nil {
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

// ParseLogWithdrawalWithTokenIdPerformed is a log parse operation binding the contract event 0xc6ba68235f3229e53f3a95cda25543ad54c0f6df2493a06c05fb930bea7966fe.
//
// Solidity: event LogWithdrawalWithTokenIdPerformed(uint256 ownerKey, uint256 assetType, uint256 tokenId, uint256 assetId, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_DepositMainnet *DepositMainnetFilterer) ParseLogWithdrawalWithTokenIdPerformed(log types.Log) (*DepositMainnetLogWithdrawalWithTokenIdPerformed, error) {
	event := new(DepositMainnetLogWithdrawalWithTokenIdPerformed)
	if err := _DepositMainnet.contract.UnpackLog(event, "LogWithdrawalWithTokenIdPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
