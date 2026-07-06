package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
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
func GetPostByID(pid int64) (data *models.ApiPostDetail, err error) {
	//查询数据并且拼接组合我们接口想要的数据
	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(pid) failed", zap.Int64("pid", pid), zap.Error(err))
	}
	//根据作者id查询作者信息
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserInfoByID(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
		return
	}
	//根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
		return
	}
	//将查询到的数据拼接
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		CommunityDetail: community,
		Post:            post,
	}
	return
}
