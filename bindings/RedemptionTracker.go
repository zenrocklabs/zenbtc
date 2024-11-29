// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RedemptionTracker

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

// RedemptionTrackerMetaData contains all meta data concerning the RedemptionTracker contract.
var RedemptionTrackerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"numBlocks\",\"type\":\"uint8\"}],\"name\":\"getRecentRedemptions\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"destinationAddrs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64[]\",\"name\":\"amounts\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"ids\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RedemptionTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use RedemptionTrackerMetaData.ABI instead.
var RedemptionTrackerABI = RedemptionTrackerMetaData.ABI

// RedemptionTracker is an auto generated Go binding around an Ethereum contract.
type RedemptionTracker struct {
	RedemptionTrackerCaller     // Read-only binding to the contract
	RedemptionTrackerTransactor // Write-only binding to the contract
	RedemptionTrackerFilterer   // Log filterer for contract events
}

// RedemptionTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RedemptionTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedemptionTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RedemptionTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedemptionTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RedemptionTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedemptionTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RedemptionTrackerSession struct {
	Contract     *RedemptionTracker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RedemptionTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RedemptionTrackerCallerSession struct {
	Contract *RedemptionTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RedemptionTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RedemptionTrackerTransactorSession struct {
	Contract     *RedemptionTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RedemptionTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RedemptionTrackerRaw struct {
	Contract *RedemptionTracker // Generic contract binding to access the raw methods on
}

// RedemptionTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RedemptionTrackerCallerRaw struct {
	Contract *RedemptionTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// RedemptionTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RedemptionTrackerTransactorRaw struct {
	Contract *RedemptionTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRedemptionTracker creates a new instance of RedemptionTracker, bound to a specific deployed contract.
func NewRedemptionTracker(address common.Address, backend bind.ContractBackend) (*RedemptionTracker, error) {
	contract, err := bindRedemptionTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RedemptionTracker{RedemptionTrackerCaller: RedemptionTrackerCaller{contract: contract}, RedemptionTrackerTransactor: RedemptionTrackerTransactor{contract: contract}, RedemptionTrackerFilterer: RedemptionTrackerFilterer{contract: contract}}, nil
}

// NewRedemptionTrackerCaller creates a new read-only instance of RedemptionTracker, bound to a specific deployed contract.
func NewRedemptionTrackerCaller(address common.Address, caller bind.ContractCaller) (*RedemptionTrackerCaller, error) {
	contract, err := bindRedemptionTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RedemptionTrackerCaller{contract: contract}, nil
}

// NewRedemptionTrackerTransactor creates a new write-only instance of RedemptionTracker, bound to a specific deployed contract.
func NewRedemptionTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*RedemptionTrackerTransactor, error) {
	contract, err := bindRedemptionTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RedemptionTrackerTransactor{contract: contract}, nil
}

// NewRedemptionTrackerFilterer creates a new log filterer instance of RedemptionTracker, bound to a specific deployed contract.
func NewRedemptionTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*RedemptionTrackerFilterer, error) {
	contract, err := bindRedemptionTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RedemptionTrackerFilterer{contract: contract}, nil
}

// bindRedemptionTracker binds a generic wrapper to an already deployed contract.
func bindRedemptionTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RedemptionTrackerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RedemptionTracker *RedemptionTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RedemptionTracker.Contract.RedemptionTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RedemptionTracker *RedemptionTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedemptionTracker.Contract.RedemptionTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RedemptionTracker *RedemptionTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RedemptionTracker.Contract.RedemptionTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RedemptionTracker *RedemptionTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RedemptionTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RedemptionTracker *RedemptionTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedemptionTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RedemptionTracker *RedemptionTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RedemptionTracker.Contract.contract.Transact(opts, method, params...)
}

// GetRecentRedemptions is a free data retrieval call binding the contract method 0xe7332f55.
//
// Solidity: function getRecentRedemptions(uint8 numBlocks) view returns(bytes[] destinationAddrs, uint64[] amounts, uint64[] ids)
func (_RedemptionTracker *RedemptionTrackerCaller) GetRecentRedemptions(opts *bind.CallOpts, numBlocks uint8) (struct {
	DestinationAddrs [][]byte
	Amounts          []uint64
	Ids              []uint64
}, error) {
	var out []interface{}
	err := _RedemptionTracker.contract.Call(opts, &out, "getRecentRedemptions", numBlocks)

	outstruct := new(struct {
		DestinationAddrs [][]byte
		Amounts          []uint64
		Ids              []uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DestinationAddrs = *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)
	outstruct.Ids = *abi.ConvertType(out[2], new([]uint64)).(*[]uint64)

	return *outstruct, err

}

// GetRecentRedemptions is a free data retrieval call binding the contract method 0xe7332f55.
//
// Solidity: function getRecentRedemptions(uint8 numBlocks) view returns(bytes[] destinationAddrs, uint64[] amounts, uint64[] ids)
func (_RedemptionTracker *RedemptionTrackerSession) GetRecentRedemptions(numBlocks uint8) (struct {
	DestinationAddrs [][]byte
	Amounts          []uint64
	Ids              []uint64
}, error) {
	return _RedemptionTracker.Contract.GetRecentRedemptions(&_RedemptionTracker.CallOpts, numBlocks)
}

// GetRecentRedemptions is a free data retrieval call binding the contract method 0xe7332f55.
//
// Solidity: function getRecentRedemptions(uint8 numBlocks) view returns(bytes[] destinationAddrs, uint64[] amounts, uint64[] ids)
func (_RedemptionTracker *RedemptionTrackerCallerSession) GetRecentRedemptions(numBlocks uint8) (struct {
	DestinationAddrs [][]byte
	Amounts          []uint64
	Ids              []uint64
}, error) {
	return _RedemptionTracker.Contract.GetRecentRedemptions(&_RedemptionTracker.CallOpts, numBlocks)
}

// RedemptionTrackerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RedemptionTracker contract.
type RedemptionTrackerInitializedIterator struct {
	Event *RedemptionTrackerInitialized // Event containing the contract specifics and raw log

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
func (it *RedemptionTrackerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedemptionTrackerInitialized)
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
		it.Event = new(RedemptionTrackerInitialized)
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
func (it *RedemptionTrackerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedemptionTrackerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedemptionTrackerInitialized represents a Initialized event raised by the RedemptionTracker contract.
type RedemptionTrackerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_RedemptionTracker *RedemptionTrackerFilterer) FilterInitialized(opts *bind.FilterOpts) (*RedemptionTrackerInitializedIterator, error) {

	logs, sub, err := _RedemptionTracker.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RedemptionTrackerInitializedIterator{contract: _RedemptionTracker.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_RedemptionTracker *RedemptionTrackerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RedemptionTrackerInitialized) (event.Subscription, error) {

	logs, sub, err := _RedemptionTracker.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedemptionTrackerInitialized)
				if err := _RedemptionTracker.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_RedemptionTracker *RedemptionTrackerFilterer) ParseInitialized(log types.Log) (*RedemptionTrackerInitialized, error) {
	event := new(RedemptionTrackerInitialized)
	if err := _RedemptionTracker.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
