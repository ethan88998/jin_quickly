package hans

import (
	"jin_quickly/models"
	"jin_quickly/utils"

	"github.com/gin-gonic/gin"
)

func UserStat(c *gin.Context) {
	var total int64
	var today int64

	// 用户总数
	if err := utils.DB.Model(&models.User{}).Count(&total).Error; err != nil {
		c.JSON(500, gin.H{"error": "总统计失败"})
		return
	}

	// 今日统计
	if err := utils.DB.Model(&models.User{}).
		Where("DATE(created_at) = CURDATE()").
		Count(&today).Error; err != nil {
		c.JSON(500, gin.H{"error": "统计失败"})
		return
	}

	c.JSON(200, gin.H{
		"total": total,
		"today": today,
	})
}
