/**
 * @Author: huangw1
 * @Date: 2019/7/25 18:51
 */

package model

type Model struct {
	Id         int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id" form:"id"`
	CreateTime int64 `json:"createTime" form:"createTime"`
	UpdateTime int64 `json:"updateTime" form:"updateTime"`
}

// t_user
const (
	UserStatusOk = iota
	UserStatusDisabled
)

const (
	UserTypeNormal = iota
	UserTypeGzh
)

type User struct {
	Model
	Username    string `json:"username" form:"username"`
	Name        string `json:"name" form:"name"`
	Avatar      string `json:"avatar" form:"avatar"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Roles       string `json:"roles" form:"roles"`
	Description string `json:"description" form:"description"`
	City        string `json:"city" form:"city"`
	Company     string `json:"company" form:"company"`
	Type        int    `json:"type" form:"type"`
	Status      int    `json:"status" form:"status"`
}

// t_third_user
const (
	ThirdTypeGithub = iota
)

type ThirdUser struct {
	Model
	UserId   int64  `json:"userId" form:"userId"`
	ThirdId  int64  `json:"thirdId" form:"thirdId"`
	Type     int    `json:"type" form:"type"`
	Username string `json:"Username" form:"Username"`
	Name     string `json:"Name" form:"Name"`
	Email    string `json:"email" form:"email"`
	Avatar   string `json:"avatar" form:"avatar"`
	Url      string `json:"url" form:"url"`
	HtmlUrl  string `json:"htmlUrl" form:"htmlUrl"`
}

// t_category
const (
	CategoryStatusOk = iota
	CategoryStatusDisabled
)

type Category struct {
	Model
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Status      int    `json:"status" form:"status"`
}

// t_tag
const (
	TagStatusOk = iota
	TagStatusDisabled
)

type Tag struct {
	Model
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Status      int    `json:"status" form:"status"`
}

// t_article
const (
	ArticleStatusPublished = iota
	ArticleStatusDeleted
	ArticleStatusDraft
)

const (
	ArticleTypeOriginal = iota
	ArticleTypeShared
)

const (
	ArticleContentTypeHtml     = "html"
	ArticleContentTypeMarkdown = "markdown"
)

type Article struct {
	Model
	CategoryId  int64  `json:"categoryId" form:"categoryId"`
	UserId      int64  `json:"userId" form:"userId"`
	Title       string `json:"title" form:"title"`
	Summary     string `json:"summary" form:"summary"`
	Content     string `json:"content" form:"content"`
	ContentType string `json:"contentType" form:"contentType"`
	Status      int    `json:"status" form:"status"`
	Type        int    `json:"type" form:"type"`
	SourceUrl   string `json:"sourceUrl" form:"sourceUrl"`
}

// t_article_tag
type ArticleTag struct {
	Model
	ArticleId int64 `json:"articleId" form:"articleId"`
	TagId     int64 `json:"tagId" form:"tagId"`
}

// t_comment
const (
	EntityTypeArticle = "article"
	EntityTypeTopic   = "topic"
)

const (
	CommentStatusOk = iota
	CommentStatusDeleted
)

type Comment struct {
	Model
	UserId     int64  `json:"userId" form:"userId"`
	EntityType string `json:"entityType" form:"entityType"`
	EntityId   int64  `json:"entityId" form:"entityId"`
	Content    string `json:"content" form:"content"`
	QuoteId    int64  `json:"quoteId" form:"quoteId"`
	Status     int    `json:"status" form:"status"`
}

// t_favorite
type Favorite struct {
	Model
	UserId     int64  `json:"userId" form:"userId"`
	EntityType string `json:"entityType" form:"entityType"`
	EntityId   int64  `json:"entityId" form:"entityId"`
}

// t_topic
const (
	TopicStatusOk = iota
	TopicStatusDeleted
)

type Topic struct {
	Model
	UserId          int64  `json:"userId" form:"userId"`
	Title           string `json:"title" form:"title"`
	Content         string `json:"content" form:"content"`
	ViewCount       int64  `json:"viewCount" form:"viewCount"`
	Status          int    `json:"status" form:"status"`
	LastCommentTime int64  `json:"lastCommentTime" form:"lastCommentTime"`
}

// t_topic_tag
type TopicTag struct {
	Model
	TopicId int64 `json:"topicId" form:"topicId"`
	TagId   int64 `json:"tagId" form:"tagId"`
}

// t_message
const (
	MsgStatusUnread = iota
	MsgStatusReaded
)

const (
	MsgTypeComment = iota
	MsgTypeSystem
)

type Message struct {
	Model
	FromId       int64  `json:"fromId" form:"fromId"`
	UserId       int64  `json:"userId" form:"userId"`
	Content      string `json:"content" form:"content"`
	QuoteContent string `json:"quoteContent" form:"quoteContent"`
	Type         int    `json:"type" form:"type"`
	ExtraData    string `json:"extraData" form:"extraData"`
	Status       int    `json:"status" form:"status"`
}

type SysConfig struct {
	Model
	Key         string `json:"key" form:"key"`
	Value       string `json:"value" form:"value"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}
