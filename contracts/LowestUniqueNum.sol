pragma solidity ^0.4.23;

import "./Game.sol";


//NEEDED: Account for scenario where no one wins (no unique number) !!!
contract LowestUniqueNum is Game {

    address NULL_ADDR = config.FEE_ADDRESS;

    // Guesses and guesser addresses
    mapping (uint => address) private guessAddrs;
    uint[] private guesses;

    struct Rules {
        uint MIN_GUESS;
        uint MAX_GUESS;
    }

    Rules public rules;

    // Debug
    uint public testLowest = 0;

    constructor(uint maxp) public Game(maxp) {
        rules.MIN_GUESS = 0;
        rules.MAX_GUESS = ~uint256(0) - 1;
    }

    function checkGuess(string guess) private {
        uint guess_num = stringToUint(guess);
        require(guess_num >= rules.MIN_GUESS && guess_num <= rules.MAX_GUESS);
    }

    function findWinners() private {

       for(uint i = 0; i < state.currNumberReveals; i++){
            uint tmp = stringToUint(gameData[gameDataKeys[i]]);

            //If guess is unique input guesser address to mapping
            if(guessAddrs[tmp] == 0){
                guessAddrs[tmp] = gameDataKeys[i];
                guesses.push(tmp);
            }

            //If guess is no longer unique input NULL_ADDR to mapping
            else {
                guessAddrs[tmp] = NULL_ADDR;
            }
        }

        uint lowest_unique_guess = rules.MAX_GUESS + 1;
        //Declared as array for interface purposes
        address[] memory winner = new address[](1);
        // We also flush the last list of winners
        delete info.lastWinners;

        for(uint j = 0; j < guesses.length; j++){
            uint cur_guess = guesses[j];

            //If guess is unique
            if(guessAddrs[cur_guess] != NULL_ADDR){
                //If guess is lower than current min
                if(guesses[j] < lowest_unique_guess){
                    lowest_unique_guess = guesses[j];
                    testLowest = guesses[j];
                    winner[0] = guessAddrs[cur_guess];
                }
            }
            //Clear mapping for next round
            delete guessAddrs[cur_guess];
        }

        //Clear array for next round
        delete guesses;

        // Lets pay ourselves some money
        uint gamefee = (address(this).balance/100) * config.GAME_FEE_PERCENT;
        config.FEE_ADDRESS.transfer(gamefee);

        // Give the rest to winner
        uint prize = address(this).balance;
        performPayout(winner, 1, prize);


    } 
}






