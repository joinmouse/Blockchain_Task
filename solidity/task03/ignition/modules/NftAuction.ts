import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

export default buildModule("NftAuctionModule", (m) => {
  const nftAuction = m.contract("NftAuction");

  return { nftAuction };
});
