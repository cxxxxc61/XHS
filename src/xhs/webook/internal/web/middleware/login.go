package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddlewareBuild struct {
	path []string
}

func NewLoginMiddlewareBuild() *LoginMiddlewareBuild {
	return &LoginMiddlewareBuild{}
}

func (LMB *LoginMiddlewareBuild) Ignorepath(path string) *LoginMiddlewareBuild {
	LMB.path = append(LMB.path, path)
	return LMB
}

func (LMB *LoginMiddlewareBuild) Build() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if c.Request.URL.Path == "/users/login" ||
		//	c.Request.URL.Path == "/users/signup" {
		//	return
		//}
		//
		for _, path := range LMB.path {
			if path == c.Request.URL.Path {
				return
			}
		}
		session := sessions.Default(c)
		id := session.Get("userId")
		if id == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		updatatime := session.Get("update_Time")
		now := time.Now().Second()
		if updatatime == nil {
			session.Set("update_Time", now)
			session.Save()
			return
		}
		////updatetimeval, _ := updatatime.(int)
		//if now-updatetimeval > 60 {
		//	session.Set("update_Time", now)
		//	session.Save()
		//	return
		//}
		return
	}

}
