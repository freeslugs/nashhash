pragma solidity ^0.4.23;

// Proxy contract for testing throws
contract ThrowProxy {
  address public target;
  bytes data;

  // Truffle looks for `initialBalance` when it compiles the test suite 
  // and funds this test contract with the specified amount on deployment.
  uint public initialBalance = 10 ether;

  constructor(address _target) public{
    target = _target;
  }

  //prime the data using the fallback function.
  function() public{
    data = msg.data;
  }

  function execute() public returns (bool) {
    //return target.call.value(1 ether)(data);
    return target.call(data);
  }
}