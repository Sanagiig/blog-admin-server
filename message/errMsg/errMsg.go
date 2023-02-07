package errMsg

const (
	ERROR = 500
)

const (
	ERROR_ARGUMENTS_WRONG = iota + 400
	ERROR_AUTH_WRONG
	ERROR_FORBIDDEN = iota + 403
)

// code = 1000  用户模块的错误
const (
	ERROR_USERNAME_USED = iota + 1000
	ERROR_USERNAME_NOT_EXIST
	ERROR_USERNAME_WRONG
	ERROR_PASSWORD_WRONG
	ERROR_USERNAME_PASSWORD_WRONG
	ERROR_TOKEN_EXIST
	ERROR_TOKEN_RUNTIME
	ERROR_TOKEN_TYPE_WRONG
	ERROR_NO_TOKEN
	ERROR_TOKEN_WRONG
	ERROT_CHECKCODE_WRONG
	ERROR_LOGIN_ANOTHER
)

// code = 2000 文章模块错误

// code = 3000 分类模块错误

func InitErrMsg(m map[int]string) {
	m[ERROR] = "操作失败"
	m[ERROR_ARGUMENTS_WRONG] = "参数错误"
	m[ERROR_AUTH_WRONG] = "权限验证错误"
	m[ERROR_FORBIDDEN] = "该账号无权访问，请联系管理员"

	m[ERROR_USERNAME_USED] = "用户名已存在"
	m[ERROR_USERNAME_NOT_EXIST] = "用户名不存在"
	m[ERROR_USERNAME_WRONG] = "用户名错误"
	m[ERROR_PASSWORD_WRONG] = "密码错误"
	m[ERROR_USERNAME_PASSWORD_WRONG] = "用户名或密码错误"
	m[ERROR_TOKEN_EXIST] = "Token 已存在"
	m[ERROR_TOKEN_RUNTIME] = "Token 已过期"
	m[ERROR_TOKEN_TYPE_WRONG] = "Token 格式错误"
	m[ERROR_NO_TOKEN] = "缺少token,请先登录"
	m[ERROR_TOKEN_WRONG] = "Token 错误"
	m[ERROT_CHECKCODE_WRONG] = "校验码错误"
	m[ERROR_LOGIN_ANOTHER] = "该账号在其他设备登录了，请重新登录并检查您的账号是否泄露"
}
