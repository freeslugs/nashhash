pragma solidity ^0.4.23;

import "truffle/Assert.sol";
import "./ThrowProxy.sol";
import "../contracts/LowestUniqueNum.sol";

// Solidity test contract, meant to test Thrower
contract TestThrower {

  // Truffle looks for `initialBalance` when it compiles the test suite 
  // and funds this test contract with the specified amount on deployment.
  uint public initialBalance = 10 ether;

  function testThrow() public{
    LowestUniqueNum lun = new LowestUniqueNum(10);
    ThrowProxy throwProxy = new ThrowProxy(address(lun)); //set Thrower as the contract to forward requests to. The target.

    //prime the proxy.
    bytes32 arg = keccak256("6");

    LowestUniqueNum(address(throwProxy)).commit.value(1 ether)(arg);
    //execute the call that is supposed to throw.
    //r will be false if it threw. r will be true if it didn't.
    //make sure you send enough gas for your contract method.
    bool r = throwProxy.execute.gas(200000)();

    //Dummy test
    //LowestUniqueNum(address(throwProxy)).throwReq();
    //bool r = throwProxy.execute.gas(200000)();

    Assert.isTrue(r, "Should be true");
  }
}
