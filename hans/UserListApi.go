package hans

//
//import (
//	"jin_quickly/models"
//	"jin_quickly/utils"
//	"net/http"
//	"strconv"
//
//	"github.com/gin-gonic/gin"
//)
//
//func UserPageAApi(c *gin.Context) {
//	// 默认参数
//	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
//	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
//
//	if page < 1 {
//		page = 1
//	}
//	if pageSize <= 0 || pageSize > 50 {
//		pageSize = 10
//	}
//
//	offset := (page - 1) * pageSize
//
//	var users []models.User
//	var total int64
//
//	// 总数
//	utils.DB.Model(&models.User{}).Count(&total)
//
//	// 分页数据
//	utils.DB.
//		Offset(offset).
//		Limit(pageSize).
//		Order("id desc").
//		Find(&users)
//
//	c.JSON(http.StatusOK, gin.H{
//		"code":     200,
//		"data":     users,
//		"page":     page,
//		"pageSize": pageSize,
//		"total":    total,
//	})
//}
