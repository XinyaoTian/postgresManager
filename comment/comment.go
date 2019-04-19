package comment

import (
	"fmt"
	"strconv"
)

type Comment struct {
	id int64
	Username string
	Post string
	PostTime string
	RecordId string
}

func Init(id int64, username, post, postTime, recordId string) *Comment {
	return &Comment{id:id, Username:username, Post:post, PostTime:postTime, RecordId:recordId}
}

func (comment Comment) GetId() int64  {
	return comment.id
}

func (comment Comment) GetRecordId() string  {
	return comment.RecordId
}

func (comment Comment) GetUsername() string  {
	return comment.Username
}

func (comment Comment) GetPost() string  {
	return comment.Post
}

func (comment Comment) GetPostTime() string  {
	return comment.PostTime
}

func (comment Comment) Show() {
	fmt.Println("id = " + strconv.FormatInt(comment.id,10))
	fmt.Println("username = " + comment.Username)
	fmt.Println("post = " + comment.Post)
	fmt.Println("postTime = " + comment.PostTime)
	fmt.Print("record_id = " + comment.RecordId)
}

func (comment *Comment) SetUsername(username string) {
	comment.Username = username
}

func (comment *Comment) SetPost(post string) {
	comment.Post = post
}

func (comment *Comment) SetPostTime(postTime string) {
	comment.PostTime = postTime
}

func (comment *Comment) SetId(id int64) {
	comment.id = id
}

func (comment *Comment) SetRecordId(recordId string) {
	comment.RecordId = recordId
}




