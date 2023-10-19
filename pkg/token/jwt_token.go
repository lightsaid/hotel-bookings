package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 16

type JWTMaker struct {
	secretKey string // 密钥
	Issuer    string // 站点域名
}

func NewJWTMaker(secretKey string, issuer string) (TokenMaker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, errors.New("密钥必须最小长度为16")
	}
	maker := &JWTMaker{
		secretKey: secretKey,
		Issuer:    issuer,
	}

	return maker, nil
}

func (maker *JWTMaker) CreateToken(userID int64, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userID, maker.Issuer, duration)
	if err != nil {
		return "", nil, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	accessToken, err := jwtToken.SignedString([]byte(maker.secretKey))
	return accessToken, payload, err
}

func (maker *JWTMaker) VerifyToken(tokenString string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, keyFunc)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	payload, ok := token.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
