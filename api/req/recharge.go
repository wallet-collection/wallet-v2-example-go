package req

import "github.com/shopspring/decimal"

type RechargeListReq struct {
	ListPageReq
	CoinSymbol string `form:"coin_symbol"` // 币种符号
	StartDate  string `form:"start_date"`  // 开始时间
	EndDate    string `form:"end_date"`    // 结束时间
}

type RechargeInfoReq struct {
	Id int64 `form:"id"` // Id
}

type RechargeAddressReq struct {
	NetworkName string `form:"network_name"` // 网络名称
	CoinSymbol  string `form:"coin_symbol"`  // 币种符号
}

type RechargeCallReq struct {
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
	Status       int             `json:"status"`         // 状态（0: 链上打包中，1：充值成功，2：充值失败）
	Sign         string          `json:"sign"`           // 签名
}
