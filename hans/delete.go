package hans

import (
	"fmt"
	u "jin_quickly/models"
	"jin_quickly/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user u.User
	if err := utils.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "用户不存在"})
		return
	}
	// 删除用户
	if err := utils.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "删除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})
	fmt.Println("删除用户:", user)
}
