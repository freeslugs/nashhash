// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gm

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

// GameABI is the input ABI used to generate the binding from.
const GameABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"gameData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"commits\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"birthBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_maxp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CommitsSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RevealsSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"error\",\"type\":\"string\"}],\"name\":\"DebugEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"DebugWinner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"last_win_l\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"win_l\",\"type\":\"uint256\"}],\"name\":\"DebugCommitState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGameState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentCommits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentReveals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakeSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_stake\",\"type\":\"uint256\"}],\"name\":\"setStakeSize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfWinners\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getLastWinners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastPrize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGameFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMaxPlayers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_max\",\"type\":\"uint256\"}],\"name\":\"setMaxPlayers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resetGame\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"forceToRevealState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"forceToPayoutState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hashedCommit\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guess\",\"type\":\"string\"},{\"name\":\"random\",\"type\":\"string\"}],\"name\":\"reveal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"payout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Game is an auto generated Go binding around an Ethereum contract.
type Game struct {
	GameCaller     // Read-only binding to the contract
	GameTransactor // Write-only binding to the contract
	GameFilterer   // Log filterer for contract events
}

// GameCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameSession struct {
	Contract     *Game             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameCallerSession struct {
	Contract *GameCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameTransactorSession struct {
	Contract     *GameTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameRaw struct {
	Contract *Game // Generic contract binding to access the raw methods on
}

// GameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameCallerRaw struct {
	Contract *GameCaller // Generic read-only contract binding to access the raw methods on
}

// GameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameTransactorRaw struct {
	Contract *GameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGame creates a new instance of Game, bound to a specific deployed contract.
func NewGame(address common.Address, backend bind.ContractBackend) (*Game, error) {
	contract, err := bindGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Game{GameCaller: GameCaller{contract: contract}, GameTransactor: GameTransactor{contract: contract}, GameFilterer: GameFilterer{contract: contract}}, nil
}

// NewGameCaller creates a new read-only instance of Game, bound to a specific deployed contract.
func NewGameCaller(address common.Address, caller bind.ContractCaller) (*GameCaller, error) {
	contract, err := bindGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameCaller{contract: contract}, nil
}

// NewGameTransactor creates a new write-only instance of Game, bound to a specific deployed contract.
func NewGameTransactor(address common.Address, transactor bind.ContractTransactor) (*GameTransactor, error) {
	contract, err := bindGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameTransactor{contract: contract}, nil
}

// NewGameFilterer creates a new log filterer instance of Game, bound to a specific deployed contract.
func NewGameFilterer(address common.Address, filterer bind.ContractFilterer) (*GameFilterer, error) {
	contract, err := bindGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameFilterer{contract: contract}, nil
}

// bindGame binds a generic wrapper to an already deployed contract.
func bindGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Game *GameRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Game.Contract.GameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Game *GameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.Contract.GameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Game *GameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Game.Contract.GameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Game *GameCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Game.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Game *GameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Game *GameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Game.Contract.contract.Transact(opts, method, params...)
}

// BirthBlock is a free data retrieval call binding the contract method 0xbba4b31e.
//
// Solidity: function birthBlock() constant returns(uint256)
func (_Game *GameCaller) BirthBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "birthBlock")
	return *ret0, err
}

// BirthBlock is a free data retrieval call binding the contract method 0xbba4b31e.
//
// Solidity: function birthBlock() constant returns(uint256)
func (_Game *GameSession) BirthBlock() (*big.Int, error) {
	return _Game.Contract.BirthBlock(&_Game.CallOpts)
}

// BirthBlock is a free data retrieval call binding the contract method 0xbba4b31e.
//
// Solidity: function birthBlock() constant returns(uint256)
func (_Game *GameCallerSession) BirthBlock() (*big.Int, error) {
	return _Game.Contract.BirthBlock(&_Game.CallOpts)
}

// Commits is a free data retrieval call binding the contract method 0x7b43a8e6.
//
// Solidity: function commits( address) constant returns(bytes32)
func (_Game *GameCaller) Commits(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "commits", arg0)
	return *ret0, err
}

// Commits is a free data retrieval call binding the contract method 0x7b43a8e6.
//
// Solidity: function commits( address) constant returns(bytes32)
func (_Game *GameSession) Commits(arg0 common.Address) ([32]byte, error) {
	return _Game.Contract.Commits(&_Game.CallOpts, arg0)
}

// Commits is a free data retrieval call binding the contract method 0x7b43a8e6.
//
// Solidity: function commits( address) constant returns(bytes32)
func (_Game *GameCallerSession) Commits(arg0 common.Address) ([32]byte, error) {
	return _Game.Contract.Commits(&_Game.CallOpts, arg0)
}

// GameData is a free data retrieval call binding the contract method 0x29bb80a0.
//
// Solidity: function gameData( address) constant returns(string)
func (_Game *GameCaller) GameData(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "gameData", arg0)
	return *ret0, err
}

// GameData is a free data retrieval call binding the contract method 0x29bb80a0.
//
// Solidity: function gameData( address) constant returns(string)
func (_Game *GameSession) GameData(arg0 common.Address) (string, error) {
	return _Game.Contract.GameData(&_Game.CallOpts, arg0)
}

// GameData is a free data retrieval call binding the contract method 0x29bb80a0.
//
// Solidity: function gameData( address) constant returns(string)
func (_Game *GameCallerSession) GameData(arg0 common.Address) (string, error) {
	return _Game.Contract.GameData(&_Game.CallOpts, arg0)
}

// GetCurrentCommits is a free data retrieval call binding the contract method 0x0d7099ea.
//
// Solidity: function getCurrentCommits() constant returns(uint256)
func (_Game *GameCaller) GetCurrentCommits(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "getCurrentCommits")
	return *ret0, err
}

// GetCurrentCommits is a free data retrieval call binding the contract method 0x0d7099ea.
//
// Solidity: function getCurrentCommits() constant returns(uint256)
func (_Game *GameSession) GetCurrentCommits() (*big.Int, error) {
	return _Game.Contract.GetCurrentCommits(&_Game.CallOpts)
}

// GetCurrentCommits is a free data retrieval call binding the contract method 0x0d7099ea.
//
// Solidity: function getCurrentCommits() constant returns(uint256)
func (_Game *GameCallerSession) GetCurrentCommits() (*big.Int, error) {
	return _Game.Contract.GetCurrentCommits(&_Game.CallOpts)
}

// GetCurrentReveals is a free data retrieval call binding the contract method 0x3e62aa72.
//
// Solidity: function getCurrentReveals() constant returns(uint256)
func (_Game *GameCaller) GetCurrentReveals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "getCurrentReveals")
	return *ret0, err
}

// GetCurrentReveals is a free data retrieval call binding the contract method 0x3e62aa72.
//
// Solidity: function getCurrentReveals() constant returns(uint256)
func (_Game *GameSession) GetCurrentReveals() (*big.Int, error) {
	return _Game.Contract.GetCurrentReveals(&_Game.CallOpts)
}

// GetCurrentReveals is a free data retrieval call binding the contract method 0x3e62aa72.
//
// Solidity: function getCurrentReveals() constant returns(uint256)
func (_Game *GameCallerSession) GetCurrentReveals() (*big.Int, error) {
	return _Game.Contract.GetCurrentReveals(&_Game.CallOpts)
}

// GetGameFee is a free data retrieval call binding the contract method 0x5bba88f1.
//
// Solidity: function getGameFee() constant returns(uint256)
func (_Game *GameCaller) GetGameFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "getGameFee")
	return *ret0, err
}

// GetGameFee is a free data retrieval call binding the contract method 0x5bba88f1.
//
// Solidity: function getGameFee() constant returns(uint256)
func (_Game *GameSession) GetGameFee() (*big.Int, error) {
	return _Game.Contract.GetGameFee(&_Game.CallOpts)
}

// GetGameFee is a free data retrieval call binding the contract method 0x5bba88f1.
//
// Solidity: function getGameFee() constant returns(uint256)
func (_Game *GameCallerSession) GetGameFee() (*big.Int, error) {
	return _Game.Contract.GetGameFee(&_Game.CallOpts)
}

// GetGameState is a free data retrieval call binding the contract method 0xb7d0628b.
//
// Solidity: function getGameState() constant returns(uint256)
func (_Game *GameCaller) GetGameState(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "getGameState")
	return *ret0, err
}

// GetGameState is a free data retrieval call binding the contract method 0xb7d0628b.
//
// Solidity: function getGameState() constant returns(uint256)
func (_Game *GameSession) GetGameState() (*big.Int, error) {
	return _Game.Contract.GetGameState(&_Game.CallOpts)
}

// GetGameState is a free data retrieval call binding the contract method 0xb7d0628b.
//
// Solidity: function getGameState() constant returns(uint256)
func (_Game *GameCallerSession) GetGameState() (*big.Int, error) {
	return _Game.Contract.GetGameState(&_Game.CallOpts)
}

// GetLastPrize is a free data retrieval call binding the contract method 0x0cfa6131.
//
// Solidity: function getLastPrize() constant returns(uint256)
func (_Game *GameCaller) GetLastPrize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "getLastPrize")
	return *ret0, err
}

// GetLastPrize is a free data retrieval call binding the contract method 0x0cfa6131.
//
// Solidity: function getLastPrize() constant returns(uint256)
func (_Game *GameSession) GetLastPrize() (*big.Int, error) {
	return _Game.Contract.GetLastPrize(&_Game.CallOpts)
}

// GetLastPrize is a free data retrieval call binding the contract method 0x0cfa6131.
//
// Solidity: function getLastPrize() constant returns(uint256)
func (_Game *GameCallerSession) GetLastPrize() (*big.Int, error) {
	return _Game.Contract.GetLastPrize(&_Game.CallOpts)
}

// GetLastWinners is a free data retrieval call binding the contract method 0x6764f6fb.
//
// Solidity: function getLastWinners(i uint256) constant returns(address)
func (_Game *GameCaller) GetLastWinners(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "getLastWinners", i)
	return *ret0, err
}

// GetLastWinners is a free data retrieval call binding the contract method 0x6764f6fb.
//
// Solidity: function getLastWinners(i uint256) constant returns(address)
func (_Game *GameSession) GetLastWinners(i *big.Int) (common.Address, error) {
	return _Game.Contract.GetLastWinners(&_Game.CallOpts, i)
}

// GetLastWinners is a free data retrieval call binding the contract method 0x6764f6fb.
//
// Solidity: function getLastWinners(i uint256) constant returns(address)
func (_Game *GameCallerSession) GetLastWinners(i *big.Int) (common.Address, error) {
	return _Game.Contract.GetLastWinners(&_Game.CallOpts, i)
}

// GetMaxPlayers is a free data retrieval call binding the contract method 0x8cbd9eb8.
//
// Solidity: function getMaxPlayers() constant returns(uint256)
func (_Game *GameCaller) GetMaxPlayers(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "getMaxPlayers")
	return *ret0, err
}

// GetMaxPlayers is a free data retrieval call binding the contract method 0x8cbd9eb8.
//
// Solidity: function getMaxPlayers() constant returns(uint256)
func (_Game *GameSession) GetMaxPlayers() (*big.Int, error) {
	return _Game.Contract.GetMaxPlayers(&_Game.CallOpts)
}

// GetMaxPlayers is a free data retrieval call binding the contract method 0x8cbd9eb8.
//
// Solidity: function getMaxPlayers() constant returns(uint256)
func (_Game *GameCallerSession) GetMaxPlayers() (*big.Int, error) {
	return _Game.Contract.GetMaxPlayers(&_Game.CallOpts)
}

// GetNumberOfWinners is a free data retrieval call binding the contract method 0x8e33cffc.
//
// Solidity: function getNumberOfWinners() constant returns(uint256)
func (_Game *GameCaller) GetNumberOfWinners(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "getNumberOfWinners")
	return *ret0, err
}

// GetNumberOfWinners is a free data retrieval call binding the contract method 0x8e33cffc.
//
// Solidity: function getNumberOfWinners() constant returns(uint256)
func (_Game *GameSession) GetNumberOfWinners() (*big.Int, error) {
	return _Game.Contract.GetNumberOfWinners(&_Game.CallOpts)
}

// GetNumberOfWinners is a free data retrieval call binding the contract method 0x8e33cffc.
//
// Solidity: function getNumberOfWinners() constant returns(uint256)
func (_Game *GameCallerSession) GetNumberOfWinners() (*big.Int, error) {
	return _Game.Contract.GetNumberOfWinners(&_Game.CallOpts)
}

// GetStakeSize is a free data retrieval call binding the contract method 0x76cc4b7e.
//
// Solidity: function getStakeSize() constant returns(uint256)
func (_Game *GameCaller) GetStakeSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "getStakeSize")
	return *ret0, err
}

// GetStakeSize is a free data retrieval call binding the contract method 0x76cc4b7e.
//
// Solidity: function getStakeSize() constant returns(uint256)
func (_Game *GameSession) GetStakeSize() (*big.Int, error) {
	return _Game.Contract.GetStakeSize(&_Game.CallOpts)
}

// GetStakeSize is a free data retrieval call binding the contract method 0x76cc4b7e.
//
// Solidity: function getStakeSize() constant returns(uint256)
func (_Game *GameCallerSession) GetStakeSize() (*big.Int, error) {
	return _Game.Contract.GetStakeSize(&_Game.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Game *GameCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Game *GameSession) Owner() (common.Address, error) {
	return _Game.Contract.Owner(&_Game.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Game *GameCallerSession) Owner() (common.Address, error) {
	return _Game.Contract.Owner(&_Game.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Game *GameCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Game *GameSession) Paused() (bool, error) {
	return _Game.Contract.Paused(&_Game.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Game *GameCallerSession) Paused() (bool, error) {
	return _Game.Contract.Paused(&_Game.CallOpts)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(hashedCommit bytes32) returns()
func (_Game *GameTransactor) Commit(opts *bind.TransactOpts, hashedCommit [32]byte) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "commit", hashedCommit)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(hashedCommit bytes32) returns()
func (_Game *GameSession) Commit(hashedCommit [32]byte) (*types.Transaction, error) {
	return _Game.Contract.Commit(&_Game.TransactOpts, hashedCommit)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(hashedCommit bytes32) returns()
func (_Game *GameTransactorSession) Commit(hashedCommit [32]byte) (*types.Transaction, error) {
	return _Game.Contract.Commit(&_Game.TransactOpts, hashedCommit)
}

// ForceToPayoutState is a paid mutator transaction binding the contract method 0x760704c1.
//
// Solidity: function forceToPayoutState() returns()
func (_Game *GameTransactor) ForceToPayoutState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "forceToPayoutState")
}

// ForceToPayoutState is a paid mutator transaction binding the contract method 0x760704c1.
//
// Solidity: function forceToPayoutState() returns()
func (_Game *GameSession) ForceToPayoutState() (*types.Transaction, error) {
	return _Game.Contract.ForceToPayoutState(&_Game.TransactOpts)
}

// ForceToPayoutState is a paid mutator transaction binding the contract method 0x760704c1.
//
// Solidity: function forceToPayoutState() returns()
func (_Game *GameTransactorSession) ForceToPayoutState() (*types.Transaction, error) {
	return _Game.Contract.ForceToPayoutState(&_Game.TransactOpts)
}

// ForceToRevealState is a paid mutator transaction binding the contract method 0x86b2b986.
//
// Solidity: function forceToRevealState() returns()
func (_Game *GameTransactor) ForceToRevealState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "forceToRevealState")
}

// ForceToRevealState is a paid mutator transaction binding the contract method 0x86b2b986.
//
// Solidity: function forceToRevealState() returns()
func (_Game *GameSession) ForceToRevealState() (*types.Transaction, error) {
	return _Game.Contract.ForceToRevealState(&_Game.TransactOpts)
}

// ForceToRevealState is a paid mutator transaction binding the contract method 0x86b2b986.
//
// Solidity: function forceToRevealState() returns()
func (_Game *GameTransactorSession) ForceToRevealState() (*types.Transaction, error) {
	return _Game.Contract.ForceToRevealState(&_Game.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Game *GameTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Game *GameSession) Pause() (*types.Transaction, error) {
	return _Game.Contract.Pause(&_Game.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Game *GameTransactorSession) Pause() (*types.Transaction, error) {
	return _Game.Contract.Pause(&_Game.TransactOpts)
}

// Payout is a paid mutator transaction binding the contract method 0x63bd1d4a.
//
// Solidity: function payout() returns()
func (_Game *GameTransactor) Payout(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "payout")
}

// Payout is a paid mutator transaction binding the contract method 0x63bd1d4a.
//
// Solidity: function payout() returns()
func (_Game *GameSession) Payout() (*types.Transaction, error) {
	return _Game.Contract.Payout(&_Game.TransactOpts)
}

// Payout is a paid mutator transaction binding the contract method 0x63bd1d4a.
//
// Solidity: function payout() returns()
func (_Game *GameTransactorSession) Payout() (*types.Transaction, error) {
	return _Game.Contract.Payout(&_Game.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Game *GameTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Game *GameSession) RenounceOwnership() (*types.Transaction, error) {
	return _Game.Contract.RenounceOwnership(&_Game.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Game *GameTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Game.Contract.RenounceOwnership(&_Game.TransactOpts)
}

// ResetGame is a paid mutator transaction binding the contract method 0xbb472219.
//
// Solidity: function resetGame() returns()
func (_Game *GameTransactor) ResetGame(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "resetGame")
}

// ResetGame is a paid mutator transaction binding the contract method 0xbb472219.
//
// Solidity: function resetGame() returns()
func (_Game *GameSession) ResetGame() (*types.Transaction, error) {
	return _Game.Contract.ResetGame(&_Game.TransactOpts)
}

// ResetGame is a paid mutator transaction binding the contract method 0xbb472219.
//
// Solidity: function resetGame() returns()
func (_Game *GameTransactorSession) ResetGame() (*types.Transaction, error) {
	return _Game.Contract.ResetGame(&_Game.TransactOpts)
}

// Reveal is a paid mutator transaction binding the contract method 0x5ba1b2c7.
//
// Solidity: function reveal(guess string, random string) returns()
func (_Game *GameTransactor) Reveal(opts *bind.TransactOpts, guess string, random string) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "reveal", guess, random)
}

// Reveal is a paid mutator transaction binding the contract method 0x5ba1b2c7.
//
// Solidity: function reveal(guess string, random string) returns()
func (_Game *GameSession) Reveal(guess string, random string) (*types.Transaction, error) {
	return _Game.Contract.Reveal(&_Game.TransactOpts, guess, random)
}

// Reveal is a paid mutator transaction binding the contract method 0x5ba1b2c7.
//
// Solidity: function reveal(guess string, random string) returns()
func (_Game *GameTransactorSession) Reveal(guess string, random string) (*types.Transaction, error) {
	return _Game.Contract.Reveal(&_Game.TransactOpts, guess, random)
}

// SetMaxPlayers is a paid mutator transaction binding the contract method 0x288dee3b.
//
// Solidity: function setMaxPlayers(new_max uint256) returns()
func (_Game *GameTransactor) SetMaxPlayers(opts *bind.TransactOpts, new_max *big.Int) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "setMaxPlayers", new_max)
}

// SetMaxPlayers is a paid mutator transaction binding the contract method 0x288dee3b.
//
// Solidity: function setMaxPlayers(new_max uint256) returns()
func (_Game *GameSession) SetMaxPlayers(new_max *big.Int) (*types.Transaction, error) {
	return _Game.Contract.SetMaxPlayers(&_Game.TransactOpts, new_max)
}

// SetMaxPlayers is a paid mutator transaction binding the contract method 0x288dee3b.
//
// Solidity: function setMaxPlayers(new_max uint256) returns()
func (_Game *GameTransactorSession) SetMaxPlayers(new_max *big.Int) (*types.Transaction, error) {
	return _Game.Contract.SetMaxPlayers(&_Game.TransactOpts, new_max)
}

// SetStakeSize is a paid mutator transaction binding the contract method 0xf3b9204e.
//
// Solidity: function setStakeSize(new_stake uint256) returns()
func (_Game *GameTransactor) SetStakeSize(opts *bind.TransactOpts, new_stake *big.Int) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "setStakeSize", new_stake)
}

// SetStakeSize is a paid mutator transaction binding the contract method 0xf3b9204e.
//
// Solidity: function setStakeSize(new_stake uint256) returns()
func (_Game *GameSession) SetStakeSize(new_stake *big.Int) (*types.Transaction, error) {
	return _Game.Contract.SetStakeSize(&_Game.TransactOpts, new_stake)
}

// SetStakeSize is a paid mutator transaction binding the contract method 0xf3b9204e.
//
// Solidity: function setStakeSize(new_stake uint256) returns()
func (_Game *GameTransactorSession) SetStakeSize(new_stake *big.Int) (*types.Transaction, error) {
	return _Game.Contract.SetStakeSize(&_Game.TransactOpts, new_stake)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Game *GameTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Game *GameSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Game.Contract.TransferOwnership(&_Game.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Game *GameTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Game.Contract.TransferOwnership(&_Game.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Game *GameTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Game *GameSession) Unpause() (*types.Transaction, error) {
	return _Game.Contract.Unpause(&_Game.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Game *GameTransactorSession) Unpause() (*types.Transaction, error) {
	return _Game.Contract.Unpause(&_Game.TransactOpts)
}

// GameCommitsSubmittedIterator is returned from FilterCommitsSubmitted and is used to iterate over the raw logs and unpacked data for CommitsSubmitted events raised by the Game contract.
type GameCommitsSubmittedIterator struct {
	Event *GameCommitsSubmitted // Event containing the contract specifics and raw log

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
func (it *GameCommitsSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameCommitsSubmitted)
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
		it.Event = new(GameCommitsSubmitted)
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
func (it *GameCommitsSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameCommitsSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameCommitsSubmitted represents a CommitsSubmitted event raised by the Game contract.
type GameCommitsSubmitted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCommitsSubmitted is a free log retrieval operation binding the contract event 0x69c81b49bf12351dbabf91faefefa2b2ee324d944cd6a8b3b8416ade02326664.
//
// Solidity: e CommitsSubmitted()
func (_Game *GameFilterer) FilterCommitsSubmitted(opts *bind.FilterOpts) (*GameCommitsSubmittedIterator, error) {

	logs, sub, err := _Game.contract.FilterLogs(opts, "CommitsSubmitted")
	if err != nil {
		return nil, err
	}
	return &GameCommitsSubmittedIterator{contract: _Game.contract, event: "CommitsSubmitted", logs: logs, sub: sub}, nil
}

// WatchCommitsSubmitted is a free log subscription operation binding the contract event 0x69c81b49bf12351dbabf91faefefa2b2ee324d944cd6a8b3b8416ade02326664.
//
// Solidity: e CommitsSubmitted()
func (_Game *GameFilterer) WatchCommitsSubmitted(opts *bind.WatchOpts, sink chan<- *GameCommitsSubmitted) (event.Subscription, error) {

	logs, sub, err := _Game.contract.WatchLogs(opts, "CommitsSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameCommitsSubmitted)
				if err := _Game.contract.UnpackLog(event, "CommitsSubmitted", log); err != nil {
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

// GameDebugCommitStateIterator is returned from FilterDebugCommitState and is used to iterate over the raw logs and unpacked data for DebugCommitState events raised by the Game contract.
type GameDebugCommitStateIterator struct {
	Event *GameDebugCommitState // Event containing the contract specifics and raw log

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
func (it *GameDebugCommitStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameDebugCommitState)
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
		it.Event = new(GameDebugCommitState)
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
func (it *GameDebugCommitStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameDebugCommitStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameDebugCommitState represents a DebugCommitState event raised by the Game contract.
type GameDebugCommitState struct {
	LastWinL *big.Int
	WinL     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDebugCommitState is a free log retrieval operation binding the contract event 0x278480b035975688c107f41ad805694077b1c71eb9ef9c10132f19f893c87c0c.
//
// Solidity: e DebugCommitState(last_win_l uint256, win_l uint256)
func (_Game *GameFilterer) FilterDebugCommitState(opts *bind.FilterOpts) (*GameDebugCommitStateIterator, error) {

	logs, sub, err := _Game.contract.FilterLogs(opts, "DebugCommitState")
	if err != nil {
		return nil, err
	}
	return &GameDebugCommitStateIterator{contract: _Game.contract, event: "DebugCommitState", logs: logs, sub: sub}, nil
}

// WatchDebugCommitState is a free log subscription operation binding the contract event 0x278480b035975688c107f41ad805694077b1c71eb9ef9c10132f19f893c87c0c.
//
// Solidity: e DebugCommitState(last_win_l uint256, win_l uint256)
func (_Game *GameFilterer) WatchDebugCommitState(opts *bind.WatchOpts, sink chan<- *GameDebugCommitState) (event.Subscription, error) {

	logs, sub, err := _Game.contract.WatchLogs(opts, "DebugCommitState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameDebugCommitState)
				if err := _Game.contract.UnpackLog(event, "DebugCommitState", log); err != nil {
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

// GameDebugEventIterator is returned from FilterDebugEvent and is used to iterate over the raw logs and unpacked data for DebugEvent events raised by the Game contract.
type GameDebugEventIterator struct {
	Event *GameDebugEvent // Event containing the contract specifics and raw log

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
func (it *GameDebugEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameDebugEvent)
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
		it.Event = new(GameDebugEvent)
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
func (it *GameDebugEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameDebugEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameDebugEvent represents a DebugEvent event raised by the Game contract.
type GameDebugEvent struct {
	Error string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDebugEvent is a free log retrieval operation binding the contract event 0x56f074d292557f2e3c567d982816e0fb5b72100ff196892f8fbd23b8a9073679.
//
// Solidity: e DebugEvent(error string)
func (_Game *GameFilterer) FilterDebugEvent(opts *bind.FilterOpts) (*GameDebugEventIterator, error) {

	logs, sub, err := _Game.contract.FilterLogs(opts, "DebugEvent")
	if err != nil {
		return nil, err
	}
	return &GameDebugEventIterator{contract: _Game.contract, event: "DebugEvent", logs: logs, sub: sub}, nil
}

// WatchDebugEvent is a free log subscription operation binding the contract event 0x56f074d292557f2e3c567d982816e0fb5b72100ff196892f8fbd23b8a9073679.
//
// Solidity: e DebugEvent(error string)
func (_Game *GameFilterer) WatchDebugEvent(opts *bind.WatchOpts, sink chan<- *GameDebugEvent) (event.Subscription, error) {

	logs, sub, err := _Game.contract.WatchLogs(opts, "DebugEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameDebugEvent)
				if err := _Game.contract.UnpackLog(event, "DebugEvent", log); err != nil {
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

// GameDebugWinnerIterator is returned from FilterDebugWinner and is used to iterate over the raw logs and unpacked data for DebugWinner events raised by the Game contract.
type GameDebugWinnerIterator struct {
	Event *GameDebugWinner // Event containing the contract specifics and raw log

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
func (it *GameDebugWinnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameDebugWinner)
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
		it.Event = new(GameDebugWinner)
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
func (it *GameDebugWinnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameDebugWinnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameDebugWinner represents a DebugWinner event raised by the Game contract.
type GameDebugWinner struct {
	Addr common.Address
	N    *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDebugWinner is a free log retrieval operation binding the contract event 0xece34cfb64cf1d40c8a91ac1e5395fb2f1586a6d85d750495dd23dd2b22a596d.
//
// Solidity: e DebugWinner(addr address, n uint256)
func (_Game *GameFilterer) FilterDebugWinner(opts *bind.FilterOpts) (*GameDebugWinnerIterator, error) {

	logs, sub, err := _Game.contract.FilterLogs(opts, "DebugWinner")
	if err != nil {
		return nil, err
	}
	return &GameDebugWinnerIterator{contract: _Game.contract, event: "DebugWinner", logs: logs, sub: sub}, nil
}

// WatchDebugWinner is a free log subscription operation binding the contract event 0xece34cfb64cf1d40c8a91ac1e5395fb2f1586a6d85d750495dd23dd2b22a596d.
//
// Solidity: e DebugWinner(addr address, n uint256)
func (_Game *GameFilterer) WatchDebugWinner(opts *bind.WatchOpts, sink chan<- *GameDebugWinner) (event.Subscription, error) {

	logs, sub, err := _Game.contract.WatchLogs(opts, "DebugWinner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameDebugWinner)
				if err := _Game.contract.UnpackLog(event, "DebugWinner", log); err != nil {
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

// GameOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Game contract.
type GameOwnershipRenouncedIterator struct {
	Event *GameOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *GameOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameOwnershipRenounced)
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
		it.Event = new(GameOwnershipRenounced)
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
func (it *GameOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameOwnershipRenounced represents a OwnershipRenounced event raised by the Game contract.
type GameOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Game *GameFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*GameOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Game.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GameOwnershipRenouncedIterator{contract: _Game.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Game *GameFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *GameOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Game.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameOwnershipRenounced)
				if err := _Game.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// GameOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Game contract.
type GameOwnershipTransferredIterator struct {
	Event *GameOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GameOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameOwnershipTransferred)
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
		it.Event = new(GameOwnershipTransferred)
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
func (it *GameOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameOwnershipTransferred represents a OwnershipTransferred event raised by the Game contract.
type GameOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Game *GameFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GameOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Game.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GameOwnershipTransferredIterator{contract: _Game.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Game *GameFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GameOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Game.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameOwnershipTransferred)
				if err := _Game.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// GamePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Game contract.
type GamePauseIterator struct {
	Event *GamePause // Event containing the contract specifics and raw log

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
func (it *GamePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePause)
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
		it.Event = new(GamePause)
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
func (it *GamePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePause represents a Pause event raised by the Game contract.
type GamePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Game *GameFilterer) FilterPause(opts *bind.FilterOpts) (*GamePauseIterator, error) {

	logs, sub, err := _Game.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &GamePauseIterator{contract: _Game.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Game *GameFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *GamePause) (event.Subscription, error) {

	logs, sub, err := _Game.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePause)
				if err := _Game.contract.UnpackLog(event, "Pause", log); err != nil {
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

// GameRevealsSubmittedIterator is returned from FilterRevealsSubmitted and is used to iterate over the raw logs and unpacked data for RevealsSubmitted events raised by the Game contract.
type GameRevealsSubmittedIterator struct {
	Event *GameRevealsSubmitted // Event containing the contract specifics and raw log

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
func (it *GameRevealsSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameRevealsSubmitted)
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
		it.Event = new(GameRevealsSubmitted)
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
func (it *GameRevealsSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameRevealsSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameRevealsSubmitted represents a RevealsSubmitted event raised by the Game contract.
type GameRevealsSubmitted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRevealsSubmitted is a free log retrieval operation binding the contract event 0x249eb5abee408afd11bfd7a0449ece5a765def88489dd844882a500f87280a24.
//
// Solidity: e RevealsSubmitted()
func (_Game *GameFilterer) FilterRevealsSubmitted(opts *bind.FilterOpts) (*GameRevealsSubmittedIterator, error) {

	logs, sub, err := _Game.contract.FilterLogs(opts, "RevealsSubmitted")
	if err != nil {
		return nil, err
	}
	return &GameRevealsSubmittedIterator{contract: _Game.contract, event: "RevealsSubmitted", logs: logs, sub: sub}, nil
}

// WatchRevealsSubmitted is a free log subscription operation binding the contract event 0x249eb5abee408afd11bfd7a0449ece5a765def88489dd844882a500f87280a24.
//
// Solidity: e RevealsSubmitted()
func (_Game *GameFilterer) WatchRevealsSubmitted(opts *bind.WatchOpts, sink chan<- *GameRevealsSubmitted) (event.Subscription, error) {

	logs, sub, err := _Game.contract.WatchLogs(opts, "RevealsSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameRevealsSubmitted)
				if err := _Game.contract.UnpackLog(event, "RevealsSubmitted", log); err != nil {
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

// GameUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Game contract.
type GameUnpauseIterator struct {
	Event *GameUnpause // Event containing the contract specifics and raw log

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
func (it *GameUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameUnpause)
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
		it.Event = new(GameUnpause)
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
func (it *GameUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameUnpause represents a Unpause event raised by the Game contract.
type GameUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Game *GameFilterer) FilterUnpause(opts *bind.FilterOpts) (*GameUnpauseIterator, error) {

	logs, sub, err := _Game.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &GameUnpauseIterator{contract: _Game.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Game *GameFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *GameUnpause) (event.Subscription, error) {

	logs, sub, err := _Game.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameUnpause)
				if err := _Game.contract.UnpackLog(event, "Unpause", log); err != nil {
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
