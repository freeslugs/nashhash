var Game = artifacts.require("TwoThirdsAverage");

module.exports = function(deployer) {
  deployer.deploy(Game, 10);
};
