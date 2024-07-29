package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/common"
	"github.com/taoshihan1991/imaptool/models"
)

func PostIpblack(c *gin.Context) {
	ip := c.PostForm("ip")
	if ip == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "請輸入IP!",
		})
		return
	}
	kefuId, _ := c.Get("kefu_name")
	models.CreateIpblack(ip, kefuId.(string))
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "添加黑名單成功!",
	})
}
func DelIpblack(c *gin.Context) {
	ip := c.Query("ip")
	if ip == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "請輸入IP!",
		})
		return
	}
	models.DeleteIpblackByIp(ip)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "刪除黑名單成功!",
	})
}
func GetIpblacks(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	count := models.CountIps(nil, nil)
	list := models.FindIps(nil, nil, uint(page), common.VisitorPageSize)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result": gin.H{
			"list":     list,
			"count":    count,
			"pagesize": common.PageSize,
		},
	})
}
func GetIpblacksByKefuId(c *gin.Context) {
	kefuId, _ := c.Get("kefu_name")
	list := models.FindIpsByKefuId(kefuId.(string))
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": list,
	})
}
