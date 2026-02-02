package hans

import (
	"fmt"
	u "jin_quickly/models"
	"jin_quickly/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册界面
func ShowRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

// 提交注册
func Register(c *gin.Context) {
	var req struct {
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
		Age      int    `form:"age" json:"age"`
		Email    string `form:"email" json:"email"`
	}

	if err := c.ShouldBind(&req); err != nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "参数错误"})
		return
	}

	// 查重
	var user u.User
	if err := utils.DB.Where("username = ?", req.Username).First(&user).Error; err == nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "用户已存在"})
		return
	}

	// 创建用户
	newUser := u.User{
		Username: req.Username,
		Password: req.Password,
		Age:      req.Age,
		Email:    req.Email,
	}
	if err := utils.DB.Create(&newUser).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "注册失败！"})
		return
	}
	c.HTML(http.StatusOK, "login.html", gin.H{
		"success":  "注册成功，请登录！",
		"username": newUser.Username,
	})
	fmt.Println("newUser:", newUser)
}
