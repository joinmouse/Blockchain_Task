package main

import (
	"context"
	"fmt"

	"github.com/blocto/solana-go-sdk/client"
)

func main() {
	query()
}

func query() {
	c := client.NewClient("https://solana-mainnet.g.alchemy.com/v2/qgU3kZi4TGL1HMLEYAPGY")
	version, _ := c.GetVersion(context.Background())
	fmt.Println("Solana node version:", version.SolanaCore)

	// 获取最新区块
	recentBlock, err := c.GetBlock(context.Background(), 0) // 0表示最新区块
	if err != nil {
		panic("查询失败: " + err.Error())
	}

	fmt.Printf("区块高度: %d\n", recentBlock.BlockHeight)
	fmt.Printf("交易数量: %d\n", len(recentBlock.Transactions))
}
