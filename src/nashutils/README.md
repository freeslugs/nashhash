# Gamemaster

## Build

1. Install golang        
        brew install go
3. Setup the GOPATH
        export GOPATH=$(pwd)
You want to set it up to some/path/nashhash/src
4. Genereate the go binding for the Game.sol
        cd gm
        go generate
5. In a separate terminal window, starth geth on rinkeby. It might take sometime to sync.
        geth --rinkeby

You will then need to find the location of the geth.ipc. On my machine, this is the location:
        "/Users/me/Library/Ethereum/rinkeby/geth.ipc"
Once you have your geth.ipc path, you need to go into gm/commmon.go and set the constatn EthClientPath to that path.
        EthClientPath = "/Users/me/Library/Ethereum/rinkeby/geth.ipc"

5. Run tests
        go generate
        go test -v


## Description

Gamemaster is a piece of server side code that manages the state of all instances
of the games. Each newly created game (aka newly deployed contract) is assigned a GameOperator. 

GameOperator runs as a separate go routine who's main purpose is to control game state transitions.
A contract can be in three different states:
1. COMMIT_STATE – the contract is accepting commits
2. REVEAL_STATE – the contract is accepting reveals
3. PAYOUT_STATE – the contract is ready for the payout() to be called by the owner.

The states transitions as follows:
COMMIT_STATE -> REVEAL_STATE -> PAYOUT_STATE -> COMMIT_STATE -> REVEAL_STATE -> ... 

The GameOperator has three different transitions it can influence. 
1. When the contract is in PAYOUT_STATE, the GameOperator is responsibile for noticing this and
calling payout() method of the contract (and thus pay the gas fee for findWinners() internal function)
2. When the contract has been in COMMIT_STATE for too long, we might want the game to start without waiting for the rest of the players. In that case the owner has to call forceToRevealState() to 
trigger a COMMIT_STATE -> REVEAL_STATE transition.
3. Similar but with forceToPayoutState()

