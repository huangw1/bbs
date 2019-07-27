/**
 * @Author: huangw1
 * @Date: 2019/7/27 17:07
 */

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/huangw1/bbs/controller"
	"github.com/huangw1/bbs/router/middleware"
)

func InitRouter(g *gin.Engine) {
	g.Use(gin.Recovery())
	g.Use(middleware.Logging())
	g.Use(middleware.Options())
	g.Use(middleware.RequestId())

	controller.RegisterIndex(g)
	controller.RegisterArticle(g)
}
