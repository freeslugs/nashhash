var Web3Utils = require('web3-utils');

class API {
  constructor(web3, assert, game) {
    this.web3 = web3
    this.assert = assert
    this.game = game
  }




  // Get the state of the game
  async isInCommitState() {
    const state = await this.getGameState();
    if (state == 0) {
      return true;
    } else {
      return false;
    }
  }

  async isInRevealState() {
    const state = await this.getGameState();
    if (state == 1) {
      return true;
    } else {
      return false;
    }
  }

  async isInPayoutState() {
    const state = await this.getGameState();
    if (state == 2) {
      return true;
    } else {
      return false;
    }
  }

  async getGameState() {
    const state = await this.game.getGameState();
    return state.toNumber();
  }
  ////


  // Get number of commits and number of reveales in the current round
  async getCurrentCommits() {
    const currNumberCommits = await this.game.getCurrentCommits();
    return currNumberCommits.toNumber();
  }

  async getCurrentReveals() {
    const curNumberReveals = await this.game.getCurrentReveals();
    return curNumberReveals.toNumber();
  }


  // Reset the game to the initial state
  async resetGame() {
    await this.game.resetGame();
  }


  // Commit and reveal functions
  async commitGuess(usr_addr, guess, random) {
    const bet = await this.getStakeSize();
    const hash = this.hashGuess(guess, random);
    await this.game.commit(hash, { value: this.web3.toWei(bet,'ether'), from: usr_addr });
  }

  async revealGuess(usr_addr, guess, random) {
    await this.game.reveal(guess, random, { from: usr_addr });
  }

  hashGuess(guess, random) {
    const hash = Web3Utils.soliditySha3({ type: 'string', value: guess }, { type: 'string', value: random });
    return hash;
  }


  // Get the size of the stake for the game
  async getStakeSize() {
    const bet = await this.game.getStakeSize();
    return this.web3.fromWei(bet.toNumber(), 'ether');
  }


  // Functiuons help getting information about the prize
  async getWinners() {
    let winners = [];

    const nw = await this.getNumberOfWinners();

    for (let i = 0; i < nw; i++) {
      const winner = await this.game.getLastWinners(i);
      winners.push(winner);
    }
    return winners;
  }

  async getNumberOfWinners() {
    const nw = await this.game.getNumberOfWinners();
    return nw.toNumber();
  }

  async getPayout(usr_addr) {
    const winners = await this.getWinners();
    const prize = await this.game.getLastPrize();
    
    for (let i = 0; i < winners.length; i++) {
      if (winners[i] == usr_addr) {
        return this.web3.fromWei(prize.toNumber(), 'ether');
      }
    }
    return 0;
  }

  async getPrizeAmount() {
    const prize = await this.game.getLastPrize();
    return this.web3.fromWei(prize.toNumber(), 'ether');
  }

  async getGameFeeAmount() {
    //console.log("10100101010101010010101");
    const fee = await this.game.getGameFee();
    return fee.toNumber();
  }


  // Control functions that can be used by the owner
  async pauseGame() {
    await this.game.pause();
  }

  async unpauseGame() {
    await this.game.unpause();
  }

  async setMaxPlayers(num) {
    await this.game.setMaxPlayers(num);
  }

  async getMaxPlayers() {
    const num = await this.game.getMaxPlayers();
    return num.toNumber();
  }


  // These functions are used by the Game Master to control the flow of the game
  async payout(){
    await this.game.payout();
  }

  async forceToRevealState(){
    await this.game.forceToRevealState();
  }

  async forceToPayoutState(){
    await this.game.forceToPayoutState();
  }

  async getCommitStageStartBlock(){
    const csb = await this.game.getCommitStageStartBlock()
    return csb
  }

  async getRevealStageStartBlock(){
    const rsb = await this.game.getRevealStageStartBlock()
    return rsb
  }


  /* Cool . 
    - ev is the event to watch for from the contract. EX. game.CommitsSubmitted
    - handler is the  that is to be called when that event is emited by the contract
    - handler_args_list is a list of argumetns to the handler
  */
  // todo: 
  /*watchEvent(e) {    
    return new Promise((resolve, reject) => {
      const event = e({}, { fromBlock: 0, toBlock: 'latest' });
      event.watch((error, result) => {
        if (!error) {
          // handler.apply(this, handler_args_list);
          resolve()
        } else {
          console.log(error);
          // assert(true == false, "event handler failed to be installed");
          reject("event handler failed to be installed")
        }
      });
    })
  }*/

  watchEvent(ev, handler, handler_args_list) {    
    const event = ev({}, { fromBlock: 0, toBlock: 'latest' });
    event.watch((error, result) => {
      if (!error) {
        console.log('success! ')
        handler.apply(this, handler_args_list);
      } else {
        console.log(error);
        this.assert(true == false, "event handler failed to be installed");
      }
    });
  }

  /*watchEvent(ev, handler) {
    var event = ev({}, { fromBlock: 0, toBlock: 'latest' });
    event.watch((error, result) => {
      if (!error) {
        handler();
      } else {
        console.log(error);
        assert(true == false, "event handler failed to be installed");
      }
    });
  }*/
}

export default API;