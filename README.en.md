![P-Blog标志](https://plutowu-blogimgs.oss-cn-guangzhou.aliyuncs.com/img/P.png "P-Blog logo")

[![Language](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![Open Source](https://forthebadge.com/images/badges/open-source.svg)](https://forthebadge.com)
[![Keybord](https://forthebadge.com/images/badges/powered-by-jeffs-keyboard.svg)](https://forthebadge.com)

<!-- TOC -->

- [What is P-Blog](#What is P-Blog)
- [Installation](#Installation)
- [Technology stack](#Technology stack)
- [Licence](#Licence)

<!-- /TOC -->

## What is P-Blog

P-Blog is a simple blog built with Gin and Vue.

P-blog provides a simple learning project of Gin framework, which aims to help beginners start the actual development of Golang faster.

![P-Blog架构](https://plutowu-blogimgs.oss-cn-guangzhou.aliyuncs.com/img/P-Blog%E6%9E%B6%E6%9E%84.png)

## Installation

1. Clone the project.

   ```bash
   git clone https://gitee.com/p1utowu/p-blog.git
   ```

2. Install related dependencies.

   ```bash
   go mod tidy
   ```

3. Configuration project settings file config.ini

   ```ini
   [server]
   AppMode = debug # debug-development mode, release-production mode
   HttpPort = :3000 # Project port
   JwtKey = 89js82js72 # JWT key, random string
   
   [database]
   Db = mysql # Database type, cannot be changed to another form
   DbHost = 127.0.0.1 # Database address
   DbPort = 3306 # Database port
   DbUser = root # Database user
   DbPassWord = root # Database password
   DbName = p_blog # Database name
   
   [qiniu]
   # Qiniu store information
   AccessKey = # AK
   SecretKey = # SK
   Bucket = 
   QiniuSever =
   ```

4. Run the SQL file in the SQL folder.

5. Start project

   ```bash
   go run main.go
   ```

6. Visit project

   ```
   Front page
   http://localhost:3000
   
   Background management page
   http://localhost:3000/admin
       
   Default administrator:admin  Password:123456
   ```

## Technology stack

- Golang
  - Gin web framework
  - GORM
  - Jwt-Go
  - Scrypt
  - Logrus
  - Go-INI
- JavaScript
  - Vue
  - Vue Router
  - Ant Design
  - Axios
  - Tinymce
  - Moment

## Licence

Apache License 2.0