package main

import (
	"BeegoProject0603/db_mysql"
	_ "BeegoProject0603/routers"
	"github.com/astaxie/beego"
)

func main() {
	//1、连接数据库
	db_mysql.Connect()

	//2、其他配置

	//3、启动应用
	beego.Run()//代码简洁
}

