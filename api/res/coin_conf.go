package res

import (
	"github.com/shopspring/decimal"
	"time"
)

// CoinConfListRes ...
type CoinConfListRes struct {
	Id              int64           `json:"id"`               // 自增ID
	CoinSymbol      string          `json:"coin_symbol"`      // 币种符号
	NetworkName     string          `json:"network_name"`     // 网络名称
	Decimals        int             `json:"decimals"`         // 币种精度
	IsWithdraw      int             `json:"is_withdraw"`      // 是否可提现（0：否，1：手动，2：自动）
	WithdrawRate    decimal.Decimal `json:"withdraw_rate"`    // 提现费率
	MinWithdrawFee  decimal.Decimal `json:"min_withdraw_fee"` // 最低提现费用
	MinWithdraw     decimal.Decimal `json:"min_withdraw"`     // 最低提现
	MaxWithdraw     decimal.Decimal `json:"max_withdraw"`     // 最大提现
	MinRecharge     decimal.Decimal `json:"min_recharge"`     // 最低充值
	IsRecharge      int             `json:"is_recharge"`      // 是否可充值（0：手动，1：自动）
	RechargeConfirm int             `json:"recharge_confirm"` // 充值确认数
	WithdrawConfirm int             `json:"withdraw_confirm"` // 提现确认数
	Sort            int             `json:"sort"`             // 排序（升序）
	Status          int             `json:"status"`           // 状态（0：禁用，1：正常）
	CreateTime      time.Time       `json:"create_time"`      // 创建时间
	ModifiedTime    time.Time       `json:"modified_time"`    // 更新时间
}

type CoinConfInfoRes struct {
	Balance decimal.Decimal    `json:"balance"` // 余额
	List    []*CoinConfListRes `json:"list"`    // 列表
}
