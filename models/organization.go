package models

import (
	"time"
)

// Organization 組織/機關 (學校、縣市教育局、承辦單位、教育署等)
type Organization struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `json:"name" gorm:"size:255;not null"`        // 組織名稱
	Code      string     `json:"code" gorm:"size:50;uniqueIndex"`      // 組織代碼 (學校代碼)
	OrgType   string     `json:"org_type" gorm:"size:20;not null"`     // school/county/ministry/contractor
	ParentID  uint       `json:"parent_id" gorm:"default:0;index"`     // 上級組織 ID
	County    string     `json:"county" gorm:"size:50"`                // 所屬縣市
	District  string     `json:"district" gorm:"size:50"`              // 所屬區域
	Level     string     `json:"level" gorm:"size:20"`                 // 學校層級: elementary/junior_high/both
	Address   string     `json:"address" gorm:"size:255"`              // 地址
	Phone     string     `json:"phone" gorm:"size:50"`                 // 電話
	Principal string     `json:"principal" gorm:"size:50"`             // 負責人/校長
	IsActive  bool       `json:"is_active" gorm:"default:true"`       // 是否啟用
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// --- CRUD ---

func CreateOrganization(org *Organization) error {
	return DB.Create(org).Error
}

func UpdateOrganization(org *Organization) error {
	return DB.Save(org).Error
}

func FindOrganizationByID(id uint) (Organization, error) {
	var org Organization
	err := DB.Where("id = ?", id).First(&org).Error
	return org, err
}

func FindOrganizationByCode(code string) (Organization, error) {
	var org Organization
	err := DB.Where("code = ?", code).First(&org).Error
	return org, err
}

func FindOrganizations(orgType string, county string, parentID uint, page, pageSize int) ([]Organization, int) {
	var orgs []Organization
	var total int
	query := DB.Model(&Organization{})
	if orgType != "" {
		query = query.Where("org_type = ?", orgType)
	}
	if county != "" {
		query = query.Where("county = ?", county)
	}
	if parentID > 0 {
		query = query.Where("parent_id = ?", parentID)
	}
	query.Count(&total)
	if page > 0 && pageSize > 0 {
		query = query.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	query.Order("id asc").Find(&orgs)
	return orgs, total
}

func DeleteOrganization(id uint) error {
	return DB.Where("id = ?", id).Delete(&Organization{}).Error
}

func FindSchoolsByCounty(county string) []Organization {
	var orgs []Organization
	DB.Where("org_type = ? AND county = ? AND is_active = ?", "school", county, true).Order("name asc").Find(&orgs)
	return orgs
}
