package model

type Posts struct {
	Id           int    `gorm:"column:id;primaryKey"`
	Typ          int32  `gorm:"column:type"`
	Title        string `gorm:"column:title"`
	Cover        string `gorm:"column:cover"`
	Content      string `gorm:"column:content"`
	Introduction string `gorm:"column:introduction"`
	CreateTime   int64  `gorm:"column:create_time"`
	UpdateTime   int64  `gorm:"column:update_time"`
	Status       int32  `gorm:"column:status"`
	RemoveTime   int64  `gorm:"column:remove_time"`
}

type Tags struct {
	Id   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}

type PostTags struct {
	Id     int `gorm:"column:id;primaryKey"`
	PostId int `gorm:"column:post_id"`
	TagId  int `gorm:"column:tag_id"`
}

type Comment struct {
	ID           int    `gorm:"primaryKey;autoIncrement;column:id"`                       // 主键自增
	Nickname     string `gorm:"type:varchar(255);not null;comment:评论用户名;column:nickname"` // 强制非空
	Email        string `gorm:"type:varchar(255);not null;comment:评论用户邮箱;column:email"`   // 强制非空
	Content      string `gorm:"type:varchar(255);not null;comment:评论内容;column:content"`   // 强制非空
	CommentID    int    `gorm:"comment:用户回复的评论;column:comment_id"`                        // 指向自身的指针（可空）
	PostID       int    `gorm:"not null;comment:文章Id;column:post_id"`                     // 强制非空
	CreateTime   int64  `gorm:"comment:评论时间;column:create_time"`                          // bigint存储时间戳
	Status       int    `gorm:"comment:状态：0-未审核，1-审核通过，2-拒绝;column:status"`               // 枚举值指针（可空）
	ReviewTime   int64  `gorm:"comment:审核时间;column:review_time"`                          // 可空时间戳
	ReviewRemark string `gorm:"type:varchar(255);comment:审核备注;column:review_remark"`      // 可空字符串
	CommentIP    string `gorm:"type:varchar(255);comment:评论用户IP;column:comment_ip"`       // 可空字符串
	Reply        string `gorm:"type:varchar(255);comment:作者回复;column:reply"`              // 可空字符串
}
