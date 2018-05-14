# NASHHASH

Game theory games on the blockchain

## Install

1. clone the app

1. `npm install -g truffle`

1. `npm install -g yarn`

1. `yarn install`

## Build



## Run locally 

1. `truffle develop`
  1. `compile; deploy`

1. `yarn start`  


## API i need to connect to 

NOTICE: 
When creating a new game, make sure you manually set the number of MAX_PLAYERS. I could not find how to test a constructor with params in trufle, so was not ready to include it in. 

I added the function trigger_payout(). This function will trigger the move to payout procedure if REVEAL_PERIOD blocks have gone by since the blocktime of the last commit.

P1: 
1. `commit(hashed guess` √ 
I added an event called SuccesfulCommit to allow you to update the UI.
I think you have to use watch() function.
That should be enough information
to dedeuce that the user was late with their bet.
2. `reveal(key, guess)` √
Same here, except SuccesfulReveal
Do not forget to check for errors.

Here is the amazing js interface I created for you Gilad:

    ////////////////////// GILADS API ///////////////////////
    async function isInCommitState(game){
        var state = await game.game_state_debug();
        if(state.toNumber() == 0){
            return true;
        }else{
            return false;
        }
    }

    async function isInRevealState(game){
        var state = await game.game_state_debug();
        if(state.toNumber() == 1){
            return true;
        }else{
            return false;
        }
    }

    async function getCurrentCommits(game){
        const curr_number_bets = await game.curr_number_bets();
        return curr_number_bets.toNumber();
    }

    async function getCurrentReveals(game){
        var cur_reveals = await game.curr_number_reveals();
        return cur_reveals.toNumber();
    }

    async function resetGame(game){
        game.reset();
    }

    async function commitGuess(game, usr_addr, guess, random){
        const bet = await getBetSize(game);
        const hash = hashGuess(guess, random);
        await game.commit(hash, { value: web3.toWei(bet,'ether'), from: usr_addr });
    }

    async function revealGuess(game, usr_addr, guess, random){
        await game.reveal(guess, random, {from: usr_addr});
    }

    function hashGuess(guess, random){
        return hash = Web3Utils.soliditySha3({type: 'string', value: guess}, {type: 'string', value: random});
    }

    async function getBetSize(game){
        var bet = await game.BET_SIZE();
        return web3.fromWei(bet.toNumber(), 'ether');
    }

    async function getWinners(game){

        var winners = new Array();

        var nw = await game.num_last_winners();
        var number_of_winners = nw.toNumber();
        var i;
        for (i = 0; i < number_of_winners; i++){
            var winner = await game.last_winners(i);
            winners.push(winner);
        }

        return winners;
    }

    async function getPayout(game, usr_addr){

        var winners = getWinners(game);
        var prize = await game.last_prize();
        var i;
        for (i = 0; i < winners.length; i++){
            if(winners[i] == usr_addr){
                return web3.fromWei(prize.toNumber(), 'ether');
            }
        }
        return 0;
    }

    async function getPrizeAmount(game){
        var prize = await game.last_prize();
        return web3.fromWei(prize.toNumber(), 'ether');
    }

    async function getGameFeeAmount(game){
        var fee = await game.GAME_FEE_PERCENT();
        return fee.toNumber();
    }

Simply amazing...


P2.5:
1. bot to always complete game 

P3: 
1. `when will we move ot next stage in game?` => maybe save variable of when game start (block #)  and game end (blcok #)




