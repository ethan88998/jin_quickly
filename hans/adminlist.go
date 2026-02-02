package hans

import (
	u "jin_quickly/models"
	"jin_quickly/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Adminlist(c *gin.Context) {
	var users []u.User

	// 查询所有用户
	if err := utils.DB.Find(&users).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "userpage.html", gin.H{"error": "查询用户失败"})
		return
	}
	c.HTML(http.StatusOK, "userpage.html", gin.H{"users": users})

}
