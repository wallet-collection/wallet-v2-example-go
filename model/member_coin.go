package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type (
	MemberCoinModel interface {
		Insert(model *MemberCoin) error
		InsertAll(model []*MemberCoin) error
		UpdateAll(model []*MemberCoin) error
		MinusFrozenAll(model []*MemberCoin) error
		FindByMemberIdAndCoinSymbol(memberId int64, coinSymbol string) (*MemberCoin, error)
		ListAllByMemberId(memberId int64, offset, limit int) (int64, []*MemberCoin, error)
		ListAllByBalance(coinSymbol string, balance decimal.Decimal) ([]*MemberCoin, error)
		UpdateVirtualBalance(memberId int64, coinSymbol string, balance decimal.Decimal) error
		UpdateBalanceAndFreeze(memberId int64, coinSymbol string, balance, freeze decimal.Decimal) error
		SumBalance(coinSymbol string) (*MemberCoinSumBalance, error)
	}

	defaultMemberCoinModel struct {
		db *gorm.DB
	}

	MemberCoin struct {
		Id             int64           //
		MemberId       int64           // 用户ID
		CoinSymbol     string          // 币种符号
		Balance        decimal.Decimal `sql:"type:decimal(48,18)"` // 可用余额
		FrozenBalance  decimal.Decimal `sql:"type:decimal(48,18)"` // 冻结余额
		VirtualBalance decimal.Decimal `sql:"type:decimal(48,18)"` // 虚拟余额
		CreateTime     time.Time       // 创建时间
		ModifiedTime   time.Time       // 更新时间
	}

	MemberCoinSumBalance struct {
		Balance decimal.Decimal // 数量
	}
)

func (m *defaultMemberCoinModel) Insert(model *MemberCoin) error {
	err := m.db.Model(&MemberCoin{}).Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultMemberCoinModel) InsertAll(model []*MemberCoin) error {
	err := m.db.Model(&MemberCoin{}).Omit("id").Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultMemberCoinModel) UpdateAll(model []*MemberCoin) error {

	var values []string

	now := time.Now().Format("2006-01-02 15:04:05")

	for _, coin := range model {
		var valuesItem []string
		valuesItem = append(valuesItem, strconv.FormatInt(coin.MemberId, 10))
		valuesItem = append(valuesItem, "'"+coin.CoinSymbol+"'")
		valuesItem = append(valuesItem, coin.Balance.String())
		valuesItem = append(valuesItem, coin.FrozenBalance.String())
		valuesItem = append(valuesItem, coin.VirtualBalance.String())
		valuesItem = append(valuesItem, "'"+now+"'")
		valuesItem = append(valuesItem, "'"+now+"'")
		values = append(values, "("+strings.Join(valuesItem, ",")+")")
	}

	sql := "INSERT INTO `member_coin`" +
		" (`member_id`,`coin_symbol`,`balance`,`frozen_balance`,`virtual_balance`,`create_time`,`modified_time`)" +
		" VALUES " +
		strings.Join(values, ",") +
		" ON DUPLICATE KEY UPDATE `balance`=`balance` + VALUES(`balance`),`virtual_balance`=`virtual_balance` + VALUES(`virtual_balance`),modified_time=?;"

	err := m.db.Exec(sql, now).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *defaultMemberCoinModel) MinusFrozenAll(model []*MemberCoin) error {

	var values []string

	now := time.Now().Format("2006-01-02 15:04:05")

	for _, coin := range model {
		var valuesItem []string
		valuesItem = append(valuesItem, strconv.FormatInt(coin.MemberId, 10))
		valuesItem = append(valuesItem, "'"+coin.CoinSymbol+"'")
		valuesItem = append(valuesItem, coin.Balance.String())
		valuesItem = append(valuesItem, coin.FrozenBalance.String())
		valuesItem = append(valuesItem, coin.VirtualBalance.String())
		valuesItem = append(valuesItem, "'"+now+"'")
		valuesItem = append(valuesItem, "'"+now+"'")
		values = append(values, "("+strings.Join(valuesItem, ",")+")")
	}

	sql := "INSERT INTO `member_coin`" +
		" (`member_id`,`coin_symbol`,`balance`,`frozen_balance`,`virtual_balance`,`create_time`,`modified_time`)" +
		" VALUES " +
		strings.Join(values, ",") +
		" ON DUPLICATE KEY UPDATE `balance`=`balance` + VALUES(`balance`),`frozen_balance`=`frozen_balance` - VALUES(`frozen_balance`),modified_time=?;"

	err := m.db.Exec(sql, now).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *defaultMemberCoinModel) FindByMemberIdAndCoinSymbol(memberId int64, coinSymbol string) (*MemberCoin, error) {
	var model *MemberCoin
	db := m.db.Model(&MemberCoin{})
	err := db.Where("`member_id` = ?", memberId).Where("coin_symbol = ?", coinSymbol).Order("`id` DESC").First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultMemberCoinModel) ListAllByMemberId(memberId int64, offset, limit int) (int64, []*MemberCoin, error) {
	var model []*MemberCoin
	db := m.db.Model(&MemberCoin{})
	if memberId > 0 {
		db.Where("`member_id` = ?", memberId)
	}
	var total int64

	if limit > 0 {

		db.Count(&total)

		db.Offset(offset).Limit(limit)
	} else {
		total = -1
	}

	err := db.Order("`id` DESC").Find(&model).Error
	if err != nil {
		return total, nil, err
	}
	return total, model, nil
}

func (m *defaultMemberCoinModel) ListAllByBalance(coinSymbol string, balance decimal.Decimal) ([]*MemberCoin, error) {
	var model []*MemberCoin
	db := m.db.Model(&MemberCoin{})
	db.Where("`coin_symbol` = ?", coinSymbol)
	if balance.Cmp(decimal.NewFromInt(0)) > 0 {
		db.Where("`balance` >= ?", balance)
	}
	err := db.Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultMemberCoinModel) UpdateVirtualBalance(memberId int64, coinSymbol string, balance decimal.Decimal) error {

	db := m.db.Model(&MemberCoin{}).Where("`member_id` = ?", memberId).Where("`coin_symbol` = ?", coinSymbol)

	data := make(map[string]interface{})
	decimal0 := decimal.NewFromInt(0)
	if balance.Cmp(decimal0) == 0 {
		return nil
	}
	if balance.Cmp(decimal0) >= 0 {
		data["virtual_balance"] = gorm.Expr("virtual_balance + " + balance.Abs().String())
	} else {
		data["virtual_balance"] = gorm.Expr("virtual_balance - " + balance.Abs().String())
		db.Where("virtual_balance >= " + balance.Abs().String())
	}
	return db.Updates(data).Error
}

func (m *defaultMemberCoinModel) UpdateBalanceAndFreeze(memberId int64, coinSymbol string, balance, freeze decimal.Decimal) error {

	db := m.db.Model(&MemberCoin{}).Where("`member_id` = ?", memberId).Where("`coin_symbol` = ?", coinSymbol)

	data := make(map[string]interface{})

	decimal0 := decimal.NewFromInt(0)
	if balance.Cmp(decimal0) > 0 {
		data["balance"] = gorm.Expr("balance + " + balance.Abs().String())
	} else if balance.Cmp(decimal0) < 0 {
		data["balance"] = gorm.Expr("balance - " + balance.Abs().String())
		db.Where("balance >= " + balance.Abs().String())
	}
	if freeze.Cmp(decimal0) > 0 {
		data["frozen_balance"] = gorm.Expr("frozen_balance + " + freeze.Abs().String())
	} else if freeze.Cmp(decimal0) < 0 {
		data["frozen_balance"] = gorm.Expr("frozen_balance - " + freeze.Abs().String())
		//db.Where("frozen_balance >= ?", freeze.Abs())
	}

	return db.Updates(data).Error
}

func (m *defaultMemberCoinModel) SumBalance(coinSymbol string) (*MemberCoinSumBalance, error) {

	db := m.db.Model(&MemberCoin{})

	db.Where("`coin_symbol` = ?", coinSymbol)

	var memberCoinSumBalance *MemberCoinSumBalance
	err := db.Select("SUM(balance) as balance").Scan(&memberCoinSumBalance).Error

	if err != nil {
		return nil, err
	}

	return memberCoinSumBalance, nil
}

func NewMemberCoinModel(db *gorm.DB) MemberCoinModel {
	return &defaultMemberCoinModel{
		db: db,
	}
}
