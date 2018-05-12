var Web3Utils = require('web3-utils');

// var keccak256 = require('js-sha3').keccak256;


var Game = artifacts.require("./Game.sol");
  
contract("Game", function([owner, donor]){

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

        // const res = await game.getSha("66", "3");
        // console.log(res)
        
        //Commit/Reveal
        await game.commit(hash, { value: 1, from: donor });
        await game.reveal("66", "3", {from: donor});
        
        const guess = await game.game_data(donor);
        console.log(guess);

        assert.equal(guess, '66', "Revealed guesses do not match");
        console.log(guess);
    })



});
