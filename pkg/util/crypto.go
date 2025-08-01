package util

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strconv"
)

// PrivateKeySign 私钥签名
func PrivateKeySign(privateKeyStr string, hash common.Hash) (string, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return "", err
	}
	msgHash := crypto.Keccak256Hash(
		[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%v", len(hash))),
		hash.Bytes(),
	)

	signature, _ := crypto.Sign(msgHash.Bytes(), privateKey)
	signature[64] += 27

	return "0x" + hex.EncodeToString(signature), nil
}

// VerifyMessage 验证签名，返回签名地址
func VerifyMessage(message string, sign string) (string, error) {
	// Hash the unsigned message using EIP-191
	hashedMessage := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + message)
	hash := crypto.Keccak256Hash(hashedMessage)

	// Get the bytes of the signed message
	decodedMessage := hexutil.MustDecode(sign)

	// Handles cases where EIP-115 is not implemented (most wallets don't implement it)
	if decodedMessage[64] == 27 || decodedMessage[64] == 28 {
		decodedMessage[64] -= 27
	}

	// Recover a public key from the signed message
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), decodedMessage)
	if sigPublicKeyECDSA == nil {
		err = errors.New("could not get a public get from the message signature")
	}
	if err != nil {
		return "", err
	}

	return crypto.PubkeyToAddress(*sigPublicKeyECDSA).String(), nil
}
