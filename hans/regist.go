package hans

import (
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
		Username string `json:"username"`
		Password string `json:"password"`
		Age      int    `json:"age"`
		Email    string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
		return
	}

	// 查重
	var user u.User
	if err := utils.DB.Where("username = ?", req.Username).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户已存在"})
		return
	}

	// 组装数据
	newUser := u.User{
		Username: req.Username,
		Password: req.Password,
		Age:      req.Age,
		Email:    req.Email,
		Status:   1, // 默认启用
	}

	if err := utils.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "注册失败"})
		return
	}

	// 生成 JWT
	token, err := utils.GenToken(newUser.ID, newUser.Username, newUser.Age, newUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "token生成失败"})
		return
	}

	// 写入 cookie
	c.SetCookie("token", token, 3600*24, "/", "", false, true)

	// 返回 JSON 成功
	c.JSON(http.StatusOK, gin.H{
		"msg":   "注册成功",
		"token": token,
		"user":  newUser,
	})
}
