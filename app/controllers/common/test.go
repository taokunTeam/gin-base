package common

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/taokunTeam/gin-base/global"
	"github.com/taokunTeam/gin-base/utils"
)

func TestSql(ctx *gin.Context) {
	utils.ServeWs(global.App.Hub, ctx.Writer, ctx.Request)
}

func TestWs(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("userId"))
	toUserId, _ := strconv.Atoi(ctx.Param("toUserId"))
	utils.ServeSingleWs(global.App.Single, userId, toUserId, ctx.Writer, ctx.Request)
}
