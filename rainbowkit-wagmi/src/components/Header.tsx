import { ConnectButton } from "@rainbow-me/rainbowkit";
import React from "react";

export const Header: React.FC = () => {
  return (
    <div className="flex justify-between">
      <div> Dapp Frontend </div>
      <div>
        <ConnectButton />
      </div>
    </div>
  );
};
