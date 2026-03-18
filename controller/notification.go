package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
	"strconv"
)

// GetNotifications 取得通知列表
func GetNotifications(c *gin.Context) {
	userID, _ := c.Get("kefu_id")
	isReadStr := c.Query("is_read")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	var isRead *bool
	if isReadStr != "" {
		b := isReadStr == "true"
		isRead = &b
	}

	notifs, total := models.FindNotifications(uint(userID.(float64)), isRead, page, pageSize)
	unread := models.CountUnreadNotifications(uint(userID.(float64)))
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": notifs,
		"total":  total,
		"unread": unread,
	})
}

// PostNotificationRead 標記通知已讀
func PostNotificationRead(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	if id > 0 {
		models.MarkNotificationRead(uint(id))
	} else {
		userID, _ := c.Get("kefu_id")
		models.MarkAllNotificationsRead(uint(userID.(float64)))
	}
	c.JSON(200, gin.H{"code": 200, "msg": "ok"})
}

// GetAuditLogs 取得操作日誌
func GetAuditLogs(c *gin.Context) {
	userID, _ := strconv.Atoi(c.DefaultQuery("user_id", "0"))
	action := c.Query("action")
	targetType := c.Query("target_type")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	logs, total := models.FindAuditLogs(uint(userID), action, targetType, page, pageSize)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": logs,
		"total":  total,
	})
}
