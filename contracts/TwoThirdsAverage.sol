pragma solidity ^0.4.23;

import "./Game.sol";

contract TwoThirdsAverage is Game {
    using SafeMath for uint256;

    struct Rules {
        uint256 MIN_GUESS;
        uint256 MAX_GUESS;
    }

    Rules public rules;

    //DEBUG
    uint256 public average23 = 0;

    constructor(
        address _feeAddress,
        uint256 _gameFeePercent,
        uint256 _stakeSize,
        uint256 _maxp, 
        uint256 _gameStageLength,
        address _nptAddress) public Game(_feeAddress, _gameFeePercent, _stakeSize, _maxp, _gameStageLength, _nptAddress) {
        rules.MIN_GUESS = 0;
        rules.MAX_GUESS = 100;
    }

    function checkGuess(string guess) private {
        uint256 guess_num = stringToUint(guess);
        require(guess_num >= rules.MIN_GUESS && guess_num <= rules.MAX_GUESS);
    }

    function findWinners() private {

        //emit DebugWinner(1, gameDataKeys.length);
        
        // Calculate the 2/3 average
        uint256 guess_sum = 0;
        for(uint256 i = 0; i < state.currNumberReveals; i++){
            uint256 tmp = stringToUint(gameData[gameDataKeys[i]]);
            guess_sum = guess_sum.add(tmp);
        }

        guess_sum = guess_sum.mul(10000);
        uint256 average = guess_sum.div(state.currNumberReveals);
        uint256 twothirdsavg = average.mul(2).div(3);
        twothirdsavg = twothirdsavg.div(10000);

        //DEBUG
        average23 = twothirdsavg;

        address[] memory winners = new address[](config.MAX_PLAYERS);
        uint256 winIndex = 0;

        // Find the guessers who are the closest to the 2/3 average
        uint256 min_diff = ~uint256(0);
        uint256 cur_diff;
        for(i = 0; i < state.currNumberReveals; i++) {
            
            uint256 cur_guess = stringToUint(gameData[gameDataKeys[i]]);

            // Find the difference between the guess and the average
            if(twothirdsavg > cur_guess){
                cur_diff = twothirdsavg.sub(cur_guess);
            }
            else{
                cur_diff = cur_guess.sub(twothirdsavg);
            }
            
            // If the difference is less than the smallest difference,
            // we delete all the winners and add the new candidate
            if(cur_diff < min_diff) {
                
                // NOT A BUG!
                //delete winners; // WTF you might ask? There is no necessity to delete elements.
                winIndex = 0;
                winners[winIndex] = gameDataKeys[i];
                winIndex = winIndex.add(1);

                min_diff = cur_diff;
            // Else, if the difference are the same, we add the candidate to the 
            // list of winners
            } else if(cur_diff == min_diff){
                winners[winIndex] = gameDataKeys[i];
                winIndex = winIndex.add(1);
            }
        }

        //emit DebugWinner(2, gameDataKeys.length);

        // winrIndex here has the number of winner in our array
        require(winIndex > 0);

        // Lets pay ourselves some money
        uint256 gamefee = (address(this).balance.mul(config.GAME_FEE_PERCENT)).div(100);
        config.FEE_ADDRESS.transfer(gamefee);

        // Split the rest equally among winners
        uint256 prize = address(this).balance.div(winIndex);
        performPayout(winners, winIndex, prize);
        
    }

}


