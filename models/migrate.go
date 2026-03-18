package models

// AutoMigrateCurriculum 自動建立課程計畫填報系統相關資料表
func AutoMigrateCurriculum() {
	DB.AutoMigrate(
		&CurriculumPlan{},
		&PlanSubmission{},
		&PlanReview{},
		&PlanFormField{},
		&Organization{},
		&PlatformUser{},
		&ReviewAssignment{},
		&Notification{},
		&AuditLog{},
	)
}
