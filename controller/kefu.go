package controller

import (
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
	"github.com/taoshihan1991/imaptool/ws"
)

func PostKefuAvator(c *gin.Context) {

	avator := c.PostForm("avator")
	if avator == "" {
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "不能為空",
			"result": "",
		})
		return
	}
	kefuName, _ := c.Get("kefu_name")
	models.UpdateUserAvator(kefuName.(string), avator)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
	})
}
func PostKefuPass(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	newPass := c.PostForm("new_pass")
	confirmNewPass := c.PostForm("confirm_new_pass")
	old_pass := c.PostForm("old_pass")
	if newPass != confirmNewPass {
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "密碼不一致",
			"result": "",
		})
		return
	}
	user := models.FindUser(kefuName.(string))
	if user.Password != tools.Md5(old_pass) {
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "舊密碼不正確",
			"result": "",
		})
		return
	}
	models.UpdateUserPass(kefuName.(string), tools.Md5(newPass))
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
	})
}
func PostKefuClient(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	clientId := c.PostForm("client_id")

	if clientId == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "client_id不能為空",
		})
		return
	}
	models.CreateUserClient(kefuName.(string), clientId)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
	})
}
func GetKefuInfo(c *gin.Context) {
	kefuId, _ := c.Get("kefu_id")
	user := models.FindUserById(kefuId)
	info := make(map[string]interface{})
	info["name"] = user.Nickname
	info["id"] = user.Name
	info["avator"] = user.Avator
	info["username"] = user.Name
	info["nickname"] = user.Nickname
	info["uid"] = user.ID
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": info,
	})
}
func GetKefuInfoAll(c *gin.Context) {
	id, _ := c.Get("kefu_id")
	userinfo := models.FindUserRole("user.avator,user.name,user.id, role.name role_name", id)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "驗證成功",
		"result": userinfo,
	})
}
func GetOtherKefuList(c *gin.Context) {
	idStr, _ := c.Get("kefu_id")
	id := idStr.(float64)
	result := make([]interface{}, 0)
	ws.SendPingToKefuClient()
	kefus := models.FindUsers()
	for _, kefu := range kefus {
		if uint(id) == kefu.ID {
			continue
		}

		item := make(map[string]interface{})
		item["name"] = kefu.Name
		item["nickname"] = kefu.Nickname
		item["avator"] = kefu.Avator
		item["status"] = "offline"
		kefu, ok := ws.KefuList[kefu.Name]
		if ok && kefu != nil {
			item["status"] = "online"
		}
		result = append(result, item)
	}
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": result,
	})
}
func PostTransKefu(c *gin.Context) {
	kefuId := c.Query("kefu_id")
	visitorId := c.Query("visitor_id")
	curKefuId, _ := c.Get("kefu_name")
	user := models.FindUser(kefuId)
	visitor := models.FindVisitorByVistorId(visitorId)
	if user.Name == "" || visitor.Name == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "訪客或客服不存在",
		})
		return
	}
	models.UpdateVisitorKefu(visitorId, kefuId)
	ws.UpdateVisitorUser(visitorId, kefuId)
	go ws.VisitorOnline(kefuId, visitor)
	go ws.VisitorOffline(curKefuId.(string), visitor.VisitorId, visitor.Name)
	go ws.VisitorNotice(visitor.VisitorId, "客服轉接到"+user.Nickname)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "轉移成功",
	})
}
func GetKefuInfoSetting(c *gin.Context) {
	kefuId := c.Query("kefu_id")
	user := models.FindUserById(kefuId)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": user,
	})
}
func PostKefuRegister(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	rePassword := c.PostForm("rePassword")
	avator := "/static/images/4.jpg"
	nickname := c.PostForm("nickname")
	captchaCode := c.PostForm("captcha")
	roleId := 1
	if name == "" || password == "" || rePassword == "" || nickname == "" || captchaCode == "" {
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "參數不能為空",
			"result": "",
		})
		return
	}
	if password != rePassword {
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "密碼不一致",
			"result": "",
		})
		return
	}
	oldUser := models.FindUser(name)
	if oldUser.Name != "" {
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "帳號已經存在",
			"result": "",
		})
		return
	}
	session := sessions.Default(c)
	if captchaId := session.Get("captcha"); captchaId != nil {
		session.Delete("captcha")
		_ = session.Save()
		if !captcha.VerifyString(captchaId.(string), captchaCode) {
			c.JSON(200, gin.H{
				"code":   400,
				"msg":    "驗證碼驗證失敗",
				"result": "",
			})
			return
		}
	} else {
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "驗證碼失效",
			"result": "",
		})
		return
	}
	//插入新用戶
	uid := models.CreateUser(name, tools.Md5(password), avator, nickname)
	if uid == 0 {
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "增加用戶失敗",
			"result": "",
		})
		return
	}
	models.CreateUserRole(uid, uint(roleId))

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "註冊完成",
		"result": "",
	})
}
func PostKefuInfo(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	password := c.PostForm("password")
	avator := c.PostForm("avator")
	nickname := c.PostForm("nickname")
	if name == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "客服帳號不能為空",
		})
		return
	}
	//插入新用戶
	if id == "" {
		uid := models.CreateUser(name, tools.Md5(password), avator, nickname)
		if uid == 0 {
			c.JSON(200, gin.H{
				"code":   400,
				"msg":    "增加用戶失敗",
				"result": "",
			})
			return
		}
	} else {
		//更新用戶
		if password != "" {
			password = tools.Md5(password)
		}
		message := &models.Message{
			KefuId: name,
		}
		models.DB.Model(&models.Message{}).Update(message)
		visitor := &models.Visitor{
			ToId: name,
		}
		models.DB.Model(&models.Visitor{}).Update(visitor)
	}

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
	})
}
func GetKefuList(c *gin.Context) {
	users := models.FindUsers()
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "獲得成功",
		"result": users,
	})
}
func DeleteKefuInfo(c *gin.Context) {
	kefuId := c.Query("id")
	models.DeleteUserById(kefuId)
	models.DeleteRoleByUserId(kefuId)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "刪除成功",
		"result": "",
	})
}
