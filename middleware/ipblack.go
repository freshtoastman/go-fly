package middleware

import (
	"github.com/freshtoastman/imaptool/models"
	"github.com/gin-gonic/gin"
)

func Ipblack(c *gin.Context) {
	ip := c.ClientIP()
	ipblack := models.FindIp(ip)
	if ipblack.IP != "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "IP已被加入黑名單",
		})
		c.Abort()
		return
	}
}
