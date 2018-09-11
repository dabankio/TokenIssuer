/*
* Copyright (c) 2018 Dabank
* Use of this work is governed by a MIT License.
* You may find a license copy under project root.
 */

package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"

	"math"

	"github.com/dabankio/TokenIssuer"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	rpcURL      string
	name        string
	symbol      string
	totalSupply uint64
	decimals    uint64
	owner       string
	prikeyStr   string

	// Version build params
	Version string
	// BuildDate build params
	BuildDate string
)

func init() {
	flag.StringVar(&rpcURL, "rpc", "https://mainnet.infura.io", "Ethereum RPC endpoint")
	flag.StringVar(&name, "name", "My Lovely Token", "token Name")
	flag.StringVar(&symbol, "symbol", "MLT", "token symbol")
	flag.Uint64Var(&totalSupply, "total", 42, "initial supply for token")
	flag.StringVar(&owner, "owner", "", "owner of token, default to message sender")
	flag.Uint64Var(&decimals, "decimals", 6, "decimals of token")
	flag.StringVar(&prikeyStr, "prikey", "", "prikey of message sender")
}

func main() {
	flag.Parse()

	fmt.Printf("issueToken(%s) by Dabank Authors, built on %s\n", Version, BuildDate)

	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	if decimals > math.MaxUint8 {
		err = fmt.Errorf("decimals should be less than 255, got %v", decimals)
		return
	}
	if totalSupply > math.MaxInt64 {
		err = fmt.Errorf("total should be less than 9223372036854775807")
		return
	}

	bigTotalSupply := new(big.Int).
		Mul(big.NewInt(int64(totalSupply)),
			new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))

	prikeyStr = strings.TrimLeft(prikeyStr, "0x")
	prikeyStr = strings.TrimLeft(prikeyStr, "0X")

	prikeyBytes, err := hex.DecodeString(prikeyStr)
	if err != nil {
		err = fmt.Errorf("can not decode prikey from hex: %v", err)
		return
	}

	prikey, err := crypto.ToECDSA(prikeyBytes)
	if err != nil {
		err = fmt.Errorf("can not convert prikey to ECDSA: %v", err)
		return
	}

	var ownerAddr common.Address
	if len(owner) == 0 {
		ownerAddr = crypto.PubkeyToAddress(prikey.PublicKey)
	} else {
		ownerAddr = common.HexToAddress(owner)
		if strings.ToLower(ownerAddr.Hex()) != strings.ToLower(owner) {
			err = fmt.Errorf("invalid owner input")
			return
		}
	}

	coon, err := ethclient.Dial(rpcURL)
	if err != nil {
		err = fmt.Errorf("can not dial Ethereum RPC: %v", err)
		return
	}

	opts := bind.NewKeyedTransactor(prikey)

	tokenAddr, tx, _, err := TokenIssuer.DeployToken(opts, coon, name, symbol, ownerAddr, bigTotalSupply, uint8(decimals))
	if err != nil {
		err = fmt.Errorf("deploy failed: %v", err)
		return
	}

	fmt.Printf(`Congradulations! Your %s(%s) will be addressed at
%s
after tx %s is mined. :)
`, name, symbol,
		tokenAddr.Hex(),
		tx.Hash().Hex())
}
