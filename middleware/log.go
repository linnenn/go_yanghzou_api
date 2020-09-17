package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
//使用zap仓库，设置存储日志的相关配置，相关错误存储到日志文件
func GinLog(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context ) {

	}
}