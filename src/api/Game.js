var Web3Utils = require('web3-utils');

class API {
  static async isInCommitState(game) {
    const state = await game.getGameState();
    if (state.toNumber() == 0) {
      return true;
    } else {
      return false;
    }
  }

  static async isInRevealState(game) {
    const state = await game.getGameState();
    if (state.toNumber() == 1) {
      return true;
    } else {
      return false;
    }
  }

  static async isInPayoutState(game) {
    const state = await game.getGameState();
    if (state.toNumber() == 2) {
      return true;
    } else {
      return false;
    }
  }

  static async getGameState() {
    const state = await game.getGameState();
    return state.toNumber();
  }

  static async getCurrentCommits(game) {
    const currNumberCommits = await game.getCurrentCommits();
    return currNumberCommits.toNumber();
  }

  static async getCurrentReveals(game) {
    const curNumberReveals = await game.getCurrentReveals();
    return curNumberReveals.toNumber();
  }

  static async resetGame(game) {
    await game.resetGame();
  }

  static async commitGuess(game, usr_addr, guess, random) {
    const bet = await this.getStakeSize(game);
    const hash = this.hashGuess(guess, random);
    await game.commit(hash, { value: web3.toWei(bet,'ether'), from: usr_addr });
  }

  static async revealGuess(game, usr_addr, guess, random) {
    await game.reveal(guess, random, { from: usr_addr });
  }

  static hashGuess(guess, random) {
    const hash = Web3Utils.soliditySha3({ type: 'string', value: guess }, { type: 'string', value: random });
    return hash;
  }

  static async getStakeSize(game) {
    const bet = await game.getStakeSize();
    return web3.fromWei(bet.toNumber(), 'ether');
  }

  static async getWinners(game) {
    let winners = [];

    const nw = await game.getNumberOfWinners();
    const number_of_winners = nw.toNumber();

    for (let i = 0; i < number_of_winners; i++) {
      const winner = await game.getLastWinners(i);
      winners.push(winner);
    }

    return winners;
  }

  static async getPayout(game, usr_addr) {
    const winners = await this.getWinners(game);
    const prize = await game.getLastPrize();
    
    for (let i = 0; i < winners.length; i++) {
      if (winners[i] == usr_addr) {
        return web3.fromWei(prize.toNumber(), 'ether');
      }
    }
    return 0;
  }

  static async getPrizeAmount(game) {
    const prize = await game.getLastPrize();
    return web3.fromWei(prize.toNumber(), 'ether');
  }

  static async getGameFeeAmount(game) {
    const fee = await game.getGameFee();
    return fee.toNumber();
  }

  static async pauseGame(game) {
    await game.pause();
  }

  static async unpauseGame(game) {
    await game.unpause();
  }

  static async setMaxPlayers(game, num) {
    await game.setMaxPlayers(num);
  }

  static async getMaxPlayers(game) {
    const num = await game.getMaxPlayers();
    return num.toNumber();
  }

  static async payout(game){
    await game.payout();
  }

  static async forceToRevealState(game){
    await game.forceToRevealState();
  }

  static async forceToPayoutState(game){
    await game.forceToPayoutState();
  }


  /* Cool . 
    - ev is the event to watch for from the contract. EX. game.CommitsSubmitted
    - handler is the  that is to be called when that event is emited by the contract
    - handler_args_list is a list of argumetns to the handler
  */
  // todo: 
  /*static watchEvent(e) {    
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

  static watchEvent(ev, handler, handler_args_list) {    
    const event = ev({}, { fromBlock: 0, toBlock: 'latest' });
    event.watch((error, result) => {
      if (!error) {
        console.log('success! ')
        handler.apply(this, handler_args_list);
      } else {
        console.log(error);
        assert(true == false, "event handler failed to be installed");
      }
    });
  }

  /*static watchEvent(ev, handler) {
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