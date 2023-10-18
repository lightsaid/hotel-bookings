package reps

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lightsaid/hotel-bookings/api/errs"
)

// Response 通用数据响应结构
type Response struct {
	Result any    `json:"result"`
	Msg    string `json:"msg"`
	Code   int    `json:"code"`
}

func ToResponse(err *errs.ApiError, result any) *Response {
	return &Response{
		Result: result,
		Msg:    err.Message(),
		Code:   err.Code(),
	}
}

// ListResponse 通用分页数据响应结构
type ListResponse struct {
	List any      `json:"list"`
	Meta Metadata `json:"meta"`
	Msg  string   `json:"msg"`
	Code int      `json:"code"`
}

func ToListResponse(list any, meteData Metadata, err *errs.ApiError) *ListResponse {
	return &ListResponse{
		List: list,
		Meta: meteData,
		Msg:  err.Message(),
		Code: err.Code(),
	}
}

// Responder 是为了约束下面 JSON 方法传参
type Responder interface {
	Response | ListResponse
}

func JSON[T Responder](c *gin.Context, err *errs.ApiError, data *T) {
	if err.StatusCode() != http.StatusOK {
		// TODO: log
		log.Println(err)
	}

	c.JSON(err.StatusCode(), data)
}

func OK(c *gin.Context, data any) {
	JSON(c, errs.ErrOK, ToResponse(errs.ErrOK, data))
}

func FAIL(c *gin.Context, err *errs.ApiError) {
	data := ToResponse(err, nil)
	JSON(c, err, data)
}

func PAGE(c *gin.Context, list any, total int64, pageNum, pageSize int32) {
	meta := CalculateMetadata(total, pageNum, pageSize)
	result := ToListResponse(list, meta, errs.ErrOK)
	JSON(c, errs.ErrOK, result)
}
