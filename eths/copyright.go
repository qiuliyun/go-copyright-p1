// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eths

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

// EthsABI is the input ABI used to generate the binding from.
const EthsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_approved\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tran\",\"type\":\"string\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"assRecode\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recode\",\"type\":\"string\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"copyRecode\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"newAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assets\",\"outputs\":[{\"name\":\"contentHash\",\"type\":\"bytes32\"},{\"name\":\"copyrightTran\",\"type\":\"string\"},{\"name\":\"tran\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAssRecode\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getCopyrecode\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Eths is an auto generated Go binding around an Ethereum contract.
type Eths struct {
	EthsCaller     // Read-only binding to the contract
	EthsTransactor // Write-only binding to the contract
	EthsFilterer   // Log filterer for contract events
}

// EthsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthsSession struct {
	Contract     *Eths             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthsCallerSession struct {
	Contract *EthsCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EthsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthsTransactorSession struct {
	Contract     *EthsTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthsRaw struct {
	Contract *Eths // Generic contract binding to access the raw methods on
}

// EthsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthsCallerRaw struct {
	Contract *EthsCaller // Generic read-only contract binding to access the raw methods on
}

// EthsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthsTransactorRaw struct {
	Contract *EthsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEths creates a new instance of Eths, bound to a specific deployed contract.
func NewEths(address common.Address, backend bind.ContractBackend) (*Eths, error) {
	contract, err := bindEths(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eths{EthsCaller: EthsCaller{contract: contract}, EthsTransactor: EthsTransactor{contract: contract}, EthsFilterer: EthsFilterer{contract: contract}}, nil
}

// NewEthsCaller creates a new read-only instance of Eths, bound to a specific deployed contract.
func NewEthsCaller(address common.Address, caller bind.ContractCaller) (*EthsCaller, error) {
	contract, err := bindEths(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthsCaller{contract: contract}, nil
}

// NewEthsTransactor creates a new write-only instance of Eths, bound to a specific deployed contract.
func NewEthsTransactor(address common.Address, transactor bind.ContractTransactor) (*EthsTransactor, error) {
	contract, err := bindEths(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthsTransactor{contract: contract}, nil
}

// NewEthsFilterer creates a new log filterer instance of Eths, bound to a specific deployed contract.
func NewEthsFilterer(address common.Address, filterer bind.ContractFilterer) (*EthsFilterer, error) {
	contract, err := bindEths(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthsFilterer{contract: contract}, nil
}

// bindEths binds a generic wrapper to an already deployed contract.
func bindEths(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eths *EthsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Eths.Contract.EthsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eths *EthsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eths.Contract.EthsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eths *EthsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eths.Contract.EthsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eths *EthsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Eths.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eths *EthsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eths.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eths *EthsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eths.Contract.contract.Transact(opts, method, params...)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) constant returns(bytes32 contentHash, string copyrightTran, string tran, string name)
func (_Eths *EthsCaller) Assets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ContentHash   [32]byte
	CopyrightTran string
	Tran          string
	Name          string
}, error) {
	ret := new(struct {
		ContentHash   [32]byte
		CopyrightTran string
		Tran          string
		Name          string
	})
	out := ret
	err := _Eths.contract.Call(opts, out, "assets", arg0)
	return *ret, err
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) constant returns(bytes32 contentHash, string copyrightTran, string tran, string name)
func (_Eths *EthsSession) Assets(arg0 *big.Int) (struct {
	ContentHash   [32]byte
	CopyrightTran string
	Tran          string
	Name          string
}, error) {
	return _Eths.Contract.Assets(&_Eths.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) constant returns(bytes32 contentHash, string copyrightTran, string tran, string name)
func (_Eths *EthsCallerSession) Assets(arg0 *big.Int) (struct {
	ContentHash   [32]byte
	CopyrightTran string
	Tran          string
	Name          string
}, error) {
	return _Eths.Contract.Assets(&_Eths.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Eths *EthsCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Eths.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Eths *EthsSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Eths.Contract.BalanceOf(&_Eths.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_Eths *EthsCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Eths.Contract.BalanceOf(&_Eths.CallOpts, _owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address)
func (_Eths *EthsCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Eths.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address)
func (_Eths *EthsSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Eths.Contract.GetApproved(&_Eths.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address)
func (_Eths *EthsCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Eths.Contract.GetApproved(&_Eths.CallOpts, _tokenId)
}

// GetAssRecode is a free data retrieval call binding the contract method 0xd3380bab.
//
// Solidity: function getAssRecode(uint256 _tokenId) constant returns(string)
func (_Eths *EthsCaller) GetAssRecode(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Eths.contract.Call(opts, out, "getAssRecode", _tokenId)
	return *ret0, err
}

// GetAssRecode is a free data retrieval call binding the contract method 0xd3380bab.
//
// Solidity: function getAssRecode(uint256 _tokenId) constant returns(string)
func (_Eths *EthsSession) GetAssRecode(_tokenId *big.Int) (string, error) {
	return _Eths.Contract.GetAssRecode(&_Eths.CallOpts, _tokenId)
}

// GetAssRecode is a free data retrieval call binding the contract method 0xd3380bab.
//
// Solidity: function getAssRecode(uint256 _tokenId) constant returns(string)
func (_Eths *EthsCallerSession) GetAssRecode(_tokenId *big.Int) (string, error) {
	return _Eths.Contract.GetAssRecode(&_Eths.CallOpts, _tokenId)
}

// GetCopyrecode is a free data retrieval call binding the contract method 0xba5bae39.
//
// Solidity: function getCopyrecode(uint256 _tokenId) constant returns(string)
func (_Eths *EthsCaller) GetCopyrecode(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Eths.contract.Call(opts, out, "getCopyrecode", _tokenId)
	return *ret0, err
}

// GetCopyrecode is a free data retrieval call binding the contract method 0xba5bae39.
//
// Solidity: function getCopyrecode(uint256 _tokenId) constant returns(string)
func (_Eths *EthsSession) GetCopyrecode(_tokenId *big.Int) (string, error) {
	return _Eths.Contract.GetCopyrecode(&_Eths.CallOpts, _tokenId)
}

// GetCopyrecode is a free data retrieval call binding the contract method 0xba5bae39.
//
// Solidity: function getCopyrecode(uint256 _tokenId) constant returns(string)
func (_Eths *EthsCallerSession) GetCopyrecode(_tokenId *big.Int) (string, error) {
	return _Eths.Contract.GetCopyrecode(&_Eths.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_Eths *EthsCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Eths.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_Eths *EthsSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Eths.Contract.IsApprovedForAll(&_Eths.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_Eths *EthsCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Eths.Contract.IsApprovedForAll(&_Eths.CallOpts, _owner, _operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address)
func (_Eths *EthsCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Eths.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address)
func (_Eths *EthsSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Eths.Contract.OwnerOf(&_Eths.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address)
func (_Eths *EthsCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Eths.Contract.OwnerOf(&_Eths.CallOpts, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) returns()
func (_Eths *EthsTransactor) Approve(opts *bind.TransactOpts, _approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Eths.contract.Transact(opts, "approve", _approved, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) returns()
func (_Eths *EthsSession) Approve(_approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Eths.Contract.Approve(&_Eths.TransactOpts, _approved, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) returns()
func (_Eths *EthsTransactorSession) Approve(_approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Eths.Contract.Approve(&_Eths.TransactOpts, _approved, _tokenId)
}

// AssRecode is a paid mutator transaction binding the contract method 0xc9e509a3.
//
// Solidity: function assRecode(string _tran, uint256 _tokenId) returns(string)
func (_Eths *EthsTransactor) AssRecode(opts *bind.TransactOpts, _tran string, _tokenId *big.Int) (*types.Transaction, error) {
	return _Eths.contract.Transact(opts, "assRecode", _tran, _tokenId)
}

// AssRecode is a paid mutator transaction binding the contract method 0xc9e509a3.
//
// Solidity: function assRecode(string _tran, uint256 _tokenId) returns(string)
func (_Eths *EthsSession) AssRecode(_tran string, _tokenId *big.Int) (*types.Transaction, error) {
	return _Eths.Contract.AssRecode(&_Eths.TransactOpts, _tran, _tokenId)
}

// AssRecode is a paid mutator transaction binding the contract method 0xc9e509a3.
//
// Solidity: function assRecode(string _tran, uint256 _tokenId) returns(string)
func (_Eths *EthsTransactorSession) AssRecode(_tran string, _tokenId *big.Int) (*types.Transaction, error) {
	return _Eths.Contract.AssRecode(&_Eths.TransactOpts, _tran, _tokenId)
}

// CopyRecode is a paid mutator transaction binding the contract method 0xaa72bdef.
//
// Solidity: function copyRecode(string _recode, uint256 _tokenId) returns(string)
func (_Eths *EthsTransactor) CopyRecode(opts *bind.TransactOpts, _recode string, _tokenId *big.Int) (*types.Transaction, error) {
	return _Eths.contract.Transact(opts, "copyRecode", _recode, _tokenId)
}

// CopyRecode is a paid mutator transaction binding the contract method 0xaa72bdef.
//
// Solidity: function copyRecode(string _recode, uint256 _tokenId) returns(string)
func (_Eths *EthsSession) CopyRecode(_recode string, _tokenId *big.Int) (*types.Transaction, error) {
	return _Eths.Contract.CopyRecode(&_Eths.TransactOpts, _recode, _tokenId)
}

// CopyRecode is a paid mutator transaction binding the contract method 0xaa72bdef.
//
// Solidity: function copyRecode(string _recode, uint256 _tokenId) returns(string)
func (_Eths *EthsTransactorSession) CopyRecode(_recode string, _tokenId *big.Int) (*types.Transaction, error) {
	return _Eths.Contract.CopyRecode(&_Eths.TransactOpts, _recode, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xee7242e3.
//
// Solidity: function mint(bytes32 _hash, string _data) returns()
func (_Eths *EthsTransactor) Mint(opts *bind.TransactOpts, _hash [32]byte, _data string) (*types.Transaction, error) {
	return _Eths.contract.Transact(opts, "mint", _hash, _data)
}

// Mint is a paid mutator transaction binding the contract method 0xee7242e3.
//
// Solidity: function mint(bytes32 _hash, string _data) returns()
func (_Eths *EthsSession) Mint(_hash [32]byte, _data string) (*types.Transaction, error) {
	return _Eths.Contract.Mint(&_Eths.TransactOpts, _hash, _data)
}

// Mint is a paid mutator transaction binding the contract method 0xee7242e3.
//
// Solidity: function mint(bytes32 _hash, string _data) returns()
func (_Eths *EthsTransactorSession) Mint(_hash [32]byte, _data string) (*types.Transaction, error) {
	return _Eths.Contract.Mint(&_Eths.TransactOpts, _hash, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Eths *EthsTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Eths.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Eths *EthsSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Eths.Contract.SafeTransferFrom(&_Eths.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Eths *EthsTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Eths.Contract.SafeTransferFrom(&_Eths.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Eths *EthsTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Eths.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Eths *EthsSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Eths.Contract.SetApprovalForAll(&_Eths.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Eths *EthsTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Eths.Contract.SetApprovalForAll(&_Eths.TransactOpts, _operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Eths *EthsTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Eths.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Eths *EthsSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Eths.Contract.TransferFrom(&_Eths.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Eths *EthsTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Eths.Contract.TransferFrom(&_Eths.TransactOpts, _from, _to, _tokenId)
}

// EthsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Eths contract.
type EthsApprovalIterator struct {
	Event *EthsApproval // Event containing the contract specifics and raw log

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
func (it *EthsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthsApproval)
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
		it.Event = new(EthsApproval)
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
func (it *EthsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthsApproval represents a Approval event raised by the Eths contract.
type EthsApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Eths *EthsFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*EthsApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Eths.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &EthsApprovalIterator{contract: _Eths.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Eths *EthsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EthsApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Eths.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthsApproval)
				if err := _Eths.contract.UnpackLog(event, "Approval", log); err != nil {
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

// EthsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Eths contract.
type EthsApprovalForAllIterator struct {
	Event *EthsApprovalForAll // Event containing the contract specifics and raw log

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
func (it *EthsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthsApprovalForAll)
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
		it.Event = new(EthsApprovalForAll)
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
func (it *EthsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthsApprovalForAll represents a ApprovalForAll event raised by the Eths contract.
type EthsApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Eths *EthsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*EthsApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Eths.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &EthsApprovalForAllIterator{contract: _Eths.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Eths *EthsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *EthsApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Eths.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthsApprovalForAll)
				if err := _Eths.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// EthsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Eths contract.
type EthsTransferIterator struct {
	Event *EthsTransfer // Event containing the contract specifics and raw log

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
func (it *EthsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthsTransfer)
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
		it.Event = new(EthsTransfer)
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
func (it *EthsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthsTransfer represents a Transfer event raised by the Eths contract.
type EthsTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Eths *EthsFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*EthsTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Eths.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &EthsTransferIterator{contract: _Eths.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Eths *EthsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EthsTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Eths.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthsTransfer)
				if err := _Eths.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// EthsNewAssetIterator is returned from FilterNewAsset and is used to iterate over the raw logs and unpacked data for NewAsset events raised by the Eths contract.
type EthsNewAssetIterator struct {
	Event *EthsNewAsset // Event containing the contract specifics and raw log

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
func (it *EthsNewAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthsNewAsset)
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
		it.Event = new(EthsNewAsset)
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
func (it *EthsNewAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthsNewAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthsNewAsset represents a NewAsset event raised by the Eths contract.
type EthsNewAsset struct {
	Hash    [32]byte
	Owner   common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewAsset is a free log retrieval operation binding the contract event 0x33f1fbd35b894d6d223bad33670fff2e57c4cdb440abe337eca70e4f5d231751.
//
// Solidity: event newAsset(bytes32 _hash, address _owner, uint256 _tokenId)
func (_Eths *EthsFilterer) FilterNewAsset(opts *bind.FilterOpts) (*EthsNewAssetIterator, error) {

	logs, sub, err := _Eths.contract.FilterLogs(opts, "newAsset")
	if err != nil {
		return nil, err
	}
	return &EthsNewAssetIterator{contract: _Eths.contract, event: "newAsset", logs: logs, sub: sub}, nil
}

// WatchNewAsset is a free log subscription operation binding the contract event 0x33f1fbd35b894d6d223bad33670fff2e57c4cdb440abe337eca70e4f5d231751.
//
// Solidity: event newAsset(bytes32 _hash, address _owner, uint256 _tokenId)
func (_Eths *EthsFilterer) WatchNewAsset(opts *bind.WatchOpts, sink chan<- *EthsNewAsset) (event.Subscription, error) {

	logs, sub, err := _Eths.contract.WatchLogs(opts, "newAsset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthsNewAsset)
				if err := _Eths.contract.UnpackLog(event, "newAsset", log); err != nil {
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
