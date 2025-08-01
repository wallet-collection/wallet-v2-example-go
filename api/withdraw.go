package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
	"wallet/api/req"
	"wallet/api/res"
	"wallet/biz"
	"wallet/client"
	"wallet/config"
	"wallet/model"
	"wallet/pkg/util"
	"wallet/redis"
)

var big0 = decimal.NewFromInt(0)

// WithdrawList 获取列表
func WithdrawList(c *gin.Context) {

	var q req.WithdrawListReq

	if err := c.ShouldBindQuery(&q); err != nil {
		res.APIResponse(c, err, nil)
		return
	}

	memberId := util.GetAuthMemberId(c)

	db := model.DB

	withdrawModel := model.NewWithdrawModel(db)

	offset := (q.Page - 1) * q.Limit

	var startDate time.Time
	var endDate time.Time
	if len(q.StartDate) != 0 {
		startDate, _ = time.Parse("2006-01-02", q.StartDate)
	}
	if len(q.EndDate) != 0 {
		endDate, _ = time.Parse("2006-01-02", q.EndDate)
	}

	list, err := withdrawModel.List(memberId, q.CoinSymbol, &startDate, &endDate, offset, q.Limit)
	if err != nil {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	resData := make([]*res.WithdrawListRes, 0)
	for _, item := range list {
		resData = append(resData, &res.WithdrawListRes{
			Id:           item.Id,
			NetworkName:  item.NetworkName,
			CoinSymbol:   item.CoinSymbol,
			Address:      item.Address,
			Amount:       item.Amount,
			Fee:          item.Fee,
			Status:       item.Status,
			ModifiedTime: item.ModifiedTime,
		})
	}

	res.APIResponse(c, nil, resData)
}

// WithdrawInfo 获取列表
func WithdrawInfo(c *gin.Context) {

	var q req.WithdrawInfoReq

	if err := c.ShouldBindQuery(&q); err != nil {
		res.APIResponse(c, err, nil)
		return
	}

	memberId := util.GetAuthMemberId(c)

	db := model.DB

	withdrawModel := model.NewWithdrawModel(db)

	item, err := withdrawModel.FindById(q.Id)
	if err != nil || item.MemberId != memberId {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	resData := &res.WithdrawInfoRes{
		Id:           item.Id,
		MemberId:     item.MemberId,
		NetworkName:  item.NetworkName,
		CoinSymbol:   item.CoinSymbol,
		Address:      item.Address,
		Amount:       item.Amount,
		Fee:          item.Fee,
		BlockHigh:    item.BlockHigh,
		Txid:         item.Txid,
		Remark:       item.Remark,
		Status:       item.Status,
		CreateTime:   item.CreateTime,
		ModifiedTime: item.ModifiedTime,
	}

	res.APIResponse(c, nil, resData)
}

// WithdrawCreate 提现
func WithdrawCreate(c *gin.Context) {

	var q req.WithdrawCreateReq

	if err := c.ShouldBindJSON(&q); err != nil {
		res.APIResponse(c, err, nil)
		return
	}

	memberId := util.GetAuthMemberId(c)

	db := model.DB

	fmt.Println(q)
	memberModel := model.NewMemberModel(db)
	member, err := memberModel.FindById(memberId)
	if err != nil {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	// 判断是否有谷歌验证码
	if len(member.GoogleKey) > 0 {
		verifyCode, err := util.NewGoogleAuth().VerifyCode(member.GoogleKey, q.GoogleCode)
		if err != nil || !verifyCode {
			res.APIResponse(c, res.ErrCodeErr, nil)
			return
		}
	}

	rdb := redis.RDB

	if member.Email != nil {
		// 验证新验证码
		err = biz.NewCodeBiz(rdb).Verify(biz.CodeTemplateSceneWithdraw, *member.Email, q.EmailCode, 0)
		if err != nil {
			res.APIResponse(c, res.ErrCodeErr, nil)
			return
		}
	}

	if member.Tel != nil {
		// 验证新验证码
		err = biz.NewCodeBiz(rdb).Verify(biz.CodeTemplateSceneWithdraw, *member.Tel, q.TelCode, 0)
		if err != nil {
			res.APIResponse(c, res.ErrCodeErr, nil)
			return
		}
	}

	walletClient := client.WalletClientGlobal

	coinConfModel := model.NewCoinConfModel(db)
	coinConf, err := coinConfModel.FindByNetworkNameAndCoinSymbol(q.CoinSymbol, q.NetworkName)
	if err != nil {
		res.APIResponse(c, res.ErrNetworkNot, nil)
		return
	}

	// 不可提现
	if coinConf.IsWithdraw == 0 || coinConf.Status != 1 {
		res.APIResponse(c, res.ErrNetworkNot, nil)
		return
	}

	// 判断最低提现
	if coinConf.MinWithdraw.Cmp(q.Amount) > 0 {
		res.APIResponse(c, res.ErrMinWithdraw, nil)
		return
	}
	// 判断最大提现
	if coinConf.MaxWithdraw.Cmp(big0) > 0 && coinConf.MaxWithdraw.Cmp(q.Amount) < 0 {
		res.APIResponse(c, res.ErrMaxWithdraw, nil)
		return
	}

	fee := coinConf.WithdrawRate.Mul(q.Amount)
	// 判断收费小于最低手续费
	if fee.Cmp(coinConf.MinWithdrawFee) < 0 {
		fee = coinConf.MinWithdrawFee
	}

	withdrawAmount := q.Amount.Sub(fee)

	if withdrawAmount.Cmp(big0) <= 0 {
		res.APIResponse(c, res.ErrMinWithdraw, nil)
		return
	}

	status := 0
	if coinConf.IsWithdraw == 2 && (coinConf.WithdrawAuto.Cmp(decimal.NewFromInt(0)) == 0 || q.Amount.Cmp(coinConf.WithdrawAuto) <= 0) {
		status = 1
	}

	withdraw := &model.Withdraw{
		MemberId:     memberId,
		NetworkName:  q.NetworkName,
		CoinSymbol:   q.CoinSymbol,
		Address:      q.Address,
		Amount:       withdrawAmount,
		Fee:          fee,
		Remark:       "",
		Status:       status,
		CreateTime:   time.Now(),
		ModifiedTime: time.Now(),
	}

	txid := ""

	err = db.Transaction(func(tx *gorm.DB) error {

		withdrawModel := model.NewWithdrawModel(tx)

		err = withdrawModel.Insert(withdraw)
		if err != nil {
			return err
		}

		businessId := strconv.FormatInt(withdraw.Id, 10)

		memberBillBiz := biz.MemberBillBiz{
			MemberId:     memberId,
			CoinSymbol:   q.CoinSymbol,
			Mode:         biz.BillModeReduce,
			BusinessType: biz.BillBusinessTypeWithdraw,
			BusinessId:   businessId,
			Balance:      q.Amount.Neg(),
			Freeze:       q.Amount,
			Remark:       "",
		}
		err = memberBillBiz.Create(tx)
		if err != nil {
			return err
		}

		// 判断配置，如果是自动的
		if coinConf.IsWithdraw == 2 && (coinConf.WithdrawAuto.Cmp(decimal.NewFromInt(0)) == 0 || q.Amount.Cmp(coinConf.WithdrawAuto) <= 0) {
			withdrawPrivateKey := ""
			pwdCode, err := util.DePwdCode(coinConf.WithdrawPrivateKey)
			if err == nil {
				withdrawPrivateKey = string(pwdCode)
			}
			callUrl := config.CONF.Wallet.CallUrl + "/withdraw/call"
			txres, err := walletClient.CreateWithdraw(q.NetworkName, q.CoinSymbol, q.Address, q.Amount, businessId, withdrawPrivateKey, callUrl)
			if err != nil {
				return err
			}
			txid = txres.Hash
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	if len(txid) > 0 {
		up := make(map[string]interface{})
		up["txid"] = txid
		_ = model.NewWithdrawModel(db).Update(withdraw.Id, up, []int{1, 3})
	}

	res.APIResponse(c, nil, nil)
}

// WithdrawCall 提现回调
func WithdrawCall(c *gin.Context) {

	var q req.WithdrawCallReq

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

	withdrawModel := model.NewWithdrawModel(db)

	id, err := strconv.ParseInt(q.BusinessId, 10, 64)
	if err != nil {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	withdraw, err := withdrawModel.FindById(id)
	if err != nil || withdraw.Status > 3 {
		fmt.Println("提现回调重复调用", q)
		c.String(http.StatusOK, "SUCCESS")
		return
	}

	// 只修改状态
	if q.Status == 0 {
		data := make(map[string]interface{})
		data["block_high"] = q.BlockHigh
		data["txid"] = q.Txid
		data["status"] = 3
		data["modified_time"] = time.Now()
		err = withdrawModel.Update(id, data, []int{1, 3})
		if err != nil {
			res.APIResponse(c, res.InternalServerError, nil)
			return
		}
		c.String(http.StatusOK, "SUCCESS")
		return
	}

	memberId := withdraw.MemberId

	amount := withdraw.Amount.Add(withdraw.Fee)

	status := 4
	freeze := amount.Neg()
	balance := decimal.NewFromInt(0)
	if q.Status == 2 {
		// 失败
		status = 5
		// 增加余额
		balance = amount
	}

	// 状态为成功
	if q.Status == 1 {
		// 查询邀请关系
	}

	// 查询当前的冻结余额
	memberCoin, err := model.NewMemberCoinModel(db).FindByMemberIdAndCoinSymbol(memberId, q.CoinSymbol)
	if err != nil {
		res.APIResponse(c, err, nil)
		return
	}

	// 判断冻结金额和当前金额相差多少，因为用精度问题
	if q.Amount.Cmp(memberCoin.FrozenBalance) > 0 {
		freeze = memberCoin.FrozenBalance.Neg()
	}

	//dbMaxLen := 5000

	err = db.Transaction(func(tx *gorm.DB) error {

		withdrawModel = model.NewWithdrawModel(tx)
		data := make(map[string]interface{})
		data["block_high"] = q.BlockHigh
		data["txid"] = q.Txid
		data["status"] = status
		data["modified_time"] = time.Now()
		err = withdrawModel.Update(id, data, []int{1, 3})
		if err != nil {
			return err
		}

		memberBillBiz := biz.MemberBillBiz{
			MemberId:     memberId,
			CoinSymbol:   withdraw.CoinSymbol,
			Mode:         biz.BillModeIncrease,
			BusinessType: biz.BillBusinessTypeWithdraw,
			BusinessId:   strconv.FormatInt(withdraw.Id, 10),
			Balance:      balance,
			Freeze:       freeze,
			Remark:       "",
		}
		err = memberBillBiz.Create(tx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		res.APIResponse(c, res.InternalServerError, nil)
		return
	}

	c.String(http.StatusOK, "SUCCESS")
}
