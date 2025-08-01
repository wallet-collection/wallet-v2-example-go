package model

import (
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type (
	CoinConfModel interface {
		Insert(model *CoinConf) error
		ListByPage(networkName string, coinSymbol string, status, offset, limit int) (int64, []*CoinConf, error)
		ListAll(coinSymbol string) ([]*CoinConf, error)
		FindById(id int64) (*CoinConf, error)
		FindByNetworkNameAndCoinSymbol(networkName string, coinSymbol string) (*CoinConf, error)
		Update(id int64, data map[string]interface{}) error
		Delete(id int64) error
	}

	defaultCoinConfModel struct {
		db *gorm.DB
	}

	CoinConf struct {
		Id                 int64           // 自增ID
		CoinSymbol         string          // 币种符号
		NetworkName        string          // 网络名称
		Decimals           int             // 币种精度
		IsWithdraw         int             // 是否可提现（0：否，1：手动，2：自动）
		WithdrawAuto       decimal.Decimal // 提现自动转的阈值
		WithdrawRate       decimal.Decimal // 提现费率
		MinWithdrawFee     decimal.Decimal // 最低提现费用
		MinWithdraw        decimal.Decimal // 最低提现
		MaxWithdraw        decimal.Decimal // 最大提现
		WithdrawPrivateKey string          // 提现私钥
		MinRecharge        decimal.Decimal // 最低充值
		IsRecharge         int             // 是否可充值（0：手动，1：自动）
		RechargeConfirm    int             // 充值确认数
		WithdrawConfirm    int             // 提现确认数
		Sort               int             // 排序（升序）
		Status             int             // 状态（0：禁用，1：正常）
		CreateTime         time.Time       // 创建时间
		ModifiedTime       time.Time       // 更新时间
	}
)

func (m *defaultCoinConfModel) Insert(model *CoinConf) error {
	err := m.db.Model(&CoinConf{}).Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultCoinConfModel) ListByPage(coinSymbol string, networkName string, status, offset, limit int) (int64, []*CoinConf, error) {
	var model []*CoinConf
	db := m.db.Model(&CoinConf{})
	if len(coinSymbol) > 0 {
		db.Where("`coin_symbol` = ?", coinSymbol)
	}
	if len(networkName) > 0 {
		db.Where("`network_name` = ?", networkName)
	}
	if status >= 0 {
		db.Where("`status` = ?", status)
	}
	var total int64
	db.Count(&total)
	err := db.Offset(offset).Limit(limit).Order("`sort` DESC").Find(&model).Error
	if err != nil {
		return 0, nil, err
	}
	return total, model, nil
}

func (m *defaultCoinConfModel) ListAll(coinSymbol string) ([]*CoinConf, error) {
	var model []*CoinConf
	db := m.db.Model(&CoinConf{})
	if len(coinSymbol) > 0 {
		db.Where("coin_symbol = ?", coinSymbol)
	}
	err := db.Order("`sort` DESC").Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultCoinConfModel) FindById(id int64) (*CoinConf, error) {
	var model *CoinConf
	err := m.db.Where("id = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultCoinConfModel) FindByNetworkNameAndCoinSymbol(coinSymbol string, networkName string) (*CoinConf, error) {
	var model *CoinConf
	err := m.db.Where("`coin_symbol` = ?", coinSymbol).Where("network_name = ?", networkName).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultCoinConfModel) Update(id int64, data map[string]interface{}) error {
	if data == nil {
		return errors.New("data not")
	}
	err := m.db.Model(&CoinConf{}).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultCoinConfModel) Delete(id int64) error {
	err := m.db.Where("id = ?", id).Delete(&CoinConf{}).Error
	if err != nil {
		return err
	}
	return nil
}

func NewCoinConfModel(db *gorm.DB) CoinConfModel {
	return &defaultCoinConfModel{
		db: db,
	}
}
