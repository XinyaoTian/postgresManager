package comment

import (
	"fmt"
	"strconv"
)

type Comment struct {
	id int64
	username string
	post string
	postTime string
}

func Init(id int64, username, post, postTime string) *Comment {
	return &Comment{id:id, username:username, post:post, postTime:postTime}
}

func (comment Comment) GetId() int64  {
	return comment.id
}

func (comment Comment) GetUsername() string  {
	return comment.username
}

func (comment Comment) GetPost() string  {
	return comment.GetPost()
}

func (comment Comment) GetPostTime() string  {
	return comment.postTime
}

func (comment Comment) Show() {
	fmt.Println("id = " + strconv.FormatInt(comment.id,10))
	fmt.Println("username = " + comment.username)
	fmt.Println("post = " + comment.post)
	fmt.Println("postTime = " + comment.postTime)
}

func (comment *Comment) SetUsername(username string) {
	comment.username = username
}

func (comment *Comment) SetPost(post string) {
	comment.post = post
}

func (comment *Comment) SetPostTime(postTime string) {
	comment.postTime = postTime
}

func (comment *Comment) SetId(id int64) {
	comment.id = id
}




