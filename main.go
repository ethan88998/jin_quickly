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

	admin := r.Group("/admin")
	admin.Use(middlewares.JWTAuth())
	{
		// 用户管理
		admin.GET("/user", hans.UserListPage)
		admin.GET("/user/api", hans.UserList)
		// 删除用户
		admin.DELETE("/user/:id", hans.DeleteUser)
		// 编辑用户
		admin.GET("/user/detail", hans.UserDetailPage)
		admin.GET("/user/detail/api", hans.UserDetailApi)
		// 更新数据
		admin.PUT("/user/detail/api", hans.UpdateUserApi)
		// 搜索用户
		admin.GET("/user/search/api", hans.SearchUserapi)

		// 用户统计
		admin.GET("/user/total/api", hans.UserStat)
	}

	r.Run(":8081")
}
