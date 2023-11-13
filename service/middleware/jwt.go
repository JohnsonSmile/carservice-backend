package middleware

import (
	"carservice/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// socialToken on header
		token := c.Request.Header.Get("Authorization")
		if len(strings.Fields(token)) < 2 {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "please login",
			})
			c.Abort()
			return
		}
		token = strings.Fields(token)[1]
		zap.L().Debug("token", zap.String("token", token))
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "please login",
			})
			c.Abort()
			return
		}
		j := util.NewJWT()
		claims, err := j.ParseToken(token)
		zap.L().Debug("claims", zap.Any("claims", claims))
		if err != nil {
			if err == util.ErrorTokenExpired {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"msg": "Authorization token expired",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, "not login")
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Set("userId", claims.ID)
		zap.L().Debug("pass the userId", zap.Uint("userId", claims.ID))
		c.Next()
	}
}

func JWTWSAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// socialToken on header
		token := c.Query("token")
		zap.L().Debug("token", zap.String("token", token))
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "please login",
			})
			c.Abort()
			return
		}
		j := util.NewJWT()
		claims, err := j.ParseToken(token)
		zap.L().Debug("claims", zap.Any("claims", claims))
		if err != nil {
			if err == util.ErrorTokenExpired {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"msg": "Authorization token expired",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, "not login")
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Set("userId", claims.ID)
		zap.L().Debug("pass the userId", zap.Uint("userId", claims.ID))
		c.Next()
	}
}
