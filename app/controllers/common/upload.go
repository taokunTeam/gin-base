package common

import (
	"github.com/taokunTeam/gin-base/app/common/request"
	"github.com/taokunTeam/gin-base/app/common/response"
	"github.com/taokunTeam/gin-base/app/services"

	"github.com/gin-gonic/gin"
)

func ImageUpload(c *gin.Context) {
	var form request.ImageUpload
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	outPut, err := services.MediaService.SaveImage(form)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, outPut)
}
