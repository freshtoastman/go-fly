package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
	"strconv"
)

// GetPlatformUsers 取得平台使用者列表
func GetPlatformUsers(c *gin.Context) {
	orgID, _ := strconv.Atoi(c.DefaultQuery("org_id", "0"))
	role := c.Query("platform_role")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	users, total := models.FindPlatformUsers(uint(orgID), role, page, pageSize)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": users,
		"total":  total,
	})
}

// GetPlatformUser 取得單一平台使用者
func GetPlatformUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	pu, err := models.FindPlatformUserByID(uint(id))
	if err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": "使用者不存在"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "ok", "result": pu})
}

// PostPlatformUser 建立/更新平台使用者
func PostPlatformUser(c *gin.Context) {
	id := c.PostForm("id")
	userID, _ := strconv.Atoi(c.PostForm("user_id"))
	orgID, _ := strconv.Atoi(c.PostForm("org_id"))
	platformRole := c.PostForm("platform_role")
	title := c.PostForm("title")
	phone := c.PostForm("phone")
	email := c.PostForm("email")

	if platformRole == "" {
		c.JSON(200, gin.H{"code": 400, "msg": "角色不可為空"})
		return
	}

	if id == "" {
		pu := &models.PlatformUser{
			UserID:       uint(userID),
			OrgID:        uint(orgID),
			PlatformRole: platformRole,
			Title:        title,
			Phone:        phone,
			Email:        email,
			IsActive:     true,
		}
		if err := models.CreatePlatformUser(pu); err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": "建立使用者失敗"})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "建立成功", "result": pu})
	} else {
		puID, _ := strconv.Atoi(id)
		pu, err := models.FindPlatformUserByID(uint(puID))
		if err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": "使用者不存在"})
			return
		}
		pu.UserID = uint(userID)
		pu.OrgID = uint(orgID)
		pu.PlatformRole = platformRole
		pu.Title = title
		pu.Phone = phone
		pu.Email = email
		if err := models.UpdatePlatformUser(&pu); err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": "更新使用者失敗"})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新成功", "result": pu})
	}
}

// GetCommitteeMembers 取得審查委員列表
func GetCommitteeMembers(c *gin.Context) {
	members := models.FindCommitteeMembers()
	c.JSON(200, gin.H{"code": 200, "msg": "ok", "result": members})
}

// PostReviewAssignment 指派審查委員
func PostReviewAssignment(c *gin.Context) {
	planID, _ := strconv.Atoi(c.PostForm("plan_id"))
	reviewerID, _ := strconv.Atoi(c.PostForm("reviewer_id"))
	schoolID, _ := strconv.Atoi(c.DefaultPostForm("school_id", "0"))
	assignedBy, _ := c.Get("kefu_id")

	ra := &models.ReviewAssignment{
		PlanID:     uint(planID),
		ReviewerID: uint(reviewerID),
		SchoolID:   uint(schoolID),
		AssignedBy: uint(assignedBy.(float64)),
		Status:     "pending",
	}
	if err := models.CreateReviewAssignment(ra); err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": "指派失敗"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "指派成功", "result": ra})
}

// GetReviewAssignments 取得審查指派列表
func GetReviewAssignments(c *gin.Context) {
	planID, _ := strconv.Atoi(c.DefaultQuery("plan_id", "0"))
	reviewerID, _ := strconv.Atoi(c.DefaultQuery("reviewer_id", "0"))

	assignments := models.FindReviewAssignments(uint(planID), uint(reviewerID))
	c.JSON(200, gin.H{"code": 200, "msg": "ok", "result": assignments})
}
