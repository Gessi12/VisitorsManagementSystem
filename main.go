package main

import (
	"VisitorsManagementSystem/dao"
	"VisitorsManagementSystem/routers"
)

func main() {
	//获取初始化的数据库
	dao.InitDB()

	router := routers.Router()
	//在9090端口启动服务
	router.Run(":9090")
}
