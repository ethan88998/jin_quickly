package middlewares

import (
	"jin_quickly/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//
//func JWTAuth() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// 1. 取 token
//		tokenStr, err := c.Cookie("token")
//		if err != nil {
//			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
//				"message": "未登录",
//			})
//			return
//		}
//
//		claims := &utils.MyClaims{}
//		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
//			return utils.JwtKey, nil
//		})
//		if err != nil || !token.Valid {
//			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
//				"message": "324",
//			})
//		}

//	}
//}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1.取token
		tokenStr, err := c.Cookie("token")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// 2.解析并验证 token
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			// token过期/ 被篡改/ 无效
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// 3.校验通过后，立刻Set
		c.Set("user", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("age", claims.Age)
		c.Set("email", claims.Email)

		// 4.放行
		c.Next()
	}
}
