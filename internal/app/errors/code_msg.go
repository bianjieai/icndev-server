package errors

const (
	EC40000 = 40000 // 错误的请求参数

	EC40001 = 40001 // 未认证

	EC40002 = 40002 // 参数转化异常

	EC40003 = 40003 // 记录已存在

	EC40004 = 40004 // 记录未找到

	EC40005 = 40005 // 操作被拒绝

	ErrRecordExists        = 40003 // 记录已存在
	ErrRecordNotFound      = 40004 // 记录未找到
	ErrOperationRejection  = 40005 // 记录未找到
	EC50001                = 50001 // 服务列表失败
	EC50002                = 50002 // 服务交易列表失败
	ErrAuthForbidden       = 50003 // 权限错误
	ErrPermissionForbidden = 50004 //接口无权限

	ErrUsernameOrPassword = 60001
	ErrRestPassword       = 60002
	ErrAccountDisabled    = 60003
	ErrUnknownSsoCode     = 60004
	ErrSsoLogout          = 60005
	ErrAppDisabled        = 60006
	ErrSsoAppStatus       = 60007
	ErrCookieInvalid      = 60008
	ErrOrgDisabled        = 60009
	ErrSsoSystem          = 60010

	ERROR           = 500
	ErrConnectChain = 60011

	ErrInvalidParams       = 40000
	ErrNotCertification    = 40001
	ErrParameterConversion = 40002
)
