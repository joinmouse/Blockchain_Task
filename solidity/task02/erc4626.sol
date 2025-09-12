// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";


// @dev 实现ERC4626代币化Vault标准: 允许用户存入基础资产并获得相应的份额代币
contract ERC4626 is ERC20, Ownable, ReentrancyGuard {
    using SafeERC20 for ERC20; // 安全的ERC20操作

    // 存款事件
    event Deposit(address indexed caller, address indexed receiver, uint256 assets, uint256 shares);
    
    // 取款事件
    event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares);

    // 基础资产代币
    IERC20 public immutable asset;

    // 构造函数，初始化Vault名称、符号和基础资产
    // @param asset_ 基础资产合约地址
    constructor(
        IERC20 asset_,
        string memory name_,
        string memory symbol_
    ) ERC20(name_, symbol_) Ownable(msg.sender) {
        require(address(asset_) != address(0), "ERC4626: asset is zero address");
        asset = asset_;
    }

    // 基础资产总数量
    function totalAssets() public view returns (uint256) {
        return asset.balanceOf(address(this));
    }

    // 计算资产转换为份额的比例
    function convertToShares(uint256 assets) public view returns (uint256) {
        uint256 totalSupply = totalSupply();  // 当前 Vault 中所有用户持有的份额代币总和
        uint256 totalAssets = totalAssets(); // 当前 Vault 中所有用户存入的资产总和
        if (supply == 0) {
            // Vault 初始为空时，公式会简化为 shares = assets（1:1 兑换）
            return assets;
        } else {
            // 计算份额，使用比例公式
            return (assets / totalAssets) * totalSupply;
        }
    }
    // 计算份额转换为资产的比例
    function convertToAssets(uint256 shares) public view returns (uint256) {
        uint256 totalSupply = totalSupply();
        if (totalSupply == 0) {
            return shares;
        } else {
            return (shares / totalSupply) * totalAssets();
        }
    }


    // 预览存入资产所获得的份额
    function previewDeposit(uint256 assets) public view returns (uint256) {
        return convertToShares(assets);
    }
    // 最大可存入资产
    function maxDeposit(address receiver) public view returns (uint256) {
        return type(uint256).max;
    }
    // 存入资产
    function deposit(uint256 assets, address receiver) public nonReentrant returns (uint256) {
        require(assets > 0, "ERC4626: assets is zero");
        require(receiver != address(0), "ERC4626: receiver is zero address");
        require(assets <= maxDeposit(receiver), "ERC4626: deposit exceeds max");

        uint256 shares = previewDeposit(assets); // 计算份额
        
        // 从发送者转移资产到合约
        asset.safeTransferFrom(msg.sender, address(this), assets);  // 发送资产到 vault 合约中

        // 铸造份额给接收者
        _mint(receiver, shares);      // 铸造份额给接收者

        emit Deposit(msg.sender, receiver, assets, shares);  // 事件上报
         
        return shares;
    }

    // 最大提取资产
    function maxWithdraw(address owner) public view returns (uint256) {
        return convertToAssets(balanceOf(owner));
    }
    // 预览提取资产所需的份额
    function previewWithdraw(uint256 assets) public view returns (uint256) {
        return convertToShares(assets);
    }
    // 提取
    function withdraw(uint256 assets, address receiver, address owner) 
        public 
        nonReentrant 
        returns (uint256) 
    {
        require(assets > 0, "ERC4626: assets is zero");
        require(receiver != address(0), "ERC4626: receiver is zero address");
        require(owner != address(0), "ERC4626: owner is zero address");
        require(assets <= maxWithdraw(owner), "ERC4626: withdraw exceeds max");

        uint256 shares = previewWithdraw(assets);  // 预览提取资产所需的份额

        // 检查提取者是否为所有者
        if (msg.sender != owner) {
            uint256 allowed = allowance(owner, msg.sender);
            require(allowed >= shares, "ERC4626: insufficient allowance");
            _approve(owner, msg.sender, allowed - shares);
        }

        _burn(owner, shares);    // 销毁份额

        // 转移资产给接收者
        asset.safeTransfer(receiver, assets);    // 转移资产给接收者

        emit Withdraw(msg.sender, receiver, owner, assets, shares);  // 事件上报

        return shares;
    }
}
