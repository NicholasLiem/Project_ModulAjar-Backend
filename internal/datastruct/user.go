package datastruct

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	UserID    uint       `gorm:"uniqueIndex:user_id" json:"user_id,omitempty"`
	Email     string     `gorm:"column:email;unique_index" json:"email,omitempty"`
	Password  string     `gorm:"column:password;not null" json:"password_hash,omitempty"`
	Documents []Document `gorm:"foreignKey:UserRefer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"documents,omitempty"`
}
