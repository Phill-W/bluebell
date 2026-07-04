package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

// GetCommunityList 获取社区列表的业务逻辑
func GetCommunityList() ([]*models.Community, error) {
	//查数据库 查找到所有的community 并返回
	return mysql.GetCommunityList()
}
