package errors

import (
	"fmt"
)

type Error interface {
	Error() string
	Code() int
	Msg() string
}

// Wrap system error
func Wrap(err error) Error {
	return stsErr{
		code: EC40005,
		msg:  err.Error(),
	}
}

// WrapDetail with code and msg
func WrapDetail(code int, msg string) Error {
	return stsErr{
		code: code,
		msg:  msg,
	}
}

func WrapBadRequest(err error) Error {
	return stsErr{
		code: EC40000,
		msg:  err.Error(),
	}
}

type stsErr struct {
	code int
	msg  string
}

func (e stsErr) Error() string {
	return fmt.Sprintf("err_code: %d, err_msg: %s", e.code, e.msg)
}

func (e stsErr) Code() int {
	return e.code
}

func (e stsErr) Msg() string {
	return e.msg
}

var MsgFlags = map[int]string{
	ERROR:                  "failed",
	ErrInvalidParams:       "Request parameter error",
	ErrNotCertification:    "Not certified",
	ErrParameterConversion: "Abnormal parameter conversion",

	ErrUsernameOrPassword: "The login email or password is incorrect, please try again",
	ErrRestPassword:       "Account password is not reset",
	ErrAccountDisabled:    "Your account is disabled, if you have any questions, please contact the administrator",
	ErrUnknownSsoCode:     "Unknown sso error code",
	ErrSsoLogout:          "Invoke sso logout Failed",
	ErrAppDisabled:        "app is disabled",
	ErrSsoAppStatus:       "Invoke sso app status failed",
	ErrCookieInvalid:      "Cookie invalid",
	ErrOrgDisabled:        "org is disabled",
	ErrSsoSystem:          "sso system err",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

var ErrEsNotFound = fmt.Errorf("not foud from es")

type vsErr struct {
	code int
	msg  string
}

func (e vsErr) Error() string {
	return fmt.Sprintf("err_code: %d, err_msg: %s", e.code, e.msg)
}

func (e vsErr) Code() int {
	return e.code
}

func (e vsErr) Msg() string {
	return e.msg
}
