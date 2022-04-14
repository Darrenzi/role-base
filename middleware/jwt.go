package middleware

import (
	"blog/common/global"
	common "blog/common/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		Authorization := c.GetHeader("Authorization")
		if Authorization == "" {
			resp := common.NewResponse(c)
			resp.UnauthorizedAuthFail()
			c.Abort()
			return
		}

		claims, err := global.ParseToken(Authorization)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				common.JSON(http.StatusUnauthorized, common.ERROR_AUTH_CHECK_TOKEN_TIMEOUT, c)
				c.Abort()
				return
			default:
				common.JSON(http.StatusUnauthorized, common.ERROR_AUTH_CHECK_TOKEN_FAIL, c)
				c.Abort()
				return
			}
		}

		c.Set("user", claims)
		c.Next()
	}
}
