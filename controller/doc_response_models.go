package controller

import "bluebell/models"

// 专门用来放接口文档用到的model
// 因为我们的接口文档返回的数据格式是一致的，但是具体的data类型不一致

// _Response 通用响应数据(data为空时使用)
type _Response struct {
	Code ResCode     `json:"code"`     // 业务响应状态码
	Msg  interface{} `json:"msg"`      // 提示信息
	Data interface{} `json:"data"`     // 数据
}

// _ResponseSignUp 注册接口响应数据
type _ResponseSignUp struct {
	Code ResCode     `json:"code"` // 业务响应状态码
	Msg  interface{} `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 数据
}

// _LoginData 登录成功返回的数据
type _LoginData struct {
	UserID   string `json:"user_id"`   // 用户id
	UserName string `json:"user_name"` // 用户名
	Token    string `json:"token"`     // JWT token
}

// _ResponseLogin 登录接口响应数据
type _ResponseLogin struct {
	Code ResCode   `json:"code"` // 业务响应状态码
	Msg  interface{} `json:"msg"`  // 提示信息
	Data _LoginData `json:"data"` // 数据
}

// _ResponseCommunityList 社区列表接口响应数据
type _ResponseCommunityList struct {
	Code ResCode             `json:"code"` // 业务响应状态码
	Msg  interface{}         `json:"msg"`  // 提示信息
	Data []*models.Community `json:"data"` // 数据
}

// _ResponseCommunityDetail 社区详情接口响应数据
type _ResponseCommunityDetail struct {
	Code ResCode                 `json:"code"` // 业务响应状态码
	Msg  interface{}             `json:"msg"`  // 提示信息
	Data *models.CommunityDetail `json:"data"` // 数据
}

// _ResponsePostDetail 帖子详情接口响应数据
type _ResponsePostDetail struct {
	Code ResCode                 `json:"code"` // 业务响应状态码
	Msg  interface{}             `json:"msg"`  // 提示信息
	Data *models.ApiPostDetail   `json:"data"` // 数据
}

// _ResponsePostList 帖子列表接口响应数据
type _ResponsePostList struct {
	Code    ResCode                 `json:"code"`    // 业务响应状态码
	Msg     interface{}             `json:"msg"`     // 提示信息
	Data    []*models.ApiPostDetail `json:"data"`    // 数据
}
