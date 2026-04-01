package hans

import (
	"fmt"
	u "jin_quickly/models"
	"jin_quickly/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUserPage(c *gin.Context) {
	c.HTML(http.StatusOK, "useradd.html", nil)
}

func AddUserApi(c *gin.Context) {
	fmt.Println(">>> AddUserApi called")

	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Age      int    `json:"age" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": "参数错误，不能为空"})
		return
	}
	//
	//if req.Username == "" || req.Password == "" {
	//	c.JSON(400, gin.H{"code": 400, "msg": "用户名和密码不能为空"})
	//	return
	//}

	var count int64
	utils.DB.Model(&u.User{}).
		Where("username = ?", req.Username).
		Count(&count)

	if count > 0 {
		c.JSON(400, gin.H{"code": 400, "msg": "用户名已存在"})
		return
	}

	// 把临时接收的数据映射到数据库模型，也就是准备写入数据库
	user := u.User{
		Username: req.Username,
		Password: req.Password, // 下一步我们可以加 bcrypt
		Age:      req.Age,
		Email:    req.Email,
	}

	if err := utils.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": 200, "msg": "创建成功", "username": user.Username})
}
