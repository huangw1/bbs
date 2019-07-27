/**
 * @Author: huangw1
 * @Date: 2019/7/25 14:51
 */

package main

import (
	"github.com/huangw1/bbs/app"
	"github.com/huangw1/bbs/database"
	"github.com/huangw1/bbs/pkg/config"
	_ "github.com/huangw1/bbs/pkg/process"
	"github.com/huangw1/bbs/pkg/template"
	"github.com/sirupsen/logrus"
)

func init() {
	initConfig()
	initLogrus()
	initTemplate()
	initDatabase()
}

func initConfig() {
	config.InitConfig("bbs.yaml")
}

func initLogrus() {
	// todo
	logrus.SetLevel(logrus.InfoLevel)
}

func initTemplate() {
	template.InitTemplate()
}

func initDatabase() {
	database.OpenDB(&database.DBConfig{
		Dialect:        "mysql",
		Url:            config.Conf.MySqlUrl,
		EnableLogModel: config.Conf.ShowSql,
	})
}

func main() {
	defer database.CloseDB()
	app.StartGin()
}
