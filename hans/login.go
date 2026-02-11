package hans

import (
	"fmt"
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
	}

	if password != loginUser.Password {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": "用户名或密码错误",
		})
		return
	}

	// 登录成功，生成JWT
	token, err := utils.GenToken(loginUser.ID, loginUser.Username, loginUser.Age, loginUser.Email)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"error": "生成token失败",
		})
		return
	}
	fmt.Println("===token===:", token)

	// JWT写入cookie
	c.SetCookie("token", token, 3600*24, "/", "", false, true)

	c.Redirect(http.StatusFound, "/admin/user")
	fmt.Println("===SetCookie===:", c.SetCookie)

}
