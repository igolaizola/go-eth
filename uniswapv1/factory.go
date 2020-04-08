// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv1

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

// FactoryABI is the input ABI used to generate the binding from.
const FactoryABI = "[{\"name\":\"NewExchange\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"indexed\":true},{\"type\":\"address\",\"name\":\"exchange\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"initializeFactory\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"template\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35725},{\"name\":\"createExchange\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":187911},{\"name\":\"getExchange\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":715},{\"name\":\"getToken\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"exchange\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":745},{\"name\":\"getTokenWithId\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"token_id\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":736},{\"name\":\"exchangeTemplate\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":633},{\"name\":\"tokenCount\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":663}]"

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// ExchangeTemplate is a free data retrieval call binding the contract method 0x1c2bbd18.
//
// Solidity: function exchangeTemplate() constant returns(address out)
func (_Factory *FactoryCaller) ExchangeTemplate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "exchangeTemplate")
	return *ret0, err
}

// ExchangeTemplate is a free data retrieval call binding the contract method 0x1c2bbd18.
//
// Solidity: function exchangeTemplate() constant returns(address out)
func (_Factory *FactorySession) ExchangeTemplate() (common.Address, error) {
	return _Factory.Contract.ExchangeTemplate(&_Factory.CallOpts)
}

// ExchangeTemplate is a free data retrieval call binding the contract method 0x1c2bbd18.
//
// Solidity: function exchangeTemplate() constant returns(address out)
func (_Factory *FactoryCallerSession) ExchangeTemplate() (common.Address, error) {
	return _Factory.Contract.ExchangeTemplate(&_Factory.CallOpts)
}

// GetExchange is a free data retrieval call binding the contract method 0x06f2bf62.
//
// Solidity: function getExchange(address token) constant returns(address out)
func (_Factory *FactoryCaller) GetExchange(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "getExchange", token)
	return *ret0, err
}

// GetExchange is a free data retrieval call binding the contract method 0x06f2bf62.
//
// Solidity: function getExchange(address token) constant returns(address out)
func (_Factory *FactorySession) GetExchange(token common.Address) (common.Address, error) {
	return _Factory.Contract.GetExchange(&_Factory.CallOpts, token)
}

// GetExchange is a free data retrieval call binding the contract method 0x06f2bf62.
//
// Solidity: function getExchange(address token) constant returns(address out)
func (_Factory *FactoryCallerSession) GetExchange(token common.Address) (common.Address, error) {
	return _Factory.Contract.GetExchange(&_Factory.CallOpts, token)
}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address exchange) constant returns(address out)
func (_Factory *FactoryCaller) GetToken(opts *bind.CallOpts, exchange common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "getToken", exchange)
	return *ret0, err
}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address exchange) constant returns(address out)
func (_Factory *FactorySession) GetToken(exchange common.Address) (common.Address, error) {
	return _Factory.Contract.GetToken(&_Factory.CallOpts, exchange)
}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address exchange) constant returns(address out)
func (_Factory *FactoryCallerSession) GetToken(exchange common.Address) (common.Address, error) {
	return _Factory.Contract.GetToken(&_Factory.CallOpts, exchange)
}

// GetTokenWithId is a free data retrieval call binding the contract method 0xaa65a6c0.
//
// Solidity: function getTokenWithId(uint256 token_id) constant returns(address out)
func (_Factory *FactoryCaller) GetTokenWithId(opts *bind.CallOpts, token_id *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "getTokenWithId", token_id)
	return *ret0, err
}

// GetTokenWithId is a free data retrieval call binding the contract method 0xaa65a6c0.
//
// Solidity: function getTokenWithId(uint256 token_id) constant returns(address out)
func (_Factory *FactorySession) GetTokenWithId(token_id *big.Int) (common.Address, error) {
	return _Factory.Contract.GetTokenWithId(&_Factory.CallOpts, token_id)
}

// GetTokenWithId is a free data retrieval call binding the contract method 0xaa65a6c0.
//
// Solidity: function getTokenWithId(uint256 token_id) constant returns(address out)
func (_Factory *FactoryCallerSession) GetTokenWithId(token_id *big.Int) (common.Address, error) {
	return _Factory.Contract.GetTokenWithId(&_Factory.CallOpts, token_id)
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() constant returns(uint256 out)
func (_Factory *FactoryCaller) TokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "tokenCount")
	return *ret0, err
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() constant returns(uint256 out)
func (_Factory *FactorySession) TokenCount() (*big.Int, error) {
	return _Factory.Contract.TokenCount(&_Factory.CallOpts)
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() constant returns(uint256 out)
func (_Factory *FactoryCallerSession) TokenCount() (*big.Int, error) {
	return _Factory.Contract.TokenCount(&_Factory.CallOpts)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x1648f38e.
//
// Solidity: function createExchange(address token) returns(address out)
func (_Factory *FactoryTransactor) CreateExchange(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "createExchange", token)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x1648f38e.
//
// Solidity: function createExchange(address token) returns(address out)
func (_Factory *FactorySession) CreateExchange(token common.Address) (*types.Transaction, error) {
	return _Factory.Contract.CreateExchange(&_Factory.TransactOpts, token)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x1648f38e.
//
// Solidity: function createExchange(address token) returns(address out)
func (_Factory *FactoryTransactorSession) CreateExchange(token common.Address) (*types.Transaction, error) {
	return _Factory.Contract.CreateExchange(&_Factory.TransactOpts, token)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0x538a3f0e.
//
// Solidity: function initializeFactory(address template) returns()
func (_Factory *FactoryTransactor) InitializeFactory(opts *bind.TransactOpts, template common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "initializeFactory", template)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0x538a3f0e.
//
// Solidity: function initializeFactory(address template) returns()
func (_Factory *FactorySession) InitializeFactory(template common.Address) (*types.Transaction, error) {
	return _Factory.Contract.InitializeFactory(&_Factory.TransactOpts, template)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0x538a3f0e.
//
// Solidity: function initializeFactory(address template) returns()
func (_Factory *FactoryTransactorSession) InitializeFactory(template common.Address) (*types.Transaction, error) {
	return _Factory.Contract.InitializeFactory(&_Factory.TransactOpts, template)
}

// FactoryNewExchangeIterator is returned from FilterNewExchange and is used to iterate over the raw logs and unpacked data for NewExchange events raised by the Factory contract.
type FactoryNewExchangeIterator struct {
	Event *FactoryNewExchange // Event containing the contract specifics and raw log

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
func (it *FactoryNewExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryNewExchange)
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
		it.Event = new(FactoryNewExchange)
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
func (it *FactoryNewExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryNewExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryNewExchange represents a NewExchange event raised by the Factory contract.
type FactoryNewExchange struct {
	Token    common.Address
	Exchange common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewExchange is a free log retrieval operation binding the contract event 0x9d42cb017eb05bd8944ab536a8b35bc68085931dd5f4356489801453923953f9.
//
// Solidity: event NewExchange(address indexed token, address indexed exchange)
func (_Factory *FactoryFilterer) FilterNewExchange(opts *bind.FilterOpts, token []common.Address, exchange []common.Address) (*FactoryNewExchangeIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var exchangeRule []interface{}
	for _, exchangeItem := range exchange {
		exchangeRule = append(exchangeRule, exchangeItem)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "NewExchange", tokenRule, exchangeRule)
	if err != nil {
		return nil, err
	}
	return &FactoryNewExchangeIterator{contract: _Factory.contract, event: "NewExchange", logs: logs, sub: sub}, nil
}

// WatchNewExchange is a free log subscription operation binding the contract event 0x9d42cb017eb05bd8944ab536a8b35bc68085931dd5f4356489801453923953f9.
//
// Solidity: event NewExchange(address indexed token, address indexed exchange)
func (_Factory *FactoryFilterer) WatchNewExchange(opts *bind.WatchOpts, sink chan<- *FactoryNewExchange, token []common.Address, exchange []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var exchangeRule []interface{}
	for _, exchangeItem := range exchange {
		exchangeRule = append(exchangeRule, exchangeItem)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "NewExchange", tokenRule, exchangeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryNewExchange)
				if err := _Factory.contract.UnpackLog(event, "NewExchange", log); err != nil {
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

// ParseNewExchange is a log parse operation binding the contract event 0x9d42cb017eb05bd8944ab536a8b35bc68085931dd5f4356489801453923953f9.
//
// Solidity: event NewExchange(address indexed token, address indexed exchange)
func (_Factory *FactoryFilterer) ParseNewExchange(log types.Log) (*FactoryNewExchange, error) {
	event := new(FactoryNewExchange)
	if err := _Factory.contract.UnpackLog(event, "NewExchange", log); err != nil {
		return nil, err
	}
	return event, nil
}
