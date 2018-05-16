var Web3Utils = require('web3-utils');

var Game = artifacts.require("./LowestUniqueNum.sol");

contract("Game", function([owner, donor]){

	var accounts;

	let game

	beforeEach('setup contract for each test', async () => {
        game = await Game.new(10);

        //watchEvent(game.RevealsSubmitted, function(){game.payout();} );
    })

    it("init", async () => {
        const count = await game.getStakeSize();
        assert.equal(count.toNumber(), web3.toWei(1,'ether'));
    });


    it("Should find winner and distribute prizes", async () => {
        // keccak256 , web3.sha3
        const bet = await game.getStakeSize();

        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        await setMaxPlayers(game, 2);

        await commitGuess(game, accounts[2], "3", "6534");
        await commitGuess(game, accounts[3], "5", "1004");
        //await commitGuess(game, accounts[8], "5", "7728");

        await revealGuess(game, accounts[2], "3", "6534");  
        await revealGuess(game, accounts[3], "5", "1004"); 
        //await revealGuess(game, accounts[8], "5", "7728");

        // console.log("HELLO");

        await game.payout();

        // //console.log(hash)

        const winner = await game.getLastWinners(0);
        console.log(winner);
        
        assert.equal(winner, accounts[2], "Winner isn't correctly selected");
    })

});


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
