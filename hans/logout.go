package hans

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 退出登录
func Logout(c *gin.Context) {
	// 清除cookie
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/login")
}
