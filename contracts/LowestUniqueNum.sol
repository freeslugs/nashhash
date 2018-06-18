pragma solidity ^0.4.23;

import "./Game.sol";


//NEEDED: Account for scenario where no one wins (no unique number) !!!
contract LowestUniqueNum is Game {
    using SafeMath for uint256;

    address NULL_ADDR = config.FEE_ADDRESS;

    // Guesses and guesser addresses
    mapping (uint256 => address) private guessAddrs;
    uint256[] private guesses;

    struct Rules {
        uint256 MIN_GUESS;
        uint256 MAX_GUESS;
    }

    Rules public rules;

    // Debug
    uint256 public testLowest = 0;

    constructor(
        address _feeAddress,
        uint256 _gameFeePercent,
        uint256 _stakeSize,
        uint256 _maxp, 
        uint256 _gameStageLength,
        address _nptAddress) public Game(_feeAddress, _gameFeePercent, _stakeSize, _maxp, _gameStageLength, _nptAddress) {
        rules.MIN_GUESS = 0;
        rules.MAX_GUESS = (~uint256(0)).sub(1);
    }

    function checkGuess(string guess) private {
        uint256 guess_num = stringToUint(guess);
        require(guess_num >= rules.MIN_GUESS && guess_num <= rules.MAX_GUESS);
    }

    function findWinners() private {

        for(uint256 i = 0; i < state.currNumberReveals; i++){
            uint256 tmp = stringToUint(gameData[gameDataKeys[i]]);

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

        uint256 lowest_unique_guess = rules.MAX_GUESS.add(1);
        //Declared as array for interface purposes
        address[] memory winner = new address[](1);

        for(uint256 j = 0; j < guesses.length; j++){
            uint256 cur_guess = guesses[j];

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
        uint256 gamefee = (address(this).balance.mul(config.GAME_FEE_PERCENT)).div(100);
        config.FEE_ADDRESS.transfer(gamefee);

        // Give the rest to winner
        uint256 prize = address(this).balance;
        performPayout(winner, 1, prize);


    } 
}