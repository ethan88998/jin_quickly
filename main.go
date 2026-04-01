package main

import (
	"jin_quickly/hans"
	"jin_quickly/middlewares"
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
	r.Static("/static", "./static")

	// 注册/登录路由
	r.GET("/", hans.ShowRegister)
	r.POST("/register", hans.Register)

	r.GET("/login", hans.ShowLogin)
	r.POST("/login", hans.Login)
	r.GET("/logout", hans.Logout)

	admin := r.Group("/admin/user")
	admin.Use(middlewares.JWTAuth())
	{
		// 用户管理
		admin.GET("/", hans.UserListPage)
		admin.GET("/api", hans.UserList)
		// 删除用户
		admin.DELETE("/:id", hans.DeleteUser)

		// 新增
		admin.GET("/add", hans.AddUserPage)
		admin.POST("/add/api", hans.AddUserApi)

		// 页面
		admin.GET("/detail", hans.UserDetailPage)

		// JSON API
		admin.GET("/detail/api", hans.UserDetailApi)
		admin.PUT("/detail/api", hans.UpdateUserApi)

		// 搜索用户
		admin.GET("/search/api", hans.SearchUserapi)

		// 用户统计
		admin.GET("/total/api", hans.UserStat)

		// 状态
		admin.POST("/status", hans.ChangeUserStatus)
	}

	r.Run("0.0.0.0:8081")
}
