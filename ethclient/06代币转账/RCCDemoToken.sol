// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyToken is ERC20, Ownable {
    // 定义兑换比率，即每 1 ETH 可以兑换 100000000 个 MyToken
    uint256 public constant RATE = 100000000; // 100000000 MyToken per 1 ETH
    // 定义最小以太币存款金额，用户至少需要发送 0.001 个以太币才能进行操作
    uint256 public constant MIN_ETH = 0.001 ether;

    /**
     * @dev 合约构造函数，初始化 ERC20 代币的名称和符号，并设置合约的所有者。
     * @param initialOwner 合约的初始所有者地址。但这里使用 Ownable(msg.sender) 实际上会将部署者设为所有者，与参数 initialOwner 可能不符，需注意。
     */
    constructor(address initialOwner) ERC20("RCCDemoToken", "RDT") Ownable(msg.sender) {
    }

    /**
     * @dev 允许用户通过发送以太币来铸造代币。
     * 要求用户发送的以太币数量至少为 MIN_ETH。
     * 铸造的代币数量根据发送的以太币数量和兑换比率 RATE 计算。
     */
    function mint() public payable {
        // 检查用户发送的以太币是否达到最小要求
        require(msg.value >= MIN_ETH, "Not enough ETH sent");
        // 计算需要铸造的代币数量
        uint256 tokensToMint = (msg.value * RATE);
        // 铸造代币并发送给调用者
        _mint(msg.sender, tokensToMint);
    }

    /**
     * @dev 允许合约所有者提取合约中所有的以太币余额。
     * 要求合约中存在以太币余额才能进行提取操作。
     */
    function withdrawETH() public onlyOwner {
        // 获取合约当前的以太币余额
        uint256 balance = address(this).balance;
        // 检查合约是否有以太币余额
        require(balance > 0, "No ETH to withdraw");
        // 将合约中的以太币余额转移给合约所有者
        payable(owner()).transfer(balance);
    }

    /**
     * @dev 接收以太币的 fallback 函数。
     * 当合约直接接收以太币时，会自动调用 mint 函数来铸造相应的代币。
     */
    receive() external payable {
        mint();
    }
}
