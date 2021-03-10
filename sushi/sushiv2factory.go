// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sushi

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SushiV2FactoryABI is the input ABI used to generate the binding from.
const SushiV2FactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairCodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SushiV2Factory is an auto generated Go binding around an Ethereum contract.
type SushiV2Factory struct {
	SushiV2FactoryCaller     // Read-only binding to the contract
	SushiV2FactoryTransactor // Write-only binding to the contract
	SushiV2FactoryFilterer   // Log filterer for contract events
}

// SushiV2FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SushiV2FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiV2FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SushiV2FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiV2FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SushiV2FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiV2FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SushiV2FactorySession struct {
	Contract     *SushiV2Factory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SushiV2FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SushiV2FactoryCallerSession struct {
	Contract *SushiV2FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SushiV2FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SushiV2FactoryTransactorSession struct {
	Contract     *SushiV2FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SushiV2FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SushiV2FactoryRaw struct {
	Contract *SushiV2Factory // Generic contract binding to access the raw methods on
}

// SushiV2FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SushiV2FactoryCallerRaw struct {
	Contract *SushiV2FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SushiV2FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SushiV2FactoryTransactorRaw struct {
	Contract *SushiV2FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSushiV2Factory creates a new instance of SushiV2Factory, bound to a specific deployed contract.
func NewSushiV2Factory(address common.Address, backend bind.ContractBackend) (*SushiV2Factory, error) {
	contract, err := bindSushiV2Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SushiV2Factory{SushiV2FactoryCaller: SushiV2FactoryCaller{contract: contract}, SushiV2FactoryTransactor: SushiV2FactoryTransactor{contract: contract}, SushiV2FactoryFilterer: SushiV2FactoryFilterer{contract: contract}}, nil
}

// NewSushiV2FactoryCaller creates a new read-only instance of SushiV2Factory, bound to a specific deployed contract.
func NewSushiV2FactoryCaller(address common.Address, caller bind.ContractCaller) (*SushiV2FactoryCaller, error) {
	contract, err := bindSushiV2Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SushiV2FactoryCaller{contract: contract}, nil
}

// NewSushiV2FactoryTransactor creates a new write-only instance of SushiV2Factory, bound to a specific deployed contract.
func NewSushiV2FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SushiV2FactoryTransactor, error) {
	contract, err := bindSushiV2Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SushiV2FactoryTransactor{contract: contract}, nil
}

// NewSushiV2FactoryFilterer creates a new log filterer instance of SushiV2Factory, bound to a specific deployed contract.
func NewSushiV2FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SushiV2FactoryFilterer, error) {
	contract, err := bindSushiV2Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SushiV2FactoryFilterer{contract: contract}, nil
}

// bindSushiV2Factory binds a generic wrapper to an already deployed contract.
func bindSushiV2Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SushiV2FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SushiV2Factory *SushiV2FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SushiV2Factory.Contract.SushiV2FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SushiV2Factory *SushiV2FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SushiV2Factory.Contract.SushiV2FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SushiV2Factory *SushiV2FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SushiV2Factory.Contract.SushiV2FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SushiV2Factory *SushiV2FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SushiV2Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SushiV2Factory *SushiV2FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SushiV2Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SushiV2Factory *SushiV2FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SushiV2Factory.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_SushiV2Factory *SushiV2FactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SushiV2Factory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_SushiV2Factory *SushiV2FactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _SushiV2Factory.Contract.AllPairs(&_SushiV2Factory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_SushiV2Factory *SushiV2FactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _SushiV2Factory.Contract.AllPairs(&_SushiV2Factory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_SushiV2Factory *SushiV2FactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SushiV2Factory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_SushiV2Factory *SushiV2FactorySession) AllPairsLength() (*big.Int, error) {
	return _SushiV2Factory.Contract.AllPairsLength(&_SushiV2Factory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_SushiV2Factory *SushiV2FactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _SushiV2Factory.Contract.AllPairsLength(&_SushiV2Factory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_SushiV2Factory *SushiV2FactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SushiV2Factory.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_SushiV2Factory *SushiV2FactorySession) FeeTo() (common.Address, error) {
	return _SushiV2Factory.Contract.FeeTo(&_SushiV2Factory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_SushiV2Factory *SushiV2FactoryCallerSession) FeeTo() (common.Address, error) {
	return _SushiV2Factory.Contract.FeeTo(&_SushiV2Factory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_SushiV2Factory *SushiV2FactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SushiV2Factory.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_SushiV2Factory *SushiV2FactorySession) FeeToSetter() (common.Address, error) {
	return _SushiV2Factory.Contract.FeeToSetter(&_SushiV2Factory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_SushiV2Factory *SushiV2FactoryCallerSession) FeeToSetter() (common.Address, error) {
	return _SushiV2Factory.Contract.FeeToSetter(&_SushiV2Factory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_SushiV2Factory *SushiV2FactoryCaller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _SushiV2Factory.contract.Call(opts, &out, "getPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_SushiV2Factory *SushiV2FactorySession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _SushiV2Factory.Contract.GetPair(&_SushiV2Factory.CallOpts, arg0, arg1)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_SushiV2Factory *SushiV2FactoryCallerSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _SushiV2Factory.Contract.GetPair(&_SushiV2Factory.CallOpts, arg0, arg1)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_SushiV2Factory *SushiV2FactoryCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SushiV2Factory.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_SushiV2Factory *SushiV2FactorySession) Migrator() (common.Address, error) {
	return _SushiV2Factory.Contract.Migrator(&_SushiV2Factory.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_SushiV2Factory *SushiV2FactoryCallerSession) Migrator() (common.Address, error) {
	return _SushiV2Factory.Contract.Migrator(&_SushiV2Factory.CallOpts)
}

// PairCodeHash is a free data retrieval call binding the contract method 0x9aab9248.
//
// Solidity: function pairCodeHash() pure returns(bytes32)
func (_SushiV2Factory *SushiV2FactoryCaller) PairCodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SushiV2Factory.contract.Call(opts, &out, "pairCodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PairCodeHash is a free data retrieval call binding the contract method 0x9aab9248.
//
// Solidity: function pairCodeHash() pure returns(bytes32)
func (_SushiV2Factory *SushiV2FactorySession) PairCodeHash() ([32]byte, error) {
	return _SushiV2Factory.Contract.PairCodeHash(&_SushiV2Factory.CallOpts)
}

// PairCodeHash is a free data retrieval call binding the contract method 0x9aab9248.
//
// Solidity: function pairCodeHash() pure returns(bytes32)
func (_SushiV2Factory *SushiV2FactoryCallerSession) PairCodeHash() ([32]byte, error) {
	return _SushiV2Factory.Contract.PairCodeHash(&_SushiV2Factory.CallOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_SushiV2Factory *SushiV2FactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _SushiV2Factory.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_SushiV2Factory *SushiV2FactorySession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _SushiV2Factory.Contract.CreatePair(&_SushiV2Factory.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_SushiV2Factory *SushiV2FactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _SushiV2Factory.Contract.CreatePair(&_SushiV2Factory.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_SushiV2Factory *SushiV2FactoryTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _SushiV2Factory.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_SushiV2Factory *SushiV2FactorySession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _SushiV2Factory.Contract.SetFeeTo(&_SushiV2Factory.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_SushiV2Factory *SushiV2FactoryTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _SushiV2Factory.Contract.SetFeeTo(&_SushiV2Factory.TransactOpts, _feeTo)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_SushiV2Factory *SushiV2FactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _SushiV2Factory.contract.Transact(opts, "setFeeToSetter", _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_SushiV2Factory *SushiV2FactorySession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _SushiV2Factory.Contract.SetFeeToSetter(&_SushiV2Factory.TransactOpts, _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_SushiV2Factory *SushiV2FactoryTransactorSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _SushiV2Factory.Contract.SetFeeToSetter(&_SushiV2Factory.TransactOpts, _feeToSetter)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_SushiV2Factory *SushiV2FactoryTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _SushiV2Factory.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_SushiV2Factory *SushiV2FactorySession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _SushiV2Factory.Contract.SetMigrator(&_SushiV2Factory.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_SushiV2Factory *SushiV2FactoryTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _SushiV2Factory.Contract.SetMigrator(&_SushiV2Factory.TransactOpts, _migrator)
}

// SushiV2FactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the SushiV2Factory contract.
type SushiV2FactoryPairCreatedIterator struct {
	Event *SushiV2FactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *SushiV2FactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushiV2FactoryPairCreated)
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
		it.Event = new(SushiV2FactoryPairCreated)
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
func (it *SushiV2FactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushiV2FactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushiV2FactoryPairCreated represents a PairCreated event raised by the SushiV2Factory contract.
type SushiV2FactoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_SushiV2Factory *SushiV2FactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*SushiV2FactoryPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _SushiV2Factory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &SushiV2FactoryPairCreatedIterator{contract: _SushiV2Factory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_SushiV2Factory *SushiV2FactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *SushiV2FactoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _SushiV2Factory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushiV2FactoryPairCreated)
				if err := _SushiV2Factory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_SushiV2Factory *SushiV2FactoryFilterer) ParsePairCreated(log types.Log) (*SushiV2FactoryPairCreated, error) {
	event := new(SushiV2FactoryPairCreated)
	if err := _SushiV2Factory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
