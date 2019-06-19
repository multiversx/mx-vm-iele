# How to call Solidity to IELE compiler

Git clone https://github.com/ElrondNetwork/solidity
Follow instructions to build compiler
The compiler executable `isolc` can be found in `<solidity_repository>/build/solc`

Run the following:

```
isolc --asm auction.sol -o . --overwrite
```
