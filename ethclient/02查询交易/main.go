package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/U3kArMqv4LilPctZa0_GvGpcO9OGaORP")
	if err != nil {
		log.Fatal(err)
	}

	// 调用 client 的 ChainID 方法，传入上下文，获取当前区块链的链 ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// blockNumber 表示要查询的以太坊区块的编号，这里指定为 5671744，后续会使用该编号调用 client.BlockByNumber 方法获取该编号对应的完整区块信息
	blockNumber := big.NewInt(5671744)
	// 取到完整的区块信息
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	// 打印区块的哈希值，哈希值是区块的唯一标识符
	fmt.Println("区块哈希值", block.Hash().Hex()) // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5

	// 遍历区块中的所有交易
	for _, tx := range block.Transactions() {
		// 打印交易的哈希值，哈希值是交易的唯一标识符
		fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		// 打印交易发送的以太币数量，以 Wei 为单位
		fmt.Println(tx.Value().String())
		// 打印交易执行所需的 Gas 数量
		fmt.Println(tx.Gas())
		// 打印每单位 Gas 的价格，单位为 Wei
		fmt.Println(tx.GasPrice().Uint64())
		// 打印发送者账户的交易序列号，用于防止重放攻击
		fmt.Println(tx.Nonce())
		// 打印交易携带的数据，通常用于智能合约调用
		fmt.Println(tx.Data())
		// 打印交易的接收者地址
		fmt.Println(tx.To().Hex())
		fmt.Println("chainID", chainID)
		// 使用 types.NewEIP155Signer 根据链 ID 创建一个 EIP-155 签名者，然后调用 types.Sender 方法从交易中恢复发送者的地址
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("sender", sender.Hex()) // 0x2CdA41645F2dBffB852a605E92B185501801FC28
		} else {
			log.Fatal(err)
		}

		// 调用 client 的 TransactionReceipt 方法，传入上下文和交易的哈希值，获取该交易的收据信息
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		// 打印交易的执行状态，1 表示成功，0 表示失败
		fmt.Println(receipt.Status) // 1
		// 打印交易产生的日志信息，通常在智能合约调用时会产生日志
		fmt.Println(receipt.Logs) // []
		// 跳出当前的交易遍历循环
		break
	}

	// 调用 TransactionCount 来了解块中有多少个事务
	blockHash := common.HexToHash(block.Hash().Hex())
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("交易数量", count)
	var transactionInBlock string
	for idx := uint(0); idx < 1; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tx.Hash().Hex())
		transactionInBlock = tx.Hash().Hex()
	}

	// TransactionByHash 在给定具体事务哈希值的情况下直接查询单个事务
	txHash := common.HexToHash(transactionInBlock)
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isPending)
	fmt.Println(tx.Hash().Hex())
}
