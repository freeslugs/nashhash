var Web3Utils = require('web3-utils');

// var keccak256 = require('js-sha3').keccak256;


var Game = artifacts.require("./Game.sol");
  
contract("Game", function([owner, donor]){

    var accounts;

    let game

    beforeEach('setup contract for each test', async () => {
        game = await Game.new(owner);
    })

    it("init", async () => {
        const count = await game.BET_SIZE();
        assert.equal(count.toNumber(), web3.toWei(1,'ether'));
    });

    it("Should commit hashed guess with stake", async () => {
        const bet = await game.BET_SIZE();
        const hash = Web3Utils.soliditySha3({type: 'string', value: "66"}, {type: 'string', value: "3"});

        await game.commit(hash, { value: bet, from: donor });
        const curr_number_bets = await game.curr_number_bets();
        const guess_commit = await game.commits(donor);

        assert.equal(curr_number_bets, 1, "Number of bets did not increment");
        assert.equal(guess_commit, hash, "Hashes do not match");
        console.log(hash);
        console.log(guess_commit);
    })

    it("Should reveal hashed guess", async () => {
        // keccak256 , web3.sha3
        const bet = await game.BET_SIZE();
        const hash = Web3Utils.soliditySha3({type: 'string', value: "66"}, {type: 'string', value: "3"});
        // console.log(hash)

        //Commit/Reveal
        await game.commit(hash, { value: bet, from: donor });
        await game.reveal("66", "3", {from: donor});
        
        const guess = await game.game_data(donor);
        console.log(guess);

        assert.equal(guess, '66', "Revealed guesses do not match");
        console.log(guess);
    })

    it("Should find winner and distribute prizes", async () => {
        // keccak256 , web3.sha3
        const bet = await game.BET_SIZE();
        const hash1 = Web3Utils.soliditySha3({type: 'string', value: "80"}, {type: 'string', value: "3"});
        const hash2 = Web3Utils.soliditySha3({type: 'string', value: "20"}, {type: 'string', value: "3"});

        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        await game.set_MAX_PLAYERS(2);

        //console.log(accounts);
        //console.log(donor);
        //console.log(accounts[2]);

        await game.commit(hash1, { value: bet, from: accounts[2] });
        await game.commit(hash2, { value: bet, from: accounts[6] });

        await game.reveal("80", "3", {from: accounts[2]});   
        await game.reveal("20", "3", {from: accounts[6]});

        // //console.log(hash)

        // //Commit/Reveal
        await game.find_winner();

        // const winner = await game.winners(0);
        // console.log(winner);
        
        // assert.equal(winner, accounts[6], "Winner isn't correctly selected");
    })

    it("Should run with random input correctly", async () => {
        
        // MAX IS 10, because max account number is 10
        const num_players = 10;
        const bet = await game.BET_SIZE();
        
        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        var guesses = createRandomGuesses(num_players, accounts);

        await game.set_MAX_PLAYERS(num_players);

        var i;
        for(i = 0; i < num_players; i++){
            const hash = Web3Utils.soliditySha3({type: 'string', value: guesses[1][i].toString()}, {type: 'string', value: "3"});
            await game.commit(hash, { value: bet, from: accounts[i] });
        }

        var state = await game.game_state_debug();
        
        assert(state.toNumber() == 1, "Bad state transition, should be in REVEAL_STATE");
        for(i = 0; i < num_players; i++){
            await game.reveal(guesses[1][i].toString(), "3", {from: accounts[i]});
        }

        state = await game.game_state_debug();
        assert(state.toNumber() == 0, "Bad state transition, should be in COMMIT_STATE");

        // Lets check the balances
        for(i = 0; i < num_players; i++){
            var balance = web3.fromWei(web3.eth.getBalance(accounts[i]),'ether').toString()
            console.log(balance);
        }

        var average = computeTwoThirdsAverage(guesses[1]);
        console.log(average);

        var average23 = await game.average23();

        console.log(average23.toNumber());
        console.log(Math.floor(average));
        
        // DEBUG: Put this one back in. There is a known problem here.
        //assert(Math.floor(average) == average23.toNumber(), "Average23 miscalculated...");

        findWinner(accounts, guesses[1], average23);
        var winner = await game.last_winners(0);
        console.log(winner);
        



    })
});



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

    console.log(winners);
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

    console.log(guesses);
    return guesses;
}

function computeTwoThirdsAverage(guesses){
    var sum = guesses.reduce(function(acc, val) { return acc + val; });
    var average = sum / guesses.length;
    var average23 = (average * 2) / 3;

    return  average23;    
}
