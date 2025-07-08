package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	sever := gin.Default()

	sever.POST("/users/signup", func(c *gin.Context) {

	})

	sever.POST("/users/login", func(c *gin.Context) {

	})

	sever.POST("/users/edit", func(c *gin.Context) {

	})
	sever.PUT("/users/profile", func(c *gin.Context) {

	})

	sever.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
