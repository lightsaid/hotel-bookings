package errs

import (
	"fmt"
	"net/http"
)

// ApiError 一个错误结构体，统一返回错误结构，包含用户错误提示、错误编码和错误信息
type ApiError struct {
	message   string
	exception error
	code      int
}

// 收集错误状态码和错误信息
var codeMaps = map[int]string{}

// 收集 http statusCode
var statusCodeMaps = make(map[int]int)

// NewApiError 往 codeMaps 添加错误，如果 code 已存在会触发panic;
// code 是自定义业务编码; msg 是人类可读错误提示; statusCode 是用于 StatusCode 方法返回的http的状态码，有且仅有一个;
// 主要是用于自定义业务码返回 statusCode
func NewApiError(code int, msg string, statusCode int) *ApiError {
	if _, ok := codeMaps[code]; ok {
		panic(fmt.Sprintf("code(%d) already exist, please replace it", code))
	}

	codeMaps[code] = msg

	statusCodeMaps[code] = statusCode

	return &ApiError{
		code:    code,
		message: msg,
	}
}

// Error 实现 error 接口
func (a *ApiError) Error() string {
	err := fmt.Sprintf("code: %d, message: %s", a.code, a.message)
	if a.exception != nil {
		err = fmt.Sprintf("%s, exception: %v", err, a.exception.Error())
	}
	return err
}

// Unwrap 解开，提供给 errors.Is 和 errors.As 使用
func (a *ApiError) Unwrap() error {
	return a.exception
}

// Code 返回错误码
func (a *ApiError) Code() int {
	return a.code
}

// Message 返回用户可读错误信息
func (a *ApiError) Message() string {
	return a.message
}

// AsMessage 修改消息，例如：错误时入参错误（InvalidParams），可以修改成具体的错误消息（exp:手机号码不正确）
// 返回一个新的 ApiError 指针
func (a *ApiError) AsMessage(msg string) *ApiError {
	return &ApiError{
		code:      a.code,
		message:   msg,
		exception: a.exception,
	}
}

// AsException 添加/追加错误, 返回一个新的 ApiError 指针
func (a *ApiError) AsException(err error, msgs ...string) *ApiError {
	var e error
	if a.exception == nil {
		e = fmt.Errorf("%w", err)
	} else {
		e = fmt.Errorf("%v | %w", a.exception, err)
	}
	newErr := &ApiError{
		code:      a.code,
		message:   a.message,
		exception: e,
	}
	if len(msgs) > 0 {
		newErr.message = msgs[0]
	}
	return newErr
}

// StatusCode 根据ApiError的Code返回 http status code
func (e *ApiError) StatusCode() int {
	if status, ok := statusCodeMaps[e.code]; ok {
		return status
	}

	return http.StatusInternalServerError
}
