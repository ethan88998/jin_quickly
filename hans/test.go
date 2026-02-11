package hans

import (
	"jin_quickly/models"
	"jin_quickly/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserListApii(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if page < 1 {
		page = 1
	}

	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	var users []models.User
	var total int64

	utils.DB.Model(&models.User{}).Count(&total)
	utils.DB.
		Limit(pageSize).
		Offset(offset).
		Order("id desc").
		Find(&users)

	c.JSON(200, gin.H{
		"total":    total,
		"users":    users,
		"page":     page,
		"pageSize": pageSize,
	})
}
