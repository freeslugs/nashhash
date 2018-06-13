/* var NPT = artifacts.require("NPT");
var GameFactory = artifacts.require("GameFactory");
var Game = artifacts.require("Game");

module.exports = function(deployer) {
  deployer.deploy(NPT).then(async () => {
  	const npt = await NPT.at(NPT.address);
	const gamefactory = await GameFactory.deployed();

	const gameAdresses = await gamefactory.getGames();

	for(var i = 0; i < gameAdresses.length; i++){
		Game curGame = Game.at(gameAdresses[i]);
		curGame.setNPTAddress(NPT.address);

		npt.addMinter(gameAdresses[i]);
	}

  })
};
*/