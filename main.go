package main

import (
	"github.com/XinyaoTian/postgresManager/config"
	"github.com/XinyaoTian/postgresManager/routeHandler"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)


func main()  {
	//dbOp, _ := dbOperation.Init("postgres","dbuser","docker","174.137.53.55","5432","testdb", "disable")
	//
	//commentsSlice, _ := dbOp.QueryFromPostByUsername("Mike")
	//for i:=0; i<len(commentsSlice); i++ {
	//	if commentsSlice[i] != nil {
	//		commentsSlice[i].Show()
	//	}
	//}

	//err = dbOp.InsertIntoPost("Mike","Hello","2012-02-02")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(255)
	//}
	//
	//commentsSlice, _ := dbOp.QueryAllFromPost()
	//for i:=0; i<len(commentsSlice); i++ {
	//	if commentsSlice[i] != nil {
	//		commentsSlice[i].Show()
	//	}
	//}

	// 加载 postgresManager/config 中的数据库配置项
	// 输出配置项到日志中，便于容器内部的运维
	var dbConfig = config.Configinfo
	log.Println("Postgres DB manager start!")
	log.Println("POSTGRES_PROTOCOL=", dbConfig.GetDbProtocol())
	log.Println("POSTGRES_IPADDR=", dbConfig.GetDbIp())
	log.Println("POSTGRES_PORT=", dbConfig.GetDbPort())
	log.Println("POSTGRES_USERNAME=", dbConfig.GetUsername())
	log.Println("POSTGRES_PASSWORD=", dbConfig.GetDbPassword())
	log.Println("POSTGRES_DBNAME=", dbConfig.GetDbName())
	log.Println("POSTGRES_SSLMODE=", dbConfig.GetSslMode())
	log.Println("GOPATH=", os.Getenv("GOPATH"))

	router := httprouter.New()

	router.GET("/", routeHandler.Manual)
	router.GET("/all-queries", routeHandler.GetAllPosts)
	router.GET("/queries/:username", routeHandler.GetPostsByUsername)
	// curl 的 post 命令
	// curl -X POST -d 'username=SAM&post=Hello&postTime=2012-02-05' http://127.0.0.1:9090/insert-post/00ac
	router.POST("/insert-post/:recordId", routeHandler.InsertNewPost)

	log.Fatal(http.ListenAndServe(":9090", router))


	//routeHandler.Manual()

}


