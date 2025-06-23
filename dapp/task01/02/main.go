package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/learn/init_order/count"
)

const (
	contractAddr = "0x7Ea32bF4d58a2BB293fC5e4784829a0f1EDa49Ef"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/U3kArMqv4LilPctZa0_GvGpcO9OGaORP")
	if err != nil {
		log.Fatal(err)
	}
	storeContract, err := count.NewCount(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("privateKey")
	if err != nil {
		log.Fatal(err)
	}

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := storeContract.Increment(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := storeContract.GetCount(callOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("valueInContract in contract:", valueInContract)

	value := big.NewInt(6)

	fmt.Println("is value saving in contract equals to origin value:", valueInContract.Cmp(value) == 0)
}
