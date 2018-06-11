const MAX_GUESS = 100;

class Helper {
  constructor(web3, assert, game, api) {
    this.web3 = web3
    this.assert = assert
    this.game = game
    this.api = api
  }

  async runGame(bet, num_players, accounts, owner) {
    var guesses = this.createRandomGuesses(num_players, accounts);
    await this.api.setMaxPlayers(num_players, owner);

    for(var i = 0; i < num_players; i++){
        await this.api.commitGuess(accounts[i], guesses[i].toString(), "3");
    }

    var state = await this.api.isInRevealState();
    
    assert(state == true, "Bad state transition, should be in REVEAL_STATE");
    for(i = 0; i < num_players; i++){
        await this.api.revealGuess(accounts[i], guesses[i].toString(), "3");
    }

    state = await this.api.isInPayoutState();
    assert(state == true, "Bad state transition, should be in PAYOUT_STATE");

    // Uncomment to check the balances
    for(i = 0; i < num_players; i++){
        var balance = web3.fromWei(web3.eth.getBalance(accounts[i]),'ether').toString()
    }

    await this.api.payout();

    var average = this.computeTwoThirdsAverage(guesses);
    var average23 = await this.api.game.average23();
    
    // DEBUG: Put this one back in. There is a known problem here.
    assert(Math.floor(average) == average23.toNumber(), "Average23 miscalculated...");

    var loc_winners = this.findWinner(accounts, guesses, average23);

    // Grab all the winners
    var winners = await this.api.getWinners();
    var number_of_winners = winners.length;

    assert(loc_winners.length == number_of_winners, "Number of winners varies");

    for (i = 0; i < number_of_winners; i++){
        var winner = winners[i];
        assert(winner == loc_winners[i], "Wrong winner");
    }

    state = await this.api.isInCommitState();
    assert(state == true, "Bad state transition, should be in COMMIT_STATE");
  }

  sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }

  findWinner(player_addrs, guesses, avrg) { 
    var min_diff = MAX_GUESS;
    var cur_diff;

    let winners = [];
    for(var i = 0; i < player_addrs.length; i++) {
      var cur_guess = guesses[i];
      // Find the difference between the guess and the average
      if(avrg > cur_guess) {
        cur_diff = avrg - cur_guess;
      } else {
        cur_diff = cur_guess - avrg;
      }
      
      // If the difference is less than the smallest difference,
      // we delete all the winners and add the new candidate
      if(cur_diff < min_diff) {
        winners = [];
        winners.push(player_addrs[i]);
        min_diff = cur_diff;
      // Else, if the difference are the same, we add the candidate to the 
      // list of winners
      } else if(cur_diff == min_diff) {
        winners.push(player_addrs[i]);
      }
    }
    return winners;
  }

  createRandomGuesses(max_players, accounts) { 
    var guesses = new Array(max_players);
    for(var i = 0; i < max_players; i++){
      guesses[i] = Math.floor(Math.random() * 101);
    }
    return guesses;
  }

  computeTwoThirdsAverage(guesses) {
    var sum = guesses.reduce(function(acc, val) { return acc + val; });
    sum = sum * 10000;
    var average = Math.floor(sum / guesses.length);
    var average23 = Math.floor((average * 2) / 3);
    average23 = Math.floor(average23/10000)

    return average23;    
  }

  // todo: this is not being used in this file
  spin(i) {
      var j;
      var sum;
      for(j = 0; j < i; j++ ){
          sum = j + i;
      }
      return sum;
  }

  async expectedThrow(promise, msg) {
    try {
        await promise;
    } catch (err) {
        return;
    }
    assert(false, msg);
  } 
}

export default Helper;