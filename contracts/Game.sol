pragma solidity ^0.4.17;

import "./Ownable.sol";

contract Game is Ownable {

    mapping (address => string) private commits;
    mapping (address => string) private game_data;

    uint public BET_SIZE = 1;
    uint public curr_number_bets = 0;

    function commit(string com) public {
        commits[msg.sender] = com;

        
    }

    function reveal(string guess, uint random) public {
        
    }



}