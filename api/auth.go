package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"time"
	"wallet/api/req"
	"wallet/api/res"
	"wallet/model"
	"wallet/pkg/util"
)

// AuthLoginByPwd 登录
func AuthLoginByPwd(c *gin.Context) {

	var q req.LoginReq

	if err := c.ShouldBindJSON(&q); err != nil {
		res.APIResponse(c, err, nil)
		return
	}

	account := q.Account

	db := model.DB

	memberModel := model.NewMemberModel(db)
	var member *model.Member
	var err error
	if strings.Contains(account, "@") {
		member, err = memberModel.FindByEmail(account)
	} else {
		member, err = memberModel.FindByTel(account)
	}
	if err != nil || member.Status != 1 {
		res.APIResponse(c, res.ErrLoginNot, nil)
		return
	}

	if member.Pwd != util.CreatePwd(q.Pwd) {
		res.APIResponse(c, res.ErrLoginNot, nil)
		return
	}

	token := util.CreateToken(member.Id)

	res.APIResponse(c, nil, &res.LoginRes{
		Id:    member.Id,
		Token: token,
	})
}

// AuthRegister 注册
func AuthRegister(c *gin.Context) {

	var q req.RegisterReq

	if err := c.ShouldBindJSON(&q); err != nil {
		res.APIResponse(c, err, nil)
		return
	}

	account := q.Account

	db := model.DB

	nickname := ""
	memberModel := model.NewMemberModel(db)
	var tel *string
	var email *string
	var info *model.Member
	var err error
	if strings.Contains(account, "@") {
		info, err = memberModel.FindByEmail(account)
		nickname = util.HideEmail(account)
		email = &account
	} else {
		info, err = memberModel.FindByTel(account)
		nickname = util.HidePhoneNumber(account)
		tel = &account
	}
	if err == nil && info != nil {
		if info.Status != 1 {
			res.APIResponse(c, res.ErrLoginNot, nil)
			return
		}
		//token := util.CreateToken(info.Id)
		//
		//res.APIResponse(c, nil, &res.LoginRes{
		//	Id:    info.Id,
		//	Token: token,
		//})
		res.APIResponse(c, res.ErrNotRepeatData, nil)
		return
	}

	pid := int64(0)

	pwd := util.CreatePwd(q.Pwd)

	member := &model.Member{
		Pid:            pid,
		Nickname:       nickname,
		Avatar:         "",
		Remark:         "",
		Tel:            tel,
		Email:          email,
		Pwd:            pwd,
		Status:         1,
		LastUpdateSafe: time.Now().AddDate(0, 0, -1),
		CreateTime:     time.Now(),
		ModifiedTime:   time.Now(),
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		err := model.NewMemberModel(tx).Insert(member)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	token := util.CreateToken(member.Id)

	res.APIResponse(c, nil, &res.LoginRes{
		Id:    member.Id,
		Token: token,
	})
}

// AuthForgotPassword 忘记密码
func AuthForgotPassword(c *gin.Context) {

	var q req.ForgotPasswordReq

	if err := c.ShouldBindJSON(&q); err != nil {
		res.APIResponse(c, err, nil)
		return
	}

	account := q.Account

	db := model.DB

	//rdb := redis.RDB
	var err error
	// 验证验证码
	//err := biz.NewCodeBiz(db, rdb).Verify(biz.CodeTemplateSceneForgotPassword, account, q.Code, 0)
	//if err != nil {
	//	res.APIResponse(c, err, nil)
	//	return
	//}

	memberModel := model.NewMemberModel(db)
	var info *model.Member
	if strings.Contains(account, "@") {
		info, err = memberModel.FindByEmail(account)
	} else {
		info, err = memberModel.FindByTel(account)
	}
	if err != nil {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}
	memberId := info.Id

	memberUp := make(map[string]interface{})
	memberUp["pwd"] = util.CreatePwd(q.Pwd)
	err = memberModel.Update(memberId, memberUp)

	if err != nil {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	res.APIResponse(c, nil, nil)
}
