pragma solidity ^0.4.23;

import "./Game.sol";

contract TwoThirdsAverage is Game {


    struct Rules {
        uint MIN_GUESS;
        uint MAX_GUESS;
    }

    Rules public rules;

    //DEBUG
    uint public average23 = 0;

    constructor(uint maxp) public Game(maxp) {
        rules.MIN_GUESS = 0;
        rules.MAX_GUESS = 100;
    }

    function checkGuess(string guess) private {
        uint guess_num = stringToUint(guess);
        require(guess_num >= rules.MIN_GUESS && guess_num <= rules.MAX_GUESS);
    }

    function findWinners() private {

        //emit DebugWinner(1, gameDataKeys.length);
        
        // Calculate the 2/3 average
        uint guess_sum = 0;
        for(uint i = 0; i < state.currNumberReveals; i++){
            uint tmp = stringToUint(gameData[gameDataKeys[i]]);
            guess_sum += tmp;
        }

        guess_sum = guess_sum * 10000;
        uint average = div(guess_sum, state.currNumberReveals);
        uint twothirdsavg = div(mul(average, 2), 3);
        twothirdsavg = twothirdsavg / 10000;

        //DEBUG
        average23 = twothirdsavg;

        address[] memory winners = new address[](config.MAX_PLAYERS);
        uint winIndex = 0;

        // Find the guessers who are the closest to the 2/3 average
        uint min_diff = ~uint256(0);
        uint cur_diff;
        for(i = 0; i < state.currNumberReveals; i++) {
            
            uint cur_guess = stringToUint(gameData[gameDataKeys[i]]);

            // Find the difference between the guess and the average
            if(twothirdsavg > cur_guess){
                cur_diff = twothirdsavg - cur_guess;
            }
            else{
                cur_diff = cur_guess - twothirdsavg;
            }
            
            // If the difference is less than the smallest difference,
            // we delete all the winners and add the new candidate
            if(cur_diff < min_diff) {
                
                // NOT A BUG!
                //delete winners; // WTF you might ask? There is no necessity to delete elements.
                winIndex = 0;
                winners[winIndex] = gameDataKeys[i];
                winIndex++;

                min_diff = cur_diff;
            // Else, if the difference are the same, we add the candidate to the 
            // list of winners
            } else if(cur_diff == min_diff){
                winners[winIndex] = gameDataKeys[i];
                winIndex++;
            }
        }

        //emit DebugWinner(2, gameDataKeys.length);

        // winrIndex here has the number of winner in our array
        require(winIndex > 0);

        // Lets pay ourselves some money
        uint gamefee = (address(this).balance/100) * config.GAME_FEE_PERCENT;
        config.FEE_ADDRESS.transfer(gamefee);

        // Split the rest equally among winners
        uint prize = address(this).balance/winIndex;
        performPayout(winners, winIndex, prize);
        
    }

}


