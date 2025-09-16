// 写一个本地部署的 deploy 脚本
import { network } from "hardhat";

const { viem } = await network.connect({
  network: "localhost",
  chainType: "l1",
});

const counter = await viem.deployContract("Counter");

viem.getPublicClient().then((publicClient) => {
  console.log("Contract deployed at:", counter.address);
});

await counter.write.inc();
await counter.write.inc();

console.log("Counter value:", await counter.read.x());
