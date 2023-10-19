package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func NewPayload(userID int64, issuer string, duration time.Duration) (payload *Payload, err error) {
	var tokenID uuid.UUID
	tokenID, err = uuid.NewRandom()
	if err != nil {
		return
	}
	payload = &Payload{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
			Subject:   issuer,
			ID:        tokenID.String(),
			Audience:  []string{issuer},
		},
	}
	return
}
