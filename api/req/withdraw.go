package req

import "github.com/shopspring/decimal"

type WithdrawListReq struct {
	ListPageReq
	CoinSymbol string `form:"coin_symbol"` // 币种符号
	StartDate  string `form:"start_date"`  // 开始时间
	EndDate    string `form:"end_date"`    // 结束时间
}

type WithdrawInfoReq struct {
	Id int64 `form:"id"` // ID
}

type WithdrawCreateReq struct {
	NetworkName string          `json:"network_name"` // 网络名称
	CoinSymbol  string          `json:"coin_symbol"`  // 币种符号
	Address     string          `json:"address"`      // 地址
	Amount      decimal.Decimal `json:"amount"`       // 提币数量
	EmailCode   string          `json:"email_code"`   // 邮箱验证码（如果绑定邮箱了需要）
	TelCode     string          `json:"tel_code"`     // 手机验证码（如果绑定手机了需要）
	GoogleCode  string          `json:"google_code"`  // 谷歌验证码（如果有的情况下）
}

type WithdrawCallReq struct {
	Appid        string          `json:"appid"`          // appid
	NetworkName  string          `json:"network_name"`   // 网络名称
	CoinSymbol   string          `json:"coin_symbol"`    // 币种符号
	Decimals     int64           `json:"decimals"`       // 币种精度
	Address      string          `json:"address"`        // 地址
	Amount       decimal.Decimal `json:"amount"`         // 充值数量
	BusinessId   string          `json:"business_id"`    // 业务ID
	MaxBlockHigh uint64          `json:"max_block_high"` // 最大区块高度
	BlockHigh    uint64          `json:"block_high"`     // 区块高度
	BlockHash    string          `json:"block_hash"`     // 区块高度
	Txid         string          `json:"txid"`           // 区块链交易哈希
	Status       int             `json:"status"`         // 状态（0: 链上打包中，1：提币成功，2：提币失败）
	Sign         string          `json:"sign"`           // 签名
}
