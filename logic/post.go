package logic

import (
	"cheese/dao/mysql"
	"cheese/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *mysql.Post) (err error) {
	// 1. 生成post id
	p.PostID = snowflake.GenID()
	// 2. 保存到数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	//err = redis.CreatePost(p.ID, p.CommunityID)
	return
	// 3. 返回
}

//// GetPostById 根据帖子id查询帖子详情数据
//func GetPostById(pid int64) (data *mysql.ApiPostDetail, err error) {
//	// 查询并组合我们接口想用的数据
//	post, err := mysql.GetPostById(pid)
//	if err != nil {
//		zap.L().Error("mysql.GetPostById(pid) failed",
//			zap.Int64("pid", pid),
//			zap.Error(err))
//		return
//	}
//	// 根据作者id查询作者信息
//	user, err := mysql.GetUserById(post.AuthorID)
//	if err != nil {
//		zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
//			zap.Int64("author_id", post.AuthorID),
//			zap.Error(err))
//		return
//	}
//	// 根据社区id查询社区详细信息
//	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
//	if err != nil {
//		zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
//			zap.Int64("community_id", post.CommunityID),
//			zap.Error(err))
//		return
//	}
//	// 接口数据拼接
//	data = &mysql.ApiPostDetail{
//		AuthorName:      user.Username,
//		Post:            post,
//		Community: community,
//	}
//	return
//}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (data []*mysql.ApiPostDetail, err error) {

	posts, err := mysql.GetPostList(page, size)

	if err != nil {
		return nil, err
	}

	data = make([]*mysql.ApiPostDetail, 0, len(posts))
	zap.L().Info("len posts = ", zap.Any("len", len(posts)))
	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		//根据社区id查询社区详细信息
		community, err := mysql.GetCommunityById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		postDetail := &mysql.ApiPostDetail{
			AuthorName: user.Username,
			Post:       post,
			Community:  &community,
		}
		data = append(data, postDetail)
	}
	return
}
