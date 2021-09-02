package controller

import (
	"cheese/logic"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func GetCommunityHandler(c *gin.Context) {
	//获取数据
	list, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic community list get error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, list)
}
func GetDetailHandler(c *gin.Context) {
	//获取数据
	idp := c.Param("id")
	id, err := strconv.ParseInt(idp, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	fmt.Printf("%d", id)
	data, err := logic.GetSingleCommunity(id)
	if err != nil {
		zap.L().Error("logic get single community error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
