// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";
import "@uniswap/v2-core/contracts/interfaces/IUniswapV2Factory.sol";
import "@uniswap/v2-core/contracts/interfaces/IUniswapV2Pair.sol";

contract MemeToken is ERC20, ERC20Burnable, Ownable(address(msg.sender)) {
    // 代币税相关参数
    uint256 public buyTaxRate = 5; // 购买税率 (5%)
    uint256 public sellTaxRate = 10; // 出售税率 (10%)
    uint256 public maxTaxRate = 15; // 最大税率限制
    
    // 流动性相关参数
    address public liquidityReceiver;
    uint256 public liquidityFee = 3; // 流动性费用占税收的比例 (3%)
    
    // 交易限制相关参数
    uint256 public maxTxAmount; // 最大单笔交易限额
    uint256 public maxWalletAmount; // 最大钱包持有限额
    mapping(address => uint256) public userLastTradeTime;
    uint256 public tradeCooldown = 30 seconds; // 交易冷却时间
    
    // 排除税收和限制的地址
    mapping(address => bool) public isExcludedFromTax;
    mapping(address => bool) public isExcludedFromLimits;
    
    // 流动性池和Uniswap相关
    IUniswapV2Router02 public uniswapV2Router;
    address public uniswapV2Pair;
    bool inSwapAndLiquify;
    bool public swapAndLiquifyEnabled = true;
    uint256 public swapThreshold;
    
    // 事件定义
    event SwapAndLiquify(uint256 tokensSwapped, uint256 ethReceived, uint256 tokensIntoLiqudity);
    event TaxRateUpdated(uint256 buyTax, uint256 sellTax);
    event TransactionLimitUpdated(uint256 maxTxAmount, uint256 maxWalletAmount);
    
    modifier lockTheSwap {
        inSwapAndLiquify = true;
        _;
        inSwapAndLiquify = false;
    }
    
    constructor() ERC20("MemeToken", "MEME") {
        // 设置Uniswap路由器和工厂地址（以太坊主网）
        IUniswapV2Router02 _uniswapV2Router = IUniswapV2Router02(0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D);
        uniswapV2Router = _uniswapV2Router;
        
        // 创建交易对
        uniswapV2Pair = IUniswapV2Factory(_uniswapV2Router.factory())
            .createPair(address(this), _uniswapV2Router.WETH());
        
        // 初始化参数
        uint256 totalSupply = 1000000000 * 10 ** decimals();
        maxTxAmount = totalSupply * 1 / 100; // 1% of total supply
        maxWalletAmount = totalSupply * 2 / 100; // 2% of total supply
        swapThreshold = totalSupply * 1 / 1000; // 0.1% of total supply
        
        // 设置接收地址
        liquidityReceiver = owner();
        
        // 排除初始地址的税收和限制
        isExcludedFromTax[owner()] = true;
        isExcludedFromTax[address(this)] = true;
        isExcludedFromTax[address(uniswapV2Router)] = true;
        
        isExcludedFromLimits[owner()] = true;
        isExcludedFromLimits[address(this)] = true;
        isExcludedFromLimits[address(uniswapV2Router)] = true;
        
        // 铸造初始代币
        _mint(owner(), totalSupply);
    }
    
    // 重写_transfer函数以实现税收和交易限制
    function _transfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual override {
        require(from != address(0), "ERC20: transfer from the zero address");
        require(to != address(0), "ERC20: transfer to the zero address");
        require(amount > 0, "Transfer amount must be greater than zero");
        
        // 检查交易限制
        if (!isExcludedFromLimits[from] && !isExcludedFromLimits[to]) {
            // 检查单笔交易限额
            require(amount <= maxTxAmount, "Transfer amount exceeds the maxTxAmount.");
            
            // 检查钱包持有限额（不包括流动性池）
            if (to != uniswapV2Pair) {
                require(balanceOf(to) + amount <= maxWalletAmount, "New balance would exceed the maxWalletAmount.");
            }
            
            // 检查交易冷却时间
            if (from == uniswapV2Pair && !isExcludedFromLimits[to]) {
                require(block.timestamp - userLastTradeTime[to] >= tradeCooldown, "Trade cooldown not expired");
                userLastTradeTime[to] = block.timestamp;
            }
        }
        
        // 计算基础税额
        uint256 taxAmount = 0;
        bool isBuy = from == uniswapV2Pair;
        bool isSell = to == uniswapV2Pair;
        
        // 应用税收（除了排除的地址）
        if (!isExcludedFromTax[from] && !isExcludedFromTax[to]) {
            if (isBuy) {
                taxAmount = amount * buyTaxRate / 100;
            } else if (isSell) {
                taxAmount = amount * sellTaxRate / 100;
            }
            
            if (taxAmount > 0) {
                super._transfer(from, address(this), taxAmount);
                amount = amount - taxAmount;
            }
        }
        
        // 处理流动性
        if (!inSwapAndLiquify && 
            swapAndLiquifyEnabled && 
            taxAmount > 0 &&
            from != uniswapV2Pair &&
            balanceOf(address(this)) >= swapThreshold
        ) {
            swapAndLiquify(swapThreshold);
        }
        
        // 执行实际转账
        super._transfer(from, to, amount);
    }
    
    // 将代币转换为ETH并添加流动性
    function swapAndLiquify(uint256 contractTokenBalance) private lockTheSwap {
        // 计算流动性份额
        uint256 tokensForLiquidity = contractTokenBalance * liquidityFee / 100;
        uint256 tokensForSwap = contractTokenBalance - tokensForLiquidity;
        
        // 记录初始ETH余额
        uint256 initialBalance = address(this).balance;
        
        // 交换代币为ETH
        swapTokensForEth(tokensForSwap);
        
        // 计算获得的ETH
        uint256 ethForLiquidity = address(this).balance - initialBalance;
        
        // 添加流动性
        if (ethForLiquidity > 0) {
            addLiquidity(tokensForLiquidity, ethForLiquidity);
            emit SwapAndLiquify(tokensForSwap, ethForLiquidity, tokensForLiquidity);
        }
    }
    
    // 交换代币为ETH
    function swapTokensForEth(uint256 tokenAmount) private {
        // 构建路径：代币 -> WETH
        address[] memory path = new address[](2);
        path[0] = address(this);
        path[1] = uniswapV2Router.WETH();
        
        _approve(address(this), address(uniswapV2Router), tokenAmount);
        
        // 执行交换
        uniswapV2Router.swapExactTokensForETHSupportingFeeOnTransferTokens(
            tokenAmount,
            0, // 接受任何数量的ETH
            path,
            address(this),
            block.timestamp
        );
    }
    
    // 添加流动性
    function addLiquidity(uint256 tokenAmount, uint256 ethAmount) private {
        // 批准合约使用代币
        _approve(address(this), address(uniswapV2Router), tokenAmount);
        
        // 添加流动性
        uniswapV2Router.addLiquidityETH{value: ethAmount}(
            address(this),
            tokenAmount,
            0, // 滑点设置
            0, // 滑点设置
            liquidityReceiver,
            block.timestamp
        );
    }
    
    // 允许合约接收ETH
    receive() external payable {}
    
    // 管理函数 - 更新税率
    function updateTaxRates(uint256 _buyTax, uint256 _sellTax) external onlyOwner {
        require(_buyTax <= maxTaxRate && _sellTax <= maxTaxRate, "Tax rate exceeds maximum");
        buyTaxRate = _buyTax;
        sellTaxRate = _sellTax;
        emit TaxRateUpdated(_buyTax, _sellTax);
    }
    
    // 管理函数 - 更新交易限制
    function updateTransactionLimits(uint256 _maxTxAmount, uint256 _maxWalletAmount) external onlyOwner {
        maxTxAmount = _maxTxAmount;
        maxWalletAmount = _maxWalletAmount;
        emit TransactionLimitUpdated(_maxTxAmount, _maxWalletAmount);
    }
    
    // 管理函数 - 更新流动性接收地址
    function updateLiquidityReceiver(address _newReceiver) external onlyOwner {
        liquidityReceiver = _newReceiver;
    }
    
    // 管理函数 - 启用/禁用自动流动性
    function setSwapAndLiquifyEnabled(bool _enabled) external onlyOwner {
        swapAndLiquifyEnabled = _enabled;
    }
    
    // 管理函数 - 从税收中排除地址
    function excludeFromTax(address _address, bool _exclude) external onlyOwner {
        isExcludedFromTax[_address] = _exclude;
    }
    
    // 管理函数 - 从限制中排除地址
    function excludeFromLimits(address _address, bool _exclude) external onlyOwner {
        isExcludedFromLimits[_address] = _exclude;
    }
    
    // 紧急提取ETH（如果有）
    function emergencyWithdrawETH(uint256 amount) external onlyOwner {
        payable(owner()).transfer(amount);
    }
    
    // 紧急提取代币
    function emergencyWithdrawTokens(address tokenAddress, uint256 amount) external onlyOwner {
        IERC20(tokenAddress).transfer(owner(), amount);
    }
}
