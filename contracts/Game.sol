pragma solidity ^0.4.23;



/*
!!!!!!!!!!!KNOWN BUGS!!!!!!!!!!!!!!!
1) 2/3 average not consistent with the js results. NEEDS FIXING.

!!!!!!!!!!!! POTENTIAL BUGS !!!!!!!!!!!!!!!!
1) Ether being brought over to the next round due to rounding issues. 
Check that this is not the case. 

!!!!!!!!!! LOGICAL IMPROVEMENTS !!!!!!!!!!!!
1) Send money to HOME address only once the fees reach a certain amount.

*/

import "./Ownable.sol";
import "./GameHelper.sol";

contract Game is Ownable, GameHelper {

    enum GameState {COMMIT_STATE, REVEAL_STATE}
    GameState public game_state = GameState.COMMIT_STATE;
    //DEBUG: Enums cannot be tested. Value mirrors the enum
    uint public game_state_debug = 0;

    mapping (address => bytes32) public commits;
    mapping (address => uint) public game_data;
    address[] public player_addrs;
    address[] public winners;
    
    address OUR_ADDRESS = 0x2540099e9ed04aF369d557a40da2D8f9c2ab928D; //Address 
    uint public constant GAME_FEE_PERCENT = 5;
    uint public MAX_PLAYERS = 1;
    uint public constant MIN_GUESS = 0;
    uint public constant MAX_GUESS = 100; 

    uint public BET_SIZE = 1 ether; //0.01 ether;
    uint public curr_number_bets = 0;
    uint public curr_number_reveals = 0;

    // DEBUG vars
    uint public average23 = 0;
    address[] public last_winners;
    uint public num_last_winners;

    function set_MAX_PLAYERS(uint new_val) public onlyOwner {
        MAX_PLAYERS = new_val;
    }

    // Commit your guess. 
    function commit(bytes32 hashed_com) public payable{
        require(game_state == GameState.COMMIT_STATE);
        require(msg.value == BET_SIZE);

        commits[msg.sender] = hashed_com;
        curr_number_bets++;

        // If we received the MAX_PLAYER number of commits, it is time for
        // us to change state.
        if (curr_number_bets == MAX_PLAYERS) {
            toRevealState();
        }
    }

    function reveal(string guess, string random) public  {
        
        require(game_state == GameState.REVEAL_STATE);
        
        // DEBUG: Need to make sure it throws if the guess is not integer
        uint guess_num = stringToUint(guess);
        
        require(guess_num >= MIN_GUESS && guess_num <= MAX_GUESS);

        // Check that the hashes match
        require(commits[msg.sender] == keccak256(guess, random));

        // When they do, we add the revealed guess to game data
        game_data[msg.sender] = guess_num;
        player_addrs.push(msg.sender);
        curr_number_reveals++;

        if(curr_number_reveals == MAX_PLAYERS){
            find_winner();
        }
    }


    function find_winner() public {
        
        // Calculate the 2/3 average
        uint guess_sum = 0;
        for(uint i = 0; i < player_addrs.length; i++){
            uint tmp = game_data[player_addrs[i]];
            guess_sum += tmp;
        }
        uint average = div(guess_sum, player_addrs.length);
        uint twothirdsavg = div(mul(average, 2), 3);

        //DEBUG
        average23 = twothirdsavg;


        // Find the guessers who are the closest to the 2/3 average
        uint min_diff = MAX_GUESS;
        uint cur_diff;
        for(i = 0; i < player_addrs.length; i++) {
            
            uint cur_guess = game_data[player_addrs[i]];

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
                delete winners;
                winners.push(player_addrs[i]);
                min_diff = cur_diff;
            // Else, if the difference are the same, we add the candidate to the 
            // list of winners
            } else if(cur_diff == min_diff){
                winners.push(player_addrs[i]);
            }
        }

        // Lets pay ourselves some money
        uint gamefee = (address(this).balance/100) * GAME_FEE_PERCENT;
        OUR_ADDRESS.transfer(gamefee);

        // Split the rest equally among winners
        uint prize = address(this).balance/winners.length;
        for(i = 0; i < winners.length; i++){
            winners[i].transfer(prize); 
        }

        // DEBUG: Make sure no ether is lost due to rounding. 



        // RESET STATE
        toCommitState();
    }

    // Call this funtion to get to COMMIT_STATE
    function toCommitState() internal {
        game_state = GameState.COMMIT_STATE;
        game_state_debug = 0;

        last_winners = winners;
        num_last_winners = winners.length;
        delete winners;
        curr_number_bets = 0;
        curr_number_reveals = 0;
    }

    // Call this function to get to REVEAL_STATE
    function toRevealState() internal {
        game_state = GameState.REVEAL_STATE;
        game_state_debug = 1;
    }

}