const MAX_GUESS = 100;

class Helper {
  constructor(web3, assert, game, api) {
    this.web3 = web3
    this.assert = assert
    this.game = game
    this.api = api
  }

  async runGame(bet, num_players, accounts, game, owner) {
      var guesses = this.createRandomGuesses(num_players, accounts);

      await this.api.setMaxPlayers(num_players, owner);

      for(var i = 0; i < num_players; i++){
          //const hash = Web3Utils.soliditySha3({type: 'string', value: guesses[1][i].toString()}, {type: 'string', value: "3"});
          await this.api.commitGuess( accounts[i], guesses[i].toString(), "3");
      }

      var state = await this.api.isInRevealState(game);
      assert(state == true, "Bad state transition, should be in REVEAL_STATE");
      for(i = 0; i < num_players; i++){
          await this.api.revealGuess( accounts[i], guesses[i].toString(), "3");
      }

      state = await this.api.isInPayoutState(game);
      assert(state == true, "Bad state transition, should be in PAYOUT_STATE");

      // Uncomment to check the balances
  /*    for(i = 0; i < num_players; i++){
          var balance = web3.fromWei(web3.eth.getBalance(accounts[i]),'ether').toString()
          console.log(balance);
      }*/

      await game.payout();

      //  >>> lowest game 
      var gamelowest = this.findLowestUniqueNum(guesses);

      var realLowest = await game.testLowest();

      var loc_winner = this.findWinner(accounts, guesses, realLowest);
      // Grab all the winners

      var winner = await this.api.getWinners(game);

      assert(winner.length == 1, "More than one winner");

      assert(winner[0] == loc_winner, "Winner is incorrectly chosen")

      state = await this.api.isInCommitState(game);
      assert(state == true, "Bad state transition, should be in COMMIT_STATE");
  }

  findWinner(player_addrs, guesses, lwst){
    var winner; 
    for(let i = 0; i < player_addrs.length; i++) {
        var cur_guess = guesses[i];

        //If current guess is lowest unique number, set winner to guesser address
        if(cur_guess == lwst){
         winner = player_addrs[i];
        }
    }
    return winner;
  }

  // todo: refactor using map
  createRandomGuesses(max_players, accounts) { 
    var guesses = new Array(max_players);
    for(var i = 0; i < max_players; i++){
      guesses[i] = Math.floor(Math.random() * 1000);
    }
    return guesses;
  }

  findLowestUniqueNum(guesses) {
    //Initialize empty sets
    var unique = new Set([]);
    var notUnique = new Set([]);


    for(var i = 0; i < guesses.length; i++){
      if(unique.has(guesses[i])){
        unique.delete(guesses[i]);
        notUnique.add(guesses[i]);
      } else if(!notUnique.has(guesses[i])){
        unique.add(guesses[i]);
      }
    }

    var lowest;
    for(let num of unique.values()){
      if(num < lowest){
        lowest = num;
      }
    }  

    return lowest;
  }
}

export default Helper;