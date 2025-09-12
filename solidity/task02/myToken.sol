// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

// 以下代码定义了一个名为 MyToken 的合约，它继承自 OpenZeppelin 库中的 ERC20 合约，用于创建一个自定义的 ERC20 代币。

// ERC20 是以太坊上的一个标准代币接口，定义了一系列方法和事件，使得不同的代币合约可以遵循统一的规范。
// 这些方法包括转账（transfer）、授权（approve）、查询余额（balanceOf）等，方便开发者在不同的应用中使用和交互。

// 定义 MyToken 合约，继承自 ERC20 合约
contract MyToken is ERC20 {
    // 定义一个公开的地址类型变量 owner，用于存储合约的所有者地址
    address public owner;

    // 合约的构造函数，在合约部署时执行
    constructor() ERC20("MyToken", "MTK") {
        // 将合约部署者的地址赋值给 owner 变量
        owner = msg.sender;
        // 调用 ERC20 合约的 _mint 方法，为合约部署者铸造 1000000 个代币（考虑小数位数）
        _mint(msg.sender, 1000000 * 10 ** uint(decimals()));
    }

    // 定义一个公开的 mint 方法，用于铸造新的代币
    function mint(address account, uint256 amount) public {
        // 检查调用者是否为合约所有者，如果不是则抛出错误信息
        require(msg.sender == owner, "Only owner can mint tokens");
        // 调用 ERC20 合约的 _mint 方法，为指定账户铸造指定数量的代币
        _mint(account, amount);
    }
}
