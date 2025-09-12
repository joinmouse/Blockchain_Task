// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ERC721 {
    // Events
    event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);
    event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId);
    event ApprovalForAll(address indexed owner, address indexed operator, bool approved);

    string private _name;  // NFT名称
    string private _symbol; // NFT简称/符号

    mapping(uint256 => address) private _owners;    // 存储每个NFT的拥有者
    mapping(address => uint256) private _balances;  // 存储每个地址拥有的NFT数量

    // tokenId => approved address 单次授权: 针对特定代币的临时授权, 转移后应清除以避免重复使用
    mapping(uint256 => address) private _tokenApprovals; 
    // owner => operator => approved  全局授权: 是账户级别的长期授权，适用于所有的 NFT, 清除会影响其他代币的操作权限
    mapping(address => mapping(address => bool)) private _operatorApprovals; 

    uint256 private _totalSupply; // NFT总供应量


    constructor(string memory name_, string memory symbol_) {
        _name = name_;
        _symbol = symbol_;
    }

    // 下面三个是可选方法
    function name() public view returns (string memory) {
        return _name;
    }
    function symbol() public view returns (string memory) {
        return _symbol;
    }
    // 查询指定 NFT 的 token uri信息
    function tokenURI(uint256 _tokenId) public view returns (string memory) {
        require(_tokenId < _totalSupply, "Token ID does not exist");  // 检查Token ID是否存在
        return string(abi.encodePacked("https://api.example.com/metadata/", _tokenId));
    }

    // 统计某个地址拥有的所有 NFT 数量
    function balanceOf(address _owner) public view returns (uint256) {
        require(_owner != address(0), "Balance query for the zero address");  // 检查地址是否为零地址
        return _balance[_owner];  // 返回该地址拥有的NFT数量
    }
    // 查看某个NFT的拥有者
    function ownerOf(uint256 _tokenId) public view returns (address) {
        require(_tokenId < _totalSupply, "Token ID does not exist");
        return _owners[_tokenId];
    }


    // 单个NFT授权操作
    function approve(address _approved, uint256 _tokenId) external payable {
        require(_approved != address(0), "Approval to the zero address"); // 检查授权地址是否为零地址
        require(msg.sender == _owners[_tokenId], "Only the owner can approve");  // 检查调用者是否为NFT的拥有者
        _tokenApprovals[_tokenId] = _approved;   // 将tokenId的授权地址设置为_approved
    }
    function getApproved(uint256 _tokenId) public view returns (address) {
        require(_tokenId < _totalSupply, "Token ID does not exist"); // 检查Token ID是否存在
        return _tokenApprovals[_tokenId];  // 返回该tokenId的授权地址
    }

    // 对所有 NFT 的批量授权，一般用于交易所用, 注意这里授权并没有转移所有权，可以理解为一种中间的状态
    function setApprovalForAll(address _operator, bool _approved) external {
        require(_operator != address(0), "Approval to the zero address");
        _operatorApprovals[msg.sender][_operator] = _approved;
        emit ApprovalForAll(msg.sender, _operator, _approved);  // 触发批量授权事件
    }
    function isApprovedForAll(address _owner, address _operator) public view returns (bool) {
        return _operatorApprovals[_owner][_operator]; // 返回操作员的授权状态
    }


    // 转移
    // TransferFrom 将_tokenId 从_from地址转移到 _to 地址;
    function TransferFrom(address _from, address _to, uint256 _tokenId) external payable {
        // 检查msg.sender是否为拥有者，或被授权者
        require(_isApprovedOrOwner(msg.sender, _tokenId), "not approved nor owner");
        _transfer(_from, _to, _tokenId);  // 内部转移函数
    }
    // 将_tokenId 从_from地址转移到 _to 地址; 附加校验_to地址是否为合约地址;
    function safeTransferFrom(address _from, address _to, uint256 _tokenId) external payable {
        require(_isApprovedOrOwner(msg.sender, _tokenId), "not approved nor owner");
        _safeTransfer(_from, _to, _tokenId, "");
    }


    // ---------------------------
    // Mint / Burn（内部）
    // ---------------------------
    function _mint(address to, uint256 tokenId) internal {
        require(to != address(0), "mint to zero"); // 检查铸造地址是否为零地址
        require(!_exists(tokenId), "token already minted");

        // 更新NFT拥有者的数量和所有者映射
        _balances[to] += 1;
        _owners[tokenId] = to;

        emit Transfer(address(0), to, tokenId);  // 触发铸造事件
    }
    function _burn(uint256 tokenId) internal {
        address owner = ownerOf(tokenId); // 获取tokenId的拥有者
        require(owner != address(0), "token does not exist");  // owner必须是有效地址

        // 清除授权
        delete _tokenApprovals[tokenId];  // 清除tokenId的授权

        // 更新拥有者的数量和所有者映射
        _balances[owner] -= 1;
        delete _owners[tokenId]; // 清除tokenId的拥有者

        emit Transfer(owner, address(0), tokenId); // 触发销毁事件
    }

    // ---------------------------
    // Internal helpers
    // ---------------------------
    function _exists(uint256 tokenId) internal view returns (bool) {
        return _owners[tokenId] != address(0);  // 检查tokenId是否存在
    }
    function _isApprovedOrOwner(address spender, uint256 tokenId) internal view returns (bool) {
        require(_exists(tokenId), "token does not exist");
        address owner = ownerOf(tokenId);
        // 检查调用者是否为token的拥有者或被授权的操作员
        return (spender == owner || getApproved(tokenId) == spender || isApprovedForAll(owner, spender));
    }
    // 转移内部实现
    function _transfer(address from, address to, uint256 tokenId) internal {
        require(ownerOf(tokenId) == from, "from is not owner");  // 检查转出地址是否为token的拥有者
        require(to != address(0), "transfer to zero"); // 检查转入地址是否为零地址

        // clear approvals
        delete _tokenApprovals[tokenId];   // 清除授权

        // update balances and owner
        _balances[from] -= 1;  // 更新转出地址的NFT数量
        _balances[to] += 1;    // 更新转入地址的NFT数量
        _owners[tokenId] = to; // 更新tokenId的拥有者

        emit Transfer(from, to, tokenId);  // 触发转账事件
    }

    function _safeTransfer(address from, address to, uint256 tokenId, bytes memory data) internal {
        _transfer(from, to, tokenId);  // 调用内部转移函数
        // 如果 to 是合约则调用 IERC721Receiver.onERC721Received
        if (_isContract(to)) {
            // 调用合约的 onERC721Received 方法,并检查返回值是否正确
            try IERC721Receiver(to).onERC721Received(msg.sender, from, tokenId, data) returns (bytes4 retval) {
                require(retval == IERC721Receiver.onERC721Received.selector, "unsafe recipient");
            } catch {
                revert("unsafe recipient");
            }
        }
    }

    // 检查地址是否为合约地址
    function _isContract(address account) internal view returns (bool) {
        return account.code.length > 0;
    }
}


// @notice 极简 IERC721Receiver 接口，用于 safeTransfer 检测
interface IERC721Receiver {
    function onERC721Received(address operator, address from, uint256 tokenId, bytes calldata data) external returns (bytes4);
}
