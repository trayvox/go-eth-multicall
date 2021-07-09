// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MultiCall

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

// Multicall2Call is an auto generated low-level Go binding around an user-defined struct.
type Multicall2Call struct {
	Target   common.Address
	CallData []byte
}

// Multicall2Result is an auto generated low-level Go binding around an user-defined struct.
type Multicall2Result struct {
	Success    bool
	ReturnData []byte
}

// MultiCall2ABI is the input ABI used to generate the binding from.
const MultiCall2ABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"blockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryAggregate\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryBlockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MultiCall2 is an auto generated Go binding around an Ethereum contract.
type MultiCall2 struct {
	MultiCall2Caller     // Read-only binding to the contract
	MultiCall2Transactor // Write-only binding to the contract
	MultiCall2Filterer   // Log filterer for contract events
}

// MultiCall2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MultiCall2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCall2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiCall2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCall2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiCall2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCall2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiCall2Session struct {
	Contract     *MultiCall2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiCall2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiCall2CallerSession struct {
	Contract *MultiCall2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MultiCall2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiCall2TransactorSession struct {
	Contract     *MultiCall2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MultiCall2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MultiCall2Raw struct {
	Contract *MultiCall2 // Generic contract binding to access the raw methods on
}

// MultiCall2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiCall2CallerRaw struct {
	Contract *MultiCall2Caller // Generic read-only contract binding to access the raw methods on
}

// MultiCall2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiCall2TransactorRaw struct {
	Contract *MultiCall2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiCall2 creates a new instance of MultiCall2, bound to a specific deployed contract.
func NewMultiCall2(address common.Address, backend bind.ContractBackend) (*MultiCall2, error) {
	contract, err := bindMultiCall2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiCall2{MultiCall2Caller: MultiCall2Caller{contract: contract}, MultiCall2Transactor: MultiCall2Transactor{contract: contract}, MultiCall2Filterer: MultiCall2Filterer{contract: contract}}, nil
}

// NewMultiCall2Caller creates a new read-only instance of MultiCall2, bound to a specific deployed contract.
func NewMultiCall2Caller(address common.Address, caller bind.ContractCaller) (*MultiCall2Caller, error) {
	contract, err := bindMultiCall2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiCall2Caller{contract: contract}, nil
}

// NewMultiCall2Transactor creates a new write-only instance of MultiCall2, bound to a specific deployed contract.
func NewMultiCall2Transactor(address common.Address, transactor bind.ContractTransactor) (*MultiCall2Transactor, error) {
	contract, err := bindMultiCall2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiCall2Transactor{contract: contract}, nil
}

// NewMultiCall2Filterer creates a new log filterer instance of MultiCall2, bound to a specific deployed contract.
func NewMultiCall2Filterer(address common.Address, filterer bind.ContractFilterer) (*MultiCall2Filterer, error) {
	contract, err := bindMultiCall2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiCall2Filterer{contract: contract}, nil
}

// bindMultiCall2 binds a generic wrapper to an already deployed contract.
func bindMultiCall2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiCall2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiCall2 *MultiCall2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiCall2.Contract.MultiCall2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiCall2 *MultiCall2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiCall2.Contract.MultiCall2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiCall2 *MultiCall2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiCall2.Contract.MultiCall2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiCall2 *MultiCall2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiCall2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiCall2 *MultiCall2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiCall2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiCall2 *MultiCall2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiCall2.Contract.contract.Transact(opts, method, params...)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MultiCall2 *MultiCall2Caller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MultiCall2.contract.Call(opts, &out, "getBlockHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MultiCall2 *MultiCall2Session) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _MultiCall2.Contract.GetBlockHash(&_MultiCall2.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MultiCall2 *MultiCall2CallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _MultiCall2.Contract.GetBlockHash(&_MultiCall2.CallOpts, blockNumber)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_MultiCall2 *MultiCall2Caller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiCall2.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_MultiCall2 *MultiCall2Session) GetBlockNumber() (*big.Int, error) {
	return _MultiCall2.Contract.GetBlockNumber(&_MultiCall2.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_MultiCall2 *MultiCall2CallerSession) GetBlockNumber() (*big.Int, error) {
	return _MultiCall2.Contract.GetBlockNumber(&_MultiCall2.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MultiCall2 *MultiCall2Caller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultiCall2.contract.Call(opts, &out, "getCurrentBlockCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MultiCall2 *MultiCall2Session) GetCurrentBlockCoinbase() (common.Address, error) {
	return _MultiCall2.Contract.GetCurrentBlockCoinbase(&_MultiCall2.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MultiCall2 *MultiCall2CallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _MultiCall2.Contract.GetCurrentBlockCoinbase(&_MultiCall2.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MultiCall2 *MultiCall2Caller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiCall2.contract.Call(opts, &out, "getCurrentBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MultiCall2 *MultiCall2Session) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _MultiCall2.Contract.GetCurrentBlockDifficulty(&_MultiCall2.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MultiCall2 *MultiCall2CallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _MultiCall2.Contract.GetCurrentBlockDifficulty(&_MultiCall2.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MultiCall2 *MultiCall2Caller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiCall2.contract.Call(opts, &out, "getCurrentBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MultiCall2 *MultiCall2Session) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _MultiCall2.Contract.GetCurrentBlockGasLimit(&_MultiCall2.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MultiCall2 *MultiCall2CallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _MultiCall2.Contract.GetCurrentBlockGasLimit(&_MultiCall2.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MultiCall2 *MultiCall2Caller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiCall2.contract.Call(opts, &out, "getCurrentBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MultiCall2 *MultiCall2Session) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _MultiCall2.Contract.GetCurrentBlockTimestamp(&_MultiCall2.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MultiCall2 *MultiCall2CallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _MultiCall2.Contract.GetCurrentBlockTimestamp(&_MultiCall2.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MultiCall2 *MultiCall2Caller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MultiCall2.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MultiCall2 *MultiCall2Session) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MultiCall2.Contract.GetEthBalance(&_MultiCall2.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MultiCall2 *MultiCall2CallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MultiCall2.Contract.GetEthBalance(&_MultiCall2.CallOpts, addr)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MultiCall2 *MultiCall2Caller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MultiCall2.contract.Call(opts, &out, "getLastBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MultiCall2 *MultiCall2Session) GetLastBlockHash() ([32]byte, error) {
	return _MultiCall2.Contract.GetLastBlockHash(&_MultiCall2.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MultiCall2 *MultiCall2CallerSession) GetLastBlockHash() ([32]byte, error) {
	return _MultiCall2.Contract.GetLastBlockHash(&_MultiCall2.CallOpts)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MultiCall2 *MultiCall2Transactor) Aggregate(opts *bind.TransactOpts, calls []Multicall2Call) (*types.Transaction, error) {
	return _MultiCall2.contract.Transact(opts, "aggregate", calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MultiCall2 *MultiCall2Session) Aggregate(calls []Multicall2Call) (*types.Transaction, error) {
	return _MultiCall2.Contract.Aggregate(&_MultiCall2.TransactOpts, calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MultiCall2 *MultiCall2TransactorSession) Aggregate(calls []Multicall2Call) (*types.Transaction, error) {
	return _MultiCall2.Contract.Aggregate(&_MultiCall2.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MultiCall2 *MultiCall2Transactor) BlockAndAggregate(opts *bind.TransactOpts, calls []Multicall2Call) (*types.Transaction, error) {
	return _MultiCall2.contract.Transact(opts, "blockAndAggregate", calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MultiCall2 *MultiCall2Session) BlockAndAggregate(calls []Multicall2Call) (*types.Transaction, error) {
	return _MultiCall2.Contract.BlockAndAggregate(&_MultiCall2.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MultiCall2 *MultiCall2TransactorSession) BlockAndAggregate(calls []Multicall2Call) (*types.Transaction, error) {
	return _MultiCall2.Contract.BlockAndAggregate(&_MultiCall2.TransactOpts, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) returns((bool,bytes)[] returnData)
func (_MultiCall2 *MultiCall2Transactor) TryAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _MultiCall2.contract.Transact(opts, "tryAggregate", requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) returns((bool,bytes)[] returnData)
func (_MultiCall2 *MultiCall2Session) TryAggregate(requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _MultiCall2.Contract.TryAggregate(&_MultiCall2.TransactOpts, requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) returns((bool,bytes)[] returnData)
func (_MultiCall2 *MultiCall2TransactorSession) TryAggregate(requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _MultiCall2.Contract.TryAggregate(&_MultiCall2.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MultiCall2 *MultiCall2Transactor) TryBlockAndAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _MultiCall2.contract.Transact(opts, "tryBlockAndAggregate", requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MultiCall2 *MultiCall2Session) TryBlockAndAggregate(requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _MultiCall2.Contract.TryBlockAndAggregate(&_MultiCall2.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MultiCall2 *MultiCall2TransactorSession) TryBlockAndAggregate(requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _MultiCall2.Contract.TryBlockAndAggregate(&_MultiCall2.TransactOpts, requireSuccess, calls)
}
