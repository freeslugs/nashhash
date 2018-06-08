pragma solidity ^0.4.23;

//import "openzeppelin-solidity/contracts/token/ERC20/MintableToken.sol";
//import "openzeppelin-solidity/contracts/ownership/Whitelist.sol";


contract NPT {// is MintableToken, Whitelist {
    mapping (address => uint256) private balances; //Should it be private?
    mapping (address => mapping (address => uint256)) private allowed;
    uint256 public totalSupply;

   	function transfer(address _to, uint256 _value) returns (bool success) {
        if (balances[msg.sender] >= _value && _value > 0) {
            balances[msg.sender] -= _value;
            balances[_to] += _value;
            //Transfer(msg.sender, _to, _value);
            return true;
        } else { return false; }
    }
}