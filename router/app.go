package router

import (
	_ "gin-gorm-oj/docs"
	"gin-gorm-oj/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	//路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) //swagger文档路由
	r.GET("/ping", service.Ping)                                         //test
	r.GET("/problemlist", service.GetProblemList)                        //获取问题页面
	r.GET("/problemdetail", service.GetProblemDetail)                    //获取问题详情

	return r
}
