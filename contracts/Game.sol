pragma solidity ^0.4.23;

import "./Ownable.sol";

contract Game is Ownable {

    enum GameState {COMMIT_STATE, REVEAL_STATE}
    GameState game_state = GameState.COMMIT_STATE;

    mapping (address => bytes32) public commits;
    mapping (address => uint) public game_data;
    address[] public player_addrs;
    address[] public winners;
    
    address OUR_ADDRESS = 0x2540099e9ed04aF369d557a40da2D8f9c2ab928D; //Address 
    uint public constant GAME_FEE_PERCENT = 1;
    uint public constant MAX_PLAYERS = 2;
    uint public constant MIN_GUESS = 0;
    uint public constant MAX_GUESS = 100; 

    uint public BET_SIZE = 1; //0.01 ether;
    uint public curr_number_bets = 0;
    uint public curr_number_reveals = 0;
    uint public number_reveals = 0;

    // Commit your guess. 
    function commit(bytes32 hashed_com) public payable{
        require(game_state == GameState.COMMIT_STATE);
        require(msg.value == BET_SIZE);

        commits[msg.sender] = hashed_com;
        curr_number_bets++;

        // If we received the MAX_PLAYER number of commits, it is time for
        // us to change state.
        if (curr_number_bets == MAX_PLAYERS) {
            game_state = GameState.REVEAL_STATE;
        }
    }

    function getSha(string one, string two) public view returns(bytes32) {
        return keccak256(one, two);
    }

    // returns (string)
    //DEBUG: remove return
    function reveal(string guess, string random) public  {
        require(game_state == GameState.REVEAL_STATE);
        
        uint guess_num = stringToUint(guess);
        
        require(guess_num >= MIN_GUESS && guess_num <= MAX_GUESS);

        // Check that the hashes match
        require(commits[msg.sender] == keccak256(guess, random));

        // When they do, we add the revealed guess to game data
        game_data[msg.sender] = guess_num;
        player_addrs.push(msg.sender);
        number_reveals++;

        if(curr_number_reveals == MAX_PLAYERS){
            find_winner();
        }
    }


    function find_winner() public {
        uint guess_sum = 0;
        for(uint i = 0; i < player_addrs.length; i++){
            uint tmp = game_data[player_addrs[i]];
            guess_sum += tmp;
        }

        uint twothirdsavg = ((guess_sum/player_addrs.length)*2)/3;

        uint min_diff = MAX_GUESS;
        for(uint j = 0; j < player_addrs.length; i++) {
            uint cur_guess = game_data[player_addrs[i]];
            uint cur_diff = MAX_GUESS;

            if(twothirdsavg > cur_guess){
                cur_diff = twothirdsavg - cur_guess;
            }
            else{
                cur_diff = cur_guess - twothirdsavg;
            }
            
            if(cur_diff < min_diff) {
                delete winners;
                winners.push(player_addrs[i]);
                min_diff = cur_diff;
            } else if(cur_diff == min_diff){
                winners.push(player_addrs[i]);
            }
        }

        uint gamefee = (address(this).balance/100) * GAME_FEE_PERCENT;

        OUR_ADDRESS.transfer(gamefee);

        uint prize = address(this).balance/winners.length;

        for(uint k = 0; k < winners.length; k++){
            winners[i].transfer(prize); 
        }
    }

    // Move to helper
    function bytes32ToString (bytes32 data) private returns (string) {
        bytes memory bytesString = new bytes(32);
        for (uint j = 0; j < 32 ;j++) {
            byte char = byte(bytes32(uint(data) * 2 ** (8 * j)));
            if (char != 0) {
                bytesString[j] = char;
            }
        }
        return string(bytesString);
    }

    function stringToUint(string s) constant private returns (uint) {
        bytes memory b = bytes(s);
        uint result = 0;
        for (uint i = 0; i < b.length; i++) { // c = b[i] was not needed
            if (b[i] >= 48 && b[i] <= 57) {
                result = result * 10 + (uint(b[i]) - 48); // bytes and int are not compatible with the operator -.
            }
        }
        return result; // this was missing
    }

}