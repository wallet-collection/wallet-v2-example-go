package res

import (
	"github.com/shopspring/decimal"
	"time"
)

// MemberCoinListRes ...
type MemberCoinListRes struct {
	MemberId       int64           `json:"member_id"`        // 用户ID
	CoinSymbol     string          `json:"coin_symbol"`      // 币种符号
	CoinName       string          `json:"coin_name"`        // 币种名称
	CoinIcon       string          `json:"coin_icon"`        // 币种图标
	UsdtPrice      decimal.Decimal `json:"usdt_price"`       // 币种价格
	Precision      int             `json:"precision"`        // 币种精度
	IsTransfer     int             `json:"is_transfer"`      // 是否可划转（0：否，1：手动，2：自动）
	TransferRate   decimal.Decimal `json:"transfer_rate"`    // 划转费率
	MinTransferFee decimal.Decimal `json:"min_transfer_fee"` // 最低划转费用
	MinTransfer    decimal.Decimal `json:"min_transfer"`     // 最低划转
	MaxTransfer    decimal.Decimal `json:"max_transfer"`     // 最大划转
	Balance        decimal.Decimal `json:"balance"`          // 余额
	FrozenBalance  decimal.Decimal `json:"frozen_balance"`   // 冻结
	VirtualBalance decimal.Decimal `json:"virtual_balance"`  // 虚拟
	CreateTime     time.Time       `json:"create_time"`      // 创建时间
	ModifiedTime   time.Time       `json:"modified_time"`    // 更新时间
}
