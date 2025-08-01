package model

import (
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type (
	RechargeModel interface {
		Insert(model *Recharge) error
		ListByPage(memberId int64, coinSymbol string, status int, businessId string, txid string, address string, offset, limit int) (int64, []*Recharge, error)
		ListByPageExport(memberId int64, coinSymbol string, status int, businessId string, txid string, address string) ([]*Recharge, error)
		List(memberId int64, coinSymbol string, startDate *time.Time, endDate *time.Time, offset, limit int) ([]*Recharge, error)
		FindByBusinessId(businessId string) (*Recharge, error)
		FindById(id int64) (*Recharge, error)
		Update(businessId string, data map[string]interface{}) error
	}

	defaultRechargeModel struct {
		db *gorm.DB
	}

	Recharge struct {
		Id           int64           // 自增ID
		BusinessId   string          // 业务ID
		MemberId     int64           // 用户ID
		NetworkName  string          // 网络名称
		CoinSymbol   string          // 币种符号
		Address      string          // 充值地址
		Amount       decimal.Decimal // 数量
		MaxBlockHigh uint64          // 最大区块高度
		BlockHigh    uint64          // 区块高度
		Txid         string          // 区块交易哈希
		Remark       string          // 备注
		Status       int             // 状态（0：区块确认中，1：充值到账，2：区块确认失败）
		CreateTime   time.Time       // 创建时间
		ModifiedTime time.Time       // 更新时间
	}
)

func (m *defaultRechargeModel) Insert(model *Recharge) error {
	err := m.db.Model(&Recharge{}).Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultRechargeModel) ListByPage(memberId int64, coinSymbol string, status int, businessId string, txid string, address string, offset, limit int) (int64, []*Recharge, error) {
	var model []*Recharge
	db := m.db.Model(&Recharge{})
	if memberId > 0 {
		db.Where("`member_id` = ?", memberId)
	}
	if len(coinSymbol) > 0 {
		db.Where("`coin_symbol` = ?", coinSymbol)
	}
	if status >= 0 {
		db.Where("`status` = ?", status)
	}
	if len(businessId) > 0 {
		db.Where("`business_id` = ?", businessId)
	}
	if len(txid) > 0 {
		db.Where("`txid` = ?", txid)
	}
	if len(address) > 0 {
		db.Where("`address` = ?", address)
	}
	var total int64
	db.Count(&total)
	err := db.Offset(offset).Limit(limit).Order("`create_time` DESC").Find(&model).Error
	if err != nil {
		return 0, nil, err
	}
	return total, model, nil
}

func (m *defaultRechargeModel) ListByPageExport(memberId int64, coinSymbol string, status int, businessId string, txid string, address string) ([]*Recharge, error) {
	var model []*Recharge
	db := m.db.Model(&Recharge{})
	if memberId > 0 {
		db.Where("`member_id` = ?", memberId)
	}
	if len(coinSymbol) > 0 {
		db.Where("`coin_symbol` = ?", coinSymbol)
	}
	if status >= 0 {
		db.Where("`status` = ?", status)
	}
	if len(businessId) > 0 {
		db.Where("`business_id` = ?", businessId)
	}
	if len(txid) > 0 {
		db.Where("`txid` = ?", txid)
	}
	if len(address) > 0 {
		db.Where("`address` = ?", address)
	}

	err := db.Order("`create_time` DESC").Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultRechargeModel) List(memberId int64, coinSymbol string, startDate *time.Time, endDate *time.Time, offset, limit int) ([]*Recharge, error) {
	var model []*Recharge
	db := m.db.Model(&Recharge{})
	if memberId > 0 {
		db.Where("`member_id` = ?", memberId)
	}
	if len(coinSymbol) > 0 {
		db.Where("`coin_symbol` = ?", coinSymbol)
	}

	if startDate != nil && endDate != nil && !startDate.IsZero() && !endDate.IsZero() {
		start := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
		end := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 0, endDate.Location())
		db.Where("`create_time` between ? and ?", start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))
	}

	err := db.Offset(offset).Limit(limit).Order("`create_time` DESC").Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultRechargeModel) FindByBusinessId(businessId string) (*Recharge, error) {
	var model *Recharge
	err := m.db.Where("`business_id` = ?", businessId).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultRechargeModel) FindById(id int64) (*Recharge, error) {
	var model *Recharge
	err := m.db.Where("`id` = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultRechargeModel) Update(businessId string, data map[string]interface{}) error {
	if data == nil {
		return errors.New("data not")
	}
	db := m.db.Model(&Recharge{}).Where("business_id = ?", businessId)
	if _, ok := data["status"]; ok {
		db.Where("`status` = ?", 0)
	}
	err := db.Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func NewRechargeModel(db *gorm.DB) RechargeModel {
	return &defaultRechargeModel{
		db: db,
	}
}
