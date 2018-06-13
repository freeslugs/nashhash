var Game = artifacts.require("TwoThirdsAverage");
var NPT = artifacts.require("NPT");

var Web3Utils = require('web3-utils');


const BET = web3.toWei(0.1, 'ether');
const MAX_PLAYERS = 5;
const HASHNASH_ADDRESS = 0x2540099e9ed04aF369d557a40da2D8f9c2ab928D;
const GAME_STAGE_LENGTH = 0;
const GAME_FEE_PERCENT = 5;

module.exports = function(deployer) {
    deployer.deploy(Game,
    HASHNASH_ADDRESS,
    GAME_FEE_PERCENT,
    BET,
    MAX_PLAYERS,
    GAME_STAGE_LENGTH,
    HASHNASH_ADDRESS
    ).then(async () => {
        const game = await Game.at(Game.address);
        const npt = await NPT.deployed();
        const NPT_ADDRESS = npt.address;

        await game.setNPTAddress(NPT_ADDRESS);
    })
};

