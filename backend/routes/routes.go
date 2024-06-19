/*
*
@author: Azir
@desc:
@date: 6/19/24
*
*/
package routes

import (
	"backend/controller"
	middleware "backend/middle"

	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	// 允许跨域访问
	r.Use(middleware.CORSMiddleware())
	// 注册
	r.POST("/register", controller.Register)
	// 登录
	r.POST("/login", controller.Login)
	// 登录获取用户信息
	r.GET("/user", middleware.AuthMiddleware(), controller.GetInfo)

	return r
}
