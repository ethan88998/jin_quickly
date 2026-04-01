
好的，我帮你整理一份 **完整、系统的 Docker 部署流程**，专门针对你这种 **Go + Gin + MySQL 用户管理项目**，目标是：

* 本地或内网可访问
* 容器化部署
* 数据库初始化可用
* 可长期维护

下面一步步来。

---

# **一、准备阶段**

## 1️⃣ 确认环境

* Docker 安装：`docker -v`
* Docker Compose 安装：`docker-compose -v`
* Go 版本 ≥ 项目要求（例如 1.25+）

---

## 2️⃣ 项目目录结构（推荐）

```
jin_quickly/
├── main.go
├── go.mod
├── go.sum
├── Dockerfile
├── docker-compose.yml
├── static/
├── templates/
└── models/
```

* `Dockerfile`：Go 容器编译 + 运行
* `docker-compose.yml`：容器编排（MySQL + Go）

---

# **二、Dockerfile（Go 容器）**

📄 `Dockerfile`

```dockerfile
FROM golang:1.25-alpine

WORKDIR /app

# 先复制 go.mod 下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 编译 Go 项目
RUN go build -o app

# 容器端口
EXPOSE 8081

# 启动程序
CMD ["./app"]
```

---

# **三、docker-compose.yml**

📄 `docker-compose.yml`

```yaml
services:
  mysql:
    image: mysql:8.0
    container_name: jin_mysql
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: root123
      MYSQL_DATABASE: jin_quickly   # 初始化数据库
    volumes:
      - mysql_data:/var/lib/mysql
    ports:
      - "127.0.0.1:3306:3306"     # 本机可访问，但外网不可访问

  app:
    build: .
    container_name: jin_app
    restart: always
    depends_on:
      - mysql
    ports:
      - "0.0.0.0:8081:8081"       # 内网可访问
    environment:
      DB_USER: root
      DB_PASSWORD: root123
      DB_HOST: jin_mysql
      DB_PORT: 3306
      DB_NAME: jin_quickly

volumes:
  mysql_data:
```

⚠️ 关键点：

* `DB_HOST` 用 **容器名**，不要用 localhost
* MySQL 容器密码必须和 Go 容器一致
* Go 容器端口暴露 0.0.0.0 → 局域网可访问

---

# **四、Go 数据库连接修改**

在 Go 项目里：

```go
dsn := fmt.Sprintf(
  "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
  os.Getenv("DB_USER"),
  os.Getenv("DB_PASSWORD"),
  os.Getenv("DB_HOST"),
  os.Getenv("DB_PORT"),
  os.Getenv("DB_NAME"),
)
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
```

✅ 使用环境变量，避免写死密码

---

# **五、构建并启动**

```bash
docker-compose up -d --build
```

检查状态：

```bash
docker ps
docker logs -f jin_app
```

* Gin 启动日志
* 数据库连接成功
* 端口监听正常

---

# **六、访问**

* 本机访问：

```
http://localhost:8081
```

* 局域网访问：

```
http://192.168.x.x:8081
```

> 你的电脑 IP 用 `ifconfig` 或 `ipconfig` 查看

---

# **七、常用运维命令**

```bash
# 查看容器日志
docker logs -f jin_app

# 停止容器
docker-compose down

# 重新构建
docker-compose up -d --build

# 重启容器
docker-compose restart
```

---

# **八、可选优化**

1. **内网安全**

    * Docker 容器端口只暴露给局域网
    * MySQL 只绑定 127.0.0.1

2. **Nginx 反向代理**

    * 内网访问统一用 80 端口
    * 避免暴露 8081

3. **数据备份**

    * MySQL 数据存 volume
    * 定期备份 `mysql_data`

4. **初始化数据**

    * `docker-entrypoint-initdb.d/` 放 SQL 文件，第一次容器启动自动执行

---

💡 **总结**

* Docker + Compose 可以完整部署你的 Go + MySQL 项目
* 内网访问安全
* 数据库初始化可控
* Go 与数据库连接通过环境变量管理
* 容器日志可直接查看，方便排错

---

如果你需要，我可以帮你直接写一份 **“完整可运行的 docker-compose + Go + MySQL 内网部署模板”**，只要拷贝就能用。

你希望我直接生成吗？
