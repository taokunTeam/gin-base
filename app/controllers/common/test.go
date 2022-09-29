package common

import (
	"github.com/gin-gonic/gin"
	"github.com/taokunTeam/gin-base/global"
	"github.com/taokunTeam/gin-base/utils"
)

func TestSql(ctx *gin.Context) {
	utils.ServeWs(global.App.Hub, ctx.Writer, ctx.Request)
}
