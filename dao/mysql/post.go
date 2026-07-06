package mysql

import "bluebell/models"

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
			 post_id, title, content, author_id, community_id
			 ) values (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)

	return

}

// GetPostByID 获取帖子详情
func GetPostByID(pid int64) (post *models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post where post_id = ?`
	post = new(models.Post)
	err = db.Get(post, sqlStr, pid)
	return
}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select
	post_id, title, content, author_id, community_id, create_time
	from post
	limit ?,?
`
	posts = make([]*models.Post, 0, 2) //不要写成make([]*models.Post, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}
