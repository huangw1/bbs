/**
 * @Author: huangw1
 * @Date: 2019/7/26 00:56
 */

package template

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huangw1/bbs/utils/extension"
	"github.com/sirupsen/logrus"
	template2 "html/template"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

var (
	ViewPath  = "website/views"
	Extension = ".html"
	Layout    = "common/layout"
	FuncMaps  = template2.FuncMap{
		"formatDate": func(timestamp int64) string {
			return extension.TimeFormat(extension.TimeFromTimestamp(timestamp), extension.FmtDateTime)
		},
	}
)

type Config struct {
	ViewPath  string
	Extension string
	Layout    string
	FuncMap   template2.FuncMap
}

func InitTemplate(conf *Config) {
	if conf.ViewPath != "" {
		ViewPath = conf.ViewPath
	}
	if conf.Extension != "" {
		Extension = conf.Extension
	}
	if conf.Layout != "" {
		Layout = conf.Layout
	}
	if conf.FuncMap != nil {
		for name, fun := range conf.FuncMap {
			FuncMaps[name] = fun
		}
	}
}

func HTML(c *gin.Context, filename string, data map[string]interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	HTMLWithLayout(c, Layout, filename, data)
}

func HTMLWithLayout(c *gin.Context, Layout, filename string, data map[string]interface{}) {
	files := fmt.Sprintf("%s,%s", Layout, filename)
	htmlFiles := strings.Split(files, ",")
	for i, file := range htmlFiles {
		htmlFiles[i] = filepath.Join(ViewPath, file+Extension)
	}
	temp, err := template2.New(filepath.Base(Layout + Extension)).Funcs(FuncMaps).Funcs(template2.FuncMap{"include": Include}).ParseFiles(htmlFiles...)
	if err != nil {
		logrus.Errorf("html failed: %s", err.Error())
	}
	var buf bytes.Buffer
	err = temp.Execute(&buf, data)
	if err != nil {
		logrus.Errorf("html execute failed: %s", err.Error())
	}
	sendHTML(c, http.StatusOK, buf.String())
}

func sendHTML(c *gin.Context, status int, content string) {
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.Writer.WriteHeader(status)
	io.WriteString(c.Writer, content)
}

func Include(filename string, data map[string]interface{}) template2.HTML {
	if data == nil {
		data = map[string]interface{}{}
	}
	temp, err := template2.New(filepath.Base(filename + Extension)).Funcs(FuncMaps).ParseFiles(filepath.Join(ViewPath, filename+Extension))
	if err != nil {
		logrus.Errorf("include failed: %s", err.Error())
	}
	var buf bytes.Buffer
	err = temp.Execute(&buf, data)
	if err != nil {
		logrus.Errorf("include execute failed: %s", err.Error())
	}
	return template2.HTML(buf.String())
}
