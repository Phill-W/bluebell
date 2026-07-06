package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// CtxUserIDKey 在gin上下文中存储当前登录用户ID的key
const CtxUserIDKey = "UserID"

// ErrorUserNotLogin 用户未登录错误
var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUserID 获得当前登录用户ID
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
