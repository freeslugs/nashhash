# NASHHASH

Game theory games on the blockchain

## Install

1. clone the app

1. `npm install -g truffle`

1. `npm install -g yarn`

1. `yarn install`

## Build



## Run locally 

1. `truffle develop`
  1. `compile; deploy`

1. `yarn start`  


## API i need to connect to 

P1: 
1. `commit(hashed guess` √ 
I added an event called SuccesfulCommit to allow you to update the UI.
I think you have to use watch() function.
That should be enough information
to dedeuce that the user was late with their bet.
2. `reveal(key, guess)` √
Same here, except SuccesfulReveal
Do not forget to check for errors.
3. `get_payout`=> how much did i win!
To get the payout there are few things you have to do:
    
      var number_of_winners = await game.num_last_winners();
      for (i = 0; i < number_of_winners; i++){
        var winner = await game.last_winners(i);
        if(winner == the_address_of_the_player){
          var prize = await game.last_prize();
        }
      }

P2: 
4. `get_status` => e.g. commit, reveal, payout
The easiest way to do this is by assigining a number to state.
0 --- COMMIT_STATE
1 --- REVEAL_STATE

Under current implementation, the contract has only two states, because the 
PAYOUT_STATE happens to be redundant i.e one function call will make transition REVEAL_STATE -> PAYOUT_STATE -> REVEAL_STATE, and since eth is fully sequential, we might as well skip the PAYOUT_STATE as the contract will never be found in it. Hope this makes sense.

So,

    var state = await game.game_state_debug();
    if (state.toNumber() == 0){
      // We are still accepting commits, the game is accepting bets
    }else if (state.toNumber() == 1){
      // We are waiting to get all the reveals
    }else{
      // We are so fucked
    }



5. `get_curr_users` => how many people have committed for htis game? 
Also a public variable, so you can get it by simply

P2.5:
1. bot to always complete game 

P3: 
1. `when will we move ot next stage in game?` => maybe save variable of when game start (block #)  and game end (blcok #)

