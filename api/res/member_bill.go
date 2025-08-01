package res

import (
	"github.com/shopspring/decimal"
	"time"
)

// MemberBillListRes ...
type MemberBillListRes struct {
	Id               int64           `json:"id"`                 // 自增ID
	MemberId         int64           `json:"member_id"`          // 用户ID
	CoinSymbol       string          `json:"coin_symbol"`        // 币种符号
	FromAccount      string          `json:"from_account"`       // 来源账户
	ToAccount        string          `json:"to_account"`         // 接收账户
	Mode             int             `json:"mode"`               // 类型（0：划转，1：收入，2：支出）
	BusinessType     string          `json:"business_type"`      // 业务类型
	BusinessTypeText string          `json:"business_type_text"` // 业务类型说明
	BusinessId       string          `json:"business_id"`        // 业务ID
	Amount           decimal.Decimal `json:"amount"`             // 数量
	Remark           string          `json:"remark"`             // 备注s
	CreateTime       time.Time       `json:"create_time"`        // 创建时间
	ModifiedTime     time.Time       `json:"modified_time"`      // 更新时间
}
