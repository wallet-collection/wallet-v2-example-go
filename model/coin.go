package model

import (
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type (
	CoinModel interface {
		Insert(model *Coin) error
		ListByPage(status, offset, limit int) (int64, []*Coin, error)
		ListAll() ([]*Coin, error)
		FindById(id int64) (*Coin, error)
		FindBySymbol(symbol string) (*Coin, error)
		Update(id int64, data map[string]interface{}) error
		Delete(id int64) error
	}

	defaultCoinModel struct {
		db *gorm.DB
	}

	Coin struct {
		Id             int64           // 自增ID
		Name           string          // 币种名称
		Symbol         string          // 币种单位
		Icon           string          // 图标
		UsdtPrice      decimal.Decimal // USDT价格
		IsAutoPrice    int             // 是否自动获取价格（0：否，1：是）
		Precision      int             // 精度
		IsTransfer     int             // 是否可划转（0：否，1：手动，2：自动）
		TransferRate   decimal.Decimal // 划转费率
		MinTransferFee decimal.Decimal // 最低划转费用
		MinTransfer    decimal.Decimal // 最低划转
		MaxTransfer    decimal.Decimal // 最大划转
		Sort           int             // 排序（升序）
		Status         int             // 状态（0：禁用，1：正常）
		CreateTime     time.Time       // 创建时间
		ModifiedTime   time.Time       // 更新时间
	}
)

func (m *defaultCoinModel) Insert(model *Coin) error {
	err := m.db.Model(&Coin{}).Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultCoinModel) ListByPage(status, offset, limit int) (int64, []*Coin, error) {
	var model []*Coin
	db := m.db.Model(&Coin{})
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

func (m *defaultCoinModel) ListAll() ([]*Coin, error) {
	var model []*Coin
	db := m.db.Model(&Coin{})
	err := db.Order("`sort` DESC").Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultCoinModel) FindById(id int64) (*Coin, error) {
	var model *Coin
	err := m.db.Where("id = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultCoinModel) FindBySymbol(symbol string) (*Coin, error) {
	var model *Coin
	err := m.db.Where("`symbol` = ?", symbol).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultCoinModel) Update(id int64, data map[string]interface{}) error {
	if data == nil {
		return errors.New("data not")
	}
	err := m.db.Model(&Coin{}).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultCoinModel) Delete(id int64) error {
	err := m.db.Where("id = ?", id).Delete(&Coin{}).Error
	if err != nil {
		return err
	}
	return nil
}

func NewCoinModel(db *gorm.DB) CoinModel {
	return &defaultCoinModel{
		db: db,
	}
}
