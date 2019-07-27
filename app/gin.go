/**
 * @Author: huangw1
 * @Date: 2019/7/27 14:16
 */

package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huangw1/bbs/pkg/config"
	"github.com/huangw1/bbs/pkg/interior"
	"github.com/huangw1/bbs/pkg/session"
	"github.com/huangw1/bbs/pkg/template"
	"github.com/huangw1/bbs/router"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
)

const EnvProd = "prod"

func StartGin() {
	if config.Conf.Env != EnvProd {
		gin.SetMode(gin.DebugMode)
	}

	g := gin.New()

	extendViewFunc()
	router.InitRouter(g)
	g.Use(session.InitSession())

	g.Static("/static", "website/static")

	g.NoRoute(func(c *gin.Context) {
		notFound := regexp.MustCompile(`\.\w+$`)
		if notFound.MatchString(c.Request.URL.Path) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		template.HTML(c, "404", nil)
	})

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", config.Conf.Port),
		Handler: g,
	}

	go func() {
		logrus.Info("server run at port ", fmt.Sprintf(":%d", config.Conf.Port))
		logrus.Error("", srv.ListenAndServe())
	}()

	graceShutdown(srv)
}

func graceShutdown(srv http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	killSignal := <-quit
	switch killSignal {
	case os.Interrupt:
		logrus.Error("INTERRUPT...")
	case syscall.SIGTERM:
		logrus.Error("SIGTERM...")
	case syscall.SIGKILL:
		logrus.Error("SIGKILL...")
	}

	srv.Shutdown(context.Background())
	os.Exit(1)
}

func extendViewFunc() {
	template.AddFunc("siteTitle", func(title string) string {
		// todo
		return "golang bbs"
	})
	template.AddFunc("baseUrl", interior.BuildBaseUrl)

	template.AddFunc("articleUrl", interior.BuildArticleUrl)
	template.AddFunc("articlesUrl", interior.BuildArticlesUrl)
	template.AddFunc("tagArticlesUrl", interior.BuildTagArticlesUrl)
	template.AddFunc("categoryArticlesUrl", interior.BuildCategoryArticlesUrl)

	template.AddFunc("topicUrl", interior.BuildTopicUrl)
	template.AddFunc("topicsUrl", interior.BuildTopicsUrl)
	template.AddFunc("topicsUrl", interior.BuildTopicsUrl)
	template.AddFunc("tagTopicsUrl", interior.BuildTagTopicsUrl)

	template.AddFunc("userUrl", interior.BuildUserUrl)
	template.AddFunc("userArticlesUrl", interior.BuildUserArticlesUrl)
	template.AddFunc("userTopicsUrl", interior.BuildUserTopicsUrl)
	template.AddFunc("userFavoritesUrl", interior.BuildUserFavoritesUrl)
}
