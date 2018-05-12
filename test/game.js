var Game = artifacts.require("./Game.sol");
  
contract("Game", function(accounts){

    it("init with two candidates", function(){
        return Game.deployed().then(function(instance) {
            return instance.BET_SIZE();
        }).then( function(count) {
            assert.equal(count, 1);
        });
    });
});