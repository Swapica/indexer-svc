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

// SwapicaMatch is an auto generated low-level Go binding around an user-defined struct.
type SwapicaMatch struct {
	Id            *big.Int
	OriginOrderId *big.Int
	Account       common.Address
	TokenToSell   common.Address
	AmountToSell  *big.Int
	OriginChain   *big.Int
}

// SwapicaOrder is an auto generated low-level Go binding around an user-defined struct.
type SwapicaOrder struct {
	Id           *big.Int
	Account      common.Address
	TokenToSell  common.Address
	TokenToBuy   common.Address
	AmountToSell *big.Int
	AmountToBuy  *big.Int
	DestChain    *big.Int
}

// SwapicaStatus is an auto generated low-level Go binding around an user-defined struct.
type SwapicaStatus struct {
	State        uint8
	ExecutedBy   *big.Int
	MatchSwapica common.Address
}

// SwapicaMetaData contains all meta data concerning the Swapica contract.
var SwapicaMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumSwapica.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"executedBy\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"matchSwapica\",\"type\":\"address\"}],\"indexed\":true,\"internalType\":\"structSwapica.Status\",\"name\":\"status\",\"type\":\"tuple\"}],\"name\":\"MatchUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumSwapica.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"executedBy\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"matchSwapica\",\"type\":\"address\"}],\"indexed\":true,\"internalType\":\"structSwapica.Status\",\"name\":\"status\",\"type\":\"tuple\"}],\"name\":\"OrderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NATIVE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers_\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"signaturesThreshold_\",\"type\":\"uint256\"}],\"name\":\"__Signers_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"}],\"name\":\"__Swapica_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers_\",\"type\":\"address[]\"}],\"name\":\"addSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"orderData\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"cancelMatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"orderData\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"createMatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenToSell\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToSell\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenToBuy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChain\",\"type\":\"uint256\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"orderData\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"executeMatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"orderData\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"executeOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenToSell\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenToBuy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getActiveOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenToSell\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenToBuy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToSell\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChain\",\"type\":\"uint256\"}],\"internalType\":\"structSwapica.Order[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrdersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumSwapica.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getUserMatches\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originOrderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenToSell\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToSell\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originChain\",\"type\":\"uint256\"}],\"internalType\":\"structSwapica.Match[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserMatchesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumSwapica.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getUserOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenToSell\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenToBuy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToSell\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChain\",\"type\":\"uint256\"}],\"internalType\":\"structSwapica.Order[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserOrdersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"matchStatus\",\"outputs\":[{\"internalType\":\"enumSwapica.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"executedBy\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"matchSwapica\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"matches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originOrderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenToSell\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToSell\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originChain\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orderStatus\",\"outputs\":[{\"internalType\":\"enumSwapica.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"executedBy\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"matchSwapica\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenToSell\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenToBuy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToSell\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChain\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers_\",\"type\":\"address[]\"}],\"name\":\"removeSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"signaturesThreshold_\",\"type\":\"uint256\"}],\"name\":\"setSignaturesThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signaturesThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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

// NATIVE is a free data retrieval call binding the contract method 0xa0cf0aea.
//
// Solidity: function NATIVE() view returns(address)
func (_Swapica *SwapicaCaller) NATIVE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "NATIVE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVE is a free data retrieval call binding the contract method 0xa0cf0aea.
//
// Solidity: function NATIVE() view returns(address)
func (_Swapica *SwapicaSession) NATIVE() (common.Address, error) {
	return _Swapica.Contract.NATIVE(&_Swapica.CallOpts)
}

// NATIVE is a free data retrieval call binding the contract method 0xa0cf0aea.
//
// Solidity: function NATIVE() view returns(address)
func (_Swapica *SwapicaCallerSession) NATIVE() (common.Address, error) {
	return _Swapica.Contract.NATIVE(&_Swapica.CallOpts)
}

// GetActiveOrders is a free data retrieval call binding the contract method 0xab9f52a1.
//
// Solidity: function getActiveOrders(address tokenToSell, address tokenToBuy, uint256 begin, uint256 end) view returns((uint256,address,address,address,uint256,uint256,uint256)[] result)
func (_Swapica *SwapicaCaller) GetActiveOrders(opts *bind.CallOpts, tokenToSell common.Address, tokenToBuy common.Address, begin *big.Int, end *big.Int) ([]SwapicaOrder, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "getActiveOrders", tokenToSell, tokenToBuy, begin, end)

	if err != nil {
		return *new([]SwapicaOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]SwapicaOrder)).(*[]SwapicaOrder)

	return out0, err

}

// GetActiveOrders is a free data retrieval call binding the contract method 0xab9f52a1.
//
// Solidity: function getActiveOrders(address tokenToSell, address tokenToBuy, uint256 begin, uint256 end) view returns((uint256,address,address,address,uint256,uint256,uint256)[] result)
func (_Swapica *SwapicaSession) GetActiveOrders(tokenToSell common.Address, tokenToBuy common.Address, begin *big.Int, end *big.Int) ([]SwapicaOrder, error) {
	return _Swapica.Contract.GetActiveOrders(&_Swapica.CallOpts, tokenToSell, tokenToBuy, begin, end)
}

// GetActiveOrders is a free data retrieval call binding the contract method 0xab9f52a1.
//
// Solidity: function getActiveOrders(address tokenToSell, address tokenToBuy, uint256 begin, uint256 end) view returns((uint256,address,address,address,uint256,uint256,uint256)[] result)
func (_Swapica *SwapicaCallerSession) GetActiveOrders(tokenToSell common.Address, tokenToBuy common.Address, begin *big.Int, end *big.Int) ([]SwapicaOrder, error) {
	return _Swapica.Contract.GetActiveOrders(&_Swapica.CallOpts, tokenToSell, tokenToBuy, begin, end)
}

// GetOrdersLength is a free data retrieval call binding the contract method 0x07c2e16c.
//
// Solidity: function getOrdersLength() view returns(uint256)
func (_Swapica *SwapicaCaller) GetOrdersLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "getOrdersLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrdersLength is a free data retrieval call binding the contract method 0x07c2e16c.
//
// Solidity: function getOrdersLength() view returns(uint256)
func (_Swapica *SwapicaSession) GetOrdersLength() (*big.Int, error) {
	return _Swapica.Contract.GetOrdersLength(&_Swapica.CallOpts)
}

// GetOrdersLength is a free data retrieval call binding the contract method 0x07c2e16c.
//
// Solidity: function getOrdersLength() view returns(uint256)
func (_Swapica *SwapicaCallerSession) GetOrdersLength() (*big.Int, error) {
	return _Swapica.Contract.GetOrdersLength(&_Swapica.CallOpts)
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

// GetUserMatches is a free data retrieval call binding the contract method 0xa435940f.
//
// Solidity: function getUserMatches(address user, uint8 state, uint256 begin, uint256 end) view returns((uint256,uint256,address,address,uint256,uint256)[] result)
func (_Swapica *SwapicaCaller) GetUserMatches(opts *bind.CallOpts, user common.Address, state uint8, begin *big.Int, end *big.Int) ([]SwapicaMatch, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "getUserMatches", user, state, begin, end)

	if err != nil {
		return *new([]SwapicaMatch), err
	}

	out0 := *abi.ConvertType(out[0], new([]SwapicaMatch)).(*[]SwapicaMatch)

	return out0, err

}

// GetUserMatches is a free data retrieval call binding the contract method 0xa435940f.
//
// Solidity: function getUserMatches(address user, uint8 state, uint256 begin, uint256 end) view returns((uint256,uint256,address,address,uint256,uint256)[] result)
func (_Swapica *SwapicaSession) GetUserMatches(user common.Address, state uint8, begin *big.Int, end *big.Int) ([]SwapicaMatch, error) {
	return _Swapica.Contract.GetUserMatches(&_Swapica.CallOpts, user, state, begin, end)
}

// GetUserMatches is a free data retrieval call binding the contract method 0xa435940f.
//
// Solidity: function getUserMatches(address user, uint8 state, uint256 begin, uint256 end) view returns((uint256,uint256,address,address,uint256,uint256)[] result)
func (_Swapica *SwapicaCallerSession) GetUserMatches(user common.Address, state uint8, begin *big.Int, end *big.Int) ([]SwapicaMatch, error) {
	return _Swapica.Contract.GetUserMatches(&_Swapica.CallOpts, user, state, begin, end)
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

// GetUserOrders is a free data retrieval call binding the contract method 0xc7efc593.
//
// Solidity: function getUserOrders(address user, uint8 state, uint256 begin, uint256 end) view returns((uint256,address,address,address,uint256,uint256,uint256)[] result)
func (_Swapica *SwapicaCaller) GetUserOrders(opts *bind.CallOpts, user common.Address, state uint8, begin *big.Int, end *big.Int) ([]SwapicaOrder, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "getUserOrders", user, state, begin, end)

	if err != nil {
		return *new([]SwapicaOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]SwapicaOrder)).(*[]SwapicaOrder)

	return out0, err

}

// GetUserOrders is a free data retrieval call binding the contract method 0xc7efc593.
//
// Solidity: function getUserOrders(address user, uint8 state, uint256 begin, uint256 end) view returns((uint256,address,address,address,uint256,uint256,uint256)[] result)
func (_Swapica *SwapicaSession) GetUserOrders(user common.Address, state uint8, begin *big.Int, end *big.Int) ([]SwapicaOrder, error) {
	return _Swapica.Contract.GetUserOrders(&_Swapica.CallOpts, user, state, begin, end)
}

// GetUserOrders is a free data retrieval call binding the contract method 0xc7efc593.
//
// Solidity: function getUserOrders(address user, uint8 state, uint256 begin, uint256 end) view returns((uint256,address,address,address,uint256,uint256,uint256)[] result)
func (_Swapica *SwapicaCallerSession) GetUserOrders(user common.Address, state uint8, begin *big.Int, end *big.Int) ([]SwapicaOrder, error) {
	return _Swapica.Contract.GetUserOrders(&_Swapica.CallOpts, user, state, begin, end)
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

// Locked is a free data retrieval call binding the contract method 0xdb20266f.
//
// Solidity: function locked(address , address ) view returns(uint256)
func (_Swapica *SwapicaCaller) Locked(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "locked", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xdb20266f.
//
// Solidity: function locked(address , address ) view returns(uint256)
func (_Swapica *SwapicaSession) Locked(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Swapica.Contract.Locked(&_Swapica.CallOpts, arg0, arg1)
}

// Locked is a free data retrieval call binding the contract method 0xdb20266f.
//
// Solidity: function locked(address , address ) view returns(uint256)
func (_Swapica *SwapicaCallerSession) Locked(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Swapica.Contract.Locked(&_Swapica.CallOpts, arg0, arg1)
}

// MatchStatus is a free data retrieval call binding the contract method 0x23d38621.
//
// Solidity: function matchStatus(uint256 ) view returns(uint8 state, uint256 executedBy, address matchSwapica)
func (_Swapica *SwapicaCaller) MatchStatus(opts *bind.CallOpts, arg0 *big.Int) (struct {
	State        uint8
	ExecutedBy   *big.Int
	MatchSwapica common.Address
}, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "matchStatus", arg0)

	outstruct := new(struct {
		State        uint8
		ExecutedBy   *big.Int
		MatchSwapica common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.State = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.ExecutedBy = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MatchSwapica = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// MatchStatus is a free data retrieval call binding the contract method 0x23d38621.
//
// Solidity: function matchStatus(uint256 ) view returns(uint8 state, uint256 executedBy, address matchSwapica)
func (_Swapica *SwapicaSession) MatchStatus(arg0 *big.Int) (struct {
	State        uint8
	ExecutedBy   *big.Int
	MatchSwapica common.Address
}, error) {
	return _Swapica.Contract.MatchStatus(&_Swapica.CallOpts, arg0)
}

// MatchStatus is a free data retrieval call binding the contract method 0x23d38621.
//
// Solidity: function matchStatus(uint256 ) view returns(uint8 state, uint256 executedBy, address matchSwapica)
func (_Swapica *SwapicaCallerSession) MatchStatus(arg0 *big.Int) (struct {
	State        uint8
	ExecutedBy   *big.Int
	MatchSwapica common.Address
}, error) {
	return _Swapica.Contract.MatchStatus(&_Swapica.CallOpts, arg0)
}

// Matches is a free data retrieval call binding the contract method 0x4768d4ef.
//
// Solidity: function matches(uint256 ) view returns(uint256 id, uint256 originOrderId, address account, address tokenToSell, uint256 amountToSell, uint256 originChain)
func (_Swapica *SwapicaCaller) Matches(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id            *big.Int
	OriginOrderId *big.Int
	Account       common.Address
	TokenToSell   common.Address
	AmountToSell  *big.Int
	OriginChain   *big.Int
}, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "matches", arg0)

	outstruct := new(struct {
		Id            *big.Int
		OriginOrderId *big.Int
		Account       common.Address
		TokenToSell   common.Address
		AmountToSell  *big.Int
		OriginChain   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OriginOrderId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Account = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TokenToSell = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.AmountToSell = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.OriginChain = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Matches is a free data retrieval call binding the contract method 0x4768d4ef.
//
// Solidity: function matches(uint256 ) view returns(uint256 id, uint256 originOrderId, address account, address tokenToSell, uint256 amountToSell, uint256 originChain)
func (_Swapica *SwapicaSession) Matches(arg0 *big.Int) (struct {
	Id            *big.Int
	OriginOrderId *big.Int
	Account       common.Address
	TokenToSell   common.Address
	AmountToSell  *big.Int
	OriginChain   *big.Int
}, error) {
	return _Swapica.Contract.Matches(&_Swapica.CallOpts, arg0)
}

// Matches is a free data retrieval call binding the contract method 0x4768d4ef.
//
// Solidity: function matches(uint256 ) view returns(uint256 id, uint256 originOrderId, address account, address tokenToSell, uint256 amountToSell, uint256 originChain)
func (_Swapica *SwapicaCallerSession) Matches(arg0 *big.Int) (struct {
	Id            *big.Int
	OriginOrderId *big.Int
	Account       common.Address
	TokenToSell   common.Address
	AmountToSell  *big.Int
	OriginChain   *big.Int
}, error) {
	return _Swapica.Contract.Matches(&_Swapica.CallOpts, arg0)
}

// OrderStatus is a free data retrieval call binding the contract method 0xbff49450.
//
// Solidity: function orderStatus(uint256 ) view returns(uint8 state, uint256 executedBy, address matchSwapica)
func (_Swapica *SwapicaCaller) OrderStatus(opts *bind.CallOpts, arg0 *big.Int) (struct {
	State        uint8
	ExecutedBy   *big.Int
	MatchSwapica common.Address
}, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "orderStatus", arg0)

	outstruct := new(struct {
		State        uint8
		ExecutedBy   *big.Int
		MatchSwapica common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.State = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.ExecutedBy = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MatchSwapica = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// OrderStatus is a free data retrieval call binding the contract method 0xbff49450.
//
// Solidity: function orderStatus(uint256 ) view returns(uint8 state, uint256 executedBy, address matchSwapica)
func (_Swapica *SwapicaSession) OrderStatus(arg0 *big.Int) (struct {
	State        uint8
	ExecutedBy   *big.Int
	MatchSwapica common.Address
}, error) {
	return _Swapica.Contract.OrderStatus(&_Swapica.CallOpts, arg0)
}

// OrderStatus is a free data retrieval call binding the contract method 0xbff49450.
//
// Solidity: function orderStatus(uint256 ) view returns(uint8 state, uint256 executedBy, address matchSwapica)
func (_Swapica *SwapicaCallerSession) OrderStatus(arg0 *big.Int) (struct {
	State        uint8
	ExecutedBy   *big.Int
	MatchSwapica common.Address
}, error) {
	return _Swapica.Contract.OrderStatus(&_Swapica.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(uint256 id, address account, address tokenToSell, address tokenToBuy, uint256 amountToSell, uint256 amountToBuy, uint256 destChain)
func (_Swapica *SwapicaCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id           *big.Int
	Account      common.Address
	TokenToSell  common.Address
	TokenToBuy   common.Address
	AmountToSell *big.Int
	AmountToBuy  *big.Int
	DestChain    *big.Int
}, error) {
	var out []interface{}
	err := _Swapica.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		Id           *big.Int
		Account      common.Address
		TokenToSell  common.Address
		TokenToBuy   common.Address
		AmountToSell *big.Int
		AmountToBuy  *big.Int
		DestChain    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Account = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenToSell = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TokenToBuy = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.AmountToSell = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AmountToBuy = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.DestChain = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(uint256 id, address account, address tokenToSell, address tokenToBuy, uint256 amountToSell, uint256 amountToBuy, uint256 destChain)
func (_Swapica *SwapicaSession) Orders(arg0 *big.Int) (struct {
	Id           *big.Int
	Account      common.Address
	TokenToSell  common.Address
	TokenToBuy   common.Address
	AmountToSell *big.Int
	AmountToBuy  *big.Int
	DestChain    *big.Int
}, error) {
	return _Swapica.Contract.Orders(&_Swapica.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(uint256 id, address account, address tokenToSell, address tokenToBuy, uint256 amountToSell, uint256 amountToBuy, uint256 destChain)
func (_Swapica *SwapicaCallerSession) Orders(arg0 *big.Int) (struct {
	Id           *big.Int
	Account      common.Address
	TokenToSell  common.Address
	TokenToBuy   common.Address
	AmountToSell *big.Int
	AmountToBuy  *big.Int
	DestChain    *big.Int
}, error) {
	return _Swapica.Contract.Orders(&_Swapica.CallOpts, arg0)
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
// Solidity: function __Signers_init(address[] signers_, uint256 signaturesThreshold_) returns()
func (_Swapica *SwapicaTransactor) SignersInit(opts *bind.TransactOpts, signers_ []common.Address, signaturesThreshold_ *big.Int) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "__Signers_init", signers_, signaturesThreshold_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x09a55841.
//
// Solidity: function __Signers_init(address[] signers_, uint256 signaturesThreshold_) returns()
func (_Swapica *SwapicaSession) SignersInit(signers_ []common.Address, signaturesThreshold_ *big.Int) (*types.Transaction, error) {
	return _Swapica.Contract.SignersInit(&_Swapica.TransactOpts, signers_, signaturesThreshold_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x09a55841.
//
// Solidity: function __Signers_init(address[] signers_, uint256 signaturesThreshold_) returns()
func (_Swapica *SwapicaTransactorSession) SignersInit(signers_ []common.Address, signaturesThreshold_ *big.Int) (*types.Transaction, error) {
	return _Swapica.Contract.SignersInit(&_Swapica.TransactOpts, signers_, signaturesThreshold_)
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
// Solidity: function addSigners(address[] signers_) returns()
func (_Swapica *SwapicaTransactor) AddSigners(opts *bind.TransactOpts, signers_ []common.Address) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "addSigners", signers_)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] signers_) returns()
func (_Swapica *SwapicaSession) AddSigners(signers_ []common.Address) (*types.Transaction, error) {
	return _Swapica.Contract.AddSigners(&_Swapica.TransactOpts, signers_)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] signers_) returns()
func (_Swapica *SwapicaTransactorSession) AddSigners(signers_ []common.Address) (*types.Transaction, error) {
	return _Swapica.Contract.AddSigners(&_Swapica.TransactOpts, signers_)
}

// CancelMatch is a paid mutator transaction binding the contract method 0xb925c02d.
//
// Solidity: function cancelMatch(bytes orderData, bytes[] signatures) returns()
func (_Swapica *SwapicaTransactor) CancelMatch(opts *bind.TransactOpts, orderData []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "cancelMatch", orderData, signatures)
}

// CancelMatch is a paid mutator transaction binding the contract method 0xb925c02d.
//
// Solidity: function cancelMatch(bytes orderData, bytes[] signatures) returns()
func (_Swapica *SwapicaSession) CancelMatch(orderData []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.Contract.CancelMatch(&_Swapica.TransactOpts, orderData, signatures)
}

// CancelMatch is a paid mutator transaction binding the contract method 0xb925c02d.
//
// Solidity: function cancelMatch(bytes orderData, bytes[] signatures) returns()
func (_Swapica *SwapicaTransactorSession) CancelMatch(orderData []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.Contract.CancelMatch(&_Swapica.TransactOpts, orderData, signatures)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 id) returns()
func (_Swapica *SwapicaTransactor) CancelOrder(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "cancelOrder", id)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 id) returns()
func (_Swapica *SwapicaSession) CancelOrder(id *big.Int) (*types.Transaction, error) {
	return _Swapica.Contract.CancelOrder(&_Swapica.TransactOpts, id)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 id) returns()
func (_Swapica *SwapicaTransactorSession) CancelOrder(id *big.Int) (*types.Transaction, error) {
	return _Swapica.Contract.CancelOrder(&_Swapica.TransactOpts, id)
}

// CreateMatch is a paid mutator transaction binding the contract method 0xf7662788.
//
// Solidity: function createMatch(bytes orderData, bytes[] signatures) payable returns()
func (_Swapica *SwapicaTransactor) CreateMatch(opts *bind.TransactOpts, orderData []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "createMatch", orderData, signatures)
}

// CreateMatch is a paid mutator transaction binding the contract method 0xf7662788.
//
// Solidity: function createMatch(bytes orderData, bytes[] signatures) payable returns()
func (_Swapica *SwapicaSession) CreateMatch(orderData []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.Contract.CreateMatch(&_Swapica.TransactOpts, orderData, signatures)
}

// CreateMatch is a paid mutator transaction binding the contract method 0xf7662788.
//
// Solidity: function createMatch(bytes orderData, bytes[] signatures) payable returns()
func (_Swapica *SwapicaTransactorSession) CreateMatch(orderData []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.Contract.CreateMatch(&_Swapica.TransactOpts, orderData, signatures)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x96d4f640.
//
// Solidity: function createOrder(address tokenToSell, uint256 amountToSell, address tokenToBuy, uint256 amountToBuy, uint256 destChain) payable returns()
func (_Swapica *SwapicaTransactor) CreateOrder(opts *bind.TransactOpts, tokenToSell common.Address, amountToSell *big.Int, tokenToBuy common.Address, amountToBuy *big.Int, destChain *big.Int) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "createOrder", tokenToSell, amountToSell, tokenToBuy, amountToBuy, destChain)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x96d4f640.
//
// Solidity: function createOrder(address tokenToSell, uint256 amountToSell, address tokenToBuy, uint256 amountToBuy, uint256 destChain) payable returns()
func (_Swapica *SwapicaSession) CreateOrder(tokenToSell common.Address, amountToSell *big.Int, tokenToBuy common.Address, amountToBuy *big.Int, destChain *big.Int) (*types.Transaction, error) {
	return _Swapica.Contract.CreateOrder(&_Swapica.TransactOpts, tokenToSell, amountToSell, tokenToBuy, amountToBuy, destChain)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x96d4f640.
//
// Solidity: function createOrder(address tokenToSell, uint256 amountToSell, address tokenToBuy, uint256 amountToBuy, uint256 destChain) payable returns()
func (_Swapica *SwapicaTransactorSession) CreateOrder(tokenToSell common.Address, amountToSell *big.Int, tokenToBuy common.Address, amountToBuy *big.Int, destChain *big.Int) (*types.Transaction, error) {
	return _Swapica.Contract.CreateOrder(&_Swapica.TransactOpts, tokenToSell, amountToSell, tokenToBuy, amountToBuy, destChain)
}

// ExecuteMatch is a paid mutator transaction binding the contract method 0x2a9ff0d5.
//
// Solidity: function executeMatch(bytes orderData, bytes[] signatures) returns()
func (_Swapica *SwapicaTransactor) ExecuteMatch(opts *bind.TransactOpts, orderData []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "executeMatch", orderData, signatures)
}

// ExecuteMatch is a paid mutator transaction binding the contract method 0x2a9ff0d5.
//
// Solidity: function executeMatch(bytes orderData, bytes[] signatures) returns()
func (_Swapica *SwapicaSession) ExecuteMatch(orderData []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.Contract.ExecuteMatch(&_Swapica.TransactOpts, orderData, signatures)
}

// ExecuteMatch is a paid mutator transaction binding the contract method 0x2a9ff0d5.
//
// Solidity: function executeMatch(bytes orderData, bytes[] signatures) returns()
func (_Swapica *SwapicaTransactorSession) ExecuteMatch(orderData []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.Contract.ExecuteMatch(&_Swapica.TransactOpts, orderData, signatures)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x14d41b45.
//
// Solidity: function executeOrder(bytes orderData, bytes[] signatures) returns()
func (_Swapica *SwapicaTransactor) ExecuteOrder(opts *bind.TransactOpts, orderData []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "executeOrder", orderData, signatures)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x14d41b45.
//
// Solidity: function executeOrder(bytes orderData, bytes[] signatures) returns()
func (_Swapica *SwapicaSession) ExecuteOrder(orderData []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.Contract.ExecuteOrder(&_Swapica.TransactOpts, orderData, signatures)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x14d41b45.
//
// Solidity: function executeOrder(bytes orderData, bytes[] signatures) returns()
func (_Swapica *SwapicaTransactorSession) ExecuteOrder(orderData []byte, signatures [][]byte) (*types.Transaction, error) {
	return _Swapica.Contract.ExecuteOrder(&_Swapica.TransactOpts, orderData, signatures)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x8d361e43.
//
// Solidity: function removeSigners(address[] signers_) returns()
func (_Swapica *SwapicaTransactor) RemoveSigners(opts *bind.TransactOpts, signers_ []common.Address) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "removeSigners", signers_)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x8d361e43.
//
// Solidity: function removeSigners(address[] signers_) returns()
func (_Swapica *SwapicaSession) RemoveSigners(signers_ []common.Address) (*types.Transaction, error) {
	return _Swapica.Contract.RemoveSigners(&_Swapica.TransactOpts, signers_)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x8d361e43.
//
// Solidity: function removeSigners(address[] signers_) returns()
func (_Swapica *SwapicaTransactorSession) RemoveSigners(signers_ []common.Address) (*types.Transaction, error) {
	return _Swapica.Contract.RemoveSigners(&_Swapica.TransactOpts, signers_)
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
// Solidity: function setSignaturesThreshold(uint256 signaturesThreshold_) returns()
func (_Swapica *SwapicaTransactor) SetSignaturesThreshold(opts *bind.TransactOpts, signaturesThreshold_ *big.Int) (*types.Transaction, error) {
	return _Swapica.contract.Transact(opts, "setSignaturesThreshold", signaturesThreshold_)
}

// SetSignaturesThreshold is a paid mutator transaction binding the contract method 0xbf1fe08f.
//
// Solidity: function setSignaturesThreshold(uint256 signaturesThreshold_) returns()
func (_Swapica *SwapicaSession) SetSignaturesThreshold(signaturesThreshold_ *big.Int) (*types.Transaction, error) {
	return _Swapica.Contract.SetSignaturesThreshold(&_Swapica.TransactOpts, signaturesThreshold_)
}

// SetSignaturesThreshold is a paid mutator transaction binding the contract method 0xbf1fe08f.
//
// Solidity: function setSignaturesThreshold(uint256 signaturesThreshold_) returns()
func (_Swapica *SwapicaTransactorSession) SetSignaturesThreshold(signaturesThreshold_ *big.Int) (*types.Transaction, error) {
	return _Swapica.Contract.SetSignaturesThreshold(&_Swapica.TransactOpts, signaturesThreshold_)
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
	Id     *big.Int
	Status SwapicaStatus
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMatchUpdated is a free log retrieval operation binding the contract event 0xa9b5f1914efdf39593d3001993d7ebde53d6c14bf4ea8b5af4c545f81bf47120.
//
// Solidity: event MatchUpdated(uint256 indexed id, (uint8,uint256,address) indexed status)
func (_Swapica *SwapicaFilterer) FilterMatchUpdated(opts *bind.FilterOpts, id []*big.Int, status []SwapicaStatus) (*SwapicaMatchUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Swapica.contract.FilterLogs(opts, "MatchUpdated", idRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &SwapicaMatchUpdatedIterator{contract: _Swapica.contract, event: "MatchUpdated", logs: logs, sub: sub}, nil
}

// WatchMatchUpdated is a free log subscription operation binding the contract event 0xa9b5f1914efdf39593d3001993d7ebde53d6c14bf4ea8b5af4c545f81bf47120.
//
// Solidity: event MatchUpdated(uint256 indexed id, (uint8,uint256,address) indexed status)
func (_Swapica *SwapicaFilterer) WatchMatchUpdated(opts *bind.WatchOpts, sink chan<- *SwapicaMatchUpdated, id []*big.Int, status []SwapicaStatus) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Swapica.contract.WatchLogs(opts, "MatchUpdated", idRule, statusRule)
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

// ParseMatchUpdated is a log parse operation binding the contract event 0xa9b5f1914efdf39593d3001993d7ebde53d6c14bf4ea8b5af4c545f81bf47120.
//
// Solidity: event MatchUpdated(uint256 indexed id, (uint8,uint256,address) indexed status)
func (_Swapica *SwapicaFilterer) ParseMatchUpdated(log types.Log) (*SwapicaMatchUpdated, error) {
	event := new(SwapicaMatchUpdated)
	if err := _Swapica.contract.UnpackLog(event, "MatchUpdated", log); err != nil {
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
	Id     *big.Int
	Status SwapicaStatus
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOrderUpdated is a free log retrieval operation binding the contract event 0x4c59495d370247039622ce9df1b0843542a363a1b9055118689976b177709635.
//
// Solidity: event OrderUpdated(uint256 indexed id, (uint8,uint256,address) indexed status)
func (_Swapica *SwapicaFilterer) FilterOrderUpdated(opts *bind.FilterOpts, id []*big.Int, status []SwapicaStatus) (*SwapicaOrderUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Swapica.contract.FilterLogs(opts, "OrderUpdated", idRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &SwapicaOrderUpdatedIterator{contract: _Swapica.contract, event: "OrderUpdated", logs: logs, sub: sub}, nil
}

// WatchOrderUpdated is a free log subscription operation binding the contract event 0x4c59495d370247039622ce9df1b0843542a363a1b9055118689976b177709635.
//
// Solidity: event OrderUpdated(uint256 indexed id, (uint8,uint256,address) indexed status)
func (_Swapica *SwapicaFilterer) WatchOrderUpdated(opts *bind.WatchOpts, sink chan<- *SwapicaOrderUpdated, id []*big.Int, status []SwapicaStatus) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Swapica.contract.WatchLogs(opts, "OrderUpdated", idRule, statusRule)
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
// Solidity: event OrderUpdated(uint256 indexed id, (uint8,uint256,address) indexed status)
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
