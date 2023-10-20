package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	reps "github.com/lightsaid/hotel-bookings/api/response"
	"github.com/lightsaid/hotel-bookings/pkg/errs"
	"github.com/lightsaid/hotel-bookings/pkg/token"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func RequireAuth(tokenMaker token.TokenMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			reps.FAIL(c, errs.ErrUnauthorized.AsMessage(errs.MsgNotAuthHeader))
			c.Abort()
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			reps.FAIL(c, errs.ErrUnauthorized.AsMessage(errs.MsgNotFormatAuthHeader))
			c.Abort()
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			reps.FAIL(c, errs.ErrUnauthorized.AsMessage(errs.MsgNotFormatAuthHeader))
			c.Abort()
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			if token.IsCustomError(err) {
				reps.FAIL(c, errs.ErrUnauthorized.AsException(err).AsMessage(err.Error()))
				c.Abort()
				return
			}
			reps.FAIL(c, errs.ErrUnauthorized.AsException(err))
			c.Abort()
			return
		}

		c.Set(AuthorizationPayloadKey, payload)
		c.Next()
	}
}
