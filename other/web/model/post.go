package model

type Post struct {
	Model
	Title      string
	Content    string
	PostCateId int
}

func NewPost() *Post {
	return &Post{}
}
