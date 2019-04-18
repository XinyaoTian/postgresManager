package dbOperation

import (
	"database/sql"
	"errors"
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

// golang 中的构造函数写法, 不与相应的 struct 绑定!
//func Init(protocol, dbUsername, dbPassword, dbIp, dbPort, dbName, sslMode string) (*dbOperator, error)  {
	//connStr := protocol + `://` + dbUsername + `:` + dbPassword + `@` + dbIp + `:` + dbPort + `/` + dbName + `?sslmode=` +sslMode
	//// 连接数据库
	//db, err := sql.Open("postgres", connStr)
	//// 验证完是否可以连接后，关闭数据库
	//defer db.Close()
	//if err != nil {
	//	return new(dbOperator{}), err
	//} else {
	//	return &dbOperator{connStr:connStr}, err
	//}
//}

// golang 中的构造函数写法, 不与相应的 struct 绑定!
func Init(protocol, dbUsername, dbPassword, dbIp, dbPort, dbName, sslMode string) *dbOperator  {
	connStr := protocol + `://` + dbUsername + `:` + dbPassword + `@` + dbIp + `:` + dbPort + `/` + dbName + `?sslmode=` +sslMode
	return &dbOperator{connStr:connStr}
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

// 向表中插入数据
func (dbOp *dbOperator) insertIntoPost(username, post, createtime string) error  {
	// 插入相关的 sql 语句
	insertSQL := `INSERT INTO post(username,post,createtime) VALUES($1,$2,$3) RETURNING id `

	// 连接数据库
	db, errOpen := sql.Open("postgres", dbOp.connStr)
	// 准备及执行相关数据库操作
	stmt, errStmt := db.Prepare(insertSQL)
	_, errRes := stmt.Exec(username, post, createtime)

	// 纠错，如果没有发生任何错误则返回 nil， 否则输出错误
	if (_checkErr(errOpen) && _checkErr(errStmt) && _checkErr(errRes)) == true {
		return nil
	} else {
		log.Fatal("Can not insert data=(" + username + "," + post + "," + createtime + ")")
		return errors.New("Can not insert data=(" + username + "," + post + "," + createtime + ")")
	}
}


