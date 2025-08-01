package res

import (
	"github.com/shopspring/decimal"
	"time"
)

// RechargeListRes ...
type RechargeListRes struct {
	Id           int64           `json:"id"`            // 自增ID
	NetworkName  string          `json:"network_name"`  // 网络名称
	CoinSymbol   string          `json:"coin_symbol"`   // 币种符号
	Address      string          `json:"address"`       // 充值地址
	Amount       decimal.Decimal `json:"amount"`        // 数量
	Status       int             `json:"status"`        // 状态（0：区块确认中，1：充值到账，2：区块确认失败）
	ModifiedTime time.Time       `json:"modified_time"` // 更新时间
}

// RechargeInfoRes ...
type RechargeInfoRes struct {
	Id           int64           `json:"id"`             // 自增ID
	BusinessId   string          `json:"business_id"`    // 业务ID
	MemberId     int64           `json:"member_id"`      // 用户ID
	NetworkName  string          `json:"network_name"`   // 网络名称
	CoinSymbol   string          `json:"coin_symbol"`    // 币种符号
	Address      string          `json:"address"`        // 充值地址
	Amount       decimal.Decimal `json:"amount"`         // 数量
	MaxBlockHigh uint64          `json:"max_block_high"` // 最大区块高度
	BlockHigh    uint64          `json:"block_high"`     // 区块高度
	Txid         string          `json:"txid"`           // 区块交易哈希
	Remark       string          `json:"remark"`         // 备注
	Status       int             `json:"status"`         // 状态（0：区块确认中，1：充值到账，2：区块确认失败）
	CreateTime   time.Time       `json:"create_time"`    // 创建时间
	ModifiedTime time.Time       `json:"modified_time"`  // 更新时间
}
