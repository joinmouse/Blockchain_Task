import React from "react";
import { useAccount, useBalance } from "wagmi";

export const Info: React.FC = () => {
  const { address } = useAccount();
  const { data, error } = useBalance({ address });
  console.log("data", address, data);

  return (
    <div>
      <div>Address: {address}</div>
      <div>ETH Balance: {data?.value}</div>
    </div>
  );
};

export default Info;
