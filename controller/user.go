package controller

import (
	"cheese/dao/mysql"
	"cheese/logic"
	"cheese/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 处理用户注册
func SignUpHandler(c *gin.Context) {
	//1. 获取参数并且进行校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with error params ", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}
	//2. 处理注册逻辑
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return 
	}
	//3. 返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 处理用户登录
func LoginHandler(c *gin.Context) {
	//1. 获取数据校验
	p := new (models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with error params ", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}
	//2. 处理登录逻辑
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//3. 返回响应
	ResponseSuccess(c,gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID), // id值大于1<<53-1  int64类型的最大值是1<<63-1
		"user_name": user.Username,
		"token":     user.Token,
	})
}
