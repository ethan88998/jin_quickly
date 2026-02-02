package hans

import (
	u "jin_quickly/models"
	"jin_quickly/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type IntString int

func (i *IntString) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	v, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*i = IntString(v)
	return nil
}

// 更新用户数据
func UpdateUser(c *gin.Context) {
	var req struct {
		Username string    `form:"username" json:"username"`
		Password string    `form:"password" json:"password"`
		Age      IntString `json:"age"`
		Email    string    `form:"email" json:"email"`
	}

	// 获取前端传递的数据
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	id := c.Param("id")

	var user u.User
	if err := utils.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
		return
	}

	// 更新用户信息
	user.Username = req.Username
	user.Password = req.Password
	user.Age = int(req.Age)
	user.Email = req.Email
	if err := utils.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "用户更新成功"})
	//c.HTML(http.StatusOK, "login.html", gin.H{"message": "更新成功！"})
}
