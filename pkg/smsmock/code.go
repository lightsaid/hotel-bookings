package smsmock

import (
	"errors"
	"time"
)

var ErrInNotExists = errors.New("验证码不存在")
var ErrInValid = errors.New("验证码无效")
var ErrInExpired = errors.New("验证码过期")

type StatusType int8

const (
	StatusNormal  = 0
	StatusInValid = -1 // 验证码无效， 被使用过
	StatusExpired = -2 // 验证码过期
)

// SMSCode 一个短信验证码结构
type SMSCode struct {
	PhoneNumber string
	Code        string
	Status      StatusType
	ExpiresAt   time.Time
}
