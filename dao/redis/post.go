package redis

import "bluebell/models"

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	//从redis获取id
	//1. 根据用户请求中携带的排序参数，确定从redis中获取数据的key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 2. 确定查询的索引起始页
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	// 3. ZRevRange 按照分数从大到小排序查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}
