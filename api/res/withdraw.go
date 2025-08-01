package res

import (
	"github.com/shopspring/decimal"
	"time"
)

// WithdrawListRes ...
type WithdrawListRes struct {
	Id           int64           `json:"id"`            // 自增ID
	NetworkName  string          `json:"network_name"`  // 网络名称
	CoinSymbol   string          `json:"coin_symbol"`   // 币种符号
	Address      string          `json:"address"`       // 充值地址
	Amount       decimal.Decimal `json:"amount"`        // 数量
	Fee          decimal.Decimal `json:"fee"`           // 手续费
	Status       int             `json:"status"`        // 状态（0：审核中，1：审核通过，2：审核不通过，3：链上打包中，4：提币成功，5：提币失败，6：手动成功）
	ModifiedTime time.Time       `json:"modified_time"` // 更新时间
}

// WithdrawInfoRes ...
type WithdrawInfoRes struct {
	Id           int64           `json:"id"`            // 自增ID
	MemberId     int64           `json:"member_id"`     // 用户ID
	NetworkName  string          `json:"network_name"`  // 网络名称
	CoinSymbol   string          `json:"coin_symbol"`   // 币种符号
	Address      string          `json:"address"`       // 提现地址
	Amount       decimal.Decimal `json:"amount"`        // 数量
	Fee          decimal.Decimal `json:"fee"`           // 手续费
	BlockHigh    uint64          `json:"block_high"`    // 区块高度
	Txid         string          `json:"txid"`          // 区块交易哈希
	Remark       string          `json:"remark"`        // 备注
	Status       int             `json:"status"`        // 状态（0：审核中，1：审核通过，2：审核不通过，3：链上打包中，4：提币成功，5：提币失败，6：手动成功）
	CreateTime   time.Time       `json:"create_time"`   // 创建时间
	ModifiedTime time.Time       `json:"modified_time"` // 更新时间
}
