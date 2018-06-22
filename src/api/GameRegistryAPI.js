import GameABI from '../contracts/Game.json'

const contract = require('truffle-contract');
var Web3Utils = require('web3-utils');

class GameRegistryAPI {
	constructor(web3, game_addrs) {
		this.web3 = web3
	    this.game_addrs = game_addrs
	}

	async configureGame(game_typ, stake) {
		let game

		console.log(this.web3);

		const GameContract = contract(GameABI);
	    GameContract.setProvider(this.web3.currentProvider);

	    if(game_typ=="TwoThirds"){
	      if(stake==0.01){
	        game = await GameContract.at(this.game_addrs[2]);
	      }
	      else if(stake==0.1){
	        game = await GameContract.at(this.game_addrs[1]);
	      }
	      else if(stake==1){
	        game = await GameContract.at(this.game_addrs[0]);
	      }
	    } else if(game_typ=="LowestUnique")
	    {
	      if(stake==0.01){
	        game = await GameContract.at(this.game_addrs[5]);
	      }
	      else if(stake==0.1){
	        game = await GameContract.at(this.game_addrs[4]);
	      }
	      else if(stake==1){
	        game = await GameContract.at(this.game_addrs[3]);
	      }      
	    }

	    return game
	}
}

export default GameRegistryAPI;
