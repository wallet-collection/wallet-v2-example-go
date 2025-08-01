package util

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"math/big"
	"testing"
)

func TestPrivateKeySign(t *testing.T) {

	privateKeyStr := "03f6ebc24e6f5bb06891d2d0068d50c8cb6dc7bb3dd96ea38a6a05c9cae64168"

	id := int64(1)
	amount := decimal.NewFromInt(100)
	address := "0xaa715F0c5083c5FEf99b097C5C75a34519e30a83"

	hash := crypto.Keccak256Hash(
		common.LeftPadBytes(big.NewInt(id).Bytes(), 32),
		common.LeftPadBytes(amount.BigInt().Bytes(), 32),
		common.HexToAddress(address).Bytes(),
	)

	signStr, err := PrivateKeySign(privateKeyStr, hash)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(hash.Hex(), signStr)

}

func TestVerifySignature(t *testing.T) {

	initdata := "登录网站"
	sign := "0x5321f24a057500605f1d894c2be7cb7f196ba2444e8f6815af261efbcb9d272f70d327f146553c3d51cf1816823dba6254d5500a69b4197e9f4839e0971cf89d1b"
	publicKey := "0x0bDCC0C6eAc88439fb57b90977714b7430c3c623"

	publicKey2, err := VerifyMessage(initdata, sign)
	fmt.Println(publicKey == publicKey2, err)
}
