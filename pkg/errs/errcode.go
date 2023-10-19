package errs

import "net/http"

/**
 * 此文件定义系统的所有错误码, 包括公共错误码和具体业务错误码
 */

var (
	// 定义公共错误码, 以 1xxx0 开头
	ErrOK                    = NewApiError(10000, "请求成功", http.StatusOK)                     // 200
	ErrRecordExists          = NewApiError(10001, "记录已存在", http.StatusConflict)              // 409
	ErrBadRequest            = NewApiError(10004, "入参错误", http.StatusBadRequest)             // 400
	ErrUnauthorized          = NewApiError(10005, "验证失败", http.StatusUnauthorized)           // 401
	ErrForbidden             = NewApiError(10006, "未经授权", http.StatusForbidden)              // 403
	ErrNotFound              = NewApiError(10007, "未找到", http.StatusNotFound)                // 404
	ErrMethodNotAllowed      = NewApiError(10008, "请求方法不支持", http.StatusMethodNotAllowed)    // 405
	ErrNotAcceptable         = NewApiError(10009, "请求头无效", http.StatusNotAcceptable)         // 406
	ErrRequestTimeout        = NewApiError(10010, "请求超时", http.StatusRequestTimeout)         // 408
	ErrRequestEntityTooLarge = NewApiError(10011, "请求体过大", http.StatusRequestEntityTooLarge) // 413
	ErrUnprocessableEntity   = NewApiError(10012, "实体错误", http.StatusUnprocessableEntity)    // 422
	ErrTooManyRequests       = NewApiError(10013, "请求繁忙", http.StatusTooManyRequests)        // 429
	ErrServerError           = NewApiError(10014, "务器内部错误", http.StatusInternalServerError)  // 500
	ErrBadGateway            = NewApiError(10015, "网关错误", http.StatusBadGateway)             // 502
	ErrServiceUnavailable    = NewApiError(10016, "无法处理请求", http.StatusServiceUnavailable)   // 503
	ErrGatewayTimeout        = NewApiError(10017, "服务器未响应", http.StatusGatewayTimeout)       // 504

	// 业务状态

)
