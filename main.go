package main

import (
	"jin_quickly/hans"
	"jin_quickly/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//var err error

// 3.主函数注册路由
func main() {
	// 创建路由
	r := gin.Default()
	// 连接数据库
	utils.InitDB()
	// 加载模板
	r.LoadHTMLGlob("templates/*")

	// 注册/登录路由
	r.GET("/", hans.ShowRegister)
	r.POST("/register", hans.Register)

	r.GET("/login", hans.ShowLogin)
	r.POST("/login", hans.Adminlist)

	admin := r.Group("/admin")
	{
		// 用户管理
		//admin.Use(hans.Adminlist)
		// 执行后后返回默认页
		admin.GET("/users", hans.Adminlist)
		// 删除用户
		admin.DELETE("/user/:id", hans.DeleteUser)
		// 编辑用户
		admin.GET("/user/edit", hans.GetEditUserPage)
		// 更新用户数据
		admin.PUT("/users/edit/:id", hans.UpdateUser)
		// 查看数据
		admin.GET("/user/look", hans.GetViewUser)
	}

	r.Run(":8081")
}
