package api

import (
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"testing"
	"time"
	"wallet-example/api/req"
	"wallet-example/biz"
	"wallet-example/config"
	"wallet-example/model"
)

func TestWithdrawCreate(t *testing.T) {

	conf, err := config.NewConfig("../config/config-example.yml")

	if err != nil {
		panic("Failed to load configuration")
	}

	db, err := model.NewDB(conf.Mysql)
	if err != nil {
		fmt.Println(err)
		return
	}

	memberId := int64(4)

	// 设置随机数的范围
	minValue := 0.1
	maxValue := 1.9

	sum := decimal.NewFromInt(0)
	for i := 0; i < 20; i++ {
		// 生成随机浮点数
		randomFloat := decimal.NewFromFloat(minValue + rand.Float64()*(maxValue-minValue)).RoundFloor(4)
		//randomFloat = decimal.NewFromFloat(0.5117)
		fmt.Println("数量：", randomFloat)
		q := &req.WithdrawCreateReq{
			NetworkName: "BEP20",
			CoinSymbol:  "USDT",
			Address:     "111",
			Amount:      randomFloat,
			EmailCode:   "",
			TelCode:     "",
			GoogleCode:  "",
		}

		sum = sum.Add(randomFloat)

		tt(memberId, q, db)

		memberCoin, err := model.NewMemberCoinModel(db).FindByMemberIdAndCoinSymbol(memberId, q.CoinSymbol)
		if err != nil {
			return
		}

		fmt.Println("余额：", memberCoin.Balance, "冻结：", memberCoin.FrozenBalance)

	}

	fmt.Println("总量：", sum)

}

/*
数量： 1.0315
数量： 0.3299
数量： 1.0936
数量： 0.156
数量： 0.102
数量： 1.8945
数量： 0.7121
数量： 1.557
数量： 0.9039
数量： 1.5309
数量： 0.811
数量： 1.2492
数量： 0.7207
数量： 0.4853
数量： 0.5943
数量： 0.1873
数量： 1.7459
数量： 1.7183
数量： 0.6405
数量： 0.5936
总量： 18.0575
*/

func tt(memberId int64, q *req.WithdrawCreateReq, db *gorm.DB) {

	withdrawAmount := q.Amount

	fee := decimal.NewFromFloat(0.1)

	withdraw := &model.Withdraw{
		MemberId:     memberId,
		NetworkName:  q.NetworkName,
		CoinSymbol:   q.CoinSymbol,
		Address:      q.Address,
		Amount:       withdrawAmount,
		Fee:          fee,
		Remark:       "",
		Status:       0,
		CreateTime:   time.Now(),
		ModifiedTime: time.Now(),
	}

	err := db.Transaction(func(tx *gorm.DB) error {

		withdrawModel := model.NewWithdrawModel(tx)

		err := withdrawModel.Insert(withdraw)
		if err != nil {
			return err
		}

		businessId := strconv.FormatInt(withdraw.Id, 10)

		memberBillBiz := biz.MemberBillBiz{
			MemberId:     memberId,
			CoinSymbol:   q.CoinSymbol,
			Mode:         biz.BillModeReduce,
			BusinessType: biz.BillBusinessTypeWithdraw,
			BusinessId:   businessId,
			Balance:      q.Amount.Neg(),
			Freeze:       q.Amount,
			Remark:       "",
		}
		err = memberBillBiz.Create(tx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}
}
