import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.jsx";

import "@rainbow-me/rainbowkit/styles.css";
import { RainbowKitProvider, midnightTheme } from "@rainbow-me/rainbowkit";
import { WagmiProvider } from "wagmi";
import { QueryClientProvider, QueryClient } from "@tanstack/react-query";

import { config } from "./wagmi.ts";

const client = new QueryClient();

createRoot(document.getElementById("root")).render(
  <WagmiProvider config={config}>
    <QueryClientProvider client={client}>
      <RainbowKitProvider theme={midnightTheme()}>
        <App />
      </RainbowKitProvider>
    </QueryClientProvider>
  </WagmiProvider>
);
