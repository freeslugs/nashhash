var Web3Utils = require('web3-utils');
import API from '../src/api/Game.js';

var Game = artifacts.require("./LowestUniqueNum.sol");

let api

contract("Lowest Unique Game", function([owner, donor]){

	var accounts;

	let game

	beforeEach('setup contract for each test', async () => {
        game = await Game.new(10);

        api = new API(web3, assert, game);

        accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });
    })

    it("init", async () => {
        const count = await api.getStakeSize(game);
        assert.equal(count, 1);
    });


    it("Single commit", async () => {
        
        const hash = api.hashGuess("66", "3");

        await api.commitGuess(donor, "66", "3");

        const curr_number_bets = await api.getCurrentCommits();
        const guess_commit = await game.commits(donor);

        assert.equal(curr_number_bets, 1, "Number of bets did not increment");
        assert.equal(guess_commit, hash, "Hashes do not match");
    })

    it("Single commit and single reveal", async () => {
       
        //Commit/Reveal
        await api.setMaxPlayers(1);

        await api.commitGuess(donor, "66", "3");

        await api.revealGuess(donor, "66", "3");
        
        const guess = await game.gameData(donor);


        assert.equal(guess, '66', "Revealed guesses do not match");

    })


    it("Three players, two pick a common low number", async () => {
        // keccak256 , web3.sha3
        const bet = await api.getStakeSize(game);

        await api.setMaxPlayers(3);

        await api.commitGuess( accounts[2], "3", "6534");
        await api.commitGuess( accounts[3], "3", "1004");
        await api.commitGuess( accounts[8], "5", "7728");

        await api.revealGuess( accounts[2], "3", "6534");  
        await api.revealGuess( accounts[3], "3", "1004"); 
        await api.revealGuess( accounts[8], "5", "7728");


        await game.payout();

        const winner = await api.getWinners(0);
        
        assert.equal(winner, accounts[8], "Winner isn't correctly selected");
    })

    it("Should play a full game with random input correctly", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 10;
        const bet = await game.getStakeSize();

        // Round 1
        await runGame(bet, num_players, accounts, game);
    })

    it("Should play multiple rounds of game correctly", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 10;
        const bet = await api.getStakeSize(game);

        // Round 1-4
        await runGame(bet, num_players, accounts, game);
        await runGame(bet, num_players, accounts, game);
        await runGame(bet, num_players, accounts, game);
        await runGame(bet, num_players, accounts, game);
    })


    it("2 players bet, one wins, one loses", async () => {
        // MAX IS 10, because max account number is 10
        const num_players = 2;
        
        await api.setMaxPlayers(num_players);

        await api.commitGuess(accounts[2], "6", "3");
        await api.commitGuess(accounts[6], "2", "3");

        await api.revealGuess(accounts[2], "6", "3");
        await api.revealGuess(accounts[6], "2", "3");

        var payout1 = await api.getPayout(accounts[2]);

        var payout2 = await api.getPayout(accounts[6]);
        var prize = await api.getPrizeAmount();

        assert(payout1 == 0, "Wrong prize amount");
        assert(payout2 == prize, "Wrong prize amount");
    })

    it("Should fail if same user commits or reveals twice", async () => {

        const bet = await api.getStakeSize(game);
        await api.setMaxPlayers(1);

        await api.commitGuess(accounts[2], "6", "6534");

        var comtwice = false;
        try {
            await api.commitGuess(accounts[2], "6", "6534");
            comtwice = true;
        } catch (e) {
        }

        assert(comtwice == false, "The user committed twice successfully");
        
        await api.revealGuess(accounts[2], "6", "6534");

        var revtwice = false;
        try {
            await api.revealGuess(accounts[2], "6", "6534");
            revtwice = true;
        } catch (e) {
        }

        assert(revtwice == false, "The user revealed twice successfully");
    })

    it("Should fail if player reveals before commit state is over", async () => {

        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        const bet = await api.getStakeSize(game);
        await api.setMaxPlayers(3);

        await api.commitGuess(accounts[2], "6", "6534");

        var reverror = false;
        try {
            await api.revealGuess(accounts[2], "17", "6534");
            reverror = true;
        } catch (e) {
        }

        assert(reverror == false, "The user revealed before commit state was over.");
    })


    it("Should fail if commit and reveal values differ.", async () => {

        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        const bet = await api.getStakeSize(game);
        await api.setMaxPlayers(3);

        await api.commitGuess( accounts[2], "3", "6534");
        await api.commitGuess( accounts[3], "3", "1004");
        await api.commitGuess( accounts[8], "5", "7728");

        var differror = false;
        try {
            await api.revealGuess(accounts[2], "17", "6534");
            differror = true;
        } catch (e) {
        }

        assert(differror == false, "The user committed and revealed different numbers.");
    })


});




/////////////////////// HELPERS /////////////////////////

async function runGame(bet, num_players, accounts, game) {

    var guesses = createRandomGuesses(num_players, accounts);

    await api.setMaxPlayers( num_players);

    for(var i = 0; i < num_players; i++){
        //const hash = Web3Utils.soliditySha3({type: 'string', value: guesses[1][i].toString()}, {type: 'string', value: "3"});
        await api.commitGuess( accounts[i], guesses[i].toString(), "3");
    }

    var state = await api.isInRevealState(game);
    assert(state == true, "Bad state transition, should be in REVEAL_STATE");
    for(i = 0; i < num_players; i++){
        await api.revealGuess( accounts[i], guesses[i].toString(), "3");
    }

    state = await api.isInPayoutState(game);
    assert(state == true, "Bad state transition, should be in PAYOUT_STATE");

    // Uncomment to check the balances
/*    for(i = 0; i < num_players; i++){
        var balance = web3.fromWei(web3.eth.getBalance(accounts[i]),'ether').toString()
        console.log(balance);
    }*/

    await game.payout();

    var gamelowest = findLowestUniqueNum(guesses);

    var realLowest = await game.testLowest();

    var loc_winner = findWinner(accounts, guesses, realLowest);
    // Grab all the winners

    var winner = await api.getWinners(game);

    assert(winner.length == 1, "More than one winner");

    assert(winner[0] == loc_winner, "Winner is incorrectly chosen")

    state = await api.isInCommitState(game);
    assert(state == true, "Bad state transition, should be in COMMIT_STATE");
}

function findWinner(player_addrs, guesses, lwst){

    var winner; 

    for(let i = 0; i < player_addrs.length; i++) {
        
        var cur_guess = guesses[i];

        //If current guess is lowest unique number, set winner to guesser address
        if(cur_guess == lwst){
        	winner = player_addrs[i];
        }
    }

    return winner;
}

function createRandomGuesses(max_players, accounts){

    var guesses = new Array(max_players);
    for(var i = 0; i < max_players; i++){
        guesses[i] = Math.floor(Math.random() * 1000);

    }

    return guesses;
}

function findLowestUniqueNum(guesses){
	//Initialize empty sets
	var unique = new Set([]);
	var notUnique = new Set([]);


	for(var i = 0; i < guesses.length; i++){
		if(unique.has(guesses[i])){
			unique.delete(guesses[i]);
			notUnique.add(guesses[i]);
		} else if(!notUnique.has(guesses[i])){
			unique.add(guesses[i]);
		}
	}

	var lowest;
    for(let num of unique.values()){
    	if(num < lowest){
    		lowest = num;
    	}
    }  

    return lowest;
}
