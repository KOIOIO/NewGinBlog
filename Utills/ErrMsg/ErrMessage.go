package ErrMsg

const (
	SUCCESS       = 200
	ERROR         = 500
	ERROR_NO_DATA = 501
	//code=1000...用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	//code=2000...文章模块的错误
	ERROR_CATENAME_USED = 2001
	//code=3000...分类模块的错误
	ERROR_ART_NOT_EXIST  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

var ErrMsg = map[int]string{
	SUCCESS:       "OK",
	ERROR:         "FAIL",
	ERROR_NO_DATA: "没用对应数据",

	ERROR_USERNAME_USED:    "用户已存在！",
	ERROR_PASSWORD_WRONG:   "密码错误！",
	ERROR_USER_NOT_EXIST:   "用户不存在！",
	ERROR_TOKEN_EXIST:      "Token不存在！",
	ERROR_TOKEN_RUNTIME:    "Token已过期！",
	ERROR_TOKEN_WRONG:      "Token不正确！",
	ERROR_TOKEN_TYPE_WRONG: "Token格式错误！",
	ERROR_USER_NO_RIGHT:    "该用户无权限",
	ERROR_ART_NOT_EXIST:    "文章已经不存在了",
	ERROR_CATE_NOT_EXIST:   "分类不存在",
	ERROR_CATENAME_USED:    "分类已存在",
}

func GetErrMessage(code int) string { return ErrMsg[code] }
