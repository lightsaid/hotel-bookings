package token

import "errors"

var (
	ErrInvalidToken = errors.New("令牌无效")
	ErrExpiredToken = errors.New("令牌已过期")
)

func IsCustomError(err error) bool {
	return errors.Is(err, ErrInvalidToken) || errors.Is(err, ErrExpiredToken)
}
