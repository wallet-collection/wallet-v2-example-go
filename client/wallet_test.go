package client

import (
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestWalletClient_CreateWithdraw(t *testing.T) {

	withdrawClient := NewWalletClient("123123", "123123", "http://127.0.1.1:10002")

	networkName := "BEP20"
	coinSymbol := "USDT"
	address := "0xaa715F0c5083c5FEf99b097C5C75a34519e20a81"
	amount := decimal.NewFromInt(1)
	businessId := "1"
	privateKey := "xxx"
	dataRes, err := withdrawClient.CreateWithdraw(networkName, coinSymbol, address, amount, businessId, privateKey, "")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(dataRes)
}

func TestName222(t *testing.T) {

	amount := decimal.NewFromInt(83)
	// 82.999999999999990000
	fmt.Println(amount.Abs())

	/*
		SELECT * FROM `withdraw` WHERE member_id = 624 ORDER BY id DESC LIMIT 1;
		SELECT * FROM `member_coin` WHERE `member_id` = 624 ORDER BY id DESC;
	*/

}
