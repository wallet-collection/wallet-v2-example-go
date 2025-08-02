package req

// MemberEditPwdReq ...
type MemberEditPwdReq struct {
	PwdOld string `json:"pwd_old"` // 旧密码
	Pwd    string `json:"pwd"`     // 新密码
}
