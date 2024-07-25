package middleware

import "github.com/gin-gonic/gin"

func CrossSite(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//服務器支持的所有跨域請求的方法
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
	//允許跨域設置可以返回其他子段，可以自定义字段
	c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
	// 允許浏覽器（客戶端）可以解析的头部 （重要）
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
	//允許客戶端傳递校驗信息比如 cookie (重要)
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Next()
}
