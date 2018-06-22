# NASHHASH ![hi](https://circleci.com/gh/freeslugs/nashhash.png?circle-token=8b2c07bf923ed462dc8f5a8edbe76cc8f5c7457a)

Game theory games on the blockchain

## Install

1. clone the app

1. `npm install -g truffle`

1. `npm install -g yarn`

1. `yarn install`

1. IMPORTANT: you're going to want to set your metamask mneumonic in an `.env` file . used for deployments too.

    1. Create a file in root dir called `.env`

    1. first line set `WORDS=<secret phrase>`

    1. nice all done

## Build



## Run locally 

1. `truffle develop`
  1. `compile; deploy`

1. `yarn start`  

## Deploy

1. Purge assets `rm -r build; rm -r src/contracts`
2. Compile `truffle compile`
3. Deploy: `truffle migrate --reset --network rinkeby`

1. Copy assets to local `cp -r build/contracts src/contracts`

## Test

1. Purge assets `rm -r build;`
2. `truffle develop`
3. `compile`
4. `migrate --reset`
5. `test`


## DEVELOPMENT REQUIREMENTS

Dear n000bs,

1. Don't push broken code to master
  1. If you feel the need to push, just throw it up on a new branch and PR it later 
2. Indent your fucking code. Like, seriously. 
3. See #2 
4. Use useful commits messages. looking at your @talaltoukan and these guys: [e551cc33b4b985d07b3133ec652a5f2c47953598](https://github.com/freeslugs/nashhash/commit/e551cc33b4b985d07b3133ec652a5f2c47953598) and [51f7d02284d230284fc82b42959d133d3094d6f0](https://github.com/freeslugs/nashhash/commit/51f7d02284d230284fc82b42959d133d3094d6f0)


## Common errors

### `INVALID ADDRESS`

make sure you include account in function call. ie. `instance.function(params, { from: user_address })`


P2.5:
1. bot to always complete game 




# Gamemaster

## Build and test the package

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
5. Run tests
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

## Game Master Daemon

### Build
    cd src/gmd
    go build

This will produce an executable gmd

### Run

    gmd -ip=<ip_addr> -port=<port_number> -key=<owner_private_key>

We can omit -ip and -port flags to run gmd on 127.0.0.1:50000
Additionally, port provided must be a private port. 

gmd is not daemonized yet. To run it in the background you will have to explicitly do so.

### Example

Lets run the gmd in the background

    nohup ./gmd -key=76a23cff887b294bb60ccde7ad1eb800f0f6ede70d33b154a53eadb20681a4e3 &>log.txt &

This above command will report the PID of the process.
Use this PID to kill the process if needed.

If we need to kill it

    kill -9 <PID>

## GameMaster Client (gmc)

1. Build

        go build

2. Run

        ./gmc -ip=<gm_ip> -port=<port>

If flags are not provided, will default to 127.0.0.1:50000

3. Example

First, make sure gmd is running. Without it, client will have nowhere to connect to.

Lets start the client:

    ./gmc

If we have a game at contract 0x7B9d950cC1ecD94eD0cF3916989B0ac56C70AB24, we connect it like so:

    gmclerk> connect  0x7B9d950cC1ecD94eD0cF3916989B0ac56C70AB24

The relevant logging should be triggered on the gmd end. If everything went well,
gmc will return the prompt to you. If something went bad, it will print the error and return the promt.  Restart gmc if you are getting this error:

    2018/06/13 04:28:47 connection is shut down

If we now want to disconnect the game, here is how we do it:

    gmclerk> disconnect  0x7B9d950cC1ecD94eD0cF3916989B0ac56C70AB24

We can also connect all games in a file:

    gmclerk> connect -f <filename>
