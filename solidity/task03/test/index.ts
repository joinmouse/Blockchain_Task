import assert from "node:assert/strict";
import { describe, it } from "node:test";
import { network } from "hardhat";

describe("Local Deployment", () => {
  it("should deploy Counter contract and increment value", async () => {
    const { viem } = await network.connect();

    const nftAuction = await viem.deployContract("NftAuction");

    // 创建一个 nft
    await nftAuction.write.createAuction([
      BigInt(60 * 1000),
      BigInt(100 * 1000),
      "0x0000000000000000000000000000000000000000",
      BigInt(1),
    ]);

    const auction = await nftAuction.read.auctions([BigInt(0)]);

    console.log({ auction });
  });
});
