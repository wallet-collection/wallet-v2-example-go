package biz

import (
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
	"wallet/model"
)

type BusinessType struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

const (
	BillBusinessTypeAdminUpdateAdd  = "admin_update_add"
	BillBusinessTypeAdminUpdateDec  = "admin_update_dec"
	BillBusinessTypeRecharge        = "recharge"
	BillBusinessTypeWithdraw        = "withdraw"
	BillBusinessTypeTransfer        = "transfer"
	BillBusinessTypeSpotOrderCreate = "spot_order"
	BillBusinessTypeSpotOrderCancel = "spot_cancel"
	BillBusinessTypeSpotOrderMatch  = "spot_match"

	BillModeZero     = 0
	BillModeIncrease = 1
	BillModeReduce   = 2
)

var BillBusinessTypeList = []BusinessType{
	{Name: "空投", Value: BillBusinessTypeAdminUpdateAdd},
	{Name: "回收", Value: BillBusinessTypeAdminUpdateDec},
	{Name: "充值", Value: BillBusinessTypeRecharge},
	{Name: "提现", Value: BillBusinessTypeWithdraw},
	{Name: "划转", Value: BillBusinessTypeTransfer},
	{Name: "下单", Value: BillBusinessTypeSpotOrderCreate},
	{Name: "撤单", Value: BillBusinessTypeSpotOrderCancel},
	{Name: "撮合", Value: BillBusinessTypeSpotOrderMatch},
}

type MemberBillBiz struct {
	MemberId     int64
	CoinSymbol   string
	FromAccount  string // 来源账户
	ToAccount    string // 接收账户
	Mode         int    // 模式，0：划转，1：转入，2：转出
	IsVirtual    bool
	BusinessType string
	BusinessId   string
	Balance      decimal.Decimal
	Freeze       decimal.Decimal
	Remark       string
}

func (m *MemberBillBiz) Create(db *gorm.DB) error {

	// 查询余额
	memberCoinBiz := &MemberCoinBiz{}

	decimal0 := decimal.NewFromInt(0)

	if m.Balance.Cmp(decimal0) < 0 {
		balance, err := memberCoinBiz.Balance(m.MemberId, m.CoinSymbol, db)
		if err != nil {
			return err
		}
		if m.IsVirtual {
			if balance.VirtualBalance.Cmp(m.Balance) < 0 {
				return errors.New("insufficient balance")
			}
		} else {
			if balance.Balance.Cmp(m.Balance) < 0 {
				return errors.New("insufficient balance")
			}
		}
	}

	memberCoinModel := model.NewMemberCoinModel(db)

	if m.IsVirtual {
		err := memberCoinModel.UpdateVirtualBalance(m.MemberId, m.CoinSymbol, m.Balance)
		if err != nil {
			return err
		}
	} else {
		err := memberCoinModel.UpdateBalanceAndFreeze(m.MemberId, m.CoinSymbol, m.Balance, m.Freeze)
		if err != nil {
			return err
		}
	}

	if m.Balance.Abs().Cmp(decimal0) <= 0 || m.Mode < 0 {
		return nil
	}
	memberBillModel := model.NewMemberBillModel(db)
	data := make([]*model.MemberBill, 0)
	data = append(data, &model.MemberBill{
		MemberId:     m.MemberId,
		CoinSymbol:   m.CoinSymbol,
		FromAccount:  m.FromAccount,
		ToAccount:    m.ToAccount,
		Mode:         m.Mode,
		BusinessType: m.BusinessType,
		BusinessId:   m.BusinessId,
		Amount:       m.Balance.Abs(),
		Remark:       m.Remark,
		CreateTime:   time.Now(),
		ModifiedTime: time.Now(),
	})

	err := memberBillModel.InsertAll(data)
	if err != nil {
		return err
	}

	return nil
}
