package main

import (
	"beegolaoshi/db_mysql"
	_ "beegolaoshi/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//1、连接数据库
	db_mysql.Connect()

	//2、其他配置
fmt.Println("hi")
	//3、启动应用
	beego.Run()//代码简洁
}

