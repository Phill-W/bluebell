package controller

import (
	"bluebell/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ---- 和社区相关的内容 ----

// CommunityHandler 社区列表处理函数
func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区（community_id, community_name) 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区详情处理函数
func CommunityDetailHandler(c *gin.Context) {
	//1. 获取社区id
	communityID := c.Param("id")
	id, err := strconv.ParseInt(communityID, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	//2. 根据id获取社区详情
	data, err := logic.GetCommunityDetailByID(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetailByID() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不轻易把服务器报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
