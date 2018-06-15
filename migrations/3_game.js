module.exports = (deployer, network, accounts) => {
	var GameRegistry = artifacts.require("GameRegistry");
	var TwoThirdsAverage = artifacts.require("TwoThirdsAverage");
	var LowestUniqueNum = artifacts.require("LowestUniqueNum");
	var NPT = artifacts.require("NPT");
	var Web3Utils = require('web3-utils');

	const BET_1 = web3.toWei(1,'ether');
	const BET_01 = web3.toWei(0.1, 'ether');
	const BET_001 = web3.toWei(0.01, 'ether');
	const MAX_PLAYERS = 10;

	const HASHNASH_ADDRESS = 0x2540099e9ed04aF369d557a40da2D8f9c2ab928D;
	const GAME_STAGE_LENGTH = 0;
	const GAME_FEE_PERCENT = 5;

	var gameAddresses = [];

	deployer.deploy(GameRegistry).then(async () => {
		const npt = await NPT.deployed(web3);
		const NPT_ADDRESS = npt.address;

		const gameRegistry = GameRegistry.at(GameRegistry.address);

		const TwoThirds1 = await TwoThirdsAverage.new(
	    	HASHNASH_ADDRESS,
	        GAME_FEE_PERCENT,
	        BET_1,
	        MAX_PLAYERS, 
	        GAME_STAGE_LENGTH,
	        NPT_ADDRESS
	    );
		gameRegistry.insert(TwoThirds1.address);
		gameAddresses.push(TwoThirds1.address);
		await npt.addMinter(TwoThirds1.address);

		const TwoThirds01 = await TwoThirdsAverage.new(
	    	HASHNASH_ADDRESS,
	        GAME_FEE_PERCENT,
	        BET_01,
	        MAX_PLAYERS, 
	        GAME_STAGE_LENGTH,
	        NPT_ADDRESS
	    );
	    gameRegistry.insert(TwoThirds01.address);
		gameAddresses.push(TwoThirds01.address);
		npt.addMinter(TwoThirds01.address);

		const TwoThirds001 = await TwoThirdsAverage.new(
	    	HASHNASH_ADDRESS,
	        GAME_FEE_PERCENT,
	        BET_001,
	        MAX_PLAYERS, 
	        GAME_STAGE_LENGTH,
	        NPT_ADDRESS
	    );
	    gameRegistry.insert(TwoThirds001.address);
	    gameAddresses.push(TwoThirds001.address);
	    npt.addMinter(TwoThirds001.address);

		const LUN1 = await LowestUniqueNum.new(
	    	HASHNASH_ADDRESS,
	        GAME_FEE_PERCENT,
	        BET_1,
	        MAX_PLAYERS, 
	        GAME_STAGE_LENGTH,
	        NPT_ADDRESS
	    );
	   	gameRegistry.insert(LUN1.address);
		gameAddresses.push(LUN1.address);
		npt.addMinter(LUN1.address);

		const LUN01 = await LowestUniqueNum.new(
	    	HASHNASH_ADDRESS,
	        GAME_FEE_PERCENT,
	        BET_01,
	        MAX_PLAYERS, 
	        GAME_STAGE_LENGTH,
	        NPT_ADDRESS
	    );
	    gameRegistry.insert(LUN01.address);
	 	gameAddresses.push(LUN01.address);
		npt.addMinter(LUN01.address);

		const LUN001 = await LowestUniqueNum.new(
	    	HASHNASH_ADDRESS,
	        GAME_FEE_PERCENT,
	        BET_001,
	        MAX_PLAYERS, 
	        GAME_STAGE_LENGTH,
	        NPT_ADDRESS
	    );
	    gameRegistry.insert(LUN001.address);
	    gameAddresses.push(LUN001.address);
	    npt.addMinter(LUN001.address);
	})
};
