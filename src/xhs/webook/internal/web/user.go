package web

import (
	"github.com/gin-gonic/gin"
)

// 业务处理
func (u *UserHandler) Registerusersroutes(sever *gin.Engine) {
	s := sever.Group("/users")
	s.POST("/signup", u.Signup)
	s.POST("/post", u.Post)
	//s.POST("/login", u.Login)
	s.POST("/login", u.LoginJWT)
	//s.POST("/edit", u.Profile)
	s.GET("/profile", u.Profile)
	s.POST("login_sms/code/send", u.SendLoginsms)
	s.POST("login_sms/", u.Loginsms)
}
