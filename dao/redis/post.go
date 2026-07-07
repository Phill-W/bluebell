package redis

import (
	"bluebell/models"

	"github.com/go-redis/redis"
)

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

// GetPostVoteData 根据ids查询每篇帖子的投赞成票数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPF + id)
	//	// 查找key中分数是1的元素数量-> 统计每篇帖子的赞成票数
	//	v1 := client.ZCount(key, "1", "1").Val()
	//	data = append(data, v1)
	//}
	// 使用pipeline一次发送多条数据减少RTT
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
