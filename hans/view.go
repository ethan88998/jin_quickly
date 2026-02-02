package hans

import (
	u "jin_quickly/models"
	"jin_quickly/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetViewUser(c *gin.Context) {
	id := c.Query("id")
	var user u.User
	if err := utils.DB.First(&user, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "login.html", gin.H{"message": "没找到用户"})
		return
	}
	c.HTML(http.StatusOK, "view.html", gin.H{"user": user})
	//fmt.Println("look:", user)
}
