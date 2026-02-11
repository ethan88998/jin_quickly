// search.go

package hans

import (
	"jin_quickly/models"
	u "jin_quickly/models"
	"jin_quickly/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询单条用户名数据
func SearchUser(c *gin.Context) {
	username := c.Query("username")

	var user u.User
	if utils.DB.Model(&models.User{}).Where("username = ?", username).First(&user).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
	})
}

// 最基础：用户名 + 年龄
func SearchUsers(c *gin.Context) {
	username := c.Query("username")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "username is required"})
		return
	}

	var users []models.User

	err := utils.DB.
		Where("username LIKE ?", "%"+username+"%").
		Order("id desc").
		Find(&users).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "db error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    users,
		"total":   len(users),
	})
}

func SearchUserApi(c *gin.Context) {
	username := c.Query("username")
	ageStr := c.Query("age")

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "age must bu number"})
		return
	}

	var users []models.User
	err = utils.DB.
		Where("username = ? AND age = ?", username, age).
		Find(&users).
		Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "db error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    users,
		"total":   len(users),
	})

}

// 实用型：用户名模糊 + 年龄精确
func SearchUserapi(c *gin.Context) {
	username := c.Query("username")
	ageStr := c.Query("age")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	var users []models.User
	db := utils.DB.Model(&models.User{})

	if username != "" {
		db = db.Where("username like ?", "%"+username+"%")
	}

	if ageStr != "" {
		if age, err := strconv.Atoi(ageStr); err == nil {
			db = db.Where("age = ?", age)
		}
	}

	if startDate != "" {
		db = db.Where("created_at >= ?", startDate+" 00:00:00")
	}

	if endDate != "" {
		db = db.Where("created_at <= ?", endDate+" 23:59:59")
	}

	err := db.Order("age desc").Find(&users).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "db error"})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    users,
		"total":   len(users),
	})
}
