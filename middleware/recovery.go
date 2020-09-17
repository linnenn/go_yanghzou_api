package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
//恢复方法中间件，报错到500
func GinRecovery(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
