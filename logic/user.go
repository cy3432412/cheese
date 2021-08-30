package logic

import (
	"cheese/dao/mysql"
	"cheese/models"
	"cheese/pkg/snowflake"
)

//处理user逻辑

//SignUp 处理用户注册
func SignUp(p *models.ParamSignUp) (err error) {
	//1. 用户不存在
	if err := mysql.CheckUser(p.Username); err != nil {
		return err
	}
	//2. 写入数据库
	userID := snowflake.GenID()
	user := &mysql.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (err error) {
	
}
