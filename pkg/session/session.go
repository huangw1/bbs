/**
 * @Author: huangw1
 * @Date: 2019/7/25 16:34
 */

package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/huangw1/bbs/model"
)

const (
	UserKey      = "userId"
	StoreKey     = "cookie"
	CookieSecret = "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5"
)

func InitSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(CookieSecret))
	return sessions.Sessions(StoreKey, store)
}

func SaveSession(c *gin.Context, userId int64) {
	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
	})
	session.Set(UserKey, userId)
	session.Save()
}

func DeleteSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(UserKey)
	session.Save()
}

func GetCurrentUser(c *gin.Context) *model.User {
	session := sessions.Default(c)
	if userId := session.Get(UserKey); userId != nil {
		// todo get user info
	}
	return nil
}
