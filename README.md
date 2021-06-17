![P-Blog标志](https://plutowu-blogimgs.oss-cn-guangzhou.aliyuncs.com/img/P.png "P-Blog logo")

[![Language](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![Open Source](https://forthebadge.com/images/badges/open-source.svg)](https://forthebadge.com)
[![Keybord](https://forthebadge.com/images/badges/powered-by-jeffs-keyboard.svg)](https://forthebadge.com)

[View English](./README.md)

<!-- TOC -->

- [P-Blog介绍](#p-blog介绍)
- [安装](#安装)
- [技术栈](#技术栈)
- [许可证](#许可证)

<!-- /TOC -->

## P-Blog介绍

P-Blog是一个使用Gin和Vue搭建的简洁博客。

P-Blog提供了一个简便的Gin框架学习项目，旨在帮助初学者更快的入手Golang的实际开发。

![P-Blog架构](https://plutowu-blogimgs.oss-cn-guangzhou.aliyuncs.com/img/P-Blog%E6%9E%B6%E6%9E%84.png)

## 安装

1. 克隆该项目。

    ```bash
    git clone https://gitee.com/p1utowu/p-blog.git
    ```

2. 安装相关依赖。

    ```bash
    go mod tidy
    ```

3. 配置项目设置文件config.ini

   ```ini
   [server]
   AppMode = debug # debug-开发模式，release-生产模式
   HttpPort = :3000 # 项目端口
   JwtKey = 89js82js72 # JWT密钥，随机字符串即可
   
   [database]
   Db = mysql # 数据库类型，不能变更为其他形式
   DbHost = 127.0.0.1 # 数据库地址
   DbPort = 3306 # 数据库端口
   DbUser = root # 数据库用户名
   DbPassWord = root # 数据库用户密码
   DbName = p_blog # 数据库名
   
   [qiniu]
   # 七牛储存信息
   AccessKey = # AK
   SecretKey = # SK
   Bucket = 
   QiniuSever =
   ```

4. 运行SQL中的SQL文件

5. 启动项目

   ```bash
   go run main.go
   ```

6. 访问项目

   ```
   前台页面
   http://localhost:3000
   后台管理页面
   http://localhost:3000/admin
       
   默认管理员:admin  密码:123456
   ```

## 技术栈

- 后端
  - Gin web framework
  - GORM
  - Jwt-Go
  - Scrypt
  - Logrus
  - Go-INI
- 前端
  - Vue
  - Vue Router
  - Ant Design
  - Axios
  - Tinymce
  - Moment

## 许可证

Apache License 2.0