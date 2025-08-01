package util

import (
	"fmt"
	"math/big"
	"testing"
	"time"
)

var (
	expDiffPeriod = big.NewInt(100000)
	big1          = big.NewInt(1)
	big2          = big.NewInt(2)
	big9          = big.NewInt(9)
	big10         = big.NewInt(10)
	bigMinus99    = big.NewInt(-99)
)

func TestCreatePwd(t *testing.T) {

	s := CreatePwd("5HWauy9yoF2hMQXz")
	fmt.Println(s)

}

func TestName(t *testing.T) {

	a := big.NewInt(0)
	a.SetString("172008650762878078186869", 10)

	b := big.NewInt(1)
	b.Mul(b, big.NewInt(1e17))
	fmt.Println(b)
	fmt.Println(a.Div(a, big.NewInt(1e18)))

}

func TestDate(t *testing.T) {

	remindTime, err := time.Parse("2006-01-02 15:04:05", "2023-01-02 13:32:00")

	fmt.Println(remindTime, err)
}
