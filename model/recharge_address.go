package model

import (
	"gorm.io/gorm"
	"time"
)

// 当前表做了触发器，不能删除不能修改，只能增加，防止黑客入侵

type (
	RechargeAddressModel interface {
		Insert(model *RechargeAddress) error
		FindByMemberIdAndNetworkNameAndCoinSymbol(memberId int64, networkName string, coinSymbol string) (*RechargeAddress, error)
		FindByNetworkNameAndCoinSymbolAndAddress(networkName string, coinSymbol string, address string) (*RechargeAddress, error)
	}

	defaultRechargeAddressModel struct {
		db *gorm.DB
	}

	RechargeAddress struct {
		Id           int64     // 自增ID
		Key          string    // 随机key
		MemberId     int64     // 用户ID
		NetworkName  string    // 网络
		CoinSymbol   string    // 币种符号
		Address      string    // 地址
		CreateTime   time.Time // 创建时间
		ModifiedTime time.Time // 更新时间
	}
)

func (m *defaultRechargeAddressModel) Insert(model *RechargeAddress) error {
	model.Key = "suijikey"
	err := m.db.Model(&RechargeAddress{}).Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultRechargeAddressModel) FindByMemberIdAndNetworkNameAndCoinSymbol(memberId int64, networkName string, coinSymbol string) (*RechargeAddress, error) {
	var model *RechargeAddress
	err := m.db.Where("`member_id` = ?", memberId).Where("`network_name` = ?", networkName).Where("`coin_symbol` = ?", coinSymbol).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultRechargeAddressModel) FindByNetworkNameAndCoinSymbolAndAddress(networkName string, coinSymbol string, address string) (*RechargeAddress, error) {
	var model *RechargeAddress
	err := m.db.Where("`network_name` = ?", networkName).Where("`coin_symbol` = ?", coinSymbol).Where("`address` = ?", address).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func NewRechargeAddressModel(db *gorm.DB) RechargeAddressModel {
	return &defaultRechargeAddressModel{
		db: db,
	}
}
