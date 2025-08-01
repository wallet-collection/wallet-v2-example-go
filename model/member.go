package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type (
	MemberModel interface {
		Insert(model *Member) error
		ListByPage(id int64, tel string, email string, status, offset, limit int) (int64, []*Member, error)
		ListByPageExport(id int64, tel string, email string, status int) ([]*Member, error)
		ListPageByPid(pid int64, offset, limit int) (int64, []*Member, error)
		ListByPid(pid int64, offset, limit int) ([]*Member, error)
		ListByInPid(pid []int64) ([]*Member, error)
		FindByInMemberIds(memberIds []int64) ([]*Member, error)
		FindById(id int64) (*Member, error)
		FindByInId(ids []int64) ([]*Member, error)
		FindByTel(tel string) (*Member, error)
		FindByEmail(email string) (*Member, error)
		Update(id int64, data map[string]interface{}) error
	}

	defaultMemberModel struct {
		db *gorm.DB
	}

	Member struct {
		Id             int64     // 用户ID
		Pid            int64     // 直推上级
		Tel            *string   // 手机号（区号用下划线隔开）
		Email          *string   // 邮箱号
		Nickname       string    // 用户昵称
		Avatar         string    // 用户头像
		Pwd            string    // 密码
		PayPwd         string    // 支付密码
		GoogleKey      string    // 谷歌验证key
		FishingCode    string    // 防钓鱼码
		LastUpdateSafe time.Time // 最后修改安全项的时间
		Remark         string    // 个性签名
		Status         int       // 状态（0：禁用，1：正常）
		CreateTime     time.Time // 创建时间
		ModifiedTime   time.Time // 更新时间
	}
)

func (m *defaultMemberModel) Insert(model *Member) error {
	err := m.db.Model(&Member{}).Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultMemberModel) ListByPage(id int64, tel string, email string, status, offset, limit int) (int64, []*Member, error) {
	var model []*Member
	db := m.db.Model(&Member{})
	if id > 0 {
		db.Where("`id` = ?", id)
	}
	if status >= 0 {
		db.Where("`status` = ?", status)
	}
	if len(tel) > 0 {
		db.Where("`tel` = ?", tel)
	}
	if len(email) > 0 {
		db.Where("`email` = ?", email)
	}
	var total int64
	db.Count(&total)
	err := db.Offset(offset).Limit(limit).Order("`id` DESC").Find(&model).Error
	if err != nil {
		return 0, nil, err
	}
	return total, model, nil
}

func (m *defaultMemberModel) ListByPageExport(id int64, tel string, email string, status int) ([]*Member, error) {
	var model []*Member
	db := m.db.Model(&Member{})
	if id > 0 {
		db.Where("`id` = ?", id)
	}
	if status >= 0 {
		db.Where("`status` = ?", status)
	}
	if len(tel) > 0 {
		db.Where("`tel` = ?", tel)
	}
	if len(email) > 0 {
		db.Where("`email` = ?", email)
	}
	var total int64
	db.Count(&total)
	err := db.Order("`id` DESC").Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultMemberModel) ListPageByPid(pid int64, offset, limit int) (int64, []*Member, error) {
	var model []*Member
	db := m.db.Model(&Member{})
	db.Where("`pid` = ?", pid)
	var total int64
	db.Count(&total)
	err := db.Offset(offset).Limit(limit).Order("`id` DESC").Find(&model).Error
	if err != nil {
		return 0, nil, err
	}
	return total, model, nil
}

func (m *defaultMemberModel) ListByPid(pid int64, offset, limit int) ([]*Member, error) {
	var model []*Member
	db := m.db.Model(&Member{})
	db.Where("`pid` = ?", pid)
	err := db.Offset(offset).Limit(limit).Order("`id` DESC").Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultMemberModel) ListByInPid(pid []int64) ([]*Member, error) {
	var model []*Member
	db := m.db.Model(&Member{})
	db.Where("`pid` IN(?)", pid)
	err := db.Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultMemberModel) FindByInMemberIds(memberIds []int64) ([]*Member, error) {
	var model []*Member
	err := m.db.Where("`id` IN (?)", memberIds).Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultMemberModel) FindById(id int64) (*Member, error) {
	var model *Member
	err := m.db.Where("id = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultMemberModel) FindByInId(ids []int64) ([]*Member, error) {
	var model []*Member
	err := m.db.Where("id IN (?)", ids).Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultMemberModel) FindByEmail(email string) (*Member, error) {
	var model *Member
	err := m.db.Where("`email` = ?", email).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultMemberModel) FindByTel(tel string) (*Member, error) {
	var model *Member
	err := m.db.Where("`tel` = ?", tel).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *defaultMemberModel) Update(id int64, data map[string]interface{}) error {
	if data == nil {
		return errors.New("data not")
	}
	err := m.db.Model(&Member{}).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func NewMemberModel(db *gorm.DB) MemberModel {
	return &defaultMemberModel{
		db: db,
	}
}
