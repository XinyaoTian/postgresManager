package dbOperation

import (
	"database/sql"
	"errors"
	"github.com/XinyaoTian/postgresManager/comment"
	_ "github.com/lib/pq"
	"log"
)

// DB table 示例
// tableName: post
// id int64
// username string
// post string
// createtime string

// golang 中的类使用 struct 结构来表示
type dbOperator struct {
	connStr string  // 通过 Init 函数组合出的 connStr 用于连接 Postgres
}

// 纠错函数
func _checkErr(err error) bool {
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		return true
	}
}

// golang 中的构造函数写法, 不与相应的 struct 绑定!
func Init(protocol, dbUsername, dbPassword, dbIp, dbPort, dbName, sslMode string) (*dbOperator, error)  {
	connStr := protocol + `://` + dbUsername + `:` + dbPassword + `@` + dbIp + `:` + dbPort + `/` + dbName + `?sslmode=` +sslMode
	// 连接数据库
	db, err := sql.Open("postgres", connStr)
	// 验证完是否可以连接后，关闭数据库
	defer db.Close()
	if err != nil {
		return &dbOperator{""}, err
	} else {
		return &dbOperator{connStr:connStr}, err
	}
}

// 向表中插入数据, 注意对外暴露的方法名称, 首字母要大写. 不能写成 insertIntoPost
func (dbOp *dbOperator) InsertIntoPost(username, post, createtime, recordId string) error  {
	// 插入相关的 sql 语句
	insertSQL := `INSERT INTO post(username,post,createtime,record_id) VALUES($1,$2,$3,$4) RETURNING id `

	// 连接数据库
	db, errOpen := sql.Open("postgres", dbOp.connStr)
	// 函数执行完毕后记得关闭 db 链接，否则会导致服务器内存溢出
	defer db.Close()

	// 准备及执行相关数据库操作
	stmt, errStmt := db.Prepare(insertSQL)
	_, errRes := stmt.Exec(username, post, createtime, recordId)

	// 纠错，如果没有发生任何错误则返回 nil， 否则输出错误
	if (_checkErr(errOpen) && _checkErr(errStmt) && _checkErr(errRes)) == true {
		return nil
	} else {
		log.Fatal("Can not insert data=(" + username + "," + post + "," + createtime + ")")
		return errors.New("Can not insert data=(" + username + "," + post + "," + createtime + ")")
	}
}

// 输出全部记录
func (dbOp *dbOperator) QueryAllFromPost() ([]*comment.Comment, error) {
	QueryStr := `SELECT * FROM post`

	// 初始化用于存储 comment 的 slice
	var commentsSlice []*comment.Comment

	// 连接数据库
	db, errOpen := sql.Open("postgres", dbOp.connStr)
	// 函数执行完毕后记得关闭 db 链接，否则会导致服务器内存溢出
	defer db.Close()
	// 查询
	rows, errQuery := db.Query(QueryStr)
	// 纠错
	if (_checkErr(errOpen) || _checkErr(errQuery) ) == false {
		log.Fatal("Can not query.")
		return commentsSlice, errors.New("can not query")
	}

	// 循环将查到的信息加入 commentsSlice
	for rows.Next() {
		var id int64
		var username string
		var post string
		var postTime string
		var recordId string
		var c *comment.Comment
		_ = rows.Scan(&id, &username, &post, &postTime, &recordId)
		// 将扫描出的数据初始化为 comment
		c = comment.Init(id, username, post, postTime, recordId)
		// 加入 slice
		commentsSlice = append(commentsSlice, c)
	}
	return commentsSlice, nil
}

// 根据用户名查询数据
func (dbOp *dbOperator) QueryFromPostByUsername(username string) ([]*comment.Comment, error) {
	QueryStr := `SELECT * FROM post WHERE username=$1`

	// 初始化用于存储 comment 的 slice
	var commentsSlice []*comment.Comment

	// 连接数据库
	db, errOpen := sql.Open("postgres", dbOp.connStr)
	// 函数执行完毕后记得关闭 db 链接，否则会导致服务器内存溢出
	defer db.Close()
	// 查询
	rows, errQuery := db.Query(QueryStr,username)
	// 纠错
	if (_checkErr(errOpen) || _checkErr(errQuery) ) == false {
		log.Fatal("Can not query.")
		return commentsSlice, errors.New("can not query")
	}

	// 循环将查到的信息加入 commentsSlice
	for rows.Next() {
		var id int64
		var username string
		var post string
		var postTime string
		var recordId string
		var c *comment.Comment
		_ = rows.Scan(&id, &username, &post, &postTime, &recordId)
		// 将扫描出的数据初始化为 comment
		c = comment.Init(id, username, post, postTime, recordId)
		// 加入 slice
		commentsSlice = append(commentsSlice, c)
	}
	return commentsSlice, nil
}

//// 根据数据库记录查询数据
//func (dbOp *dbOperator) QueryFromPostById(username, post, createtime string)  {
//
//}
//

//
//// 根据内容查询数据
//func (dbOp *dbOperator) QueryFromPostById(username, post, createtime string)  {
//
//}
//
//// 根据创建时间查询数据
//func (dbOp *dbOperator) QueryFromPostById(username, post, createtime string)  {
//
//}

