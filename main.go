package main

import "github.com/cxxxxc61/study/webook/internal/web"

func main() {
	sever := web.Registerroutes()
	sever.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务

	text()

	//sever := gin.Default()
	//u := &web.UserHandler{}
	//sever.POST("/users/login", func(c *gin.Context) {
	//	type Signupreq struct {
	//		Email           string `json:"email"`
	//		ConfirmPassword string `json:"confirm_password"`
	//		Password        string `json:"password"`
	//	}
	//
	//	var req Signupreq
	//	if err := c.Bind(&req); err != nil {
	//		return
	//	}
	//	c.String(http.StatusOK, "注册成功")
	//})
	//sever.POST("/post", u.Post)
	//sever.POST("/users/signup", u.Signup)
	//sever.POST("/login", u.Login)
	//sever.POST("/edit", u.Profile)
	//sever.GET("/profile", u.Profile)
	//sever.Run(":8080")
}
