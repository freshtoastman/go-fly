package controller

import (
	"github.com/freshtoastman/imaptool/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	jump := models.FindConfig("JumpLang")
	if jump != "cn" {
		jump = "en"
	}
	c.Redirect(302, "/index_"+jump)
}
