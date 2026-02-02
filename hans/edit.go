package hans

import (
	"fmt"
	u "jin_quickly/models"
	"jin_quickly/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取用户编辑页面的数据
func GetEditUserPage(c *gin.Context) {
	id := c.DefaultQuery("id", "0")

	var user u.User
	if err := utils.DB.First(&user, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "userpage.html", gin.H{"message": "用户不存在"})
		return
	}
	c.HTML(http.StatusOK, "edit.html", gin.H{"user": user})
	fmt.Println("---编辑结果查询：---", user)
}

func ShowEditUserPage(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	var user u.User
	if err := utils.DB.First(&user, id).Error; err != nil {
		c.HTML(http.StatusOK, "edit.html", gin.H{"user": user})
	}
}
