/*
*
@author: Azir
@desc:
@date: 6/19/24
*
*/
package middleware

import (
	"backend/common"
	"backend/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "没有权限",
			})
			c.Abort()
			return
		}
		if tokenString == "" || len(tokenString) < 7 || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.Where("id = ?", userId).First(&user)
		// 将用户信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}
