// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package helloworld

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

// IHelloWorldServiceManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IHelloWorldServiceManagerTask struct {
	Name             string
	TaskCreatedBlock uint32
}

// IPaymentCoordinatorRangePayment is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorRangePayment struct {
	StrategiesAndMultipliers []IPaymentCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IPaymentCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// HelloWorldMetaData contains all meta data concerning the HelloWorld contract.
var HelloWorldMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorHasMinimumWeight\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIHelloWorldServiceManager.Task\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIHelloWorldServiceManager.Task\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIHelloWorldServiceManager.Task\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// HelloWorldABI is the input ABI used to generate the binding from.
// Deprecated: Use HelloWorldMetaData.ABI instead.
var HelloWorldABI = HelloWorldMetaData.ABI

// HelloWorld is an auto generated Go binding around an Ethereum contract.
type HelloWorld struct {
	HelloWorldCaller     // Read-only binding to the contract
	HelloWorldTransactor // Write-only binding to the contract
	HelloWorldFilterer   // Log filterer for contract events
}

// HelloWorldCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelloWorldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloWorldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloWorldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloWorldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloWorldSession struct {
	Contract     *HelloWorld             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloWorldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloWorldCallerSession struct {
	Contract *HelloWorldCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HelloWorldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloWorldTransactorSession struct {
	Contract     *HelloWorldTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloWorldRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelloWorldRaw struct {
	Contract *HelloWorld // Generic contract binding to access the raw methods on
}

// HelloWorldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloWorldCallerRaw struct {
	Contract *HelloWorldCaller // Generic read-only contract binding to access the raw methods on
}

// HelloWorldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloWorldTransactorRaw struct {
	Contract *HelloWorldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHelloWorld creates a new instance of HelloWorld, bound to a specific deployed contract.
func NewHelloWorld(address common.Address, backend bind.ContractBackend) (*HelloWorld, error) {
	contract, err := bindHelloWorld(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HelloWorld{HelloWorldCaller: HelloWorldCaller{contract: contract}, HelloWorldTransactor: HelloWorldTransactor{contract: contract}, HelloWorldFilterer: HelloWorldFilterer{contract: contract}}, nil
}

// NewHelloWorldCaller creates a new read-only instance of HelloWorld, bound to a specific deployed contract.
func NewHelloWorldCaller(address common.Address, caller bind.ContractCaller) (*HelloWorldCaller, error) {
	contract, err := bindHelloWorld(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloWorldCaller{contract: contract}, nil
}

// NewHelloWorldTransactor creates a new write-only instance of HelloWorld, bound to a specific deployed contract.
func NewHelloWorldTransactor(address common.Address, transactor bind.ContractTransactor) (*HelloWorldTransactor, error) {
	contract, err := bindHelloWorld(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloWorldTransactor{contract: contract}, nil
}

// NewHelloWorldFilterer creates a new log filterer instance of HelloWorld, bound to a specific deployed contract.
func NewHelloWorldFilterer(address common.Address, filterer bind.ContractFilterer) (*HelloWorldFilterer, error) {
	contract, err := bindHelloWorld(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloWorldFilterer{contract: contract}, nil
}

// bindHelloWorld binds a generic wrapper to an already deployed contract.
func bindHelloWorld(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HelloWorldMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HelloWorld *HelloWorldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HelloWorld.Contract.HelloWorldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HelloWorld *HelloWorldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HelloWorld.Contract.HelloWorldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HelloWorld *HelloWorldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HelloWorld.Contract.HelloWorldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HelloWorld *HelloWorldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HelloWorld.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HelloWorld *HelloWorldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HelloWorld.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HelloWorld *HelloWorldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HelloWorld.Contract.contract.Transact(opts, method, params...)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_HelloWorld *HelloWorldCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _HelloWorld.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_HelloWorld *HelloWorldSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _HelloWorld.Contract.AllTaskHashes(&_HelloWorld.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_HelloWorld *HelloWorldCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _HelloWorld.Contract.AllTaskHashes(&_HelloWorld.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0xc20bab7f.
//
// Solidity: function allTaskResponses(address , uint32 ) view returns(bytes)
func (_HelloWorld *HelloWorldCaller) AllTaskResponses(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) ([]byte, error) {
	var out []interface{}
	err := _HelloWorld.contract.Call(opts, &out, "allTaskResponses", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0xc20bab7f.
//
// Solidity: function allTaskResponses(address , uint32 ) view returns(bytes)
func (_HelloWorld *HelloWorldSession) AllTaskResponses(arg0 common.Address, arg1 uint32) ([]byte, error) {
	return _HelloWorld.Contract.AllTaskResponses(&_HelloWorld.CallOpts, arg0, arg1)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0xc20bab7f.
//
// Solidity: function allTaskResponses(address , uint32 ) view returns(bytes)
func (_HelloWorld *HelloWorldCallerSession) AllTaskResponses(arg0 common.Address, arg1 uint32) ([]byte, error) {
	return _HelloWorld.Contract.AllTaskResponses(&_HelloWorld.CallOpts, arg0, arg1)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_HelloWorld *HelloWorldCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HelloWorld.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_HelloWorld *HelloWorldSession) AvsDirectory() (common.Address, error) {
	return _HelloWorld.Contract.AvsDirectory(&_HelloWorld.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_HelloWorld *HelloWorldCallerSession) AvsDirectory() (common.Address, error) {
	return _HelloWorld.Contract.AvsDirectory(&_HelloWorld.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_HelloWorld *HelloWorldCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, _operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _HelloWorld.contract.Call(opts, &out, "getOperatorRestakedStrategies", _operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_HelloWorld *HelloWorldSession) GetOperatorRestakedStrategies(_operator common.Address) ([]common.Address, error) {
	return _HelloWorld.Contract.GetOperatorRestakedStrategies(&_HelloWorld.CallOpts, _operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_HelloWorld *HelloWorldCallerSession) GetOperatorRestakedStrategies(_operator common.Address) ([]common.Address, error) {
	return _HelloWorld.Contract.GetOperatorRestakedStrategies(&_HelloWorld.CallOpts, _operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_HelloWorld *HelloWorldCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _HelloWorld.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_HelloWorld *HelloWorldSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _HelloWorld.Contract.GetRestakeableStrategies(&_HelloWorld.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_HelloWorld *HelloWorldCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _HelloWorld.Contract.GetRestakeableStrategies(&_HelloWorld.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_HelloWorld *HelloWorldCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _HelloWorld.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_HelloWorld *HelloWorldSession) LatestTaskNum() (uint32, error) {
	return _HelloWorld.Contract.LatestTaskNum(&_HelloWorld.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_HelloWorld *HelloWorldCallerSession) LatestTaskNum() (uint32, error) {
	return _HelloWorld.Contract.LatestTaskNum(&_HelloWorld.CallOpts)
}

// OperatorHasMinimumWeight is a free data retrieval call binding the contract method 0xd5dba76d.
//
// Solidity: function operatorHasMinimumWeight(address operator) view returns(bool)
func (_HelloWorld *HelloWorldCaller) OperatorHasMinimumWeight(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _HelloWorld.contract.Call(opts, &out, "operatorHasMinimumWeight", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorHasMinimumWeight is a free data retrieval call binding the contract method 0xd5dba76d.
//
// Solidity: function operatorHasMinimumWeight(address operator) view returns(bool)
func (_HelloWorld *HelloWorldSession) OperatorHasMinimumWeight(operator common.Address) (bool, error) {
	return _HelloWorld.Contract.OperatorHasMinimumWeight(&_HelloWorld.CallOpts, operator)
}

// OperatorHasMinimumWeight is a free data retrieval call binding the contract method 0xd5dba76d.
//
// Solidity: function operatorHasMinimumWeight(address operator) view returns(bool)
func (_HelloWorld *HelloWorldCallerSession) OperatorHasMinimumWeight(operator common.Address) (bool, error) {
	return _HelloWorld.Contract.OperatorHasMinimumWeight(&_HelloWorld.CallOpts, operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HelloWorld *HelloWorldCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HelloWorld.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HelloWorld *HelloWorldSession) Owner() (common.Address, error) {
	return _HelloWorld.Contract.Owner(&_HelloWorld.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HelloWorld *HelloWorldCallerSession) Owner() (common.Address, error) {
	return _HelloWorld.Contract.Owner(&_HelloWorld.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_HelloWorld *HelloWorldCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _HelloWorld.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_HelloWorld *HelloWorldSession) Paused(index uint8) (bool, error) {
	return _HelloWorld.Contract.Paused(&_HelloWorld.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_HelloWorld *HelloWorldCallerSession) Paused(index uint8) (bool, error) {
	return _HelloWorld.Contract.Paused(&_HelloWorld.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_HelloWorld *HelloWorldCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HelloWorld.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_HelloWorld *HelloWorldSession) Paused0() (*big.Int, error) {
	return _HelloWorld.Contract.Paused0(&_HelloWorld.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_HelloWorld *HelloWorldCallerSession) Paused0() (*big.Int, error) {
	return _HelloWorld.Contract.Paused0(&_HelloWorld.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_HelloWorld *HelloWorldCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HelloWorld.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_HelloWorld *HelloWorldSession) PauserRegistry() (common.Address, error) {
	return _HelloWorld.Contract.PauserRegistry(&_HelloWorld.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_HelloWorld *HelloWorldCallerSession) PauserRegistry() (common.Address, error) {
	return _HelloWorld.Contract.PauserRegistry(&_HelloWorld.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_HelloWorld *HelloWorldCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HelloWorld.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_HelloWorld *HelloWorldSession) StakeRegistry() (common.Address, error) {
	return _HelloWorld.Contract.StakeRegistry(&_HelloWorld.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_HelloWorld *HelloWorldCallerSession) StakeRegistry() (common.Address, error) {
	return _HelloWorld.Contract.StakeRegistry(&_HelloWorld.CallOpts)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x85edf874.
//
// Solidity: function createNewTask(string name) returns()
func (_HelloWorld *HelloWorldTransactor) CreateNewTask(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _HelloWorld.contract.Transact(opts, "createNewTask", name)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x85edf874.
//
// Solidity: function createNewTask(string name) returns()
func (_HelloWorld *HelloWorldSession) CreateNewTask(name string) (*types.Transaction, error) {
	return _HelloWorld.Contract.CreateNewTask(&_HelloWorld.TransactOpts, name)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x85edf874.
//
// Solidity: function createNewTask(string name) returns()
func (_HelloWorld *HelloWorldTransactorSession) CreateNewTask(name string) (*types.Transaction, error) {
	return _HelloWorld.Contract.CreateNewTask(&_HelloWorld.TransactOpts, name)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_HelloWorld *HelloWorldTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _HelloWorld.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_HelloWorld *HelloWorldSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _HelloWorld.Contract.DeregisterOperatorFromAVS(&_HelloWorld.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_HelloWorld *HelloWorldTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _HelloWorld.Contract.DeregisterOperatorFromAVS(&_HelloWorld.TransactOpts, operator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_HelloWorld *HelloWorldTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _HelloWorld.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_HelloWorld *HelloWorldSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _HelloWorld.Contract.Pause(&_HelloWorld.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_HelloWorld *HelloWorldTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _HelloWorld.Contract.Pause(&_HelloWorld.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_HelloWorld *HelloWorldTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HelloWorld.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_HelloWorld *HelloWorldSession) PauseAll() (*types.Transaction, error) {
	return _HelloWorld.Contract.PauseAll(&_HelloWorld.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_HelloWorld *HelloWorldTransactorSession) PauseAll() (*types.Transaction, error) {
	return _HelloWorld.Contract.PauseAll(&_HelloWorld.TransactOpts)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_HelloWorld *HelloWorldTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _HelloWorld.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_HelloWorld *HelloWorldSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _HelloWorld.Contract.PayForRange(&_HelloWorld.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_HelloWorld *HelloWorldTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _HelloWorld.Contract.PayForRange(&_HelloWorld.TransactOpts, rangePayments)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_HelloWorld *HelloWorldTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _HelloWorld.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_HelloWorld *HelloWorldSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _HelloWorld.Contract.RegisterOperatorToAVS(&_HelloWorld.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_HelloWorld *HelloWorldTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _HelloWorld.Contract.RegisterOperatorToAVS(&_HelloWorld.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HelloWorld *HelloWorldTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HelloWorld.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HelloWorld *HelloWorldSession) RenounceOwnership() (*types.Transaction, error) {
	return _HelloWorld.Contract.RenounceOwnership(&_HelloWorld.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HelloWorld *HelloWorldTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HelloWorld.Contract.RenounceOwnership(&_HelloWorld.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x3415a49c.
//
// Solidity: function respondToTask((string,uint32) task, uint32 referenceTaskIndex, bytes signature) returns()
func (_HelloWorld *HelloWorldTransactor) RespondToTask(opts *bind.TransactOpts, task IHelloWorldServiceManagerTask, referenceTaskIndex uint32, signature []byte) (*types.Transaction, error) {
	return _HelloWorld.contract.Transact(opts, "respondToTask", task, referenceTaskIndex, signature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x3415a49c.
//
// Solidity: function respondToTask((string,uint32) task, uint32 referenceTaskIndex, bytes signature) returns()
func (_HelloWorld *HelloWorldSession) RespondToTask(task IHelloWorldServiceManagerTask, referenceTaskIndex uint32, signature []byte) (*types.Transaction, error) {
	return _HelloWorld.Contract.RespondToTask(&_HelloWorld.TransactOpts, task, referenceTaskIndex, signature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x3415a49c.
//
// Solidity: function respondToTask((string,uint32) task, uint32 referenceTaskIndex, bytes signature) returns()
func (_HelloWorld *HelloWorldTransactorSession) RespondToTask(task IHelloWorldServiceManagerTask, referenceTaskIndex uint32, signature []byte) (*types.Transaction, error) {
	return _HelloWorld.Contract.RespondToTask(&_HelloWorld.TransactOpts, task, referenceTaskIndex, signature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_HelloWorld *HelloWorldTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _HelloWorld.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_HelloWorld *HelloWorldSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _HelloWorld.Contract.SetPauserRegistry(&_HelloWorld.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_HelloWorld *HelloWorldTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _HelloWorld.Contract.SetPauserRegistry(&_HelloWorld.TransactOpts, newPauserRegistry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HelloWorld *HelloWorldTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HelloWorld.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HelloWorld *HelloWorldSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HelloWorld.Contract.TransferOwnership(&_HelloWorld.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HelloWorld *HelloWorldTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HelloWorld.Contract.TransferOwnership(&_HelloWorld.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_HelloWorld *HelloWorldTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _HelloWorld.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_HelloWorld *HelloWorldSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _HelloWorld.Contract.Unpause(&_HelloWorld.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_HelloWorld *HelloWorldTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _HelloWorld.Contract.Unpause(&_HelloWorld.TransactOpts, newPausedStatus)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_HelloWorld *HelloWorldTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _HelloWorld.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_HelloWorld *HelloWorldSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _HelloWorld.Contract.UpdateAVSMetadataURI(&_HelloWorld.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_HelloWorld *HelloWorldTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _HelloWorld.Contract.UpdateAVSMetadataURI(&_HelloWorld.TransactOpts, _metadataURI)
}

// HelloWorldInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the HelloWorld contract.
type HelloWorldInitializedIterator struct {
	Event *HelloWorldInitialized // Event containing the contract specifics and raw log

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
func (it *HelloWorldInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloWorldInitialized)
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
		it.Event = new(HelloWorldInitialized)
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
func (it *HelloWorldInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloWorldInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloWorldInitialized represents a Initialized event raised by the HelloWorld contract.
type HelloWorldInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HelloWorld *HelloWorldFilterer) FilterInitialized(opts *bind.FilterOpts) (*HelloWorldInitializedIterator, error) {

	logs, sub, err := _HelloWorld.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &HelloWorldInitializedIterator{contract: _HelloWorld.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HelloWorld *HelloWorldFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *HelloWorldInitialized) (event.Subscription, error) {

	logs, sub, err := _HelloWorld.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloWorldInitialized)
				if err := _HelloWorld.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HelloWorld *HelloWorldFilterer) ParseInitialized(log types.Log) (*HelloWorldInitialized, error) {
	event := new(HelloWorldInitialized)
	if err := _HelloWorld.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HelloWorldNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the HelloWorld contract.
type HelloWorldNewTaskCreatedIterator struct {
	Event *HelloWorldNewTaskCreated // Event containing the contract specifics and raw log

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
func (it *HelloWorldNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloWorldNewTaskCreated)
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
		it.Event = new(HelloWorldNewTaskCreated)
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
func (it *HelloWorldNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloWorldNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloWorldNewTaskCreated represents a NewTaskCreated event raised by the HelloWorld contract.
type HelloWorldNewTaskCreated struct {
	TaskIndex uint32
	Task      IHelloWorldServiceManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x58180a6a0403a63c2b5ce4b85d129d46a80d37851b2216bd0a98b59e7309b847.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (string,uint32) task)
func (_HelloWorld *HelloWorldFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*HelloWorldNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _HelloWorld.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &HelloWorldNewTaskCreatedIterator{contract: _HelloWorld.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x58180a6a0403a63c2b5ce4b85d129d46a80d37851b2216bd0a98b59e7309b847.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (string,uint32) task)
func (_HelloWorld *HelloWorldFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *HelloWorldNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _HelloWorld.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloWorldNewTaskCreated)
				if err := _HelloWorld.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
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

// ParseNewTaskCreated is a log parse operation binding the contract event 0x58180a6a0403a63c2b5ce4b85d129d46a80d37851b2216bd0a98b59e7309b847.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (string,uint32) task)
func (_HelloWorld *HelloWorldFilterer) ParseNewTaskCreated(log types.Log) (*HelloWorldNewTaskCreated, error) {
	event := new(HelloWorldNewTaskCreated)
	if err := _HelloWorld.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HelloWorldOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HelloWorld contract.
type HelloWorldOwnershipTransferredIterator struct {
	Event *HelloWorldOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HelloWorldOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloWorldOwnershipTransferred)
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
		it.Event = new(HelloWorldOwnershipTransferred)
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
func (it *HelloWorldOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloWorldOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloWorldOwnershipTransferred represents a OwnershipTransferred event raised by the HelloWorld contract.
type HelloWorldOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HelloWorld *HelloWorldFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HelloWorldOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HelloWorld.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HelloWorldOwnershipTransferredIterator{contract: _HelloWorld.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HelloWorld *HelloWorldFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HelloWorldOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HelloWorld.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloWorldOwnershipTransferred)
				if err := _HelloWorld.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_HelloWorld *HelloWorldFilterer) ParseOwnershipTransferred(log types.Log) (*HelloWorldOwnershipTransferred, error) {
	event := new(HelloWorldOwnershipTransferred)
	if err := _HelloWorld.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HelloWorldPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the HelloWorld contract.
type HelloWorldPausedIterator struct {
	Event *HelloWorldPaused // Event containing the contract specifics and raw log

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
func (it *HelloWorldPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloWorldPaused)
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
		it.Event = new(HelloWorldPaused)
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
func (it *HelloWorldPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloWorldPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloWorldPaused represents a Paused event raised by the HelloWorld contract.
type HelloWorldPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_HelloWorld *HelloWorldFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*HelloWorldPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _HelloWorld.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &HelloWorldPausedIterator{contract: _HelloWorld.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_HelloWorld *HelloWorldFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *HelloWorldPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _HelloWorld.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloWorldPaused)
				if err := _HelloWorld.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_HelloWorld *HelloWorldFilterer) ParsePaused(log types.Log) (*HelloWorldPaused, error) {
	event := new(HelloWorldPaused)
	if err := _HelloWorld.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HelloWorldPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the HelloWorld contract.
type HelloWorldPauserRegistrySetIterator struct {
	Event *HelloWorldPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *HelloWorldPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloWorldPauserRegistrySet)
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
		it.Event = new(HelloWorldPauserRegistrySet)
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
func (it *HelloWorldPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloWorldPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloWorldPauserRegistrySet represents a PauserRegistrySet event raised by the HelloWorld contract.
type HelloWorldPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_HelloWorld *HelloWorldFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*HelloWorldPauserRegistrySetIterator, error) {

	logs, sub, err := _HelloWorld.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &HelloWorldPauserRegistrySetIterator{contract: _HelloWorld.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_HelloWorld *HelloWorldFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *HelloWorldPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _HelloWorld.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloWorldPauserRegistrySet)
				if err := _HelloWorld.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_HelloWorld *HelloWorldFilterer) ParsePauserRegistrySet(log types.Log) (*HelloWorldPauserRegistrySet, error) {
	event := new(HelloWorldPauserRegistrySet)
	if err := _HelloWorld.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HelloWorldTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the HelloWorld contract.
type HelloWorldTaskRespondedIterator struct {
	Event *HelloWorldTaskResponded // Event containing the contract specifics and raw log

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
func (it *HelloWorldTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloWorldTaskResponded)
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
		it.Event = new(HelloWorldTaskResponded)
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
func (it *HelloWorldTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloWorldTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloWorldTaskResponded represents a TaskResponded event raised by the HelloWorld contract.
type HelloWorldTaskResponded struct {
	TaskIndex uint32
	Task      IHelloWorldServiceManagerTask
	Operator  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x8eb2d2fcccf5801e10ff58cd73e8781ba923122963789378771f03c1148b023e.
//
// Solidity: event TaskResponded(uint32 indexed taskIndex, (string,uint32) task, address operator)
func (_HelloWorld *HelloWorldFilterer) FilterTaskResponded(opts *bind.FilterOpts, taskIndex []uint32) (*HelloWorldTaskRespondedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _HelloWorld.contract.FilterLogs(opts, "TaskResponded", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &HelloWorldTaskRespondedIterator{contract: _HelloWorld.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x8eb2d2fcccf5801e10ff58cd73e8781ba923122963789378771f03c1148b023e.
//
// Solidity: event TaskResponded(uint32 indexed taskIndex, (string,uint32) task, address operator)
func (_HelloWorld *HelloWorldFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *HelloWorldTaskResponded, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _HelloWorld.contract.WatchLogs(opts, "TaskResponded", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloWorldTaskResponded)
				if err := _HelloWorld.contract.UnpackLog(event, "TaskResponded", log); err != nil {
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

// ParseTaskResponded is a log parse operation binding the contract event 0x8eb2d2fcccf5801e10ff58cd73e8781ba923122963789378771f03c1148b023e.
//
// Solidity: event TaskResponded(uint32 indexed taskIndex, (string,uint32) task, address operator)
func (_HelloWorld *HelloWorldFilterer) ParseTaskResponded(log types.Log) (*HelloWorldTaskResponded, error) {
	event := new(HelloWorldTaskResponded)
	if err := _HelloWorld.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HelloWorldUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the HelloWorld contract.
type HelloWorldUnpausedIterator struct {
	Event *HelloWorldUnpaused // Event containing the contract specifics and raw log

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
func (it *HelloWorldUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HelloWorldUnpaused)
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
		it.Event = new(HelloWorldUnpaused)
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
func (it *HelloWorldUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HelloWorldUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HelloWorldUnpaused represents a Unpaused event raised by the HelloWorld contract.
type HelloWorldUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_HelloWorld *HelloWorldFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*HelloWorldUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _HelloWorld.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &HelloWorldUnpausedIterator{contract: _HelloWorld.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_HelloWorld *HelloWorldFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *HelloWorldUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _HelloWorld.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HelloWorldUnpaused)
				if err := _HelloWorld.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_HelloWorld *HelloWorldFilterer) ParseUnpaused(log types.Log) (*HelloWorldUnpaused, error) {
	event := new(HelloWorldUnpaused)
	if err := _HelloWorld.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
