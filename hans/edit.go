package hans

import (
	u "jin_quickly/models"
	"jin_quickly/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 页面：只返回 HTML
func UserDetailPage(c *gin.Context) {
	c.HTML(http.StatusOK, "user_detail.html", nil)
}

// API：只返回 JSON
func UserDetailApi(c *gin.Context) {
	id := c.Query("id")

	var user u.User
	if err := utils.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "用户不存在",
		})
		return
	}

	user.Password = ""

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": user,
	})
}

func UpdateUserApi(c *gin.Context) {
	id := c.Query("id")

	var user u.User
	if err := utils.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"code": 404, "msg": "用户不存在"})
		return
	}

	var req struct {
		Username string `json:"username" binding:"required"`
		Age      int    `json:"age" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"code": 400, "msg": "参数错误"})
		return
	}

	user.Username = req.Username
	user.Age = req.Age
	user.Email = req.Email
	if err := utils.DB.Save(&user).Error; err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	user.Password = ""
	c.JSON(200, gin.H{"code": 200, "msg": "保存成功", "data": user})
}
