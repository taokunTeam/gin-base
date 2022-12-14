package routes

import (
	"net/http"
	"time"

	"github.com/taokunTeam/gin-base/app/controllers/app"
	"github.com/taokunTeam/gin-base/app/controllers/common"
	"github.com/taokunTeam/gin-base/app/middleware"
	"github.com/taokunTeam/gin-base/app/services"

	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.GET("/test", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "success")
	})

	router.POST("/auth/register", app.Register)

	router.POST("/auth/login", app.Login)
	router.POST("/image_upload", common.ImageUpload)
	router.POST("/test", common.Test)
	router.POST("/sql", common.TestSql)
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
		//authRouter.POST("/image_upload", common.ImageUpload)
	}

	router.Any("/ws", common.TestSql)
	router.Any("/ws/:userId/:toUserId", common.TestWs)
}
