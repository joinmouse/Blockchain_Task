// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Voting {
    // 使用mapping存储每个候选人的得票数
    // key: 候选人名字（字符串）
    // value: 得票数（无符号整数）
    mapping(string => uint256) private votes;
    
    // 定义投票事件，当有人投票时触发
    // 记录被投票的候选人和新的得票数
    event VoteCast(string candidate, uint256 newVoteCount);
    
    // 投票函数
    // 参数：candidate - 候选人名字
    // 功能：给指定候选人的得票数加1
    function vote(string memory candidate) public {
        votes[candidate] += 1;
        emit VoteCast(candidate, votes[candidate]);  // 合约中发生的重要操作记录到区块链上
    }
    
    // 查询得票数函数
    // 参数：candidate - 候选人名字
    // 返回值：该候选人的得票数
    function getVotes(string memory candidate) public view returns (uint256) {
        return votes[candidate];
    }
    
    // 重置单个候选人的得票数
    // 参数：candidate - 候选人名字
    // 功能：将该候选人的得票数设为0
    function resetVotes(string memory candidate) public {
        votes[candidate] = 0;
    }
    
    // 重置所有候选人的得票数
    // 参数：candidates - 候选人名字数组
    // 功能：将所有指定候选人的得票数设为0
    // 注意：此函数需要预先知道所有候选人的名字
    function resetAllVotes(string[] memory candidates) public {
        for (uint i = 0; i < candidates.length; i++) {
            votes[candidates[i]] = 0;
        }
    }
}
