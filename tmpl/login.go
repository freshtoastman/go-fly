package tmpl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/tools"
)

// 登錄界面
func PageLogin(c *gin.Context) {
	if noExist, _ := tools.IsFileNotExist("./install.lock"); noExist {
		c.Redirect(302, "/install")
	}
	c.HTML(http.StatusOK, "login.html", nil)
}

// 绑定界面
func PageBind(c *gin.Context) {
	c.HTML(http.StatusOK, "bind.html", gin.H{})
}
