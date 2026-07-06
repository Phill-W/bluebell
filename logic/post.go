package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	//1. 生成post id
	p.ID = snowflake.GenID()
	//2. 保存到数据库
	return mysql.CreatePost(p)
	//3.返回
}

// GetPostByID 获取帖子详情
func GetPostByID(pid int64) (post *models.Post, err error) {
	return mysql.GetPostByID(pid)
}
