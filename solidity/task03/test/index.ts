import assert from "node:assert/strict";
import { describe, it } from "node:test";
import { network } from "hardhat";
import { encodeFunctionData, getContract, getAddress } from "viem";

describe("NftAuction via Proxy", () => {
  it("should deploy NftAuction via proxy and create auction", async () => {
    const { viem } = await network.connect();

    // 1. 部署逻辑合约
    const nftAuctionImpl = await viem.deployContract("NftAuction");

    // 2. 编码初始化数据
    const initData = encodeFunctionData({
      abi: nftAuctionImpl.abi,
      functionName: "initialize",
    });

    // 3. 使用 TestProxy（继承 ERC1967Proxy）部署代理
    const proxy = await viem.deployContract("TestProxy", [
      nftAuctionImpl.address,
      initData,
    ]);

    // 4. 通过代理地址访问 NftAuction
    const publicClient = await viem.getPublicClient();
    const walletClient = await viem.getWalletClient();

    const nftAuction = getContract({
      address: proxy.address,
      abi: nftAuctionImpl.abi,
      client: { public: publicClient, wallet: walletClient },
    });

    // 5. 验证 admin 已设置
    const admin = await nftAuction.read.admin();
    console.log("admin:", admin);

    // 6. 创建拍卖 - 直接使用 walletClient.writeContract 以包含 account
    const [account] = await walletClient.getAddresses();
    await walletClient.writeContract({
      address: proxy.address,
      abi: nftAuctionImpl.abi,
      functionName: "createAuction",
      args: [
        BigInt(60 * 1000),
        BigInt(100 * 1000),
        "0x0000000000000000000000000000000000000000",
        BigInt(1),
      ],
      account,
    });

    const auction = await nftAuction.read.auctions([BigInt(0)]);
    console.log({ auction });

    // 7. 验证拍卖创建成功
    assert.equal(auction[1], BigInt(60 * 1000), "duration should match");
    assert.equal(auction[3], BigInt(100 * 1000), "startPrice should match");
    assert.equal(auction[4], false, "auction should not be ended");
  });
});
