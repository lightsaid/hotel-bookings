package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetFromContext[T any](c *gin.Context, key string) T {
	value, ok := c.Get(key)
	if !ok {
		panic(fmt.Sprintf("获取 %s 失败", key))
	}
	return value.(T)
}
