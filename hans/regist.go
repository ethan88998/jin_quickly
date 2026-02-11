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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
	}

	// 注册成功生成JWT
	token, err := utils.GenToken(
		newUser.ID,
		newUser.Username,
		newUser.Age,
		newUser.Email,
	)

	if err != nil {
		c.String(http.StatusInternalServerError, "token 生成失败")
		return
	}

	// 生成token
	c.SetCookie("token", token, 3600, "/", "", false, true)

	// 进入后台用户列表
	c.Redirect(http.StatusFound, "/admin/user")

}
