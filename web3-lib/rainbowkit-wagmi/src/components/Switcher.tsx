import React from "react";
import { useAccount, useSwitchChain } from "wagmi";

export const NetworkSwitcher: React.FC = () => {
  const { chain, chainId, isConnected } = useAccount();
  const { chains, switchChain } = useSwitchChain();
  console.log({ chain, chains }, "chain");
  
  return (
    <div className="flex flex-col items-center">
      <div>Current ChainId: {chainId}</div>
      <div>Current Chain Name: {chain?.name}</div>
      {isConnected && (
        <div className="">
          {chains
            .filter((v) => v.id !== chainId)
            .map((v) => (
              <button
                className=" bg-amber-400 p-2 m-2 rounded-md"
                onClick={() => switchChain({ chainId: v.id })}
                key={v.id}
              >
                switch to {v.name}
              </button>
            ))}
        </div>
      )}
    </div>
  );
};
