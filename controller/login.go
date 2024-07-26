package controller

import (
	"log"
	"net/url"
	"time"

	"github.com/freshtoastman/imaptool/tools"
	"github.com/gin-gonic/gin"
)

// @Summary 登錄驗證接口
// @Produce  json
// @Accept multipart/form-data
// @Param username formData   string true "帳號"
// @Param password formData   string true "密碼"
// @Param type formData   string true "类型"
// @Success 200 {object} controller.Response
// @Failure 200 {object} controller.Response
// @Router /check [post]
// 驗證接口
func LoginCheckPass(c *gin.Context) {
	password := c.PostForm("password")
	username := c.PostForm("username")

	info, uRole, ok := CheckKefuPass(username, password)
	userinfo := make(map[string]interface{})
	if !ok {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "驗證失敗",
		})
		return
	}

	userinfo["name"] = info.Name
	userinfo["kefu_id"] = info.ID
	userinfo["type"] = "kefu"
	if uRole.RoleId != 0 {
		userinfo["role_id"] = uRole.RoleId
	} else {
		userinfo["role_id"] = 2
	}
	userinfo["create_time"] = time.Now().Unix()

	token, _ := tools.MakeToken(userinfo)
	userinfo["ref_token"] = true
	refToken, _ := tools.MakeToken(userinfo)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "驗證成功,正在跳轉",
		"result": gin.H{
			"token":       token,
			"ref_token":   refToken,
			"create_time": userinfo["create_time"],
		},
	})
}

// 遠程請求
func PostBindOfficial(c *gin.Context) {
	api := "https://gofly.v1kf.com/2/officialBindIp"

	phone := c.PostForm("phone")
	password := c.PostForm("password")
	host := c.Request.Host
	data := url.Values{}
	data.Set("phone", phone)
	data.Set("password", password)
	data.Set("host", host)
	res, err := tools.PostForm(api, data)
	if err != nil {
		log.Println("绑定官網帳號發送認證連接錯誤")
	}
	c.Writer.Write([]byte(res))
}
