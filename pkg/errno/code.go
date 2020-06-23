package errno

/*
1						00			02
服务级错误（1为系统级错误）	服务模块代码	具体错误代码

服务级别错误：1 为系统级错误；2 为普通错误，通常是由用户非法操作引起的
服务模块为两位数：一个大型系统的服务模块通常不超过两位数，如果超过，说明这个系统该拆分了
错误码为两位数：防止一个模块定制过多的错误码，后期不好维护
code = 0 说明是正确返回，code > 0 说明是错误返回
错误通常包括系统级错误码和服务级错误码
建议代码中按服务模块将错误分类
错误码均为 >= 0 的数
在本项目中 HTTP Code 固定为 http.StatusOK，错误码通过 code 来表示。
*/

//nolint: golint
var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrParam            = &Errno{Code: 10003, Message: "参数有误"}

	ErrValidation         = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase           = &Errno{Code: 20002, Message: "Database error."}
	ErrToken              = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrInvalidTransaction = &Errno{Code: 20004, Message: "invalid transaction."}

	// user errors
	ErrEncrypt               = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound          = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid          = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect     = &Errno{Code: 20104, Message: "The password was incorrect."}
	ErrAreaCodeEmpty         = &Errno{Code: 20105, Message: "手机区号不能为空"}
	ErrPhoneEmpty            = &Errno{Code: 20106, Message: "手机号不能为空"}
	ErrGenVCode              = &Errno{Code: 20107, Message: "生成验证码错误"}
	ErrSendSMS               = &Errno{Code: 20108, Message: "发送短信错误"}
	ErrSendSMSTooMany        = &Errno{Code: 20109, Message: "已超出当日限制，请明天再试"}
	ErrVerifyCode            = &Errno{Code: 20110, Message: "验证码错误"}
	ErrEmailOrPassword       = &Errno{Code: 20111, Message: "邮箱或密码错误"}
	ErrTwicePasswordNotMatch = &Errno{Code: 20112, Message: "两次密码输入不一致"}
	ErrRegisterFailed        = &Errno{Code: 20113, Message: "注册失败"}
)
