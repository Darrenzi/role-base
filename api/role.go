package api

import (
	"blog/model"
	"blog/service"

	common "blog/common/model"

	"github.com/gin-gonic/gin"
)

var roleService service.RoleService

func SaveRole(c *gin.Context) {
	name := c.PostForm("name")
	id, err := roleService.SaveRole(&model.Role{Name: name})

	resp := common.NewResponse(c)
	if err != nil {
		resp.Error()
		return
	}

	resp.Success(id)
}
