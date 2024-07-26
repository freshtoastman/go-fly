package middleware

import (
	"fmt"
	"strings"

	"github.com/freshtoastman/imaptool/models"
	"github.com/gin-gonic/gin"
)

func RbacAuth(c *gin.Context) {
	roleId, _ := c.Get("role_id")
	role := models.FindRole(roleId)
	var flag bool
	rPaths := strings.Split(c.Request.RequestURI, "?")
	uriParam := fmt.Sprintf("%s:%s", c.Request.Method, rPaths[0])
	if role.Method != "*" || role.Path != "*" {
		paths := strings.Split(role.Path, ",")
		for _, p := range paths {
			if uriParam == p {
				flag = true
				break
			}
		}
		if !flag {
			c.JSON(200, gin.H{
				"code": 403,
				"msg":  "沒有權限:" + uriParam,
			})
			c.Abort()
			return
		}
		//methods := strings.Split(role.Method, ",")
		//for _, m := range methods {
		//	if c.Request.Method == m {
		//		methodFlag = true
		//		break
		//	}
		//}
		//if !methodFlag {
		//	c.JSON(200, gin.H{
		//		"code": 403,
		//		"msg":  "沒有權限:" + c.Request.Method + "," + rPaths[0],
		//	})
		//	c.Abort()
		//	return
		//}
	}
	//var flag bool
	//if role.Path != "*" {
	//	paths := strings.Split(role.Path, ",")
	//	for _, p := range paths {
	//		if rPaths[0] == p {
	//			flag = true
	//			break
	//		}
	//	}
	//	if !flag {
	//		c.JSON(200, gin.H{
	//			"code": 403,
	//			"msg":  "沒有權限:" + rPaths[0],
	//		})
	//		c.Abort()
	//		return
	//	}
	//}
}
