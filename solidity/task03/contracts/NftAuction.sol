// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IAuction {
    function initialize(address _seller, address token, uint amount) external;
}

// 实现一个简单的 NFT 拍卖合约
contract NftAuction {
    // 创建拍卖：允许用户将 NFT 上架拍卖
}
