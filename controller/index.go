/**
 * @Author: huangw1
 * @Date: 2019/7/29 14:39
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/huangw1/bbs/controller/render"
	"github.com/huangw1/bbs/model"
	"github.com/huangw1/bbs/pkg/template"
	"github.com/huangw1/bbs/service"
	"github.com/sirupsen/logrus"
)

func RegisterIndex(g *gin.Engine) {
	g.GET("/", index)
}

func index(c *gin.Context) {
	articles, err := service.ArticleService.QueryCondition(model.NewCondition("status = ?", model.ArticleStatusPublished))
	if err != nil {
		logrus.Errorf("controller index articles failed: %s", err.Error())
	}
	userIds, err := service.UserService.GetActiveUserIds()
	if err != nil {
		logrus.Errorf("controller index activeUsers failed: %s", err.Error())
	}
	template.HTML(c, "index", gin.H{
		"Articles":    render.BuildArticles(articles),
		"ActiveUsers": render.BuildUserInIds(userIds),
	})
}
