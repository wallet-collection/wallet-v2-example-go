package api

import (
	"github.com/gin-gonic/gin"
	"time"
	"wallet-example/api/req"
	"wallet-example/api/res"
	"wallet-example/biz"
	"wallet-example/model"
	"wallet-example/pkg/util"
)

// MemberBillBusinessType 获取业务类型列表
func MemberBillBusinessType(c *gin.Context) {

	var businessTypeList []string
	for _, businessType := range biz.BillBusinessTypeList {
		businessTypeList = append(businessTypeList, businessType.Value)
	}

	res.APIResponse(c, nil, businessTypeList)
}

// MemberBillList 获取列表
func MemberBillList(c *gin.Context) {

	var q req.MemberBillListReq

	if err := c.ShouldBindQuery(&q); err != nil {
		res.APIResponse(c, err, nil)
		return
	}

	memberId := util.GetAuthMemberId(c)

	db := model.DB

	memberBillModel := model.NewMemberBillModel(db)

	offset := (q.Page - 1) * q.Limit

	var startDate time.Time
	var endDate time.Time
	if len(q.StartDate) != 0 {
		startDate, _ = time.Parse("2006-01-02", q.StartDate)
	}
	if len(q.EndDate) != 0 {
		endDate, _ = time.Parse("2006-01-02", q.EndDate)
	}

	list, err := memberBillModel.List(memberId, q.CoinSymbol, q.Mode, q.BusinessType, &startDate, &endDate, offset, q.Limit)
	if err != nil {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	businessTypeMap := make(map[string]string)
	for _, businessType := range biz.BillBusinessTypeList {
		businessTypeMap[businessType.Value] = businessType.Name
	}

	resData := make([]*res.MemberBillListRes, 0)
	for _, item := range list {

		resData = append(resData, &res.MemberBillListRes{
			Id:               item.Id,
			MemberId:         item.MemberId,
			CoinSymbol:       item.CoinSymbol,
			FromAccount:      item.FromAccount,
			ToAccount:        item.ToAccount,
			Mode:             item.Mode,
			BusinessType:     item.BusinessType,
			BusinessTypeText: businessTypeMap[item.BusinessType],
			BusinessId:       item.BusinessId,
			Amount:           item.Amount,
			Remark:           item.Remark,
			CreateTime:       item.CreateTime,
			ModifiedTime:     item.ModifiedTime,
		})
	}

	res.APIResponse(c, nil, resData)
}
