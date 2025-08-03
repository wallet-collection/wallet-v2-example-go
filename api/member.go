package api

import (
	"github.com/gin-gonic/gin"
	"wallet-example/api/res"
	"wallet-example/model"
	"wallet-example/pkg/util"
)

// MemberLoginInfo 登录
func MemberLoginInfo(c *gin.Context) {

	memberId := util.GetAuthMemberId(c)
	if memberId <= 0 {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	db := model.DB

	memberModel := model.NewMemberModel(db)
	member, err := memberModel.FindById(memberId)
	if err != nil {
		res.APIResponse(c, res.ErrToken, nil)
		return
	}

	if member.Status != 1 {
		res.APIResponse(c, res.ErrToken, nil)
		return
	}

	token := util.CreateToken(member.Id)

	isPayPassword := 0
	isGoogleKey := 0
	FishingCode := ""
	if len(member.PayPwd) > 0 {
		isPayPassword = 1
	}
	if len(member.GoogleKey) > 0 {
		isGoogleKey = 1
	}
	if len(member.FishingCode) > 0 {
		FishingCode = member.FishingCode
	}

	tel := ""
	if member.Tel != nil {
		tel = *member.Tel
	}
	email := ""
	if member.Email != nil {
		email = *member.Email
	}

	// 查询今日数据
	resData := &res.MemberLoginInfoRes{
		Id:          memberId,
		Nickname:    member.Nickname,
		Avatar:      member.Avatar,
		Tel:         tel,
		Email:       email,
		Invite:      util.IdToCode(memberId),
		Token:       token,
		IsPayPwd:    isPayPassword,
		IsGoogleKey: isGoogleKey,
		FishingCode: FishingCode,
	}

	res.APIResponse(c, nil, resData)
}
