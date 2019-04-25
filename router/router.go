package router

import "github.com/gin-gonic/gin"

// Init is 初始化路由
func Init() {
	router := gin.Default()
	// CrossDomain跨域处理，options请求处理
	//router.Use(middleware.CrossDomain())

	v1 := router.Group("/v1")
	{
		v1.GET("/login", login.Login)
	}
	router.Run(":8000")
}
