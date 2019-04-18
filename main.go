package main

import (
	"fmt"
	"./pkg"
	"os"
	
)


func main()  {
	dbOp := dbOperation.Init("postgres","dbuser","docker","174.137.53.55","5432","testdb", "disable")

	err = dbOp.insertIntoPost("Mike","Hello","2012-02-02")
	if err != nil {
		fmt.Println(err)
		os.Exit(255)
	}
}
