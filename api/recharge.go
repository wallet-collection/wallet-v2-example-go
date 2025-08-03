package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
	"wallet-example/api/req"
	"wallet-example/api/res"
	"wallet-example/biz"
	"wallet-example/client"
	"wallet-example/config"
	"wallet-example/model"
	"wallet-example/pkg/util"
)

// RechargeList 获取列表
func RechargeList(c *gin.Context) {

	var q req.RechargeListReq

	if err := c.ShouldBindQuery(&q); err != nil {
		res.APIResponse(c, err, nil)
		return
	}

	memberId := util.GetAuthMemberId(c)

	db := model.DB

	rechargeModel := model.NewRechargeModel(db)

	offset := (q.Page - 1) * q.Limit

	var startDate time.Time
	var endDate time.Time
	if len(q.StartDate) != 0 {
		startDate, _ = time.Parse("2006-01-02", q.StartDate)
	}
	if len(q.EndDate) != 0 {
		endDate, _ = time.Parse("2006-01-02", q.EndDate)
	}

	list, err := rechargeModel.List(memberId, q.CoinSymbol, &startDate, &endDate, offset, q.Limit)
	if err != nil {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	resData := make([]*res.RechargeListRes, 0)
	for _, item := range list {
		resData = append(resData, &res.RechargeListRes{
			Id:           item.Id,
			NetworkName:  item.NetworkName,
			CoinSymbol:   item.CoinSymbol,
			Address:      item.Address,
			Amount:       item.Amount,
			Status:       item.Status,
			ModifiedTime: item.ModifiedTime,
		})
	}

	res.APIResponse(c, nil, resData)
}

// RechargeInfo 获取列表
func RechargeInfo(c *gin.Context) {

	var q req.RechargeInfoReq

	if err := c.ShouldBindQuery(&q); err != nil {
		res.APIResponse(c, err, nil)
		return
	}

	memberId := util.GetAuthMemberId(c)

	db := model.DB

	rechargeModel := model.NewRechargeModel(db)

	item, err := rechargeModel.FindById(q.Id)
	if err != nil || item.MemberId != memberId {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	resData := &res.RechargeInfoRes{
		Id:           item.Id,
		BusinessId:   item.BusinessId,
		MemberId:     item.MemberId,
		NetworkName:  item.NetworkName,
		CoinSymbol:   item.CoinSymbol,
		Address:      item.Address,
		Amount:       item.Amount,
		MaxBlockHigh: item.MaxBlockHigh,
		BlockHigh:    item.BlockHigh,
		Txid:         item.Txid,
		Remark:       item.Remark,
		Status:       item.Status,
		CreateTime:   item.CreateTime,
		ModifiedTime: item.ModifiedTime,
	}

	res.APIResponse(c, nil, resData)
}

// RechargeAddress 获取充值地址
func RechargeAddress(c *gin.Context) {

	var q req.RechargeAddressReq

	if err := c.ShouldBindQuery(&q); err != nil {
		res.APIResponse(c, err, nil)
		return
	}

	memberId := util.GetAuthMemberId(c)

	db := model.DB

	walletClient := client.WalletClientGlobal

	rechargeAddressModel := model.NewRechargeAddressModel(db)

	rechargeAddress, err := rechargeAddressModel.FindByMemberIdAndNetworkNameAndCoinSymbol(memberId, q.NetworkName, q.CoinSymbol)
	if err != nil {

		callUrl := config.CONF.Wallet.CallUrl + "/recharge/call"
		// 获取
		createWallet, err := walletClient.CreateWallet(strconv.FormatInt(memberId, 10), q.NetworkName, q.CoinSymbol, callUrl)
		if err != nil {
			res.APIResponse(c, res.InternalServerError, nil)
			return
		}

		rechargeAddress = &model.RechargeAddress{
			MemberId:     memberId,
			NetworkName:  q.NetworkName,
			CoinSymbol:   q.CoinSymbol,
			Address:      createWallet.Address,
			CreateTime:   time.Now(),
			ModifiedTime: time.Now(),
		}

		err = rechargeAddressModel.Insert(rechargeAddress)
		if err != nil {
			res.APIResponse(c, res.InternalServerError, nil)
			return
		}
	}

	res.APIResponse(c, nil, rechargeAddress.Address)
}

// RechargeCall 充值回调
func RechargeCall(c *gin.Context) {

	var q req.RechargeCallReq

	if err := c.ShouldBindJSON(&q); err != nil {
		res.APIResponse(c, err, nil)
		return
	}

	// 验签
	signData := make(map[string]interface{})
	signData["appid"] = q.Appid
	signData["network_name"] = q.NetworkName
	signData["coin_symbol"] = q.CoinSymbol
	signData["decimals"] = q.Decimals
	signData["address"] = q.Address
	signData["amount"] = q.Amount.String()
	signData["business_id"] = q.BusinessId
	signData["max_block_high"] = q.MaxBlockHigh
	signData["block_high"] = q.BlockHigh
	signData["block_hash"] = q.BlockHash
	signData["txid"] = q.Txid
	signData["status"] = q.Status
	signData["sign"] = q.Sign
	signData["secret_key"] = config.CONF.Wallet.SecretKey
	signBool := client.VerifySign(signData)
	if !signBool {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	db := model.DB
	rechargeModel := model.NewRechargeModel(db)
	recharge, err := rechargeModel.FindByBusinessId(q.BusinessId)
	if err == nil && recharge.Status != 0 {
		fmt.Println("充值回调重复调用", q)
		c.String(http.StatusOK, "SUCCESS")
		return
	}

	decimals := decimal.NewFromInt(q.Decimals)
	pow := decimal.NewFromInt(10).Pow(decimals)

	amountB := q.Amount.Div(pow)

	// 根据地址和网络获取用户ID
	rechargeAddressModel := model.NewRechargeAddressModel(db)
	rechargeAddress, err := rechargeAddressModel.FindByNetworkNameAndCoinSymbolAndAddress(q.NetworkName, q.CoinSymbol, q.Address)

	if err != nil || rechargeAddress.MemberId <= 0 {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}
	memberId := rechargeAddress.MemberId

	status := q.Status
	// 成功
	if status == 1 {
		// 查询邀请关系
	}
	//fmt.Println(q)

	//dbMaxLen := 5000

	err = db.Transaction(func(tx *gorm.DB) error {
		rechargeModel = model.NewRechargeModel(tx)
		if recharge == nil {
			recharge = &model.Recharge{
				BusinessId:   q.BusinessId,
				MemberId:     memberId,
				NetworkName:  q.NetworkName,
				CoinSymbol:   q.CoinSymbol,
				Address:      q.Address,
				Amount:       amountB,
				MaxBlockHigh: q.MaxBlockHigh,
				BlockHigh:    q.BlockHigh,
				Txid:         q.Txid,
				Remark:       "",
				Status:       status,
				CreateTime:   time.Now(),
				ModifiedTime: time.Now(),
			}
			err = rechargeModel.Insert(recharge)
			if err != nil {
				return err
			}
		} else {
			rechargeMap := make(map[string]interface{})
			rechargeMap["status"] = status
			err = rechargeModel.Update(q.BusinessId, rechargeMap)
			if err != nil {
				return err
			}
		}

		if status == 1 {
			memberBillBiz := biz.MemberBillBiz{
				MemberId:     memberId,
				CoinSymbol:   q.CoinSymbol,
				Mode:         biz.BillModeIncrease,
				BusinessType: biz.BillBusinessTypeRecharge,
				BusinessId:   strconv.FormatInt(recharge.Id, 10),
				Balance:      amountB,
				Freeze:       decimal.NewFromInt(0),
				Remark:       "",
			}
			err = memberBillBiz.Create(tx)
			if err != nil {
				return err
			}

		}

		return nil
	})

	if err != nil {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	c.String(http.StatusOK, "SUCCESS")
}
