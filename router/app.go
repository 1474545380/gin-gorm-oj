package router

import (
	"gin-gorm-oj/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	//路由
	r.GET("/ping", service.Ping)
	r.GET("/getproblemlist", service.GetProblemList)
	return r
}
