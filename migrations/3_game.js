var GameFactory = artifacts.require("GameFactory");
var Web3Utils = require('web3-utils');

const TWO_THIRDS_TYP = "TTA";
const LUN_TYP = "LUN";
const BET_1 = web3.toWei(1,'ether');
const BET_01 = web3.toWei(0.1, 'ether');
const BET_001 = web3.toWei(0.01, 'ether');
const MAX_PLAYERS = 10;

const HASHNASH_ADDRESS = 0x2540099e9ed04aF369d557a40da2D8f9c2ab928D;
const GAME_STAGE_LENGTH = 0;
const GAME_FEE_PERCENT = 5;

var gameAddresses = [];

module.exports = function(deployer) {
	deployer.deploy(GameFactory).then(async () => {
		const gameFactory = await GameFactory.at(GameFactory.address);

		const TwoThirds1 = await gameFactory.createGame(
			TWO_THIRDS_TYP,
    	0x2540099e9ed04aF369d557a40da2D8f9c2ab928D,
      GAME_FEE_PERCENT,
      BET_1,
      MAX_PLAYERS, 
      GAME_STAGE_LENGTH
    );
		gameAddresses.push(TwoThirds1.address);

		const TwoThirds01 = await gameFactory.createGame(
			TWO_THIRDS_TYP,
			0x2540099e9ed04aF369d557a40da2D8f9c2ab928D,
			GAME_FEE_PERCENT,
			BET_01,
			MAX_PLAYERS, 
			GAME_STAGE_LENGTH
    );
		gameAddresses.push(TwoThirds01.address);

		const TwoThirds001 = await gameFactory.createGame(
			TWO_THIRDS_TYP,
    	0x2540099e9ed04aF369d557a40da2D8f9c2ab928D,
      GAME_FEE_PERCENT,
      BET_001,
      MAX_PLAYERS, 
      GAME_STAGE_LENGTH
    );
    gameAddresses.push(TwoThirds001.address);

		const LUN1 = await gameFactory.createGame(
			LUN_TYP,
    	0x2540099e9ed04aF369d557a40da2D8f9c2ab928D,
      GAME_FEE_PERCENT,
      BET_1,
      MAX_PLAYERS, 
      GAME_STAGE_LENGTH
    );
		gameAddresses.push(LUN1.address);

		const LUN01 = await gameFactory.createGame(
			LUN_TYP,
    	0x2540099e9ed04aF369d557a40da2D8f9c2ab928D,
      GAME_FEE_PERCENT,
      BET_01,
      MAX_PLAYERS, 
      GAME_STAGE_LENGTH
    );
		gameAddresses.push(LUN01.address);

		const LUN001 = await gameFactory.createGame(
			LUN_TYP,
    	0x2540099e9ed04aF369d557a40da2D8f9c2ab928D,
      GAME_FEE_PERCENT,
      BET_001,
      MAX_PLAYERS, 
      GAME_STAGE_LENGTH
    );
    gameAddresses.push(LUN001.address);
	})
};
