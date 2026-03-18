package models

import (
	"time"
)

// PlatformUser 平台使用者 (擴充原有 User，支援多角色)
type PlatformUser struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	UserID       uint       `json:"user_id" gorm:"not null;index"`           // 關聯 user 表
	OrgID        uint       `json:"org_id" gorm:"not null;index"`            // 所屬組織
	PlatformRole string     `json:"platform_role" gorm:"size:30;not null"`   // school_admin/school_user/county_admin/county_user/committee/ministry/contractor
	Title        string     `json:"title" gorm:"size:100"`                   // 職稱
	Phone        string     `json:"phone" gorm:"size:50"`                    // 聯絡電話
	Email        string     `json:"email" gorm:"size:100"`                   // Email
	IsActive     bool       `json:"is_active" gorm:"default:true"`           // 是否啟用
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at"`
}

// ReviewAssignment 審查委員指派
type ReviewAssignment struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	PlanID       uint       `json:"plan_id" gorm:"not null;index"`           // 計畫 ID
	ReviewerID   uint       `json:"reviewer_id" gorm:"not null;index"`       // 審查者 PlatformUser ID
	SchoolID     uint       `json:"school_id" gorm:"default:0"`             // 指定審查的學校 (0=全部)
	AssignedBy   uint       `json:"assigned_by"`                             // 指派者 ID
	Status       string     `json:"status" gorm:"size:20;default:'pending'"` // pending/accepted/completed
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at"`
}

// --- PlatformUser CRUD ---

func CreatePlatformUser(pu *PlatformUser) error {
	return DB.Create(pu).Error
}

func UpdatePlatformUser(pu *PlatformUser) error {
	return DB.Save(pu).Error
}

func FindPlatformUserByID(id uint) (PlatformUser, error) {
	var pu PlatformUser
	err := DB.Where("id = ?", id).First(&pu).Error
	return pu, err
}

func FindPlatformUserByUserID(userID uint) (PlatformUser, error) {
	var pu PlatformUser
	err := DB.Where("user_id = ?", userID).First(&pu).Error
	return pu, err
}

func FindPlatformUsers(orgID uint, role string, page, pageSize int) ([]PlatformUser, int) {
	var users []PlatformUser
	var total int
	query := DB.Model(&PlatformUser{})
	if orgID > 0 {
		query = query.Where("org_id = ?", orgID)
	}
	if role != "" {
		query = query.Where("platform_role = ?", role)
	}
	query.Count(&total)
	if page > 0 && pageSize > 0 {
		query = query.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	query.Order("id desc").Find(&users)
	return users, total
}

func FindCommitteeMembers() []PlatformUser {
	var users []PlatformUser
	DB.Where("platform_role = ? AND is_active = ?", "committee", true).Find(&users)
	return users
}

// --- ReviewAssignment CRUD ---

func CreateReviewAssignment(ra *ReviewAssignment) error {
	return DB.Create(ra).Error
}

func FindReviewAssignments(planID uint, reviewerID uint) []ReviewAssignment {
	var assignments []ReviewAssignment
	query := DB.Model(&ReviewAssignment{})
	if planID > 0 {
		query = query.Where("plan_id = ?", planID)
	}
	if reviewerID > 0 {
		query = query.Where("reviewer_id = ?", reviewerID)
	}
	query.Order("id desc").Find(&assignments)
	return assignments
}

func UpdateReviewAssignment(ra *ReviewAssignment) error {
	return DB.Save(ra).Error
}
