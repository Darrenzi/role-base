package router

import (
	"blog/api"
	"blog/common/model"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		resp := model.NewResponse(c)
		resp.Success(nil)
	})

	r.POST("/login", api.Login)
	authApi := r.Group("/")
	authApi.Use(middleware.JWT())
	authApi.Use(middleware.CasbinMiddleware())
	{
		userApi := authApi.Group("/user")
		{
			userApi.POST("", api.AddUser)
			userApi.DELETE("", api.DeleteUser)
			userApi.GET("/condition", api.FindUsers)
			userApi.GET("/role", api.GetRoles)
			userApi.PUT("/role", api.Authorize)
		}

		roleAPi := authApi.Group("/role")
		{
			roleAPi.POST("", api.SaveRole)
		}

		policyApi := authApi.Group("/policy")
		{
			policyApi.POST("", api.AddPolicy)
			policyApi.DELETE("", api.DeletePolicy)
		}
	}
	return r
}
