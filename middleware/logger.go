package middleware

import (
	"time"

	"github.com/freshtoastman/imaptool/tools"
	"github.com/gin-gonic/gin"
)

func NewMidLogger() gin.HandlerFunc {
	logger := tools.Logger()
	return func(c *gin.Context) {
		// 开始時间
		startTime := time.Now()

		// 處理請求
		c.Next()

		// 结束時间
		endTime := time.Now()

		// 执行時间
		latencyTime := endTime.Sub(startTime)

		// 請求方式
		reqMethod := c.Request.Method

		// 請求路由
		reqUri := c.Request.RequestURI

		// 狀態碼
		statusCode := c.Writer.Status()

		// 請求IP
		clientIP := c.ClientIP()

		//日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
