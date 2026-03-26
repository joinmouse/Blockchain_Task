import { encodeFunctionData } from "viem";
import { network }  from "hardhat";

async function main() {
  // 连接到本地网络
  const { viem } = await network.connect({
    network: "localhost",
    chainType: "l1",
  });
  
  // 部署逻辑合约（V1版本）
  console.log("Deploying NftAuctionV1 implementation...");
  const nftAuctionV1 = await viem.deployContract("NftAuction");
  console.log(`NftAuctionV1 deployed to: ${nftAuctionV1.address}`);

  // 部署 ProxyAdmin
  const proxyAdmin = await viem.deployContract("ProxyAdmin");
  console.log(`ProxyAdmin deployed to: ${proxyAdmin.address}`)

  // 3. 编码初始化函数数据
  const initializerData = encodeFunctionData({
    abi: [
      {
        name: "initialize",
        type: "function",
        inputs: [],
      },
    ],
    functionName: "initialize",
  });
  console.log(`初始化数据: ${initializerData}`);

  // 部署透明代理
  const transparentProxy = await viem.deployContract("TransparentUpgradeableProxy", [
    nftAuctionV1.address,
    proxyAdmin.address,
    initializerData    // 初始化数据
  ]);
  console.log(`TransparentUpgradeableProxy deployed to: ${transparentProxy.address}`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });



