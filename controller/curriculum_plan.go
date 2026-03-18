package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
	"strconv"
	"time"
)

// GetCurriculumPlans 取得計畫列表
func GetCurriculumPlans(c *gin.Context) {
	schoolYear := c.Query("school_year")
	planType := c.Query("plan_type")
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	plans, total := models.FindCurriculumPlans(schoolYear, planType, status, page, pageSize)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": plans,
		"total":  total,
	})
}

// GetCurriculumPlan 取得單一計畫
func GetCurriculumPlan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	plan, err := models.FindCurriculumPlanByID(uint(id))
	if err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": "計畫不存在"})
		return
	}
	fields := models.FindPlanFormFields(uint(id))
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": plan,
		"fields": fields,
	})
}

// PostCurriculumPlan 建立/更新計畫
func PostCurriculumPlan(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	description := c.PostForm("description")
	schoolYear := c.PostForm("school_year")
	semester, _ := strconv.Atoi(c.DefaultPostForm("semester", "1"))
	planType := c.PostForm("plan_type")
	orgID, _ := strconv.Atoi(c.DefaultPostForm("org_id", "0"))
	submitDeadlineStr := c.PostForm("submit_deadline")
	reviewDeadlineStr := c.PostForm("review_deadline")

	if name == "" || schoolYear == "" {
		c.JSON(200, gin.H{"code": 400, "msg": "計畫名稱與學年度不可為空"})
		return
	}

	createdBy, _ := c.Get("kefu_id")

	var submitDeadline, reviewDeadline *time.Time
	if submitDeadlineStr != "" {
		t, err := time.Parse("2006-01-02", submitDeadlineStr)
		if err == nil {
			submitDeadline = &t
		}
	}
	if reviewDeadlineStr != "" {
		t, err := time.Parse("2006-01-02", reviewDeadlineStr)
		if err == nil {
			reviewDeadline = &t
		}
	}

	if id == "" {
		plan := &models.CurriculumPlan{
			Name:           name,
			Description:    description,
			SchoolYear:     schoolYear,
			Semester:       uint(semester),
			PlanType:       planType,
			OrgID:          uint(orgID),
			CreatedBy:      uint(createdBy.(float64)),
			SubmitDeadline: submitDeadline,
			ReviewDeadline: reviewDeadline,
			Status:         "draft",
		}
		if err := models.CreateCurriculumPlan(plan); err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": "建立計畫失敗"})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "建立成功", "result": plan})
	} else {
		planID, _ := strconv.Atoi(id)
		plan, err := models.FindCurriculumPlanByID(uint(planID))
		if err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": "計畫不存在"})
			return
		}
		plan.Name = name
		plan.Description = description
		plan.SchoolYear = schoolYear
		plan.Semester = uint(semester)
		plan.PlanType = planType
		plan.OrgID = uint(orgID)
		plan.SubmitDeadline = submitDeadline
		plan.ReviewDeadline = reviewDeadline
		if err := models.UpdateCurriculumPlan(&plan); err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": "更新計畫失敗"})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新成功", "result": plan})
	}
}

// DeleteCurriculumPlan 刪除計畫
func DeleteCurriculumPlan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if err := models.DeleteCurriculumPlan(uint(id)); err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": "刪除失敗"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "刪除成功"})
}

// PostPlanFormField 建立/更新表單欄位
func PostPlanFormField(c *gin.Context) {
	id := c.PostForm("id")
	planID, _ := strconv.Atoi(c.PostForm("plan_id"))
	fieldName := c.PostForm("field_name")
	fieldType := c.PostForm("field_type")
	fieldLabel := c.PostForm("field_label")
	required := c.PostForm("required") == "true"
	options := c.PostForm("options")
	sortOrder, _ := strconv.Atoi(c.DefaultPostForm("sort_order", "0"))
	groupName := c.PostForm("group_name")
	placeholder := c.PostForm("placeholder")

	if fieldName == "" || fieldType == "" {
		c.JSON(200, gin.H{"code": 400, "msg": "欄位名稱與類型不可為空"})
		return
	}

	if id == "" {
		field := &models.PlanFormField{
			PlanID:      uint(planID),
			FieldName:   fieldName,
			FieldType:   fieldType,
			FieldLabel:  fieldLabel,
			Required:    required,
			Options:     options,
			SortOrder:   uint(sortOrder),
			GroupName:   groupName,
			Placeholder: placeholder,
		}
		models.CreatePlanFormField(field)
		c.JSON(200, gin.H{"code": 200, "msg": "建立成功", "result": field})
	} else {
		fieldID, _ := strconv.Atoi(id)
		field := &models.PlanFormField{
			ID:          uint(fieldID),
			PlanID:      uint(planID),
			FieldName:   fieldName,
			FieldType:   fieldType,
			FieldLabel:  fieldLabel,
			Required:    required,
			Options:     options,
			SortOrder:   uint(sortOrder),
			GroupName:   groupName,
			Placeholder: placeholder,
		}
		models.UpdatePlanFormField(field)
		c.JSON(200, gin.H{"code": 200, "msg": "更新成功", "result": field})
	}
}

// DeletePlanFormField 刪除表單欄位
func DeletePlanFormField(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	models.DeletePlanFormField(uint(id))
	c.JSON(200, gin.H{"code": 200, "msg": "刪除成功"})
}

// GetPlanFormFields 取得計畫表單欄位
func GetPlanFormFields(c *gin.Context) {
	planID, _ := strconv.Atoi(c.Query("plan_id"))
	fields := models.FindPlanFormFields(uint(planID))
	c.JSON(200, gin.H{"code": 200, "msg": "ok", "result": fields})
}
