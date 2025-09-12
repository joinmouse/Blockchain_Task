package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	// 调用 crypto 包的 GenerateKey 函数生成一个新的 ECDSA 私钥。
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// 将生成的 ECDSA 私钥转换为字节切片
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// 对私钥字节切片进行十六进制编码，并去掉前缀 '0x' 后打印输出
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])
	// 从私钥中获取对应的公钥
	publicKey := privateKey.Public()
	// 将公钥断言为 *ecdsa.PublicKey 类型
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 将 ECDSA 公钥转换为字节切片
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// 打印去掉前缀 '0x04' 后的公钥十六进制编码
	fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'
	// 将公钥转换为以太坊地址，并获取其十六进制表示
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	// 打印以太坊地址的十六进制表示
	fmt.Println(address)
	// 创建一个新的 Keccak-256 哈希对象
	hash := sha3.NewLegacyKeccak256()
	// 对去掉第一个字节后的公钥字节切片进行哈希计算
	hash.Write(publicKeyBytes[1:])
	// 打印完整的哈希结果的十六进制编码
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	// 打印截取前 12 位后剩余的 20 位哈希结果的十六进制编码，原哈希结果长度为 32 位
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 原长32位，截去12位，保留后20位
}
