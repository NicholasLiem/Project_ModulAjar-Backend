package datastruct

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	UserID    uint       `gorm:"uniqueIndex:user_id" json:"user_id,omitempty"`
	Username  string     `gorm:"column:username" json:"username,omitempty"`
	Email     string     `gorm:"column:email;unique_index" json:"email,omitempty"`
	Password  string     `gorm:"column:password;not null" json:"password_hash,omitempty"`
	Documents []Document `gorm:"foreignKey:UserRefer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"documents,omitempty"`
}

//func (model *UserModel) SetPassword(password string) error {
//	if len(password) == 0 {
//		return errors.New("password can not be empty")
//	}
//
//	bytePassword := []byte(password)
//	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
//	model.PasswordHash = string(passwordHash)
//	return nil
//}
//
//func (model *UserModel) CheckPassword(password string) error {
//	bytePassword := []byte(password)
//	byteHashedPassword := []byte(model.PasswordHash)
//	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
//}
