package web

import (
	"github.com/gin-gonic/gin"
)

// 初始化

func Registerroutes() *gin.Engine {

	db := initdb()
	u := initUser(db)
	sever := initwebsever()

	registerusersroutes(sever, u)
	return sever
}

// 业务处理
func registerusersroutes(sever *gin.Engine, u *UserHandler) {
	s := sever.Group("/users")
	s.POST("/signup", u.Signup)
	s.POST("/post", u.Post)
	s.POST("/login", u.Login)
	s.POST("/edit", u.Profile)
	s.GET("/profile", u.Profile)
}
