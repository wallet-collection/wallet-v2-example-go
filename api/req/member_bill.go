package req

type MemberBillListReq struct {
	ListPageReq
	CoinSymbol   string `form:"coin_symbol"`   // 币种
	Mode         int    `form:"mode"`          // 类型（0：划转，1：收入，2：支出）
	BusinessType string `form:"business_type"` // 业务类型
	StartDate    string `form:"start_date"`    // 开始时间
	EndDate      string `form:"end_date"`      // 结束时间
}
