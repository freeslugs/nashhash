var Web3Utils = require('web3-utils');
import API from '../src/api/Game.js';

var Game = artifacts.require("./LowestUniqueNum.sol");

contract("Lowest Unique Game", function([owner, donor]){

	var accounts;

	let game

	beforeEach('setup contract for each test', async () => {
        game = await Game.new(10);
    })

    it("init", async () => {
        const count = await API.getStakeSize(game);
        assert.equal(count, 1);
    });


    it("Should find winner and distribute prizes", async () => {
        // keccak256 , web3.sha3
        const bet = await API.getStakeSize(game);

        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        await API.setMaxPlayers(game, 2);

        await API.commitGuess(game, accounts[2], "3", "6534");
        await API.commitGuess(game, accounts[3], "5", "1004");
        //await commitGuess(game, accounts[8], "5", "7728");

        await API.revealGuess(game, accounts[2], "3", "6534");  
        await API.revealGuess(game, accounts[3], "5", "1004"); 
        //await revealGuess(game, accounts[8], "5", "7728");

        // console.log("HELLO");

        await game.payout();

        // //console.log(hash)

        const winner = await API.getWinners(game, 0);
        console.log(winner);
        
        assert.equal(winner, accounts[2], "Winner isn't correctly selected");
    })

});