# NASHHASH ![hi](https://circleci.com/gh/freeslugs/nashhash.png?circle-token=8b2c07bf923ed462dc8f5a8edbe76cc8f5c7457a)

Game theory games on the blockchain

## Install

1. clone the app

1. `npm install -g truffle`

1. `npm install -g yarn`

1. `yarn install`

1. IMPORTANT: you're going to want to set your metamask mneumonic in an `.env` file . used for deployments too.

    1. Create a file in root dir called `.env`

    1. first line set `WORDS='my words here'`

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

## Common errors

### `INVALID ADDRESS`

make sure you include account in function call. ie. `instance.function(params, { from: user_address })`


P1: 

P2.5:
1. bot to always complete game 

P3: 
1. `when will we move ot next stage in game?` => maybe save variable of when game start (block #)  and game end (blcok #)



## DEVELOPMENT REQUIREMENTS

### Dear n000bs

1. Don't push broken code to master
  1. If you feel the need to push, just throw it up on a new branch and PR it later 
2. Indent your fucking code. Like, seriously. 
3. See #2 
