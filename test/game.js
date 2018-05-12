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
        assert.equal(count, 1);
    });

    it("Should commit hashed guess with stake", async () => {
        const hash = Web3Utils.soliditySha3({type: 'string', value: "66"}, {type: 'string', value: "3"});

        await game.commit(hash, { value: 1, from: donor });
        const curr_number_bets = await game.curr_number_bets();
        const guess_commit = await game.commits(donor);

        assert.equal(curr_number_bets, 1, "Number of bets did not increment");
        assert.equal(guess_commit, hash, "Hashes do not match");
        console.log(hash);
        console.log(guess_commit);
    })

    it("Should reveal hashed guess", async () => {
        // keccak256 , web3.sha3
        const hash = Web3Utils.soliditySha3({type: 'string', value: "66"}, {type: 'string', value: "3"});
        // console.log(hash)

        //Commit/Reveal
        await game.commit(hash, { value: 1, from: donor });
        await game.reveal("66", "3", {from: donor});
        
        const guess = await game.game_data(donor);
        console.log(guess);

        assert.equal(guess, '66', "Revealed guesses do not match");
        console.log(guess);
    })

    it("Should find winner and distribute prizes", async () => {
        // keccak256 , web3.sha3
        const hash1 = Web3Utils.soliditySha3({type: 'string', value: "80"}, {type: 'string', value: "3"});
        const hash2 = Web3Utils.soliditySha3({type: 'string', value: "20"}, {type: 'string', value: "3"});

        const accounts = await new Promise(function(resolve, reject) {
            web3.eth.getAccounts( function (err, accounts) { resolve(accounts) })
        });

        await game.set_MAX_PLAYERS(2);

        console.log(accounts);
        console.log(donor);
        console.log(accounts[2]);

        await game.commit(hash1, { value: 1, from: accounts[2] });
        await game.commit(hash2, { value: 1, from: accounts[6] });

        await game.reveal("80", "3", {from: accounts[2]});   
        await game.reveal("20", "3", {from: accounts[6]});

        // //console.log(hash)

        // //Commit/Reveal
        await game.find_winner();

        // const winner = await game.winners(0);
        // console.log(winner);
        
        // assert.equal(winner, accounts[6], "Winner isn't correctly selected");
    })
});
