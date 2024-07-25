package controller

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
)

func GetCheckWeixinSign(c *gin.Context) {
	token := models.FindConfig("WeixinToken")
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")
	//將token、timestamp、nonce三个參數進行字典序排序
	var tempArray = []string{token, timestamp, nonce}
	sort.Strings(tempArray)
	//將三个參數字符串拼接成一个字符串進行sha1加密
	var sha1String string = ""
	for _, v := range tempArray {
		sha1String += v
	}
	h := sha1.New()
	h.Write([]byte(sha1String))
	sha1String = hex.EncodeToString(h.Sum([]byte("")))
	//获得加密后的字符串可与signature對比
	if sha1String == signature {
		c.Writer.Write([]byte(echostr))
	} else {
		log.Println("微信API驗證失敗")
	}
}
