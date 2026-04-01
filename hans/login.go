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
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
		return
	}

	var loginUser u.User
	utils.DB.First(&loginUser, "username = ?", req.Username)

	if loginUser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户不存在"})
		return
	}

	if loginUser.Password != req.Password {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户名或密码错误"})
		return
	}

	if loginUser.Status == 0 {
		c.JSON(http.StatusForbidden, gin.H{"msg": "账号已被禁用，请联系客服"})
		return
	}

	// 登录成功，生成JWT
	token, err := utils.GenToken(loginUser.ID, loginUser.Username, loginUser.Age, loginUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "生成token失败"})
		return
	}

	// 写入 cookie
	c.SetCookie("token", token, 3600*24, "/", "", false, true)

	// 返回 JSON
	c.JSON(http.StatusOK, gin.H{"msg": "登录成功", "token": token})
}
