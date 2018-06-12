var Web3Utils = require('web3-utils');
import API from '../src/api/Game.js';
import Helper from './gameHelper.js';

const FIXED_BET = web3.toWei(1,'ether');
const MAX_PLAYERS = 10;
const HASHNASH_ADDRESS = 0x2540099e9ed04aF369d557a40da2D8f9c2ab928D;
const GAME_STAGE_LENGTH = 0;
const GAME_FEE_PERCENT = 5;

var Game = artifacts.require("./TwoThirdsAverage.sol");

let api, helper

contract("2/3 of the Average Game", function([owner, donor]){

    let accounts, game

    beforeEach('setup contract for each test', async () => {
        game = await Game.new(HASHNASH_ADDRESS,
            GAME_FEE_PERCENT,
            FIXED_BET,
            MAX_PLAYERS,
            GAME_STAGE_LENGTH);
        // init api
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

    it("Single commit and single reveal", async () => {
       
        //Commit/Reveal
        await api.setMaxPlayers(1, owner);

        await api.commitGuess(donor, "66", "3");
        await api.revealGuess(donor, "66", "3");
        
        const guess = await game.gameData(donor);

        assert.equal(guess, '66', "Revealed guesses do not match");
    })

    it("Winner is found correclty, prizes distributed correctly. Two players, different guesses", async () => {
        
        await api.setMaxPlayers(2, owner);

        await api.commitGuess(accounts[2], "80", "3");
        await api.commitGuess(accounts[6], "20", "3");

        await api.revealGuess(accounts[2], "80", "3");   
        await api.revealGuess(accounts[6], "20", "3");

        await api.payout();

        const winners = await api.getWinners();
        
        assert.equal(winners[0], accounts[6], "Winner isn't correctly selected");
    })

    it("Plays the full game, 10 players, random guesses", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 10;
        const bet = await api.getStakeSize();

        // Round 1
        await helper.runGame(bet, num_players, accounts, owner);
    })

    it("Multiple rounds of the game are played correclty", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 10;
        const bet = await api.getStakeSize();
        
        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        // Round 1-4
        await helper.runGame(bet, num_players, accounts, owner);
        await helper.runGame(bet, num_players, accounts, owner);

    })

    it("resetGame() works correcly", async () => {
        

        await api.setMaxPlayers(3, owner);

        await api.commitGuess(accounts[2], "30", "3");
        await api.commitGuess(accounts[6], "25", "3");

        var cur_bets = await api.getCurrentCommits();
        assert(cur_bets == 2, "Number of commits does not mathc");

        await api.resetGame(owner);

        cur_bets = await api.getCurrentCommits();
        assert(cur_bets == 0, "State was not reset properly");

        await api.setMaxPlayers(2, owner);
        await api.commitGuess(accounts[2], "30", "3");
        await api.commitGuess(accounts[6], "25", "3");

        await api.revealGuess(accounts[2], "30", "3");

        cur_bets = await api.getCurrentCommits();
        assert(cur_bets == 2, "Number of commits does not match in REVEAL_STATE");

        var cur_reveals = await api.getCurrentReveals();
        assert(cur_reveals == 1, "Number of reveals does not match in REVEAL_STATE");

        await api.resetGame(owner);

        cur_bets = await api.getCurrentCommits();
        assert(cur_bets == 0, "failed to reset: commits");
        var cur_reveals = await api.getCurrentReveals();
        assert(cur_reveals == 0, "failed to reset: reveals");

    })

    it("Makes a correct payout", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 3;
        const bet = await api.getStakeSize();

        // Round 1-4
        await helper.runGame(bet, num_players, accounts, owner);

        var prize = await api.getPrizeAmount();
        var fee = await api.getGameFeeAmount();

        var numWinners = await api.getNumberOfWinners();
      
        var expectedTotalPrize = (bet*num_players) - ((bet*num_players) / 100.0) * fee;
        var expectedPrize = expectedTotalPrize / numWinners;

        assert(prize == expectedPrize, `prizes are not equal, prize: ${prize} fee: ${fee}, bet: ${bet}`);
    })

    it("2 players bet, one wins, one loses", async () => {
        // MAX IS 10, because max account number is 10
        const num_players = 2;
        
        await api.setMaxPlayers(num_players, owner);

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

    it("2 players betting the same number", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 2;
        
        await api.setMaxPlayers(num_players, owner);

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
    
    it("Can be paused and unpaused", async () => {
        const num_players = 2;
        const bet = await api.getStakeSize();
        
        await api.setMaxPlayers(1, owner);
        await api.pauseGame();

        await helper.expectedThrow(api.commitGuess(accounts[2], "30", "3"), "The commit executed in paused state");

        await api.unpauseGame();
        await api.commitGuess(accounts[2], "30", "3");
        await api.pauseGame();

        await helper.expectedThrow(api.revealGuess(accounts[2], "30", "3"), "The reveal executed in paused state");
    })

    // it("TODO: Emit event", async () => {
    //     // MAX IS 10, because max account number is 10
    //     const num_players = 2;
    //     const bet = await api.getStakeSize();
        
    //     var error = 1;

    //     api.watchEvent(api.game.CommitsSubmitted, function(args){error = 0; console.log("Commits received");}, ['All commits submitted']);
    //     api.watchEvent(api.game.RevealsSubmitted, function(args){error = 0; console.log("hi");}, ['All reveals submitted']);

    //     await api.commitGuess(accounts[2], "30", "3");
    //     await api.commitGuess(accounts[6], "30", "3");

    //     await helper.sleep(30000);

    //     assert(error = 0, "CommitsSubmitted event not emmited");

    //     error = 1;

    //     assert(error = 0, "RevealsSubmitted event not emmited");
    // })

    it("Should be forcable into states by owner", async () => {
        const num_players = 3;

        await api.setMaxPlayers(num_players, owner);

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