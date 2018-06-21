pragma solidity ^0.4.23;


/**
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
contract GameHelper {

    function getSha(string one, string two) internal pure returns(bytes32) {
        return keccak256(one, two);
    }

    // Move to helper
    function bytes32ToString (bytes32 data) internal pure returns (string) {
        bytes memory bytesString = new bytes(32);
        for (uint j = 0; j < 32 ;j++) {
            byte char = byte(bytes32(uint(data) * 2 ** (8 * j)));
            if (char != 0) {
                bytesString[j] = char;
            }
        }
        return string(bytesString);
    }

    function stringToUint(string s) internal pure returns (uint) {
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