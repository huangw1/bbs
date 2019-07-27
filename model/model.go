/**
 * @Author: huangw1
 * @Date: 2019/7/25 18:51
 */

package model

type Model struct {
	Id         int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;column:id" json:"id" form:"id"`
	CreateTime int64 `gorm:"column:createTime" json:"createTime" form:"createTime"`
	UpdateTime int64 `gorm:"column:updateTime" json:"updateTime" form:"updateTime"`
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
	Username    string `gorm:"column:username" json:"username" form:"username"`
	Name        string `gorm:"column:name" json:"name" form:"name"`
	Avatar      string `gorm:"column:avatar" json:"avatar" form:"avatar"`
	Email       string `gorm:"column:email" json:"email" form:"email"`
	Password    string `gorm:"column:password" json:"password" form:"password"`
	Roles       string `gorm:"column:roles" json:"roles" form:"roles"`
	Description string `gorm:"column:description" json:"description" form:"description"`
	City        string `gorm:"column:city" json:"city" form:"city"`
	Company     string `gorm:"column:company" json:"company" form:"company"`
	Type        int    `gorm:"column:type" json:"type" form:"type"`
	Status      int    `gorm:"column:status" json:"status" form:"status"`
}

// t_third_user
const (
	ThirdTypeGithub = iota
)

type ThirdUser struct {
	Model
	UserId   int64  `gorm:"column:userId" json:"userId" form:"userId"`
	ThirdId  int64  `gorm:"column:thirdId" json:"thirdId" form:"thirdId"`
	Type     int    `gorm:"column:type" json:"type" form:"type"`
	Username string `gorm:"column:username" json:"username" form:"username"`
	Name     string `gorm:"column:name" json:"name" form:"name"`
	Email    string `gorm:"column:email" json:"email" form:"email"`
	Avatar   string `gorm:"column:avatar" json:"avatar" form:"avatar"`
	Url      string `gorm:"column:url" json:"url" form:"url"`
	HtmlUrl  string `gorm:"column:htmlUrl" json:"htmlUrl" form:"htmlUrl"`
}

// t_category
const (
	CategoryStatusOk = iota
	CategoryStatusDisabled
)

type Category struct {
	Model
	Name        string `gorm:"column:name" json:"name" form:"name"`
	Description string `gorm:"column:description" json:"description" form:"description"`
	Status      int    `gorm:"column:status" json:"status" form:"status"`
}

// t_tag
const (
	TagStatusOk = iota
	TagStatusDisabled
)

type Tag struct {
	Model
	Name        string `gorm:"column:name" json:"name" form:"name"`
	Description string `gorm:"column:description" json:"description" form:"description"`
	Status      int    `gorm:"column:status" json:"status" form:"status"`
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
	CategoryId  int64  `gorm:"column:categoryId" json:"categoryId" form:"categoryId"`
	UserId      int64  `gorm:"column:userId" json:"userId" form:"userId"`
	Title       string `gorm:"column:title" json:"title" form:"title"`
	Summary     string `gorm:"column:summary" json:"summary" form:"summary"`
	Content     string `gorm:"column:content" json:"content" form:"content"`
	ContentType string `gorm:"column:contentType" json:"contentType" form:"contentType"`
	Status      int    `gorm:"column:status" json:"status" form:"status"`
	Type        int    `gorm:"column:type" json:"type" form:"type"`
	SourceUrl   string `gorm:"column:sourceUrl" json:"sourceUrl" form:"sourceUrl"`
}

// t_article_tag
type ArticleTag struct {
	Model
	ArticleId int64 `gorm:"column:articleId" json:"articleId" form:"articleId"`
	TagId     int64 `gorm:"column:tagId" json:"tagId" form:"tagId"`
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
	UserId     int64  `gorm:"column:userId" json:"userId" form:"userId"`
	EntityType string `gorm:"column:entityType" json:"entityType" form:"entityType"`
	EntityId   int64  `gorm:"column:entityId" json:"entityId" form:"entityId"`
	Content    string `gorm:"column:content" json:"content" form:"content"`
	QuoteId    int64  `gorm:"column:quoteId" json:"quoteId" form:"quoteId"`
	Status     int    `gorm:"column:status" json:"status" form:"status"`
}

// t_favorite
type Favorite struct {
	Model
	UserId     int64  `gorm:"column:userId" json:"userId" form:"userId"`
	EntityType string `gorm:"column:entityType" json:"entityType" form:"entityType"`
	EntityId   int64  `gorm:"column:entityId" json:"entityId" form:"entityId"`
}

// t_topic
const (
	TopicStatusOk = iota
	TopicStatusDeleted
)

type Topic struct {
	Model
	UserId          int64  `gorm:"column:userId" json:"userId" form:"userId"`
	Title           string `gorm:"column:title" json:"title" form:"title"`
	Content         string `gorm:"column:content" json:"content" form:"content"`
	ViewCount       int64  `gorm:"column:viewCount" json:"viewCount" form:"viewCount"`
	Status          int    `gorm:"column:status" json:"status" form:"status"`
	LastCommentTime int64  `gorm:"column:lastCommentTime" json:"lastCommentTime" form:"lastCommentTime"`
}

// t_topic_tag
type TopicTag struct {
	Model
	TopicId int64 `gorm:"column:topicId" json:"topicId" form:"topicId"`
	TagId   int64 `gorm:"column:tagId" json:"tagId" form:"tagId"`
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
	FromId       int64  `gorm:"column:fromId" json:"fromId" form:"fromId"`
	UserId       int64  `gorm:"column:userId" json:"userId" form:"userId"`
	Content      string `gorm:"column:content" json:"content" form:"content"`
	QuoteContent string `gorm:"column:quoteContent" json:"quoteContent" form:"quoteContent"`
	Type         int    `gorm:"column:type" json:"type" form:"type"`
	ExtraData    string `gorm:"column:extraData" json:"extraData" form:"extraData"`
	Status       int    `gorm:"column:status" json:"status" form:"status"`
}

type SystemConfig struct {
	Model
	Key         string `gorm:"column:key" json:"key" form:"key"`
	Value       string `gorm:"column:value" json:"value" form:"value"`
	Name        string `gorm:"column:name" json:"name" form:"name"`
	Description string `gorm:"column:description" json:"description" form:"description"`
}
