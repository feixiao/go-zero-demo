package code

var (

	// [1000:1100)
	ErrInternalServer = NewError(1000, "服务内部错误")

	// User [1100:1200)
	ErrDuplicateUsername  = NewError(1100, "用户已经注册")
	ErrDuplicateMobile    = NewError(1101, "手机号已经被占用")
	ErrUsernameUnRegister = NewError(1102, "用户未注册")
	ErrIncorrectPassword  = NewError(1103, "用户密码错误")
	ErrUserNotFound       = NewError(1104, "用户不存在")
)
