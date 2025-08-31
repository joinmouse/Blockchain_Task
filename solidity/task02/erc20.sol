// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleERC20 {
    string private _name;
    string private _symbol;

    uint256 private _totalSupply; // 初始供应量

    mapping(address => uint256) private _balances; // 记录每个地址的代币余额
    mapping(address => mapping(address => uint256)) private _allowances; // 记录每个地址对其他地址的授权额度

    // 事件
    event Transfer(address indexed from, address indexed to, uint256 value); // 触发转账事件
    event Approval(address indexed owner, address indexed spender, uint256 value); // 触发授权事件

    // 临时且非持久: memory 中的数据只存在于一次外部调用/内部调用的执行期间，函数返回后就会被丢弃，不会写入链上持久存储。
    constructor(string memory name_, string memory symbol_, uint256 initialSupply_) {
        _name = name_;
        _symbol = symbol_;
        _totalSupply = initialSupply_; // 初始化代币总量

        // 初始化的供应量分配给合约的部署者
        _balances[msg.sender] = _totalSupply;
    }

    // 代币的全名, 例如"USD Coin"
    function name() public view returns (string memory) {
        return _name;
    }
    // 代币简称/符号, 例如"USDC"
    function symbol() public view returns (string memory) {
        return _symbol;
    }
    // 代币总量
    function totalSupply() public view returns (uint256) {
        return _totalSupply;
    }

    // 查询某个地址的代币余额
    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

    // transfer 转账
    function transfer(address to_, uint256 amount_) public returns (bool) {
        address owner = msg.sender; // 获取调用者地址
        require(owner != address(0), "Transfer from the zero address"); // 检查调用者地址是否为零地址
        require(to_ != address(0), "Transfer to the zero address"); // 检查接收者地址是否为零地址
        require(_balances[owner] >= amount_, "Transfer amount exceeds balance"); // 检查调用者余额是否足够

        _balances[owner] -= amount_; // 从调用者余额中扣除转账金额
        _balances[to_] += amount_; // 将转账金额添加到接收者余额中

        emit Transfer(owner, to_, amount_); // 触发转账事件, 记录到日志中

        return true; // 返回成功标志
    }

    // approve 授权
    function approve(address spender, uint256 amount) public returns (bool) {
        address owner = msg.sender;  // 获取调用者地址
        require(owner != address(0), "Approve from the zero address");
        require(spender != address(0), "Approve to the zero address");

        _allowances[owner][spender] = amount;  // 设置授权额度

        emit Approval(owner, spender, amount);
        return true;
    }

    // 查询授权额度
    function allowance(address owner, address spender) public view returns (uint256) {
        return _allowances[owner][spender];
    }

    // transferFrom 转账
    function transferFrom(address from, address to, uint256 amount) public returns (bool) {
        address spender = msg.sender; // 获取调用者地址
        require(from != address(0), "Transfer from the zero address");  // 检查发送者地址是否为零地址
        require(to != address(0), "Transfer to the zero address"); // 检查接收者地址是否为零地址
        require(_balances[from] >= amount, "Transfer amount exceeds balance"); // 检查发送者余额是否足够
        require(_allowances[from][spender] >= amount, "Transfer amount exceeds allowance"); // 检查授权额度是否足够

        _balances[from] -= amount;  // 从发送者余额中扣除转账金额
        _balances[to] += amount;    // 将转账金额添加到接收者余额中

        _allowances[from][spender] -= amount;  // 扣除授权额度

        emit Transfer(from, to, amount);
        return true;
    }
}
