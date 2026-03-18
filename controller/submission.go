package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
	"strconv"
	"time"
)

// GetPlanSubmissions 取得填報列表
func GetPlanSubmissions(c *gin.Context) {
	planID, _ := strconv.Atoi(c.DefaultQuery("plan_id", "0"))
	schoolID, _ := strconv.Atoi(c.DefaultQuery("school_id", "0"))
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	subs, total := models.FindPlanSubmissions(uint(planID), uint(schoolID), status, page, pageSize)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": subs,
		"total":  total,
	})
}

// GetPlanSubmission 取得單一填報紀錄
func GetPlanSubmission(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	sub, err := models.FindPlanSubmissionByID(uint(id))
	if err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": "填報紀錄不存在"})
		return
	}
	reviews := models.FindPlanReviews(uint(id))
	c.JSON(200, gin.H{
		"code":    200,
		"msg":     "ok",
		"result":  sub,
		"reviews": reviews,
	})
}

// PostPlanSubmission 建立/更新填報
func PostPlanSubmission(c *gin.Context) {
	id := c.PostForm("id")
	planID, _ := strconv.Atoi(c.PostForm("plan_id"))
	schoolID, _ := strconv.Atoi(c.PostForm("school_id"))
	content := c.PostForm("content")
	attachments := c.PostForm("attachments")
	remark := c.PostForm("remark")

	submitterID, _ := c.Get("kefu_id")

	if id == "" {
		sub := &models.PlanSubmission{
			PlanID:      uint(planID),
			SchoolID:    uint(schoolID),
			SubmitterID: uint(submitterID.(float64)),
			Content:     content,
			Attachments: attachments,
			Remark:      remark,
			Status:      "draft",
		}
		if err := models.CreatePlanSubmission(sub); err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": "建立填報失敗"})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "建立成功", "result": sub})
	} else {
		subID, _ := strconv.Atoi(id)
		sub, err := models.FindPlanSubmissionByID(uint(subID))
		if err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": "填報紀錄不存在"})
			return
		}
		sub.Content = content
		sub.Attachments = attachments
		sub.Remark = remark
		if err := models.UpdatePlanSubmission(&sub); err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": "更新填報失敗"})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新成功", "result": sub})
	}
}

// PostSubmitPlan 送出填報 (變更狀態為 submitted)
func PostSubmitPlan(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	sub, err := models.FindPlanSubmissionByID(uint(id))
	if err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": "填報紀錄不存在"})
		return
	}
	if sub.Status != "draft" && sub.Status != "returned" {
		c.JSON(200, gin.H{"code": 400, "msg": "目前狀態無法送出"})
		return
	}
	now := time.Now()
	sub.Status = "submitted"
	sub.SubmittedAt = &now
	models.UpdatePlanSubmission(&sub)
	c.JSON(200, gin.H{"code": 200, "msg": "送出成功"})
}

// PostPlanReview 提交審查結果
func PostPlanReview(c *gin.Context) {
	submissionID, _ := strconv.Atoi(c.PostForm("submission_id"))
	result := c.PostForm("result")
	score, _ := strconv.Atoi(c.DefaultPostForm("score", "0"))
	comment := c.PostForm("comment")
	reviewerRole := c.PostForm("reviewer_role")

	reviewerID, _ := c.Get("kefu_id")

	if result == "" {
		c.JSON(200, gin.H{"code": 400, "msg": "審查結果不可為空"})
		return
	}

	now := time.Now()
	review := &models.PlanReview{
		SubmissionID: uint(submissionID),
		ReviewerID:   uint(reviewerID.(float64)),
		ReviewerRole: reviewerRole,
		Result:       result,
		Score:        uint(score),
		Comment:      comment,
		ReviewedAt:   &now,
	}
	if err := models.CreatePlanReview(review); err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": "提交審查失敗"})
		return
	}

	// 更新填報狀態
	sub, _ := models.FindPlanSubmissionByID(uint(submissionID))
	sub.Status = result // approved/rejected/returned
	models.UpdatePlanSubmission(&sub)

	c.JSON(200, gin.H{"code": 200, "msg": "審查完成", "result": review})
}

// GetPlanReviews 取得填報的審查紀錄
func GetPlanReviews(c *gin.Context) {
	submissionID, _ := strconv.Atoi(c.Query("submission_id"))
	reviews := models.FindPlanReviews(uint(submissionID))
	c.JSON(200, gin.H{"code": 200, "msg": "ok", "result": reviews})
}
