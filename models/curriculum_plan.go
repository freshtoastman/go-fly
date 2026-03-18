package models

import (
	"time"
)

// CurriculumPlan 課程計畫
type CurriculumPlan struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	Name         string     `json:"name" gorm:"size:255;not null"`           // 計畫名稱
	Description  string     `json:"description" gorm:"type:text"`            // 計畫說明
	SchoolYear   string     `json:"school_year" gorm:"size:20;not null"`     // 學年度
	Semester     uint       `json:"semester"`                                // 學期 1=上學期 2=下學期
	PlanType     string     `json:"plan_type" gorm:"size:50"`               // 計畫類型: curriculum/special/experiment
	Status       string     `json:"status" gorm:"size:20;default:'draft'"`   // draft/submitted/reviewing/approved/rejected/returned
	OrgID        uint       `json:"org_id"`                                  // 承辦單位 ID
	CreatedBy    uint       `json:"created_by"`                              // 建立者 ID
	SubmitDeadline *time.Time `json:"submit_deadline"`                       // 填報截止日期
	ReviewDeadline *time.Time `json:"review_deadline"`                       // 審查截止日期
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at"`
}

// PlanSubmission 學校填報紀錄
type PlanSubmission struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	PlanID       uint       `json:"plan_id" gorm:"not null;index"`           // 所屬計畫
	SchoolID     uint       `json:"school_id" gorm:"not null;index"`         // 學校 ID
	SubmitterID  uint       `json:"submitter_id"`                            // 填報人 ID
	Status       string     `json:"status" gorm:"size:20;default:'draft'"`   // draft/submitted/reviewing/approved/rejected/returned
	Content      string     `json:"content" gorm:"type:longtext"`            // 填報內容 (JSON)
	Attachments  string     `json:"attachments" gorm:"type:text"`            // 附件路徑 (JSON array)
	SubmittedAt  *time.Time `json:"submitted_at"`                            // 送出時間
	Remark       string     `json:"remark" gorm:"type:text"`                // 備註
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at"`
}

// PlanReview 審查紀錄
type PlanReview struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	SubmissionID uint       `json:"submission_id" gorm:"not null;index"`     // 填報紀錄 ID
	ReviewerID   uint       `json:"reviewer_id" gorm:"not null;index"`      // 審查者 ID
	ReviewerRole string     `json:"reviewer_role" gorm:"size:20"`           // county/committee/ministry
	Result       string     `json:"result" gorm:"size:20"`                  // approved/rejected/returned
	Score        uint       `json:"score"`                                   // 評分 (0-100)
	Comment      string     `json:"comment" gorm:"type:text"`              // 審查意見
	ReviewedAt   *time.Time `json:"reviewed_at"`                            // 審查時間
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at"`
}

// PlanFormField 計畫表單欄位定義
type PlanFormField struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	PlanID       uint       `json:"plan_id" gorm:"not null;index"`
	FieldName    string     `json:"field_name" gorm:"size:100;not null"`     // 欄位名稱
	FieldType    string     `json:"field_type" gorm:"size:50;not null"`      // text/textarea/select/radio/checkbox/file/date/number
	FieldLabel   string     `json:"field_label" gorm:"size:255"`             // 欄位標籤
	Required     bool       `json:"required" gorm:"default:false"`           // 是否必填
	Options      string     `json:"options" gorm:"type:text"`               // 選項 (JSON array, 用於 select/radio/checkbox)
	SortOrder    uint       `json:"sort_order" gorm:"default:0"`            // 排序
	GroupName    string     `json:"group_name" gorm:"size:100"`             // 分組名稱
	Placeholder  string     `json:"placeholder" gorm:"size:255"`            // 提示文字
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at"`
}

// --- CRUD functions ---

func CreateCurriculumPlan(plan *CurriculumPlan) error {
	return DB.Create(plan).Error
}

func UpdateCurriculumPlan(plan *CurriculumPlan) error {
	return DB.Save(plan).Error
}

func FindCurriculumPlanByID(id uint) (CurriculumPlan, error) {
	var plan CurriculumPlan
	err := DB.Where("id = ?", id).First(&plan).Error
	return plan, err
}

func FindCurriculumPlans(schoolYear string, planType string, status string, page, pageSize int) ([]CurriculumPlan, int) {
	var plans []CurriculumPlan
	var total int
	query := DB.Model(&CurriculumPlan{})
	if schoolYear != "" {
		query = query.Where("school_year = ?", schoolYear)
	}
	if planType != "" {
		query = query.Where("plan_type = ?", planType)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	query.Count(&total)
	if page > 0 && pageSize > 0 {
		query = query.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	query.Order("id desc").Find(&plans)
	return plans, total
}

func DeleteCurriculumPlan(id uint) error {
	return DB.Where("id = ?", id).Delete(&CurriculumPlan{}).Error
}

// --- Submission CRUD ---

func CreatePlanSubmission(sub *PlanSubmission) error {
	return DB.Create(sub).Error
}

func UpdatePlanSubmission(sub *PlanSubmission) error {
	return DB.Save(sub).Error
}

func FindPlanSubmissionByID(id uint) (PlanSubmission, error) {
	var sub PlanSubmission
	err := DB.Where("id = ?", id).First(&sub).Error
	return sub, err
}

func FindPlanSubmissions(planID uint, schoolID uint, status string, page, pageSize int) ([]PlanSubmission, int) {
	var subs []PlanSubmission
	var total int
	query := DB.Model(&PlanSubmission{})
	if planID > 0 {
		query = query.Where("plan_id = ?", planID)
	}
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	query.Count(&total)
	if page > 0 && pageSize > 0 {
		query = query.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	query.Order("id desc").Find(&subs)
	return subs, total
}

// --- Review CRUD ---

func CreatePlanReview(review *PlanReview) error {
	return DB.Create(review).Error
}

func FindPlanReviews(submissionID uint) []PlanReview {
	var reviews []PlanReview
	DB.Where("submission_id = ?", submissionID).Order("id desc").Find(&reviews)
	return reviews
}

// --- Form Field CRUD ---

func CreatePlanFormField(field *PlanFormField) error {
	return DB.Create(field).Error
}

func FindPlanFormFields(planID uint) []PlanFormField {
	var fields []PlanFormField
	DB.Where("plan_id = ?", planID).Order("sort_order asc").Find(&fields)
	return fields
}

func UpdatePlanFormField(field *PlanFormField) error {
	return DB.Save(field).Error
}

func DeletePlanFormField(id uint) error {
	return DB.Where("id = ?", id).Delete(&PlanFormField{}).Error
}
