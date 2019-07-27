/**
 * @Author: huangw1
 * @Date: 2019/7/27 14:16
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"regexp"
	"time"
)

var ignoreRegexp = regexp.MustCompile(`.css|.js|.jpeg|.jpg|.png|.ico`)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		if ignoreRegexp.MatchString(c.Request.URL.Path) {
			c.Next()
			return
		}
		start := time.Now().UTC()
		path := c.Request.URL.Path
		method := c.Request.Method
		ip := c.ClientIP()
		c.Next()
		latency := time.Now().UTC().Sub(start)
		logrus.Infof("%s %s: cost %s, ip %s", method, path, latency, ip)
	}
}
