// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// 打赏合约: BeggingContract
// 1、记录每个捐赠者的地址和捐赠金额。
// 2、允许合约所有者提取所有捐赠的资金。

// 分析实现
// 一个 mapping 来记录每个捐赠者的捐赠金额。
// 一个 donate 函数，允许用户向合约发送以太币，并记录捐赠信息。
// 一个 withdraw 函数，允许合约所有者提取所有资金。
// 一个 getDonation 函数，允许查询某个地址的捐赠金额。
// 使用 payable 修饰符和 address.transfer 实现支付和提款。

contract BeggingContract {
    address public immutable owner;   // 记录合约所有者地址

    mapping(address => uint256) private _donations;  // 记录捐赠者地址和金额

    event Donate(address indexed from, uint256 value);

    constructor() {
        owner = msg.sender;
    }

    // donate 捐赠函数
    function donate() external payable {
        _donations[msg.sender] += msg.value;  // 记录
        // 记录事件
        emit Donate(msg.sender, msg.value);
    }

    // getDonation 查询某个地址的捐赠金额
    function getDonation() external view returns (uint256) {
        return _donations[msg.sender];
    }

    // 提现
    function withdraw() external {
        require(msg.sender == owner, "only owner opertion"); // 仅仅合约所有者可以提现
        payable(owner).transfer(address(this).balance);  // 合约账户（address(this)）中的所有余额，转移到合约部署者的个人钱包地址（owner）
        
    }
}