var Web3Utils = require('web3-utils');

// var keccak256 = require('js-sha3').keccak256;
const FIXED_BET = web3.toWei(1,'ether');
const MAX_PLAYERS = 10;
const HASHNASH_ADDRESS = 0x2540099e9ed04aF369d557a40da2D8f9c2ab928D;

var Game = artifacts.require("./Game.sol");

// function testHandler(args){
//     console.log(args);
// }
  
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

    it("Should commit hashed guess with stake", async () => {
        
        const hash = hashGuess("66", "3");

        await commitGuess(game, donor, "66", "3");

        const curr_number_bets = await game.getCurrentCommits();
        const guess_commit = await game.commits(donor);

        assert.equal(curr_number_bets, 1, "Number of bets did not increment");
        assert.equal(guess_commit, hash, "Hashes do not match");
    })

    it("Should reveal hashed guess", async () => {
        const bet = await game.getStakeSize();
        const hash = Web3Utils.soliditySha3({type: 'string', value: "66"}, {type: 'string', value: "3"});

        //Commit/Reveal
        await setMaxPlayers(game, 1);

        await commitGuess(game, donor, "66", "3");

        //await game.commit(hash, { value: bet, from: donor });
        await revealGuess(game, donor, "66", "3");
        
        //await game.reveal("66", "3", {from: donor});
        
        const guess = await game.gameData(donor);
        //console.log(guess);

        assert.equal(guess, '66', "Revealed guesses do not match");
        //console.log(guess);
    })

    it("Should find winner and distribute prizes", async () => {
        // keccak256 , web3.sha3
        const bet = await game.getStakeSize();

        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

         await setMaxPlayers(game, 2);

        await commitGuess(game, accounts[2], "80", "3");
        await commitGuess(game, accounts[6], "20", "3");

        await revealGuess(game, accounts[2], "80", "3");   
        await revealGuess(game, accounts[6], "20", "3");

        await game.payout();

        // //console.log(hash)

        const winner = await game.getLastWinners(0);
        //console.log(winner);
        
        assert.equal(winner, accounts[6], "Winner isn't correctly selected");
    })

    it("Should play a full game with random input correctly", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 10;
        const bet = await game.getStakeSize();
        
        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        // Round 1
        await runGame(bet, num_players, accounts, game);
    })

    it("Should reset correctly", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 10;
        const bet = await game.getStakeSize();
        
        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        // Round 1-4
        await runGame(bet, num_players, accounts, game);
        await runGame(bet, num_players, accounts, game);
        await runGame(bet, num_players, accounts, game);
        //await runGame(bet, num_players, accounts, game);
    })

    it("Should go back to init", async () => {
        
        // MAX IS 10, because max account number is 10
        const bet = await getStakeSize(game);

        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });


         await setMaxPlayers(game, 3);

        await commitGuess(game, accounts[2], "30", "3");
        await commitGuess(game, accounts[6], "25", "3");

        var cur_bets = await getCurrentCommits(game);
        assert(cur_bets == 2, "Number of commits does not mathc");

        await game.reset();

        cur_bets = await getCurrentCommits(game);
        assert(cur_bets == 0, "State was not reset properly");

        await game.set_MAX_PLAYERS(2);
        await commitGuess(game, accounts[2], "30", "3");
        await commitGuess(game, accounts[6], "25", "3");

        await revealGuess(game, accounts[2], "30", "3");

        cur_bets = await getCurrentCommits(game);
        assert(cur_bets == 2, "Number of commits does not match in REVEAL_STATE");

        var cur_reveals = await getCurrentReveals(game);
        assert(cur_reveals == 1, "Number of reveals does not match in REVEAL_STATE");

        await game.reset();

        cur_bets = await getCurrentCommits(game);
        assert(cur_bets == 0, "failed to reset: commits");
        var cur_reveals = await getCurrentReveals(game);
        assert(cur_reveals == 0, "failed to reset: reveals");

    })

    it("Should make a correct payout", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 3;
        const bet = await getStakeSize(game);
        
        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        // Round 1-4
        await runGame(bet, num_players, accounts, game);

        var prize = await getPrizeAmount(game);
        var fee = await getGameFeeAmount(game);
      
        var expected_prize = (bet*num_players) - ((bet*num_players) / 100.0) * fee;
        assert(prize == expected_prize);
    })

    it("Everyone betting the same number", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 2;
        const bet = await getStakeSize(game);
        
        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });


        await setMaxPlayers(game, num_players);

        await commitGuess(game, accounts[2], "30", "3");
        await commitGuess(game, accounts[6], "30", "3");



        await revealGuess(game, accounts[2], "30", "3");
        await revealGuess(game, accounts[6], "30", "3");

        var payout1 = await getPayout(game, accounts[2]);

        var payout2 = await getPayout(game, accounts[6]);
        var prize = await getPrizeAmount(game);

        assert(payout1 == prize, "Wrong prize amount");
        assert(payout2 == prize, "Wrong prize amount");
    })
    
    it("Should not be pausable and unpausable", async () => {
        const num_players = 2;
        const bet = await getStakeSize(game);
        
        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });
        
        await setMaxPlayers(game, 1);

        await pauseGame(game);

        var error = false;
        try {
            await commitGuess(game, accounts[2], "30", "3");
            error = true;
        } catch (e) {
        }

        assert(error == false, "The commit executed in paused state");

        await unpauseGame(game);
        await commitGuess(game, accounts[2], "30", "3");
        await pauseGame(game);

        try {
             await revealGuess(game, accounts[2], "30", "3");
             error = true;
        } catch (e) {
        }

        assert(error == false, "The reveal executed in paused state");
    })

    it("Should emmit events and API should watch for them", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 2;
        const bet = await getStakeSize(game);
        
        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        watchEvent(game.CommitsSubmitted, function(args){}, ["All commits submitted"]);
        watchEvent(game.RevealsSubmitted, function(args){}, ['All reveals submitted']);


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




/////////////////////// HELPERS /////////////////////////

async function runGame(bet, num_players, accounts, game) {

    var guesses = createRandomGuesses(num_players, accounts);

    await game.set_MAX_PLAYERS(num_players);

    var i;
    for(i = 0; i < num_players; i++){
        //const hash = Web3Utils.soliditySha3({type: 'string', value: guesses[1][i].toString()}, {type: 'string', value: "3"});
        await commitGuess(game, accounts[i], guesses[1][i].toString(), "3");
    }

    var state = await isInRevealState(game);
    
    assert(state == true, "Bad state transition, should be in REVEAL_STATE");
    for(i = 0; i < num_players; i++){
        await revealGuess(game, accounts[i], guesses[1][i].toString(), "3");
    }

    state = await isInPayoutState(game);
    assert(state == true, "Bad state transition, should be in PAYOUT_STATE");

    // Lets check the balances
    for(i = 0; i < num_players; i++){
        var balance = web3.fromWei(web3.eth.getBalance(accounts[i]),'ether').toString()
        //console.log(balance);
    }

    await game.payout();

    var average = computeTwoThirdsAverage(guesses[1]);
    //console.log(average);

    var average23 = await game.average23();
    
    // DEBUG: Put this one back in. There is a known problem here.
    assert(Math.floor(average) == average23.toNumber(), "Average23 miscalculated...");


    var loc_winners = findWinner(accounts, guesses[1], average23);

    //console.log(loc_winners);

    // Grab all the winners

    var winners = await getWinners(game);


    var number_of_winners = winners.length;

    assert(loc_winners.length == number_of_winners, "Number of winners varies");

    for (i = 0; i < number_of_winners; i++){
        var winner = winners[i];

        assert(winner == loc_winners[i], "Wrong winner");
    }

    state = await isInCommitState(game);
    assert(state == true, "Bad state transition, should be in COMMIT_STATE");


    //console.log("Done.");
}




/////////////////////// TESTER FUNCTIONS /////////////////

const MAX_GUESS = 100;

function findWinner(player_addrs, guesses, avrg){
    
    var twothirdsavg = computeTwoThirdsAverage(guesses);

    var min_diff = MAX_GUESS;
    var cur_diff;

    var winner = [];
    for(i = 0; i < player_addrs.length; i++) {
        
        var cur_guess = guesses[i];

        // Find the difference between the guess and the average
        if(twothirdsavg > cur_guess){
            cur_diff = twothirdsavg - cur_guess;
        }
        else{
            cur_diff = cur_guess - twothirdsavg;
        }
        
        // If the difference is less than the smallest difference,
        // we delete all the winners and add the new candidate
        if(cur_diff < min_diff) {
            winners = [];
            winners.push(player_addrs[i]);
            min_diff = cur_diff;
        // Else, if the difference are the same, we add the candidate to the 
        // list of winners
        } else if(cur_diff == min_diff){
            winners.push(player_addrs[i]);
        }
    }

    return winners;
}

function createRandomGuesses(max_players, accounts){

    var guesses = new Array(2);
    guesses[0] = new Array(max_players);
    guesses[1] = new Array(max_players);
    
    var i;
    for(i = 0; i < max_players; i++){
        guesses[0][i] = accounts[i];
        guesses[1][i] = Math.floor(Math.random() * 101);
    }

    //console.log(guesses);
    return guesses;
}

function computeTwoThirdsAverage(guesses){
    var sum = guesses.reduce(function(acc, val) { return acc + val; });
    sum = sum * 10000;
    var average = Math.floor(sum / guesses.length);
    var average23 = Math.floor((average * 2) / 3);
    average23 = Math.floor(average23/10000)

    return  average23;    
}
