package hans

import (
	u "jin_quickly/models"
	"jin_quickly/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetViewUserPage(c *gin.Context) {
	c.HTML(http.StatusOK, "view.html", nil)

}

func GetViewUser(c *gin.Context) {
	id := c.Query("id")
	var user u.User
	if err := utils.DB.First(&user, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "login.html", gin.H{"message": "没找到用户"})
		return
	}
	//c.HTML(http.StatusOK, "view.html", gin.H{"user": user})
	c.JSON(http.StatusOK, gin.H{
		"data": user,
		"code": 200,
		"msg":  "ok",
	})

}
