// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC777/ERC777.sol";

contract MyERC777 is ERC777 {
    constructor(
        uint256 initialSupply,
        address[] memory defaultOperators
    ) ERC777("MyERC777", "M777", defaultOperators) {
        _mint(msg.sender, initialSupply, "", "");
    }
}
