package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/U3kArMqv4LilPctZa0_GvGpcO9OGaORP")
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
	blockHash := common.HexToHash(block.Hash().Hex())

	// 调用 BlockReceipts 方法就可以得到指定区块中所有的收据列表
	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}
	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("通过区块哈希获取的收据列表", receiptByHash[0])
	fmt.Println("通过区块编号获取的收据列表", receiptsByNum[0])
	fmt.Println("通过区块哈希获取的收据列表和通过区块编号获取的收据列表是否相等", receiptByHash[0] == receiptsByNum[0]) // true

	// 遍历通过区块哈希获取的交易收据列表
	var receiptTxHash string
	for _, receipt := range receiptByHash {
		// 打印交易的执行状态，1 通常表示成功
		fmt.Println(receipt.Status)
		// 打印交易产生的日志列表，这里为空
		fmt.Println(receipt.Logs)
		// 打印交易的哈希值
		fmt.Println(receipt.TxHash.Hex())
		receiptTxHash = receipt.TxHash.Hex()
		// 打印交易在区块中的索引位置
		fmt.Println(receipt.TransactionIndex)
		// 打印交易创建的合约地址，若未创建合约则为 0x0000...
		fmt.Println(receipt.ContractAddress.Hex())
		// 仅处理第一个收据，跳出循环
		break
	}

	// 将字符串类型的交易哈希转换为 common.Hash 类型
	txHash := common.HexToHash(receiptTxHash)
	// 调用 TransactionReceipt 方法，根据交易哈希获取该交易的收据信息
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	// 检查获取交易收据信息是否出错，若出错则输出错误信息并终止程序
	if err != nil {
		log.Fatal(err)
	}
	// 打印交易的执行状态，1 通常表示成功
	fmt.Println(receipt.Status)
	// 打印交易产生的日志列表，这里为空
	fmt.Println(receipt.Logs)
	// 打印交易的哈希值
	fmt.Println(receipt.TxHash.Hex())
	// 打印交易在区块中的索引位置
	fmt.Println(receipt.TransactionIndex)
	// 打印交易创建的合约地址，若未创建合约则为 0x0000...
	fmt.Println(receipt.ContractAddress.Hex())
}
