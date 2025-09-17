// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol"

// 实现一个简单的 NFT 拍卖合约
contract NftAuction is Initializable {
    // 结构体定义
    struct Auction {
        address seller;     // 出售者
        uint256 duration;   // 拍卖持续时间
        uint256 startTime;  // 拍卖开始时间
        uint256 startPrice; // 拍卖开始价格

        bool ended;            // 拍卖是否结束
        uint256 highestBid;    // 最高出价
        address highestBidder; // 最高出价者

        address nftContract;   // NFT 合约地址
        uint256 nftId;         // NFT ID
    }

    // tokenId -> Auction 存储变量
    mapping(uint256 => Auction) public auctions;
    // 下一个拍卖Id
    uint256 public nextAuction;
    // 管理员地址
    address public admin;

    constructor() {
        admin = msg.sender;
    }

    // 创建拍卖
    function createAuction(uint256 _duration, uint256 _startPrice, address _nftAddress, uint256 _nftId) external {
        // 只能管理员创建拍卖
        require(msg.sender == admin, "only admin");
        // 拍卖时间必须大于0
        require(_duration > 0, "duration must be > 0");
        // 价格必须大于0
        require(_startPrice > 0, "startPrice must be > 0");
        // 创建拍卖
        auctions[nextAuction] = Auction({
            seller: msg.sender,
            duration: _duration,
            startPrice: _startPrice,
            startTime: block.timestamp,
            ended: false,
            highestBid: 0,
            highestBidder: address(0),
            nftContract: _nftAddress,
            nftId: _nftId
        });
        nextAuction++;
    }
 
    // 买家买单
    function placeBid(uint256 _auctionId) external payable {
        Auction storage auction = auctions[_auctionId];
        // 拍卖必须存在
        require(auction.seller != address(0), "auction not exist");
        // 拍卖未结束且当前时间小于拍卖结束时间
        require(!auction.ended && auction.startTime + auction.duration > block.timestamp,  "auction ended");
        // 出价必须高于起拍价和当前最高价
        require(msg.value > auction.startPrice && msg.value > auction.highestBid, "bid too low");

        // 如果有最高出价者，退还之前的最高出价
        if (auction.highestBidder != address(0)) {
            payable(auction.highestBidder).transfer(auction.highestBid);    
        }
        // 更新最高出价和最高出价者
        auction.highestBid = msg.value;
        auction.highestBidder = msg.sender;
    }
}
