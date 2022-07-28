package main

import (
	"gin.test.com/src/utils"
	"github.com/gin-gonic/gin"
)

func main()  {
	var a = 1
	println(a)
	println("Hello world !!!!!")
	utils.GetToday()
	//Default返回一个默认的路由引擎
	//gin.SetMode(gin.ReleaseMode)  // set debug mod off
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}