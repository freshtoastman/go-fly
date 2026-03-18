package models

import (
	"time"
)

// Notification 系統通知
type Notification struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	UserID     uint       `json:"user_id" gorm:"not null;index"`           // 接收者 User ID
	Title      string     `json:"title" gorm:"size:255;not null"`          // 通知標題
	Content    string     `json:"content" gorm:"type:text"`                // 通知內容
	NotiType   string     `json:"noti_type" gorm:"size:30"`               // system/plan/review/deadline/announcement
	RefType    string     `json:"ref_type" gorm:"size:30"`                // 關聯類型: plan/submission/review
	RefID      uint       `json:"ref_id" gorm:"default:0"`                // 關聯 ID
	IsRead     bool       `json:"is_read" gorm:"default:false"`           // 是否已讀
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index" json:"deleted_at"`
}

// AuditLog 操作日誌
type AuditLog struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	UserID     uint       `json:"user_id" gorm:"not null;index"`
	Action     string     `json:"action" gorm:"size:50;not null"`          // create/update/delete/submit/review/approve/reject
	TargetType string     `json:"target_type" gorm:"size:50"`             // plan/submission/review/organization/user
	TargetID   uint       `json:"target_id"`
	Detail     string     `json:"detail" gorm:"type:text"`                // 操作詳情 (JSON)
	IPAddress  string     `json:"ip_address" gorm:"size:50"`
	CreatedAt  time.Time  `json:"created_at"`
}

// --- Notification CRUD ---

func CreateNotification(n *Notification) error {
	return DB.Create(n).Error
}

func FindNotifications(userID uint, isRead *bool, page, pageSize int) ([]Notification, int) {
	var notifs []Notification
	var total int
	query := DB.Model(&Notification{}).Where("user_id = ?", userID)
	if isRead != nil {
		query = query.Where("is_read = ?", *isRead)
	}
	query.Count(&total)
	if page > 0 && pageSize > 0 {
		query = query.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	query.Order("id desc").Find(&notifs)
	return notifs, total
}

func MarkNotificationRead(id uint) error {
	return DB.Model(&Notification{}).Where("id = ?", id).Update("is_read", true).Error
}

func MarkAllNotificationsRead(userID uint) error {
	return DB.Model(&Notification{}).Where("user_id = ? AND is_read = ?", userID, false).Update("is_read", true).Error
}

func CountUnreadNotifications(userID uint) int {
	var count int
	DB.Model(&Notification{}).Where("user_id = ? AND is_read = ?", userID, false).Count(&count)
	return count
}

// --- AuditLog CRUD ---

func CreateAuditLog(log *AuditLog) error {
	return DB.Create(log).Error
}

func FindAuditLogs(userID uint, action string, targetType string, page, pageSize int) ([]AuditLog, int) {
	var logs []AuditLog
	var total int
	query := DB.Model(&AuditLog{})
	if userID > 0 {
		query = query.Where("user_id = ?", userID)
	}
	if action != "" {
		query = query.Where("action = ?", action)
	}
	if targetType != "" {
		query = query.Where("target_type = ?", targetType)
	}
	query.Count(&total)
	if page > 0 && pageSize > 0 {
		query = query.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	query.Order("id desc").Find(&logs)
	return logs, total
}
