package datastruct

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	UserID    uint       `gorm:"uniqueIndex:user_id;autoIncrement" json:"user_id,omitempty"`
	Email     string     `gorm:"uniqueIndex:email" json:"email,omitempty"`
	Password  string     `gorm:"column:password;not null" json:"password_hash,omitempty"`
	Role      Role       `gorm:"column:role;not null;default:'user'" json:"role,omitempty"`
	Documents []Document `gorm:"foreignKey:UserRefer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"documents,omitempty"`
}

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)
