/**
 * @Author: huangw1
 * @Date: 2019/7/29 14:20
 */

package service

import (
	"github.com/huangw1/bbs/dao"
	"github.com/huangw1/bbs/database"
	"github.com/huangw1/bbs/model"
	"github.com/huangw1/bbs/pkg/extension"
	"github.com/jinzhu/gorm"
)

var ArticleService = NewArticleService()

func NewArticleService() *articleService {
	return &articleService{}
}

type articleService struct{}

func (service *articleService) Get(id int64) (*model.Article, error) {
	return dao.ArticleDaoImpl.Get(id)
}

func (service *articleService) Take(where ...interface{}) (*model.Article, error) {
	return dao.ArticleDaoImpl.Take(where...)
}

func (service *articleService) QueryCondition(condition *model.Condition) ([]*model.Article, error) {
	return dao.ArticleDaoImpl.QueryCondition(condition)
}

func (service *articleService) QueryParams(params *model.Params) ([]*model.Article, *model.Paging, error) {
	return dao.ArticleDaoImpl.QueryParams(params)
}

func (service *articleService) Create(article *model.Article) error {
	return dao.ArticleDaoImpl.Create(article)
}

func (service *articleService) Update(article *model.Article) error {
	return dao.ArticleDaoImpl.Update(article)
}

func (service *articleService) Updates(id int64, columns map[string]interface{}) error {
	return dao.ArticleDaoImpl.Updates(id, columns)
}

func (service *articleService) UpdateColumn(id int64, name string, value interface{}) error {
	return dao.ArticleDaoImpl.UpdateColumn(id, name, value)
}

func (service *articleService) Delete(id int64) (err error) {
	return dao.ArticleDaoImpl.Delete(id)
}

func (service *articleService) GetArticleTags(articleId int64) ([]*model.Tag, error) {
	articleTags, err := dao.ArticleTagDaoImpl.GetByArticleId(articleId)
	if err != nil {
		return make([]*model.Tag, 0), err
	}
	var tagIds []int64
	for _, articleTag := range articleTags {
		tagIds = append(tagIds, articleTag.Id)
	}
	return dao.TagDaoImpl.GetTagInIds(tagIds)
}

func (service *articleService) GetArticleInIds(articleIds []int64) ([]*model.Article, error) {
	if len(articleIds) == 0 {
		return make([]*model.Article, 0), nil
	}
	return service.QueryCondition(model.NewCondition("id in (?)", articleIds))
}

func (service *articleService) GetTagArticles(tagId int64, page int) ([]*model.Article, *model.Paging, error) {
	params := model.NewParams().Eq("tagId", tagId).Page(page, 20).Desc("id")
	articleTags, paging, err := dao.ArticleTagDaoImpl.QueryParams(params)
	if len(articleTags) == 0 || err != nil {
		return make([]*model.Article, 0), paging, err
	}
	var articleIds []int64
	for _, article := range articleTags {
		articleIds = append(articleIds, article.ArticleId)
	}
	articles, err := service.GetArticleInIds(articleIds)
	return articles, paging, err
}

func (service *articleService) Publish(userId int64, articleRequest *model.ArticleRequest) (*model.Article, error) {
	article := &model.Article{
		UserId:      userId,
		Title:       articleRequest.Title,
		Summary:     articleRequest.Summary,
		Content:     articleRequest.Content,
		CategoryId:  articleRequest.CategoryId,
		ContentType: model.ArticleContentTypeMarkdown,
		Status:      model.ArticleStatusPublished,
		Type:        model.ArticleTypeOriginal,
		Model: model.Model{
			CreateTime: extension.NowTimestamp(),
			UpdateTime: extension.NowTimestamp(),
		},
	}
	err := database.Tx(database.GetDB(), func(db *gorm.DB) error {
		tagIds, err := dao.TagDaoImpl.GetOrCreates(articleRequest.Tags)
		if err != nil {
			return err
		}
		err = service.Create(article)
		if err != nil {
			return err
		}
		err = dao.ArticleTagDaoImpl.AddArticleTags(article.Id, tagIds)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return article, err
	}
	return article, nil
}
