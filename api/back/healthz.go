package back

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthZ struct{}

func (*HealthZ) HealthZ(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusText(http.StatusOK)})
}
