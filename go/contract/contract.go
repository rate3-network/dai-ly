// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DelegatableDaiABI is the input ABI used to generate the binding from.
const DelegatableDaiABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TransferPreSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ApprovalPreSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"transferPreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"approvePreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"increaseApprovalPreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"decreaseApprovalPreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"transferFromPreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"transferPreSignedHashing\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"approvePreSignedHashing\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"increaseApprovalPreSignedHashing\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"decreaseApprovalPreSignedHashing\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"transferFromPreSignedHashing\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// DelegatableDai is an auto generated Go binding around an Ethereum contract.
type DelegatableDai struct {
	DelegatableDaiCaller     // Read-only binding to the contract
	DelegatableDaiTransactor // Write-only binding to the contract
	DelegatableDaiFilterer   // Log filterer for contract events
}

// DelegatableDaiCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegatableDaiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatableDaiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegatableDaiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatableDaiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegatableDaiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatableDaiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegatableDaiSession struct {
	Contract     *DelegatableDai   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelegatableDaiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegatableDaiCallerSession struct {
	Contract *DelegatableDaiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DelegatableDaiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegatableDaiTransactorSession struct {
	Contract     *DelegatableDaiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DelegatableDaiRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegatableDaiRaw struct {
	Contract *DelegatableDai // Generic contract binding to access the raw methods on
}

// DelegatableDaiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegatableDaiCallerRaw struct {
	Contract *DelegatableDaiCaller // Generic read-only contract binding to access the raw methods on
}

// DelegatableDaiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegatableDaiTransactorRaw struct {
	Contract *DelegatableDaiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegatableDai creates a new instance of DelegatableDai, bound to a specific deployed contract.
func NewDelegatableDai(address common.Address, backend bind.ContractBackend) (*DelegatableDai, error) {
	contract, err := bindDelegatableDai(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelegatableDai{DelegatableDaiCaller: DelegatableDaiCaller{contract: contract}, DelegatableDaiTransactor: DelegatableDaiTransactor{contract: contract}, DelegatableDaiFilterer: DelegatableDaiFilterer{contract: contract}}, nil
}

// NewDelegatableDaiCaller creates a new read-only instance of DelegatableDai, bound to a specific deployed contract.
func NewDelegatableDaiCaller(address common.Address, caller bind.ContractCaller) (*DelegatableDaiCaller, error) {
	contract, err := bindDelegatableDai(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegatableDaiCaller{contract: contract}, nil
}

// NewDelegatableDaiTransactor creates a new write-only instance of DelegatableDai, bound to a specific deployed contract.
func NewDelegatableDaiTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegatableDaiTransactor, error) {
	contract, err := bindDelegatableDai(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegatableDaiTransactor{contract: contract}, nil
}

// NewDelegatableDaiFilterer creates a new log filterer instance of DelegatableDai, bound to a specific deployed contract.
func NewDelegatableDaiFilterer(address common.Address, filterer bind.ContractFilterer) (*DelegatableDaiFilterer, error) {
	contract, err := bindDelegatableDai(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegatableDaiFilterer{contract: contract}, nil
}

// bindDelegatableDai binds a generic wrapper to an already deployed contract.
func bindDelegatableDai(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DelegatableDaiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegatableDai *DelegatableDaiRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DelegatableDai.Contract.DelegatableDaiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegatableDai *DelegatableDaiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegatableDai.Contract.DelegatableDaiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegatableDai *DelegatableDaiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegatableDai.Contract.DelegatableDaiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegatableDai *DelegatableDaiCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DelegatableDai.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegatableDai *DelegatableDaiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegatableDai.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegatableDai *DelegatableDaiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegatableDai.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_DelegatableDai *DelegatableDaiCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DelegatableDai.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_DelegatableDai *DelegatableDaiSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _DelegatableDai.Contract.Allowance(&_DelegatableDai.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_DelegatableDai *DelegatableDaiCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _DelegatableDai.Contract.Allowance(&_DelegatableDai.CallOpts, _owner, _spender)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_DelegatableDai *DelegatableDaiCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DelegatableDai.contract.Call(opts, out, "allowed", arg0, arg1)
	return *ret0, err
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_DelegatableDai *DelegatableDaiSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DelegatableDai.Contract.Allowed(&_DelegatableDai.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_DelegatableDai *DelegatableDaiCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DelegatableDai.Contract.Allowed(&_DelegatableDai.CallOpts, arg0, arg1)
}

// ApprovePreSignedHashing is a free data retrieval call binding the contract method 0xf7ac9c2e.
//
// Solidity: function approvePreSignedHashing(_token address, _spender address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_DelegatableDai *DelegatableDaiCaller) ApprovePreSignedHashing(opts *bind.CallOpts, _token common.Address, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DelegatableDai.contract.Call(opts, out, "approvePreSignedHashing", _token, _spender, _value, _fee, _nonce)
	return *ret0, err
}

// ApprovePreSignedHashing is a free data retrieval call binding the contract method 0xf7ac9c2e.
//
// Solidity: function approvePreSignedHashing(_token address, _spender address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_DelegatableDai *DelegatableDaiSession) ApprovePreSignedHashing(_token common.Address, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _DelegatableDai.Contract.ApprovePreSignedHashing(&_DelegatableDai.CallOpts, _token, _spender, _value, _fee, _nonce)
}

// ApprovePreSignedHashing is a free data retrieval call binding the contract method 0xf7ac9c2e.
//
// Solidity: function approvePreSignedHashing(_token address, _spender address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_DelegatableDai *DelegatableDaiCallerSession) ApprovePreSignedHashing(_token common.Address, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _DelegatableDai.Contract.ApprovePreSignedHashing(&_DelegatableDai.CallOpts, _token, _spender, _value, _fee, _nonce)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_DelegatableDai *DelegatableDaiCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DelegatableDai.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_DelegatableDai *DelegatableDaiSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DelegatableDai.Contract.BalanceOf(&_DelegatableDai.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_DelegatableDai *DelegatableDaiCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DelegatableDai.Contract.BalanceOf(&_DelegatableDai.CallOpts, _owner)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_DelegatableDai *DelegatableDaiCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DelegatableDai.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_DelegatableDai *DelegatableDaiSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _DelegatableDai.Contract.Balances(&_DelegatableDai.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_DelegatableDai *DelegatableDaiCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _DelegatableDai.Contract.Balances(&_DelegatableDai.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_DelegatableDai *DelegatableDaiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _DelegatableDai.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_DelegatableDai *DelegatableDaiSession) Decimals() (uint8, error) {
	return _DelegatableDai.Contract.Decimals(&_DelegatableDai.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_DelegatableDai *DelegatableDaiCallerSession) Decimals() (uint8, error) {
	return _DelegatableDai.Contract.Decimals(&_DelegatableDai.CallOpts)
}

// DecreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0x59388d78.
//
// Solidity: function decreaseApprovalPreSignedHashing(_token address, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_DelegatableDai *DelegatableDaiCaller) DecreaseApprovalPreSignedHashing(opts *bind.CallOpts, _token common.Address, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DelegatableDai.contract.Call(opts, out, "decreaseApprovalPreSignedHashing", _token, _spender, _subtractedValue, _fee, _nonce)
	return *ret0, err
}

// DecreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0x59388d78.
//
// Solidity: function decreaseApprovalPreSignedHashing(_token address, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_DelegatableDai *DelegatableDaiSession) DecreaseApprovalPreSignedHashing(_token common.Address, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _DelegatableDai.Contract.DecreaseApprovalPreSignedHashing(&_DelegatableDai.CallOpts, _token, _spender, _subtractedValue, _fee, _nonce)
}

// DecreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0x59388d78.
//
// Solidity: function decreaseApprovalPreSignedHashing(_token address, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_DelegatableDai *DelegatableDaiCallerSession) DecreaseApprovalPreSignedHashing(_token common.Address, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _DelegatableDai.Contract.DecreaseApprovalPreSignedHashing(&_DelegatableDai.CallOpts, _token, _spender, _subtractedValue, _fee, _nonce)
}

// IncreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0xa45f71ff.
//
// Solidity: function increaseApprovalPreSignedHashing(_token address, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_DelegatableDai *DelegatableDaiCaller) IncreaseApprovalPreSignedHashing(opts *bind.CallOpts, _token common.Address, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DelegatableDai.contract.Call(opts, out, "increaseApprovalPreSignedHashing", _token, _spender, _addedValue, _fee, _nonce)
	return *ret0, err
}

// IncreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0xa45f71ff.
//
// Solidity: function increaseApprovalPreSignedHashing(_token address, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_DelegatableDai *DelegatableDaiSession) IncreaseApprovalPreSignedHashing(_token common.Address, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _DelegatableDai.Contract.IncreaseApprovalPreSignedHashing(&_DelegatableDai.CallOpts, _token, _spender, _addedValue, _fee, _nonce)
}

// IncreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0xa45f71ff.
//
// Solidity: function increaseApprovalPreSignedHashing(_token address, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_DelegatableDai *DelegatableDaiCallerSession) IncreaseApprovalPreSignedHashing(_token common.Address, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _DelegatableDai.Contract.IncreaseApprovalPreSignedHashing(&_DelegatableDai.CallOpts, _token, _spender, _addedValue, _fee, _nonce)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DelegatableDai *DelegatableDaiCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DelegatableDai.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DelegatableDai *DelegatableDaiSession) Name() (string, error) {
	return _DelegatableDai.Contract.Name(&_DelegatableDai.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DelegatableDai *DelegatableDaiCallerSession) Name() (string, error) {
	return _DelegatableDai.Contract.Name(&_DelegatableDai.CallOpts)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(hash bytes32, sig bytes) constant returns(address)
func (_DelegatableDai *DelegatableDaiCaller) Recover(opts *bind.CallOpts, hash [32]byte, sig []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DelegatableDai.contract.Call(opts, out, "recover", hash, sig)
	return *ret0, err
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(hash bytes32, sig bytes) constant returns(address)
func (_DelegatableDai *DelegatableDaiSession) Recover(hash [32]byte, sig []byte) (common.Address, error) {
	return _DelegatableDai.Contract.Recover(&_DelegatableDai.CallOpts, hash, sig)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(hash bytes32, sig bytes) constant returns(address)
func (_DelegatableDai *DelegatableDaiCallerSession) Recover(hash [32]byte, sig []byte) (common.Address, error) {
	return _DelegatableDai.Contract.Recover(&_DelegatableDai.CallOpts, hash, sig)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DelegatableDai *DelegatableDaiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DelegatableDai.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DelegatableDai *DelegatableDaiSession) Symbol() (string, error) {
	return _DelegatableDai.Contract.Symbol(&_DelegatableDai.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DelegatableDai *DelegatableDaiCallerSession) Symbol() (string, error) {
	return _DelegatableDai.Contract.Symbol(&_DelegatableDai.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DelegatableDai *DelegatableDaiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DelegatableDai.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DelegatableDai *DelegatableDaiSession) TotalSupply() (*big.Int, error) {
	return _DelegatableDai.Contract.TotalSupply(&_DelegatableDai.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DelegatableDai *DelegatableDaiCallerSession) TotalSupply() (*big.Int, error) {
	return _DelegatableDai.Contract.TotalSupply(&_DelegatableDai.CallOpts)
}

// TransferFromPreSignedHashing is a free data retrieval call binding the contract method 0xb7656dc5.
//
// Solidity: function transferFromPreSignedHashing(_token address, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_DelegatableDai *DelegatableDaiCaller) TransferFromPreSignedHashing(opts *bind.CallOpts, _token common.Address, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DelegatableDai.contract.Call(opts, out, "transferFromPreSignedHashing", _token, _from, _to, _value, _fee, _nonce)
	return *ret0, err
}

// TransferFromPreSignedHashing is a free data retrieval call binding the contract method 0xb7656dc5.
//
// Solidity: function transferFromPreSignedHashing(_token address, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_DelegatableDai *DelegatableDaiSession) TransferFromPreSignedHashing(_token common.Address, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _DelegatableDai.Contract.TransferFromPreSignedHashing(&_DelegatableDai.CallOpts, _token, _from, _to, _value, _fee, _nonce)
}

// TransferFromPreSignedHashing is a free data retrieval call binding the contract method 0xb7656dc5.
//
// Solidity: function transferFromPreSignedHashing(_token address, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_DelegatableDai *DelegatableDaiCallerSession) TransferFromPreSignedHashing(_token common.Address, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _DelegatableDai.Contract.TransferFromPreSignedHashing(&_DelegatableDai.CallOpts, _token, _from, _to, _value, _fee, _nonce)
}

// TransferPreSignedHashing is a free data retrieval call binding the contract method 0x15420b71.
//
// Solidity: function transferPreSignedHashing(_token address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_DelegatableDai *DelegatableDaiCaller) TransferPreSignedHashing(opts *bind.CallOpts, _token common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DelegatableDai.contract.Call(opts, out, "transferPreSignedHashing", _token, _to, _value, _fee, _nonce)
	return *ret0, err
}

// TransferPreSignedHashing is a free data retrieval call binding the contract method 0x15420b71.
//
// Solidity: function transferPreSignedHashing(_token address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_DelegatableDai *DelegatableDaiSession) TransferPreSignedHashing(_token common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _DelegatableDai.Contract.TransferPreSignedHashing(&_DelegatableDai.CallOpts, _token, _to, _value, _fee, _nonce)
}

// TransferPreSignedHashing is a free data retrieval call binding the contract method 0x15420b71.
//
// Solidity: function transferPreSignedHashing(_token address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_DelegatableDai *DelegatableDaiCallerSession) TransferPreSignedHashing(_token common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _DelegatableDai.Contract.TransferPreSignedHashing(&_DelegatableDai.CallOpts, _token, _to, _value, _fee, _nonce)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_DelegatableDai *DelegatableDaiTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_DelegatableDai *DelegatableDaiSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.Contract.Approve(&_DelegatableDai.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_DelegatableDai *DelegatableDaiTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.Contract.Approve(&_DelegatableDai.TransactOpts, _spender, _value)
}

// ApprovePreSigned is a paid mutator transaction binding the contract method 0x617b390b.
//
// Solidity: function approvePreSigned(_signature bytes, _spender address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_DelegatableDai *DelegatableDaiTransactor) ApprovePreSigned(opts *bind.TransactOpts, _signature []byte, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.contract.Transact(opts, "approvePreSigned", _signature, _spender, _value, _fee, _nonce)
}

// ApprovePreSigned is a paid mutator transaction binding the contract method 0x617b390b.
//
// Solidity: function approvePreSigned(_signature bytes, _spender address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_DelegatableDai *DelegatableDaiSession) ApprovePreSigned(_signature []byte, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.Contract.ApprovePreSigned(&_DelegatableDai.TransactOpts, _signature, _spender, _value, _fee, _nonce)
}

// ApprovePreSigned is a paid mutator transaction binding the contract method 0x617b390b.
//
// Solidity: function approvePreSigned(_signature bytes, _spender address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_DelegatableDai *DelegatableDaiTransactorSession) ApprovePreSigned(_signature []byte, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.Contract.ApprovePreSigned(&_DelegatableDai.TransactOpts, _signature, _spender, _value, _fee, _nonce)
}

// DecreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0x8be52783.
//
// Solidity: function decreaseApprovalPreSigned(_signature bytes, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_DelegatableDai *DelegatableDaiTransactor) DecreaseApprovalPreSigned(opts *bind.TransactOpts, _signature []byte, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.contract.Transact(opts, "decreaseApprovalPreSigned", _signature, _spender, _subtractedValue, _fee, _nonce)
}

// DecreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0x8be52783.
//
// Solidity: function decreaseApprovalPreSigned(_signature bytes, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_DelegatableDai *DelegatableDaiSession) DecreaseApprovalPreSigned(_signature []byte, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.Contract.DecreaseApprovalPreSigned(&_DelegatableDai.TransactOpts, _signature, _spender, _subtractedValue, _fee, _nonce)
}

// DecreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0x8be52783.
//
// Solidity: function decreaseApprovalPreSigned(_signature bytes, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_DelegatableDai *DelegatableDaiTransactorSession) DecreaseApprovalPreSigned(_signature []byte, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.Contract.DecreaseApprovalPreSigned(&_DelegatableDai.TransactOpts, _signature, _spender, _subtractedValue, _fee, _nonce)
}

// IncreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0xadb8249e.
//
// Solidity: function increaseApprovalPreSigned(_signature bytes, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_DelegatableDai *DelegatableDaiTransactor) IncreaseApprovalPreSigned(opts *bind.TransactOpts, _signature []byte, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.contract.Transact(opts, "increaseApprovalPreSigned", _signature, _spender, _addedValue, _fee, _nonce)
}

// IncreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0xadb8249e.
//
// Solidity: function increaseApprovalPreSigned(_signature bytes, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_DelegatableDai *DelegatableDaiSession) IncreaseApprovalPreSigned(_signature []byte, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.Contract.IncreaseApprovalPreSigned(&_DelegatableDai.TransactOpts, _signature, _spender, _addedValue, _fee, _nonce)
}

// IncreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0xadb8249e.
//
// Solidity: function increaseApprovalPreSigned(_signature bytes, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_DelegatableDai *DelegatableDaiTransactorSession) IncreaseApprovalPreSigned(_signature []byte, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.Contract.IncreaseApprovalPreSigned(&_DelegatableDai.TransactOpts, _signature, _spender, _addedValue, _fee, _nonce)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_DelegatableDai *DelegatableDaiTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_DelegatableDai *DelegatableDaiSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.Contract.Transfer(&_DelegatableDai.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_DelegatableDai *DelegatableDaiTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.Contract.Transfer(&_DelegatableDai.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_DelegatableDai *DelegatableDaiTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_DelegatableDai *DelegatableDaiSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.Contract.TransferFrom(&_DelegatableDai.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_DelegatableDai *DelegatableDaiTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.Contract.TransferFrom(&_DelegatableDai.TransactOpts, _from, _to, _value)
}

// TransferFromPreSigned is a paid mutator transaction binding the contract method 0xbca50515.
//
// Solidity: function transferFromPreSigned(_signature bytes, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_DelegatableDai *DelegatableDaiTransactor) TransferFromPreSigned(opts *bind.TransactOpts, _signature []byte, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.contract.Transact(opts, "transferFromPreSigned", _signature, _from, _to, _value, _fee, _nonce)
}

// TransferFromPreSigned is a paid mutator transaction binding the contract method 0xbca50515.
//
// Solidity: function transferFromPreSigned(_signature bytes, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_DelegatableDai *DelegatableDaiSession) TransferFromPreSigned(_signature []byte, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.Contract.TransferFromPreSigned(&_DelegatableDai.TransactOpts, _signature, _from, _to, _value, _fee, _nonce)
}

// TransferFromPreSigned is a paid mutator transaction binding the contract method 0xbca50515.
//
// Solidity: function transferFromPreSigned(_signature bytes, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_DelegatableDai *DelegatableDaiTransactorSession) TransferFromPreSigned(_signature []byte, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.Contract.TransferFromPreSigned(&_DelegatableDai.TransactOpts, _signature, _from, _to, _value, _fee, _nonce)
}

// TransferPreSigned is a paid mutator transaction binding the contract method 0x1296830d.
//
// Solidity: function transferPreSigned(_signature bytes, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_DelegatableDai *DelegatableDaiTransactor) TransferPreSigned(opts *bind.TransactOpts, _signature []byte, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.contract.Transact(opts, "transferPreSigned", _signature, _to, _value, _fee, _nonce)
}

// TransferPreSigned is a paid mutator transaction binding the contract method 0x1296830d.
//
// Solidity: function transferPreSigned(_signature bytes, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_DelegatableDai *DelegatableDaiSession) TransferPreSigned(_signature []byte, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.Contract.TransferPreSigned(&_DelegatableDai.TransactOpts, _signature, _to, _value, _fee, _nonce)
}

// TransferPreSigned is a paid mutator transaction binding the contract method 0x1296830d.
//
// Solidity: function transferPreSigned(_signature bytes, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_DelegatableDai *DelegatableDaiTransactorSession) TransferPreSigned(_signature []byte, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _DelegatableDai.Contract.TransferPreSigned(&_DelegatableDai.TransactOpts, _signature, _to, _value, _fee, _nonce)
}

// DelegatableDaiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DelegatableDai contract.
type DelegatableDaiApprovalIterator struct {
	Event *DelegatableDaiApproval // Event containing the contract specifics and raw log

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
func (it *DelegatableDaiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatableDaiApproval)
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
		it.Event = new(DelegatableDaiApproval)
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
func (it *DelegatableDaiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatableDaiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatableDaiApproval represents a Approval event raised by the DelegatableDai contract.
type DelegatableDaiApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_DelegatableDai *DelegatableDaiFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*DelegatableDaiApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _DelegatableDai.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &DelegatableDaiApprovalIterator{contract: _DelegatableDai.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_DelegatableDai *DelegatableDaiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DelegatableDaiApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _DelegatableDai.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatableDaiApproval)
				if err := _DelegatableDai.contract.UnpackLog(event, "Approval", log); err != nil {
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

// DelegatableDaiApprovalPreSignedIterator is returned from FilterApprovalPreSigned and is used to iterate over the raw logs and unpacked data for ApprovalPreSigned events raised by the DelegatableDai contract.
type DelegatableDaiApprovalPreSignedIterator struct {
	Event *DelegatableDaiApprovalPreSigned // Event containing the contract specifics and raw log

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
func (it *DelegatableDaiApprovalPreSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatableDaiApprovalPreSigned)
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
		it.Event = new(DelegatableDaiApprovalPreSigned)
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
func (it *DelegatableDaiApprovalPreSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatableDaiApprovalPreSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatableDaiApprovalPreSigned represents a ApprovalPreSigned event raised by the DelegatableDai contract.
type DelegatableDaiApprovalPreSigned struct {
	From     common.Address
	To       common.Address
	Delegate common.Address
	Amount   *big.Int
	Fee      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalPreSigned is a free log retrieval operation binding the contract event 0x43a220267705e74ee2ceafd46afc841850db6f85a662189a7def697bbdd90ffb.
//
// Solidity: e ApprovalPreSigned(from indexed address, to indexed address, delegate indexed address, amount uint256, fee uint256)
func (_DelegatableDai *DelegatableDaiFilterer) FilterApprovalPreSigned(opts *bind.FilterOpts, from []common.Address, to []common.Address, delegate []common.Address) (*DelegatableDaiApprovalPreSignedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _DelegatableDai.contract.FilterLogs(opts, "ApprovalPreSigned", fromRule, toRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return &DelegatableDaiApprovalPreSignedIterator{contract: _DelegatableDai.contract, event: "ApprovalPreSigned", logs: logs, sub: sub}, nil
}

// WatchApprovalPreSigned is a free log subscription operation binding the contract event 0x43a220267705e74ee2ceafd46afc841850db6f85a662189a7def697bbdd90ffb.
//
// Solidity: e ApprovalPreSigned(from indexed address, to indexed address, delegate indexed address, amount uint256, fee uint256)
func (_DelegatableDai *DelegatableDaiFilterer) WatchApprovalPreSigned(opts *bind.WatchOpts, sink chan<- *DelegatableDaiApprovalPreSigned, from []common.Address, to []common.Address, delegate []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _DelegatableDai.contract.WatchLogs(opts, "ApprovalPreSigned", fromRule, toRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatableDaiApprovalPreSigned)
				if err := _DelegatableDai.contract.UnpackLog(event, "ApprovalPreSigned", log); err != nil {
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

// DelegatableDaiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DelegatableDai contract.
type DelegatableDaiTransferIterator struct {
	Event *DelegatableDaiTransfer // Event containing the contract specifics and raw log

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
func (it *DelegatableDaiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatableDaiTransfer)
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
		it.Event = new(DelegatableDaiTransfer)
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
func (it *DelegatableDaiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatableDaiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatableDaiTransfer represents a Transfer event raised by the DelegatableDai contract.
type DelegatableDaiTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_DelegatableDai *DelegatableDaiFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*DelegatableDaiTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _DelegatableDai.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &DelegatableDaiTransferIterator{contract: _DelegatableDai.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_DelegatableDai *DelegatableDaiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DelegatableDaiTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _DelegatableDai.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatableDaiTransfer)
				if err := _DelegatableDai.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// DelegatableDaiTransferPreSignedIterator is returned from FilterTransferPreSigned and is used to iterate over the raw logs and unpacked data for TransferPreSigned events raised by the DelegatableDai contract.
type DelegatableDaiTransferPreSignedIterator struct {
	Event *DelegatableDaiTransferPreSigned // Event containing the contract specifics and raw log

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
func (it *DelegatableDaiTransferPreSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatableDaiTransferPreSigned)
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
		it.Event = new(DelegatableDaiTransferPreSigned)
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
func (it *DelegatableDaiTransferPreSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatableDaiTransferPreSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatableDaiTransferPreSigned represents a TransferPreSigned event raised by the DelegatableDai contract.
type DelegatableDaiTransferPreSigned struct {
	From     common.Address
	To       common.Address
	Delegate common.Address
	Amount   *big.Int
	Fee      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferPreSigned is a free log retrieval operation binding the contract event 0xec5a73fd1f178be20c1bca1b406cbf4b5c20d833b66e582fc122fb4baa0fc2a4.
//
// Solidity: e TransferPreSigned(from indexed address, to indexed address, delegate indexed address, amount uint256, fee uint256)
func (_DelegatableDai *DelegatableDaiFilterer) FilterTransferPreSigned(opts *bind.FilterOpts, from []common.Address, to []common.Address, delegate []common.Address) (*DelegatableDaiTransferPreSignedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _DelegatableDai.contract.FilterLogs(opts, "TransferPreSigned", fromRule, toRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return &DelegatableDaiTransferPreSignedIterator{contract: _DelegatableDai.contract, event: "TransferPreSigned", logs: logs, sub: sub}, nil
}

// WatchTransferPreSigned is a free log subscription operation binding the contract event 0xec5a73fd1f178be20c1bca1b406cbf4b5c20d833b66e582fc122fb4baa0fc2a4.
//
// Solidity: e TransferPreSigned(from indexed address, to indexed address, delegate indexed address, amount uint256, fee uint256)
func (_DelegatableDai *DelegatableDaiFilterer) WatchTransferPreSigned(opts *bind.WatchOpts, sink chan<- *DelegatableDaiTransferPreSigned, from []common.Address, to []common.Address, delegate []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _DelegatableDai.contract.WatchLogs(opts, "TransferPreSigned", fromRule, toRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatableDaiTransferPreSigned)
				if err := _DelegatableDai.contract.UnpackLog(event, "TransferPreSigned", log); err != nil {
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
