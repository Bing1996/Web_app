package controller

type ResponseCode int

const (
	CodeSuccess ResponseCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeInvalidToken
	CodeEmptyAuthorizationHeader
	CodeInvalidTokenFormat

	CodeCtxCurrentUserNotFound
)

var codeMsgMap = map[ResponseCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",

	CodeInvalidToken:             "无效的Token",
	CodeEmptyAuthorizationHeader: "请求头中Authoritarian格式为空",
	CodeInvalidTokenFormat:       "请求头中Authoritarian格式有误",

	CodeCtxCurrentUserNotFound: "context中无用户信息",
}

func (r ResponseCode) Msg() string {
	msg, ok := codeMsgMap[r]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
