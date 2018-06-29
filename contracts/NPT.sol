pragma solidity ^0.4.23;

import "./RBACMintableToken.sol";

contract NPT is RBACMintableToken {
    string public constant name = "NashPoints";
    string public constant symbol = "NPT";
    uint8 public constant decimals = 18;  // 18 is the most common number of decimal places
}