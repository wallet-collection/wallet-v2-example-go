package req

type LoginReq struct {
	Account string `json:"account" binding:"required"` // 地址
	Pwd     string `json:"pwd" binding:"required"`     // 密码
}

type RegisterReq struct {
	Account string `json:"account" binding:"required"` // 账号（邮箱/手机号）手机需要带区号如 86_15123246666
	Invite  string `json:"invite"`                     // 邀请码
	Pwd     string `json:"pwd" binding:"required"`     // 密码
}

type ForgotPasswordReq struct {
	Account string `json:"account" binding:"required"` // 账号（邮箱/手机号）手机需要带区号如 86_15123246666
	Code    string `json:"code" binding:"required"`    // 验证码
	Pwd     string `json:"pwd" binding:"required"`     // 密码
}
