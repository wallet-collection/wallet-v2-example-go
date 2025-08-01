package model

import (
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type (
	WithdrawModel interface {
		Insert(model *Withdraw) error
		ListByPage(memberId int64, coinSymbol string, status int, businessId string, txid string, address string, offset, limit int) (int64, []*Withdraw, error)
		ListByPageExport(memberId int64, coinSymbol string, status int, businessId string, txid string, address string) ([]*Withdraw, error)
		List(memberId int64, coinSymbol string, startDate *time.Time, endDate *time.Time, offset, limit int) ([]*Withdraw, error)
		ListByInIds(ids []int64) ([]*Withdraw, error)
		FindByLastId(id int64) (*Withdraw, error)
		FindById(id int64) (*Withdraw, error)
		Update(id int64, data map[string]interface{}, status []int) error
	}

	defaultWithdrawModel struct {
		db *gorm.DB
	}

	Withdraw struct {
		Id           int64           // 自增ID
		MemberId     int64           // 用户ID
		NetworkName  string          // 网络名称
		CoinSymbol   string          // 币种符号
		Address      string          // 提现地址
		Amount       decimal.Decimal // 数量
		Fee          decimal.Decimal // 手续费
		BlockHigh    uint64          // 区块高度
		Txid         string          // 区块交易哈希
		Remark       string          // 备注
		Status       int             // 状态（0：审核中，1：审核通过，2：审核不通过，3：链上打包中，4：提币成功，5：提币失败，6：手动成功）
		CreateTime   time.Time       // 创建时间
		ModifiedTime time.Time       // 更新时间
	}
)

func (m *defaultWithdrawModel) Insert(model *Withdraw) error {
	err := m.db.Model(&Withdraw{}).Omit("business_id").Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultWithdrawModel) ListByPage(memberId int64, coinSymbol string, status int, businessId string, txid string, address string, offset, limit int) (int64, []*Withdraw, error) {
	var model []*Withdraw
	db := m.db.Model(&Withdraw{})
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

func (m *defaultWithdrawModel) ListByPageExport(memberId int64, coinSymbol string, status int, businessId string, txid string, address string) ([]*Withdraw, error) {
	var model []*Withdraw
	db := m.db.Model(&Withdraw{})
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

func (m *defaultWithdrawModel) List(memberId int64, coinSymbol string, startDate *time.Time, endDate *time.Time, offset, limit int) ([]*Withdraw, error) {
	var model []*Withdraw
	db := m.db.Model(&Withdraw{})
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

func (m *defaultWithdrawModel) ListByInIds(ids []int64) ([]*Withdraw, error) {
	var model []*Withdraw
	db := m.db.Model(&Withdraw{})
	db.Where("`id` IN(?)", ids)
	err := db.Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultWithdrawModel) FindByLastId(id int64) (*Withdraw, error) {
	var model *Withdraw
	db := m.db
	if id > 0 {
		db = db.Where("`id` > ?", id)
	}
	err := db.Order("`id` DESC").Limit(1).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultWithdrawModel) FindById(id int64) (*Withdraw, error) {
	var model *Withdraw
	err := m.db.Where("`id` = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultWithdrawModel) Update(id int64, data map[string]interface{}, status []int) error {
	if data == nil {
		return errors.New("data not")
	}
	err := m.db.Model(&Withdraw{}).Where("id = ?", id).Where("`status` IN (?)", status).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func NewWithdrawModel(db *gorm.DB) WithdrawModel {
	return &defaultWithdrawModel{
		db: db,
	}
}
