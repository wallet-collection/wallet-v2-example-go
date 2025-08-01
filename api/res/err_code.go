package res

import "fmt"

// nolint: golint
var (
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 1, Message: "网络错误，请稍后"}
	ErrToken            = &Errno{Code: 10001, Message: "登录失效"}
	ErrParam            = &Errno{Code: 10002, Message: "参数错误"}
	ErrNotData          = &Errno{Code: 10003, Message: "未找到记录"}
	ErrNotChangeData    = &Errno{Code: 10004, Message: "数据没有变化"}
	ErrNotRepeatData    = &Errno{Code: 10005, Message: "当前邮箱已被注册"}
	ErrLoginNot         = &Errno{Code: 10006, Message: "用户名/密码错误"}
	ErrOldPasswordNot   = &Errno{Code: 10007, Message: "旧密码错误"}
	ErrPasswordNot      = &Errno{Code: 10008, Message: "密码错误"}
	ErrInviteNot        = &Errno{Code: 10009, Message: "当前邀请码无效"}
	ErrCodeErr          = &Errno{Code: 10010, Message: "验证码错误"}
	ErrBalanceNot       = &Errno{Code: 10011, Message: "余额不足"}
	ErrNetworkNot       = &Errno{Code: 10012, Message: "网络暂停"}
	ErrMinWithdraw      = &Errno{Code: 10013, Message: "最低提现"}
	ErrMaxWithdraw      = &Errno{Code: 10014, Message: "最大提现"}
	ErrMinTransfer      = &Errno{Code: 10015, Message: "最低划转"}
	ErrMaxTransfer      = &Errno{Code: 10016, Message: "最大划转"}
	ErrTransferNot      = &Errno{Code: 10017, Message: "该币不可划转"}
)

// Errno ...
type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

// Err represents an error
type Err struct {
	Code    int
	Message string
	Err     error
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// DecodeErr ...
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
