package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type (
	MemberBillModel interface {
		InsertAll(model []*MemberBill) error
		ListByPage(memberId int64, coinSymbol string, mode int, businessType string, offset, limit int) (int64, []*MemberBill, error)
		List(memberId int64, coinSymbol string, mode int, businessType string, startDate *time.Time, endDate *time.Time, offset, limit int) ([]*MemberBill, error)
	}

	defaultMemberBillModel struct {
		db *gorm.DB
	}

	MemberBill struct {
		Id           int64           // 自增ID
		MemberId     int64           // 用户ID
		CoinSymbol   string          // 币种符号
		FromAccount  string          // 来源账户
		ToAccount    string          // 接收账户
		Mode         int             // 类型（0：划转，1：收入，2：支出）
		BusinessType string          // 业务类型
		BusinessId   string          // 业务ID
		Amount       decimal.Decimal // 数量
		Remark       string          // 备注
		CreateTime   time.Time       // 创建时间
		ModifiedTime time.Time       // 更新时间
	}
)

func (m *defaultMemberBillModel) InsertAll(model []*MemberBill) error {
	err := m.db.Model(&MemberBill{}).Omit("id").Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultMemberBillModel) ListByPage(memberId int64, coinSymbol string, mode int, businessType string, offset, limit int) (int64, []*MemberBill, error) {
	var model []*MemberBill
	db := m.db.Model(&MemberBill{})
	if memberId > 0 {
		db.Where("`member_id` = ?", memberId)
	}
	if len(coinSymbol) > 0 {
		db.Where("`coin_symbol` = ?", coinSymbol)
	}
	if mode > 0 {
		db.Where("`mode` = ?", mode)
	}
	if len(businessType) > 0 {
		db.Where("`business_type` = ?", businessType)
	}
	var total int64
	db.Count(&total)
	err := db.Offset(offset).Limit(limit).Order("`create_time` DESC").Find(&model).Error
	if err != nil {
		return 0, nil, err
	}
	return total, model, nil
}

func (m *defaultMemberBillModel) List(memberId int64, coinSymbol string, mode int, businessType string, startDate *time.Time, endDate *time.Time, offset, limit int) ([]*MemberBill, error) {
	var model []*MemberBill
	db := m.db.Model(&MemberBill{})
	if memberId > 0 {
		db.Where("`member_id` = ?", memberId)
	}
	if len(coinSymbol) > 0 {
		db.Where("`coin_symbol` = ?", coinSymbol)
	}
	if mode > 0 {
		db.Where("`mode` = ?", mode)
	}
	if len(businessType) > 0 {
		db.Where("`business_type` = ?", businessType)
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

func NewMemberBillModel(db *gorm.DB) MemberBillModel {
	return &defaultMemberBillModel{
		db: db,
	}
}
