package api

import (
	"github.com/gin-gonic/gin"
	"wallet/api/req"
	"wallet/api/res"
	"wallet/biz"
	"wallet/model"
	"wallet/pkg/util"
)

// MemberCoinInfo 获取列表
func MemberCoinInfo(c *gin.Context) {

	var q req.MemberCoinInfoReq

	if err := c.ShouldBindQuery(&q); err != nil {
		res.APIResponse(c, err, nil)
		return
	}

	memberId := util.GetAuthMemberId(c)

	db := model.DB

	memberCoinBiz := biz.MemberCoinBiz{}
	memberCoin, err := memberCoinBiz.Balance(memberId, q.CoinSymbol, db)
	if err != nil {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	resData := &res.MemberCoinListRes{
		MemberId:       memberId,
		CoinSymbol:     q.CoinSymbol,
		CoinName:       memberCoin.CoinName,
		CoinIcon:       memberCoin.CoinIcon,
		UsdtPrice:      memberCoin.UsdtPrice,
		Precision:      memberCoin.Precision,
		IsTransfer:     memberCoin.IsTransfer,
		TransferRate:   memberCoin.TransferRate,
		MinTransferFee: memberCoin.MinTransfer,
		MinTransfer:    memberCoin.MinTransferFee,
		MaxTransfer:    memberCoin.MaxTransfer,
		Balance:        memberCoin.Balance,
		FrozenBalance:  memberCoin.FrozenBalance,
		VirtualBalance: memberCoin.VirtualBalance,
		CreateTime:     memberCoin.CreateTime,
		ModifiedTime:   memberCoin.ModifiedTime,
	}

	res.APIResponse(c, nil, resData)
}

// MemberCoinList 获取列表
func MemberCoinList(c *gin.Context) {

	memberId := util.GetAuthMemberId(c)

	db := model.DB

	memberCoinBiz := biz.MemberCoinBiz{}
	_, list, err := memberCoinBiz.ListAll(memberId, 0, 0, db)
	if err != nil {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	resData := make([]*res.MemberCoinListRes, 0)
	for _, item := range list {
		resData = append(resData, &res.MemberCoinListRes{
			MemberId:       item.MemberId,
			CoinSymbol:     item.CoinSymbol,
			CoinName:       item.CoinName,
			CoinIcon:       item.CoinIcon,
			UsdtPrice:      item.UsdtPrice,
			Precision:      item.Precision,
			IsTransfer:     item.IsTransfer,
			TransferRate:   item.TransferRate,
			MinTransferFee: item.MinTransferFee,
			MinTransfer:    item.MinTransfer,
			MaxTransfer:    item.MaxTransfer,
			Balance:        item.Balance,
			FrozenBalance:  item.FrozenBalance,
			VirtualBalance: item.VirtualBalance,
			CreateTime:     item.CreateTime,
			ModifiedTime:   item.ModifiedTime,
		})
	}

	res.APIResponse(c, nil, resData)
}
