package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID       int64  `json:"id" gorm:"primaryKey;column:id;type:bigint(20) auto_increment"`
	UserID   int64  `json:"user_id" gorm:"column:user_id;type:bigint(20);not null;uniqueIndex:idx_user_id"`
	Username string `json:"username" gorm:"column:username;type:varchar(64);not null;uniqueIndex:idx_username"`
	Password string `json:"password" gorm:"column:password;type:varchar(64);not null"`
	Email    string `json:"email" gorm:"column:email;type:varchar(64)"`
	Gender   int32  `json:"gender" gorm:"column:gender;type:tinyint(4);not null;default:0"`
}

// 存放对用户数据的增删改查
var secret = "caoyang"

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
func CheckUser(username string) (err error) {
	var user User
	res := db.Select("id").Where("username = ?", username).First(&user)
	if res.RowsAffected > 0 {
		return ErrorUserExist
	}
	return nil
}

func InsertUser(user *User) (err error) {
	user.Password = encryptPassword(user.Password)
	err = db.Create(&user).Error

	return err
}

func CheckLogin(user *User) (err error) {
	oPassword := user.Password
	err = db.Where("username = ?", user.Username).First(&user).Error
	if err != nil {
		return err
	}

	if encryptPassword(oPassword) != user.Password {
		return ErrorInvalidPassword
	}

	return nil
}
