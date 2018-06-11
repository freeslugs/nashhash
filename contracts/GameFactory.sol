pragma solidity ^0.4.23;

import "./LowestUniqueNum.sol"
import "./TwoThirdsAverage.sol"


contract GameFactory is Ownable{
    mapping(address => address) games;

    function createGame(string type) public onlyOwner{
    	if (keecak256(type) == keecak256("LUN"){
        	counters[msg.sender] = new LowestUn(msg.sender);
    	}
    	else if (keecak256(type) ==)
    }
    
    function increment() public {
        require (counters[msg.sender] != 0);
        Counter(counters[msg.sender]).increment(msg.sender);
    }
    
    function getCount(address account) public constant returns (uint) {
        if (counters[account] != 0) {
            return (Counter(counters[account]).getCount());
        }
    }
}