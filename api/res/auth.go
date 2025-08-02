package res

// LoginRes ...
type LoginRes struct {
	Id    int64  `json:"id"`    // 用户ID
	Token string `json:"token"` // 登录的token
}
