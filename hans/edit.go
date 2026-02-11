package hans

import (
	u "jin_quickly/models"
	"jin_quickly/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取用户编辑页面的数据
func UserDetailPage(c *gin.Context) {
	c.HTML(http.StatusOK, "user_detail.html", nil)

}

func ShowEditUserPage(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	var user u.User
	if err := utils.DB.First(&user, id).Error; err != nil {
		c.HTML(http.StatusOK, "edit.html", gin.H{"user": user})
	}
}

// 获取用户数据信息
func UserDetailApi(c *gin.Context) {
	id := c.Query("id")

	var user u.User
	if err := utils.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "用户不存在"})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"data": user,
		"code": 200,
		"msg":  "ok",
	})
}

// 保存编辑数据
func UpdateUserApi(c *gin.Context) {
	id := c.Query("id")

	var user u.User
	if err := utils.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "用户不存在",
			"code": 404,
		})
		return
	}

	var req struct {
		Username string `json:"username"`
		Age      int    `json:"age"`
		Email    string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数错误2",
			"code": 400,
		})
		return
	}

	user.Username = req.Username
	user.Age = req.Age
	user.Email = req.Email

	utils.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "保存成功",
	})
}
