var Web3Utils = require('web3-utils');

var Game = artifacts.require("./LowestUniqueNum.sol");

import API from '../src/api/Game.js'

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

        await API.setMaxPlayers(game, 3);

        await API.commitGuess(game, accounts[2], "3", "6534");
        await API.commitGuess(game, accounts[3], "3", "1004");
        await API.commitGuess(game, accounts[8], "5", "7728");

        await API.revealGuess(game, accounts[2], "3", "6534");  
        await API.revealGuess(game, accounts[3], "3", "1004"); 
        await API.revealGuess(game, accounts[8], "5", "7728");

        // console.log("HELLO");

        await game.payout();

        // //console.log(hash)

        const winner = await game.getLastWinners(0);
        console.log(winner);
        
        assert.equal(winner, accounts[8], "Winner isn't correctly selected");
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

});




/////////////////////// HELPERS /////////////////////////

async function runGame(bet, num_players, accounts, game) {

    var guesses = createRandomGuesses(num_players, accounts);

    await API.setMaxPlayers(game, num_players);

    for(var i = 0; i < num_players; i++){
        //const hash = Web3Utils.soliditySha3({type: 'string', value: guesses[1][i].toString()}, {type: 'string', value: "3"});
        await API.commitGuess(game, accounts[i], guesses[i].toString(), "3");
    }

    var state = await API.isInRevealState(game);
    assert(state == true, "Bad state transition, should be in REVEAL_STATE");
    for(i = 0; i < num_players; i++){
        await API.revealGuess(game, accounts[i], guesses[i].toString(), "3");
    }

    state = await API.isInPayoutState(game);
    assert(state == true, "Bad state transition, should be in PAYOUT_STATE");

    // Lets check the balances
    for(i = 0; i < num_players; i++){
        var balance = web3.fromWei(web3.eth.getBalance(accounts[i]),'ether').toString()
    }

    await game.payout();

    var gamelowest = findLowestUniqueNum(guesses);

    var realLowest = await game.testLowest();

    var loc_winner = findWinner(accounts, guesses, realLowest);

    // Grab all the winners

    var winner = await API.getWinners(game);

    assert(winner.length == 1, "More than one winner");

    assert(winner[0] == loc_winner, "Winner is incorrectly chosen")

    state = await API.isInCommitState(game);
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



