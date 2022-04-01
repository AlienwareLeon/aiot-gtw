package main

import (
	"aiot_gtw/src/main/tools"
	"database/sql"
	"fmt"
	"log"
	"time"
)




func main() {
	//// 连接数据库
	//
	//// 1.创建路由
	//r := gin.Default()
	//// 2.绑定路由规则，执行的函数
	//// gin.Context，封装了request和response
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})
	//r.GET("/xxxpost", func(context *gin.Context) {
	//	context.String(http.StatusCreated,fmt.Sprint(service.GetTypeAndProductInfo()))
	//})
	//r.PUT("/xxxput")
	//
	//fmt.Println("uuid:" + util.GetUUID())
	//
	//// 3.监听端口，默认在8080
	//// Run("里面不指定端口号默认为8080")
	//r.Run(":8000")
	//fmt.Println(service.GetTypeAndProductInfo())
	//service.GetAppCache()
	//service.GetTypeAndProductInfo()
	for true {
		//1S
		time.Sleep(time.Duration(10)*time.Second)
	}
}

func redisConnectionDev() {
	db,_:=sql.Open("mysql","root:root@(127.0.0.1:3306)/golang")
	err :=db.Ping()
	if err != nil{
		fmt.Println("数据库链接失败")
	}
	defer db.Close()
}

func init()  {
	tools.SetUp()
	tools.SetUpConnection()
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds)
}


