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
const GameABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getLastPrize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentCommits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_max\",\"type\":\"uint256\"}],\"name\":\"setMaxPlayers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"gameData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGameStageLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentReveals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGameStateInfo\",\"outputs\":[{\"name\":\"_state\",\"type\":\"uint256\"},{\"name\":\"_currNumberCommits\",\"type\":\"uint256\"},{\"name\":\"_currNumberReveals\",\"type\":\"uint256\"},{\"name\":\"_commitStageStartBlock\",\"type\":\"uint256\"},{\"name\":\"_revealStageStartBlock\",\"type\":\"uint256\"},{\"name\":\"_stageLength\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guess\",\"type\":\"string\"},{\"name\":\"random\",\"type\":\"string\"}],\"name\":\"reveal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGameFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"payout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"npt_addr\",\"type\":\"address\"}],\"name\":\"setNPTAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getLastWinners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCommitStageStartBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"forceToPayoutState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakeSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"commits\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"forceToRevealState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMaxPlayers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfWinners\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRevealStageStartBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGameState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resetGame\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"birthBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hashedCommit\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_stake\",\"type\":\"uint256\"}],\"name\":\"setStakeSize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_feeAddress\",\"type\":\"address\"},{\"name\":\"_gameFeePercent\",\"type\":\"uint256\"},{\"name\":\"_stakeSize\",\"type\":\"uint256\"},{\"name\":\"_maxp\",\"type\":\"uint256\"},{\"name\":\"_gameStageLength\",\"type\":\"uint256\"},{\"name\":\"_nptAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CommitsSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RevealsSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NewRoundStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// GameBin is the compiled bytecode used for deploying new contracts.
const GameBin = `0x`

// DeployGame deploys a new Ethereum contract, binding an instance of Game to it.
func DeployGame(auth *bind.TransactOpts, backend bind.ContractBackend, _feeAddress common.Address, _gameFeePercent *big.Int, _stakeSize *big.Int, _maxp *big.Int, _gameStageLength *big.Int, _nptAddress common.Address) (common.Address, *types.Transaction, *Game, error) {
	parsed, err := abi.JSON(strings.NewReader(GameABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GameBin), backend, _feeAddress, _gameFeePercent, _stakeSize, _maxp, _gameStageLength, _nptAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Game{GameCaller: GameCaller{contract: contract}, GameTransactor: GameTransactor{contract: contract}, GameFilterer: GameFilterer{contract: contract}}, nil
}

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

// GetCommitStageStartBlock is a free data retrieval call binding the contract method 0x6cb9a390.
//
// Solidity: function getCommitStageStartBlock() constant returns(uint256)
func (_Game *GameCaller) GetCommitStageStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "getCommitStageStartBlock")
	return *ret0, err
}

// GetCommitStageStartBlock is a free data retrieval call binding the contract method 0x6cb9a390.
//
// Solidity: function getCommitStageStartBlock() constant returns(uint256)
func (_Game *GameSession) GetCommitStageStartBlock() (*big.Int, error) {
	return _Game.Contract.GetCommitStageStartBlock(&_Game.CallOpts)
}

// GetCommitStageStartBlock is a free data retrieval call binding the contract method 0x6cb9a390.
//
// Solidity: function getCommitStageStartBlock() constant returns(uint256)
func (_Game *GameCallerSession) GetCommitStageStartBlock() (*big.Int, error) {
	return _Game.Contract.GetCommitStageStartBlock(&_Game.CallOpts)
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

// GetGameStageLength is a free data retrieval call binding the contract method 0x349570d5.
//
// Solidity: function getGameStageLength() constant returns(uint256)
func (_Game *GameCaller) GetGameStageLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "getGameStageLength")
	return *ret0, err
}

// GetGameStageLength is a free data retrieval call binding the contract method 0x349570d5.
//
// Solidity: function getGameStageLength() constant returns(uint256)
func (_Game *GameSession) GetGameStageLength() (*big.Int, error) {
	return _Game.Contract.GetGameStageLength(&_Game.CallOpts)
}

// GetGameStageLength is a free data retrieval call binding the contract method 0x349570d5.
//
// Solidity: function getGameStageLength() constant returns(uint256)
func (_Game *GameCallerSession) GetGameStageLength() (*big.Int, error) {
	return _Game.Contract.GetGameStageLength(&_Game.CallOpts)
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

// GetGameStateInfo is a free data retrieval call binding the contract method 0x5086f330.
//
// Solidity: function getGameStateInfo() constant returns(_state uint256, _currNumberCommits uint256, _currNumberReveals uint256, _commitStageStartBlock uint256, _revealStageStartBlock uint256, _stageLength uint256)
func (_Game *GameCaller) GetGameStateInfo(opts *bind.CallOpts) (struct {
	State                 *big.Int
	CurrNumberCommits     *big.Int
	CurrNumberReveals     *big.Int
	CommitStageStartBlock *big.Int
	RevealStageStartBlock *big.Int
	StageLength           *big.Int
}, error) {
	ret := new(struct {
		State                 *big.Int
		CurrNumberCommits     *big.Int
		CurrNumberReveals     *big.Int
		CommitStageStartBlock *big.Int
		RevealStageStartBlock *big.Int
		StageLength           *big.Int
	})
	out := ret
	err := _Game.contract.Call(opts, out, "getGameStateInfo")
	return *ret, err
}

// GetGameStateInfo is a free data retrieval call binding the contract method 0x5086f330.
//
// Solidity: function getGameStateInfo() constant returns(_state uint256, _currNumberCommits uint256, _currNumberReveals uint256, _commitStageStartBlock uint256, _revealStageStartBlock uint256, _stageLength uint256)
func (_Game *GameSession) GetGameStateInfo() (struct {
	State                 *big.Int
	CurrNumberCommits     *big.Int
	CurrNumberReveals     *big.Int
	CommitStageStartBlock *big.Int
	RevealStageStartBlock *big.Int
	StageLength           *big.Int
}, error) {
	return _Game.Contract.GetGameStateInfo(&_Game.CallOpts)
}

// GetGameStateInfo is a free data retrieval call binding the contract method 0x5086f330.
//
// Solidity: function getGameStateInfo() constant returns(_state uint256, _currNumberCommits uint256, _currNumberReveals uint256, _commitStageStartBlock uint256, _revealStageStartBlock uint256, _stageLength uint256)
func (_Game *GameCallerSession) GetGameStateInfo() (struct {
	State                 *big.Int
	CurrNumberCommits     *big.Int
	CurrNumberReveals     *big.Int
	CommitStageStartBlock *big.Int
	RevealStageStartBlock *big.Int
	StageLength           *big.Int
}, error) {
	return _Game.Contract.GetGameStateInfo(&_Game.CallOpts)
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

// GetRevealStageStartBlock is a free data retrieval call binding the contract method 0x9ebefb68.
//
// Solidity: function getRevealStageStartBlock() constant returns(uint256)
func (_Game *GameCaller) GetRevealStageStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Game.contract.Call(opts, out, "getRevealStageStartBlock")
	return *ret0, err
}

// GetRevealStageStartBlock is a free data retrieval call binding the contract method 0x9ebefb68.
//
// Solidity: function getRevealStageStartBlock() constant returns(uint256)
func (_Game *GameSession) GetRevealStageStartBlock() (*big.Int, error) {
	return _Game.Contract.GetRevealStageStartBlock(&_Game.CallOpts)
}

// GetRevealStageStartBlock is a free data retrieval call binding the contract method 0x9ebefb68.
//
// Solidity: function getRevealStageStartBlock() constant returns(uint256)
func (_Game *GameCallerSession) GetRevealStageStartBlock() (*big.Int, error) {
	return _Game.Contract.GetRevealStageStartBlock(&_Game.CallOpts)
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

// SetNPTAddress is a paid mutator transaction binding the contract method 0x6588ba5b.
//
// Solidity: function setNPTAddress(npt_addr address) returns()
func (_Game *GameTransactor) SetNPTAddress(opts *bind.TransactOpts, npt_addr common.Address) (*types.Transaction, error) {
	return _Game.contract.Transact(opts, "setNPTAddress", npt_addr)
}

// SetNPTAddress is a paid mutator transaction binding the contract method 0x6588ba5b.
//
// Solidity: function setNPTAddress(npt_addr address) returns()
func (_Game *GameSession) SetNPTAddress(npt_addr common.Address) (*types.Transaction, error) {
	return _Game.Contract.SetNPTAddress(&_Game.TransactOpts, npt_addr)
}

// SetNPTAddress is a paid mutator transaction binding the contract method 0x6588ba5b.
//
// Solidity: function setNPTAddress(npt_addr address) returns()
func (_Game *GameTransactorSession) SetNPTAddress(npt_addr common.Address) (*types.Transaction, error) {
	return _Game.Contract.SetNPTAddress(&_Game.TransactOpts, npt_addr)
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

// GameNewRoundStartedIterator is returned from FilterNewRoundStarted and is used to iterate over the raw logs and unpacked data for NewRoundStarted events raised by the Game contract.
type GameNewRoundStartedIterator struct {
	Event *GameNewRoundStarted // Event containing the contract specifics and raw log

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
func (it *GameNewRoundStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameNewRoundStarted)
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
		it.Event = new(GameNewRoundStarted)
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
func (it *GameNewRoundStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameNewRoundStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameNewRoundStarted represents a NewRoundStarted event raised by the Game contract.
type GameNewRoundStarted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewRoundStarted is a free log retrieval operation binding the contract event 0xc80d42c5999c8ad689a434b47a3eecdc54fa3b27a31b49cb02b8c5075929fecc.
//
// Solidity: e NewRoundStarted()
func (_Game *GameFilterer) FilterNewRoundStarted(opts *bind.FilterOpts) (*GameNewRoundStartedIterator, error) {

	logs, sub, err := _Game.contract.FilterLogs(opts, "NewRoundStarted")
	if err != nil {
		return nil, err
	}
	return &GameNewRoundStartedIterator{contract: _Game.contract, event: "NewRoundStarted", logs: logs, sub: sub}, nil
}

// WatchNewRoundStarted is a free log subscription operation binding the contract event 0xc80d42c5999c8ad689a434b47a3eecdc54fa3b27a31b49cb02b8c5075929fecc.
//
// Solidity: e NewRoundStarted()
func (_Game *GameFilterer) WatchNewRoundStarted(opts *bind.WatchOpts, sink chan<- *GameNewRoundStarted) (event.Subscription, error) {

	logs, sub, err := _Game.contract.WatchLogs(opts, "NewRoundStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameNewRoundStarted)
				if err := _Game.contract.UnpackLog(event, "NewRoundStarted", log); err != nil {
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

// GameHelperABI is the input ABI used to generate the binding from.
const GameHelperABI = "[]"

// GameHelperBin is the compiled bytecode used for deploying new contracts.
const GameHelperBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a7230582032efe2ea0cf6cb66aa419d74a4d2935e34c8b6fbb79d815aa271db0f39619c540029`

// DeployGameHelper deploys a new Ethereum contract, binding an instance of GameHelper to it.
func DeployGameHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GameHelper, error) {
	parsed, err := abi.JSON(strings.NewReader(GameHelperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GameHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GameHelper{GameHelperCaller: GameHelperCaller{contract: contract}, GameHelperTransactor: GameHelperTransactor{contract: contract}, GameHelperFilterer: GameHelperFilterer{contract: contract}}, nil
}

// GameHelper is an auto generated Go binding around an Ethereum contract.
type GameHelper struct {
	GameHelperCaller     // Read-only binding to the contract
	GameHelperTransactor // Write-only binding to the contract
	GameHelperFilterer   // Log filterer for contract events
}

// GameHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameHelperSession struct {
	Contract     *GameHelper       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameHelperCallerSession struct {
	Contract *GameHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GameHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameHelperTransactorSession struct {
	Contract     *GameHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GameHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameHelperRaw struct {
	Contract *GameHelper // Generic contract binding to access the raw methods on
}

// GameHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameHelperCallerRaw struct {
	Contract *GameHelperCaller // Generic read-only contract binding to access the raw methods on
}

// GameHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameHelperTransactorRaw struct {
	Contract *GameHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGameHelper creates a new instance of GameHelper, bound to a specific deployed contract.
func NewGameHelper(address common.Address, backend bind.ContractBackend) (*GameHelper, error) {
	contract, err := bindGameHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GameHelper{GameHelperCaller: GameHelperCaller{contract: contract}, GameHelperTransactor: GameHelperTransactor{contract: contract}, GameHelperFilterer: GameHelperFilterer{contract: contract}}, nil
}

// NewGameHelperCaller creates a new read-only instance of GameHelper, bound to a specific deployed contract.
func NewGameHelperCaller(address common.Address, caller bind.ContractCaller) (*GameHelperCaller, error) {
	contract, err := bindGameHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameHelperCaller{contract: contract}, nil
}

// NewGameHelperTransactor creates a new write-only instance of GameHelper, bound to a specific deployed contract.
func NewGameHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*GameHelperTransactor, error) {
	contract, err := bindGameHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameHelperTransactor{contract: contract}, nil
}

// NewGameHelperFilterer creates a new log filterer instance of GameHelper, bound to a specific deployed contract.
func NewGameHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*GameHelperFilterer, error) {
	contract, err := bindGameHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameHelperFilterer{contract: contract}, nil
}

// bindGameHelper binds a generic wrapper to an already deployed contract.
func bindGameHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GameHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameHelper *GameHelperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GameHelper.Contract.GameHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameHelper *GameHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameHelper.Contract.GameHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameHelper *GameHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameHelper.Contract.GameHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameHelper *GameHelperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GameHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameHelper *GameHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameHelper *GameHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameHelper.Contract.contract.Transact(opts, method, params...)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a031916331790556101ff806100326000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663715018a6811461005b5780638da5cb5b14610072578063f2fde38b146100a3575b600080fd5b34801561006757600080fd5b506100706100c4565b005b34801561007e57600080fd5b50610087610130565b60408051600160a060020a039092168252519081900360200190f35b3480156100af57600080fd5b50610070600160a060020a036004351661013f565b600054600160a060020a031633146100db57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a0316331461015657600080fd5b600160a060020a038116151561016b57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820d98b9284cff0c4dc28958097d7057ac0500cb6144c3fed612731a7064e00a8080029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Ownable contract.
type OwnableOwnershipRenouncedIterator struct {
	Event *OwnableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipRenounced)
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
		it.Event = new(OwnableOwnershipRenounced)
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
func (it *OwnableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipRenounced represents a OwnershipRenounced event raised by the Ownable contract.
type OwnableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*OwnableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipRenouncedIterator{contract: _Ownable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipRenounced)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// PausableBin is the compiled bytecode used for deploying new contracts.
const PausableBin = `0x608060405260008054600160a860020a031916331790556103b8806100256000396000f3006080604052600436106100775763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633f4ba83a811461007c5780635c975abb14610093578063715018a6146100bc5780638456cb59146100d15780638da5cb5b146100e6578063f2fde38b14610117575b600080fd5b34801561008857600080fd5b50610091610138565b005b34801561009f57600080fd5b506100a86101bf565b604080519115158252519081900360200190f35b3480156100c857600080fd5b506100916101e0565b3480156100dd57600080fd5b5061009161024c565b3480156100f257600080fd5b506100fb6102e9565b60408051600160a060020a039092168252519081900360200190f35b34801561012357600080fd5b50610091600160a060020a03600435166102f8565b600054600160a060020a0316331461014f57600080fd5b60005474010000000000000000000000000000000000000000900460ff16151561017857600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a1565b60005474010000000000000000000000000000000000000000900460ff1681565b600054600160a060020a031633146101f757600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a0316331461026357600080fd5b60005474010000000000000000000000000000000000000000900460ff161561028b57600080fd5b6000805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a1565b600054600160a060020a031681565b600054600160a060020a0316331461030f57600080fd5b600160a060020a038116151561032457600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058204a36167412d3a4e27fdb097c2901e599be1947777f57f9252136a6edccfaea9f0029`

// DeployPausable deploys a new Ethereum contract, binding an instance of Pausable to it.
func DeployPausable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pausable, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableCallerSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactorSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pausable *PausableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pausable *PausableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pausable.Contract.RenounceOwnership(&_Pausable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pausable *PausableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pausable.Contract.RenounceOwnership(&_Pausable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pausable *PausableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pausable *PausableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pausable *PausableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// PausableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Pausable contract.
type PausableOwnershipRenouncedIterator struct {
	Event *PausableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PausableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableOwnershipRenounced)
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
		it.Event = new(PausableOwnershipRenounced)
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
func (it *PausableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableOwnershipRenounced represents a OwnershipRenounced event raised by the Pausable contract.
type PausableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Pausable *PausableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PausableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableOwnershipRenouncedIterator{contract: _Pausable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Pausable *PausableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PausableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableOwnershipRenounced)
				if err := _Pausable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// PausableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pausable contract.
type PausableOwnershipTransferredIterator struct {
	Event *PausableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PausableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableOwnershipTransferred)
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
		it.Event = new(PausableOwnershipTransferred)
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
func (it *PausableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableOwnershipTransferred represents a OwnershipTransferred event raised by the Pausable contract.
type PausableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Pausable *PausableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PausableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableOwnershipTransferredIterator{contract: _Pausable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Pausable *PausableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PausableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableOwnershipTransferred)
				if err := _Pausable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PausablePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Pausable contract.
type PausablePauseIterator struct {
	Event *PausablePause // Event containing the contract specifics and raw log

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
func (it *PausablePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePause)
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
		it.Event = new(PausablePause)
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
func (it *PausablePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePause represents a Pause event raised by the Pausable contract.
type PausablePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Pausable *PausableFilterer) FilterPause(opts *bind.FilterOpts) (*PausablePauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausablePauseIterator{contract: _Pausable.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Pausable *PausableFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausablePause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePause)
				if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PausableUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Pausable contract.
type PausableUnpauseIterator struct {
	Event *PausableUnpause // Event containing the contract specifics and raw log

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
func (it *PausableUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpause)
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
		it.Event = new(PausableUnpause)
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
func (it *PausableUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpause represents a Unpause event raised by the Pausable contract.
type PausableUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Pausable *PausableFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableUnpauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableUnpauseIterator{contract: _Pausable.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Pausable *PausableFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableUnpause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpause)
				if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
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
