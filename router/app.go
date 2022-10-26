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
	//文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) //swagger文档路由
	//问题
	r.GET("/problemlist", service.GetProblemList)     //获取问题页面
	r.GET("/problemdetail", service.GetProblemDetail) //获取问题详情
	//用户
	r.GET("/userdetail", service.GetUserDetail) //获取用户详情页面
	r.POST("/login", service.Login)             //用户登录
	//提交记录
	r.GET("/submitlist", service.GetSubmitList) //提交记录数据
	return r
}
