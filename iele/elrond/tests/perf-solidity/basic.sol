pragma solidity ^0.4.0;

contract BasicPerformanceTester {

    function testNop(int exponent, int seed, uint n) external returns (int) {
        for (uint i = 0; i < n; i += 1) {}
        return seed;
    }

    function testDivAdd(uint x, uint y, uint k, uint n) external returns (uint) {
        var r = x;
        for (uint i = 0; i < n; i += 1) {
            r /= y;
            r += k;
        }
        return r;
    }
   
}
