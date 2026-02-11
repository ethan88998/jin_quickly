package hans

import (
	"jin_quickly/models"
	"jin_quickly/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//func Adminlist(c *gin.Context) {
//	var users []u.User
//
//	// 查询所有用户
//	if err := utils.DB.Find(&users).Error; err != nil {
//		c.HTML(http.StatusInternalServerError, "userlist.html", gin.H{"error": "查询用户失败"})
//		return
//	}
//	c.HTML(http.StatusOK, "userlist.html", gin.H{"users": users})
//
//}

func Adminlist(c *gin.Context) {
	// 1. 取当前登录用户
	uid, _ := c.Get("user")
	username, _ := c.Get("username")
	age, _ := c.Get("age")
	email, _ := c.Get("email")

	// 2.查账号列表
	var users []models.User
	utils.DB.Find(&users)

	// 3.渲染页面
	c.HTML(http.StatusOK, "userlist.html", gin.H{
		"users":      users,
		"login_user": username,
		"uid":        uid,
		"age":        age,
		"email":      email,
	})

}

func UserListPage(c *gin.Context) {
	username, _ := c.Get("username")
	email, _ := c.Get("email")
	c.HTML(http.StatusOK, "userall.html", gin.H{
		"username": username,
		"email":    email,
	})
}

//func UserListApi(c *gin.Context) {
//	// 1. 取当前登录用户
//	uid, _ := c.Get("user")
//	username, _ := c.Get("username")
//	age, _ := c.Get("age")
//	email, _ := c.Get("email")
//
//	var users []models.User
//	utils.DB.Find(&users)
//
//	c.JSON(http.StatusOK, gin.H{
//		"data":     users,
//		"uid":      uid,
//		"username": username,
//		"age":      age,
//		"email":    email,
//		"code":     200,
//		"msg":      "ok",
//	})
//}

func UserListApi(c *gin.Context) {
	username := c.Query("username")
	ageStr := c.Query("age")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if page < 1 {
		page = 1
	}

	var users []models.User
	var total int64

	db := utils.DB.Model(&models.User{})

	// 用户名模糊
	if username != "" {
		db = db.Where("username LIKE ?", "%"+username+"%")
	}

	// 年龄
	if ageStr != "" {
		if age, err := strconv.Atoi(ageStr); err == nil {
			db = db.Where("age = ?", age)
		}
	}

	// 注册时间
	if startDate != "" {
		db = db.Where("created_at >= ?", startDate+" 00:00:00")
	}
	if endDate != "" {
		db = db.Where("created_at <= ?", endDate+" 23:59:59")
	}

	// 先统计
	db.Count(&total)

	// 再分页查询
	offset := (page - 1) * pageSize
	err := db.
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&users).Error

	if err != nil {
		c.JSON(500, gin.H{"message": "db error"})
		return
	}

	c.JSON(200, gin.H{
		"list":  users,
		"total": total,
		"page":  page,
	})
}

func UserList(c *gin.Context) {
	// ========= 分页参数 =========
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// ========= 搜索参数 =========
	username := c.Query("username")
	ageStr := c.Query("age")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	db := utils.DB.Model(&models.User{})

	// 用户名模糊
	if username != "" {
		db = db.Where("username LIKE ?", "%"+username+"%")
	}

	// 年龄精确
	if ageStr != "" {
		if age, err := strconv.Atoi(ageStr); err == nil {
			db = db.Where("age = ?", age)
		}
	}

	// 注册时间范围
	if startDate != "" {
		t, _ := time.Parse("2006-01-02", startDate)
		db = db.Where("created_at >= ?", t)
	}

	if endDate != "" {
		t, _ := time.Parse("2006-01-02", endDate)
		db = db.Where("created_at <= ?", t.Add(23*time.Hour+59*time.Minute+59*time.Second))
	}

	// ========= 统计总数 =========
	var total int64
	if err := db.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "count failed",
		})
		return
	}

	// ========= 查询列表 =========
	var users []models.User
	err := db.
		Order("id DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&users).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "query failed",
		})
		return
	}

	// ========= 返回 =========
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"msg":      "success",
		"list":     users,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}
