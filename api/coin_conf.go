package api

import (
	"github.com/gin-gonic/gin"
	"wallet/api/req"
	"wallet/api/res"
	"wallet/model"
)

// CoinConfList 获取列表
func CoinConfList(c *gin.Context) {

	var q req.CoinConfListReq

	if err := c.ShouldBindQuery(&q); err != nil {
		res.APIResponse(c, err, nil)
		return
	}

	db := model.DB

	coinConfModel := model.NewCoinConfModel(db)

	list, err := coinConfModel.ListAll(q.CoinSymbol)
	if err != nil {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	resData := make([]*res.CoinConfListRes, 0)
	for _, item := range list {

		resData = append(resData, &res.CoinConfListRes{
			Id:              item.Id,
			CoinSymbol:      item.CoinSymbol,
			NetworkName:     item.NetworkName,
			Decimals:        item.Decimals,
			IsWithdraw:      item.IsWithdraw,
			WithdrawRate:    item.WithdrawRate,
			MinWithdrawFee:  item.MinWithdrawFee,
			MinWithdraw:     item.MinWithdraw,
			MaxWithdraw:     item.MaxWithdraw,
			MinRecharge:     item.MinRecharge,
			IsRecharge:      item.IsRecharge,
			RechargeConfirm: item.RechargeConfirm,
			WithdrawConfirm: item.WithdrawConfirm,
			Sort:            item.Sort,
			Status:          item.Status,
			CreateTime:      item.CreateTime,
			ModifiedTime:    item.ModifiedTime,
		})
	}

	res.APIResponse(c, nil, resData)
}
