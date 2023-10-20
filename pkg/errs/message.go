package errs

var (
	MsgHotelCodeExists        = "酒店编码已存在"
	MsgHotelRoomNumberExists  = "同一个酒店下，房号已存在"
	MsgErrorPhoneNumber       = "手机号码格式不正确"
	MsgSMSLoginNotImplemented = "验证码登录未实现"
	MsgMismatchedPaswd        = "账号或者密码不匹配"
	MsgCreateTokenFailed      = "登录时创建Token失败"
	MsgNotAuthHeader          = "请先登录"
	MsgNotFormatAuthHeader    = "Token 格式错误"
	MsgCreateSessionFail      = "登录时，创建session错误"
	MsgRefreshTokenNoMis      = "Token 不匹配"
	MsgUserNotMis             = "用户不匹配"
)
