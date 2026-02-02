package hans

import (
	u "jin_quickly/models"
	"jin_quickly/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录界面
func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var loginUser u.User
	utils.DB.First(&loginUser, "username = ?", username)
	if loginUser.ID == 0 {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "用户不存在"})
		return
	} else {
		if password == loginUser.Password {
			c.HTML(http.StatusOK, "userpage.html", gin.H{
				"username": loginUser.Username,
			})
		} else {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "用户名密码错误"})
		}
	}
}
