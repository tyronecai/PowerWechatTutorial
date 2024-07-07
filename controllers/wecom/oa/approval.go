package oa

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/approval/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func APIApprovalOaGetTemplateDetail(c *gin.Context) {

}

func APIApprovalOaGetApprovalInfo(c *gin.Context) {

}

func APIApprovalOaGetApprovalDetail(c *gin.Context) {

}

func APIApprovalOaGetApprovalData(c *gin.Context) {

}

func APIApprovalVacationGetCorpConf(c *gin.Context) {

}

func APIApprovalVacationGetUserVacationQuota(c *gin.Context) {

}

func APIApprovalVacationSetOneUserQuota(c *gin.Context) {

}

func APIApprovalUpdateTemplate(c *gin.Context) {
	options := &request.RequestUpdateTemplate{}
	res, err := services.WeComApp.OAApproval.UpdateTemplate(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
