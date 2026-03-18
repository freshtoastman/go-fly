package tmpl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 計畫列表頁
func PageCurriculumPlans(c *gin.Context) {
	c.HTML(http.StatusOK, "curriculum_plans.html", gin.H{
		"tab_index": "5-1",
		"action":    "curriculum_plans",
	})
}

// 計畫詳情/編輯頁
func PageCurriculumPlanEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "curriculum_plan_edit.html", gin.H{
		"tab_index": "5-1",
		"action":    "curriculum_plan_edit",
	})
}

// 填報列表頁
func PageCurriculumSubmissions(c *gin.Context) {
	c.HTML(http.StatusOK, "curriculum_submissions.html", gin.H{
		"tab_index": "5-2",
		"action":    "curriculum_submissions",
	})
}

// 填報編輯頁
func PageCurriculumSubmissionEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "curriculum_submission_edit.html", gin.H{
		"tab_index": "5-2",
		"action":    "curriculum_submission_edit",
	})
}

// 審查頁面
func PageCurriculumReview(c *gin.Context) {
	c.HTML(http.StatusOK, "curriculum_review.html", gin.H{
		"tab_index": "5-3",
		"action":    "curriculum_review",
	})
}

// 組織管理頁
func PageOrganizations(c *gin.Context) {
	c.HTML(http.StatusOK, "curriculum_organizations.html", gin.H{
		"tab_index": "5-4",
		"action":    "curriculum_organizations",
	})
}

// 統計頁面
func PageCurriculumStats(c *gin.Context) {
	c.HTML(http.StatusOK, "curriculum_statistics.html", gin.H{
		"tab_index": "5-5",
		"action":    "curriculum_statistics",
	})
}
