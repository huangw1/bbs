/**
 * @Author: huangw1
 * @Date: 2019/7/25 14:51
 */

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/huangw1/bbs/database"
	"github.com/huangw1/bbs/utils/config"
	_ "github.com/huangw1/bbs/utils/process"
	"github.com/huangw1/bbs/utils/template"
	"net/http"
)

func main() {
	config.InitConfig("bbs.yaml")

	database.OpenDB(&database.DBConfig{Dialect: "mysql", Url: config.Conf.MySqlUrl})
	defer database.CloseDB()

	template.InitTemplate(&template.Config{})

	g := gin.New()
	g.NoRoute(func(c *gin.Context) {
		template.HTML(c, "index", nil)
	})
	http.ListenAndServe(":8080", g)
}
