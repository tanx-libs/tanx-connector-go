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

// DepositPolygonMetaData contains all meta data concerning the DepositPolygon contract.
var DepositPolygonMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgedDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"withdrawalId\",\"type\":\"string\"}],\"name\":\"BridgedWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addFundsNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"authorize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"processedWithdrawalIds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeFundsNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"withdrawalId\",\"type\":\"string\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"withdrawalId\",\"type\":\"string\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DepositPolygonABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositPolygonMetaData.ABI instead.
var DepositPolygonABI = DepositPolygonMetaData.ABI

// DepositPolygon is an auto generated Go binding around an Ethereum contract.
type DepositPolygon struct {
	DepositPolygonCaller     // Read-only binding to the contract
	DepositPolygonTransactor // Write-only binding to the contract
	DepositPolygonFilterer   // Log filterer for contract events
}

// DepositPolygonCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositPolygonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositPolygonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositPolygonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositPolygonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositPolygonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositPolygonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositPolygonSession struct {
	Contract     *DepositPolygon   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositPolygonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositPolygonCallerSession struct {
	Contract *DepositPolygonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DepositPolygonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositPolygonTransactorSession struct {
	Contract     *DepositPolygonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DepositPolygonRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositPolygonRaw struct {
	Contract *DepositPolygon // Generic contract binding to access the raw methods on
}

// DepositPolygonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositPolygonCallerRaw struct {
	Contract *DepositPolygonCaller // Generic read-only contract binding to access the raw methods on
}

// DepositPolygonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositPolygonTransactorRaw struct {
	Contract *DepositPolygonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositPolygon creates a new instance of DepositPolygon, bound to a specific deployed contract.
func NewDepositPolygon(address common.Address, backend bind.ContractBackend) (*DepositPolygon, error) {
	contract, err := bindDepositPolygon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DepositPolygon{DepositPolygonCaller: DepositPolygonCaller{contract: contract}, DepositPolygonTransactor: DepositPolygonTransactor{contract: contract}, DepositPolygonFilterer: DepositPolygonFilterer{contract: contract}}, nil
}

// NewDepositPolygonCaller creates a new read-only instance of DepositPolygon, bound to a specific deployed contract.
func NewDepositPolygonCaller(address common.Address, caller bind.ContractCaller) (*DepositPolygonCaller, error) {
	contract, err := bindDepositPolygon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositPolygonCaller{contract: contract}, nil
}

// NewDepositPolygonTransactor creates a new write-only instance of DepositPolygon, bound to a specific deployed contract.
func NewDepositPolygonTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositPolygonTransactor, error) {
	contract, err := bindDepositPolygon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositPolygonTransactor{contract: contract}, nil
}

// NewDepositPolygonFilterer creates a new log filterer instance of DepositPolygon, bound to a specific deployed contract.
func NewDepositPolygonFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositPolygonFilterer, error) {
	contract, err := bindDepositPolygon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositPolygonFilterer{contract: contract}, nil
}

// bindDepositPolygon binds a generic wrapper to an already deployed contract.
func bindDepositPolygon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DepositPolygonMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositPolygon *DepositPolygonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositPolygon.Contract.DepositPolygonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositPolygon *DepositPolygonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositPolygon.Contract.DepositPolygonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositPolygon *DepositPolygonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositPolygon.Contract.DepositPolygonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositPolygon *DepositPolygonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositPolygon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositPolygon *DepositPolygonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositPolygon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositPolygon *DepositPolygonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositPolygon.Contract.contract.Transact(opts, method, params...)
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) view returns(bool)
func (_DepositPolygon *DepositPolygonCaller) Authorized(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DepositPolygon.contract.Call(opts, &out, "authorized", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) view returns(bool)
func (_DepositPolygon *DepositPolygonSession) Authorized(arg0 common.Address) (bool, error) {
	return _DepositPolygon.Contract.Authorized(&_DepositPolygon.CallOpts, arg0)
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) view returns(bool)
func (_DepositPolygon *DepositPolygonCallerSession) Authorized(arg0 common.Address) (bool, error) {
	return _DepositPolygon.Contract.Authorized(&_DepositPolygon.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DepositPolygon *DepositPolygonCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DepositPolygon.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DepositPolygon *DepositPolygonSession) Owner() (common.Address, error) {
	return _DepositPolygon.Contract.Owner(&_DepositPolygon.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DepositPolygon *DepositPolygonCallerSession) Owner() (common.Address, error) {
	return _DepositPolygon.Contract.Owner(&_DepositPolygon.CallOpts)
}

// ProcessedWithdrawalIds is a free data retrieval call binding the contract method 0x7729d644.
//
// Solidity: function processedWithdrawalIds(string ) view returns(bool)
func (_DepositPolygon *DepositPolygonCaller) ProcessedWithdrawalIds(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _DepositPolygon.contract.Call(opts, &out, "processedWithdrawalIds", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedWithdrawalIds is a free data retrieval call binding the contract method 0x7729d644.
//
// Solidity: function processedWithdrawalIds(string ) view returns(bool)
func (_DepositPolygon *DepositPolygonSession) ProcessedWithdrawalIds(arg0 string) (bool, error) {
	return _DepositPolygon.Contract.ProcessedWithdrawalIds(&_DepositPolygon.CallOpts, arg0)
}

// ProcessedWithdrawalIds is a free data retrieval call binding the contract method 0x7729d644.
//
// Solidity: function processedWithdrawalIds(string ) view returns(bool)
func (_DepositPolygon *DepositPolygonCallerSession) ProcessedWithdrawalIds(arg0 string) (bool, error) {
	return _DepositPolygon.Contract.ProcessedWithdrawalIds(&_DepositPolygon.CallOpts, arg0)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_DepositPolygon *DepositPolygonCaller) RenounceOwnership(opts *bind.CallOpts) error {
	var out []interface{}
	err := _DepositPolygon.contract.Call(opts, &out, "renounceOwnership")

	if err != nil {
		return err
	}

	return err

}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_DepositPolygon *DepositPolygonSession) RenounceOwnership() error {
	return _DepositPolygon.Contract.RenounceOwnership(&_DepositPolygon.CallOpts)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_DepositPolygon *DepositPolygonCallerSession) RenounceOwnership() error {
	return _DepositPolygon.Contract.RenounceOwnership(&_DepositPolygon.CallOpts)
}

// AddFunds is a paid mutator transaction binding the contract method 0xbc4b3365.
//
// Solidity: function addFunds(address token, uint256 amount) returns()
func (_DepositPolygon *DepositPolygonTransactor) AddFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositPolygon.contract.Transact(opts, "addFunds", token, amount)
}

// AddFunds is a paid mutator transaction binding the contract method 0xbc4b3365.
//
// Solidity: function addFunds(address token, uint256 amount) returns()
func (_DepositPolygon *DepositPolygonSession) AddFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositPolygon.Contract.AddFunds(&_DepositPolygon.TransactOpts, token, amount)
}

// AddFunds is a paid mutator transaction binding the contract method 0xbc4b3365.
//
// Solidity: function addFunds(address token, uint256 amount) returns()
func (_DepositPolygon *DepositPolygonTransactorSession) AddFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositPolygon.Contract.AddFunds(&_DepositPolygon.TransactOpts, token, amount)
}

// AddFundsNative is a paid mutator transaction binding the contract method 0x447e346f.
//
// Solidity: function addFundsNative() payable returns()
func (_DepositPolygon *DepositPolygonTransactor) AddFundsNative(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositPolygon.contract.Transact(opts, "addFundsNative")
}

// AddFundsNative is a paid mutator transaction binding the contract method 0x447e346f.
//
// Solidity: function addFundsNative() payable returns()
func (_DepositPolygon *DepositPolygonSession) AddFundsNative() (*types.Transaction, error) {
	return _DepositPolygon.Contract.AddFundsNative(&_DepositPolygon.TransactOpts)
}

// AddFundsNative is a paid mutator transaction binding the contract method 0x447e346f.
//
// Solidity: function addFundsNative() payable returns()
func (_DepositPolygon *DepositPolygonTransactorSession) AddFundsNative() (*types.Transaction, error) {
	return _DepositPolygon.Contract.AddFundsNative(&_DepositPolygon.TransactOpts)
}

// Authorize is a paid mutator transaction binding the contract method 0x2d1fb389.
//
// Solidity: function authorize(address user, bool value) returns()
func (_DepositPolygon *DepositPolygonTransactor) Authorize(opts *bind.TransactOpts, user common.Address, value bool) (*types.Transaction, error) {
	return _DepositPolygon.contract.Transact(opts, "authorize", user, value)
}

// Authorize is a paid mutator transaction binding the contract method 0x2d1fb389.
//
// Solidity: function authorize(address user, bool value) returns()
func (_DepositPolygon *DepositPolygonSession) Authorize(user common.Address, value bool) (*types.Transaction, error) {
	return _DepositPolygon.Contract.Authorize(&_DepositPolygon.TransactOpts, user, value)
}

// Authorize is a paid mutator transaction binding the contract method 0x2d1fb389.
//
// Solidity: function authorize(address user, bool value) returns()
func (_DepositPolygon *DepositPolygonTransactorSession) Authorize(user common.Address, value bool) (*types.Transaction, error) {
	return _DepositPolygon.Contract.Authorize(&_DepositPolygon.TransactOpts, user, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_DepositPolygon *DepositPolygonTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositPolygon.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_DepositPolygon *DepositPolygonSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositPolygon.Contract.Deposit(&_DepositPolygon.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_DepositPolygon *DepositPolygonTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositPolygon.Contract.Deposit(&_DepositPolygon.TransactOpts, token, amount)
}

// DepositNative is a paid mutator transaction binding the contract method 0xdb6b5246.
//
// Solidity: function depositNative() payable returns()
func (_DepositPolygon *DepositPolygonTransactor) DepositNative(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositPolygon.contract.Transact(opts, "depositNative")
}

// DepositNative is a paid mutator transaction binding the contract method 0xdb6b5246.
//
// Solidity: function depositNative() payable returns()
func (_DepositPolygon *DepositPolygonSession) DepositNative() (*types.Transaction, error) {
	return _DepositPolygon.Contract.DepositNative(&_DepositPolygon.TransactOpts)
}

// DepositNative is a paid mutator transaction binding the contract method 0xdb6b5246.
//
// Solidity: function depositNative() payable returns()
func (_DepositPolygon *DepositPolygonTransactorSession) DepositNative() (*types.Transaction, error) {
	return _DepositPolygon.Contract.DepositNative(&_DepositPolygon.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DepositPolygon *DepositPolygonTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositPolygon.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DepositPolygon *DepositPolygonSession) Initialize() (*types.Transaction, error) {
	return _DepositPolygon.Contract.Initialize(&_DepositPolygon.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DepositPolygon *DepositPolygonTransactorSession) Initialize() (*types.Transaction, error) {
	return _DepositPolygon.Contract.Initialize(&_DepositPolygon.TransactOpts)
}

// RemoveFunds is a paid mutator transaction binding the contract method 0xd6c9b6a5.
//
// Solidity: function removeFunds(address token, address to, uint256 amount) returns()
func (_DepositPolygon *DepositPolygonTransactor) RemoveFunds(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositPolygon.contract.Transact(opts, "removeFunds", token, to, amount)
}

// RemoveFunds is a paid mutator transaction binding the contract method 0xd6c9b6a5.
//
// Solidity: function removeFunds(address token, address to, uint256 amount) returns()
func (_DepositPolygon *DepositPolygonSession) RemoveFunds(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositPolygon.Contract.RemoveFunds(&_DepositPolygon.TransactOpts, token, to, amount)
}

// RemoveFunds is a paid mutator transaction binding the contract method 0xd6c9b6a5.
//
// Solidity: function removeFunds(address token, address to, uint256 amount) returns()
func (_DepositPolygon *DepositPolygonTransactorSession) RemoveFunds(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositPolygon.Contract.RemoveFunds(&_DepositPolygon.TransactOpts, token, to, amount)
}

// RemoveFundsNative is a paid mutator transaction binding the contract method 0x143531c0.
//
// Solidity: function removeFundsNative(address to, uint256 amount) returns()
func (_DepositPolygon *DepositPolygonTransactor) RemoveFundsNative(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositPolygon.contract.Transact(opts, "removeFundsNative", to, amount)
}

// RemoveFundsNative is a paid mutator transaction binding the contract method 0x143531c0.
//
// Solidity: function removeFundsNative(address to, uint256 amount) returns()
func (_DepositPolygon *DepositPolygonSession) RemoveFundsNative(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositPolygon.Contract.RemoveFundsNative(&_DepositPolygon.TransactOpts, to, amount)
}

// RemoveFundsNative is a paid mutator transaction binding the contract method 0x143531c0.
//
// Solidity: function removeFundsNative(address to, uint256 amount) returns()
func (_DepositPolygon *DepositPolygonTransactorSession) RemoveFundsNative(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositPolygon.Contract.RemoveFundsNative(&_DepositPolygon.TransactOpts, to, amount)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address newOwner) returns()
func (_DepositPolygon *DepositPolygonTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DepositPolygon.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address newOwner) returns()
func (_DepositPolygon *DepositPolygonSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _DepositPolygon.Contract.TransferOwner(&_DepositPolygon.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address newOwner) returns()
func (_DepositPolygon *DepositPolygonTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _DepositPolygon.Contract.TransferOwner(&_DepositPolygon.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DepositPolygon *DepositPolygonTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DepositPolygon.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DepositPolygon *DepositPolygonSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DepositPolygon.Contract.TransferOwnership(&_DepositPolygon.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DepositPolygon *DepositPolygonTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DepositPolygon.Contract.TransferOwnership(&_DepositPolygon.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4f065632.
//
// Solidity: function withdraw(address token, address to, uint256 amount, string withdrawalId) returns()
func (_DepositPolygon *DepositPolygonTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, withdrawalId string) (*types.Transaction, error) {
	return _DepositPolygon.contract.Transact(opts, "withdraw", token, to, amount, withdrawalId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4f065632.
//
// Solidity: function withdraw(address token, address to, uint256 amount, string withdrawalId) returns()
func (_DepositPolygon *DepositPolygonSession) Withdraw(token common.Address, to common.Address, amount *big.Int, withdrawalId string) (*types.Transaction, error) {
	return _DepositPolygon.Contract.Withdraw(&_DepositPolygon.TransactOpts, token, to, amount, withdrawalId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4f065632.
//
// Solidity: function withdraw(address token, address to, uint256 amount, string withdrawalId) returns()
func (_DepositPolygon *DepositPolygonTransactorSession) Withdraw(token common.Address, to common.Address, amount *big.Int, withdrawalId string) (*types.Transaction, error) {
	return _DepositPolygon.Contract.Withdraw(&_DepositPolygon.TransactOpts, token, to, amount, withdrawalId)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x7a78b9c7.
//
// Solidity: function withdrawNative(address to, uint256 amount, string withdrawalId) returns()
func (_DepositPolygon *DepositPolygonTransactor) WithdrawNative(opts *bind.TransactOpts, to common.Address, amount *big.Int, withdrawalId string) (*types.Transaction, error) {
	return _DepositPolygon.contract.Transact(opts, "withdrawNative", to, amount, withdrawalId)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x7a78b9c7.
//
// Solidity: function withdrawNative(address to, uint256 amount, string withdrawalId) returns()
func (_DepositPolygon *DepositPolygonSession) WithdrawNative(to common.Address, amount *big.Int, withdrawalId string) (*types.Transaction, error) {
	return _DepositPolygon.Contract.WithdrawNative(&_DepositPolygon.TransactOpts, to, amount, withdrawalId)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x7a78b9c7.
//
// Solidity: function withdrawNative(address to, uint256 amount, string withdrawalId) returns()
func (_DepositPolygon *DepositPolygonTransactorSession) WithdrawNative(to common.Address, amount *big.Int, withdrawalId string) (*types.Transaction, error) {
	return _DepositPolygon.Contract.WithdrawNative(&_DepositPolygon.TransactOpts, to, amount, withdrawalId)
}

// DepositPolygonBridgedDepositIterator is returned from FilterBridgedDeposit and is used to iterate over the raw logs and unpacked data for BridgedDeposit events raised by the DepositPolygon contract.
type DepositPolygonBridgedDepositIterator struct {
	Event *DepositPolygonBridgedDeposit // Event containing the contract specifics and raw log

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
func (it *DepositPolygonBridgedDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositPolygonBridgedDeposit)
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
		it.Event = new(DepositPolygonBridgedDeposit)
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
func (it *DepositPolygonBridgedDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositPolygonBridgedDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositPolygonBridgedDeposit represents a BridgedDeposit event raised by the DepositPolygon contract.
type DepositPolygonBridgedDeposit struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridgedDeposit is a free log retrieval operation binding the contract event 0x573284f4c36da6a8d8d84cd06662235f8a770cc98e8c80e304b8f382fdc3dca2.
//
// Solidity: event BridgedDeposit(address indexed user, address indexed token, uint256 amount)
func (_DepositPolygon *DepositPolygonFilterer) FilterBridgedDeposit(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*DepositPolygonBridgedDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DepositPolygon.contract.FilterLogs(opts, "BridgedDeposit", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &DepositPolygonBridgedDepositIterator{contract: _DepositPolygon.contract, event: "BridgedDeposit", logs: logs, sub: sub}, nil
}

// WatchBridgedDeposit is a free log subscription operation binding the contract event 0x573284f4c36da6a8d8d84cd06662235f8a770cc98e8c80e304b8f382fdc3dca2.
//
// Solidity: event BridgedDeposit(address indexed user, address indexed token, uint256 amount)
func (_DepositPolygon *DepositPolygonFilterer) WatchBridgedDeposit(opts *bind.WatchOpts, sink chan<- *DepositPolygonBridgedDeposit, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DepositPolygon.contract.WatchLogs(opts, "BridgedDeposit", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositPolygonBridgedDeposit)
				if err := _DepositPolygon.contract.UnpackLog(event, "BridgedDeposit", log); err != nil {
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

// ParseBridgedDeposit is a log parse operation binding the contract event 0x573284f4c36da6a8d8d84cd06662235f8a770cc98e8c80e304b8f382fdc3dca2.
//
// Solidity: event BridgedDeposit(address indexed user, address indexed token, uint256 amount)
func (_DepositPolygon *DepositPolygonFilterer) ParseBridgedDeposit(log types.Log) (*DepositPolygonBridgedDeposit, error) {
	event := new(DepositPolygonBridgedDeposit)
	if err := _DepositPolygon.contract.UnpackLog(event, "BridgedDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositPolygonBridgedWithdrawalIterator is returned from FilterBridgedWithdrawal and is used to iterate over the raw logs and unpacked data for BridgedWithdrawal events raised by the DepositPolygon contract.
type DepositPolygonBridgedWithdrawalIterator struct {
	Event *DepositPolygonBridgedWithdrawal // Event containing the contract specifics and raw log

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
func (it *DepositPolygonBridgedWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositPolygonBridgedWithdrawal)
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
		it.Event = new(DepositPolygonBridgedWithdrawal)
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
func (it *DepositPolygonBridgedWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositPolygonBridgedWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositPolygonBridgedWithdrawal represents a BridgedWithdrawal event raised by the DepositPolygon contract.
type DepositPolygonBridgedWithdrawal struct {
	User         common.Address
	Token        common.Address
	Amount       *big.Int
	WithdrawalId string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBridgedWithdrawal is a free log retrieval operation binding the contract event 0xe4f4f1fb3534fe80225d336f6e5a73007dc992e5f6740152bf13ed2a08f3851a.
//
// Solidity: event BridgedWithdrawal(address indexed user, address indexed token, uint256 amount, string withdrawalId)
func (_DepositPolygon *DepositPolygonFilterer) FilterBridgedWithdrawal(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*DepositPolygonBridgedWithdrawalIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DepositPolygon.contract.FilterLogs(opts, "BridgedWithdrawal", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &DepositPolygonBridgedWithdrawalIterator{contract: _DepositPolygon.contract, event: "BridgedWithdrawal", logs: logs, sub: sub}, nil
}

// WatchBridgedWithdrawal is a free log subscription operation binding the contract event 0xe4f4f1fb3534fe80225d336f6e5a73007dc992e5f6740152bf13ed2a08f3851a.
//
// Solidity: event BridgedWithdrawal(address indexed user, address indexed token, uint256 amount, string withdrawalId)
func (_DepositPolygon *DepositPolygonFilterer) WatchBridgedWithdrawal(opts *bind.WatchOpts, sink chan<- *DepositPolygonBridgedWithdrawal, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DepositPolygon.contract.WatchLogs(opts, "BridgedWithdrawal", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositPolygonBridgedWithdrawal)
				if err := _DepositPolygon.contract.UnpackLog(event, "BridgedWithdrawal", log); err != nil {
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

// ParseBridgedWithdrawal is a log parse operation binding the contract event 0xe4f4f1fb3534fe80225d336f6e5a73007dc992e5f6740152bf13ed2a08f3851a.
//
// Solidity: event BridgedWithdrawal(address indexed user, address indexed token, uint256 amount, string withdrawalId)
func (_DepositPolygon *DepositPolygonFilterer) ParseBridgedWithdrawal(log types.Log) (*DepositPolygonBridgedWithdrawal, error) {
	event := new(DepositPolygonBridgedWithdrawal)
	if err := _DepositPolygon.contract.UnpackLog(event, "BridgedWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositPolygonOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DepositPolygon contract.
type DepositPolygonOwnershipTransferredIterator struct {
	Event *DepositPolygonOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DepositPolygonOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositPolygonOwnershipTransferred)
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
		it.Event = new(DepositPolygonOwnershipTransferred)
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
func (it *DepositPolygonOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositPolygonOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositPolygonOwnershipTransferred represents a OwnershipTransferred event raised by the DepositPolygon contract.
type DepositPolygonOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DepositPolygon *DepositPolygonFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DepositPolygonOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DepositPolygon.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DepositPolygonOwnershipTransferredIterator{contract: _DepositPolygon.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DepositPolygon *DepositPolygonFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DepositPolygonOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DepositPolygon.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositPolygonOwnershipTransferred)
				if err := _DepositPolygon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DepositPolygon *DepositPolygonFilterer) ParseOwnershipTransferred(log types.Log) (*DepositPolygonOwnershipTransferred, error) {
	event := new(DepositPolygonOwnershipTransferred)
	if err := _DepositPolygon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
