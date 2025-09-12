package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 1、尝试连接到Rinkeby测试网络的Infura节点，创建一个以太坊客户端实例
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/U3kArMqv4LilPctZa0_GvGpcO9OGaORP")
	// 检查连接过程中是否出现错误
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("连接成功:", client)

	// 2、将十六进制格式的私钥字符串转换为ECDSA私钥。这里使用的私钥是
	privateKey, err := crypto.HexToECDSA("privateKey")
	// 检查转换过程中是否出现错误，如果出现错误则终止程序并输出错误信息
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("私钥转换成功:", privateKey)

	// 获取私钥对应的公钥
	publicKey := privateKey.Public()
	// 若后续需要使用 publicKeyECDSA，可在此处添加使用代码
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fmt.Println("公钥转换成功:", publicKeyECDSA)

	// 将公钥转换为以太坊地址，该地址将作为交易的发送方地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 从以太坊客户端获取发送方地址的待处理交易的随机数（nonce），该随机数用于确保交易的唯一性
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nonce:", nonce)

	// 定义本次交易要发送的以太币数量，单位为wei，这里设置为1以太币
	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	// 定义本次交易的 gas 限制，单位为 gas 单位，这里设置为标准的 21000 单位
	gasLimit := uint64(21000) // in units
	// 从以太坊客户端获取建议的 gas 价格，该价格会根据当前网络状况动态调整
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("gasPrice:", gasPrice)

	// 将十六进制格式的地址字符串转换为以太坊地址，该地址将作为交易的接收方地址
	toAddress := common.HexToAddress("0xbf8164FA7982F6C878c0c779fB9ac126908B548E")
	log.Println("toAddress:", toAddress)
	// 定义交易附带的数据，这里初始化为空切片，表示不附带额外数据
	var data []byte
	// 使用之前获取的随机数、接收方地址、转账金额、gas限制、gas价格和交易数据创建一个新的以太坊交易
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// 从以太坊客户端获取当前网络的链ID，链ID用于交易签名时确保交易在正确的网络上执行
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("NetworkID:", err)
	}
	fmt.Println("chainID:", chainID)

	// 使用 EIP155 签名器对之前创建的以太坊交易进行签名，需要传入交易对象、EIP155 签名器实例和发送方的私钥
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal("sign:", err)
	}
	fmt.Println("signedTx:", signedTx)

	// 通过以太坊客户端将签名后的交易发送到以太坊网络
	err = client.SendTransaction(context.Background(), signedTx)
	fmt.Println("success", err)
	// 检查发送交易过程中是否出现错误，如果出现错误则终止程序并输出错误信息
	if err != nil {
		log.Fatal("send", err)
	}

	// 打印已发送交易的哈希值，方便后续查询交易状态
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // 0x80f69fec967636917ebde14985329fa5b88745e94dc6aa1b834e5af1217d3b10
}
