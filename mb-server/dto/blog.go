package dto

type PostsStatus int32

const (
	PostStatus_Draft   PostsStatus = 0
	PostStatus_Publish PostsStatus = 1
	PostStatus_Removed PostsStatus = 2
)

type PostsQueryDto struct {
	Page  BasePageQuery
	Title string `json:"title"`
}

type PoststDto struct {
	Id           int64       `json:"id"`
	Title        string      `json:"title"`
	Cover        string      `json:"cover"`
	Content      string      `json:"content"`
	Introduction string      `json:"introduction"`
	CreateTime   int64       `json:"create_time"`
	UpdateTime   int64       `json:"update_time"`
	RemoveTime   int64       `json:"remove_time"`
	Status       PostsStatus `json:"status"`
}
