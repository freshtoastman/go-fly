package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
	"strconv"
)

// GetCurriculumStatistics 取得課程計畫統計資訊
func GetCurriculumStatistics(c *gin.Context) {
	planID, _ := strconv.Atoi(c.DefaultQuery("plan_id", "0"))
	county := c.Query("county")

	stats := make(map[string]interface{})

	// 計畫總數
	plans, planTotal := models.FindCurriculumPlans("", "", "", 0, 0)
	stats["total_plans"] = planTotal
	_ = plans

	if planID > 0 {
		// 該計畫填報統計
		allSubs, totalSubs := models.FindPlanSubmissions(uint(planID), 0, "", 0, 0)
		_ = allSubs
		_, submitted := models.FindPlanSubmissions(uint(planID), 0, "submitted", 0, 0)
		_, approved := models.FindPlanSubmissions(uint(planID), 0, "approved", 0, 0)
		_, rejected := models.FindPlanSubmissions(uint(planID), 0, "rejected", 0, 0)
		_, returned := models.FindPlanSubmissions(uint(planID), 0, "returned", 0, 0)
		_, draft := models.FindPlanSubmissions(uint(planID), 0, "draft", 0, 0)

		stats["plan_id"] = planID
		stats["total_submissions"] = totalSubs
		stats["submitted"] = submitted
		stats["approved"] = approved
		stats["rejected"] = rejected
		stats["returned"] = returned
		stats["draft"] = draft
	}

	if county != "" {
		// 縣市學校統計
		schools := models.FindSchoolsByCounty(county)
		stats["county"] = county
		stats["total_schools"] = len(schools)
	}

	// 組織統計
	_, totalSchools := models.FindOrganizations("school", "", 0, 0, 0)
	_, totalCounties := models.FindOrganizations("county", "", 0, 0, 0)
	stats["total_schools"] = totalSchools
	stats["total_counties"] = totalCounties

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": stats,
	})
}
