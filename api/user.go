package api

import (
	common "blog/common/model"
	"blog/model"
	"blog/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var userService service.UserService

func AddUser(c *gin.Context) {
	var user model.User
	c.Bind(&user)
	resp := common.NewResponse(c)
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		resp.InvalidWithDesc(err.Error())
		return
	}

	roles := c.PostFormArray("roles[]")

	if len(roles) == 0 {
		resp.InvalidWithDesc("用户角色不能为空")
	}
	id, err := userService.Add(&user, roles)

	if err != nil {
		resp.InvalidWithDesc(err.Error())
		return
	}

	resp.Success(id)

}

func DeleteUser(c *gin.Context) {
	var user model.User
	c.Bind(&user)

	_, err := userService.Delete(&user)

	resp := common.NewResponse(c)
	if err != nil {
		resp.Error()
		return
	}

	resp.Success(nil)

}

func FindUsers(c *gin.Context) {
	var user model.User
	c.Bind(&user)

	users, err := userService.Find(&user)
	resp := common.NewResponse(c)
	if err != nil {
		resp.Error()
		return
	}

	resp.Success(users)
}

func GetRoles(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Query("ID"))

	roles, err := userService.Roles(ID)

	resp := common.NewResponse(c)
	if err != nil {
		resp.Error()
		return
	}

	resp.Success(roles)
}

func Authorize(c *gin.Context) {
	userID, _ := strconv.Atoi(c.PostForm("userID"))
	roles := c.PostFormArray("roles[]")
	fmt.Println(roles)
	_, err := userService.Authorize(userID, roles)

	resp := common.NewResponse(c)
	if err != nil {
		resp.Error()
		return
	}

	resp.Success(nil)
}

func Login(c *gin.Context) {
	var form model.LoginForm

	c.Bind(&form)

	token, err := userService.Login(form)
	if err != nil {
		c.JSON(http.StatusUnauthorized, common.Response{
			Code: common.UnauthorizedAuthFail,
			Desc: err.Error(),
			Data: nil,
		})
		return
	}
	resp := common.NewResponse(c)
	resp.Success(token)
}
