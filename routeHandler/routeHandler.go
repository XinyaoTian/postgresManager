package routeHandler

import (
	"encoding/json"
	"fmt"
	"github.com/XinyaoTian/postgresManager/config"
	"github.com/XinyaoTian/postgresManager/dbOperation"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// 纠错函数
func _checkErr(err error) bool {
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		return true
	}
}

func Manual(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	// 用法的结构体
	type usage struct {
		Url string
		Description string
	}
	// 用于输出 json 的结构体
	type manual struct {
		Welcome string
		Usages []usage
	}

	// 用法文档
	usage_allQueries := usage{Url:"/all-queries", Description:"Get all queries from post db."}
	usage_insertPost := usage{Url:"/insert-post/:recordId", Description:"Insert a new post into db."}
	usage_queryByName := usage{Url:"/queries/:username", Description:"Get queries by username."}

	// 如需添加新的 API 用法说明，请按照上面的格式在下面继续添加
	// 并将其载入 usagesSlice 中
	// ...

	usagesSlice := []usage{usage_allQueries, usage_insertPost, usage_queryByName}

	m := manual{Welcome:"Welcome to use postgres DB management API.", Usages:usagesSlice}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal("json err:", err)
		fmt.Fprintf(w,`["Info":"Page Error."]`)
	} else {
		fmt.Fprintf(w, string(b))
	}
}

func GetAllPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 加载 postgresManager/config 中的数据库配置项
	var dbConfig = config.Configinfo
	dbOp, errOpen := dbOperation.Init(
		dbConfig.GetDbProtocol(), dbConfig.GetUsername(),
		dbConfig.GetDbPassword(), dbConfig.GetDbIp(),
		dbConfig.GetDbPort(), dbConfig.GetDbName(),
		dbConfig.GetSslMode(),
		)
	// 执行 postgresManager/dbOperation 中的数据库操作
	comments, errQuery := dbOp.QueryAllFromPost()
	b, errMar := json.Marshal(comments)
	// 若所有步骤都没有报错
	if _checkErr(errOpen) && _checkErr(errMar) && _checkErr(errQuery) == true {
		// 则输出 json 格式信息
		fmt.Fprintf(w, string(b))
	} else {
		// 否则报错 记录错误日志
		log.Fatal(errQuery)
		log.Fatal(errOpen)
		log.Fatal(errMar)
		fmt.Fprintf(w,`["Info":"Page Error."]`)
	}
}

func GetPostsByUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	// 载入由 heeprouter 从 url 中获取到的参数信息
	username := ps.ByName("username")
	// 加载 postgresManager/config 中的数据库配置项
	var dbConfig = config.Configinfo
	dbOp, errOpen := dbOperation.Init(
		dbConfig.GetDbProtocol(), dbConfig.GetUsername(),
		dbConfig.GetDbPassword(), dbConfig.GetDbIp(),
		dbConfig.GetDbPort(), dbConfig.GetDbName(),
		dbConfig.GetSslMode(),
	)
	// 执行 postgresManager/dbOperation 中的数据库操作
	comments, errQuery := dbOp.QueryFromPostByUsername(username)
	b, errMar := json.Marshal(comments)
	// 若所有步骤都没有报错
	if _checkErr(errOpen) && _checkErr(errMar) && _checkErr(errQuery) == true {
		// 则输出 json 格式信息
		fmt.Fprintf(w, string(b))
	} else {
		// 否则报错 记录错误日志
		log.Fatal(errQuery)
		log.Fatal(errOpen)
		log.Fatal(errMar)
		fmt.Fprintf(w,`["Info":"Page Error."]`)
	}
}

func InsertNewPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	// 获取必要的信息
	// 从 url 传递的使用 ps.ByName("name");
	recordId := ps.ByName("recordId")
	// 从表单来的数据使用 r.FormValue("name")
	username := r.FormValue("username")
	post := r.FormValue("post")
	postTime := r.FormValue("postTime")

	// 加载 postgresManager/config 中的数据库配置项
	var dbConfig = config.Configinfo
	dbOp, errOpen := dbOperation.Init(
		dbConfig.GetDbProtocol(), dbConfig.GetUsername(),
		dbConfig.GetDbPassword(), dbConfig.GetDbIp(),
		dbConfig.GetDbPort(), dbConfig.GetDbName(),
		dbConfig.GetSslMode(),
	)
	// 执行 postgresManager/dbOperation 中的数据库操作
	errInsert := dbOp.InsertIntoPost(username,post,postTime, recordId)
	if _checkErr(errInsert) && _checkErr(errOpen) == true {
		fmt.Fprintf(w, `["Info":"Insert Success.","Error":"false"]` )
	} else {
		fmt.Fprintf(w, `["Info":"Insert failed.","Error":"true"]`)
	}
}


