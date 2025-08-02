package res

// MemberLoginInfoRes ...
type MemberLoginInfoRes struct {
	Id          int64  `json:"id"`            // 用户ID
	Nickname    string `json:"nickname"`      // 用户昵称
	Avatar      string `json:"avatar"`        // 用户头像
	Tel         string `json:"tel"`           // 手机号
	Email       string `json:"email"`         // 地址
	Invite      string `json:"invite"`        // 邀请码
	Token       string `json:"token"`         // 登录的token
	IsPayPwd    int    `json:"is_pay_pwd"`    // 是否设置了支付密码
	IsGoogleKey int    `json:"is_google_key"` // 是否设置了谷歌key
	FishingCode string `json:"fishing_code"`  // 防钓鱼码
}
