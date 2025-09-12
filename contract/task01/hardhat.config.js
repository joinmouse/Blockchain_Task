require("@nomicfoundation/hardhat-toolbox");

module.exports = {
  solidity: "0.8.20", // 与合约版本匹配
  paths: {
    sources: "./contracts", // 合约目录
    artifacts: "./artifacts", // 编译输出目录
  },
};
