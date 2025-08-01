package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type (
	CoinNetworkModel interface {
		Insert(model *CoinNetwork) error
		ListByPage(status, offset, limit int) (int64, []*CoinNetwork, error)
		ListAll() ([]*CoinNetwork, error)
		FindById(id int64) (*CoinNetwork, error)
		FindByName(name string) (*CoinNetwork, error)
		Update(id int64, data map[string]interface{}) error
		Delete(id int64) error
	}

	defaultCoinNetworkModel struct {
		db *gorm.DB
	}

	CoinNetwork struct {
		Id           int64     // 自增ID
		Name         string    // 网络名称
		Sort         int       // 排序（升序）
		Status       int       // 状态（0：禁用，1：正常）
		CreateTime   time.Time // 创建时间
		ModifiedTime time.Time // 更新时间
	}
)

func (m *defaultCoinNetworkModel) Insert(model *CoinNetwork) error {
	err := m.db.Model(&CoinNetwork{}).Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultCoinNetworkModel) ListByPage(status, offset, limit int) (int64, []*CoinNetwork, error) {
	var model []*CoinNetwork
	db := m.db.Model(&CoinNetwork{})
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

func (m *defaultCoinNetworkModel) ListAll() ([]*CoinNetwork, error) {
	var model []*CoinNetwork
	db := m.db.Model(&CoinNetwork{})
	err := db.Order("`sort` DESC").Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultCoinNetworkModel) FindById(id int64) (*CoinNetwork, error) {
	var model *CoinNetwork
	err := m.db.Where("id = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultCoinNetworkModel) FindByName(name string) (*CoinNetwork, error) {
	var model *CoinNetwork
	err := m.db.Where("`name` = ?", name).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultCoinNetworkModel) Update(id int64, data map[string]interface{}) error {
	if data == nil {
		return errors.New("data not")
	}
	err := m.db.Model(&CoinNetwork{}).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultCoinNetworkModel) Delete(id int64) error {
	err := m.db.Where("id = ?", id).Delete(&CoinNetwork{}).Error
	if err != nil {
		return err
	}
	return nil
}

func NewCoinNetworkModel(db *gorm.DB) CoinNetworkModel {
	return &defaultCoinNetworkModel{
		db: db,
	}
}
