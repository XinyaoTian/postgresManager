package main

import (
	"fmt"
	"github.com/XinyaoTian/postgresManager/comment"
)

func main()  {
	//dbOp, err := dbOperation.Init("postgres","dbuser","docker","174.137.53.55","5432","testdb", "disable")

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

	c := comment.Init(23,"Lisa","Good","datetime")
	c.Show()
	fmt.Println(c.GetId())

}


