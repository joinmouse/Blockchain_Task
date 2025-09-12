package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/U3kArMqv4LilPctZa0_GvGpcO9OGaORP")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("xx")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	// 假设根据网络拥堵情况，将建议Gas价格提高20%
	suggestedGasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	adjustedGasPrice := big.NewInt(0).Mul(suggestedGasPrice, big.NewInt(120))
	adjustedGasPrice = big.NewInt(0).Div(adjustedGasPrice, big.NewInt(100))
	// 将 adjustedGasPrice 与额外的 Gas 价格（10000000000000）相加，得到最终的 Gas 价格
	gasPrice := big.NewInt(0).Add(adjustedGasPrice, big.NewInt(10000000000000))

	// 转账地址：在 ERC20 代币转账操作中，此地址为接收代币的目标账户地址。当调用代币合约的 transfer 方法时，代币将从发起转账的账户转移到该地址对应的账户。
	toAddress := common.HexToAddress("0xbf8164FA7982F6C878c0c779fB9ac126908B548E")
	// 代币合约地址：在以太坊网络中，每个 ERC20 代币都由一个智能合约管理。此地址指向该代币对应的智能合约，在进行代币转账等操作时，需要调用该合约的相关方法。
	tokenAddress := common.HexToAddress("0x63e466D034e421499c59EA689f2B9D539EA59198")
	// 定义 ERC20 代币合约的 transfer 方法签名，该方法用于转账
	transferFnSignature := []byte("transfer(address,uint256)")
	// 创建一个 Keccak256 哈希对象，用于计算方法签名的哈希值
	hash := sha3.NewLegacyKeccak256()
	// 将方法签名写入哈希对象进行哈希计算
	hash.Write(transferFnSignature)
	// 截取哈希结果的前 4 个字节作为方法 ID，用于识别调用的方法
	methodID := hash.Sum(nil)[:4]
	// 打印方法 ID 的十六进制编码
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb
	// 将转账目标地址填充到 32 字节长度，以满足 EVM 调用的数据格式要求
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	// 打印填充后的目标地址的十六进制编码
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d
	// 创建一个新的大整数对象，用于表示转账的代币数量
	amount := new(big.Int)
	// 将字符串形式的代币数量转换为大整数，这里表示 100 个代币
	amount.SetString("100000000000000000000", 10) // 100 tokens
	// 将代币数量填充到 32 字节长度，以满足 EVM 调用的数据格式要求
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	// 打印填充后的代币数量的十六进制编码
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
	// 定义一个字节切片，用于存储最终的调用数据
	var data []byte
	// 将方法 ID 添加到调用数据中
	data = append(data, methodID...)
	// 将填充后的目标地址添加到调用数据中
	data = append(data, paddedAddress...)
	// 将填充后的代币数量添加到调用数据中
	data = append(data, paddedAmount...)

	// 调用客户端的 EstimateGas 方法，用于估算执行该交易所需的 gas 量。
	// 该方法接收一个上下文对象和一个 CallMsg 结构体作为参数。
	// CallMsg 结构体包含交易的目标地址和调用数据。
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		// 设置交易的目标地址
		To: &toAddress,
		// 设置交易的调用数据，包含方法 ID、填充后的目标地址和填充后的代币数量
		Data: data,
	})
	// 检查估算 gas 量时是否发生错误
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasLimit:", gasLimit) // 23256
	gasLimit = uint64(100000)
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xeea6792293880fdfc3d7f2202deafe819fe5005879b5fe34e8a6804daa434357
}
