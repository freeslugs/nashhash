# NASHHASH

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
            var state = await game.getGameState();
            if(state.toNumber() == 0){
                return true;
            }else{
                return false;
            }
        }

        async function isInRevealState(game){
            var state = await game.getGameState();
            if(state.toNumber() == 1){
                return true;
            }else{
                return false;
            }
        }

        async function isInPayoutState(game){
            var state =  await game.getGameState();
            if(state.toNumber() == 2){
                return true;
            }else{
                return false;
            }
            
        }

        async function getCurrentCommits(game){
            const currNumberCommits = await game.getCurrentCommits();
            return currNumberCommits.toNumber();
        }

        async function getCurrentReveals(game){
            var curNumberReveals = await game.getCurrentReveals();
            return curNumberReveals.toNumber();
        }

        async function resetGame(game){
            game.reset();
        }

        async function commitGuess(game, usr_addr, guess, random){
            const bet = await getStakeSize(game);
            const hash = hashGuess(guess, random);
            await game.commit(hash, { value: web3.toWei(bet,'ether'), from: usr_addr });
        }

        async function revealGuess(game, usr_addr, guess, random){
            await game.reveal(guess, random, {from: usr_addr});
        }

        function hashGuess(guess, random){
            return hash = Web3Utils.soliditySha3({type: 'string', value: guess}, {type: 'string', value: random});
        }

        async function getStakeSize(game){
            var bet = await game.getStakeSize();
            return web3.fromWei(bet.toNumber(), 'ether');
        }

        async function getWinners(game){

            var winners = new Array();

            var nw = await game.getNumberOfWinners();
            var number_of_winners = nw.toNumber();
            var i;
            for (i = 0; i < number_of_winners; i++){
                //var winner = await game.last_winners(i);
                var winner = await game.getLastWinners(i);
                winners.push(winner);
            }

            return winners;
        }

        async function getPayout(game, usr_addr){

            var winners = await getWinners(game);
            var prize = await game.getLastPrize();
            var i;
            for (i = 0; i < winners.length; i++){
                if(winners[i] == usr_addr){
                    return web3.fromWei(prize.toNumber(), 'ether');
                }
            }
            return 0;
        }

        async function getPrizeAmount(game){
            var prize = await game.getLastPrize();
            return web3.fromWei(prize.toNumber(), 'ether');
        }

        async function getGameFeeAmount(game){
            var fee = await game.getGameFee();
            return fee.toNumber();
        }

        async function pauseGame(game){
            await game.pause();
        }

        async function unpauseGame(game){
            await game.unpause();
        }

        async function setMaxPlayers(game, num){
            await game.setMaxPlayers(num);
        }

        async function getMaxPlayers(game){
            var num = await game.getMaxPlayers();
            return num.toNumber();
        }

        /* Cool function. 
            - ev is the event to watch for from the contract. EX. game.CommitsSubmitted
            - handler is the function that is to be called when that event is emited by the contract
            - handler_args_list is a list of argumetns to the handler

        */
        function watchEvent(ev, handler, handler_args_list){
            
            var event = ev({}, {fromBlock: 0, toBlock: 'latest'});
            event.watch(function(error, result){
                if(!error){
                    handler.apply(this, handler_args_list);
                }else{
                    console.log(error);
                    assert(true == false, "event handler failed to be installed");
                }
            });
        }

        function watchEvent(ev, handler){
            
            var event = ev({}, {fromBlock: 0, toBlock: 'latest'});
            event.watch(function(error, result){
                if(!error){
                    handler();
                }else{
                    console.log(error);
                    assert(true == false, "event handler failed to be installed");
                }
            });

        }



Simply amazing...


P2.5:
1. bot to always complete game 

P3: 
1. `when will we move ot next stage in game?` => maybe save variable of when game start (block #)  and game end (blcok #)




