/**
 * @Author: huangw1
 * @Date: 2019/7/31 15:58
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/huangw1/bbs/controller/render"
	"github.com/huangw1/bbs/model"
	"github.com/huangw1/bbs/pkg/extension"
	"github.com/huangw1/bbs/pkg/template"
	"github.com/huangw1/bbs/service"
)

func RegisterArticle(g *gin.Engine) {
	g.GET("/article/:articleId", articleDetail)
}

func articleDetail(c *gin.Context) {
	articleId := extension.MustInt64(c.Param("articleId"))
	article, err := service.ArticleService.Get(articleId)
	if err == nil && article.Status != model.ArticleStatusPublished {
		template.HTML(c, "404", nil)
		return
	}
	template.HTML(c, "article/detail", gin.H{
		"Article": render.BuildArticle(article),
	})
}
