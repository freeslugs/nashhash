pragma solidity ^0.4.23;

import "./LowestUniqueNum.sol";
import "./TwoThirdsAverage.sol";


contract GameFactory is Ownable{
    address[] public games;

    function createGame(string typ,
				    	address _feeAddress,
				        uint _gameFeePercent,
				        uint _stakeSize,
				        uint _maxp, 
				        uint _gameStageLength
				        ) public onlyOwner{

    	if (keccak256(typ) == keccak256("TTA")){
        	games.push(new TwoThirdsAverage(_feeAddress, _gameFeePercent, _stakeSize, _maxp, _gameStageLength));
    	}
    	else if (keccak256(typ) == keccak256("LUN")){
    		games.push(new LowestUniqueNum(_feeAddress, _gameFeePercent, _stakeSize, _maxp, _gameStageLength));
    	}
    }
    
    function getGames() public view returns (address[]) {
    	return games;
    }
}