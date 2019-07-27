/**
 * @Author: huangw1
 * @Date: 2019/7/29 16:30
 */

package model

import "html/template"

type UserInfo struct {
	Id          int64    `json:"id"`
	Name        string   `json:"name"`
	Username    string   `json:"username"`
	Avatar      string   `json:"avatar"`
	Email       string   `json:"email"`
	Type        int      `json:"type"`
	Roles       []string `json:"roles"`
	Description string   `json:"description"`
	CreateTime  int64    `json:"createTime"`
}

type CategoryResponse struct {
	CategoryId   int64  `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}

type TagResponse struct {
	TagId   int64  `json:"tagId"`
	TagName string `json:"tagName"`
}

type ArticleResponse struct {
	ArticleId  int64             `json:"articleId"`
	User       *UserInfo         `json:"user"`
	Category   *CategoryResponse `json:"category"`
	Tags       []*TagResponse    `json:"tags"`
	Title      string            `json:"title"`
	Summary    string            `json:"summary"`
	Content    template.HTML     `json:"content"`
	Toc        template.HTML     `json:"toc"`
	Type       int               `json:"type"`
	SourceUrl  string            `json:"sourceUrl"`
	CreateTime int64             `json:"createTime"`
}

type TopicResponse struct {
	TopicId         int64          `json:"topicId"`
	User            *UserInfo      `json:"user"`
	Tags            []*TagResponse `json:"tags"`
	Title           string         `json:"title"`
	Content         template.HTML  `json:"content"`
	Toc             template.HTML  `json:"toc"`
	LastCommentTime int64          `json:"lastCommentTime"`
	ViewCount       int64          `json:"viewCount"`
	CreateTime      int64          `json:"createTime"`
}

type CommentResponse struct {
	CommentId    int64            `json:"commentId"`
	User         *UserInfo        `json:"user"`
	EntityType   string           `json:"entityType"`
	EntityId     int64            `json:"entityId"`
	Content      template.HTML    `json:"content"`
	QuoteId      int64            `json:"quoteId"`
	Quote        *CommentResponse `json:"quote"`
	QuoteContent template.HTML    `json:"quoteContent"`
	Status       int              `json:"status"`
	CreateTime   int64            `json:"createTime"`
}

type FavoriteResponse struct {
	FavoriteId int64     `json:"favoriteId"`
	EntityType string    `json:"entityType"`
	EntityId   int64     `json:"entityId"`
	Deleted    bool      `json:"deleted"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	User       *UserInfo `json:"user"`
	Url        string    `json:"url"`
	CreateTime int64     `json:"createTime"`
}

type MessageResponse struct {
	MessageId    int64     `json:"messageId"`
	From         *UserInfo `json:"from"`
	UserId       int64     `json:"userId"`
	Content      string    `json:"content"`
	QuoteContent string    `json:"quoteContent"`
	Type         int       `json:"type"`
	DetailUrl    string    `json:"detailUrl"`
	ExtraData    string    `json:"extraData"`
	Status       int       `json:"status"`
	CreateTime   int64     `json:"createTime"`
}
