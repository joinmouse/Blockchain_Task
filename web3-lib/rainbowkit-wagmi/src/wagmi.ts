import { getDefaultConfig } from "@rainbow-me/rainbowkit";
import { mainnet, polygon, optimism, sepolia } from "wagmi/chains";

export const config = getDefaultConfig({
  appName: "My RainbowKit App",
  projectId: "d4f094e9a01a1c260e8705cbe639a8a9",
  chains: [mainnet, sepolia, polygon, optimism],
});
