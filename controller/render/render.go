/**
 * @Author: huangw1
 * @Date: 2019/7/29 16:34
 */

package render

import (
	"github.com/huangw1/bbs/dao"
	"github.com/huangw1/bbs/model"
	"github.com/huangw1/bbs/pkg/avatar"
	"github.com/huangw1/bbs/pkg/markdown"
	"github.com/huangw1/bbs/service"
	template2 "html/template"
	"strings"
)

func BuildUserById(id int64) *model.UserInfo {
	user, err := service.UserService.Get(id)
	if err != nil {
		return nil
	}
	return BuildUser(user)
}

func BuildUserInIds(ids []int64) []*model.UserInfo {
	var responses []*model.UserInfo
	users, err := service.UserService.QueryCondition(model.NewCondition("id in (?)", ids))
	if err != nil {
		return responses
	}
	for _, user := range users {
		responses = append(responses, BuildUser(user))
	}
	return responses
}

func BuildUser(user *model.User) *model.UserInfo {
	if user == nil {
		return nil
	}
	avatarURL := user.Avatar
	if len(avatarURL) == 0 {
		avatarURL = avatar.GetDefaultAvatar(user.Id)
	}
	roles := strings.Split(user.Roles, ",")
	return &model.UserInfo{
		Id:          user.Id,
		Username:    user.Username,
		Name:        user.Name,
		Avatar:      avatarURL,
		Email:       user.Email,
		Type:        user.Type,
		Roles:       roles,
		Description: user.Description,
		CreateTime:  user.CreateTime,
	}
}

func BuildTagInIds(tagIds []int64) []*model.TagResponse {
	var responses []*model.TagResponse
	if len(tagIds) == 0 {
		return responses
	}
	tags, err := dao.TagDaoImpl.GetTagInIds(tagIds)
	if err != nil {
		return responses
	}
	return BuildTags(tags)
}

func BuildTag(tag *model.Tag) *model.TagResponse {
	if tag == nil {
		return nil
	}
	return &model.TagResponse{TagId: tag.Id, TagName: tag.Name}
}

func BuildTags(tags []*model.Tag) []*model.TagResponse {
	var responses []*model.TagResponse
	if len(tags) == 0 {
		return responses
	}
	for _, tag := range tags {
		responses = append(responses, BuildTag(tag))
	}
	return responses
}

func BuildArticle(article *model.Article) *model.ArticleResponse {
	if article == nil {
		return nil
	}
	res := &model.ArticleResponse{}
	res.ArticleId = article.Id
	res.Type = article.Type
	res.Title = article.Title
	res.Summary = article.Summary
	res.SourceUrl = article.SourceUrl
	res.CreateTime = article.CreateTime

	res.User = BuildUserById(article.UserId)
	articleTags, err := dao.ArticleTagDaoImpl.GetByArticleId(article.Id)
	var tagIds []int64
	if err == nil {
		for _, articleTag := range articleTags {
			tagIds = append(tagIds, articleTag.TagId)
		}
		res.Tags = BuildTagInIds(tagIds)
	}

	if article.ContentType == model.ArticleContentTypeMarkdown {
		md := markdown.NewMD(markdown.MdWithTOC()).Run(article.Content)
		res.Content = template2.HTML(md.ContentHtml)
		res.Toc = template2.HTML(md.TocHtml)
		if len(res.Summary) == 0 {
			res.Summary = md.SummaryText
		}
	}

	return res
}

func BuildArticles(articles []*model.Article) []*model.ArticleResponse {
	var responses []*model.ArticleResponse
	if len(articles) == 0 {
		return responses
	}
	for _, article := range articles {
		responses = append(responses, BuildArticle(article))
	}
	return responses
}
