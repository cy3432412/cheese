package mysql

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Post struct {
	gorm.Model
	ID          int64     `json:"id" gorm:"primaryKey;column:id;type:bigint(20) auto_increment"`
	PostID      int64     `json:"post_id" gorm:"column:post_id;type:bigint(20);not null;uniqueIndex:idx_post_id;comment:帖子id"`          // 帖子id
	Title       string    `json:"title" gorm:"column:title;type:varchar(128);not null;comment:标题"`                                      // 标题
	Content     string    `json:"content" gorm:"column:content;type:varchar(8192);not null;comment:内容"`                                 // 内容
	AuthorID    int64     `json:"author_id" gorm:"column:author_id;type:bigint(20);not null;index:idx_author_id;comment:作者的用户id"`       // 作者的用户id
	CommunityID int64     `json:"community_id" gorm:"column:community_id;type:bigint(20);not null;index:idx_community_id;comment:所属社区"` // 所属社区
	Status      int32     `json:"status" gorm:"column:status;type:tinyint(4);not null;default:1;comment:帖子状态"`                          // 帖子状态
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间"`            // 创建时间
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:更新时间"`            // 更新时间
	DeletedAt   time.Time `json:"deleted_at" gorm:"column:deleted_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:删除时间"`            // 删除时间
}

// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName string             `json:"author_name"` // 作者
	VoteNum    int64              `json:"vote_num"`    // 投票数
	*Post                         // 嵌入帖子结构体
	*Community `json:"community"` // 嵌入社区信息
}

func CreatePost(post *Post) (err error) {
	err = db.Create(&post).Error
	return err
}

// GetPostList 查询帖子列表函数
func GetPostList(page, size int64) (posts []*Post, err error) {
	posts = make([]*Post, 0, 2) // 不要写成make([]*models.Post, 2)
	db.Find(&posts)
	return
}
