// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gobind

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

// ISwapicaCreateOrderRequest is an auto generated low-level Go binding around an user-defined struct.
type ISwapicaCreateOrderRequest struct {
	UseRelayer       bool
	TokenToSell      common.Address
	AmountToSell     *big.Int
	TokenToBuy       common.Address
	AmountToBuy      *big.Int
	DestinationChain *big.Int
}

// ISwapicaMatch is an auto generated low-level Go binding around an user-defined struct.
type ISwapicaMatch struct {
	State         uint8
	MatchId       *big.Int
	OriginOrderId *big.Int
	Creator       common.Address
	TokenToSell   common.Address
	AmountToSell  *big.Int
	OriginChainId *big.Int
}

// ISwapicaOrder is an auto generated low-level Go binding around an user-defined struct.
type ISwapicaOrder struct {
	Status           ISwapicaOrderStatus
	OrderId          *big.Int
	Creator          common.Address
	TokenToSell      common.Address
	AmountToSell     *big.Int
	TokenToBuy       common.Address
	AmountToBuy      *big.Int
	DestinationChain *big.Int
}

// ISwapicaOrderStatus is an auto generated low-level Go binding around an user-defined struct.
type ISwapicaOrderStatus struct {
	State        uint8
	MatchId      *big.Int
	MatchSwapica common.Address
}

// SwapicaMetaData contains all meta data concerning the Swapica contract.
var SwapicaMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"enumISwapica.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originOrderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenToSell\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToSell\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISwapica.Match\",\"name\":\"match_\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useRelayer\",\"type\":\"bool\"}],\"name\":\"MatchCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumISwapica.State\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"MatchUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumISwapica.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"matchSwapica\",\"type\":\"address\"}],\"internalType\":\"structISwapica.OrderStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenToSell\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToSell\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenToBuy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChain\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISwapica.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useRelayer\",\"type\":\"bool\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumISwapica.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"matchSwapica\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structISwapica.OrderStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"name\":\"OrderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_signaturesThreshold\",\"type\":\"uint256\"}],\"name\":\"__Signers_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"}],\"name\":\"__Swapica_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"}],\"name\":\"addSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"cancelMatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"checkSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"createMatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"useRelayer\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"tokenToSell\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToSell\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenToBuy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChain\",\"type\":\"uint256\"}],\"internalType\":\"structISwapica.CreateOrderRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"executeMatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"executeOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getAllMatches\",\"outputs\":[{\"components\":[{\"internalType\":\"enumISwapica.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originOrderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenToSell\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToSell\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"}],\"internalType\":\"structISwapica.Match[]\",\"name\":\"allMatches\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllMatchesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getAllOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumISwapica.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"matchSwapica\",\"type\":\"address\"}],\"internalType\":\"structISwapica.OrderStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenToSell\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToSell\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenToBuy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChain\",\"type\":\"uint256\"}],\"internalType\":\"structISwapica.Order[]\",\"name\":\"allOrders\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllOrdersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getUserMatches\",\"outputs\":[{\"components\":[{\"internalType\":\"enumISwapica.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originOrderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenToSell\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToSell\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"}],\"internalType\":\"structISwapica.Match[]\",\"name\":\"userMatches\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserMatchesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getUserOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumISwapica.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"matchSwapica\",\"type\":\"address\"}],\"internalType\":\"structISwapica.OrderStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenToSell\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToSell\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenToBuy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChain\",\"type\":\"uint256\"}],\"internalType\":\"structISwapica.Order[]\",\"name\":\"userOrders\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserOrdersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"}],\"name\":\"removeSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_signaturesThreshold\",\"type\":\"uint256\"}],\"name\":\"setSignaturesThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signaturesThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// SwapicaABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapicaMetaData.ABI instead.
var SwapicaABI = SwapicaMetaData.ABI

// Swapica is an auto generated Go binding around an Ethereum contract.
type Swapica struct {
	SwapicaCaller     // Read-only binding to the contract
	SwapicaTransactor // Write-only binding to the contract
	SwapicaFilterer   // Log filterer for contract events
}

// SwapicaCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapicaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapicaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapicaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapicaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapicaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapicaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapicaSession struct {
	Contract     *Swapica          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapicaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapicaCallerSession struct {
	Contract *SwapicaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SwapicaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapicaTransactorSession struct {
	Contract     *SwapicaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SwapicaRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapicaRaw struct {
	Contract *Swapica // Generic contract binding to access the raw methods on
}

// SwapicaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapicaCallerRaw struct {
	Contract *SwapicaCaller // Generic read-only contract binding to access the raw methods on
}

// SwapicaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapicaTransactorRaw struct {
	Contract *SwapicaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapica creates a new instance of Swapica, bound to a specific deployed contract.
func NewSwapica(address common.Address, backend bind.ContractBackend) (*Swapica, error) {
	contract, err := bindSwapica(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swapica{SwapicaCaller: SwapicaCaller{contract: contract}, SwapicaTransactor: SwapicaTransactor{contract: contract}, SwapicaFilterer: SwapicaFilterer{contract: contract}}, nil
}

// NewSwapicaCaller creates a new read-only instance of Swapica, bound to a specific deployed contract.
func NewSwapicaCaller(address common.Address, caller bind.ContractCaller) (*SwapicaCaller, error) {
	contract, err := bindSwapica(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapicaCaller{contract: contract}, nil
}

// NewSwapicaTransactor creates a new write-only instance of Swapica, bound to a specific deployed contract.
func NewSwapicaTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapicaTransactor, error) {
	contract, err := bindSwapica(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapicaTransactor{contract: contract}, nil
}

// NewSwapicaFilterer creates a new log filterer instance of Swapica, bound to a specific deployed contract.
func NewSwapicaFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapicaFilterer, error) {
	contract, err := bindSwapica(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapicaFilterer{contract: contract}, nil
}

// bindSwapica binds a generic wrapper to an already deployed contract.
func bindSwapica(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapicaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapica *SwapicaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapica.Contract.SwapicaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapica *SwapicaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapica.Contract.SwapicaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapica *SwapicaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapica.Contract.SwapicaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapica *SwapicaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapica.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapica *SwapicaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapica.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapica *SwapicaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapica.Contract.contract.Transact(opts, method, params...)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x5a0f8830.
//
// Solidity: function checkSignatures(bytes32 signHash, bytes[] signatures) view returns()
func (_Swapica *SwapicaCaller) CheckSignatures(opts *bind.CallOpts, signHash [32]byte, signatures [][]byte) error {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "checkSignatures", signHash, signatures)

	if err != nil {
		return err
	}

	return err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x5a0f8830.
//
// Solidity: function checkSignatures(bytes32 signHash, bytes[] signatures) view returns()
func (_Swapica *SwapicaSession) CheckSignatures(signHash [32]byte, signatures [][]byte) error {
	return _Swapica.Contract.CheckSignatures(&_Swapica.CallOpts, signHash, signatures)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x5a0f8830.
//
// Solidity: function checkSignatures(bytes32 signHash, bytes[] signatures) view returns()
func (_Swapica *SwapicaCallerSession) CheckSignatures(signHash [32]byte, signatures [][]byte) error {
	return _Swapica.Contract.CheckSignatures(&_Swapica.CallOpts, signHash, signatures)
}

// GetAllMatches is a free data retrieval call binding the contract method 0xc8866955.
//
// Solidity: function getAllMatches(uint256 offset, uint256 limit) view returns((uint8,uint256,uint256,address,address,uint256,uint256)[] allMatches)
func (_Swapica *SwapicaCaller) GetAllMatches(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]ISwapicaMatch, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "getAllMatches", offset, limit)

	if err != nil {
		return *new([]ISwapicaMatch), err
	}

	out0 := *abi.ConvertType(out[0], new([]ISwapicaMatch)).(*[]ISwapicaMatch)

	return out0, err

}

// GetAllMatches is a free data retrieval call binding the contract method 0xc8866955.
//
// Solidity: function getAllMatches(uint256 offset, uint256 limit) view returns((uint8,uint256,uint256,address,address,uint256,uint256)[] allMatches)
func (_Swapica *SwapicaSession) GetAllMatches(offset *big.Int, limit *big.Int) ([]ISwapicaMatch, error) {
	return _Swapica.Contract.GetAllMatches(&_Swapica.CallOpts, offset, limit)
}

// GetAllMatches is a free data retrieval call binding the contract method 0xc8866955.
//
// Solidity: function getAllMatches(uint256 offset, uint256 limit) view returns((uint8,uint256,uint256,address,address,uint256,uint256)[] allMatches)
func (_Swapica *SwapicaCallerSession) GetAllMatches(offset *big.Int, limit *big.Int) ([]ISwapicaMatch, error) {
	return _Swapica.Contract.GetAllMatches(&_Swapica.CallOpts, offset, limit)
}

// GetAllMatchesLength is a free data retrieval call binding the contract method 0xf3ed9e50.
//
// Solidity: function getAllMatchesLength() view returns(uint256)
func (_Swapica *SwapicaCaller) GetAllMatchesLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "getAllMatchesLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllMatchesLength is a free data retrieval call binding the contract method 0xf3ed9e50.
//
// Solidity: function getAllMatchesLength() view returns(uint256)
func (_Swapica *SwapicaSession) GetAllMatchesLength() (*big.Int, error) {
	return _Swapica.Contract.GetAllMatchesLength(&_Swapica.CallOpts)
}

// GetAllMatchesLength is a free data retrieval call binding the contract method 0xf3ed9e50.
//
// Solidity: function getAllMatchesLength() view returns(uint256)
func (_Swapica *SwapicaCallerSession) GetAllMatchesLength() (*big.Int, error) {
	return _Swapica.Contract.GetAllMatchesLength(&_Swapica.CallOpts)
}

// GetAllOrders is a free data retrieval call binding the contract method 0x418f2845.
//
// Solidity: function getAllOrders(uint256 offset, uint256 limit) view returns(((uint8,uint256,address),uint256,address,address,uint256,address,uint256,uint256)[] allOrders)
func (_Swapica *SwapicaCaller) GetAllOrders(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]ISwapicaOrder, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "getAllOrders", offset, limit)

	if err != nil {
		return *new([]ISwapicaOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]ISwapicaOrder)).(*[]ISwapicaOrder)

	return out0, err

}

// GetAllOrders is a free data retrieval call binding the contract method 0x418f2845.
//
// Solidity: function getAllOrders(uint256 offset, uint256 limit) view returns(((uint8,uint256,address),uint256,address,address,uint256,address,uint256,uint256)[] allOrders)
func (_Swapica *SwapicaSession) GetAllOrders(offset *big.Int, limit *big.Int) ([]ISwapicaOrder, error) {
	return _Swapica.Contract.GetAllOrders(&_Swapica.CallOpts, offset, limit)
}

// GetAllOrders is a free data retrieval call binding the contract method 0x418f2845.
//
// Solidity: function getAllOrders(uint256 offset, uint256 limit) view returns(((uint8,uint256,address),uint256,address,address,uint256,address,uint256,uint256)[] allOrders)
func (_Swapica *SwapicaCallerSession) GetAllOrders(offset *big.Int, limit *big.Int) ([]ISwapicaOrder, error) {
	return _Swapica.Contract.GetAllOrders(&_Swapica.CallOpts, offset, limit)
}

// GetAllOrdersLength is a free data retrieval call binding the contract method 0xb06d498e.
//
// Solidity: function getAllOrdersLength() view returns(uint256)
func (_Swapica *SwapicaCaller) GetAllOrdersLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "getAllOrdersLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllOrdersLength is a free data retrieval call binding the contract method 0xb06d498e.
//
// Solidity: function getAllOrdersLength() view returns(uint256)
func (_Swapica *SwapicaSession) GetAllOrdersLength() (*big.Int, error) {
	return _Swapica.Contract.GetAllOrdersLength(&_Swapica.CallOpts)
}

// GetAllOrdersLength is a free data retrieval call binding the contract method 0xb06d498e.
//
// Solidity: function getAllOrdersLength() view returns(uint256)
func (_Swapica *SwapicaCallerSession) GetAllOrdersLength() (*big.Int, error) {
	return _Swapica.Contract.GetAllOrdersLength(&_Swapica.CallOpts)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_Swapica *SwapicaCaller) GetSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "getSigners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_Swapica *SwapicaSession) GetSigners() ([]common.Address, error) {
	return _Swapica.Contract.GetSigners(&_Swapica.CallOpts)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_Swapica *SwapicaCallerSession) GetSigners() ([]common.Address, error) {
	return _Swapica.Contract.GetSigners(&_Swapica.CallOpts)
}

// GetUserMatches is a free data retrieval call binding the contract method 0xecf1d9f9.
//
// Solidity: function getUserMatches(address user, uint256 offset, uint256 limit) view returns((uint8,uint256,uint256,address,address,uint256,uint256)[] userMatches)
func (_Swapica *SwapicaCaller) GetUserMatches(opts *bind.CallOpts, user common.Address, offset *big.Int, limit *big.Int) ([]ISwapicaMatch, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "getUserMatches", user, offset, limit)

	if err != nil {
		return *new([]ISwapicaMatch), err
	}

	out0 := *abi.ConvertType(out[0], new([]ISwapicaMatch)).(*[]ISwapicaMatch)

	return out0, err

}

// GetUserMatches is a free data retrieval call binding the contract method 0xecf1d9f9.
//
// Solidity: function getUserMatches(address user, uint256 offset, uint256 limit) view returns((uint8,uint256,uint256,address,address,uint256,uint256)[] userMatches)
func (_Swapica *SwapicaSession) GetUserMatches(user common.Address, offset *big.Int, limit *big.Int) ([]ISwapicaMatch, error) {
	return _Swapica.Contract.GetUserMatches(&_Swapica.CallOpts, user, offset, limit)
}

// GetUserMatches is a free data retrieval call binding the contract method 0xecf1d9f9.
//
// Solidity: function getUserMatches(address user, uint256 offset, uint256 limit) view returns((uint8,uint256,uint256,address,address,uint256,uint256)[] userMatches)
func (_Swapica *SwapicaCallerSession) GetUserMatches(user common.Address, offset *big.Int, limit *big.Int) ([]ISwapicaMatch, error) {
	return _Swapica.Contract.GetUserMatches(&_Swapica.CallOpts, user, offset, limit)
}

// GetUserMatchesLength is a free data retrieval call binding the contract method 0x026f562d.
//
// Solidity: function getUserMatchesLength(address user) view returns(uint256)
func (_Swapica *SwapicaCaller) GetUserMatchesLength(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "getUserMatchesLength", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserMatchesLength is a free data retrieval call binding the contract method 0x026f562d.
//
// Solidity: function getUserMatchesLength(address user) view returns(uint256)
func (_Swapica *SwapicaSession) GetUserMatchesLength(user common.Address) (*big.Int, error) {
	return _Swapica.Contract.GetUserMatchesLength(&_Swapica.CallOpts, user)
}

// GetUserMatchesLength is a free data retrieval call binding the contract method 0x026f562d.
//
// Solidity: function getUserMatchesLength(address user) view returns(uint256)
func (_Swapica *SwapicaCallerSession) GetUserMatchesLength(user common.Address) (*big.Int, error) {
	return _Swapica.Contract.GetUserMatchesLength(&_Swapica.CallOpts, user)
}

// GetUserOrders is a free data retrieval call binding the contract method 0xc7efbe36.
//
// Solidity: function getUserOrders(address user, uint256 offset, uint256 limit) view returns(((uint8,uint256,address),uint256,address,address,uint256,address,uint256,uint256)[] userOrders)
func (_Swapica *SwapicaCaller) GetUserOrders(opts *bind.CallOpts, user common.Address, offset *big.Int, limit *big.Int) ([]ISwapicaOrder, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "getUserOrders", user, offset, limit)

	if err != nil {
		return *new([]ISwapicaOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]ISwapicaOrder)).(*[]ISwapicaOrder)

	return out0, err

}

// GetUserOrders is a free data retrieval call binding the contract method 0xc7efbe36.
//
// Solidity: function getUserOrders(address user, uint256 offset, uint256 limit) view returns(((uint8,uint256,address),uint256,address,address,uint256,address,uint256,uint256)[] userOrders)
func (_Swapica *SwapicaSession) GetUserOrders(user common.Address, offset *big.Int, limit *big.Int) ([]ISwapicaOrder, error) {
	return _Swapica.Contract.GetUserOrders(&_Swapica.CallOpts, user, offset, limit)
}

// GetUserOrders is a free data retrieval call binding the contract method 0xc7efbe36.
//
// Solidity: function getUserOrders(address user, uint256 offset, uint256 limit) view returns(((uint8,uint256,address),uint256,address,address,uint256,address,uint256,uint256)[] userOrders)
func (_Swapica *SwapicaCallerSession) GetUserOrders(user common.Address, offset *big.Int, limit *big.Int) ([]ISwapicaOrder, error) {
	return _Swapica.Contract.GetUserOrders(&_Swapica.CallOpts, user, offset, limit)
}

// GetUserOrdersLength is a free data retrieval call binding the contract method 0x7a7007a8.
//
// Solidity: function getUserOrdersLength(address user) view returns(uint256)
func (_Swapica *SwapicaCaller) GetUserOrdersLength(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "getUserOrdersLength", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserOrdersLength is a free data retrieval call binding the contract method 0x7a7007a8.
//
// Solidity: function getUserOrdersLength(address user) view returns(uint256)
func (_Swapica *SwapicaSession) GetUserOrdersLength(user common.Address) (*big.Int, error) {
	return _Swapica.Contract.GetUserOrdersLength(&_Swapica.CallOpts, user)
}

// GetUserOrdersLength is a free data retrieval call binding the contract method 0x7a7007a8.
//
// Solidity: function getUserOrdersLength(address user) view returns(uint256)
func (_Swapica *SwapicaCallerSession) GetUserOrdersLength(user common.Address) (*big.Int, error) {
	return _Swapica.Contract.GetUserOrdersLength(&_Swapica.CallOpts, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swapica *SwapicaCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swapica *SwapicaSession) Owner() (common.Address, error) {
	return _Swapica.Contract.Owner(&_Swapica.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swapica *SwapicaCallerSession) Owner() (common.Address, error) {
	return _Swapica.Contract.Owner(&_Swapica.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Swapica *SwapicaCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Swapica *SwapicaSession) ProxiableUUID() ([32]byte, error) {
	return _Swapica.Contract.ProxiableUUID(&_Swapica.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Swapica *SwapicaCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Swapica.Contract.ProxiableUUID(&_Swapica.CallOpts)
}

// SignaturesThreshold is a free data retrieval call binding the contract method 0x39ce73c7.
//
// Solidity: function signaturesThreshold() view returns(uint256)
func (_Swapica *SwapicaCaller) SignaturesThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "signaturesThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignaturesThreshold is a free data retrieval call binding the contract method 0x39ce73c7.
//
// Solidity: function signaturesThreshold() view returns(uint256)
func (_Swapica *SwapicaSession) SignaturesThreshold() (*big.Int, error) {
	return _Swapica.Contract.SignaturesThreshold(&_Swapica.CallOpts)
}

// SignaturesThreshold is a free data retrieval call binding the contract method 0x39ce73c7.
//
// Solidity: function signaturesThreshold() view returns(uint256)
func (_Swapica *SwapicaCallerSession) SignaturesThreshold() (*big.Int, error) {
	return _Swapica.Contract.SignaturesThreshold(&_Swapica.CallOpts)
}

// SignersInit is a paid mutator transaction binding the contract method 0x09a55841.
//
// Solidity: function __Signers_init(address[] signers, uint256 _signaturesThreshold) returns()
func (_Swapica *SwapicaTransactor) SignersInit(opts *bind.TransactOpts, signers []common.Address, _signaturesThreshold *big.Int) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "__Signers_init", signers, _signaturesThreshold)
}

// SignersInit is a paid mutator transaction binding the contract method 0x09a55841.
//
// Solidity: function __Signers_init(address[] signers, uint256 _signaturesThreshold) returns()
func (_Swapica *SwapicaSession) SignersInit(signers []common.Address, _signaturesThreshold *big.Int) (*types.Transaction, error) {
	return _Swapica.Contract.SignersInit(&_Swapica.TransactOpts, signers, _signaturesThreshold)
}

// SignersInit is a paid mutator transaction binding the contract method 0x09a55841.
//
// Solidity: function __Signers_init(address[] signers, uint256 _signaturesThreshold) returns()
func (_Swapica *SwapicaTransactorSession) SignersInit(signers []common.Address, _signaturesThreshold *big.Int) (*types.Transaction, error) {
	return _Swapica.Contract.SignersInit(&_Swapica.TransactOpts, signers, _signaturesThreshold)
}

// SwapicaInit is a paid mutator transaction binding the contract method 0x86b2790c.
//
// Solidity: function __Swapica_init(address[] signers) returns()
func (_Swapica *SwapicaTransactor) SwapicaInit(opts *bind.TransactOpts, signers []common.Address) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "__Swapica_init", signers)
}

// SwapicaInit is a paid mutator transaction binding the contract method 0x86b2790c.
//
// Solidity: function __Swapica_init(address[] signers) returns()
func (_Swapica *SwapicaSession) SwapicaInit(signers []common.Address) (*types.Transaction, error) {
	return _Swapica.Contract.SwapicaInit(&_Swapica.TransactOpts, signers)
}

// SwapicaInit is a paid mutator transaction binding the contract method 0x86b2790c.
//
// Solidity: function __Swapica_init(address[] signers) returns()
func (_Swapica *SwapicaTransactorSession) SwapicaInit(signers []common.Address) (*types.Transaction, error) {
	return _Swapica.Contract.SwapicaInit(&_Swapica.TransactOpts, signers)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] signers) returns()
func (_Swapica *SwapicaTransactor) AddSigners(opts *bind.TransactOpts, signers []common.Address) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "addSigners", signers)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] signers) returns()
func (_Swapica *SwapicaSession) AddSigners(signers []common.Address) (*types.Transaction, error) {
	return _Swapica.Contract.AddSigners(&_Swapica.TransactOpts, signers)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] signers) returns()
func (_Swapica *SwapicaTransactorSession) AddSigners(signers []common.Address) (*types.Transaction, error) {
	return _Swapica.Contract.AddSigners(&_Swapica.TransactOpts, signers)
}

// CancelMatch is a paid mutator transaction binding the contract method 0xb925c02d.
//
// Solidity: function cancelMatch(bytes data, bytes[] signatures) returns()
func (_Swapica *SwapicaTransactor) CancelMatch(opts *bind.TransactOpts, data []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "cancelMatch", data, signatures)
}

// CancelMatch is a paid mutator transaction binding the contract method 0xb925c02d.
//
// Solidity: function cancelMatch(bytes data, bytes[] signatures) returns()
func (_Swapica *SwapicaSession) CancelMatch(data []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.Contract.CancelMatch(&_Swapica.TransactOpts, data, signatures)
}

// CancelMatch is a paid mutator transaction binding the contract method 0xb925c02d.
//
// Solidity: function cancelMatch(bytes data, bytes[] signatures) returns()
func (_Swapica *SwapicaTransactorSession) CancelMatch(data []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.Contract.CancelMatch(&_Swapica.TransactOpts, data, signatures)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 orderId) returns()
func (_Swapica *SwapicaTransactor) CancelOrder(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "cancelOrder", orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 orderId) returns()
func (_Swapica *SwapicaSession) CancelOrder(orderId *big.Int) (*types.Transaction, error) {
	return _Swapica.Contract.CancelOrder(&_Swapica.TransactOpts, orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 orderId) returns()
func (_Swapica *SwapicaTransactorSession) CancelOrder(orderId *big.Int) (*types.Transaction, error) {
	return _Swapica.Contract.CancelOrder(&_Swapica.TransactOpts, orderId)
}

// CreateMatch is a paid mutator transaction binding the contract method 0xf7662788.
//
// Solidity: function createMatch(bytes data, bytes[] signatures) payable returns()
func (_Swapica *SwapicaTransactor) CreateMatch(opts *bind.TransactOpts, data []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "createMatch", data, signatures)
}

// CreateMatch is a paid mutator transaction binding the contract method 0xf7662788.
//
// Solidity: function createMatch(bytes data, bytes[] signatures) payable returns()
func (_Swapica *SwapicaSession) CreateMatch(data []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.Contract.CreateMatch(&_Swapica.TransactOpts, data, signatures)
}

// CreateMatch is a paid mutator transaction binding the contract method 0xf7662788.
//
// Solidity: function createMatch(bytes data, bytes[] signatures) payable returns()
func (_Swapica *SwapicaTransactorSession) CreateMatch(data []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.Contract.CreateMatch(&_Swapica.TransactOpts, data, signatures)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x61014c65.
//
// Solidity: function createOrder((bool,address,uint256,address,uint256,uint256) request) payable returns()
func (_Swapica *SwapicaTransactor) CreateOrder(opts *bind.TransactOpts, request ISwapicaCreateOrderRequest) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "createOrder", request)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x61014c65.
//
// Solidity: function createOrder((bool,address,uint256,address,uint256,uint256) request) payable returns()
func (_Swapica *SwapicaSession) CreateOrder(request ISwapicaCreateOrderRequest) (*types.Transaction, error) {
	return _Swapica.Contract.CreateOrder(&_Swapica.TransactOpts, request)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x61014c65.
//
// Solidity: function createOrder((bool,address,uint256,address,uint256,uint256) request) payable returns()
func (_Swapica *SwapicaTransactorSession) CreateOrder(request ISwapicaCreateOrderRequest) (*types.Transaction, error) {
	return _Swapica.Contract.CreateOrder(&_Swapica.TransactOpts, request)
}

// ExecuteMatch is a paid mutator transaction binding the contract method 0x2a9ff0d5.
//
// Solidity: function executeMatch(bytes data, bytes[] signatures) returns()
func (_Swapica *SwapicaTransactor) ExecuteMatch(opts *bind.TransactOpts, data []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "executeMatch", data, signatures)
}

// ExecuteMatch is a paid mutator transaction binding the contract method 0x2a9ff0d5.
//
// Solidity: function executeMatch(bytes data, bytes[] signatures) returns()
func (_Swapica *SwapicaSession) ExecuteMatch(data []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.Contract.ExecuteMatch(&_Swapica.TransactOpts, data, signatures)
}

// ExecuteMatch is a paid mutator transaction binding the contract method 0x2a9ff0d5.
//
// Solidity: function executeMatch(bytes data, bytes[] signatures) returns()
func (_Swapica *SwapicaTransactorSession) ExecuteMatch(data []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.Contract.ExecuteMatch(&_Swapica.TransactOpts, data, signatures)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x14d41b45.
//
// Solidity: function executeOrder(bytes data, bytes[] signatures) returns()
func (_Swapica *SwapicaTransactor) ExecuteOrder(opts *bind.TransactOpts, data []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "executeOrder", data, signatures)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x14d41b45.
//
// Solidity: function executeOrder(bytes data, bytes[] signatures) returns()
func (_Swapica *SwapicaSession) ExecuteOrder(data []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.Contract.ExecuteOrder(&_Swapica.TransactOpts, data, signatures)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x14d41b45.
//
// Solidity: function executeOrder(bytes data, bytes[] signatures) returns()
func (_Swapica *SwapicaTransactorSession) ExecuteOrder(data []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.Contract.ExecuteOrder(&_Swapica.TransactOpts, data, signatures)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x8d361e43.
//
// Solidity: function removeSigners(address[] signers) returns()
func (_Swapica *SwapicaTransactor) RemoveSigners(opts *bind.TransactOpts, signers []common.Address) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "removeSigners", signers)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x8d361e43.
//
// Solidity: function removeSigners(address[] signers) returns()
func (_Swapica *SwapicaSession) RemoveSigners(signers []common.Address) (*types.Transaction, error) {
	return _Swapica.Contract.RemoveSigners(&_Swapica.TransactOpts, signers)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x8d361e43.
//
// Solidity: function removeSigners(address[] signers) returns()
func (_Swapica *SwapicaTransactorSession) RemoveSigners(signers []common.Address) (*types.Transaction, error) {
	return _Swapica.Contract.RemoveSigners(&_Swapica.TransactOpts, signers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swapica *SwapicaTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swapica *SwapicaSession) RenounceOwnership() (*types.Transaction, error) {
	return _Swapica.Contract.RenounceOwnership(&_Swapica.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swapica *SwapicaTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Swapica.Contract.RenounceOwnership(&_Swapica.TransactOpts)
}

// SetSignaturesThreshold is a paid mutator transaction binding the contract method 0xbf1fe08f.
//
// Solidity: function setSignaturesThreshold(uint256 _signaturesThreshold) returns()
func (_Swapica *SwapicaTransactor) SetSignaturesThreshold(opts *bind.TransactOpts, _signaturesThreshold *big.Int) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "setSignaturesThreshold", _signaturesThreshold)
}

// SetSignaturesThreshold is a paid mutator transaction binding the contract method 0xbf1fe08f.
//
// Solidity: function setSignaturesThreshold(uint256 _signaturesThreshold) returns()
func (_Swapica *SwapicaSession) SetSignaturesThreshold(_signaturesThreshold *big.Int) (*types.Transaction, error) {
	return _Swapica.Contract.SetSignaturesThreshold(&_Swapica.TransactOpts, _signaturesThreshold)
}

// SetSignaturesThreshold is a paid mutator transaction binding the contract method 0xbf1fe08f.
//
// Solidity: function setSignaturesThreshold(uint256 _signaturesThreshold) returns()
func (_Swapica *SwapicaTransactorSession) SetSignaturesThreshold(_signaturesThreshold *big.Int) (*types.Transaction, error) {
	return _Swapica.Contract.SetSignaturesThreshold(&_Swapica.TransactOpts, _signaturesThreshold)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swapica *SwapicaTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swapica *SwapicaSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Swapica.Contract.TransferOwnership(&_Swapica.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swapica *SwapicaTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Swapica.Contract.TransferOwnership(&_Swapica.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Swapica *SwapicaTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Swapica *SwapicaSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Swapica.Contract.UpgradeTo(&_Swapica.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Swapica *SwapicaTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Swapica.Contract.UpgradeTo(&_Swapica.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Swapica *SwapicaTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Swapica *SwapicaSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Swapica.Contract.UpgradeToAndCall(&_Swapica.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Swapica *SwapicaTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Swapica.Contract.UpgradeToAndCall(&_Swapica.TransactOpts, newImplementation, data)
}

// SwapicaAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Swapica contract.
type SwapicaAdminChangedIterator struct {
	Event *SwapicaAdminChanged // Event containing the contract specifics and raw log

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
func (it *SwapicaAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapicaAdminChanged)
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
		it.Event = new(SwapicaAdminChanged)
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
func (it *SwapicaAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapicaAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapicaAdminChanged represents a AdminChanged event raised by the Swapica contract.
type SwapicaAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Swapica *SwapicaFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*SwapicaAdminChangedIterator, error) {

	logs, sub, err := _Swapica.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &SwapicaAdminChangedIterator{contract: _Swapica.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Swapica *SwapicaFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *SwapicaAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Swapica.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapicaAdminChanged)
				if err := _Swapica.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Swapica *SwapicaFilterer) ParseAdminChanged(log types.Log) (*SwapicaAdminChanged, error) {
	event := new(SwapicaAdminChanged)
	if err := _Swapica.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapicaBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Swapica contract.
type SwapicaBeaconUpgradedIterator struct {
	Event *SwapicaBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *SwapicaBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapicaBeaconUpgraded)
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
		it.Event = new(SwapicaBeaconUpgraded)
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
func (it *SwapicaBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapicaBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapicaBeaconUpgraded represents a BeaconUpgraded event raised by the Swapica contract.
type SwapicaBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Swapica *SwapicaFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*SwapicaBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Swapica.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &SwapicaBeaconUpgradedIterator{contract: _Swapica.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Swapica *SwapicaFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *SwapicaBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Swapica.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapicaBeaconUpgraded)
				if err := _Swapica.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Swapica *SwapicaFilterer) ParseBeaconUpgraded(log types.Log) (*SwapicaBeaconUpgraded, error) {
	event := new(SwapicaBeaconUpgraded)
	if err := _Swapica.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapicaInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Swapica contract.
type SwapicaInitializedIterator struct {
	Event *SwapicaInitialized // Event containing the contract specifics and raw log

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
func (it *SwapicaInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapicaInitialized)
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
		it.Event = new(SwapicaInitialized)
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
func (it *SwapicaInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapicaInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapicaInitialized represents a Initialized event raised by the Swapica contract.
type SwapicaInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Swapica *SwapicaFilterer) FilterInitialized(opts *bind.FilterOpts) (*SwapicaInitializedIterator, error) {

	logs, sub, err := _Swapica.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SwapicaInitializedIterator{contract: _Swapica.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Swapica *SwapicaFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SwapicaInitialized) (event.Subscription, error) {

	logs, sub, err := _Swapica.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapicaInitialized)
				if err := _Swapica.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Swapica *SwapicaFilterer) ParseInitialized(log types.Log) (*SwapicaInitialized, error) {
	event := new(SwapicaInitialized)
	if err := _Swapica.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapicaMatchCreatedIterator is returned from FilterMatchCreated and is used to iterate over the raw logs and unpacked data for MatchCreated events raised by the Swapica contract.
type SwapicaMatchCreatedIterator struct {
	Event *SwapicaMatchCreated // Event containing the contract specifics and raw log

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
func (it *SwapicaMatchCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapicaMatchCreated)
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
		it.Event = new(SwapicaMatchCreated)
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
func (it *SwapicaMatchCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapicaMatchCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapicaMatchCreated represents a MatchCreated event raised by the Swapica contract.
type SwapicaMatchCreated struct {
	Match      ISwapicaMatch
	UseRelayer bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMatchCreated is a free log retrieval operation binding the contract event 0x32619f811334dd1aff2e3cc0ed3b2697603d2439fcf63f47d1de5cc5971c8f89.
//
// Solidity: event MatchCreated((uint8,uint256,uint256,address,address,uint256,uint256) match_, bool useRelayer)
func (_Swapica *SwapicaFilterer) FilterMatchCreated(opts *bind.FilterOpts) (*SwapicaMatchCreatedIterator, error) {

	logs, sub, err := _Swapica.contract.FilterLogs(opts, "MatchCreated")
	if err != nil {
		return nil, err
	}
	return &SwapicaMatchCreatedIterator{contract: _Swapica.contract, event: "MatchCreated", logs: logs, sub: sub}, nil
}

// WatchMatchCreated is a free log subscription operation binding the contract event 0x32619f811334dd1aff2e3cc0ed3b2697603d2439fcf63f47d1de5cc5971c8f89.
//
// Solidity: event MatchCreated((uint8,uint256,uint256,address,address,uint256,uint256) match_, bool useRelayer)
func (_Swapica *SwapicaFilterer) WatchMatchCreated(opts *bind.WatchOpts, sink chan<- *SwapicaMatchCreated) (event.Subscription, error) {

	logs, sub, err := _Swapica.contract.WatchLogs(opts, "MatchCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapicaMatchCreated)
				if err := _Swapica.contract.UnpackLog(event, "MatchCreated", log); err != nil {
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

// ParseMatchCreated is a log parse operation binding the contract event 0x32619f811334dd1aff2e3cc0ed3b2697603d2439fcf63f47d1de5cc5971c8f89.
//
// Solidity: event MatchCreated((uint8,uint256,uint256,address,address,uint256,uint256) match_, bool useRelayer)
func (_Swapica *SwapicaFilterer) ParseMatchCreated(log types.Log) (*SwapicaMatchCreated, error) {
	event := new(SwapicaMatchCreated)
	if err := _Swapica.contract.UnpackLog(event, "MatchCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapicaMatchUpdatedIterator is returned from FilterMatchUpdated and is used to iterate over the raw logs and unpacked data for MatchUpdated events raised by the Swapica contract.
type SwapicaMatchUpdatedIterator struct {
	Event *SwapicaMatchUpdated // Event containing the contract specifics and raw log

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
func (it *SwapicaMatchUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapicaMatchUpdated)
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
		it.Event = new(SwapicaMatchUpdated)
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
func (it *SwapicaMatchUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapicaMatchUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapicaMatchUpdated represents a MatchUpdated event raised by the Swapica contract.
type SwapicaMatchUpdated struct {
	MatchId *big.Int
	Status  uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMatchUpdated is a free log retrieval operation binding the contract event 0xf194a47302430b75146eed29b4833f4c3c52e4c13f14dce86af88d6e1df1f20e.
//
// Solidity: event MatchUpdated(uint256 indexed matchId, uint8 status)
func (_Swapica *SwapicaFilterer) FilterMatchUpdated(opts *bind.FilterOpts, matchId []*big.Int) (*SwapicaMatchUpdatedIterator, error) {

	var matchIdRule []interface{}
	for _, matchIdItem := range matchId {
		matchIdRule = append(matchIdRule, matchIdItem)
	}

	logs, sub, err := _Swapica.contract.FilterLogs(opts, "MatchUpdated", matchIdRule)
	if err != nil {
		return nil, err
	}
	return &SwapicaMatchUpdatedIterator{contract: _Swapica.contract, event: "MatchUpdated", logs: logs, sub: sub}, nil
}

// WatchMatchUpdated is a free log subscription operation binding the contract event 0xf194a47302430b75146eed29b4833f4c3c52e4c13f14dce86af88d6e1df1f20e.
//
// Solidity: event MatchUpdated(uint256 indexed matchId, uint8 status)
func (_Swapica *SwapicaFilterer) WatchMatchUpdated(opts *bind.WatchOpts, sink chan<- *SwapicaMatchUpdated, matchId []*big.Int) (event.Subscription, error) {

	var matchIdRule []interface{}
	for _, matchIdItem := range matchId {
		matchIdRule = append(matchIdRule, matchIdItem)
	}

	logs, sub, err := _Swapica.contract.WatchLogs(opts, "MatchUpdated", matchIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapicaMatchUpdated)
				if err := _Swapica.contract.UnpackLog(event, "MatchUpdated", log); err != nil {
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

// ParseMatchUpdated is a log parse operation binding the contract event 0xf194a47302430b75146eed29b4833f4c3c52e4c13f14dce86af88d6e1df1f20e.
//
// Solidity: event MatchUpdated(uint256 indexed matchId, uint8 status)
func (_Swapica *SwapicaFilterer) ParseMatchUpdated(log types.Log) (*SwapicaMatchUpdated, error) {
	event := new(SwapicaMatchUpdated)
	if err := _Swapica.contract.UnpackLog(event, "MatchUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapicaOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the Swapica contract.
type SwapicaOrderCreatedIterator struct {
	Event *SwapicaOrderCreated // Event containing the contract specifics and raw log

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
func (it *SwapicaOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapicaOrderCreated)
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
		it.Event = new(SwapicaOrderCreated)
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
func (it *SwapicaOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapicaOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapicaOrderCreated represents a OrderCreated event raised by the Swapica contract.
type SwapicaOrderCreated struct {
	Order      ISwapicaOrder
	UseRelayer bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0x9488f6f0c9882f87541e717b5129ed27dd5cdd85e7dbfa256947795fda7cc276.
//
// Solidity: event OrderCreated(((uint8,uint256,address),uint256,address,address,uint256,address,uint256,uint256) order, bool useRelayer)
func (_Swapica *SwapicaFilterer) FilterOrderCreated(opts *bind.FilterOpts) (*SwapicaOrderCreatedIterator, error) {

	logs, sub, err := _Swapica.contract.FilterLogs(opts, "OrderCreated")
	if err != nil {
		return nil, err
	}
	return &SwapicaOrderCreatedIterator{contract: _Swapica.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x9488f6f0c9882f87541e717b5129ed27dd5cdd85e7dbfa256947795fda7cc276.
//
// Solidity: event OrderCreated(((uint8,uint256,address),uint256,address,address,uint256,address,uint256,uint256) order, bool useRelayer)
func (_Swapica *SwapicaFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *SwapicaOrderCreated) (event.Subscription, error) {

	logs, sub, err := _Swapica.contract.WatchLogs(opts, "OrderCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapicaOrderCreated)
				if err := _Swapica.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0x9488f6f0c9882f87541e717b5129ed27dd5cdd85e7dbfa256947795fda7cc276.
//
// Solidity: event OrderCreated(((uint8,uint256,address),uint256,address,address,uint256,address,uint256,uint256) order, bool useRelayer)
func (_Swapica *SwapicaFilterer) ParseOrderCreated(log types.Log) (*SwapicaOrderCreated, error) {
	event := new(SwapicaOrderCreated)
	if err := _Swapica.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapicaOrderUpdatedIterator is returned from FilterOrderUpdated and is used to iterate over the raw logs and unpacked data for OrderUpdated events raised by the Swapica contract.
type SwapicaOrderUpdatedIterator struct {
	Event *SwapicaOrderUpdated // Event containing the contract specifics and raw log

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
func (it *SwapicaOrderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapicaOrderUpdated)
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
		it.Event = new(SwapicaOrderUpdated)
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
func (it *SwapicaOrderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapicaOrderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapicaOrderUpdated represents a OrderUpdated event raised by the Swapica contract.
type SwapicaOrderUpdated struct {
	OrderId *big.Int
	Status  ISwapicaOrderStatus
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderUpdated is a free log retrieval operation binding the contract event 0x4c59495d370247039622ce9df1b0843542a363a1b9055118689976b177709635.
//
// Solidity: event OrderUpdated(uint256 indexed orderId, (uint8,uint256,address) status)
func (_Swapica *SwapicaFilterer) FilterOrderUpdated(opts *bind.FilterOpts, orderId []*big.Int) (*SwapicaOrderUpdatedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Swapica.contract.FilterLogs(opts, "OrderUpdated", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &SwapicaOrderUpdatedIterator{contract: _Swapica.contract, event: "OrderUpdated", logs: logs, sub: sub}, nil
}

// WatchOrderUpdated is a free log subscription operation binding the contract event 0x4c59495d370247039622ce9df1b0843542a363a1b9055118689976b177709635.
//
// Solidity: event OrderUpdated(uint256 indexed orderId, (uint8,uint256,address) status)
func (_Swapica *SwapicaFilterer) WatchOrderUpdated(opts *bind.WatchOpts, sink chan<- *SwapicaOrderUpdated, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Swapica.contract.WatchLogs(opts, "OrderUpdated", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapicaOrderUpdated)
				if err := _Swapica.contract.UnpackLog(event, "OrderUpdated", log); err != nil {
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

// ParseOrderUpdated is a log parse operation binding the contract event 0x4c59495d370247039622ce9df1b0843542a363a1b9055118689976b177709635.
//
// Solidity: event OrderUpdated(uint256 indexed orderId, (uint8,uint256,address) status)
func (_Swapica *SwapicaFilterer) ParseOrderUpdated(log types.Log) (*SwapicaOrderUpdated, error) {
	event := new(SwapicaOrderUpdated)
	if err := _Swapica.contract.UnpackLog(event, "OrderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapicaOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Swapica contract.
type SwapicaOwnershipTransferredIterator struct {
	Event *SwapicaOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SwapicaOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapicaOwnershipTransferred)
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
		it.Event = new(SwapicaOwnershipTransferred)
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
func (it *SwapicaOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapicaOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapicaOwnershipTransferred represents a OwnershipTransferred event raised by the Swapica contract.
type SwapicaOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swapica *SwapicaFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SwapicaOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Swapica.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwapicaOwnershipTransferredIterator{contract: _Swapica.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swapica *SwapicaFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SwapicaOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Swapica.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapicaOwnershipTransferred)
				if err := _Swapica.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Swapica *SwapicaFilterer) ParseOwnershipTransferred(log types.Log) (*SwapicaOwnershipTransferred, error) {
	event := new(SwapicaOwnershipTransferred)
	if err := _Swapica.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapicaUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Swapica contract.
type SwapicaUpgradedIterator struct {
	Event *SwapicaUpgraded // Event containing the contract specifics and raw log

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
func (it *SwapicaUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapicaUpgraded)
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
		it.Event = new(SwapicaUpgraded)
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
func (it *SwapicaUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapicaUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapicaUpgraded represents a Upgraded event raised by the Swapica contract.
type SwapicaUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Swapica *SwapicaFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SwapicaUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Swapica.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SwapicaUpgradedIterator{contract: _Swapica.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Swapica *SwapicaFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SwapicaUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Swapica.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapicaUpgraded)
				if err := _Swapica.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Swapica *SwapicaFilterer) ParseUpgraded(log types.Log) (*SwapicaUpgraded, error) {
	event := new(SwapicaUpgraded)
	if err := _Swapica.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
