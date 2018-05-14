pragma solidity ^0.4.23;



/*
!!!!!!!!!!! KNOWN BUGS !!!!!!!!!!!!!!!
1) 2/3 average consistent with the js results. BUT. The way we calculate it
in solidity seems to be rounded down at the taking of the average step, and then
again rounded down after division by 3. It is better to fix this because this skews
the average ever so slightly south. 
2) Protect against repeated reveal call

!!!!!!!!!!!! POTENTIAL BUGS !!!!!!!!!!!!!!!!
1) Ether being brought over to the next round due to rounding issues. 
Check that this is not the case. 

!!!!!!!!!! LOGICAL IMPROVEMENTS !!!!!!!!!!!!
1) Send money to HOME address only once the fees reach a certain amount.

*/

import "./Ownable.sol";
import "./GameHelper.sol";

contract Game is Ownable, GameHelper {

    // Tracks the state of the game. 
    enum GameState {COMMIT_STATE, REVEAL_STATE}
    GameState public game_state = GameState.COMMIT_STATE;
    //DEBUG: Enums cannot be tested. Value mirrors the enum
    uint public game_state_debug = 0;

    // Commit/Reveal Protocol vars
    mapping (address => bytes32) public commits;
    mapping (address => uint) public game_data;
    address[] internal player_addrs;
    address[] internal winners;
    uint public curr_number_bets = 0;
    uint public curr_number_reveals = 0;
    uint public final_commit_block = 0;
    uint constant REVEAL_PERIOD = 5;
    
    // Money Specifics
    address OUR_ADDRESS = 0x2540099e9ed04aF369d557a40da2D8f9c2ab928D; //Address 
    uint public constant GAME_FEE_PERCENT = 5;
    uint public BET_SIZE = 1 ether; 
    
    // Rules
    uint public MAX_PLAYERS = 10;
    uint public constant MIN_GUESS = 0;
    uint public constant MAX_GUESS = 100;

    // UI vars
    address[] public last_winners;
    uint public num_last_winners = 0;
    uint public last_prize = 0;


    ////// DEBUG vars and debug functions
    uint public average23 = 0;

    function set_MAX_PLAYERS(uint new_val) public onlyOwner {
        MAX_PLAYERS = new_val;
    }

    // function is used to trigger a payout in a situation where somone
    // forgets to send the reveal.
    function trigger_payout() public onlyOwner {
        require(game_state == GameState.REVEAL_STATE);

        // If the REVEAL_PERIOD blocks has gone by, while unfair, 
        // keep the money of nonrevealers, play the game with the
        // rest of the players.
        if(block.number > final_commit_block + REVEAL_PERIOD){
            find_winner();
        }
    }

    // Reset the contract to the initial state
    function reset() public onlyOwner {  
        toCommitState();
        delete last_winners;
        num_last_winners = 0;
        last_prize = 0;
    }

    ////

    // constructor(uint num_players, uint bet_size) public {
    //     MAX_PLAYERS = num_players;
    //     BET_SIZE = bet_size;
    //    // OUR_ADDRESS = fee_addr;
    //     owner = msg.sender;
    // }

    // Commit your guess. 
    event SuccesfulCommit(
        bytes32 hashed_commit
    );

    function commit(bytes32 hashed_com) public payable{
        require(game_state == GameState.COMMIT_STATE);
        require(msg.value == BET_SIZE);

        commits[msg.sender] = hashed_com;
        curr_number_bets++;

        // Notify the user that their bet reached us
        emit SuccesfulCommit(hashed_com);

        // If we received the MAX_PLAYER number of commits, it is time for
        // us to change state.
        if (curr_number_bets == MAX_PLAYERS) {
            toRevealState();
        }
    }

    event SuccesfulReveal(
        string guess,
        string random
    );
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

        emit SuccesfulReveal(guess, random);

        if(curr_number_reveals == MAX_PLAYERS){
            find_winner();
        }
    }

    event DebugWinner(address addr, uint n);


    function find_winner() private {

        emit DebugWinner(5, player_addrs.length);
        
        // Calculate the 2/3 average
        uint guess_sum = 0;
        for(uint i = 0; i < player_addrs.length; i++){
            uint tmp = game_data[player_addrs[i]];
            guess_sum += tmp;
        }

        guess_sum = guess_sum * 10000;
        uint average = div(guess_sum, player_addrs.length);
        uint twothirdsavg = div(mul(average, 2), 3);
        twothirdsavg = twothirdsavg / 10000;

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
            //emit DebugWinner(winners[i], winners.length);
        }
        last_prize = prize;

        // DEBUG: Make sure no ether is lost due to rounding. 

        // RESET STATE
        toCommitState();
    }

    // Call this funtion to get to COMMIT_STATE
    event DebugCommitState(uint last_win_l, uint win_l);

    function toCommitState() internal {
        game_state_debug = 0;

        delete player_addrs;
        delete last_winners;
        last_winners = winners;
        num_last_winners = winners.length;
        delete winners;
        curr_number_bets = 0;
        curr_number_reveals = 0;

        game_state = GameState.COMMIT_STATE;
        //DebugCommitState(num_last_winners, winners.length);
    }

    // Call this function to get to REVEAL_STATE
    function toRevealState() internal {
        game_state = GameState.REVEAL_STATE;
        game_state_debug = 1;
        final_commit_block = block.number;
    }

}