/**
 * @Author: huangw1
 * @Date: 2019/7/26 00:56
 */

package template

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huangw1/bbs/pkg/extension"
	"github.com/sirupsen/logrus"
	template2 "html/template"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

type TempOption func(conf *Config)

type Config struct {
	ViewPath  string
	Extension string
	Layout    string
	FuncMaps  template2.FuncMap
}

var conf = &Config{
	ViewPath:  "website/views",
	Extension: ".html",
	Layout:    "common/layout",
	FuncMaps: template2.FuncMap{
		"formatDate": func(timestamp int64) string {
			return extension.TimeFormat(extension.TimeFromTimestamp(timestamp), extension.FmtDateTime)
		},
		"prettyTime": func(timestamp int64) string {
			return extension.PrettyTime(timestamp)
		},
	},
}

func InitTemplate(options ...TempOption) {
	if len(options) > 0 {
		for _, tempOption := range options {
			tempOption(conf)
		}
	}
}

func TempWithViewPath(viewPath string) TempOption {
	return func(conf *Config) {
		conf.ViewPath = viewPath
	}
}

func TempWithExtension(extension string) TempOption {
	return func(conf *Config) {
		conf.Extension = extension
	}
}

func TempWithLayout(layout string) TempOption {
	return func(conf *Config) {
		conf.Layout = layout
	}
}

func TempWithFuncMaps(funcMaps template2.FuncMap) TempOption {
	return func(conf *Config) {
		for name, fun := range funcMaps {
			conf.FuncMaps[name] = fun
		}
	}
}

func AddFunc(name string, fun interface{}) {
	conf.FuncMaps[name] = fun
}

func HTML(c *gin.Context, filename string, data map[string]interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	HTMLWithLayout(c, conf.Layout, filename, data)
}

func HTMLWithLayout(c *gin.Context, Layout, filename string, data map[string]interface{}) {
	files := fmt.Sprintf("%s,%s", Layout, filename)
	htmlFiles := strings.Split(files, ",")
	for i, file := range htmlFiles {
		htmlFiles[i] = filepath.Join(conf.ViewPath, file+conf.Extension)
	}
	temp, err := template2.New(filepath.Base(Layout + conf.Extension)).Funcs(conf.FuncMaps).Funcs(template2.FuncMap{"include": Include}).ParseFiles(htmlFiles...)
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
	temp, err := template2.New(filepath.Base(filename + conf.Extension)).Funcs(conf.FuncMaps).ParseFiles(filepath.Join(conf.ViewPath, filename+conf.Extension))
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
