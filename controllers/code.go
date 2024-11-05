package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeInvalidToken
	CodeNeedLogin
)

var codeMsg = map[ResCode]string{
	CodeSuccess:         "成功",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务器繁忙",
	CodeInvalidToken:    "无效Token",
	CodeNeedLogin:       "需要登录",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsg[c]
	if !ok {
		return codeMsg[CodeServerBusy]
	}
	return msg
}
