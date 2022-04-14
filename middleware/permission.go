package middleware

import (
	"blog/common/global"
	"blog/common/model"
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		gormAdapter, _ := gormadapter.NewAdapterByDB(global.GetDB())
		enforcer, _ := casbin.NewEnforcer(global.Config.Casbin.Path, gormAdapter)
		enforcer.LoadPolicy()

		claims, _ := c.Get("user")
		user := claims.(*global.Claims)
		var pass bool
		var err error
		for _, role := range user.Roles {
			pass, err = enforcer.Enforce(role, c.Request.URL.Path, c.Request.Method)
			if !pass || err != nil {
				break
			}
		}

		resp := model.NewResponse(c)
		if err != nil {
			log.Warn("鉴权失败：", err)
			resp.ErrorWithDesc(fmt.Sprintf("鉴权失败: ", err))
			c.Abort()
			return
		}

		if !pass {
			resp.UnauthorizedAuthFail()
			c.Abort()
			return
		}

		c.Next()
	}
}
