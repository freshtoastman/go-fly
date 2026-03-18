package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
	"strconv"
)

// GetOrganizations 取得組織列表
func GetOrganizations(c *gin.Context) {
	orgType := c.Query("org_type")
	county := c.Query("county")
	parentID, _ := strconv.Atoi(c.DefaultQuery("parent_id", "0"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	orgs, total := models.FindOrganizations(orgType, county, uint(parentID), page, pageSize)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": orgs,
		"total":  total,
	})
}

// GetOrganization 取得單一組織
func GetOrganization(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	org, err := models.FindOrganizationByID(uint(id))
	if err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": "組織不存在"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "ok", "result": org})
}

// PostOrganization 建立/更新組織
func PostOrganization(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	code := c.PostForm("code")
	orgType := c.PostForm("org_type")
	parentID, _ := strconv.Atoi(c.DefaultPostForm("parent_id", "0"))
	county := c.PostForm("county")
	district := c.PostForm("district")
	level := c.PostForm("level")
	address := c.PostForm("address")
	phone := c.PostForm("phone")
	principal := c.PostForm("principal")

	if name == "" || orgType == "" {
		c.JSON(200, gin.H{"code": 400, "msg": "組織名稱與類型不可為空"})
		return
	}

	if id == "" {
		org := &models.Organization{
			Name:      name,
			Code:      code,
			OrgType:   orgType,
			ParentID:  uint(parentID),
			County:    county,
			District:  district,
			Level:     level,
			Address:   address,
			Phone:     phone,
			Principal: principal,
			IsActive:  true,
		}
		if err := models.CreateOrganization(org); err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": "建立組織失敗"})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "建立成功", "result": org})
	} else {
		orgID, _ := strconv.Atoi(id)
		org, err := models.FindOrganizationByID(uint(orgID))
		if err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": "組織不存在"})
			return
		}
		org.Name = name
		org.Code = code
		org.OrgType = orgType
		org.ParentID = uint(parentID)
		org.County = county
		org.District = district
		org.Level = level
		org.Address = address
		org.Phone = phone
		org.Principal = principal
		if err := models.UpdateOrganization(&org); err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": "更新組織失敗"})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新成功", "result": org})
	}
}

// DeleteOrganization 刪除組織
func DeleteOrganization(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if err := models.DeleteOrganization(uint(id)); err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": "刪除失敗"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "刪除成功"})
}

// GetSchoolsByCounty 依縣市取得學校列表
func GetSchoolsByCounty(c *gin.Context) {
	county := c.Query("county")
	schools := models.FindSchoolsByCounty(county)
	c.JSON(200, gin.H{"code": 200, "msg": "ok", "result": schools})
}
