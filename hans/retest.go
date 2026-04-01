package hans

import (
	"jin_quickly/models"
	"jin_quickly/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 用户数据列表+分页
func ReAdminList(c *gin.Context) {
	// 1.读取查询参数
	username := c.Query("username")
	ageStr := c.Query("age")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	// 2.读取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	// 3.验证分页数据合法性
	if page < 0 {
		page = 1
	}

	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// 4.构建DB

	db := utils.DB.Model(&models.User{})

	// 5.判断用户名/年龄
	if username != "" {
		db = db.Where("username LIKE ?", "%"+username+"%")
	}

	if ageStr != "" {
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "年龄必须是数字"})
		}
		db = db.Where("age >= ?", age)
	}

	if startDate != "" {
		db = db.Where("startDate >= ?", startDate+" 00:00:00")
	}

	if endDate != "" {
		db = db.Where("endDate >= ?", endDate+" 23:59:59")
	}

	// 6.统计
	var total int64
	db.Count(&total)

	// 7.分页查询
	offset := (page - 1) * pageSize
	var users []models.User
	err := db.
		Order("create_time desc").
		Limit(pageSize).
		Offset(offset).
		Find(&users).Error

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 8.返回json
	c.JSON(200, gin.H{
		"users": users,
		"total": total,
		"page":  page,
	})
}

// 状态接口实现
type ReChangeStatusReq struct {
	ID     int64 `json:"id" binding:"required, gt=0"`
	Status int   `json:"status"`
}

func ReChangeStatus(c *gin.Context) {
	// 1.接收数据
	var req ReChangeStatusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
		return
	}

	// 2.校验状态合法
	if req.Status != 0 && req.Status != 1 {
		c.JSON(400, gin.H{"error": "非法状态"})
		return
	}

	// 3.构建DB，SQL
	db := utils.DB.Model(&models.User{})

	err := db.Where("id=?", req.ID).
		Update("status", req.Status).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "状态更新失败"})
		return
	}

	// 返回json
	c.JSON(200, gin.H{
		"status": req.Status,
		"id":     req.ID,
		"msg":    "状态更新成功",
	})

}

// 删除接口
func ReDelete(c *gin.Context) {
	// 1.读取删除ID
	id := c.Query("id")

	// 2.创建变量 存储数据库返回的数据
	var user models.User

	// 3.校验ID是否存在
	if err := utils.DB.Where("id=?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户不存在"})
		return
	}

	// 4.删除操作
	if err := utils.DB.Delete(&user).Error; err != nil {
		c.JSON(400, gin.H{"msg": "删除失败"})
		return
	}

	// 5.返回json
	c.JSON(200, gin.H{
		"code":     200,
		"msg":      "删除成功",
		"id":       id,
		"username": user.Username,
	})

}

// 新增用户接口
func ReAddUserT(c *gin.Context) {

	// 1.创建变更接收数据
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Age      int    `json:"age" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}

	// 2.接收并解析数据
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// 3.变更count
	var count int64
	utils.DB.Where("username = ?", req.Username).
		Count(&count)

	if count > 0 {
		c.JSON(400, gin.H{"msg": "用户已存在"})
		return
	}

	// 4.把临时存储数据映射到数据库表
	user := models.User{
		Username: req.Username,
		Password: req.Password,
		Age:      req.Age,
		Email:    req.Email,
	}

	// 5.处理创建报错
	if err := utils.DB.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	// 6.创建成功返回json
	c.JSON(200, gin.H{
		"code":     200,
		"username": user.Username,
		"msg":      "success",
	})
}

// 用户统计接口
func ReTatol(c *gin.Context) {
	var tatal int64
	var today int64

	// 用户总数
	err := utils.DB.Model(&models.User{}).Count(&tatal).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "用户总数统计失败"})
		return
	}

	// 今日用户
	if err := utils.DB.Model(&models.User{}).
		Where("DATE(created_at) = CURDATE()").
		Count(&today).Error; err != nil {
		c.JSON(400, gin.H{"error": "今日统计失败"})
		return
	}

	// 返回json
	c.JSON(200, gin.H{
		"code":  200,
		"today": today,
		"total": tatal,
	})

}

// 编辑接口
func ReEditUser(c *gin.Context) {
	// 1.接收ID
	id := c.Query("id")

	// 2.通过ID查询是否存在
	var user models.User
	err := utils.DB.First(&models.User{}, id).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "用户不存在"})
		return
	}

	user.Password = ""

	// 返回json
	c.JSON(200, gin.H{
		"code": 200,
		"user": user,
	})
}

// 更新
func ReUpdateUser(c *gin.Context) {
	id := c.Query("id")

	var user models.User
	err := utils.DB.First(&models.User{}, id).Error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 接收数据
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Age      int    `json:"age" binding:"required"`
	}

	// 解析
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	user.Username = req.Username
	user.Password = req.Password
	user.Age = req.Age

	if err := utils.DB.Save(&user).Error; err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	user.Password = ""

	c.JSON(200, gin.H{
		"code": 200,
		"user": user,
	})

}
