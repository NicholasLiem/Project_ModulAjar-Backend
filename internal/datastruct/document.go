package datastruct

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	DocumentID   string `gorm:"uniqueIndex:document_id" json:"document_id,omitempty"`
	UserRefer    uint   `gorm:"column:user_refer" json:"user_refer,omitempty"`
	DocumentPath string `gorm:"column:document_path" json:"document_path,omitempty"`
}
