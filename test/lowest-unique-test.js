var Web3Utils = require('web3-utils');
import API from '../src/api/Game.js';
import Helper from './lowestUniqueHelper.js';

const FIXED_BET = web3.toWei(1,'ether');
const MAX_PLAYERS = 10;
const HASHNASH_ADDRESS = 0x2540099e9ed04aF369d557a40da2D8f9c2ab928D;
const GAME_STAGE_LENGTH = 0;
const GAME_FEE_PERCENT = 5;

var Game = artifacts.require("./LowestUniqueNum.sol");
var NPT = artifacts.require("./NPT.sol");

let api, helper

contract("Lowest Unique Game", function([owner, donor]){

    let accounts, game, npt

    beforeEach('setup contract for each test', async () => {
        npt = await NPT.new();

        game = await Game.new(HASHNASH_ADDRESS,
            GAME_FEE_PERCENT,
            FIXED_BET,
            MAX_PLAYERS,
            GAME_STAGE_LENGTH,
            npt.address);

        //Give game minting permission
        npt.addMinter(game.address);

        api = new API(web3, assert, game);
        helper = new Helper(web3, assert, game, api);

        accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });
    })

    it("init", async () => {
        const count = await api.getStakeSize();
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


    it("Nashpoints correctly transferred for single commit", async () => {
        
        const hash = api.hashGuess("66", "3");
        await api.commitGuess(donor, "66", "3");

        /*
        var nptMintEvent = npt.Mint();
        console.log("to\t\tamount");
        nptMintEvent.watch(function(error, result){
            console.log(result.args._to + "\t" + 
            result.args._amount + "\t");
        });

        var nptTransferEvent = npt.Transfer();
        console.log("address\t\tfrom\t\tto\t\tamount");
        nptTransferEvent.watch(function(error, result){
            console.log(result.address + "\t" + result.args._from + "\t" + result.args._to + "\t" + 
            result.args._amount / 1e16 + "\t");
        });
        */

        const total_nashpoints = await npt.totalSupply();

        const nashpoints_awarded = await npt.balanceOf(donor);

        assert.equal(total_nashpoints.toNumber(), 10, "Incorrect number of Nashpoints minted");
        assert.equal(nashpoints_awarded.toNumber(), 10, "Incorrect number of Nashpoints awarded");
    })


    it("Single commit and single reveal", async () => {
       
        //Commit/Reveal
        await api.setMaxPlayers(1, owner);

        await api.commitGuess(donor, "66", "3");

        await api.revealGuess(donor, "66", "3");
        
        const guess = await game.gameData(donor);


        assert.equal(guess, '66', "Revealed guesses do not match");

    })


    it("Three players, two pick a common low number", async () => {
        // keccak256 , web3.sha3
        const bet = await api.getStakeSize(game);

        await api.setMaxPlayers(3, owner);

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
        await helper.runGame(bet, num_players, accounts, game, owner);
    })

    it("Should play multiple rounds of game correctly", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 10;
        const bet = await api.getStakeSize(game);

        // Round 1-4
        await helper.runGame(bet, num_players, accounts, game, owner);
        await helper.runGame(bet, num_players, accounts, game, owner);
        await helper.runGame(bet, num_players, accounts, game, owner);
        await helper.runGame(bet, num_players, accounts, game, owner);
    })


    it("2 players bet, one wins, one loses", async () => {
        // MAX IS 10, because max account number is 10
        const num_players = 2;
        
        await api.setMaxPlayers(num_players, owner);

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
        await api.setMaxPlayers(1, owner);

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
        await api.setMaxPlayers(3, owner);

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
        await api.setMaxPlayers(3, owner);

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