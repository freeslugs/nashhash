pragma solidity ^0.4.23;

import "./Game.sol";

contract TwoThirdsAverage is Game {

    constructor(uint maxp) public Game(maxp) {}

    function findWinners() private {

        emit DebugWinner(1, player_addrs.length);
        
        // Calculate the 2/3 average
        uint guess_sum = 0;
        for(uint i = 0; i < player_addrs.length; i++){
            uint tmp = gameData[player_addrs[i]];
            guess_sum += tmp;
        }

        guess_sum = guess_sum * 10000;
        uint average = div(guess_sum, player_addrs.length);
        uint twothirdsavg = div(mul(average, 2), 3);
        twothirdsavg = twothirdsavg / 10000;

        //DEBUG
        average23 = twothirdsavg;

        address[] memory winners = new address[](config.MAX_PLAYERS);
        uint winIndex = 0;
        // We also flush the last list of winner
        delete info.lastWinners;


        // Find the guessers who are the closest to the 2/3 average
        uint min_diff = config.MAX_GUESS;
        uint cur_diff;
        for(i = 0; i < player_addrs.length; i++) {
            
            uint cur_guess = gameData[player_addrs[i]];

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
                winners[winIndex] = player_addrs[i];
                winIndex++;

                min_diff = cur_diff;
            // Else, if the difference are the same, we add the candidate to the 
            // list of winners
            } else if(cur_diff == min_diff){
                winners[winIndex] = player_addrs[i];
                winIndex++;
            }
        }

        emit DebugWinner(2, player_addrs.length);

        // winrIndex here has the number of winner in our array
        require(winIndex > 0);

        // Lets pay ourselves some money
        uint gamefee = (address(this).balance/100) * config.GAME_FEE_PERCENT;
        config.FEE_ADDRESS.transfer(gamefee);

        // Split the rest equally among winners
        uint prize = address(this).balance/winIndex;
        for(i = 0; i < winIndex; i++){
            winners[i].transfer(prize); 
            info.lastWinners.push(winners[i]);
            //emit DebugWinner(winners[i], winners.length);
        }
        info.lastPrize = prize;

        // DEBUG: Make sure no ether is lost due to rounding. 

        // RESET STATE
        toCommitState();
    }

}


