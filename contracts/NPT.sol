pragma solidity ^0.4.23;

import "openzeppelin-solidity/contracts/token/ERC20/RBACMintableToken.sol";
import "openzeppelin-solidity/contracts/ownership/Whitelist.sol";


contract NPT is RBACMintableToken, Whitelist {
    mapping (address => uint256) private balances; //Should it be private?
    mapping (address => mapping (address => uint256)) private allowed;
    uint256 public totalSupply;

    string public constant name = "NashPoints";
    string public constant symbol = "NPT";
    uint8 public constant decimals = 18;  // 18 is the most common number of decimal places

   	function transfer(address _to, uint256 _value) public returns (bool success) {
        if (balances[msg.sender] >= _value && _value > 0) {
            balances[msg.sender] -= _value;
            balances[_to] += _value;
            emit Transfer(msg.sender, _to, _value);
            return true;
        } else { 
            return false; 
        }
    }

    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
        if (balances[_from] >= _value && allowed[_from][msg.sender] >= _value && _value > 0) {
            balances[_to] += _value;
            balances[_from] -= _value;
            allowed[_from][msg.sender] -= _value;
            emit Transfer(_from, _to, _value);
            return true;
        } else { return false; }
    }

    function totalSupply() public view returns(uint256){
        return totalSupply;
    }

    function balanceOf(address _owner) public view returns (uint256 balance) {
        return balances[_owner];
    }

    function approve(address _spender, uint256 _value) public returns (bool success) {
        allowed[msg.sender][_spender] = _value;
        return true;
    }

    function allowance(address _owner, address _spender) public view returns (uint256 remaining) {
        return allowed[_owner][_spender];
    } 
}