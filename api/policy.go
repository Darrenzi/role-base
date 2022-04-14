package api

import (
	common "blog/common/model"
	"blog/model"
	"blog/service"

	"github.com/gin-gonic/gin"
)

var policyService service.PolicyService

func AddPolicy(c *gin.Context) {
	var policy model.Policy
	c.Bind(&policy)

	err := policyService.Add(&policy)

	resp := common.NewResponse(c)
	if err != nil {
		resp.ErrorWithDesc(err.Error())
		return
	}

	resp.Success(nil)
}

func DeletePolicy(c *gin.Context) {
	var policy model.Policy
	c.Bind(&policy)

	err := policyService.Delete(&policy)

	resp := common.NewResponse(c)
	if err != nil {
		resp.ErrorWithDesc(err.Error())
		return
	}

	resp.Success(nil)
}
