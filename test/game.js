var Web3Utils = require('web3-utils');
import API from '../src/api/Game.js';
//var api = require('../src/api/Game.js');

// var keccak256 = require('js-sha3').keccak256;
const FIXED_BET = web3.toWei(1,'ether');
const MAX_PLAYERS = 10;
const HASHNASH_ADDRESS = 0x2540099e9ed04aF369d557a40da2D8f9c2ab928D;

var Game = artifacts.require("./TwoThirdsAverage.sol");
  
let api

contract("2/3 of the Average Game", function([owner, donor]){

    var accounts;

    let game

    beforeEach('setup contract for each test', async () => {
        game = await Game.new(10);
        // init api
        api = new API(web3, assert, game);

        // watchEvent(game.RevealsSubmitted, function(){game.payout();} );
    })

    it("init", async () => {
        const count = await api.getStakeSize();
        // console.log(count);
        assert.equal(count, 1);
    });

    it("Should commit hashed guess with stake", async () => {
        
        const hash = api.hashGuess("66", "3");

        await api.commitGuess(donor, "66", "3");

        const curr_number_bets = await game.getCurrentCommits();
        const guess_commit = await game.commits(donor);

        assert.equal(curr_number_bets, 1, "Number of bets did not increment");
        assert.equal(guess_commit, hash, "Hashes do not match");
    })

    it("Should reveal hashed guess", async () => {
        const bet = await api.getStakeSize();
        const hash = Web3Utils.soliditySha3({type: 'string', value: "66"}, {type: 'string', value: "3"});

        //Commit/Reveal
        await api.setMaxPlayers(1);

        await api.commitGuess(donor, "66", "3");

        //await game.commit(hash, { value: bet, from: donor });
        await api.revealGuess(donor, "66", "3");
        
        //await game.reveal("66", "3", {from: donor});
        
        const guess = await game.gameData(donor);
        //console.log(guess);

        assert.equal(guess, '66', "Revealed guesses do not match");
        //console.log(guess);
    })

    it("Should find winner and distribute prizes", async () => {
        // keccak256 , web3.sha3
        const bet = await api.getStakeSize();

        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        await api.setMaxPlayers(2);

        await api.commitGuess(accounts[2], "80", "3");
        await api.commitGuess(accounts[6], "20", "3");

        await api.revealGuess(accounts[2], "80", "3");   
        await api.revealGuess(accounts[6], "20", "3");

        await api.payout();

        // //console.log(hash)

        const winners = await api.getWinners();
        //console.log(winner);
        
        assert.equal(winners[0], accounts[6], "Winner isn't correctly selected");
    })

    it("Should play a full game with random input correctly", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 10;
        const bet = await api.getStakeSize();
        
        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        // Round 1
        await runGame(bet, num_players, accounts, game);
    })

    it("Should reset correctly", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 10;
        const bet = await api.getStakeSize();
        
        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        // Round 1-4
        await runGame(bet, num_players, accounts, game);
        await runGame(bet, num_players, accounts, game);
        // await runGame(bet, num_players, accounts, game);
        //await runGame(bet, num_players, accounts, game);
    })

    it("Should go back to init", async () => {
        
        // MAX IS 10, because max account number is 10
        const bet = await api.getStakeSize();

        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });


        await api.setMaxPlayers(3);

        await api.commitGuess(accounts[2], "30", "3");
        await api.commitGuess(accounts[6], "25", "3");

        var cur_bets = await api.getCurrentCommits();
        assert(cur_bets == 2, "Number of commits does not mathc");

        await api.resetGame();

        cur_bets = await api.getCurrentCommits();
        assert(cur_bets == 0, "State was not reset properly");

        await api.setMaxPlayers(2);
        await api.commitGuess(accounts[2], "30", "3");
        await api.commitGuess(accounts[6], "25", "3");

        await api.revealGuess(accounts[2], "30", "3");

        cur_bets = await api.getCurrentCommits();
        assert(cur_bets == 2, "Number of commits does not match in REVEAL_STATE");

        var cur_reveals = await api.getCurrentReveals();
        assert(cur_reveals == 1, "Number of reveals does not match in REVEAL_STATE");

        await api.resetGame();

        cur_bets = await api.getCurrentCommits();
        assert(cur_bets == 0, "failed to reset: commits");
        var cur_reveals = await api.getCurrentReveals();
        assert(cur_reveals == 0, "failed to reset: reveals");

    })

    it("Should make a correct payout", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 3;
        const bet = await api.getStakeSize();
        
        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        // Round 1-4
        await runGame(bet, num_players, accounts, game);

        var prize = await api.getPrizeAmount();
        var fee = await api.getGameFeeAmount();
      
        var expected_prize = (bet*num_players) - ((bet*num_players) / 100.0) * fee;
        assert(prize == expected_prize);
    })

    it("2 players bet, one wins, one loses", async () => {
        // MAX IS 10, because max account number is 10
        const num_players = 2;
        const bet = await api.getStakeSize();
        
        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        await api.setMaxPlayers(num_players);

        await api.commitGuess(accounts[2], "100", "3");
        await api.commitGuess(accounts[6], "2", "3");



        await api.revealGuess(accounts[2], "100", "3");
        await api.revealGuess(accounts[6], "2", "3");

        var payout1 = await api.getPayout(accounts[2]);

        var payout2 = await api.getPayout(accounts[6]);
        var prize = await api.getPrizeAmount();

        assert(payout1 == 0, "Wrong prize amount");
        assert(payout2 == prize, "Wrong prize amount");
    })

    it("Everyone betting the same number", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 2;
        const bet = await api.getStakeSize();
        
        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        await api.setMaxPlayers(num_players);

        await api.commitGuess(accounts[2], "30", "3");
        await api.commitGuess(accounts[6], "30", "3");



        await api.revealGuess(accounts[2], "30", "3");
        await api.revealGuess(accounts[6], "30", "3");

        var payout1 = await api.getPayout(accounts[2]);

        var payout2 = await api.getPayout(accounts[6]);
        var prize = await api.getPrizeAmount();

        assert(payout1 == prize, "Wrong prize amount");
        assert(payout2 == prize, "Wrong prize amount");
    })
    
    it("Should not be pausable and unpausable", async () => {
        const num_players = 2;
        const bet = await api.getStakeSize();
        
        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });
        
        await api.setMaxPlayers(1);

        await api.pauseGame();

        var error = false;
        try {
            await api.commitGuess(accounts[2], "30", "3");
            error = true;
        } catch (e) {
        }

        assert(error == false, "The commit executed in paused state");

        await api.unpauseGame();
        await api.commitGuess(accounts[2], "30", "3");
        await api.pauseGame();

        try {
             await api.revealGuess(accounts[2], "30", "3");
             error = true;
        } catch (e) {
        }

        assert(error == false, "The reveal executed in paused state");
    })

    it("Should emmit events and api should watch for them", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 2;
        const bet = await api.getStakeSize();
        
        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        // todo: 
        // api.watchEvent(game.CommitsSubmitted).then(() => console.log("All commits submitted"))
        // api.watchEvent(game.RevealsSubmitted).then(() => console.log("All reveals submitted"))
        api.watchEvent(game.CommitsSubmitted, function(args){}, ['All commits submitted']);
        api.watchEvent(game.RevealsSubmitted, function(args){}, ['All reveals submitted']);
    })

    it("Should be forcable into states by owner", async () => {
        const num_players = 3;
        const bet = await api.getStakeSize();
        
        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        await api.setMaxPlayers(num_players);

        await api.commitGuess(accounts[2], "45", "3");
        await api.commitGuess(accounts[6], "30", "3");

        await api.forceToRevealState();

        var commitNumber = await api.getCurrentCommits();
        assert(commitNumber == 2, "Wrong commit number");

        await api.revealGuess(accounts[2], "45", "3");

        await api.forceToPayoutState();

        await api.payout();

        var winners = await api.getWinners();

        assert(winners.length == 1, winners.length);
        assert(winners[0] == accounts[2], winners[0]);
    })



});

/////////////////////// HELPERS /////////////////////////

async function runGame(bet, num_players, accounts) {

    var guesses = createRandomGuesses(num_players, accounts);

    await api.setMaxPlayers(num_players);

    for(var i = 0; i < num_players; i++){
        //const hash = Web3Utils.soliditySha3({type: 'string', value: guesses[1][i].toString()}, {type: 'string', value: "3"});
        await api.commitGuess(accounts[i], guesses[i].toString(), "3");
    }

    var state = await api.isInRevealState();
    
    assert(state == true, "Bad state transition, should be in REVEAL_STATE");
    for(i = 0; i < num_players; i++){
        await api.revealGuess(accounts[i], guesses[i].toString(), "3");
    }

    state = await api.isInPayoutState();
    assert(state == true, "Bad state transition, should be in PAYOUT_STATE");

    // Uncomment to check the balances
    for(i = 0; i < num_players; i++){
        var balance = web3.fromWei(web3.eth.getBalance(accounts[i]),'ether').toString()
        //console.log(balance);
    }

    await api.payout();

    var average = computeTwoThirdsAverage(guesses);
    //console.log(average);

    var average23 = await api.game.average23();
    
    // DEBUG: Put this one back in. There is a known problem here.
    assert(Math.floor(average) == average23.toNumber(), "Average23 miscalculated...");


    var loc_winners = findWinner(accounts, guesses, average23);

    //console.log(loc_winners);

    // Grab all the winners

    var winners = await api.getWinners();


    var number_of_winners = winners.length;

    assert(loc_winners.length == number_of_winners, "Number of winners varies");

    for (i = 0; i < number_of_winners; i++){
        var winner = winners[i];

        assert(winner == loc_winners[i], "Wrong winner");
    }

    state = await api.isInCommitState();
    assert(state == true, "Bad state transition, should be in COMMIT_STATE");


    //console.log("Done.");
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}




/////////////////////// TESTER FUNCTIONS /////////////////

const MAX_GUESS = 100;

function findWinner(player_addrs, guesses, avrg){
    
    var min_diff = MAX_GUESS;
    var cur_diff;

    let winners = [];
    for(var i = 0; i < player_addrs.length; i++) {
        
        var cur_guess = guesses[i];

        // Find the difference between the guess and the average
        if(avrg > cur_guess){
            cur_diff = avrg - cur_guess;
        }
        else{
            cur_diff = cur_guess - avrg;
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

    var guesses = new Array(max_players);
    for(var i = 0; i < max_players; i++){
        guesses[i] = Math.floor(Math.random() * 101);

    }

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


